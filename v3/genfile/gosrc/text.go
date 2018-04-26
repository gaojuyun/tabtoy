package gosrc

// 报错行号+3
const templateText = `// Generated by github.com/davyxu/tabtoy
// DO NOT EDIT!!
// Version: {{.Version}}
package {{.PackageName}}

{{range $sn, $objName := $.Symbols.EnumNames}}
// table: {{$.Symbols.ObjectAtTable $objName}}
type {{$objName}} int32
const (	{{range $fi,$field := $.Symbols.Fields $objName}}
	{{$objName}}_{{$field.FieldName}} = {{$field.DefaultValue}} // {{$field.Name}} {{end}}
)

var (
	{{$objName}}MapperValueByName = map[string]int32{ {{range $fi,$field := $.Symbols.Fields $objName}}
		"{{$field.FieldName}}": {{$field.DefaultValue}}, // {{$field.Name}} {{end}}
	}

	{{$objName}}MapperNameByValue = map[int32]string{ {{range $fi,$field := $.Symbols.Fields $objName}}
		 {{$field.DefaultValue}}: "{{$field.FieldName}}", // {{$field.Name}} {{end}}
	}
)
{{end}}

{{range $sn, $objName := $.Symbols.StructNames}}
// table: {{$.Symbols.ObjectAtTable $objName}}
type {{$objName}} struct{ {{range $fi,$field := $.Symbols.Fields $objName}}
	{{$field.FieldName}} {{$field.FieldType}} // {{$field.Name}} {{end}} 
}
{{end}}



`
