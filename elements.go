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

// Begin of generated elements


func Abbr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("abbr", attrs, children...)
}

func Abbr_(children ...HTML) HTML {
    return Abbr(a.Attr(), children...)
}

func Address(attrs []a.Attribute, children ...HTML) HTML {
    return Element("address", attrs, children...)
}

func Address_(children ...HTML) HTML {
    return Address(a.Attr(), children...)
}

func Article(attrs []a.Attribute, children ...HTML) HTML {
    return Element("article", attrs, children...)
}

func Article_(children ...HTML) HTML {
    return Article(a.Attr(), children...)
}

func Aside(attrs []a.Attribute, children ...HTML) HTML {
    return Element("aside", attrs, children...)
}

func Aside_(children ...HTML) HTML {
    return Aside(a.Attr(), children...)
}

func Audio(attrs []a.Attribute, children ...HTML) HTML {
    return Element("audio", attrs, children...)
}

func Audio_(children ...HTML) HTML {
    return Audio(a.Attr(), children...)
}

func B(attrs []a.Attribute, children ...HTML) HTML {
    return Element("b", attrs, children...)
}

func B_(children ...HTML) HTML {
    return B(a.Attr(), children...)
}

func Bdi(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bdi", attrs, children...)
}

func Bdi_(children ...HTML) HTML {
    return Bdi(a.Attr(), children...)
}

func Bdo(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bdo", attrs, children...)
}

func Bdo_(children ...HTML) HTML {
    return Bdo(a.Attr(), children...)
}

func Blockquote(attrs []a.Attribute, children ...HTML) HTML {
    return Element("blockquote", attrs, children...)
}

func Blockquote_(children ...HTML) HTML {
    return Blockquote(a.Attr(), children...)
}

func Body(attrs []a.Attribute, children ...HTML) HTML {
    return Element("body", attrs, children...)
}

func Body_(children ...HTML) HTML {
    return Body(a.Attr(), children...)
}

func Button(attrs []a.Attribute, children ...HTML) HTML {
    return Element("button", attrs, children...)
}

func Button_(children ...HTML) HTML {
    return Button(a.Attr(), children...)
}

func Canvas(attrs []a.Attribute, children ...HTML) HTML {
    return Element("canvas", attrs, children...)
}

func Canvas_(children ...HTML) HTML {
    return Canvas(a.Attr(), children...)
}

func Caption(attrs []a.Attribute, children ...HTML) HTML {
    return Element("caption", attrs, children...)
}

func Caption_(children ...HTML) HTML {
    return Caption(a.Attr(), children...)
}

func Cite(attrs []a.Attribute, children ...HTML) HTML {
    return Element("cite", attrs, children...)
}

func Cite_(children ...HTML) HTML {
    return Cite(a.Attr(), children...)
}

func Code(attrs []a.Attribute, children ...HTML) HTML {
    return Element("code", attrs, children...)
}

func Code_(children ...HTML) HTML {
    return Code(a.Attr(), children...)
}

func Colgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("colgroup", attrs, children...)
}

func Colgroup_(children ...HTML) HTML {
    return Colgroup(a.Attr(), children...)
}

func Command(attrs []a.Attribute, children ...HTML) HTML {
    return Element("command", attrs, children...)
}

func Command_(children ...HTML) HTML {
    return Command(a.Attr(), children...)
}

func Datalist(attrs []a.Attribute, children ...HTML) HTML {
    return Element("datalist", attrs, children...)
}

func Datalist_(children ...HTML) HTML {
    return Datalist(a.Attr(), children...)
}

func Dd(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dd", attrs, children...)
}

func Dd_(children ...HTML) HTML {
    return Dd(a.Attr(), children...)
}

func Del(attrs []a.Attribute, children ...HTML) HTML {
    return Element("del", attrs, children...)
}

func Del_(children ...HTML) HTML {
    return Del(a.Attr(), children...)
}

func Details(attrs []a.Attribute, children ...HTML) HTML {
    return Element("details", attrs, children...)
}

func Details_(children ...HTML) HTML {
    return Details(a.Attr(), children...)
}

func Dfn(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dfn", attrs, children...)
}

func Dfn_(children ...HTML) HTML {
    return Dfn(a.Attr(), children...)
}

func Div(attrs []a.Attribute, children ...HTML) HTML {
    return Element("div", attrs, children...)
}

func Div_(children ...HTML) HTML {
    return Div(a.Attr(), children...)
}

func Dl(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dl", attrs, children...)
}

func Dl_(children ...HTML) HTML {
    return Dl(a.Attr(), children...)
}

func Dt(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dt", attrs, children...)
}

func Dt_(children ...HTML) HTML {
    return Dt(a.Attr(), children...)
}

func Em(attrs []a.Attribute, children ...HTML) HTML {
    return Element("em", attrs, children...)
}

func Em_(children ...HTML) HTML {
    return Em(a.Attr(), children...)
}

func Fieldset(attrs []a.Attribute, children ...HTML) HTML {
    return Element("fieldset", attrs, children...)
}

func Fieldset_(children ...HTML) HTML {
    return Fieldset(a.Attr(), children...)
}

func Figcaption(attrs []a.Attribute, children ...HTML) HTML {
    return Element("figcaption", attrs, children...)
}

func Figcaption_(children ...HTML) HTML {
    return Figcaption(a.Attr(), children...)
}

func Figure(attrs []a.Attribute, children ...HTML) HTML {
    return Element("figure", attrs, children...)
}

func Figure_(children ...HTML) HTML {
    return Figure(a.Attr(), children...)
}

func Footer(attrs []a.Attribute, children ...HTML) HTML {
    return Element("footer", attrs, children...)
}

func Footer_(children ...HTML) HTML {
    return Footer(a.Attr(), children...)
}

func Form(attrs []a.Attribute, children ...HTML) HTML {
    return Element("form", attrs, children...)
}

func Form_(children ...HTML) HTML {
    return Form(a.Attr(), children...)
}

func H1(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h1", attrs, children...)
}

func H1_(children ...HTML) HTML {
    return H1(a.Attr(), children...)
}

func H2(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h2", attrs, children...)
}

func H2_(children ...HTML) HTML {
    return H2(a.Attr(), children...)
}

func H3(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h3", attrs, children...)
}

func H3_(children ...HTML) HTML {
    return H3(a.Attr(), children...)
}

func H4(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h4", attrs, children...)
}

func H4_(children ...HTML) HTML {
    return H4(a.Attr(), children...)
}

func H5(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h5", attrs, children...)
}

func H5_(children ...HTML) HTML {
    return H5(a.Attr(), children...)
}

func H6(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h6", attrs, children...)
}

func H6_(children ...HTML) HTML {
    return H6(a.Attr(), children...)
}

func Head(attrs []a.Attribute, children ...HTML) HTML {
    return Element("head", attrs, children...)
}

func Head_(children ...HTML) HTML {
    return Head(a.Attr(), children...)
}

func Header(attrs []a.Attribute, children ...HTML) HTML {
    return Element("header", attrs, children...)
}

func Header_(children ...HTML) HTML {
    return Header(a.Attr(), children...)
}

func Hgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("hgroup", attrs, children...)
}

func Hgroup_(children ...HTML) HTML {
    return Hgroup(a.Attr(), children...)
}

func Html(attrs []a.Attribute, children ...HTML) HTML {
    return Element("html", attrs, children...)
}

func Html_(children ...HTML) HTML {
    return Html(a.Attr(), children...)
}

func I(attrs []a.Attribute, children ...HTML) HTML {
    return Element("i", attrs, children...)
}

func I_(children ...HTML) HTML {
    return I(a.Attr(), children...)
}

func Iframe(attrs []a.Attribute, children ...HTML) HTML {
    return Element("iframe", attrs, children...)
}

func Iframe_(children ...HTML) HTML {
    return Iframe(a.Attr(), children...)
}

func Ins(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ins", attrs, children...)
}

func Ins_(children ...HTML) HTML {
    return Ins(a.Attr(), children...)
}


// Begin of generated void elements


func Area(attrs []a.Attribute) HTML {
    return VoidElement("area", attrs)
}
func Area_() HTML {
    return Area(a.Attr())
}

func Base(attrs []a.Attribute) HTML {
    return VoidElement("base", attrs)
}
func Base_() HTML {
    return Base(a.Attr())
}

func Br(attrs []a.Attribute) HTML {
    return VoidElement("br", attrs)
}
func Br_() HTML {
    return Br(a.Attr())
}

func Col(attrs []a.Attribute) HTML {
    return VoidElement("col", attrs)
}
func Col_() HTML {
    return Col(a.Attr())
}

func Embed(attrs []a.Attribute) HTML {
    return VoidElement("embed", attrs)
}
func Embed_() HTML {
    return Embed(a.Attr())
}

func Hr(attrs []a.Attribute) HTML {
    return VoidElement("hr", attrs)
}
func Hr_() HTML {
    return Hr(a.Attr())
}

func Img(attrs []a.Attribute) HTML {
    return VoidElement("img", attrs)
}
func Img_() HTML {
    return Img(a.Attr())
}

func Input(attrs []a.Attribute) HTML {
    return VoidElement("input", attrs)
}
func Input_() HTML {
    return Input(a.Attr())
}

