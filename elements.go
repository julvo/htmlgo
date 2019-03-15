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

func Text_(s string) HTML {
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

func JavaScript_(templs ...string) JS {
    return JavaScript(nil, templs...)
}

// Begin of generated elements


func A(attrs []a.Attribute, children ...HTML) HTML {
    return Element("a", attrs, children...)
}

func A_(children ...HTML) HTML {
    return A(Attr(), children...)
}

func Abbr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("abbr", attrs, children...)
}

func Abbr_(children ...HTML) HTML {
    return Abbr(Attr(), children...)
}

func Acronym(attrs []a.Attribute, children ...HTML) HTML {
    return Element("acronym", attrs, children...)
}

func Acronym_(children ...HTML) HTML {
    return Acronym(Attr(), children...)
}

func Address(attrs []a.Attribute, children ...HTML) HTML {
    return Element("address", attrs, children...)
}

func Address_(children ...HTML) HTML {
    return Address(Attr(), children...)
}

func Applet(attrs []a.Attribute, children ...HTML) HTML {
    return Element("applet", attrs, children...)
}

func Applet_(children ...HTML) HTML {
    return Applet(Attr(), children...)
}

func Article(attrs []a.Attribute, children ...HTML) HTML {
    return Element("article", attrs, children...)
}

func Article_(children ...HTML) HTML {
    return Article(Attr(), children...)
}

func Aside(attrs []a.Attribute, children ...HTML) HTML {
    return Element("aside", attrs, children...)
}

func Aside_(children ...HTML) HTML {
    return Aside(Attr(), children...)
}

func Audio(attrs []a.Attribute, children ...HTML) HTML {
    return Element("audio", attrs, children...)
}

func Audio_(children ...HTML) HTML {
    return Audio(Attr(), children...)
}

func B(attrs []a.Attribute, children ...HTML) HTML {
    return Element("b", attrs, children...)
}

func B_(children ...HTML) HTML {
    return B(Attr(), children...)
}

func Basefont(attrs []a.Attribute, children ...HTML) HTML {
    return Element("basefont", attrs, children...)
}

func Basefont_(children ...HTML) HTML {
    return Basefont(Attr(), children...)
}

func Bdi(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bdi", attrs, children...)
}

func Bdi_(children ...HTML) HTML {
    return Bdi(Attr(), children...)
}

func Bdo(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bdo", attrs, children...)
}

func Bdo_(children ...HTML) HTML {
    return Bdo(Attr(), children...)
}

func Bgsound(attrs []a.Attribute, children ...HTML) HTML {
    return Element("bgsound", attrs, children...)
}

func Bgsound_(children ...HTML) HTML {
    return Bgsound(Attr(), children...)
}

func Big(attrs []a.Attribute, children ...HTML) HTML {
    return Element("big", attrs, children...)
}

func Big_(children ...HTML) HTML {
    return Big(Attr(), children...)
}

func Blink(attrs []a.Attribute, children ...HTML) HTML {
    return Element("blink", attrs, children...)
}

func Blink_(children ...HTML) HTML {
    return Blink(Attr(), children...)
}

func Blockquote(attrs []a.Attribute, children ...HTML) HTML {
    return Element("blockquote", attrs, children...)
}

func Blockquote_(children ...HTML) HTML {
    return Blockquote(Attr(), children...)
}

func Body(attrs []a.Attribute, children ...HTML) HTML {
    return Element("body", attrs, children...)
}

func Body_(children ...HTML) HTML {
    return Body(Attr(), children...)
}

func Button(attrs []a.Attribute, children ...HTML) HTML {
    return Element("button", attrs, children...)
}

func Button_(children ...HTML) HTML {
    return Button(Attr(), children...)
}

func Canvas(attrs []a.Attribute, children ...HTML) HTML {
    return Element("canvas", attrs, children...)
}

func Canvas_(children ...HTML) HTML {
    return Canvas(Attr(), children...)
}

func Caption(attrs []a.Attribute, children ...HTML) HTML {
    return Element("caption", attrs, children...)
}

func Caption_(children ...HTML) HTML {
    return Caption(Attr(), children...)
}

func Center(attrs []a.Attribute, children ...HTML) HTML {
    return Element("center", attrs, children...)
}

func Center_(children ...HTML) HTML {
    return Center(Attr(), children...)
}

func Cite(attrs []a.Attribute, children ...HTML) HTML {
    return Element("cite", attrs, children...)
}

func Cite_(children ...HTML) HTML {
    return Cite(Attr(), children...)
}

func Code(attrs []a.Attribute, children ...HTML) HTML {
    return Element("code", attrs, children...)
}

func Code_(children ...HTML) HTML {
    return Code(Attr(), children...)
}

func Colgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("colgroup", attrs, children...)
}

func Colgroup_(children ...HTML) HTML {
    return Colgroup(Attr(), children...)
}

func Datalist(attrs []a.Attribute, children ...HTML) HTML {
    return Element("datalist", attrs, children...)
}

func Datalist_(children ...HTML) HTML {
    return Datalist(Attr(), children...)
}

func Dd(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dd", attrs, children...)
}

func Dd_(children ...HTML) HTML {
    return Dd(Attr(), children...)
}

func Del(attrs []a.Attribute, children ...HTML) HTML {
    return Element("del", attrs, children...)
}

func Del_(children ...HTML) HTML {
    return Del(Attr(), children...)
}

func Details(attrs []a.Attribute, children ...HTML) HTML {
    return Element("details", attrs, children...)
}

func Details_(children ...HTML) HTML {
    return Details(Attr(), children...)
}

func Dfn(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dfn", attrs, children...)
}

func Dfn_(children ...HTML) HTML {
    return Dfn(Attr(), children...)
}

func Dir(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dir", attrs, children...)
}

func Dir_(children ...HTML) HTML {
    return Dir(Attr(), children...)
}

func Div(attrs []a.Attribute, children ...HTML) HTML {
    return Element("div", attrs, children...)
}

func Div_(children ...HTML) HTML {
    return Div(Attr(), children...)
}

func Dl(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dl", attrs, children...)
}

func Dl_(children ...HTML) HTML {
    return Dl(Attr(), children...)
}

func Dt(attrs []a.Attribute, children ...HTML) HTML {
    return Element("dt", attrs, children...)
}

func Dt_(children ...HTML) HTML {
    return Dt(Attr(), children...)
}

func Em(attrs []a.Attribute, children ...HTML) HTML {
    return Element("em", attrs, children...)
}

func Em_(children ...HTML) HTML {
    return Em(Attr(), children...)
}

func Fieldset(attrs []a.Attribute, children ...HTML) HTML {
    return Element("fieldset", attrs, children...)
}

func Fieldset_(children ...HTML) HTML {
    return Fieldset(Attr(), children...)
}

func Figcaption(attrs []a.Attribute, children ...HTML) HTML {
    return Element("figcaption", attrs, children...)
}

func Figcaption_(children ...HTML) HTML {
    return Figcaption(Attr(), children...)
}

func Figure(attrs []a.Attribute, children ...HTML) HTML {
    return Element("figure", attrs, children...)
}

func Figure_(children ...HTML) HTML {
    return Figure(Attr(), children...)
}

func Font(attrs []a.Attribute, children ...HTML) HTML {
    return Element("font", attrs, children...)
}

func Font_(children ...HTML) HTML {
    return Font(Attr(), children...)
}

func Footer(attrs []a.Attribute, children ...HTML) HTML {
    return Element("footer", attrs, children...)
}

func Footer_(children ...HTML) HTML {
    return Footer(Attr(), children...)
}

func Form(attrs []a.Attribute, children ...HTML) HTML {
    return Element("form", attrs, children...)
}

func Form_(children ...HTML) HTML {
    return Form(Attr(), children...)
}

func Frame(attrs []a.Attribute, children ...HTML) HTML {
    return Element("frame", attrs, children...)
}

func Frame_(children ...HTML) HTML {
    return Frame(Attr(), children...)
}

func Frameset(attrs []a.Attribute, children ...HTML) HTML {
    return Element("frameset", attrs, children...)
}

func Frameset_(children ...HTML) HTML {
    return Frameset(Attr(), children...)
}

func H1(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h1", attrs, children...)
}

func H1_(children ...HTML) HTML {
    return H1(Attr(), children...)
}

func H2(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h2", attrs, children...)
}

func H2_(children ...HTML) HTML {
    return H2(Attr(), children...)
}

func H3(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h3", attrs, children...)
}

func H3_(children ...HTML) HTML {
    return H3(Attr(), children...)
}

func H4(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h4", attrs, children...)
}

func H4_(children ...HTML) HTML {
    return H4(Attr(), children...)
}

func H5(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h5", attrs, children...)
}

func H5_(children ...HTML) HTML {
    return H5(Attr(), children...)
}

func H6(attrs []a.Attribute, children ...HTML) HTML {
    return Element("h6", attrs, children...)
}

func H6_(children ...HTML) HTML {
    return H6(Attr(), children...)
}

func Head(attrs []a.Attribute, children ...HTML) HTML {
    return Element("head", attrs, children...)
}

func Head_(children ...HTML) HTML {
    return Head(Attr(), children...)
}

func Header(attrs []a.Attribute, children ...HTML) HTML {
    return Element("header", attrs, children...)
}

func Header_(children ...HTML) HTML {
    return Header(Attr(), children...)
}

func Hgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("hgroup", attrs, children...)
}

func Hgroup_(children ...HTML) HTML {
    return Hgroup(Attr(), children...)
}

func Html(attrs []a.Attribute, children ...HTML) HTML {
    return Element("html", attrs, children...)
}

func Html_(children ...HTML) HTML {
    return Html(Attr(), children...)
}

func I(attrs []a.Attribute, children ...HTML) HTML {
    return Element("i", attrs, children...)
}

func I_(children ...HTML) HTML {
    return I(Attr(), children...)
}

func Iframe(attrs []a.Attribute, children ...HTML) HTML {
    return Element("iframe", attrs, children...)
}

func Iframe_(children ...HTML) HTML {
    return Iframe(Attr(), children...)
}

func Ins(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ins", attrs, children...)
}

func Ins_(children ...HTML) HTML {
    return Ins(Attr(), children...)
}

func Isindex(attrs []a.Attribute, children ...HTML) HTML {
    return Element("isindex", attrs, children...)
}

func Isindex_(children ...HTML) HTML {
    return Isindex(Attr(), children...)
}

func Kbd(attrs []a.Attribute, children ...HTML) HTML {
    return Element("kbd", attrs, children...)
}

func Kbd_(children ...HTML) HTML {
    return Kbd(Attr(), children...)
}

func Keygen(attrs []a.Attribute, children ...HTML) HTML {
    return Element("keygen", attrs, children...)
}

func Keygen_(children ...HTML) HTML {
    return Keygen(Attr(), children...)
}

func Label(attrs []a.Attribute, children ...HTML) HTML {
    return Element("label", attrs, children...)
}

func Label_(children ...HTML) HTML {
    return Label(Attr(), children...)
}

func Legend(attrs []a.Attribute, children ...HTML) HTML {
    return Element("legend", attrs, children...)
}

func Legend_(children ...HTML) HTML {
    return Legend(Attr(), children...)
}

func Li(attrs []a.Attribute, children ...HTML) HTML {
    return Element("li", attrs, children...)
}

func Li_(children ...HTML) HTML {
    return Li(Attr(), children...)
}

func Listing(attrs []a.Attribute, children ...HTML) HTML {
    return Element("listing", attrs, children...)
}

func Listing_(children ...HTML) HTML {
    return Listing(Attr(), children...)
}

func Main(attrs []a.Attribute, children ...HTML) HTML {
    return Element("main", attrs, children...)
}

func Main_(children ...HTML) HTML {
    return Main(Attr(), children...)
}

func Map(attrs []a.Attribute, children ...HTML) HTML {
    return Element("map", attrs, children...)
}

func Map_(children ...HTML) HTML {
    return Map(Attr(), children...)
}

func Mark(attrs []a.Attribute, children ...HTML) HTML {
    return Element("mark", attrs, children...)
}

func Mark_(children ...HTML) HTML {
    return Mark(Attr(), children...)
}

func Marquee(attrs []a.Attribute, children ...HTML) HTML {
    return Element("marquee", attrs, children...)
}

func Marquee_(children ...HTML) HTML {
    return Marquee(Attr(), children...)
}

func Menu(attrs []a.Attribute, children ...HTML) HTML {
    return Element("menu", attrs, children...)
}

func Menu_(children ...HTML) HTML {
    return Menu(Attr(), children...)
}

func Meter(attrs []a.Attribute, children ...HTML) HTML {
    return Element("meter", attrs, children...)
}

func Meter_(children ...HTML) HTML {
    return Meter(Attr(), children...)
}

func Nav(attrs []a.Attribute, children ...HTML) HTML {
    return Element("nav", attrs, children...)
}

func Nav_(children ...HTML) HTML {
    return Nav(Attr(), children...)
}

func Nobr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("nobr", attrs, children...)
}

func Nobr_(children ...HTML) HTML {
    return Nobr(Attr(), children...)
}

func Noframes(attrs []a.Attribute, children ...HTML) HTML {
    return Element("noframes", attrs, children...)
}

func Noframes_(children ...HTML) HTML {
    return Noframes(Attr(), children...)
}

func Noscript(attrs []a.Attribute, children ...HTML) HTML {
    return Element("noscript", attrs, children...)
}

func Noscript_(children ...HTML) HTML {
    return Noscript(Attr(), children...)
}

func Object(attrs []a.Attribute, children ...HTML) HTML {
    return Element("object", attrs, children...)
}

func Object_(children ...HTML) HTML {
    return Object(Attr(), children...)
}

func Ol(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ol", attrs, children...)
}

func Ol_(children ...HTML) HTML {
    return Ol(Attr(), children...)
}

func Optgroup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("optgroup", attrs, children...)
}

func Optgroup_(children ...HTML) HTML {
    return Optgroup(Attr(), children...)
}

func Option(attrs []a.Attribute, children ...HTML) HTML {
    return Element("option", attrs, children...)
}

func Option_(children ...HTML) HTML {
    return Option(Attr(), children...)
}

func Output(attrs []a.Attribute, children ...HTML) HTML {
    return Element("output", attrs, children...)
}

func Output_(children ...HTML) HTML {
    return Output(Attr(), children...)
}

func P(attrs []a.Attribute, children ...HTML) HTML {
    return Element("p", attrs, children...)
}

func P_(children ...HTML) HTML {
    return P(Attr(), children...)
}

func Plaintext(attrs []a.Attribute, children ...HTML) HTML {
    return Element("plaintext", attrs, children...)
}

func Plaintext_(children ...HTML) HTML {
    return Plaintext(Attr(), children...)
}

func Pre(attrs []a.Attribute, children ...HTML) HTML {
    return Element("pre", attrs, children...)
}

func Pre_(children ...HTML) HTML {
    return Pre(Attr(), children...)
}

func Progress(attrs []a.Attribute, children ...HTML) HTML {
    return Element("progress", attrs, children...)
}

func Progress_(children ...HTML) HTML {
    return Progress(Attr(), children...)
}

func Q(attrs []a.Attribute, children ...HTML) HTML {
    return Element("q", attrs, children...)
}

func Q_(children ...HTML) HTML {
    return Q(Attr(), children...)
}

func Rp(attrs []a.Attribute, children ...HTML) HTML {
    return Element("rp", attrs, children...)
}

func Rp_(children ...HTML) HTML {
    return Rp(Attr(), children...)
}

func Rt(attrs []a.Attribute, children ...HTML) HTML {
    return Element("rt", attrs, children...)
}

func Rt_(children ...HTML) HTML {
    return Rt(Attr(), children...)
}

func Ruby(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ruby", attrs, children...)
}

func Ruby_(children ...HTML) HTML {
    return Ruby(Attr(), children...)
}

func S(attrs []a.Attribute, children ...HTML) HTML {
    return Element("s", attrs, children...)
}

func S_(children ...HTML) HTML {
    return S(Attr(), children...)
}

func Samp(attrs []a.Attribute, children ...HTML) HTML {
    return Element("samp", attrs, children...)
}

func Samp_(children ...HTML) HTML {
    return Samp(Attr(), children...)
}

func Section(attrs []a.Attribute, children ...HTML) HTML {
    return Element("section", attrs, children...)
}

func Section_(children ...HTML) HTML {
    return Section(Attr(), children...)
}

func Select(attrs []a.Attribute, children ...HTML) HTML {
    return Element("select", attrs, children...)
}

func Select_(children ...HTML) HTML {
    return Select(Attr(), children...)
}

func Small(attrs []a.Attribute, children ...HTML) HTML {
    return Element("small", attrs, children...)
}

func Small_(children ...HTML) HTML {
    return Small(Attr(), children...)
}

func Spacer(attrs []a.Attribute, children ...HTML) HTML {
    return Element("spacer", attrs, children...)
}

func Spacer_(children ...HTML) HTML {
    return Spacer(Attr(), children...)
}

func Span(attrs []a.Attribute, children ...HTML) HTML {
    return Element("span", attrs, children...)
}

func Span_(children ...HTML) HTML {
    return Span(Attr(), children...)
}

func Strike(attrs []a.Attribute, children ...HTML) HTML {
    return Element("strike", attrs, children...)
}

func Strike_(children ...HTML) HTML {
    return Strike(Attr(), children...)
}

func Strong(attrs []a.Attribute, children ...HTML) HTML {
    return Element("strong", attrs, children...)
}

func Strong_(children ...HTML) HTML {
    return Strong(Attr(), children...)
}

func Style(attrs []a.Attribute, children ...HTML) HTML {
    return Element("style", attrs, children...)
}

func Style_(children ...HTML) HTML {
    return Style(Attr(), children...)
}

func Sub(attrs []a.Attribute, children ...HTML) HTML {
    return Element("sub", attrs, children...)
}

func Sub_(children ...HTML) HTML {
    return Sub(Attr(), children...)
}

func Summary(attrs []a.Attribute, children ...HTML) HTML {
    return Element("summary", attrs, children...)
}

func Summary_(children ...HTML) HTML {
    return Summary(Attr(), children...)
}

func Sup(attrs []a.Attribute, children ...HTML) HTML {
    return Element("sup", attrs, children...)
}

func Sup_(children ...HTML) HTML {
    return Sup(Attr(), children...)
}

func Table(attrs []a.Attribute, children ...HTML) HTML {
    return Element("table", attrs, children...)
}

func Table_(children ...HTML) HTML {
    return Table(Attr(), children...)
}

func Tbody(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tbody", attrs, children...)
}

func Tbody_(children ...HTML) HTML {
    return Tbody(Attr(), children...)
}

func Td(attrs []a.Attribute, children ...HTML) HTML {
    return Element("td", attrs, children...)
}

func Td_(children ...HTML) HTML {
    return Td(Attr(), children...)
}

func Textarea(attrs []a.Attribute, children ...HTML) HTML {
    return Element("textarea", attrs, children...)
}

func Textarea_(children ...HTML) HTML {
    return Textarea(Attr(), children...)
}

func Tfoot(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tfoot", attrs, children...)
}

func Tfoot_(children ...HTML) HTML {
    return Tfoot(Attr(), children...)
}

func Th(attrs []a.Attribute, children ...HTML) HTML {
    return Element("th", attrs, children...)
}

func Th_(children ...HTML) HTML {
    return Th(Attr(), children...)
}

func Thead(attrs []a.Attribute, children ...HTML) HTML {
    return Element("thead", attrs, children...)
}

func Thead_(children ...HTML) HTML {
    return Thead(Attr(), children...)
}

func Time(attrs []a.Attribute, children ...HTML) HTML {
    return Element("time", attrs, children...)
}

func Time_(children ...HTML) HTML {
    return Time(Attr(), children...)
}

func Title(attrs []a.Attribute, children ...HTML) HTML {
    return Element("title", attrs, children...)
}

func Title_(children ...HTML) HTML {
    return Title(Attr(), children...)
}

func Tr(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tr", attrs, children...)
}

func Tr_(children ...HTML) HTML {
    return Tr(Attr(), children...)
}

func Tt(attrs []a.Attribute, children ...HTML) HTML {
    return Element("tt", attrs, children...)
}

func Tt_(children ...HTML) HTML {
    return Tt(Attr(), children...)
}

func U(attrs []a.Attribute, children ...HTML) HTML {
    return Element("u", attrs, children...)
}

func U_(children ...HTML) HTML {
    return U(Attr(), children...)
}

func Ul(attrs []a.Attribute, children ...HTML) HTML {
    return Element("ul", attrs, children...)
}

func Ul_(children ...HTML) HTML {
    return Ul(Attr(), children...)
}

func Var(attrs []a.Attribute, children ...HTML) HTML {
    return Element("var", attrs, children...)
}

func Var_(children ...HTML) HTML {
    return Var(Attr(), children...)
}

func Video(attrs []a.Attribute, children ...HTML) HTML {
    return Element("video", attrs, children...)
}

func Video_(children ...HTML) HTML {
    return Video(Attr(), children...)
}


// Begin of generated void elements


func Area(attrs []a.Attribute) HTML {
    return VoidElement("area", attrs)
}
func Area_() HTML {
    return Area(Attr())
}

func Base(attrs []a.Attribute) HTML {
    return VoidElement("base", attrs)
}
func Base_() HTML {
    return Base(Attr())
}

func Br(attrs []a.Attribute) HTML {
    return VoidElement("br", attrs)
}
func Br_() HTML {
    return Br(Attr())
}

func Col(attrs []a.Attribute) HTML {
    return VoidElement("col", attrs)
}
func Col_() HTML {
    return Col(Attr())
}

func Embed(attrs []a.Attribute) HTML {
    return VoidElement("embed", attrs)
}
func Embed_() HTML {
    return Embed(Attr())
}

func Hr(attrs []a.Attribute) HTML {
    return VoidElement("hr", attrs)
}
func Hr_() HTML {
    return Hr(Attr())
}

func Img(attrs []a.Attribute) HTML {
    return VoidElement("img", attrs)
}
func Img_() HTML {
    return Img(Attr())
}

func Input(attrs []a.Attribute) HTML {
    return VoidElement("input", attrs)
}
func Input_() HTML {
    return Input(Attr())
}

func Link(attrs []a.Attribute) HTML {
    return VoidElement("link", attrs)
}
func Link_() HTML {
    return Link(Attr())
}

func Meta(attrs []a.Attribute) HTML {
    return VoidElement("meta", attrs)
}
func Meta_() HTML {
    return Meta(Attr())
}

func Param(attrs []a.Attribute) HTML {
    return VoidElement("param", attrs)
}
func Param_() HTML {
    return Param(Attr())
}

func Source(attrs []a.Attribute) HTML {
    return VoidElement("source", attrs)
}
func Source_() HTML {
    return Source(Attr())
}

func Track(attrs []a.Attribute) HTML {
    return VoidElement("track", attrs)
}
func Track_() HTML {
    return Track(Attr())
}

func Wbr(attrs []a.Attribute) HTML {
    return VoidElement("wbr", attrs)
}
func Wbr_() HTML {
    return Wbr(Attr())
}

