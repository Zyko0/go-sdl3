package main

import (
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
	Name     string
	Receiver *FuncArg
	Params   []*FuncArg
	Return   *FuncArg
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("err: ", err)
	}
	path = filepath.Join(path, "sdl_functions.gen.go")

	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Fatal("cannot parse "+path+": ", err)
	}

	funcsByReceiver := map[string][]*Function{}
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
						name := ts.Names[0].Name[1:]  // Remove "i" prefix
						switch tt := ts.Type.(type) { //ts.Values[0].(type) {
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
										default:
											raw := fmt.Sprintf("%v", xt)
											typ = "*" + raw
										}
									case *ast.SliceExpr:
										typ = fmt.Sprintf("[]%v", ft.X)
									}
									fn.Return = &FuncArg{
										Type: typ,
									}
								}
							}
							if len(tt.Params.List) > 0 {
								for i, f := range tt.Params.List {
									typ := fmt.Sprintf("%v", f.Type)
									switch ft := f.Type.(type) {
									case *ast.StarExpr:
										switch xt := ft.X.(type) {
										case *ast.StarExpr:
											typ = fmt.Sprintf("**%v", xt.X)
										default:
											raw := fmt.Sprintf("%v", xt)
											typ = "*" + raw
											if i == 0 && raw[0] >= 'A' && raw[0] <= 'Z' {
												fn.Receiver = &FuncArg{
													Name: f.Names[0].Name,
													Type: typ,
												}
											}
										}
									case *ast.SliceExpr:
										typ = fmt.Sprintf("[]%v", ft.X)
									default:
										if i == 0 && typ[0] >= 'A' && typ[0] <= 'Z' {
											fn.Receiver = &FuncArg{
												Name: f.Names[0].Name,
												Type: typ,
											}
										}
									}
									if i != 0 || fn.Receiver == nil {
										fn.Params = append(fn.Params, &FuncArg{
											Name: f.Names[0].Name,
											Type: typ,
										})
									}
								}
								fmt.Println()
							}
							if fn.Receiver != nil {
								funcsByReceiver[fn.Receiver.Type] = append(funcsByReceiver[fn.Receiver.Type], fn)
							}
						}
					}
				}
			}
		}
		return true
	})

	count := 0
	for _, funcs := range funcsByReceiver {
		count += len(funcs)
	}
	fmt.Println("count:", len(funcsByReceiver), "total methods:", count)

	f := jen.NewFile("sdl")
	for receiver, funcs := range funcsByReceiver {
		receiver = strings.ReplaceAll(receiver, "*", "")
		f.Comment("// " + receiver).Line()
		for _, fn := range funcs {
			newFnName := strings.Replace(fn.Name, receiver, "", 1)
			if newFnName != "Get" && strings.HasPrefix(newFnName, "Get") {
				newFnName = newFnName[3:]
			}
			f.Func().
				Params(
					jen.Id(fn.Receiver.Name).Id(fn.Receiver.Type),
				).
				Id(newFnName).
				ParamsFunc(func(g *jen.Group) {
					for _, arg := range fn.Params {
						g.Add(jen.Id(arg.Name).Id(arg.Type))
					}
				}).
				ParamsFunc(func(g *jen.Group) {
					if fn.Return != nil {
						g.Add(jen.Id(fn.Return.Name).Id(fn.Return.Type))
					}
				}).
				BlockFunc(func(g *jen.Group) {
					g.Add(jen.Panic(jen.Lit("not implemented")))
					call := jen.Id("i" + fn.Name).CallFunc(func(g *jen.Group) {
						g.Add(jen.Id(fn.Receiver.Name))
						for _, arg := range fn.Params {
							g.Add(jen.Id(arg.Name))
						}
					})
					if fn.Return != nil {
						g.Add(jen.Return(call))
					} else {
						g.Add(call)
					}
				}).
				Line()
		}
	}

	err = os.WriteFile("methods_generated_please_rename.go", []byte(f.GoString()), 0666)
	if err != nil {
		log.Fatal("err: ", err)
	}
}
