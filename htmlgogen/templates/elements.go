package htmlgo

import (
    "fmt"
    "io"
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

func WriteTo(w io.Writer, h HTML) {
    w.Write([]byte(h))
}

// Build a slice of type []Attribute for cosmetic purposes
func Attr(attrs ...a.Attribute) []a.Attribute {
    return attrs
}

func prepareAttributes(attrs []a.Attribute) (string, string, map[string]interface{}) {
    data := map[string]interface{}{}
    templ := ""
    defs := ""

    for _, attr := range attrs {
        data[attr.Name] = attr.Data
        templ += ` {{template "` + attr.Name + `" .` + attr.Name + `}}`
        defs += attr.Templ
    }

    return templ, defs, data
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

func buildElement(tag string, attrs []a.Attribute, content string, close_ bool) string {
    templ, defs, data := prepareAttributes(attrs)
    complTempl := defs + "\n<" + tag + templ + ">" + content
    if close_ {
        complTempl += "\n</" + tag +">"
    }

    t, _ := template.New(tag).Parse(complTempl)

    buf := new(bytes.Buffer)
    _ = t.Execute(buf, data)
    return buf.String()
}

func Element(tag string, attrs []a.Attribute, children ...HTML) HTML {
    return HTML(buildElement(tag, attrs, 
                indent(insertChildren(children...), "  "), true))
}

func VoidElement(tag string, attrs []a.Attribute) HTML {
    return HTML(buildElement(tag, attrs, "", false))
}

// Produce HTML from plain text by escaping
func Text(v interface{}) HTML {
    return HTML("\n" + html.EscapeString(fmt.Sprint(v)))
}

func TextRaw(s string) HTML {
    return HTML(s)
}

// Begin of manually defined elements

func Html5(attrs []a.Attribute, children ...HTML) HTML {
    return DoctypeHtml5 + Html(attrs, children...)
}

func Html5_(children ...HTML) HTML {
    return Html5(Attr(), children...)
}

func Doctype(t string) HTML {
    return HTML("<!DOCTYPE " + t + ">")
}

const DoctypeHtml5 HTML = "<!DOCTYPE HTML>"

func Script(attrs []a.Attribute, js JS) HTML {
    if js.data == nil {
        return Element("script", attrs, HTML("\n" + js.templ))
    }

    complTempl := buildElement("script", attrs,
                               indent("\n" + js.templ, "  "), true)
    
    // TODO set verbosity level to enable logging
    t, err := template.New("_").Delims("{%$", "$%}").Parse(complTempl)
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
    return Script(Attr(), js)
}

func JavaScript(data interface{}, templs ...string) JS {
    js := JS{ data: data }
    if len(templs) == 0 {
        js.templ = "{%$.$%}"
    } else {
        js.templ = strings.Replace(
                     strings.Replace(
                        strings.Join(templs, "\n"),
                        "{{", "{%$", -1),
                     "}}", "$%}", -1)
    }
    return js
}

func JavaScriptRaw(templs ...string) JS {
    return JavaScript(nil, templs...)
}

// Begin of generated elements

[[ range .ElementFuncs ]]
func [[.FuncName]](attrs []a.Attribute, children ...HTML) HTML {
    return Element("[[.TagName]]", attrs, children...)
}

func [[.FuncName]]_(children ...HTML) HTML {
    return [[.FuncName]](Attr(), children...)
}
[[ end ]]

// Begin of generated void elements

[[ range .VoidElementFuncs ]]
func [[.FuncName]](attrs []a.Attribute) HTML {
    return VoidElement("[[.TagName]]", attrs)
}
func [[.FuncName]]_() HTML {
    return [[.FuncName]](Attr())
}
[[ end ]]
