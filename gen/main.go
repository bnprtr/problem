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
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the input Go file.")
	}
	inputFile := os.Args[1]

	// Parse the input file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, inputFile, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	// Determine the package name (from input file)
	packageName := extractPackageName(file)
	if packageName == "" {
		log.Fatal("Couldn't determine package name.")
	}

	// Output file
	outputFile := filepath.Join(filepath.Dir(inputFile), strings.Replace(inputFile, ".go", ".gen.go", 1))
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	// Write package declaration
	fmt.Fprintf(output, "package %s\n\n", packageName)
	fmt.Fprintln(output, `import "net/http"`)

	var typeNames []string

	// Traverse AST and find string types
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if ident, ok := typeSpec.Type.(*ast.Ident); ok && ident.Name == "string" {
						typeNames = append(typeNames, typeSpec.Name.Name)
						generateErrorMethod(output, typeSpec.Name.Name)
						generateStatusCodeMethod(output, typeSpec.Name.Name)
						generateNewMethod(output, typeSpec.Name.Name)
					}
				}
			}
		}
	}
}

func generateStatusCodeMethod(output *os.File, typeName string) {
	fmt.Fprintf(output, "func (e %s) StatusCode() int {\n", typeName)
	fmt.Fprintf(output, "    return http.%s\n", typeName)
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
}

// extractPackageName extracts the package name from the parsed Go file.
func extractPackageName(file *ast.File) string {
	return file.Name.Name
}

// generateErrorMethod creates an Error() string method for the identified string-based type.
func generateErrorMethod(output *os.File, typeName string) {
	fmt.Fprintf(output, "func (e %s) Error() string {\n", typeName)
	fmt.Fprintf(output, "    return string(e)\n")
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
}

// generateNewMethod creates an Error() string method for the identified string-based type.
func generateNewMethod(output *os.File, typeName string) {
	fmt.Fprintf(output, "func (e %s) New(message string) Error[%s] {\n", typeName, typeName)
	fmt.Fprintf(output, "    return New(e).WithMessage(message)\n")
	fmt.Fprintln(output, "}")
	fmt.Fprintln(output)
}
