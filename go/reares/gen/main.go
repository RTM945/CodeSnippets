package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type XMLPackage struct {
	XMLName xml.Name `xml:"package"`
	Name    string   `xml:"name,attr"`
	Msgs    []XMLMsg `xml:"msg"`
}

type TemplateData struct {
	Package string
	Name    string
	Type    int
	Vars    []struct {
		GoName string
		GoType string
	}
}

type XMLMsg struct {
	Name    string   `xml:"name,attr"`
	Type    int      `xml:"type,attr"`
	MaxSize int      `xml:"maxsize,attr"`
	Vars    []XMLVar `xml:"var"`
}

type XMLVar struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

const msgTemplate = `package {{.Package}}

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	shard "reares/cmd/go-netty"
	"reares/protobuf"
)

type {{.Name}} struct {
	*shard.Msg
	Processor
	{{range .Vars}}{{.GoName}} {{.GoType}}
{{end}}}

func Init{{.Name}}(msgProcessor Processor) *{{.Name}} {
	header := &shard.MsgHeader{}
	header.TypeId = {{.Type}}
	return &{{.Name}}{
		Msg:       shard.NewMsg(header),
		Processor: msgProcessor,
	}
}

func New{{.Name}}() *{{.Name}} {
	header := &shard.MsgHeader{}
	header.TypeId = {{.Type}}
	return &{{.Name}}{
		Msg: shard.NewMsg(header),
	}
}

func (msg *{{.Name}}) Decode(src *bytes.Buffer) error {
	tmp := &protobuf.{{.Name}}{}
	err := proto.Unmarshal(src.Bytes(), tmp)
	if err != nil {
		return err
	}
	{{range .Vars}}msg.{{.GoName}} = tmp.{{.GoName}}
{{end}}	return nil
}

func (msg *{{.Name}}) Encode(dst *bytes.Buffer) error {
	tmp := &protobuf.{{.Name}}{
		{{range .Vars}}{{.GoName}}: msg.{{.GoName}},
{{end}}	}
	data, err := proto.Marshal(tmp)
	if err != nil {
		return err
	}
	_, err = dst.Write(data)
	return err
}

func (msg *{{.Name}}) Dispatch() {
	msg.Process()
}

func (msg *{{.Name}}) Process() error {
	return msg.Process{{.Name}}(msg)
}
`

const processorTemplate = `package {{.Name}}

type Processor interface {
	{{range .Msgs}}Process{{.Name}}(msg *{{.Name}}) error
	{{end}}
}`

func main() {
	// Parse XML input
	xmlData := `<?xml version="1.0" encoding="UTF-8" ?>
<package name="client_switch">
	<msg name="RsaKeyExchange" type="1" maxsize="1024">
        <var name="key" type="binary"/>
    </msg>
    <msg name="KeyExchange" type="2" maxsize="1024">
        <var name="key" type="binary"/>
    </msg>
</package>`

	var pkg XMLPackage
	err := xml.Unmarshal([]byte(xmlData), &pkg)
	if err != nil {
		panic(err)
	}

	// Generate Go code for each message
	for _, msg := range pkg.Msgs {
		// Convert variables to Go types
		var goVars []struct {
			GoName string
			GoType string
		}
		for _, v := range msg.Vars {
			goType := convertType(v.Type)
			goName := strings.Title(v.Name)
			goVars = append(goVars, struct {
				GoName string
				GoType string
			}{
				GoName: goName,
				GoType: goType,
			})
		}

		// Create template data
		data := TemplateData{
			Package: pkg.Name, // 从XML中获取的package名
			Name:    msg.Name,
			Type:    msg.Type,
			Vars:    goVars,
		}

		// Generate message code
		tmpl := template.Must(template.New("msg").Parse(msgTemplate))
		fmt.Printf("// Generated message code for %s\n", msg.Name)
		tmpl.Execute(os.Stdout, data)
		fmt.Println()
	}

	// Generate processor interface
	tmpl := template.Must(template.New("processor").Parse(processorTemplate))
	fmt.Println("// Generated processor interface")
	tmpl.Execute(os.Stdout, pkg)
}

func convertType(xmlType string) string {
	switch xmlType {
	case "binary":
		return "[]byte"
	case "string":
		return "string"
	case "int":
		return "int32"
	case "long":
		return "int64"
	default:
		return "interface{}"
	}
}
