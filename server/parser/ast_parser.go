package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// APIEndpoint represents a parsed HTTP endpoint from source code
type APIEndpoint struct {
	Method      string
	Path        string
	Handler     string
	Description string
	Package     string
}

// ParseGoDirectory walks a directory and extracts API endpoint definitions
func ParseGoDirectory(dir string) ([]APIEndpoint, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var endpoints []APIEndpoint
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			endpoints = append(endpoints, extractEndpoints(file, pkg.Name)...)
		}
	}
	return endpoints, nil
}

func extractEndpoints(file *ast.File, pkgName string) []APIEndpoint {
	var endpoints []APIEndpoint

	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		method := sel.Sel.Name
		if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
			return true
		}

		if len(call.Args) >= 2 {
			pathLit, ok := call.Args[0].(*ast.BasicLit)
			if ok {
				endpoints = append(endpoints, APIEndpoint{
					Method:  strings.ToUpper(method),
					Path:    strings.Trim(pathLit.Value, "\""),
					Package: pkgName,
				})
			}
		}
		return true
	})

	return endpoints
}
