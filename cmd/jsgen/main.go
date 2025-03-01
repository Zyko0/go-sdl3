package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

type FuncArg struct {
	Name string
	Type string
}

type Function struct {
	Name   string
	Params []*FuncArg
	Return *FuncArg
}

const (
	internalQual = "github.com/Zyko0/go-sdl3/internal"
)

func main() {
	var (
		libraryName string
		dir         string
	)

	flag.StringVar(&dir, "dir", "", "base directory to generate from/to")
	flag.StringVar(&libraryName, "library", "", "library name (e.g: sdl)")
	flag.Parse()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("err: ", err)
	}
	path = filepath.Join(path, dir, libraryName+"_functions.gen.go")

	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Fatal("cannot parse "+path+": ", err)
	}

	var funcs []*Function
	ast.Inspect(root, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch tn := n.(type) {
		case *ast.GenDecl:
			if tn.Tok == token.VAR {
				for _, s := range tn.Specs {
					switch ts := s.(type) {
					case *ast.ValueSpec:
						name := ts.Names[0].Name
						switch tt := ts.Type.(type) {
						case *ast.FuncType:
							fn := &Function{
								Name: name,
							}
							if tt.Results != nil && len(tt.Results.List) > 0 {
								for _, f := range tt.Results.List {
									typ := fmt.Sprintf("%v", f.Type)
									switch ft := f.Type.(type) {
									case *ast.StarExpr:
										switch xt := ft.X.(type) {
										case *ast.StarExpr:
											typ = fmt.Sprintf("**%v", xt.X)
										case *ast.SelectorExpr:
											typ = fmt.Sprintf("*%v.%v", xt.X, xt.Sel)
										default:
											raw := fmt.Sprintf("%v", xt)
											typ = "*" + raw
										}
									case *ast.SliceExpr:
										typ = fmt.Sprintf("[]%v", ft.X)
									case *ast.SelectorExpr:
										typ = fmt.Sprintf("%v.%v", ft.X, ft.Sel)
									}
									fn.Return = &FuncArg{
										Type: typ,
									}
								}
							}
							if len(tt.Params.List) > 0 {
								for _, f := range tt.Params.List {
									typ := fmt.Sprintf("%v", f.Type)
									switch ft := f.Type.(type) {
									case *ast.StarExpr:
										switch xt := ft.X.(type) {
										case *ast.StarExpr:
											typ = fmt.Sprintf("**%v", xt.X)
										case *ast.SelectorExpr:
											typ = fmt.Sprintf("*%v.%v", xt.X, xt.Sel)
										default:
											raw := fmt.Sprintf("%v", xt)
											typ = "*" + raw
										}
									case *ast.SliceExpr:
										typ = fmt.Sprintf("[]%v", ft.X)
									case *ast.SelectorExpr:
										typ = fmt.Sprintf("*%v.%v", ft.X, ft.Sel)
									default:
									}
									fn.Params = append(fn.Params, &FuncArg{
										Name: f.Names[0].Name,
										Type: typ,
									})
								}
							}
							funcs = append(funcs, fn)
						}
					}
				}
			}
		}
		return true
	})

	f := jen.NewFile(libraryName)
	f.HeaderComment("//go:build js")
	f.Func().Id("init").Params().BlockFunc(func(g *jen.Group) {
		for _, fn := range funcs {
			g.Id(fn.Name).Op("=").Func().
				ParamsFunc(func(h *jen.Group) {
					for _, arg := range fn.Params {
						h.Add(jen.Id(arg.Name).Id(arg.Type))
					}
				}).
				ParamsFunc(func(h *jen.Group) {
					if fn.Return != nil {
						h.Add(jen.Id(fn.Return.Name).Id(fn.Return.Type))
					}
				}).
				BlockFunc(func(h *jen.Group) {
					h.Panic(jen.Lit("not implemented on js"))
					h.Add(jen.Qual(internalQual, "StackSave").Call())
					h.Add(jen.Defer().Qual(internalQual, "StackRestore").Call())
					for _, arg := range fn.Params {
						switch {
						case strings.HasPrefix(arg.Type, "*"):
							h.Add(
								jen.Id("_"+arg.Name).Op(",").Id("ok").Op(":=").
									Qual(internalQual, "GetJSPointer").
									Call(jen.Id(arg.Name)),
							)
							h.Add(jen.If(jen.Op("!").Id("ok")).BlockFunc(func(i *jen.Group) {
								var n jen.Code
								switch {
								case strings.HasPrefix(arg.Type, "**"):
									n = jen.Lit(4) // Size of int32
								case arg.Type == "*string":

								default:
									n = jen.Int().Parens(jen.Qual("unsafe", "Sizeof").Call(
										jen.Id("*" + arg.Name),
									))
								}
								i.Id("_"+arg.Name).Op("=").
									Qual(internalQual, "StackAlloc").
									Call(n)
							}))
						default:
							switch arg.Type {
							case "bool":
								h.Add(
									jen.Id("_"+arg.Name).Op(":=").
										Qual(internalQual, "NewBoolean").
										Call(jen.Id(arg.Name)),
								)
							case "int64", "uint64", "uintptr":
								h.Add(
									jen.Id("_"+arg.Name).Op(":=").
										Qual(internalQual, "NewBigInt").
										Call(jen.Id(arg.Name)),
								)
							case "string":
								h.Add(
									jen.Id("_"+arg.Name).Op(":=").
										Qual(internalQual, "StringOnStackGoToJS").
										Call(jen.Id(arg.Name)),
								)
							default:
								h.Add(
									jen.Id("_" + arg.Name).Op(":=").Int32().Parens(
										jen.Id(arg.Name),
									),
								)
							}
						}
					}
					// Function call
					call := jen.Null()
					if fn.Return != nil {
						call.Id("ret").Op(":=")
					}
					call.Qual("syscall/js", "Global").Call().Op(".").Id("Get").
						Call(jen.Lit("Module")).Op(".").Id("Call").
						CallFunc(func(i *jen.Group) {
							i.Add(jen.Line().Lit("_SDL_" + fn.Name[1:]))
							for _, a := range fn.Params {
								i.Add(jen.Line().Id("_" + a.Name))
							}
							i.Line()
						})
					h.Add(call)
					// Return
					if fn.Return != nil {
						h.Line()
						switch fn.Return.Type {
						case "bool":
							h.Return(jen.Qual(internalQual, "GetBool").
								Call(jen.Id("ret")),
							)
						case "int", "int8", "int16", "int32", "byte",
							"uint", "uint8", "uint16", "uint32":
							h.Return(jen.Id(fn.Return.Type).Parens(
								jen.Id("ret").Op(".").Id("Int").Call()),
							)
						case "int64", "uint64", "uintptr":
							h.Return(jen.Id(fn.Return.Type).Parens(jen.Qual(internalQual, "GetInt64").
								Call(jen.Id("ret")),
							))
							// TODO: ^
						case "string":
							h.Return(jen.Qual(internalQual, "UTF8ToStringJSToGo").
								Call(jen.Id("ret")),
							)
						default:
							switch {
							case strings.HasPrefix(fn.Return.Type, "*"):
								h.Id("_obj").Op(":=").Qual(internalQual, "NewPointer").
									Types(jen.Id(strings.ReplaceAll(fn.Return.Type, "*", ""))).
									Call(jen.Id("ret"))
								h.Return(jen.Id("_obj"))
							// TODO: case it's a known type
							default:
								h.Return(jen.Id(fn.Return.Type).Parens(
									jen.Id("ret").Op(".").Id("Int").Call(),
								))
							}
						}
					}
				}).
				Line()
		}
	})

	err = os.WriteFile(libraryName+"_functions_js_please_rename.go", []byte(f.GoString()), 0666)
	if err != nil {
		log.Fatal("err: ", err)
	}
}
