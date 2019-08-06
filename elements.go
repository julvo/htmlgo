package htmlgo

import (
    "html/template"
    "io"
    "strings"
		"strconv"
)

type Node interface {
	BuildTemplateTo(templ *strings.Builder, vals Values, indent string)
	RenderTo(w io.Writer) error
}

type Values map[string]interface{}
type Template string
type Dataset map[string]string

func HTML(html string) template.HTML {
	return template.HTML(html)
}
func HTMLAttr(attr string) template.HTMLAttr {
	return template.HTMLAttr(attr)
}
func JS(js string) template.JS {
	return template.JS(js)
}
func JSStr(js string) template.JSStr {
	return template.JSStr(js)
}
func CSS(css string) template.CSS {
	return template.CSS(css)
}
func URL(url string) template.URL {
	return template.URL(url)
}
func Srcset(srcset string) template.Srcset {
	return template.Srcset(srcset)
}


type HTMLNode struct {
	Attributes Attr
	Tag        string
	Children   []Node
}

func (n *HTMLNode) RenderTo(w io.Writer) error {
	return RenderNodeTo(n, w)
}

func (n *HTMLNode) BuildTemplateTo(templ *strings.Builder, vals Values, indent string) {
	templ.WriteString(indent + "<" + n.Tag)
	n.Attributes.buildTemplateTo(templ, vals)
	templ.WriteString(">\n")
	for _, child := range n.Children {
		child.BuildTemplateTo(templ, vals, indent+"  ")
	}
	templ.WriteString(indent + "</" + n.Tag + ">\n")
}

type VoidHTMLNode struct {
	Attributes Attr
	Tag        string
}

func (n *VoidHTMLNode) BuildTemplateTo(templ *strings.Builder, vals Values, indent string) {
	templ.WriteString(indent + "<" + n.Tag)
	n.Attributes.buildTemplateTo(templ, vals)
	templ.WriteString("/>\n")
}
func (n *VoidHTMLNode) RenderTo(w io.Writer) error {
	return RenderNodeTo(n, w)
}

func JavaScript(s string) *TextNode {
	return Text(template.JS(s))
}

func Text(chunks ...interface{}) *TextNode {
	return &TextNode{chunks}
}

type NodeSlice struct {
	nodes []Node
}

func (ns *NodeSlice) BuildTemplateTo(templ *strings.Builder, vals Values, indent string) {
	for _, n := range ns.nodes {
		n.BuildTemplateTo(templ, vals, indent)
	}
}

func (ns *NodeSlice) RenderTo(w io.Writer) error {
	for _, n := range ns.nodes {
		if err := n.RenderTo(w); err != nil {
			return err
		}
	}
	return nil
}

func Nodes(nodes ...Node) *NodeSlice {
	ns := NodeSlice{nodes}
	return &ns
}

func (ns *NodeSlice) Append(nodes ...Node) *NodeSlice {
	ns.nodes = append(ns.nodes, nodes...)
	return ns
}

func (ns *NodeSlice) Prepend(nodes ...Node) *NodeSlice {
	ns.nodes = append(nodes, ns.nodes...)
	return ns
}

type TextNode struct {
	Chunks []interface{}
}

func (n *TextNode) BuildTemplateTo(templ *strings.Builder, vals Values, indent string) {
	templ.WriteString(indent)
	for _, chunk := range n.Chunks {
		placeholder := "P" + strconv.Itoa(len(vals))
		templ.WriteString("{{." + placeholder + "}}")
		vals[placeholder] = chunk
	}
	templ.WriteString("\n")
}

func (n *TextNode) RenderTo(w io.Writer) error {
	return RenderNodeTo(n, w)
}

func RawText(s string) *RawTextNode {
	return &RawTextNode{s}
}

type RawTextNode struct {
	Text string
}

func (n *RawTextNode) BuildTemplateTo(templ *strings.Builder, vals Values, indent string) {
	templ.WriteString(indent + n.Text + "\n")
}

func (n *RawTextNode) RenderTo(w io.Writer) error {
	return RenderNodeTo(n, w)
}

type DeclarationNode struct {
	Declaration string
}

func (n *DeclarationNode) BuildTemplateTo(templ *strings.Builder, vals Values, indent string) {
	templ.WriteString(indent + "<" + n.Declaration + ">\n")
}

func (n *DeclarationNode) RenderTo(w io.Writer) error {
	return RenderNodeTo(n, w)
}

type DocumentNode struct {
	Children []Node
}

func (d *DocumentNode) RenderTo(w io.Writer) error {
	return RenderNodeTo(d, w)
}

func (d *DocumentNode) BuildTemplateTo(templ *strings.Builder, vals Values, indent string) {
	for _, child := range d.Children {
		child.BuildTemplateTo(templ, vals, indent)
	}
}

func RenderNodeTo(n Node, w io.Writer) error {
	templBuilder := strings.Builder{}
	vals := Values{}

	n.BuildTemplateTo(&templBuilder, vals, "")

	if len(vals) == 0 {
		_, err := w.Write([]byte(templBuilder.String()))
		return err
	}

	templ, err := template.New("node").Parse(templBuilder.String())
	if err != nil {
		return err
	}
	return templ.Execute(w, vals)
}

func Document(children ...Node) *DocumentNode {
	return &DocumentNode{children}
}

func Doctype(t string) *DeclarationNode {
	return &DeclarationNode{
		Declaration: "!DOCTYPE " + t,
	}
}

type Attr struct {
	Dataset   Dataset
	DisabledBoolean  bool
	templData map[string]interface{}

	Accept string

	AcceptCharset string

	Accesskey string

	Action string

	Align string

	Alt string

	AriaExpanded string

	AriaHidden string

	AriaLabel string

	Async string

	Autocomplete string

	Autofocus string

	Autoplay string

	Bgcolor string

	Border string

	Charset string

	Checked string

	Cite string

	Class string

	Color string

	Cols string

	Colspan string

	Content string

	Contenteditable string

	Controls string

	Coords string

	Data string

	Datetime string

	Default string

	Defer string

	Dir string

	Dirname string

	Disabled string

	Download string

	Draggable string

	Dropzone string

	Enctype string

	For string

	Form string

	Formaction string

	Headers string

	Height string

	Hidden string

	High string

	Href string

	Hreflang string

	HttpEquiv string

	Id string

	InitialScale string

	Ismap string

	Kind string

	Label string

	Lang string

	List string

	Loop string

	Low string

	Max string

	Maxlength string

	Media string

	Method string

	Min string

	Multiple string

	Muted string

	Name string

	Novalidate string

	Onabort string

	Onafterprint string

	Onbeforeprint string

	Onbeforeunload string

	Onblur string

	Oncanplay string

	Oncanplaythrough string

	Onchange string

	Onclick string

	Oncontextmenu string

	Oncopy string

	Oncuechange string

	Oncut string

	Ondblclick string

	Ondrag string

	Ondragend string

	Ondragenter string

	Ondragleave string

	Ondragover string

	Ondragstart string

	Ondrop string

	Ondurationchange string

	Onemptied string

	Onended string

	Onerror string

	Onfocus string

	Onhashchange string

	Oninput string

	Oninvalid string

	Onkeydown string

	Onkeypress string

	Onkeyup string

	Onload string

	Onloadeddata string

	Onloadedmetadata string

	Onloadstart string

	Onmousedown string

	Onmousemove string

	Onmouseout string

	Onmouseover string

	Onmouseup string

	Onmousewheel string

	Onoffline string

	Ononline string

	Onpagehide string

	Onpageshow string

	Onpaste string

	Onpause string

	Onplay string

	Onplaying string

	Onpopstate string

	Onprogress string

	Onratechange string

	Onreset string

	Onresize string

	Onscroll string

	Onsearch string

	Onseeked string

	Onseeking string

	Onselect string

	Onstalled string

	Onstorage string

	Onsubmit string

	Onsuspend string

	Ontimeupdate string

	Ontoggle string

	Onunload string

	Onvolumechange string

	Onwaiting string

	Onwheel string

	Open string

	Optimum string

	Pattern string

	Placeholder string

	Poster string

	Preload string

	Readonly string

	Rel string

	Required string

	Reversed string

	Role string

	Rows string

	Rowspan string

	Sandbox string

	Scope string

	Selected string

	Shape string

	Size string

	Sizes string

	Span string

	Spellcheck string

	Src string

	Srcdoc string

	Srclang string

	Srcset string

	Start string

	Step string

	Style string

	Tabindex string

	Target string

	Title string

	Translate string

	Type string

	Usemap string

	Value string

	Width string

	Wrap string

}

func (a Attr) Bind(key string, value interface{}) Attr {
	if a.templData == nil {
		a.templData = make(map[string]interface{})
	}
	a.templData[key] = value
	return a
}

func (a Attr) Bind_(value interface{}) Attr {
	return a.Bind("", value)
}

func (a *Attr) buildTemplateTo(templ *strings.Builder, vals Values) {
	if len(a.templData) > 0 {
		placeholder := "P" + strconv.Itoa(len(vals))

		a.Accept = strings.Replace(a.Accept, "{{.", "{{."+placeholder, -1)

		a.AcceptCharset = strings.Replace(a.AcceptCharset, "{{.", "{{."+placeholder, -1)

		a.Accesskey = strings.Replace(a.Accesskey, "{{.", "{{."+placeholder, -1)

		a.Action = strings.Replace(a.Action, "{{.", "{{."+placeholder, -1)

		a.Align = strings.Replace(a.Align, "{{.", "{{."+placeholder, -1)

		a.Alt = strings.Replace(a.Alt, "{{.", "{{."+placeholder, -1)

		a.AriaExpanded = strings.Replace(a.AriaExpanded, "{{.", "{{."+placeholder, -1)

		a.AriaHidden = strings.Replace(a.AriaHidden, "{{.", "{{."+placeholder, -1)

		a.AriaLabel = strings.Replace(a.AriaLabel, "{{.", "{{."+placeholder, -1)

		a.Async = strings.Replace(a.Async, "{{.", "{{."+placeholder, -1)

		a.Autocomplete = strings.Replace(a.Autocomplete, "{{.", "{{."+placeholder, -1)

		a.Autofocus = strings.Replace(a.Autofocus, "{{.", "{{."+placeholder, -1)

		a.Autoplay = strings.Replace(a.Autoplay, "{{.", "{{."+placeholder, -1)

		a.Bgcolor = strings.Replace(a.Bgcolor, "{{.", "{{."+placeholder, -1)

		a.Border = strings.Replace(a.Border, "{{.", "{{."+placeholder, -1)

		a.Charset = strings.Replace(a.Charset, "{{.", "{{."+placeholder, -1)

		a.Checked = strings.Replace(a.Checked, "{{.", "{{."+placeholder, -1)

		a.Cite = strings.Replace(a.Cite, "{{.", "{{."+placeholder, -1)

		a.Class = strings.Replace(a.Class, "{{.", "{{."+placeholder, -1)

		a.Color = strings.Replace(a.Color, "{{.", "{{."+placeholder, -1)

		a.Cols = strings.Replace(a.Cols, "{{.", "{{."+placeholder, -1)

		a.Colspan = strings.Replace(a.Colspan, "{{.", "{{."+placeholder, -1)

		a.Content = strings.Replace(a.Content, "{{.", "{{."+placeholder, -1)

		a.Contenteditable = strings.Replace(a.Contenteditable, "{{.", "{{."+placeholder, -1)

		a.Controls = strings.Replace(a.Controls, "{{.", "{{."+placeholder, -1)

		a.Coords = strings.Replace(a.Coords, "{{.", "{{."+placeholder, -1)

		a.Data = strings.Replace(a.Data, "{{.", "{{."+placeholder, -1)

		a.Datetime = strings.Replace(a.Datetime, "{{.", "{{."+placeholder, -1)

		a.Default = strings.Replace(a.Default, "{{.", "{{."+placeholder, -1)

		a.Defer = strings.Replace(a.Defer, "{{.", "{{."+placeholder, -1)

		a.Dir = strings.Replace(a.Dir, "{{.", "{{."+placeholder, -1)

		a.Dirname = strings.Replace(a.Dirname, "{{.", "{{."+placeholder, -1)

		a.Disabled = strings.Replace(a.Disabled, "{{.", "{{."+placeholder, -1)

		a.Download = strings.Replace(a.Download, "{{.", "{{."+placeholder, -1)

		a.Draggable = strings.Replace(a.Draggable, "{{.", "{{."+placeholder, -1)

		a.Dropzone = strings.Replace(a.Dropzone, "{{.", "{{."+placeholder, -1)

		a.Enctype = strings.Replace(a.Enctype, "{{.", "{{."+placeholder, -1)

		a.For = strings.Replace(a.For, "{{.", "{{."+placeholder, -1)

		a.Form = strings.Replace(a.Form, "{{.", "{{."+placeholder, -1)

		a.Formaction = strings.Replace(a.Formaction, "{{.", "{{."+placeholder, -1)

		a.Headers = strings.Replace(a.Headers, "{{.", "{{."+placeholder, -1)

		a.Height = strings.Replace(a.Height, "{{.", "{{."+placeholder, -1)

		a.Hidden = strings.Replace(a.Hidden, "{{.", "{{."+placeholder, -1)

		a.High = strings.Replace(a.High, "{{.", "{{."+placeholder, -1)

		a.Href = strings.Replace(a.Href, "{{.", "{{."+placeholder, -1)

		a.Hreflang = strings.Replace(a.Hreflang, "{{.", "{{."+placeholder, -1)

		a.HttpEquiv = strings.Replace(a.HttpEquiv, "{{.", "{{."+placeholder, -1)

		a.Id = strings.Replace(a.Id, "{{.", "{{."+placeholder, -1)

		a.InitialScale = strings.Replace(a.InitialScale, "{{.", "{{."+placeholder, -1)

		a.Ismap = strings.Replace(a.Ismap, "{{.", "{{."+placeholder, -1)

		a.Kind = strings.Replace(a.Kind, "{{.", "{{."+placeholder, -1)

		a.Label = strings.Replace(a.Label, "{{.", "{{."+placeholder, -1)

		a.Lang = strings.Replace(a.Lang, "{{.", "{{."+placeholder, -1)

		a.List = strings.Replace(a.List, "{{.", "{{."+placeholder, -1)

		a.Loop = strings.Replace(a.Loop, "{{.", "{{."+placeholder, -1)

		a.Low = strings.Replace(a.Low, "{{.", "{{."+placeholder, -1)

		a.Max = strings.Replace(a.Max, "{{.", "{{."+placeholder, -1)

		a.Maxlength = strings.Replace(a.Maxlength, "{{.", "{{."+placeholder, -1)

		a.Media = strings.Replace(a.Media, "{{.", "{{."+placeholder, -1)

		a.Method = strings.Replace(a.Method, "{{.", "{{."+placeholder, -1)

		a.Min = strings.Replace(a.Min, "{{.", "{{."+placeholder, -1)

		a.Multiple = strings.Replace(a.Multiple, "{{.", "{{."+placeholder, -1)

		a.Muted = strings.Replace(a.Muted, "{{.", "{{."+placeholder, -1)

		a.Name = strings.Replace(a.Name, "{{.", "{{."+placeholder, -1)

		a.Novalidate = strings.Replace(a.Novalidate, "{{.", "{{."+placeholder, -1)

		a.Onabort = strings.Replace(a.Onabort, "{{.", "{{."+placeholder, -1)

		a.Onafterprint = strings.Replace(a.Onafterprint, "{{.", "{{."+placeholder, -1)

		a.Onbeforeprint = strings.Replace(a.Onbeforeprint, "{{.", "{{."+placeholder, -1)

		a.Onbeforeunload = strings.Replace(a.Onbeforeunload, "{{.", "{{."+placeholder, -1)

		a.Onblur = strings.Replace(a.Onblur, "{{.", "{{."+placeholder, -1)

		a.Oncanplay = strings.Replace(a.Oncanplay, "{{.", "{{."+placeholder, -1)

		a.Oncanplaythrough = strings.Replace(a.Oncanplaythrough, "{{.", "{{."+placeholder, -1)

		a.Onchange = strings.Replace(a.Onchange, "{{.", "{{."+placeholder, -1)

		a.Onclick = strings.Replace(a.Onclick, "{{.", "{{."+placeholder, -1)

		a.Oncontextmenu = strings.Replace(a.Oncontextmenu, "{{.", "{{."+placeholder, -1)

		a.Oncopy = strings.Replace(a.Oncopy, "{{.", "{{."+placeholder, -1)

		a.Oncuechange = strings.Replace(a.Oncuechange, "{{.", "{{."+placeholder, -1)

		a.Oncut = strings.Replace(a.Oncut, "{{.", "{{."+placeholder, -1)

		a.Ondblclick = strings.Replace(a.Ondblclick, "{{.", "{{."+placeholder, -1)

		a.Ondrag = strings.Replace(a.Ondrag, "{{.", "{{."+placeholder, -1)

		a.Ondragend = strings.Replace(a.Ondragend, "{{.", "{{."+placeholder, -1)

		a.Ondragenter = strings.Replace(a.Ondragenter, "{{.", "{{."+placeholder, -1)

		a.Ondragleave = strings.Replace(a.Ondragleave, "{{.", "{{."+placeholder, -1)

		a.Ondragover = strings.Replace(a.Ondragover, "{{.", "{{."+placeholder, -1)

		a.Ondragstart = strings.Replace(a.Ondragstart, "{{.", "{{."+placeholder, -1)

		a.Ondrop = strings.Replace(a.Ondrop, "{{.", "{{."+placeholder, -1)

		a.Ondurationchange = strings.Replace(a.Ondurationchange, "{{.", "{{."+placeholder, -1)

		a.Onemptied = strings.Replace(a.Onemptied, "{{.", "{{."+placeholder, -1)

		a.Onended = strings.Replace(a.Onended, "{{.", "{{."+placeholder, -1)

		a.Onerror = strings.Replace(a.Onerror, "{{.", "{{."+placeholder, -1)

		a.Onfocus = strings.Replace(a.Onfocus, "{{.", "{{."+placeholder, -1)

		a.Onhashchange = strings.Replace(a.Onhashchange, "{{.", "{{."+placeholder, -1)

		a.Oninput = strings.Replace(a.Oninput, "{{.", "{{."+placeholder, -1)

		a.Oninvalid = strings.Replace(a.Oninvalid, "{{.", "{{."+placeholder, -1)

		a.Onkeydown = strings.Replace(a.Onkeydown, "{{.", "{{."+placeholder, -1)

		a.Onkeypress = strings.Replace(a.Onkeypress, "{{.", "{{."+placeholder, -1)

		a.Onkeyup = strings.Replace(a.Onkeyup, "{{.", "{{."+placeholder, -1)

		a.Onload = strings.Replace(a.Onload, "{{.", "{{."+placeholder, -1)

		a.Onloadeddata = strings.Replace(a.Onloadeddata, "{{.", "{{."+placeholder, -1)

		a.Onloadedmetadata = strings.Replace(a.Onloadedmetadata, "{{.", "{{."+placeholder, -1)

		a.Onloadstart = strings.Replace(a.Onloadstart, "{{.", "{{."+placeholder, -1)

		a.Onmousedown = strings.Replace(a.Onmousedown, "{{.", "{{."+placeholder, -1)

		a.Onmousemove = strings.Replace(a.Onmousemove, "{{.", "{{."+placeholder, -1)

		a.Onmouseout = strings.Replace(a.Onmouseout, "{{.", "{{."+placeholder, -1)

		a.Onmouseover = strings.Replace(a.Onmouseover, "{{.", "{{."+placeholder, -1)

		a.Onmouseup = strings.Replace(a.Onmouseup, "{{.", "{{."+placeholder, -1)

		a.Onmousewheel = strings.Replace(a.Onmousewheel, "{{.", "{{."+placeholder, -1)

		a.Onoffline = strings.Replace(a.Onoffline, "{{.", "{{."+placeholder, -1)

		a.Ononline = strings.Replace(a.Ononline, "{{.", "{{."+placeholder, -1)

		a.Onpagehide = strings.Replace(a.Onpagehide, "{{.", "{{."+placeholder, -1)

		a.Onpageshow = strings.Replace(a.Onpageshow, "{{.", "{{."+placeholder, -1)

		a.Onpaste = strings.Replace(a.Onpaste, "{{.", "{{."+placeholder, -1)

		a.Onpause = strings.Replace(a.Onpause, "{{.", "{{."+placeholder, -1)

		a.Onplay = strings.Replace(a.Onplay, "{{.", "{{."+placeholder, -1)

		a.Onplaying = strings.Replace(a.Onplaying, "{{.", "{{."+placeholder, -1)

		a.Onpopstate = strings.Replace(a.Onpopstate, "{{.", "{{."+placeholder, -1)

		a.Onprogress = strings.Replace(a.Onprogress, "{{.", "{{."+placeholder, -1)

		a.Onratechange = strings.Replace(a.Onratechange, "{{.", "{{."+placeholder, -1)

		a.Onreset = strings.Replace(a.Onreset, "{{.", "{{."+placeholder, -1)

		a.Onresize = strings.Replace(a.Onresize, "{{.", "{{."+placeholder, -1)

		a.Onscroll = strings.Replace(a.Onscroll, "{{.", "{{."+placeholder, -1)

		a.Onsearch = strings.Replace(a.Onsearch, "{{.", "{{."+placeholder, -1)

		a.Onseeked = strings.Replace(a.Onseeked, "{{.", "{{."+placeholder, -1)

		a.Onseeking = strings.Replace(a.Onseeking, "{{.", "{{."+placeholder, -1)

		a.Onselect = strings.Replace(a.Onselect, "{{.", "{{."+placeholder, -1)

		a.Onstalled = strings.Replace(a.Onstalled, "{{.", "{{."+placeholder, -1)

		a.Onstorage = strings.Replace(a.Onstorage, "{{.", "{{."+placeholder, -1)

		a.Onsubmit = strings.Replace(a.Onsubmit, "{{.", "{{."+placeholder, -1)

		a.Onsuspend = strings.Replace(a.Onsuspend, "{{.", "{{."+placeholder, -1)

		a.Ontimeupdate = strings.Replace(a.Ontimeupdate, "{{.", "{{."+placeholder, -1)

		a.Ontoggle = strings.Replace(a.Ontoggle, "{{.", "{{."+placeholder, -1)

		a.Onunload = strings.Replace(a.Onunload, "{{.", "{{."+placeholder, -1)

		a.Onvolumechange = strings.Replace(a.Onvolumechange, "{{.", "{{."+placeholder, -1)

		a.Onwaiting = strings.Replace(a.Onwaiting, "{{.", "{{."+placeholder, -1)

		a.Onwheel = strings.Replace(a.Onwheel, "{{.", "{{."+placeholder, -1)

		a.Open = strings.Replace(a.Open, "{{.", "{{."+placeholder, -1)

		a.Optimum = strings.Replace(a.Optimum, "{{.", "{{."+placeholder, -1)

		a.Pattern = strings.Replace(a.Pattern, "{{.", "{{."+placeholder, -1)

		a.Placeholder = strings.Replace(a.Placeholder, "{{.", "{{."+placeholder, -1)

		a.Poster = strings.Replace(a.Poster, "{{.", "{{."+placeholder, -1)

		a.Preload = strings.Replace(a.Preload, "{{.", "{{."+placeholder, -1)

		a.Readonly = strings.Replace(a.Readonly, "{{.", "{{."+placeholder, -1)

		a.Rel = strings.Replace(a.Rel, "{{.", "{{."+placeholder, -1)

		a.Required = strings.Replace(a.Required, "{{.", "{{."+placeholder, -1)

		a.Reversed = strings.Replace(a.Reversed, "{{.", "{{."+placeholder, -1)

		a.Role = strings.Replace(a.Role, "{{.", "{{."+placeholder, -1)

		a.Rows = strings.Replace(a.Rows, "{{.", "{{."+placeholder, -1)

		a.Rowspan = strings.Replace(a.Rowspan, "{{.", "{{."+placeholder, -1)

		a.Sandbox = strings.Replace(a.Sandbox, "{{.", "{{."+placeholder, -1)

		a.Scope = strings.Replace(a.Scope, "{{.", "{{."+placeholder, -1)

		a.Selected = strings.Replace(a.Selected, "{{.", "{{."+placeholder, -1)

		a.Shape = strings.Replace(a.Shape, "{{.", "{{."+placeholder, -1)

		a.Size = strings.Replace(a.Size, "{{.", "{{."+placeholder, -1)

		a.Sizes = strings.Replace(a.Sizes, "{{.", "{{."+placeholder, -1)

		a.Span = strings.Replace(a.Span, "{{.", "{{."+placeholder, -1)

		a.Spellcheck = strings.Replace(a.Spellcheck, "{{.", "{{."+placeholder, -1)

		a.Src = strings.Replace(a.Src, "{{.", "{{."+placeholder, -1)

		a.Srcdoc = strings.Replace(a.Srcdoc, "{{.", "{{."+placeholder, -1)

		a.Srclang = strings.Replace(a.Srclang, "{{.", "{{."+placeholder, -1)

		a.Srcset = strings.Replace(a.Srcset, "{{.", "{{."+placeholder, -1)

		a.Start = strings.Replace(a.Start, "{{.", "{{."+placeholder, -1)

		a.Step = strings.Replace(a.Step, "{{.", "{{."+placeholder, -1)

		a.Style = strings.Replace(a.Style, "{{.", "{{."+placeholder, -1)

		a.Tabindex = strings.Replace(a.Tabindex, "{{.", "{{."+placeholder, -1)

		a.Target = strings.Replace(a.Target, "{{.", "{{."+placeholder, -1)

		a.Title = strings.Replace(a.Title, "{{.", "{{."+placeholder, -1)

		a.Translate = strings.Replace(a.Translate, "{{.", "{{."+placeholder, -1)

		a.Type = strings.Replace(a.Type, "{{.", "{{."+placeholder, -1)

		a.Usemap = strings.Replace(a.Usemap, "{{.", "{{."+placeholder, -1)

		a.Value = strings.Replace(a.Value, "{{.", "{{."+placeholder, -1)

		a.Width = strings.Replace(a.Width, "{{.", "{{."+placeholder, -1)

		a.Wrap = strings.Replace(a.Wrap, "{{.", "{{."+placeholder, -1)


		for k, v := range a.templData {
			vals[placeholder+k] = v
		}
	}

	switch {

    case a.Accept != "":
		templ.WriteString(` accept="` + a.Accept + `"`)

    case a.AcceptCharset != "":
		templ.WriteString(` accept-charset="` + a.AcceptCharset + `"`)

    case a.Accesskey != "":
		templ.WriteString(` accesskey="` + a.Accesskey + `"`)

    case a.Action != "":
		templ.WriteString(` action="` + a.Action + `"`)

    case a.Align != "":
		templ.WriteString(` align="` + a.Align + `"`)

    case a.Alt != "":
		templ.WriteString(` alt="` + a.Alt + `"`)

    case a.AriaExpanded != "":
		templ.WriteString(` aria-expanded="` + a.AriaExpanded + `"`)

    case a.AriaHidden != "":
		templ.WriteString(` aria-hidden="` + a.AriaHidden + `"`)

    case a.AriaLabel != "":
		templ.WriteString(` aria-label="` + a.AriaLabel + `"`)

    case a.Async != "":
		templ.WriteString(` async="` + a.Async + `"`)

    case a.Autocomplete != "":
		templ.WriteString(` autocomplete="` + a.Autocomplete + `"`)

    case a.Autofocus != "":
		templ.WriteString(` autofocus="` + a.Autofocus + `"`)

    case a.Autoplay != "":
		templ.WriteString(` autoplay="` + a.Autoplay + `"`)

    case a.Bgcolor != "":
		templ.WriteString(` bgcolor="` + a.Bgcolor + `"`)

    case a.Border != "":
		templ.WriteString(` border="` + a.Border + `"`)

    case a.Charset != "":
		templ.WriteString(` charset="` + a.Charset + `"`)

    case a.Checked != "":
		templ.WriteString(` checked="` + a.Checked + `"`)

    case a.Cite != "":
		templ.WriteString(` cite="` + a.Cite + `"`)

    case a.Class != "":
		templ.WriteString(` class="` + a.Class + `"`)

    case a.Color != "":
		templ.WriteString(` color="` + a.Color + `"`)

    case a.Cols != "":
		templ.WriteString(` cols="` + a.Cols + `"`)

    case a.Colspan != "":
		templ.WriteString(` colspan="` + a.Colspan + `"`)

    case a.Content != "":
		templ.WriteString(` content="` + a.Content + `"`)

    case a.Contenteditable != "":
		templ.WriteString(` contenteditable="` + a.Contenteditable + `"`)

    case a.Controls != "":
		templ.WriteString(` controls="` + a.Controls + `"`)

    case a.Coords != "":
		templ.WriteString(` coords="` + a.Coords + `"`)

    case a.Data != "":
		templ.WriteString(` data="` + a.Data + `"`)

    case a.Datetime != "":
		templ.WriteString(` datetime="` + a.Datetime + `"`)

    case a.Default != "":
		templ.WriteString(` default="` + a.Default + `"`)

    case a.Defer != "":
		templ.WriteString(` defer="` + a.Defer + `"`)

    case a.Dir != "":
		templ.WriteString(` dir="` + a.Dir + `"`)

    case a.Dirname != "":
		templ.WriteString(` dirname="` + a.Dirname + `"`)

    case a.Disabled != "":
		templ.WriteString(` disabled="` + a.Disabled + `"`)

    case a.Download != "":
		templ.WriteString(` download="` + a.Download + `"`)

    case a.Draggable != "":
		templ.WriteString(` draggable="` + a.Draggable + `"`)

    case a.Dropzone != "":
		templ.WriteString(` dropzone="` + a.Dropzone + `"`)

    case a.Enctype != "":
		templ.WriteString(` enctype="` + a.Enctype + `"`)

    case a.For != "":
		templ.WriteString(` for="` + a.For + `"`)

    case a.Form != "":
		templ.WriteString(` form="` + a.Form + `"`)

    case a.Formaction != "":
		templ.WriteString(` formaction="` + a.Formaction + `"`)

    case a.Headers != "":
		templ.WriteString(` headers="` + a.Headers + `"`)

    case a.Height != "":
		templ.WriteString(` height="` + a.Height + `"`)

    case a.Hidden != "":
		templ.WriteString(` hidden="` + a.Hidden + `"`)

    case a.High != "":
		templ.WriteString(` high="` + a.High + `"`)

    case a.Href != "":
		templ.WriteString(` href="` + a.Href + `"`)

    case a.Hreflang != "":
		templ.WriteString(` hreflang="` + a.Hreflang + `"`)

    case a.HttpEquiv != "":
		templ.WriteString(` http-equiv="` + a.HttpEquiv + `"`)

    case a.Id != "":
		templ.WriteString(` id="` + a.Id + `"`)

    case a.InitialScale != "":
		templ.WriteString(` initial-scale="` + a.InitialScale + `"`)

    case a.Ismap != "":
		templ.WriteString(` ismap="` + a.Ismap + `"`)

    case a.Kind != "":
		templ.WriteString(` kind="` + a.Kind + `"`)

    case a.Label != "":
		templ.WriteString(` label="` + a.Label + `"`)

    case a.Lang != "":
		templ.WriteString(` lang="` + a.Lang + `"`)

    case a.List != "":
		templ.WriteString(` list="` + a.List + `"`)

    case a.Loop != "":
		templ.WriteString(` loop="` + a.Loop + `"`)

    case a.Low != "":
		templ.WriteString(` low="` + a.Low + `"`)

    case a.Max != "":
		templ.WriteString(` max="` + a.Max + `"`)

    case a.Maxlength != "":
		templ.WriteString(` maxlength="` + a.Maxlength + `"`)

    case a.Media != "":
		templ.WriteString(` media="` + a.Media + `"`)

    case a.Method != "":
		templ.WriteString(` method="` + a.Method + `"`)

    case a.Min != "":
		templ.WriteString(` min="` + a.Min + `"`)

    case a.Multiple != "":
		templ.WriteString(` multiple="` + a.Multiple + `"`)

    case a.Muted != "":
		templ.WriteString(` muted="` + a.Muted + `"`)

    case a.Name != "":
		templ.WriteString(` name="` + a.Name + `"`)

    case a.Novalidate != "":
		templ.WriteString(` novalidate="` + a.Novalidate + `"`)

    case a.Onabort != "":
		templ.WriteString(` onabort="` + a.Onabort + `"`)

    case a.Onafterprint != "":
		templ.WriteString(` onafterprint="` + a.Onafterprint + `"`)

    case a.Onbeforeprint != "":
		templ.WriteString(` onbeforeprint="` + a.Onbeforeprint + `"`)

    case a.Onbeforeunload != "":
		templ.WriteString(` onbeforeunload="` + a.Onbeforeunload + `"`)

    case a.Onblur != "":
		templ.WriteString(` onblur="` + a.Onblur + `"`)

    case a.Oncanplay != "":
		templ.WriteString(` oncanplay="` + a.Oncanplay + `"`)

    case a.Oncanplaythrough != "":
		templ.WriteString(` oncanplaythrough="` + a.Oncanplaythrough + `"`)

    case a.Onchange != "":
		templ.WriteString(` onchange="` + a.Onchange + `"`)

    case a.Onclick != "":
		templ.WriteString(` onclick="` + a.Onclick + `"`)

    case a.Oncontextmenu != "":
		templ.WriteString(` oncontextmenu="` + a.Oncontextmenu + `"`)

    case a.Oncopy != "":
		templ.WriteString(` oncopy="` + a.Oncopy + `"`)

    case a.Oncuechange != "":
		templ.WriteString(` oncuechange="` + a.Oncuechange + `"`)

    case a.Oncut != "":
		templ.WriteString(` oncut="` + a.Oncut + `"`)

    case a.Ondblclick != "":
		templ.WriteString(` ondblclick="` + a.Ondblclick + `"`)

    case a.Ondrag != "":
		templ.WriteString(` ondrag="` + a.Ondrag + `"`)

    case a.Ondragend != "":
		templ.WriteString(` ondragend="` + a.Ondragend + `"`)

    case a.Ondragenter != "":
		templ.WriteString(` ondragenter="` + a.Ondragenter + `"`)

    case a.Ondragleave != "":
		templ.WriteString(` ondragleave="` + a.Ondragleave + `"`)

    case a.Ondragover != "":
		templ.WriteString(` ondragover="` + a.Ondragover + `"`)

    case a.Ondragstart != "":
		templ.WriteString(` ondragstart="` + a.Ondragstart + `"`)

    case a.Ondrop != "":
		templ.WriteString(` ondrop="` + a.Ondrop + `"`)

    case a.Ondurationchange != "":
		templ.WriteString(` ondurationchange="` + a.Ondurationchange + `"`)

    case a.Onemptied != "":
		templ.WriteString(` onemptied="` + a.Onemptied + `"`)

    case a.Onended != "":
		templ.WriteString(` onended="` + a.Onended + `"`)

    case a.Onerror != "":
		templ.WriteString(` onerror="` + a.Onerror + `"`)

    case a.Onfocus != "":
		templ.WriteString(` onfocus="` + a.Onfocus + `"`)

    case a.Onhashchange != "":
		templ.WriteString(` onhashchange="` + a.Onhashchange + `"`)

    case a.Oninput != "":
		templ.WriteString(` oninput="` + a.Oninput + `"`)

    case a.Oninvalid != "":
		templ.WriteString(` oninvalid="` + a.Oninvalid + `"`)

    case a.Onkeydown != "":
		templ.WriteString(` onkeydown="` + a.Onkeydown + `"`)

    case a.Onkeypress != "":
		templ.WriteString(` onkeypress="` + a.Onkeypress + `"`)

    case a.Onkeyup != "":
		templ.WriteString(` onkeyup="` + a.Onkeyup + `"`)

    case a.Onload != "":
		templ.WriteString(` onload="` + a.Onload + `"`)

    case a.Onloadeddata != "":
		templ.WriteString(` onloadeddata="` + a.Onloadeddata + `"`)

    case a.Onloadedmetadata != "":
		templ.WriteString(` onloadedmetadata="` + a.Onloadedmetadata + `"`)

    case a.Onloadstart != "":
		templ.WriteString(` onloadstart="` + a.Onloadstart + `"`)

    case a.Onmousedown != "":
		templ.WriteString(` onmousedown="` + a.Onmousedown + `"`)

    case a.Onmousemove != "":
		templ.WriteString(` onmousemove="` + a.Onmousemove + `"`)

    case a.Onmouseout != "":
		templ.WriteString(` onmouseout="` + a.Onmouseout + `"`)

    case a.Onmouseover != "":
		templ.WriteString(` onmouseover="` + a.Onmouseover + `"`)

    case a.Onmouseup != "":
		templ.WriteString(` onmouseup="` + a.Onmouseup + `"`)

    case a.Onmousewheel != "":
		templ.WriteString(` onmousewheel="` + a.Onmousewheel + `"`)

    case a.Onoffline != "":
		templ.WriteString(` onoffline="` + a.Onoffline + `"`)

    case a.Ononline != "":
		templ.WriteString(` ononline="` + a.Ononline + `"`)

    case a.Onpagehide != "":
		templ.WriteString(` onpagehide="` + a.Onpagehide + `"`)

    case a.Onpageshow != "":
		templ.WriteString(` onpageshow="` + a.Onpageshow + `"`)

    case a.Onpaste != "":
		templ.WriteString(` onpaste="` + a.Onpaste + `"`)

    case a.Onpause != "":
		templ.WriteString(` onpause="` + a.Onpause + `"`)

    case a.Onplay != "":
		templ.WriteString(` onplay="` + a.Onplay + `"`)

    case a.Onplaying != "":
		templ.WriteString(` onplaying="` + a.Onplaying + `"`)

    case a.Onpopstate != "":
		templ.WriteString(` onpopstate="` + a.Onpopstate + `"`)

    case a.Onprogress != "":
		templ.WriteString(` onprogress="` + a.Onprogress + `"`)

    case a.Onratechange != "":
		templ.WriteString(` onratechange="` + a.Onratechange + `"`)

    case a.Onreset != "":
		templ.WriteString(` onreset="` + a.Onreset + `"`)

    case a.Onresize != "":
		templ.WriteString(` onresize="` + a.Onresize + `"`)

    case a.Onscroll != "":
		templ.WriteString(` onscroll="` + a.Onscroll + `"`)

    case a.Onsearch != "":
		templ.WriteString(` onsearch="` + a.Onsearch + `"`)

    case a.Onseeked != "":
		templ.WriteString(` onseeked="` + a.Onseeked + `"`)

    case a.Onseeking != "":
		templ.WriteString(` onseeking="` + a.Onseeking + `"`)

    case a.Onselect != "":
		templ.WriteString(` onselect="` + a.Onselect + `"`)

    case a.Onstalled != "":
		templ.WriteString(` onstalled="` + a.Onstalled + `"`)

    case a.Onstorage != "":
		templ.WriteString(` onstorage="` + a.Onstorage + `"`)

    case a.Onsubmit != "":
		templ.WriteString(` onsubmit="` + a.Onsubmit + `"`)

    case a.Onsuspend != "":
		templ.WriteString(` onsuspend="` + a.Onsuspend + `"`)

    case a.Ontimeupdate != "":
		templ.WriteString(` ontimeupdate="` + a.Ontimeupdate + `"`)

    case a.Ontoggle != "":
		templ.WriteString(` ontoggle="` + a.Ontoggle + `"`)

    case a.Onunload != "":
		templ.WriteString(` onunload="` + a.Onunload + `"`)

    case a.Onvolumechange != "":
		templ.WriteString(` onvolumechange="` + a.Onvolumechange + `"`)

    case a.Onwaiting != "":
		templ.WriteString(` onwaiting="` + a.Onwaiting + `"`)

    case a.Onwheel != "":
		templ.WriteString(` onwheel="` + a.Onwheel + `"`)

    case a.Open != "":
		templ.WriteString(` open="` + a.Open + `"`)

    case a.Optimum != "":
		templ.WriteString(` optimum="` + a.Optimum + `"`)

    case a.Pattern != "":
		templ.WriteString(` pattern="` + a.Pattern + `"`)

    case a.Placeholder != "":
		templ.WriteString(` placeholder="` + a.Placeholder + `"`)

    case a.Poster != "":
		templ.WriteString(` poster="` + a.Poster + `"`)

    case a.Preload != "":
		templ.WriteString(` preload="` + a.Preload + `"`)

    case a.Readonly != "":
		templ.WriteString(` readonly="` + a.Readonly + `"`)

    case a.Rel != "":
		templ.WriteString(` rel="` + a.Rel + `"`)

    case a.Required != "":
		templ.WriteString(` required="` + a.Required + `"`)

    case a.Reversed != "":
		templ.WriteString(` reversed="` + a.Reversed + `"`)

    case a.Role != "":
		templ.WriteString(` role="` + a.Role + `"`)

    case a.Rows != "":
		templ.WriteString(` rows="` + a.Rows + `"`)

    case a.Rowspan != "":
		templ.WriteString(` rowspan="` + a.Rowspan + `"`)

    case a.Sandbox != "":
		templ.WriteString(` sandbox="` + a.Sandbox + `"`)

    case a.Scope != "":
		templ.WriteString(` scope="` + a.Scope + `"`)

    case a.Selected != "":
		templ.WriteString(` selected="` + a.Selected + `"`)

    case a.Shape != "":
		templ.WriteString(` shape="` + a.Shape + `"`)

    case a.Size != "":
		templ.WriteString(` size="` + a.Size + `"`)

    case a.Sizes != "":
		templ.WriteString(` sizes="` + a.Sizes + `"`)

    case a.Span != "":
		templ.WriteString(` span="` + a.Span + `"`)

    case a.Spellcheck != "":
		templ.WriteString(` spellcheck="` + a.Spellcheck + `"`)

    case a.Src != "":
		templ.WriteString(` src="` + a.Src + `"`)

    case a.Srcdoc != "":
		templ.WriteString(` srcdoc="` + a.Srcdoc + `"`)

    case a.Srclang != "":
		templ.WriteString(` srclang="` + a.Srclang + `"`)

    case a.Srcset != "":
		templ.WriteString(` srcset="` + a.Srcset + `"`)

    case a.Start != "":
		templ.WriteString(` start="` + a.Start + `"`)

    case a.Step != "":
		templ.WriteString(` step="` + a.Step + `"`)

    case a.Style != "":
		templ.WriteString(` style="` + a.Style + `"`)

    case a.Tabindex != "":
		templ.WriteString(` tabindex="` + a.Tabindex + `"`)

    case a.Target != "":
		templ.WriteString(` target="` + a.Target + `"`)

    case a.Title != "":
		templ.WriteString(` title="` + a.Title + `"`)

    case a.Translate != "":
		templ.WriteString(` translate="` + a.Translate + `"`)

    case a.Type != "":
		templ.WriteString(` type="` + a.Type + `"`)

    case a.Usemap != "":
		templ.WriteString(` usemap="` + a.Usemap + `"`)

    case a.Value != "":
		templ.WriteString(` value="` + a.Value + `"`)

    case a.Width != "":
		templ.WriteString(` width="` + a.Width + `"`)

    case a.Wrap != "":
		templ.WriteString(` wrap="` + a.Wrap + `"`)

	case a.DisabledBoolean:
		templ.WriteString(" disabled")
	}
	for k, v := range a.Dataset {
		templ.WriteString(" data-" + k + `="` + v + `"`)
	}

}

// Begin of generated elements



func A(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "a",
		Children:   children,
	}
}

func A_(children ...Node) *HTMLNode {
    return A(Attr{}, children...)
}




func Abbr(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "abbr",
		Children:   children,
	}
}

func Abbr_(children ...Node) *HTMLNode {
    return Abbr(Attr{}, children...)
}




func Acronym(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "acronym",
		Children:   children,
	}
}

func Acronym_(children ...Node) *HTMLNode {
    return Acronym(Attr{}, children...)
}




func Address(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "address",
		Children:   children,
	}
}

func Address_(children ...Node) *HTMLNode {
    return Address(Attr{}, children...)
}




func Applet(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "applet",
		Children:   children,
	}
}

func Applet_(children ...Node) *HTMLNode {
    return Applet(Attr{}, children...)
}




func Area(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "area",
	}
}

func Area_() *VoidHTMLNode {
    return Area(Attr{})
}




func Article(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "article",
		Children:   children,
	}
}

func Article_(children ...Node) *HTMLNode {
    return Article(Attr{}, children...)
}




func Aside(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "aside",
		Children:   children,
	}
}

func Aside_(children ...Node) *HTMLNode {
    return Aside(Attr{}, children...)
}




func Audio(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "audio",
		Children:   children,
	}
}

func Audio_(children ...Node) *HTMLNode {
    return Audio(Attr{}, children...)
}




func B(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "b",
		Children:   children,
	}
}

func B_(children ...Node) *HTMLNode {
    return B(Attr{}, children...)
}




func Base(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "base",
	}
}

func Base_() *VoidHTMLNode {
    return Base(Attr{})
}




func Basefont(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "basefont",
		Children:   children,
	}
}

func Basefont_(children ...Node) *HTMLNode {
    return Basefont(Attr{}, children...)
}




func Bdi(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "bdi",
		Children:   children,
	}
}

func Bdi_(children ...Node) *HTMLNode {
    return Bdi(Attr{}, children...)
}




func Bdo(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "bdo",
		Children:   children,
	}
}

func Bdo_(children ...Node) *HTMLNode {
    return Bdo(Attr{}, children...)
}




func Bgsound(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "bgsound",
		Children:   children,
	}
}

func Bgsound_(children ...Node) *HTMLNode {
    return Bgsound(Attr{}, children...)
}




func Big(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "big",
		Children:   children,
	}
}

func Big_(children ...Node) *HTMLNode {
    return Big(Attr{}, children...)
}




func Blink(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "blink",
		Children:   children,
	}
}

func Blink_(children ...Node) *HTMLNode {
    return Blink(Attr{}, children...)
}




func Blockquote(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "blockquote",
		Children:   children,
	}
}

func Blockquote_(children ...Node) *HTMLNode {
    return Blockquote(Attr{}, children...)
}




func Body(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "body",
		Children:   children,
	}
}

func Body_(children ...Node) *HTMLNode {
    return Body(Attr{}, children...)
}




func Br(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "br",
	}
}

func Br_() *VoidHTMLNode {
    return Br(Attr{})
}




func Button(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "button",
		Children:   children,
	}
}

func Button_(children ...Node) *HTMLNode {
    return Button(Attr{}, children...)
}




func Canvas(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "canvas",
		Children:   children,
	}
}

func Canvas_(children ...Node) *HTMLNode {
    return Canvas(Attr{}, children...)
}




func Caption(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "caption",
		Children:   children,
	}
}

func Caption_(children ...Node) *HTMLNode {
    return Caption(Attr{}, children...)
}




func Center(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "center",
		Children:   children,
	}
}

func Center_(children ...Node) *HTMLNode {
    return Center(Attr{}, children...)
}




func Cite(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "cite",
		Children:   children,
	}
}

func Cite_(children ...Node) *HTMLNode {
    return Cite(Attr{}, children...)
}




func Code(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "code",
		Children:   children,
	}
}

func Code_(children ...Node) *HTMLNode {
    return Code(Attr{}, children...)
}




func Col(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "col",
	}
}

func Col_() *VoidHTMLNode {
    return Col(Attr{})
}




func Colgroup(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "colgroup",
		Children:   children,
	}
}

func Colgroup_(children ...Node) *HTMLNode {
    return Colgroup(Attr{}, children...)
}




func Datalist(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "datalist",
		Children:   children,
	}
}

func Datalist_(children ...Node) *HTMLNode {
    return Datalist(Attr{}, children...)
}




func Dd(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dd",
		Children:   children,
	}
}

func Dd_(children ...Node) *HTMLNode {
    return Dd(Attr{}, children...)
}




func Del(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "del",
		Children:   children,
	}
}

func Del_(children ...Node) *HTMLNode {
    return Del(Attr{}, children...)
}




func Details(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "details",
		Children:   children,
	}
}

func Details_(children ...Node) *HTMLNode {
    return Details(Attr{}, children...)
}




func Dfn(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dfn",
		Children:   children,
	}
}

func Dfn_(children ...Node) *HTMLNode {
    return Dfn(Attr{}, children...)
}




func Dir(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dir",
		Children:   children,
	}
}

func Dir_(children ...Node) *HTMLNode {
    return Dir(Attr{}, children...)
}




func Div(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "div",
		Children:   children,
	}
}

func Div_(children ...Node) *HTMLNode {
    return Div(Attr{}, children...)
}




func Dl(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dl",
		Children:   children,
	}
}

func Dl_(children ...Node) *HTMLNode {
    return Dl(Attr{}, children...)
}




func Dt(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dt",
		Children:   children,
	}
}

func Dt_(children ...Node) *HTMLNode {
    return Dt(Attr{}, children...)
}




func Em(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "em",
		Children:   children,
	}
}

func Em_(children ...Node) *HTMLNode {
    return Em(Attr{}, children...)
}




func Embed(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "embed",
	}
}

func Embed_() *VoidHTMLNode {
    return Embed(Attr{})
}




func Fieldset(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "fieldset",
		Children:   children,
	}
}

func Fieldset_(children ...Node) *HTMLNode {
    return Fieldset(Attr{}, children...)
}




func Figcaption(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "figcaption",
		Children:   children,
	}
}

func Figcaption_(children ...Node) *HTMLNode {
    return Figcaption(Attr{}, children...)
}




func Figure(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "figure",
		Children:   children,
	}
}

func Figure_(children ...Node) *HTMLNode {
    return Figure(Attr{}, children...)
}




func Font(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "font",
		Children:   children,
	}
}

func Font_(children ...Node) *HTMLNode {
    return Font(Attr{}, children...)
}




func Footer(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "footer",
		Children:   children,
	}
}

func Footer_(children ...Node) *HTMLNode {
    return Footer(Attr{}, children...)
}




func Form(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "form",
		Children:   children,
	}
}

func Form_(children ...Node) *HTMLNode {
    return Form(Attr{}, children...)
}




func Frame(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "frame",
		Children:   children,
	}
}

func Frame_(children ...Node) *HTMLNode {
    return Frame(Attr{}, children...)
}




func Frameset(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "frameset",
		Children:   children,
	}
}

func Frameset_(children ...Node) *HTMLNode {
    return Frameset(Attr{}, children...)
}




func H1(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h1",
		Children:   children,
	}
}

func H1_(children ...Node) *HTMLNode {
    return H1(Attr{}, children...)
}




func H2(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h2",
		Children:   children,
	}
}

func H2_(children ...Node) *HTMLNode {
    return H2(Attr{}, children...)
}




func H3(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h3",
		Children:   children,
	}
}

func H3_(children ...Node) *HTMLNode {
    return H3(Attr{}, children...)
}




func H4(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h4",
		Children:   children,
	}
}

func H4_(children ...Node) *HTMLNode {
    return H4(Attr{}, children...)
}




func H5(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h5",
		Children:   children,
	}
}

func H5_(children ...Node) *HTMLNode {
    return H5(Attr{}, children...)
}




func H6(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h6",
		Children:   children,
	}
}

func H6_(children ...Node) *HTMLNode {
    return H6(Attr{}, children...)
}




func Head(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "head",
		Children:   children,
	}
}

func Head_(children ...Node) *HTMLNode {
    return Head(Attr{}, children...)
}




func Header(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "header",
		Children:   children,
	}
}

func Header_(children ...Node) *HTMLNode {
    return Header(Attr{}, children...)
}




func Hgroup(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "hgroup",
		Children:   children,
	}
}

func Hgroup_(children ...Node) *HTMLNode {
    return Hgroup(Attr{}, children...)
}




func Hr(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "hr",
	}
}

func Hr_() *VoidHTMLNode {
    return Hr(Attr{})
}




func Html(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "html",
		Children:   children,
	}
}

func Html_(children ...Node) *HTMLNode {
    return Html(Attr{}, children...)
}




func I(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "i",
		Children:   children,
	}
}

func I_(children ...Node) *HTMLNode {
    return I(Attr{}, children...)
}




func Iframe(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "iframe",
		Children:   children,
	}
}

func Iframe_(children ...Node) *HTMLNode {
    return Iframe(Attr{}, children...)
}




func Img(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "img",
	}
}

func Img_() *VoidHTMLNode {
    return Img(Attr{})
}




func Input(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "input",
	}
}

func Input_() *VoidHTMLNode {
    return Input(Attr{})
}




func Ins(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ins",
		Children:   children,
	}
}

func Ins_(children ...Node) *HTMLNode {
    return Ins(Attr{}, children...)
}




func Isindex(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "isindex",
		Children:   children,
	}
}

func Isindex_(children ...Node) *HTMLNode {
    return Isindex(Attr{}, children...)
}




func Kbd(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "kbd",
		Children:   children,
	}
}

func Kbd_(children ...Node) *HTMLNode {
    return Kbd(Attr{}, children...)
}




func Keygen(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "keygen",
		Children:   children,
	}
}

func Keygen_(children ...Node) *HTMLNode {
    return Keygen(Attr{}, children...)
}




func Label(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "label",
		Children:   children,
	}
}

func Label_(children ...Node) *HTMLNode {
    return Label(Attr{}, children...)
}




func Legend(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "legend",
		Children:   children,
	}
}

func Legend_(children ...Node) *HTMLNode {
    return Legend(Attr{}, children...)
}




func Li(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "li",
		Children:   children,
	}
}

func Li_(children ...Node) *HTMLNode {
    return Li(Attr{}, children...)
}




func Link(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "link",
	}
}

func Link_() *VoidHTMLNode {
    return Link(Attr{})
}




func Listing(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "listing",
		Children:   children,
	}
}

func Listing_(children ...Node) *HTMLNode {
    return Listing(Attr{}, children...)
}




func Main(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "main",
		Children:   children,
	}
}

func Main_(children ...Node) *HTMLNode {
    return Main(Attr{}, children...)
}




func Map(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "map",
		Children:   children,
	}
}

func Map_(children ...Node) *HTMLNode {
    return Map(Attr{}, children...)
}




func Mark(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "mark",
		Children:   children,
	}
}

func Mark_(children ...Node) *HTMLNode {
    return Mark(Attr{}, children...)
}




func Marquee(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "marquee",
		Children:   children,
	}
}

func Marquee_(children ...Node) *HTMLNode {
    return Marquee(Attr{}, children...)
}




func Menu(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "menu",
		Children:   children,
	}
}

func Menu_(children ...Node) *HTMLNode {
    return Menu(Attr{}, children...)
}




func Meta(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "meta",
	}
}

func Meta_() *VoidHTMLNode {
    return Meta(Attr{})
}




func Meter(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "meter",
		Children:   children,
	}
}

func Meter_(children ...Node) *HTMLNode {
    return Meter(Attr{}, children...)
}




func Nav(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "nav",
		Children:   children,
	}
}

func Nav_(children ...Node) *HTMLNode {
    return Nav(Attr{}, children...)
}




func Nobr(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "nobr",
		Children:   children,
	}
}

func Nobr_(children ...Node) *HTMLNode {
    return Nobr(Attr{}, children...)
}




func Noframes(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "noframes",
		Children:   children,
	}
}

func Noframes_(children ...Node) *HTMLNode {
    return Noframes(Attr{}, children...)
}




func Noscript(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "noscript",
		Children:   children,
	}
}

func Noscript_(children ...Node) *HTMLNode {
    return Noscript(Attr{}, children...)
}




func Object(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "object",
		Children:   children,
	}
}

func Object_(children ...Node) *HTMLNode {
    return Object(Attr{}, children...)
}




func Ol(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ol",
		Children:   children,
	}
}

func Ol_(children ...Node) *HTMLNode {
    return Ol(Attr{}, children...)
}




func Optgroup(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "optgroup",
		Children:   children,
	}
}

func Optgroup_(children ...Node) *HTMLNode {
    return Optgroup(Attr{}, children...)
}




func Option(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "option",
		Children:   children,
	}
}

func Option_(children ...Node) *HTMLNode {
    return Option(Attr{}, children...)
}




func Output(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "output",
		Children:   children,
	}
}

func Output_(children ...Node) *HTMLNode {
    return Output(Attr{}, children...)
}




func P(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "p",
		Children:   children,
	}
}

func P_(children ...Node) *HTMLNode {
    return P(Attr{}, children...)
}




func Param(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "param",
	}
}

func Param_() *VoidHTMLNode {
    return Param(Attr{})
}




func Plaintext(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "plaintext",
		Children:   children,
	}
}

func Plaintext_(children ...Node) *HTMLNode {
    return Plaintext(Attr{}, children...)
}




func Pre(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "pre",
		Children:   children,
	}
}

func Pre_(children ...Node) *HTMLNode {
    return Pre(Attr{}, children...)
}




func Progress(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "progress",
		Children:   children,
	}
}

func Progress_(children ...Node) *HTMLNode {
    return Progress(Attr{}, children...)
}




func Q(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "q",
		Children:   children,
	}
}

func Q_(children ...Node) *HTMLNode {
    return Q(Attr{}, children...)
}




func Rp(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "rp",
		Children:   children,
	}
}

func Rp_(children ...Node) *HTMLNode {
    return Rp(Attr{}, children...)
}




func Rt(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "rt",
		Children:   children,
	}
}

func Rt_(children ...Node) *HTMLNode {
    return Rt(Attr{}, children...)
}




func Ruby(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ruby",
		Children:   children,
	}
}

func Ruby_(children ...Node) *HTMLNode {
    return Ruby(Attr{}, children...)
}




func S(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "s",
		Children:   children,
	}
}

func S_(children ...Node) *HTMLNode {
    return S(Attr{}, children...)
}




func Samp(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "samp",
		Children:   children,
	}
}

func Samp_(children ...Node) *HTMLNode {
    return Samp(Attr{}, children...)
}




func Section(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "section",
		Children:   children,
	}
}

func Section_(children ...Node) *HTMLNode {
    return Section(Attr{}, children...)
}




func Select(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "select",
		Children:   children,
	}
}

func Select_(children ...Node) *HTMLNode {
    return Select(Attr{}, children...)
}




func Small(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "small",
		Children:   children,
	}
}

func Small_(children ...Node) *HTMLNode {
    return Small(Attr{}, children...)
}




func Source(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "source",
	}
}

func Source_() *VoidHTMLNode {
    return Source(Attr{})
}




func Spacer(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "spacer",
		Children:   children,
	}
}

func Spacer_(children ...Node) *HTMLNode {
    return Spacer(Attr{}, children...)
}




func Span(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "span",
		Children:   children,
	}
}

func Span_(children ...Node) *HTMLNode {
    return Span(Attr{}, children...)
}




func Strike(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "strike",
		Children:   children,
	}
}

func Strike_(children ...Node) *HTMLNode {
    return Strike(Attr{}, children...)
}




func Strong(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "strong",
		Children:   children,
	}
}

func Strong_(children ...Node) *HTMLNode {
    return Strong(Attr{}, children...)
}




func Style(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "style",
		Children:   children,
	}
}

func Style_(children ...Node) *HTMLNode {
    return Style(Attr{}, children...)
}




func Sub(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "sub",
		Children:   children,
	}
}

func Sub_(children ...Node) *HTMLNode {
    return Sub(Attr{}, children...)
}




func Summary(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "summary",
		Children:   children,
	}
}

func Summary_(children ...Node) *HTMLNode {
    return Summary(Attr{}, children...)
}




func Sup(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "sup",
		Children:   children,
	}
}

func Sup_(children ...Node) *HTMLNode {
    return Sup(Attr{}, children...)
}




func Table(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "table",
		Children:   children,
	}
}

func Table_(children ...Node) *HTMLNode {
    return Table(Attr{}, children...)
}




func Tbody(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tbody",
		Children:   children,
	}
}

func Tbody_(children ...Node) *HTMLNode {
    return Tbody(Attr{}, children...)
}




func Td(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "td",
		Children:   children,
	}
}

func Td_(children ...Node) *HTMLNode {
    return Td(Attr{}, children...)
}




func Textarea(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "textarea",
		Children:   children,
	}
}

func Textarea_(children ...Node) *HTMLNode {
    return Textarea(Attr{}, children...)
}




func Tfoot(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tfoot",
		Children:   children,
	}
}

func Tfoot_(children ...Node) *HTMLNode {
    return Tfoot(Attr{}, children...)
}




func Th(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "th",
		Children:   children,
	}
}

func Th_(children ...Node) *HTMLNode {
    return Th(Attr{}, children...)
}




func Thead(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "thead",
		Children:   children,
	}
}

func Thead_(children ...Node) *HTMLNode {
    return Thead(Attr{}, children...)
}




func Time(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "time",
		Children:   children,
	}
}

func Time_(children ...Node) *HTMLNode {
    return Time(Attr{}, children...)
}




func Title(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "title",
		Children:   children,
	}
}

func Title_(children ...Node) *HTMLNode {
    return Title(Attr{}, children...)
}




func Tr(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tr",
		Children:   children,
	}
}

func Tr_(children ...Node) *HTMLNode {
    return Tr(Attr{}, children...)
}




func Track(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "track",
	}
}

func Track_() *VoidHTMLNode {
    return Track(Attr{})
}




func Tt(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tt",
		Children:   children,
	}
}

func Tt_(children ...Node) *HTMLNode {
    return Tt(Attr{}, children...)
}




func U(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "u",
		Children:   children,
	}
}

func U_(children ...Node) *HTMLNode {
    return U(Attr{}, children...)
}




func Ul(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ul",
		Children:   children,
	}
}

func Ul_(children ...Node) *HTMLNode {
    return Ul(Attr{}, children...)
}




func Var(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "var",
		Children:   children,
	}
}

func Var_(children ...Node) *HTMLNode {
    return Var(Attr{}, children...)
}




func Video(attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "video",
		Children:   children,
	}
}

func Video_(children ...Node) *HTMLNode {
    return Video(Attr{}, children...)
}




func Wbr(attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "wbr",
	}
}

func Wbr_() *VoidHTMLNode {
    return Wbr(Attr{})
}


