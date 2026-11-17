// Command jsgen updates <library>_functions_js.go in place.
//
// A function whose body starts with panic("not implemented on js") is a stub
// and gets regenerated; any other body is hand-written and kept byte for
// byte. Functions missing from the file are appended.
//
// A stub becomes a working binding only when every argument and the result
// map to wasm32 without any knowledge of what the function does: scalars,
// strings and pointers to SDL-allocated objects. Other stubs keep panicking,
// with TODO markers on the arguments that need a decision.
//
// A generated working body has no panic, so later runs keep it like a
// hand-written one: restore the file from git to regenerate it.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

const stubPanic = "not implemented on js"

// sdlAllocated lists non-opaque structs that only SDL allocates, so that a
// Go pointer to one always comes from internal.NewObject.
var sdlAllocated = map[string]bool{
	"sdl.Surface":   true,
	"sdl.Palette":   true,
	"sdl.Texture":   true,
	"ttf.Text":      true,
	"img.Animation": true,
}

type kind int

const (
	kindUnsupported kind = iota
	kindBool
	kindInt32
	kindInt64
	kindUint64
	kindFloat32
	kindFloat64
	kindString
	kindHandle
	// kindPointer is a pointer to anything but a handle: an input, an
	// output or an array.
	kindPointer
	// kindUintptr is a size_t, a void*, a buffer or a callback.
	kindUintptr
)

type param struct {
	Name string
	Type string
	Kind kind
}

type function struct {
	Name   string
	Symbol string
	Params []param
	Result *param
	// Manual is set for signatures written by hand in *_functions_ex.go:
	// they encode the native ABI, e.g. SDL_Color passed by value as a
	// uint32, which does not hold on wasm32.
	Manual bool
}

func (f *function) mechanical() bool {
	if f.Manual {
		return false
	}
	for _, p := range f.Params {
		if !p.Kind.mechanical() {
			return false
		}
	}
	return f.Result == nil || f.Result.Kind.mechanical()
}

// destroys reports whether f frees the object passed first, whose Go
// stand-in must then leave the registry.
func (f *function) destroys() bool {
	return (strings.HasPrefix(f.Name, "iDestroy") || strings.HasPrefix(f.Name, "iClose")) &&
		len(f.Params) > 0 && f.Params[0].Kind == kindHandle
}

func (k kind) mechanical() bool {
	return k != kindUnsupported && k != kindPointer && k != kindUintptr
}

func main() {
	var libraryName, dir string
	flag.StringVar(&dir, "dir", "", "package directory to update (e.g: sdl)")
	flag.StringVar(&libraryName, "library", "", "library name (e.g: sdl)")
	flag.Parse()

	prefix, ok := map[string]string{
		"sdl":   "SDL_",
		"ttf":   "TTF_",
		"mixer": "MIX_",
		"img":   "IMG_",
	}[libraryName]
	if !ok {
		log.Fatalf("unknown library %q", libraryName)
	}

	res := &resolver{
		local: libraryName,
		pkgs: map[string]map[string]ast.Expr{
			libraryName: loadTypes(dir),
			"sdl":       loadTypes(filepath.Join(dir, "..", "sdl")),
		},
	}

	var funcs []*function
	for _, name := range []string{"_functions.gen.go", "_functions_ex.go"} {
		path := filepath.Join(dir, libraryName+name)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		funcs = append(funcs, loadFunctions(path, prefix, res, name == "_functions_ex.go")...)
	}

	path := filepath.Join(dir, libraryName+"_functions_js.go")
	src, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	out, err := merge(path, src, funcs)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(path, out, 0666); err != nil {
		log.Fatal(err)
	}
}

func loadTypes(dir string) map[string]ast.Expr {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	typs := map[string]ast.Expr{}
	fset := token.NewFileSet()
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(dir, e.Name()), nil, parser.SkipObjectResolution)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, s := range gd.Specs {
				ts := s.(*ast.TypeSpec)
				typs[ts.Name.Name] = ts.Type
			}
		}
	}
	return typs
}

type resolver struct {
	local string
	pkgs  map[string]map[string]ast.Expr
}

var builtinKinds = map[string]kind{
	"bool":    kindBool,
	"int8":    kindInt32,
	"int16":   kindInt32,
	"int32":   kindInt32,
	"uint8":   kindInt32,
	"uint16":  kindInt32,
	"uint32":  kindInt32,
	"byte":    kindInt32,
	"rune":    kindInt32,
	"int":     kindInt32, // C int
	"uint":    kindInt32, // C unsigned int
	"int64":   kindInt64,
	"uint64":  kindUint64,
	"float32": kindFloat32,
	"float64": kindFloat64,
	"string":  kindString,
	"uintptr": kindUintptr,
}

func (r *resolver) kind(expr ast.Expr, pkg string) kind {
	switch e := expr.(type) {
	case *ast.StarExpr:
		if r.isHandle(e.X, pkg) {
			return kindHandle
		}
		return kindPointer
	case *ast.Ident:
		def := r.pkgs[pkg][e.Name]
		if def == nil {
			if k, ok := builtinKinds[e.Name]; ok {
				return k
			}
			return kindUnsupported
		}
		// A named pointer type (GLContext, GUID) does not unify with the
		// *T parameters of the internal helpers.
		if _, ok := def.(*ast.StarExpr); ok {
			return kindUnsupported
		}
		return r.kind(def, pkg)
	case *ast.SelectorExpr:
		x, ok := e.X.(*ast.Ident)
		if !ok || r.pkgs[x.Name] == nil {
			return kindUnsupported
		}
		return r.kind(e.Sel, x.Name)
	default:
		return kindUnsupported
	}
}

func (r *resolver) isHandle(expr ast.Expr, pkg string) bool {
	var name string
	switch e := expr.(type) {
	case *ast.Ident:
		name = e.Name
	case *ast.SelectorExpr:
		x, ok := e.X.(*ast.Ident)
		if !ok {
			return false
		}
		pkg, name = x.Name, e.Sel.Name
	default:
		return false
	}
	if sdlAllocated[pkg+"."+name] {
		return true
	}
	st, ok := r.pkgs[pkg][name].(*ast.StructType)
	return ok && len(st.Fields.List) == 0
}

func loadFunctions(path, prefix string, res *resolver, manual bool) []*function {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	var funcs []*function
	for _, d := range f.Decls {
		gd, ok := d.(*ast.GenDecl)
		if !ok || gd.Tok != token.VAR {
			continue
		}
		for _, s := range gd.Specs {
			vs := s.(*ast.ValueSpec)
			ft, ok := vs.Type.(*ast.FuncType)
			if !ok {
				continue
			}
			fn := &function{
				Name:   vs.Names[0].Name,
				Symbol: prefix + strings.TrimPrefix(vs.Names[0].Name, "i"),
				Manual: manual,
			}
			if vs.Doc != nil {
				for _, c := range vs.Doc.List {
					if _, sym, ok := strings.Cut(c.Text, "//puregogen:function symbol="); ok {
						fn.Symbol = strings.Fields(sym)[0]
					}
				}
			}
			for _, field := range ft.Params.List {
				for _, n := range field.Names {
					fn.Params = append(fn.Params, param{
						Name: n.Name,
						Type: types.ExprString(field.Type),
						Kind: res.kind(field.Type, res.local),
					})
				}
			}
			if ft.Results != nil && len(ft.Results.List) > 0 {
				typ := ft.Results.List[0].Type
				fn.Result = &param{
					Type: types.ExprString(typ),
					Kind: res.kind(typ, res.local),
				}
			}
			funcs = append(funcs, fn)
		}
	}
	return funcs
}

func isStub(lit *ast.FuncLit) bool {
	if len(lit.Body.List) == 0 {
		return false
	}
	es, ok := lit.Body.List[0].(*ast.ExprStmt)
	if !ok {
		return false
	}
	call, ok := es.X.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return false
	}
	if id, ok := call.Fun.(*ast.Ident); !ok || id.Name != "panic" {
		return false
	}
	lit2, ok := call.Args[0].(*ast.BasicLit)
	return ok && lit2.Value == strconv.Quote(stubPanic)
}

type edit struct {
	start, end int
	text       string
}

func merge(path string, src []byte, funcs []*function) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	var init *ast.FuncDecl
	for _, d := range file.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok && fd.Name.Name == "initialize" && fd.Recv == nil {
			init = fd
		}
	}
	if init == nil {
		return nil, fmt.Errorf("%s: no initialize function", path)
	}
	off := func(p token.Pos) int { return fset.Position(p).Offset }

	byName := map[string]*function{}
	for _, f := range funcs {
		byName[f.Name] = f
	}
	var (
		edits                                 []edit
		seen                                  = map[string]bool{}
		kept, working, stubs, added, orphaned int
	)
	count := func(f *function) {
		if f.mechanical() {
			working++
		} else {
			stubs++
		}
	}
	for _, stmt := range init.Body.List {
		as, ok := stmt.(*ast.AssignStmt)
		if !ok || len(as.Lhs) != 1 || len(as.Rhs) != 1 {
			continue
		}
		id, ok1 := as.Lhs[0].(*ast.Ident)
		lit, ok2 := as.Rhs[0].(*ast.FuncLit)
		if !ok1 || !ok2 {
			continue
		}
		f := byName[id.Name]
		if f == nil {
			orphaned++
			log.Printf("%s: %s has no declaration, kept as is", path, id.Name)
			continue
		}
		seen[f.Name] = true
		if !isStub(lit) {
			kept++
			continue
		}
		count(f)
		edits = append(edits, edit{off(stmt.Pos()), off(stmt.End()), generate(f)})
	}
	var appended strings.Builder
	for _, f := range funcs {
		if seen[f.Name] {
			continue
		}
		added++
		count(f)
		appended.WriteString("\n" + generate(f) + "\n")
	}
	if appended.Len() > 0 {
		pos := off(init.Body.Rbrace)
		edits = append(edits, edit{pos, pos, appended.String()})
	}

	out := applyEdits(src, edits)
	out, err = fixImports(path, out)
	if err != nil {
		return nil, err
	}
	formatted, err := format.Source(out)
	if err != nil {
		return nil, fmt.Errorf("%s: format: %w", path, err)
	}
	log.Printf("%s: %d hand-written kept, %d generated (%d new): %d working, %d stubs, %d orphans",
		path, kept, working+stubs, added, working, stubs, orphaned)
	return formatted, nil
}

func applyEdits(src []byte, edits []edit) []byte {
	slices.SortFunc(edits, func(a, b edit) int { return b.start - a.start })
	out := slices.Clone(src)
	for _, e := range edits {
		out = slices.Concat(out[:e.start], []byte(e.text), out[e.end:])
	}
	return out
}

// generatedImports are the packages generated bodies may refer to.
var generatedImports = map[string]string{
	"js":       `js "syscall/js"`,
	"unsafe":   `"unsafe"`,
	"internal": `internal "github.com/Zyko0/go-sdl3/internal"`,
}

// fixImports drops imports that regenerated bodies stopped using, e.g.
// unsafe once no stub calls unsafe.Sizeof, and adds the ones they need.
func fixImports(path string, src []byte) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	used := map[string]bool{}
	ast.Inspect(file, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				used[id.Name] = true
			}
		}
		return true
	})
	var (
		edits    []edit
		imported = map[string]bool{}
	)
	for _, imp := range file.Imports {
		p, _ := strconv.Unquote(imp.Path.Value)
		name := p[strings.LastIndex(p, "/")+1:]
		if imp.Name != nil {
			name = imp.Name.Name
		}
		imported[name] = true
		if name == "_" || used[name] {
			continue
		}
		start := fset.Position(imp.Pos()).Offset
		end := fset.Position(imp.End()).Offset
		start = bytes.LastIndexByte(src[:start], '\n') + 1
		if i := bytes.IndexByte(src[end:], '\n'); i >= 0 {
			end += i + 1
		}
		edits = append(edits, edit{start, end, ""})
	}
	var missing []string
	for name, spec := range generatedImports {
		if used[name] && !imported[name] {
			missing = append(missing, "\t"+spec+"\n")
		}
	}
	if len(missing) > 0 {
		gd, ok := file.Decls[0].(*ast.GenDecl)
		if !ok || gd.Tok != token.IMPORT || !gd.Rparen.IsValid() {
			return nil, fmt.Errorf("%s: expected a parenthesized import block first", path)
		}
		slices.Sort(missing)
		pos := fset.Position(gd.Rparen).Offset
		edits = append(edits, edit{pos, pos, strings.Join(missing, "")})
	}
	return applyEdits(src, edits), nil
}

func convert(to, typ, expr string) string {
	if typ == to {
		return expr
	}
	return to + "(" + expr + ")"
}

func generate(f *function) string {
	var b strings.Builder
	w := func(format string, args ...any) { fmt.Fprintf(&b, format+"\n", args...) }

	params := make([]string, len(f.Params))
	for i, p := range f.Params {
		params[i] = p.Name + " " + p.Type
	}
	result := ""
	if f.Result != nil {
		result = " " + f.Result.Type
	}
	w("%s = func(%s)%s {", f.Name, strings.Join(params, ", "), result)
	if !f.mechanical() {
		w("panic(%q)", stubPanic)
		if f.Manual {
			w("// TODO: this signature follows the native ABI, check the C prototype.")
		}
	}

	needsStack := slices.ContainsFunc(f.Params, func(p param) bool {
		return p.Kind == kindString || p.Kind == kindPointer
	})
	if needsStack {
		w("internal.StackSave()")
		w("defer internal.StackRestore()")
	}
	for _, p := range f.Params {
		v := "_" + p.Name
		switch p.Kind {
		case kindBool:
			w("%s := internal.NewBoolean(%s)", v, convert("bool", p.Type, p.Name))
		case kindInt32:
			w("%s := %s", v, convert("int32", p.Type, p.Name))
		case kindInt64:
			w("%s := internal.NewBigInt(%s)", v, convert("int64", p.Type, p.Name))
		case kindUint64:
			w("%s := internal.NewBigInt(%s)", v, convert("uint64", p.Type, p.Name))
		case kindFloat32:
			w("%s := %s", v, convert("float32", p.Type, p.Name))
		case kindFloat64:
			w("%s := %s", v, convert("float64", p.Type, p.Name))
		case kindString:
			w("%s := internal.StringOnJSStack(%s)", v, convert("string", p.Type, p.Name))
		case kindHandle:
			w("%s, _ := internal.GetJSPointer(%s)", v, p.Name)
		case kindPointer:
			w("%s := %s // TODO: input, output or array?", v, p.Name)
		case kindUintptr:
			w("%s := %s // TODO: size, buffer, userdata or callback?", v, p.Name)
		case kindUnsupported:
			w("%s := %s // TODO: unsupported type", v, p.Name)
		}
	}

	call := "js.Global().Get(\"Module\").Call("
	if f.Result != nil {
		call = "ret := " + call
	}
	w("%s", call)
	w("%q,", "_"+f.Symbol)
	for _, p := range f.Params {
		w("_%s,", p.Name)
	}
	w(")")
	if f.destroys() {
		w("internal.DeleteJSPointer(uintptr(unsafe.Pointer(%s)))", f.Params[0].Name)
	}

	if r := f.Result; r != nil {
		switch r.Kind {
		case kindBool:
			w("return %s", convert(r.Type, "bool", "internal.GetBool(ret)"))
		case kindInt32:
			w("return %s(ret.Int())", r.Type)
		case kindInt64:
			w("return %s", convert(r.Type, "int64", "internal.GetInt64(ret)"))
		case kindUint64:
			w("return %s", convert(r.Type, "uint64", "internal.GetUint64(ret)"))
		case kindFloat32, kindFloat64:
			w("return %s(ret.Float())", r.Type)
		case kindString:
			w("return %s", convert(r.Type, "string", "internal.UTF8JSToString(ret)"))
		case kindHandle:
			w("return internal.NewObject[%s](ret)", r.Type[1:])
		case kindPointer:
			w("return internal.NewObject[%s](ret) // TODO: single value or array?", r.Type[1:])
		case kindUintptr:
			w("return %s // TODO: size, or pointer through internal.NewUintptr?",
				convert(r.Type, "uintptr", "uintptr(uint32(ret.Int()))"))
		case kindUnsupported:
			w("_ = ret")
			w("return *new(%s) // TODO: unsupported type", r.Type)
		}
	}
	w("}")
	return strings.TrimSuffix(b.String(), "\n")
}
