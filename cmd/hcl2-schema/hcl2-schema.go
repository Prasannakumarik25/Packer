package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

var (
	typeNames  = flag.String("type", "", "comma-separated list of type names; must be set")
	output     = flag.String("output", "", "output file name; default srcdir/<type>_hcl2.go")
	trimprefix = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of stringer:\n")
	fmt.Fprintf(os.Stderr, "\tstringer [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tstringer [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("hcl2-schema: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	types := strings.Split(*typeNames, ",")

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{os.Getenv("GOFILE")}
	}
	fname := args[0]
	outputPath := fname[:len(fname)-2] + "hcl2spec.go"

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("ReadFile: %+v", err)
		os.Exit(1)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, b, parser.ParseComments)
	if err != nil {
		fmt.Printf("ParseFile: %+v", err)
		os.Exit(1)
	}

	res := []StructDef{}

	for _, t := range types {
		for _, decl := range f.Decls {
			typeDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			typeSpec, ok := typeDecl.Specs[0].(*ast.TypeSpec)
			if !ok {
				continue
			}
			structDecl, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			if typeSpec.Name.String() != t {
				continue
			}
			sd := StructDef{StructName: t}
			fields := structDecl.Fields.List
			for _, field := range fields {
				if len(field.Names) == 0 {
					continue
				}
				fieldName := field.Names[0].Name
				fieldType := string(b[field.Type.Pos()-1 : field.Type.End()-1])
				if !unicode.IsUpper([]rune(fieldName)[0]) {
					continue
				}
				fd := FieldDef{Name: fieldName}

				switch fieldType {
				case "[]string":
					fd.Spec = fmt.Sprintf("%#v", &hcldec.AttrSpec{
						Name:     fieldName,
						Type:     cty.List(cty.String),
						Required: false,
					})
				case "[]byte", "string", "time.Duration":
					fd.Spec = fmt.Sprintf("%#v", &hcldec.AttrSpec{
						Name:     fieldName,
						Type:     cty.String,
						Required: false,
					})
					// fd.Type = "hcl2template.Type" + strings.Title(fieldType)
				case "int", "int32", "float":
					fd.Spec = fmt.Sprintf("%#v", &hcldec.AttrSpec{
						Name:     fieldName,
						Type:     cty.Number,
						Required: false,
					})
				case "bool", "config.Trilean":
					fd.Spec = fmt.Sprintf("%#v", &hcldec.AttrSpec{
						Name:     fieldName,
						Type:     cty.Bool,
						Required: false,
					})
				default:
					if strings.Contains(fieldType, "func") {
						continue
					}
					fd.Spec = `nil`
					// fd.Type = "hcl2template.TypeBlock"
					// fd.Elem = "(*" + fieldType + ").HCL2Schema(nil)"
				}

				sd.Fields = append(sd.Fields, fd)
			}
			if len(sd.Fields) == 0 {
				continue
			}
			res = append(res, sd)
		}
	}
	if len(res) == 0 {
		return
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = structDocsTemplate.Execute(outputFile, Output{
		Package:    f.Name.String(),
		StructDefs: res,
	})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}

type Output struct {
	Package    string
	StructDefs []StructDef
}

type FieldDef struct {
	Name string
	Spec string
}

type StructDef struct {
	StructName string
	Fields     []FieldDef
}

var structDocsTemplate = template.Must(template.New("structDocsTemplate").
	Funcs(template.FuncMap{
		// "indent": indent,
	}).
	Parse(`// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package {{ .Package }}

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)
{{ range .StructDefs }}
func (*{{ .StructName }}) HCL2Spec() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		{{- range .Fields}}
		"{{ .Name }}": {{ .Spec }},
		{{- end }}
	}
	return hcldec.ObjectSpec(s)
}
{{end}}`))
