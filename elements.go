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


func A(attrs []a.Attribute, children ...HTML) HTML {
    return Element("a", attrs, children...)
}

func A_(children ...HTML) HTML {
    return A(a.Attr(), children...)
}

func Abbr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("abbr", attrs, children...)
}

func Abbr_(children ...HTML) HTML {
    return Abbr(a.Attr(), children...)
}

func Acronym(attrs []a.Attribute, children ...HTML) HTML {
    return Element("acronym", attrs, children...)
}

func Acronym_(children ...HTML) HTML {
    return Acronym(a.Attr(), children...)
}

func Address(attrs []a.Attribute, children ...HTML) HTML {
    return Element("address", attrs, children...)
}

func Address_(children ...HTML) HTML {
    return Address(a.Attr(), children...)
}

func Applet(attrs []a.Attribute, children ...HTML) HTML {
    return Element("applet", attrs, children...)
}

func Applet_(children ...HTML) HTML {
    return Applet(a.Attr(), children...)
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

func Basefont(attrs []a.Attribute, children ...HTML) HTML {
    return Element("basefont", attrs, children...)
}

func Basefont_(children ...HTML) HTML {
    return Basefont(a.Attr(), children...)
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

func Bgsound(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bgsound", attrs, children...)
}

func Bgsound_(children ...HTML) HTML {
    return Bgsound(a.Attr(), children...)
}

func Big(attrs []a.Attribute, children ...HTML) HTML {
    return Element("big", attrs, children...)
}

func Big_(children ...HTML) HTML {
    return Big(a.Attr(), children...)
}

func Blink(attrs []a.Attribute, children ...HTML) HTML {
    return Element("blink", attrs, children...)
}

func Blink_(children ...HTML) HTML {
    return Blink(a.Attr(), children...)
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

func Center(attrs []a.Attribute, children ...HTML) HTML {
    return Element("center", attrs, children...)
}

func Center_(children ...HTML) HTML {
    return Center(a.Attr(), children...)
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

func Dir(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dir", attrs, children...)
}

func Dir_(children ...HTML) HTML {
    return Dir(a.Attr(), children...)
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

func Font(attrs []a.Attribute, children ...HTML) HTML {
    return Element("font", attrs, children...)
}

func Font_(children ...HTML) HTML {
    return Font(a.Attr(), children...)
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

func Frame(attrs []a.Attribute, children ...HTML) HTML {
    return Element("frame", attrs, children...)
}

func Frame_(children ...HTML) HTML {
    return Frame(a.Attr(), children...)
}

func Frameset(attrs []a.Attribute, children ...HTML) HTML {
    return Element("frameset", attrs, children...)
}

func Frameset_(children ...HTML) HTML {
    return Frameset(a.Attr(), children...)
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

func Isindex(attrs []a.Attribute, children ...HTML) HTML {
    return Element("isindex", attrs, children...)
}

func Isindex_(children ...HTML) HTML {
    return Isindex(a.Attr(), children...)
}

func Kbd(attrs []a.Attribute, children ...HTML) HTML {
    return Element("kbd", attrs, children...)
}

func Kbd_(children ...HTML) HTML {
    return Kbd(a.Attr(), children...)
}

func Keygen(attrs []a.Attribute, children ...HTML) HTML {
    return Element("keygen", attrs, children...)
}

func Keygen_(children ...HTML) HTML {
    return Keygen(a.Attr(), children...)
}

func Label(attrs []a.Attribute, children ...HTML) HTML {
    return Element("label", attrs, children...)
}

func Label_(children ...HTML) HTML {
    return Label(a.Attr(), children...)
}

func Legend(attrs []a.Attribute, children ...HTML) HTML {
    return Element("legend", attrs, children...)
}

func Legend_(children ...HTML) HTML {
    return Legend(a.Attr(), children...)
}

func Li(attrs []a.Attribute, children ...HTML) HTML {
    return Element("li", attrs, children...)
}

func Li_(children ...HTML) HTML {
    return Li(a.Attr(), children...)
}

func Listing(attrs []a.Attribute, children ...HTML) HTML {
    return Element("listing", attrs, children...)
}

func Listing_(children ...HTML) HTML {
    return Listing(a.Attr(), children...)
}

func Main(attrs []a.Attribute, children ...HTML) HTML {
    return Element("main", attrs, children...)
}

func Main_(children ...HTML) HTML {
    return Main(a.Attr(), children...)
}

func Map(attrs []a.Attribute, children ...HTML) HTML {
    return Element("map", attrs, children...)
}

func Map_(children ...HTML) HTML {
    return Map(a.Attr(), children...)
}

func Mark(attrs []a.Attribute, children ...HTML) HTML {
    return Element("mark", attrs, children...)
}

func Mark_(children ...HTML) HTML {
    return Mark(a.Attr(), children...)
}

func Marquee(attrs []a.Attribute, children ...HTML) HTML {
    return Element("marquee", attrs, children...)
}

func Marquee_(children ...HTML) HTML {
    return Marquee(a.Attr(), children...)
}

func Menu(attrs []a.Attribute, children ...HTML) HTML {
    return Element("menu", attrs, children...)
}

func Menu_(children ...HTML) HTML {
    return Menu(a.Attr(), children...)
}

func Meter(attrs []a.Attribute, children ...HTML) HTML {
    return Element("meter", attrs, children...)
}

func Meter_(children ...HTML) HTML {
    return Meter(a.Attr(), children...)
}

func Nav(attrs []a.Attribute, children ...HTML) HTML {
    return Element("nav", attrs, children...)
}

func Nav_(children ...HTML) HTML {
    return Nav(a.Attr(), children...)
}

func Nobr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("nobr", attrs, children...)
}

func Nobr_(children ...HTML) HTML {
    return Nobr(a.Attr(), children...)
}

func Noframes(attrs []a.Attribute, children ...HTML) HTML {
    return Element("noframes", attrs, children...)
}

func Noframes_(children ...HTML) HTML {
    return Noframes(a.Attr(), children...)
}

func Noscript(attrs []a.Attribute, children ...HTML) HTML {
    return Element("noscript", attrs, children...)
}

func Noscript_(children ...HTML) HTML {
    return Noscript(a.Attr(), children...)
}

func Object(attrs []a.Attribute, children ...HTML) HTML {
    return Element("object", attrs, children...)
}

func Object_(children ...HTML) HTML {
    return Object(a.Attr(), children...)
}

func Ol(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ol", attrs, children...)
}

func Ol_(children ...HTML) HTML {
    return Ol(a.Attr(), children...)
}

func Optgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("optgroup", attrs, children...)
}

func Optgroup_(children ...HTML) HTML {
    return Optgroup(a.Attr(), children...)
}

func Option(attrs []a.Attribute, children ...HTML) HTML {
    return Element("option", attrs, children...)
}

func Option_(children ...HTML) HTML {
    return Option(a.Attr(), children...)
}

func Output(attrs []a.Attribute, children ...HTML) HTML {
    return Element("output", attrs, children...)
}

func Output_(children ...HTML) HTML {
    return Output(a.Attr(), children...)
}

func P(attrs []a.Attribute, children ...HTML) HTML {
    return Element("p", attrs, children...)
}

func P_(children ...HTML) HTML {
    return P(a.Attr(), children...)
}

func Plaintext(attrs []a.Attribute, children ...HTML) HTML {
    return Element("plaintext", attrs, children...)
}

func Plaintext_(children ...HTML) HTML {
    return Plaintext(a.Attr(), children...)
}

func Pre(attrs []a.Attribute, children ...HTML) HTML {
    return Element("pre", attrs, children...)
}

func Pre_(children ...HTML) HTML {
    return Pre(a.Attr(), children...)
}

func Progress(attrs []a.Attribute, children ...HTML) HTML {
    return Element("progress", attrs, children...)
}

func Progress_(children ...HTML) HTML {
    return Progress(a.Attr(), children...)
}

func Q(attrs []a.Attribute, children ...HTML) HTML {
    return Element("q", attrs, children...)
}

func Q_(children ...HTML) HTML {
    return Q(a.Attr(), children...)
}

func Rp(attrs []a.Attribute, children ...HTML) HTML {
    return Element("rp", attrs, children...)
}

func Rp_(children ...HTML) HTML {
    return Rp(a.Attr(), children...)
}

func Rt(attrs []a.Attribute, children ...HTML) HTML {
    return Element("rt", attrs, children...)
}

func Rt_(children ...HTML) HTML {
    return Rt(a.Attr(), children...)
}

func Ruby(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ruby", attrs, children...)
}

func Ruby_(children ...HTML) HTML {
    return Ruby(a.Attr(), children...)
}

func S(attrs []a.Attribute, children ...HTML) HTML {
    return Element("s", attrs, children...)
}

func S_(children ...HTML) HTML {
    return S(a.Attr(), children...)
}

func Samp(attrs []a.Attribute, children ...HTML) HTML {
    return Element("samp", attrs, children...)
}

func Samp_(children ...HTML) HTML {
    return Samp(a.Attr(), children...)
}

func Section(attrs []a.Attribute, children ...HTML) HTML {
    return Element("section", attrs, children...)
}

func Section_(children ...HTML) HTML {
    return Section(a.Attr(), children...)
}

func Select(attrs []a.Attribute, children ...HTML) HTML {
    return Element("select", attrs, children...)
}

func Select_(children ...HTML) HTML {
    return Select(a.Attr(), children...)
}

func Small(attrs []a.Attribute, children ...HTML) HTML {
    return Element("small", attrs, children...)
}

func Small_(children ...HTML) HTML {
    return Small(a.Attr(), children...)
}

func Spacer(attrs []a.Attribute, children ...HTML) HTML {
    return Element("spacer", attrs, children...)
}

func Spacer_(children ...HTML) HTML {
    return Spacer(a.Attr(), children...)
}

func Span(attrs []a.Attribute, children ...HTML) HTML {
    return Element("span", attrs, children...)
}

func Span_(children ...HTML) HTML {
    return Span(a.Attr(), children...)
}

func Strike(attrs []a.Attribute, children ...HTML) HTML {
    return Element("strike", attrs, children...)
}

func Strike_(children ...HTML) HTML {
    return Strike(a.Attr(), children...)
}

func Strong(attrs []a.Attribute, children ...HTML) HTML {
    return Element("strong", attrs, children...)
}

func Strong_(children ...HTML) HTML {
    return Strong(a.Attr(), children...)
}

func Style(attrs []a.Attribute, children ...HTML) HTML {
    return Element("style", attrs, children...)
}

func Style_(children ...HTML) HTML {
    return Style(a.Attr(), children...)
}

func Sub(attrs []a.Attribute, children ...HTML) HTML {
    return Element("sub", attrs, children...)
}

func Sub_(children ...HTML) HTML {
    return Sub(a.Attr(), children...)
}

func Summary(attrs []a.Attribute, children ...HTML) HTML {
    return Element("summary", attrs, children...)
}

func Summary_(children ...HTML) HTML {
    return Summary(a.Attr(), children...)
}

func Sup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("sup", attrs, children...)
}

func Sup_(children ...HTML) HTML {
    return Sup(a.Attr(), children...)
}

func Table(attrs []a.Attribute, children ...HTML) HTML {
    return Element("table", attrs, children...)
}

func Table_(children ...HTML) HTML {
    return Table(a.Attr(), children...)
}

func Tbody(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tbody", attrs, children...)
}

func Tbody_(children ...HTML) HTML {
    return Tbody(a.Attr(), children...)
}

func Td(attrs []a.Attribute, children ...HTML) HTML {
    return Element("td", attrs, children...)
}

func Td_(children ...HTML) HTML {
    return Td(a.Attr(), children...)
}

func Textarea(attrs []a.Attribute, children ...HTML) HTML {
    return Element("textarea", attrs, children...)
}

func Textarea_(children ...HTML) HTML {
    return Textarea(a.Attr(), children...)
}

func Tfoot(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tfoot", attrs, children...)
}

func Tfoot_(children ...HTML) HTML {
    return Tfoot(a.Attr(), children...)
}

func Th(attrs []a.Attribute, children ...HTML) HTML {
    return Element("th", attrs, children...)
}

func Th_(children ...HTML) HTML {
    return Th(a.Attr(), children...)
}

func Thead(attrs []a.Attribute, children ...HTML) HTML {
    return Element("thead", attrs, children...)
}

func Thead_(children ...HTML) HTML {
    return Thead(a.Attr(), children...)
}

func Time(attrs []a.Attribute, children ...HTML) HTML {
    return Element("time", attrs, children...)
}

func Time_(children ...HTML) HTML {
    return Time(a.Attr(), children...)
}

func Title(attrs []a.Attribute, children ...HTML) HTML {
    return Element("title", attrs, children...)
}

func Title_(children ...HTML) HTML {
    return Title(a.Attr(), children...)
}

func Tr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tr", attrs, children...)
}

func Tr_(children ...HTML) HTML {
    return Tr(a.Attr(), children...)
}

func Tt(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tt", attrs, children...)
}

func Tt_(children ...HTML) HTML {
    return Tt(a.Attr(), children...)
}

func U(attrs []a.Attribute, children ...HTML) HTML {
    return Element("u", attrs, children...)
}

func U_(children ...HTML) HTML {
    return U(a.Attr(), children...)
}

func Ul(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ul", attrs, children...)
}

func Ul_(children ...HTML) HTML {
    return Ul(a.Attr(), children...)
}

func Var(attrs []a.Attribute, children ...HTML) HTML {
    return Element("var", attrs, children...)
}

func Var_(children ...HTML) HTML {
    return Var(a.Attr(), children...)
}

func Video(attrs []a.Attribute, children ...HTML) HTML {
    return Element("video", attrs, children...)
}

func Video_(children ...HTML) HTML {
    return Video(a.Attr(), children...)
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

func Link(attrs []a.Attribute) HTML {
    return VoidElement("link", attrs)
}
func Link_() HTML {
    return Link(a.Attr())
}

func Meta(attrs []a.Attribute) HTML {
    return VoidElement("meta", attrs)
}
func Meta_() HTML {
    return Meta(a.Attr())
}

func Param(attrs []a.Attribute) HTML {
    return VoidElement("param", attrs)
}
func Param_() HTML {
    return Param(a.Attr())
}

func Source(attrs []a.Attribute) HTML {
    return VoidElement("source", attrs)
}
func Source_() HTML {
    return Source(a.Attr())
}

func Track(attrs []a.Attribute) HTML {
    return VoidElement("track", attrs)
}
func Track_() HTML {
    return Track(a.Attr())
}

func Wbr(attrs []a.Attribute) HTML {
    return VoidElement("wbr", attrs)
}
func Wbr_() HTML {
    return Wbr(a.Attr())
}

