package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "../../proto/echo.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Error parsing Go file: %v", err)
	}
	hasLogicImport := false
	for _, imp := range file.Imports {
		if strings.Contains(imp.Path.Value, "reares/logic/echo") {
			hasLogicImport = true
			break
		}
	}
	if !hasLogicImport {
		file.Imports = append(file.Imports, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: `"reares/logic/echo"`,
			},
		})
	}
	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			log.Println(fn.Name.Name)
		}
		if ok && fn.Name.Name == "Process" && strings.HasPrefix(fn.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name, "Echo") {
			fn.Body.List = []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "echo"},
							Sel: &ast.Ident{Name: "ProcessEcho"},
						},
						Args: []ast.Expr{
							&ast.Ident{Name: "msg"},
						},
					},
				},
			}
		}
		return true
	})

	output, err := os.Create("../../proto_override/echo_override.go")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	if err := printer.Fprint(output, fset, file); err != nil {
		log.Fatalf("Error printing Go file: %v", err)
	}
	output.Close()
}
