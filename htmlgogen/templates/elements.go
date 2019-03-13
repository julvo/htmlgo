package htmlgo

import (
    "fmt"
    "strings"
    "html"
    a "github.com/julvo/htmlgo/attributes"
)

type HTML string

func insertAttributes(attrs []a.Attribute) string {
    s := ""
    for _, a := range attrs {
        s += " " + string(a)
    }
    return s
}

func insertChildren(children []HTML) string {
    s := ""
    for _, c := range children {
        s += string(c)
    }
    return s
}

func indent(s, indentation string) string {
    return strings.Replace(s, "\n", "\n" + indentation, -1)
}

func Element(tag string, attrs []a.Attribute, children ...HTML) HTML {
    return HTML(
        "\n<" + tag + insertAttributes(attrs) + ">" +
        indent(insertChildren(children), "  ") +
        "\n</" + tag + ">")
}

func VoidElement(tag string, attrs []a.Attribute) HTML {
    return HTML(
        "\n<" + tag + insertAttributes(attrs) + ">")
}

func Text(v interface{}) HTML {
    return HTML("\n" + html.EscapeString(fmt.Sprint(v)))
}

func Doctype(t string) HTML {
    return HTML("<!DOCTYPE " + t + ">")
}
const DoctypeHtml5 HTML = "<!DOCTYPE HTML>"

{{ range .ElementFuncs }}
func {{.FuncName}}(attrs []a.Attribute, children ...HTML) HTML {
    return Element("{{.TagName}}", attrs, children...)
}
{{ end }}

{{ range .VoidElementFuncs }}
func {{.FuncName}}(attrs []a.Attribute) HTML {
    return VoidElement("{{.TagName}}", attrs)
}
{{ end }}
