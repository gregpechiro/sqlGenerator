package sqlGenerator

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ernesto-jimenez/gogen/cleanimports"
	"github.com/ernesto-jimenez/gogen/gogenutil"
	"github.com/ernesto-jimenez/gogen/importer"
	"github.com/ernesto-jimenez/gogen/imports"
	"github.com/gregpechiro/structFields"
)

// SqlGenerator will generate the UnmarshalMap function
type SqlGenerator struct {
	name       string
	targetName string
	namePkg    string
	pkg        *types.Package
	target     *types.Struct
}

// NewSqlGenerator initializes a SqlGenerator
func NewSqlGenerator(pkg, target string) (*SqlGenerator, error) {
	var err error
	if pkg == "" || pkg[0] == '.' {
		pkg, err = filepath.Abs(filepath.Clean(pkg))
		if err != nil {
			return nil, err
		}
		pkg = gogenutil.StripGopath(pkg)
	}
	p, err := importer.Default().Import(pkg)
	if err != nil {
		return nil, err
	}
	obj := p.Scope().Lookup(target)
	if obj == nil {
		return nil, fmt.Errorf("struct %s missing", target)
	}
	if _, ok := obj.Type().Underlying().(*types.Struct); !ok {
		return nil, fmt.Errorf("%s should be an struct, was %s", target, obj.Type().Underlying())
	}
	return &SqlGenerator{
		targetName: target,
		pkg:        p,
		target:     obj.Type().Underlying().(*types.Struct),
	}, nil
}

func (g SqlGenerator) Fields() []structFields.Field {
	numFields := g.target.NumFields()

	fields := make([]structFields.Field, 0)
	for i := 0; i < numFields; i++ {
		f := structFields.Field{&g, g.target.Tag(i), g.target.Field(i)}
		if f.Field() != "" {
			fields = append(fields, f)
		}
	}
	return fields
}

func (g SqlGenerator) Qf(pkg *types.Package) string {
	if g.pkg == pkg {
		return ""
	}
	return pkg.Name()
}

func (g SqlGenerator) Name() string {
	name := g.targetName
	return name
}

func (g SqlGenerator) Package() string {
	if g.namePkg != "" {
		return g.namePkg
	}
	return g.pkg.Name()
}

func (g *SqlGenerator) SetPackage(name string) {
	g.namePkg = name
}

func (g SqlGenerator) Imports() map[string]string {
	imports := imports.New(g.Package())
	fields := g.Fields()
	for i := 0; i < len(fields); i++ {
		m := fields[i]
		imports.AddImportsFrom(m.V.Type())
		imports.AddImportsFrom(m.UnderlyingType())
		if sub := m.UnderlyingTarget(); sub != nil {
			fields = append(fields, sub.Fields()...)
		}
	}
	return imports.Imports()
}

func (g SqlGenerator) QuestionMarks() string {
	var q []string
	for i := 0; i < len(g.Fields()); i++ {
		q = append(q, "?")
	}
	return strings.Join(q, ", ")
}

/*func (g SqlGenerator) CreateTable() string {
	var c string
	for _, f := range g.Fields() {
		c += ToSnake(f.Name())
		tag := reflect.StructTag(f.Tag)
		t := tag.Get("sqltyp")
		if t == "" {
			t = GetSqlType(f.Type())
		}
		c += " " + t
	}
}*/

func (g SqlGenerator) Write(wr io.Writer) error {
	var buf bytes.Buffer
	if err := fnTmpl.Execute(&buf, g); err != nil {
		return err
	}
	return cleanimports.Clean(wr, buf.Bytes())
}

var (
	tempFuncs = template.FuncMap{
		"toLowerFirst": toLowerFirst,
		"sqlName":      ToSnake,
	}

	fnTmpl = template.Must(template.New("func").Funcs(tempFuncs).Parse(`/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/sqlGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package {{ .Package }}

import (
"fmt"
"log"
{{ range $path, $name := .Imports }}
{{ $name }} "{{ $path }}"{{ end }}
)
{{ $fields := .Fields }}{{ $strct := .Name }}
func Create{{ $strct }}sTable() {
   _, err := db.Exec("CREATE TABLE {{ sqlName $strct }}()")
   if err != nil {
	   fmt.Printf("Create{{ $strct }}sTable >> %v\n", err)
   }
}

func Add{{ $strct }}({{ toLowerFirst $strct }} {{ $strct }}) {
	_, err := db.Exec("INSERT INTO user VALUES( {{ .QuestionMarks }} )", {{ range $fields }}{{ toLowerFirst $strct }}.{{ .Field }},{{ end }})
	if err != nil {
		fmt.Printf("Add{{ $strct }} >> %v\n", err)
	}
}

func GetAll{{ $strct }}s() []{{ $strct }} {
	rows, err := db.Query("SELECT * FROM {{ sqlName $strct }}")
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var {{ toLowerFirst $strct }}s []{{ $strct }}
	for rows.Next() {
		var {{ toLowerFirst $strct }} {{ $strct }}
		err = rows.Scan({{ range $fields }}&{{ toLowerFirst $strct }}.{{ .Field }},{{ end }})
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		{{ toLowerFirst $strct }}s = append({{ toLowerFirst $strct }}s, {{ toLowerFirst $strct }})
	}
	return {{ toLowerFirst $strct }}s
}
{{ range .Fields }}{{ if .UnderlyingIsBasic }}
func Get{{ $strct }}By{{ .Name }}({{ toLowerFirst .Name }} {{ .Type }}) {{ $strct }} {
	{{ toLowerFirst $strct }} := {{ $strct }}{}
	err := db.QueryRow("SELECT * FROM {{ sqlName $strct }} WHERE {{ sqlName .Name }}=?", {{ toLowerFirst .Name }}).Scan({{ range $fields }}&{{ toLowerFirst $strct }}.{{ .Field }},{{ end }})
	if err != nil {
		fmt.Printf("Get{{ .Name }}ById >> %v\n", err)
	}
	return {{ toLowerFirst $strct }}
}

func GetAll{{ $strct }}sBy{{ .Name }}({{ toLowerFirst .Name }} {{ .Type }}) []{{ $strct }} {
	rows, err := db.Query("SELECT * FROM {{ sqlName $strct }} WHERE {{ sqlName .Name }}=?", {{ toLowerFirst .Name }})
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var {{ toLowerFirst $strct }}s []{{ $strct }}
	for rows.Next() {
		var {{ toLowerFirst $strct }} {{ $strct }}
		err = rows.Scan({{ range $fields }}&{{ toLowerFirst $strct }}.{{ .Field }},{{ end }})
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		{{ toLowerFirst $strct }}s = append({{ toLowerFirst $strct }}s, {{ toLowerFirst $strct }})
	}
	return {{ toLowerFirst $strct }}s
}
{{ end }}{{ end }}
`))

//
// 	fnTmpl = template.Must(template.New("func").Parse(`/*
// * CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/unmarshalmap
// * THIS FILE SHOULD NOT BE EDITED BY HAND
// */
//
// package {{.Package}}
//
// import (
// 	"fmt"
// {{range $path, $name := .Imports}}
// 	{{$name}} "{{$path}}"{{end}}
// )
//
// {{define "UNMARSHALFIELDS"}}
// {{range .Fields}}
// {{if .IsAnonymous}}
// 	// Anonymous {{.Name}}
// 	if scoped := true; scoped {
// 		var s *{{.Type}} = &s.{{.Name}}
// 		// Fill object
// 		{{template "UNMARSHALFIELDS" .UnderlyingTarget}}
// 	}
// {{else if .IsArrayOrSlice}}
// 	// ArrayOrSlice {{.Name}}
// 	{{if .UnderlyingIsBasic}}
// 	if v, ok := m["{{.Field}}"].([]{{.UnderlyingType}}); ok {
// 		{{if .IsSlice}}
// 		s.{{.Name}} = make({{.Type}}, len(v))
// 		{{else}}
// 		if len(s.{{.Name}}) < len(v) {
// 			return fmt.Errorf("expected field {{.Field}} to be an array with %d elements, but got an array with %d", len(s.{{.Name}}), len(v))
// 		}
// 		{{end}}
// 		for i, el := range v {
// 			s.{{.Name}}[i] = el
// 		}
// 	} else if v, ok := m["{{.Field}}"].([]interface{}); ok {
// 		{{if .IsSlice}}
// 		s.{{.Name}} = make({{.Type}}, len(v))
// 		{{else}}
// 		if len(s.{{.Name}}) < len(v) {
// 			return fmt.Errorf("expected field {{.Field}} to be an array with %d elements, but got an array with %d", len(s.{{.Name}}), len(v))
// 		}
// 		{{end}}
// 		for i, el := range v {
// 			if v, ok := el.({{.UnderlyingType}}); ok {
// 				s.{{.Name}}[i] = v
// 			{{if .UnderlyingConvertibleFromFloat64}}
// 			} else if m, ok := el.(float64); ok {
// 				v := {{.UnderlyingType}}(m)
// 				s.{{.Name}} = v
// 			{{end}}
// 			} else {
// 				return fmt.Errorf("expected field {{.Field}}[%d] to be {{.UnderlyingType}} but got %T", i, el)
// 			}
// 		}
// 	} else if v, exists := m["{{.Field}}"]; exists && v != nil {
// 		return fmt.Errorf("expected field {{.Field}} to be []{{.UnderlyingType}} but got %T", m["{{.Field}}"])
// 	}
// 	{{else}}
// 	if v, ok := m["{{.Field}}"].([]interface{}); ok {
// 		{{if .IsSlice}}
// 		s.{{.Name}} = make({{.Type}}, len(v))
// 		{{else}}
// 		if len(s.{{.Name}}) < len(v) {
// 			return fmt.Errorf("expected field {{.Field}} to be an array with %d elements, but got an array with %d", len(s.{{.Name}}), len(v))
// 		}
// 		{{end}}
// 		prev := s
// 		for i, el := range v {
// 			var s *{{.UnderlyingTypeName}}
// 			{{if .UnderlyingIsPointer}}
// 			if el == nil {
// 				continue
// 			}
// 			prev.{{.Name}}[i] = &{{.UnderlyingTypeName}}{}
// 			s = prev.{{.Name}}[i]
// 			{{else}}
// 			s = &prev.{{.Name}}[i]
// 			{{end}}
// 			if m, ok := el.(map[string]interface{}); ok {
// 				// Fill object
// 				{{template "UNMARSHALFIELDS" .UnderlyingTarget}}
// 			}
// 		}
// 	} else if v, exists := m["{{.Field}}"]; exists && v != nil {
// 		return fmt.Errorf("expected field {{.Field}} to be []interface{} but got %T", m["{{.Field}}"])
// 	}
// 	{{end}}
// {{else if .IsPointer}}
// 	// Pointer {{.Name}}
// 	if p, ok := m["{{.Field}}"]; ok {
// 		{{if .UnderlyingIsBasic}}
// 		if m, ok := p.({{.UnderlyingType}}); ok {
// 			s.{{.Name}} = &m
// 		{{if .UnderlyingConvertibleFromFloat64}}
// 		} else if m, ok := p.(float64); ok {
// 			v := {{.UnderlyingType}}(m)
// 			s.{{.Name}} = &v
// 		{{end}}
// 		} else if p == nil {
// 			s.{{.Name}} = nil
// 		}
// 		{{else}}
// 		if m, ok := p.(map[string]interface{}); ok {
// 			if s.{{.Name}} == nil {
// 				s.{{.Name}} = &{{.UnderlyingTypeName}}{}
// 			}
// 			s := s.{{.Name}}
// 			{{template "UNMARSHALFIELDS" .UnderlyingTarget}}
// 		} else if p == nil {
// 			s.{{.Name}} = nil
// 		} else {
// 			return fmt.Errorf("expected field {{.Field}} to be map[string]interface{} but got %T", p)
// 		}
// 		{{end}}
// 	}
// {{else if .IsStruct}}
// 	// Struct {{.Name}}
// 	if m, ok := m["{{.Field}}"].(map[string]interface{}); ok {
// 		var s *{{.Type}} = &s.{{.Name}}
// 		// Fill object
// 		{{template "UNMARSHALFIELDS" .UnderlyingTarget}}
// 	} else if v, exists := m["{{.Field}}"]; exists && v != nil {
// 		return fmt.Errorf("expected field {{.Field}} to be map[string]interface{} but got %T", m["{{.Field}}"])
// 	}
// {{else}}
// 	if v, ok := m["{{.Field}}"].({{.Type}}); ok {
// 		s.{{.Name}} = v
// 	{{if .ConvertibleFromFloat64}}
// 	} else if p, ok := m["{{.Field}}"].(float64); ok {
// 		v := {{.Type}}(p)
// 		s.{{.Name}} = v
// 	{{end}}
// 	} else if v, exists := m["{{.Field}}"]; exists && v != nil {
// 		return fmt.Errorf("expected field {{.Field}} to be {{.Type}} but got %T", m["{{.Field}}"])
// 	}
// {{end}}
// {{end}}
// {{end}}
//
// // UnmarshalMap takes a map and unmarshals the fieds into the struct
// func (s *{{.Name}}) UnmarshalMap(m map[string]interface{}) error {
// 	{{template "UNMARSHALFIELDS" .}}
// 	return nil
// }
// `))

)
