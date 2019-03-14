package htmlgo

import (
    "fmt"
    "strings"
    "html"
    "html/template"
    "bytes"

    a "github.com/julvo/htmlgo/attributes"
)

type HTML string
type JS struct {
    templ       string
    data        interface{}
}

func insertAttributes(attrs []a.Attribute) string {
    s := ""
    for _, a := range attrs {
        s += " " + string(a)
    }
    return s
}

func insertChildren(children ...HTML) string {
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
        indent(insertChildren(children...), "  ") +
        "\n</" + tag + ">")
}

func VoidElement(tag string, attrs []a.Attribute) HTML {
    return HTML(
        "\n<" + tag + insertAttributes(attrs) + ">")
}

// Produce HTML from plain text by escaping
func Text(v interface{}) HTML {
    return HTML("\n" + html.EscapeString(fmt.Sprint(v)))
}

// Begin of manually defined elements

func Html5(attrs []a.Attribute, children ...HTML) HTML {
    return DoctypeHtml5 + Html(attrs, children...)
}

func Html5_(children ...HTML) HTML {
    return Html5(a.Attr(), children...)
}

func Doctype(t string) HTML {
    return HTML("<!DOCTYPE " + t + ">")
}

const DoctypeHtml5 HTML = "<!DOCTYPE HTML>"

func Script(attrs []a.Attribute, js JS) HTML {
    if js.data == nil {
        return Element("script", attrs, HTML("\n" + js.templ))
    }

    // TODO set verbosity level to enable logging
    t, err := template.New("_").Parse(
        "\n<script" + insertAttributes(attrs) + ">" +
        indent("\n" + js.templ, "  ") + "\n</script>")
    if err != nil {
        return Element("script", attrs)
    }
    buf := new(bytes.Buffer)
    err = t.Execute(buf, js.data)
    if err != nil {
        return Element("script", attrs)
    }
    return HTML(buf.String())
}

func Script_(js JS) HTML {
    return Script(a.Attr(), js)
}

func Javascript(data interface{}, templs ...string) JS {
    js := JS{ data: data }
    if len(templs) == 0 {
        js.templ = "{{.}}"
    } else {
        js.templ = strings.Join(templs, "\n")
    }
    return js
}

func Javascript_(templs ...string) JS {
    return Javascript(nil, templs...)
}

// Begin of generated elements

[[ range .ElementFuncs ]]
func [[.FuncName]](attrs []a.Attribute, children ...HTML) HTML {
    return Element("[[.TagName]]", attrs, children...)
}

func [[.FuncName]]_(children ...HTML) HTML {
    return [[.FuncName]](a.Attr(), children...)
}
[[ end ]]

// Begin of generated void elements

[[ range .VoidElementFuncs ]]
func [[.FuncName]](attrs []a.Attribute) HTML {
    return VoidElement("[[.TagName]]", attrs)
}
func [[.FuncName]]_() HTML {
    return [[.FuncName]](a.Attr())
}
[[ end ]]
