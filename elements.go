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


func Abbr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("abbr", attrs, children...)
}

func Address(attrs []a.Attribute, children ...HTML) HTML {
    return Element("address", attrs, children...)
}

func Article(attrs []a.Attribute, children ...HTML) HTML {
    return Element("article", attrs, children...)
}

func Aside(attrs []a.Attribute, children ...HTML) HTML {
    return Element("aside", attrs, children...)
}

func Audio(attrs []a.Attribute, children ...HTML) HTML {
    return Element("audio", attrs, children...)
}

func B(attrs []a.Attribute, children ...HTML) HTML {
    return Element("b", attrs, children...)
}

func Bdi(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bdi", attrs, children...)
}

func Bdo(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bdo", attrs, children...)
}

func Blockquote(attrs []a.Attribute, children ...HTML) HTML {
    return Element("blockquote", attrs, children...)
}

func Body(attrs []a.Attribute, children ...HTML) HTML {
    return Element("body", attrs, children...)
}

func Button(attrs []a.Attribute, children ...HTML) HTML {
    return Element("button", attrs, children...)
}

func Canvas(attrs []a.Attribute, children ...HTML) HTML {
    return Element("canvas", attrs, children...)
}

func Caption(attrs []a.Attribute, children ...HTML) HTML {
    return Element("caption", attrs, children...)
}

func Cite(attrs []a.Attribute, children ...HTML) HTML {
    return Element("cite", attrs, children...)
}

func Code(attrs []a.Attribute, children ...HTML) HTML {
    return Element("code", attrs, children...)
}

func Colgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("colgroup", attrs, children...)
}

func Command(attrs []a.Attribute, children ...HTML) HTML {
    return Element("command", attrs, children...)
}

func Datalist(attrs []a.Attribute, children ...HTML) HTML {
    return Element("datalist", attrs, children...)
}

func Dd(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dd", attrs, children...)
}

func Del(attrs []a.Attribute, children ...HTML) HTML {
    return Element("del", attrs, children...)
}

func Details(attrs []a.Attribute, children ...HTML) HTML {
    return Element("details", attrs, children...)
}

func Dfn(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dfn", attrs, children...)
}

func Div(attrs []a.Attribute, children ...HTML) HTML {
    return Element("div", attrs, children...)
}

func Dl(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dl", attrs, children...)
}

func Dt(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dt", attrs, children...)
}

func Em(attrs []a.Attribute, children ...HTML) HTML {
    return Element("em", attrs, children...)
}

func Fieldset(attrs []a.Attribute, children ...HTML) HTML {
    return Element("fieldset", attrs, children...)
}

func Figcaption(attrs []a.Attribute, children ...HTML) HTML {
    return Element("figcaption", attrs, children...)
}

func Figure(attrs []a.Attribute, children ...HTML) HTML {
    return Element("figure", attrs, children...)
}

func Footer(attrs []a.Attribute, children ...HTML) HTML {
    return Element("footer", attrs, children...)
}

func Form(attrs []a.Attribute, children ...HTML) HTML {
    return Element("form", attrs, children...)
}

func H1(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h1", attrs, children...)
}

func H2(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h2", attrs, children...)
}

func H3(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h3", attrs, children...)
}

func H4(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h4", attrs, children...)
}

func H5(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h5", attrs, children...)
}

func H6(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h6", attrs, children...)
}

func Head(attrs []a.Attribute, children ...HTML) HTML {
    return Element("head", attrs, children...)
}

func Header(attrs []a.Attribute, children ...HTML) HTML {
    return Element("header", attrs, children...)
}

func Hgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("hgroup", attrs, children...)
}

func Html(attrs []a.Attribute, children ...HTML) HTML {
    return Element("html", attrs, children...)
}

func I(attrs []a.Attribute, children ...HTML) HTML {
    return Element("i", attrs, children...)
}

func Iframe(attrs []a.Attribute, children ...HTML) HTML {
    return Element("iframe", attrs, children...)
}

func Ins(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ins", attrs, children...)
}



func Area(attrs []a.Attribute) HTML {
    return VoidElement("area", attrs)
}

func Base(attrs []a.Attribute) HTML {
    return VoidElement("base", attrs)
}

func Br(attrs []a.Attribute) HTML {
    return VoidElement("br", attrs)
}

func Col(attrs []a.Attribute) HTML {
    return VoidElement("col", attrs)
}

func Embed(attrs []a.Attribute) HTML {
    return VoidElement("embed", attrs)
}

func Hr(attrs []a.Attribute) HTML {
    return VoidElement("hr", attrs)
}

func Img(attrs []a.Attribute) HTML {
    return VoidElement("img", attrs)
}

func Input(attrs []a.Attribute) HTML {
    return VoidElement("input", attrs)
}

