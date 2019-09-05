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
//func Srcset(srcset string) template.Srcset {
//	return template.Srcset(srcset)
//}


type HTMLNode struct {
	Attributes Attributes
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
	Attributes Attributes
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

type Attributes struct {
	Slice []Attribute
}

type Attribute struct {
	Name string
	Data interface{}
	Templ string
}

// TODO Dataset
// TODO boolean attributes


func (a Attributes) Accept(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "accept",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) AcceptCharset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "accept-charset",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Accesskey(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "accesskey",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Action(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "action",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Align(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "align",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Alt(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "alt",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) AriaExpanded(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "aria-expanded",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) AriaHidden(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "aria-hidden",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) AriaLabel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "aria-label",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Async(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "async",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Autocomplete(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "autocomplete",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Autofocus(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "autofocus",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Autoplay(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "autoplay",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Bgcolor(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "bgcolor",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Border(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "border",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Charset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "charset",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Checked(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "checked",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Cite(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "cite",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Class(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "class",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Color(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "color",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Cols(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "cols",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Colspan(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "colspan",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Content(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "content",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Contenteditable(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "contenteditable",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Controls(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "controls",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Coords(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "coords",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Data(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "data",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Datetime(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "datetime",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Default(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "default",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Defer(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "defer",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Dir(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "dir",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Dirname(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "dirname",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Disabled(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "disabled",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Download(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "download",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Draggable(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "draggable",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Dropzone(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "dropzone",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Enctype(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "enctype",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) For(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "for",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Form(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "form",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Formaction(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "formaction",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Headers(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "headers",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Height(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "height",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Hidden(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "hidden",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) High(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "high",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Href(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "href",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Hreflang(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "hreflang",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) HttpEquiv(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "http-equiv",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Id(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "id",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) InitialScale(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "initial-scale",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ismap(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ismap",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Kind(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "kind",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Label(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "label",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Lang(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "lang",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) List(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "list",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Loop(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "loop",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Low(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "low",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Max(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "max",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Maxlength(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "maxlength",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Media(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "media",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Method(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "method",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Min(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "min",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Multiple(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "multiple",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Muted(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "muted",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Name(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "name",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Novalidate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "novalidate",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onabort(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onabort",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onafterprint(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onafterprint",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onbeforeprint(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onbeforeprint",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onbeforeunload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onbeforeunload",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onblur(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onblur",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncanplay(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oncanplay",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncanplaythrough(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oncanplaythrough",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onchange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onchange",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onclick(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onclick",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncontextmenu(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oncontextmenu",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncopy(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oncopy",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncuechange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oncuechange",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncut(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oncut",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondblclick(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondblclick",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondrag(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondrag",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragend(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragend",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragenter(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragenter",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragleave(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragleave",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragover(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragover",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragstart(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragstart",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondrop(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondrop",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondurationchange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ondurationchange",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onemptied(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onemptied",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onended(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onended",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onerror(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onerror",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onfocus(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onfocus",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onhashchange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onhashchange",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oninput(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oninput",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oninvalid(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "oninvalid",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onkeydown(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onkeydown",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onkeypress(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onkeypress",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onkeyup(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onkeyup",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onload",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onloadeddata(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onloadeddata",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onloadedmetadata(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onloadedmetadata",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onloadstart(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onloadstart",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmousedown(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onmousedown",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmousemove(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onmousemove",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmouseout(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onmouseout",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmouseover(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onmouseover",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmouseup(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onmouseup",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmousewheel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onmousewheel",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onoffline(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onoffline",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ononline(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ononline",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpagehide(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onpagehide",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpageshow(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onpageshow",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpaste(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onpaste",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpause(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onpause",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onplay(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onplay",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onplaying(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onplaying",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpopstate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onpopstate",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onprogress(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onprogress",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onratechange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onratechange",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onreset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onreset",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onresize(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onresize",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onscroll(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onscroll",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onsearch(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onsearch",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onseeked(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onseeked",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onseeking(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onseeking",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onselect(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onselect",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onstalled(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onstalled",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onstorage(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onstorage",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onsubmit(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onsubmit",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onsuspend(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onsuspend",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ontimeupdate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ontimeupdate",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ontoggle(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "ontoggle",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onunload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onunload",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onvolumechange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onvolumechange",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onwaiting(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onwaiting",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onwheel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "onwheel",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Open(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "open",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Optimum(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "optimum",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Pattern(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "pattern",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Placeholder(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "placeholder",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Poster(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "poster",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Preload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "preload",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Readonly(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "readonly",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Rel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "rel",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Required(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "required",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Reversed(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "reversed",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Role(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "role",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Rows(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "rows",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Rowspan(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "rowspan",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Sandbox(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "sandbox",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Scope(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "scope",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Selected(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "selected",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Shape(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "shape",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Size(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "size",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Sizes(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "sizes",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Span(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "span",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Spellcheck(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "spellcheck",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Src(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "src",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Srcdoc(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "srcdoc",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Srclang(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "srclang",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Srcset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "srcset",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Start(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "start",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Step(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "step",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Style(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "style",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Tabindex(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "tabindex",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Target(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "target",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Title(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "title",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Translate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "translate",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Type(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "type",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Usemap(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "usemap",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Value(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "value",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Width(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "width",
		Data: value,
		Templ: templ,
	})
	return a
}

func (a Attributes) Wrap(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "wrap",
		Data: value,
		Templ: templ,
	})
	return a
}



func (a Attributes) Accept_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "accept",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) AcceptCharset_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "accept-charset",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Accesskey_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "accesskey",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Action_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "action",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Align_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "align",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Alt_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "alt",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) AriaExpanded_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "aria-expanded",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) AriaHidden_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "aria-hidden",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) AriaLabel_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "aria-label",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Async_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "async",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Autocomplete_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "autocomplete",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Autofocus_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "autofocus",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Autoplay_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "autoplay",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Bgcolor_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "bgcolor",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Border_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "border",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Charset_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "charset",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Checked_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "checked",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Cite_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "cite",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Class_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "class",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Color_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "color",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Cols_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "cols",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Colspan_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "colspan",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Content_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "content",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Contenteditable_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "contenteditable",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Controls_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "controls",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Coords_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "coords",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Data_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "data",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Datetime_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "datetime",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Default_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "default",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Defer_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "defer",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Dir_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "dir",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Dirname_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "dirname",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Disabled_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "disabled",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Download_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "download",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Draggable_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "draggable",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Dropzone_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "dropzone",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Enctype_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "enctype",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) For_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "for",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Form_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "form",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Formaction_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "formaction",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Headers_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "headers",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Height_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "height",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Hidden_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "hidden",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) High_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "high",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Href_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "href",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Hreflang_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "hreflang",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) HttpEquiv_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "http-equiv",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Id_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "id",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) InitialScale_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "initial-scale",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ismap_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ismap",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Kind_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "kind",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Label_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "label",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Lang_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "lang",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) List_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "list",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Loop_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "loop",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Low_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "low",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Max_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "max",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Maxlength_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "maxlength",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Media_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "media",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Method_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "method",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Min_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "min",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Multiple_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "multiple",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Muted_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "muted",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Name_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "name",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Novalidate_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "novalidate",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onabort_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onabort",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onafterprint_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onafterprint",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onbeforeprint_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onbeforeprint",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onbeforeunload_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onbeforeunload",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onblur_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onblur",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncanplay_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oncanplay",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncanplaythrough_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oncanplaythrough",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onchange_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onchange",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onclick_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onclick",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncontextmenu_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oncontextmenu",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncopy_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oncopy",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncuechange_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oncuechange",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oncut_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oncut",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondblclick_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondblclick",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondrag_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondrag",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragend_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragend",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragenter_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragenter",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragleave_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragleave",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragover_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragover",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondragstart_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondragstart",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondrop_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondrop",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ondurationchange_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ondurationchange",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onemptied_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onemptied",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onended_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onended",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onerror_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onerror",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onfocus_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onfocus",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onhashchange_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onhashchange",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oninput_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oninput",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Oninvalid_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "oninvalid",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onkeydown_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onkeydown",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onkeypress_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onkeypress",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onkeyup_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onkeyup",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onload_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onload",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onloadeddata_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onloadeddata",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onloadedmetadata_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onloadedmetadata",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onloadstart_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onloadstart",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmousedown_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onmousedown",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmousemove_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onmousemove",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmouseout_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onmouseout",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmouseover_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onmouseover",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmouseup_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onmouseup",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onmousewheel_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onmousewheel",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onoffline_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onoffline",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ononline_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ononline",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpagehide_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onpagehide",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpageshow_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onpageshow",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpaste_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onpaste",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpause_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onpause",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onplay_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onplay",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onplaying_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onplaying",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onpopstate_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onpopstate",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onprogress_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onprogress",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onratechange_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onratechange",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onreset_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onreset",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onresize_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onresize",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onscroll_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onscroll",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onsearch_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onsearch",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onseeked_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onseeked",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onseeking_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onseeking",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onselect_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onselect",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onstalled_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onstalled",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onstorage_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onstorage",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onsubmit_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onsubmit",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onsuspend_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onsuspend",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ontimeupdate_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ontimeupdate",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Ontoggle_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "ontoggle",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onunload_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onunload",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onvolumechange_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onvolumechange",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onwaiting_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onwaiting",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Onwheel_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "onwheel",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Open_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "open",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Optimum_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "optimum",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Pattern_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "pattern",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Placeholder_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "placeholder",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Poster_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "poster",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Preload_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "preload",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Readonly_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "readonly",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Rel_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "rel",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Required_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "required",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Reversed_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "reversed",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Role_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "role",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Rows_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "rows",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Rowspan_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "rowspan",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Sandbox_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "sandbox",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Scope_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "scope",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Selected_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "selected",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Shape_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "shape",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Size_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "size",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Sizes_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "sizes",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Span_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "span",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Spellcheck_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "spellcheck",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Src_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "src",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Srcdoc_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "srcdoc",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Srclang_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "srclang",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Srcset_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "srcset",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Start_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "start",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Step_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "step",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Style_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "style",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Tabindex_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "tabindex",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Target_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "target",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Title_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "title",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Translate_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "translate",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Type_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "type",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Usemap_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "usemap",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Value_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "value",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Width_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "width",
		Data: nil,
		Templ: templ,
	})
	return a
}

func (a Attributes) Wrap_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "wrap",
		Data: nil,
		Templ: templ,
	})
	return a
}




func Accept(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "accept",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func AcceptCharset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "accept-charset",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Accesskey(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "accesskey",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Action(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "action",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Align(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "align",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Alt(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "alt",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func AriaExpanded(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "aria-expanded",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func AriaHidden(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "aria-hidden",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func AriaLabel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "aria-label",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Async(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "async",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Autocomplete(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "autocomplete",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Autofocus(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "autofocus",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Autoplay(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "autoplay",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Bgcolor(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "bgcolor",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Border(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "border",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Charset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "charset",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Checked(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "checked",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Class(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "class",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Color(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "color",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Cols(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "cols",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Colspan(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "colspan",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Content(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "content",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Contenteditable(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "contenteditable",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Controls(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "controls",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Coords(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "coords",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Data(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "data",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Datetime(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "datetime",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Default(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "default",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Defer(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "defer",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Dirname(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "dirname",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Disabled(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "disabled",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Download(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "download",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Draggable(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "draggable",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Dropzone(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "dropzone",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Enctype(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "enctype",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func For(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "for",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Formaction(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "formaction",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Headers(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "headers",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Height(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "height",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Hidden(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "hidden",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func High(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "high",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Href(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "href",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Hreflang(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "hreflang",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func HttpEquiv(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "http-equiv",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Id(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "id",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func InitialScale(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "initial-scale",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ismap(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ismap",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Kind(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "kind",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Lang(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "lang",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func List(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "list",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Loop(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "loop",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Low(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "low",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Max(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "max",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Maxlength(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "maxlength",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Media(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "media",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Method(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "method",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Min(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "min",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Multiple(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "multiple",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Muted(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "muted",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Name(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "name",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Novalidate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "novalidate",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onabort(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onabort",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onafterprint(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onafterprint",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onbeforeprint(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onbeforeprint",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onbeforeunload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onbeforeunload",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onblur(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onblur",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncanplay(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncanplay",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncanplaythrough(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncanplaythrough",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onchange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onchange",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onclick(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onclick",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncontextmenu(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncontextmenu",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncopy(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncopy",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncuechange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncuechange",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncut(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncut",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondblclick(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondblclick",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondrag(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondrag",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragend(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragend",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragenter(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragenter",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragleave(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragleave",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragover(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragover",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragstart(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragstart",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondrop(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondrop",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondurationchange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondurationchange",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onemptied(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onemptied",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onended(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onended",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onerror(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onerror",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onfocus(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onfocus",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onhashchange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onhashchange",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oninput(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oninput",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Oninvalid(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oninvalid",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onkeydown(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onkeydown",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onkeypress(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onkeypress",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onkeyup(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onkeyup",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onload",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onloadeddata(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onloadeddata",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onloadedmetadata(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onloadedmetadata",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onloadstart(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onloadstart",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmousedown(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmousedown",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmousemove(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmousemove",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmouseout(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmouseout",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmouseover(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmouseover",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmouseup(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmouseup",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmousewheel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmousewheel",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onoffline(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onoffline",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ononline(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ononline",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpagehide(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpagehide",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpageshow(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpageshow",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpaste(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpaste",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpause(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpause",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onplay(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onplay",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onplaying(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onplaying",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpopstate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpopstate",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onprogress(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onprogress",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onratechange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onratechange",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onreset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onreset",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onresize(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onresize",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onscroll(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onscroll",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onsearch(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onsearch",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onseeked(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onseeked",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onseeking(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onseeking",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onselect(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onselect",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onstalled(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onstalled",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onstorage(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onstorage",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onsubmit(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onsubmit",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onsuspend(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onsuspend",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ontimeupdate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ontimeupdate",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Ontoggle(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ontoggle",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onunload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onunload",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onvolumechange(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onvolumechange",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onwaiting(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onwaiting",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Onwheel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onwheel",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Open(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "open",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Optimum(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "optimum",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Pattern(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "pattern",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Placeholder(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "placeholder",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Poster(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "poster",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Preload(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "preload",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Readonly(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "readonly",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Rel(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "rel",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Required(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "required",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Reversed(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "reversed",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Role(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "role",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Rows(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "rows",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Rowspan(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "rowspan",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Sandbox(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "sandbox",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Scope(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "scope",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Selected(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "selected",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Shape(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "shape",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Size(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "size",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Sizes(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "sizes",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Spellcheck(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "spellcheck",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Src(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "src",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Srcdoc(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "srcdoc",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Srclang(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "srclang",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Srcset(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "srcset",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Start(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "start",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Step(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "step",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Tabindex(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "tabindex",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Target(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "target",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Translate(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "translate",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Type(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "type",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Usemap(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "usemap",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Value(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "value",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Width(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "width",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}



func Wrap(value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "wrap",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}





func Accept_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "accept",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func AcceptCharset_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "accept-charset",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Accesskey_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "accesskey",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Action_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "action",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Align_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "align",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Alt_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "alt",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func AriaExpanded_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "aria-expanded",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func AriaHidden_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "aria-hidden",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func AriaLabel_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "aria-label",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Async_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "async",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Autocomplete_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "autocomplete",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Autofocus_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "autofocus",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Autoplay_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "autoplay",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Bgcolor_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "bgcolor",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Border_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "border",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Charset_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "charset",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Checked_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "checked",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}





func Class_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "class",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Color_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "color",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Cols_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "cols",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Colspan_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "colspan",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Content_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "content",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Contenteditable_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "contenteditable",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Controls_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "controls",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Coords_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "coords",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Data_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "data",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Datetime_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "datetime",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Default_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "default",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Defer_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "defer",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}





func Dirname_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "dirname",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Disabled_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "disabled",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Download_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "download",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Draggable_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "draggable",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Dropzone_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "dropzone",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Enctype_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "enctype",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func For_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "for",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}





func Formaction_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "formaction",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Headers_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "headers",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Height_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "height",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Hidden_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "hidden",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func High_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "high",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Href_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "href",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Hreflang_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "hreflang",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func HttpEquiv_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "http-equiv",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Id_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "id",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func InitialScale_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "initial-scale",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ismap_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ismap",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Kind_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "kind",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}





func Lang_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "lang",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func List_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "list",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Loop_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "loop",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Low_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "low",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Max_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "max",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Maxlength_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "maxlength",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Media_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "media",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Method_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "method",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Min_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "min",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Multiple_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "multiple",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Muted_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "muted",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Name_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "name",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Novalidate_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "novalidate",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onabort_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onabort",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onafterprint_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onafterprint",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onbeforeprint_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onbeforeprint",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onbeforeunload_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onbeforeunload",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onblur_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onblur",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncanplay_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncanplay",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncanplaythrough_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncanplaythrough",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onchange_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onchange",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onclick_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onclick",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncontextmenu_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncontextmenu",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncopy_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncopy",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncuechange_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncuechange",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oncut_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oncut",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondblclick_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondblclick",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondrag_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondrag",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragend_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragend",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragenter_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragenter",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragleave_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragleave",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragover_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragover",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondragstart_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondragstart",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondrop_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondrop",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ondurationchange_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ondurationchange",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onemptied_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onemptied",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onended_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onended",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onerror_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onerror",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onfocus_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onfocus",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onhashchange_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onhashchange",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oninput_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oninput",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Oninvalid_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "oninvalid",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onkeydown_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onkeydown",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onkeypress_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onkeypress",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onkeyup_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onkeyup",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onload_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onload",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onloadeddata_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onloadeddata",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onloadedmetadata_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onloadedmetadata",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onloadstart_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onloadstart",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmousedown_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmousedown",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmousemove_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmousemove",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmouseout_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmouseout",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmouseover_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmouseover",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmouseup_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmouseup",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onmousewheel_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onmousewheel",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onoffline_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onoffline",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ononline_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ononline",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpagehide_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpagehide",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpageshow_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpageshow",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpaste_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpaste",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpause_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpause",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onplay_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onplay",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onplaying_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onplaying",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onpopstate_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onpopstate",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onprogress_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onprogress",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onratechange_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onratechange",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onreset_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onreset",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onresize_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onresize",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onscroll_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onscroll",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onsearch_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onsearch",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onseeked_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onseeked",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onseeking_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onseeking",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onselect_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onselect",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onstalled_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onstalled",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onstorage_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onstorage",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onsubmit_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onsubmit",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onsuspend_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onsuspend",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ontimeupdate_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ontimeupdate",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Ontoggle_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "ontoggle",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onunload_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onunload",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onvolumechange_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onvolumechange",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onwaiting_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onwaiting",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Onwheel_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "onwheel",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Open_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "open",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Optimum_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "optimum",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Pattern_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "pattern",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Placeholder_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "placeholder",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Poster_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "poster",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Preload_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "preload",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Readonly_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "readonly",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Rel_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "rel",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Required_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "required",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Reversed_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "reversed",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Role_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "role",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Rows_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "rows",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Rowspan_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "rowspan",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Sandbox_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "sandbox",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Scope_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "scope",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Selected_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "selected",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Shape_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "shape",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Size_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "size",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Sizes_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "sizes",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}





func Spellcheck_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "spellcheck",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Src_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "src",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Srcdoc_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "srcdoc",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Srclang_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "srclang",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Srcset_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "srcset",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Start_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "start",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Step_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "step",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}





func Tabindex_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "tabindex",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Target_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "target",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}





func Translate_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "translate",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Type_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "type",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Usemap_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "usemap",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Value_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "value",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Width_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "width",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Wrap_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "wrap",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}



func Attr() Attributes {
	return Attributes{
		Slice: []Attribute{},
	}
}

type OldAttr struct {
	Dataset  Dataset
	DisabledBoolean  bool
	templData map[string]interface{}
}

func (a *Attributes) buildTemplateTo(templ *strings.Builder, vals Values) {
	for _, attr := range a.Slice {
		placeholder := "P" + strconv.Itoa(len(vals))
		templ.WriteString(" " + attr.Name + `="` + strings.Replace(attr.Templ, "{{.", "{{."+placeholder, -1) + `"`)
		vals[placeholder] = attr.Data
	}
}

// Begin of generated elements



func A(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "a",
		Children:   children,
	}
}

func A_(children ...Node) *HTMLNode {
    return A(Attr(), children...)
}




func Abbr(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "abbr",
		Children:   children,
	}
}

func Abbr_(children ...Node) *HTMLNode {
    return Abbr(Attr(), children...)
}




func Acronym(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "acronym",
		Children:   children,
	}
}

func Acronym_(children ...Node) *HTMLNode {
    return Acronym(Attr(), children...)
}




func Address(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "address",
		Children:   children,
	}
}

func Address_(children ...Node) *HTMLNode {
    return Address(Attr(), children...)
}




func Applet(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "applet",
		Children:   children,
	}
}

func Applet_(children ...Node) *HTMLNode {
    return Applet(Attr(), children...)
}




func Area(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "area",
	}
}

func Area_() *VoidHTMLNode {
    return Area(Attr())
}




func Article(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "article",
		Children:   children,
	}
}

func Article_(children ...Node) *HTMLNode {
    return Article(Attr(), children...)
}




func Aside(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "aside",
		Children:   children,
	}
}

func Aside_(children ...Node) *HTMLNode {
    return Aside(Attr(), children...)
}




func Audio(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "audio",
		Children:   children,
	}
}

func Audio_(children ...Node) *HTMLNode {
    return Audio(Attr(), children...)
}




func B(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "b",
		Children:   children,
	}
}

func B_(children ...Node) *HTMLNode {
    return B(Attr(), children...)
}




func Base(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "base",
	}
}

func Base_() *VoidHTMLNode {
    return Base(Attr())
}




func Basefont(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "basefont",
		Children:   children,
	}
}

func Basefont_(children ...Node) *HTMLNode {
    return Basefont(Attr(), children...)
}




func Bdi(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "bdi",
		Children:   children,
	}
}

func Bdi_(children ...Node) *HTMLNode {
    return Bdi(Attr(), children...)
}




func Bdo(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "bdo",
		Children:   children,
	}
}

func Bdo_(children ...Node) *HTMLNode {
    return Bdo(Attr(), children...)
}




func Bgsound(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "bgsound",
		Children:   children,
	}
}

func Bgsound_(children ...Node) *HTMLNode {
    return Bgsound(Attr(), children...)
}




func Big(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "big",
		Children:   children,
	}
}

func Big_(children ...Node) *HTMLNode {
    return Big(Attr(), children...)
}




func Blink(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "blink",
		Children:   children,
	}
}

func Blink_(children ...Node) *HTMLNode {
    return Blink(Attr(), children...)
}




func Blockquote(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "blockquote",
		Children:   children,
	}
}

func Blockquote_(children ...Node) *HTMLNode {
    return Blockquote(Attr(), children...)
}




func Body(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "body",
		Children:   children,
	}
}

func Body_(children ...Node) *HTMLNode {
    return Body(Attr(), children...)
}




func Br(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "br",
	}
}

func Br_() *VoidHTMLNode {
    return Br(Attr())
}




func Button(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "button",
		Children:   children,
	}
}

func Button_(children ...Node) *HTMLNode {
    return Button(Attr(), children...)
}




func Canvas(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "canvas",
		Children:   children,
	}
}

func Canvas_(children ...Node) *HTMLNode {
    return Canvas(Attr(), children...)
}




func Caption(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "caption",
		Children:   children,
	}
}

func Caption_(children ...Node) *HTMLNode {
    return Caption(Attr(), children...)
}




func Center(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "center",
		Children:   children,
	}
}

func Center_(children ...Node) *HTMLNode {
    return Center(Attr(), children...)
}




func Cite(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "cite",
		Children:   children,
	}
}

func Cite_(children ...Node) *HTMLNode {
    return Cite(Attr(), children...)
}




func Code(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "code",
		Children:   children,
	}
}

func Code_(children ...Node) *HTMLNode {
    return Code(Attr(), children...)
}




func Col(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "col",
	}
}

func Col_() *VoidHTMLNode {
    return Col(Attr())
}




func Colgroup(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "colgroup",
		Children:   children,
	}
}

func Colgroup_(children ...Node) *HTMLNode {
    return Colgroup(Attr(), children...)
}




func Datalist(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "datalist",
		Children:   children,
	}
}

func Datalist_(children ...Node) *HTMLNode {
    return Datalist(Attr(), children...)
}




func Dd(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dd",
		Children:   children,
	}
}

func Dd_(children ...Node) *HTMLNode {
    return Dd(Attr(), children...)
}




func Del(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "del",
		Children:   children,
	}
}

func Del_(children ...Node) *HTMLNode {
    return Del(Attr(), children...)
}




func Details(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "details",
		Children:   children,
	}
}

func Details_(children ...Node) *HTMLNode {
    return Details(Attr(), children...)
}




func Dfn(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dfn",
		Children:   children,
	}
}

func Dfn_(children ...Node) *HTMLNode {
    return Dfn(Attr(), children...)
}




func Dir(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dir",
		Children:   children,
	}
}

func Dir_(children ...Node) *HTMLNode {
    return Dir(Attr(), children...)
}




func Div(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "div",
		Children:   children,
	}
}

func Div_(children ...Node) *HTMLNode {
    return Div(Attr(), children...)
}




func Dl(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dl",
		Children:   children,
	}
}

func Dl_(children ...Node) *HTMLNode {
    return Dl(Attr(), children...)
}




func Dt(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "dt",
		Children:   children,
	}
}

func Dt_(children ...Node) *HTMLNode {
    return Dt(Attr(), children...)
}




func Em(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "em",
		Children:   children,
	}
}

func Em_(children ...Node) *HTMLNode {
    return Em(Attr(), children...)
}




func Embed(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "embed",
	}
}

func Embed_() *VoidHTMLNode {
    return Embed(Attr())
}




func Fieldset(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "fieldset",
		Children:   children,
	}
}

func Fieldset_(children ...Node) *HTMLNode {
    return Fieldset(Attr(), children...)
}




func Figcaption(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "figcaption",
		Children:   children,
	}
}

func Figcaption_(children ...Node) *HTMLNode {
    return Figcaption(Attr(), children...)
}




func Figure(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "figure",
		Children:   children,
	}
}

func Figure_(children ...Node) *HTMLNode {
    return Figure(Attr(), children...)
}




func Font(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "font",
		Children:   children,
	}
}

func Font_(children ...Node) *HTMLNode {
    return Font(Attr(), children...)
}




func Footer(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "footer",
		Children:   children,
	}
}

func Footer_(children ...Node) *HTMLNode {
    return Footer(Attr(), children...)
}




func Form(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "form",
		Children:   children,
	}
}

func Form_(children ...Node) *HTMLNode {
    return Form(Attr(), children...)
}




func Frame(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "frame",
		Children:   children,
	}
}

func Frame_(children ...Node) *HTMLNode {
    return Frame(Attr(), children...)
}




func Frameset(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "frameset",
		Children:   children,
	}
}

func Frameset_(children ...Node) *HTMLNode {
    return Frameset(Attr(), children...)
}




func H1(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h1",
		Children:   children,
	}
}

func H1_(children ...Node) *HTMLNode {
    return H1(Attr(), children...)
}




func H2(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h2",
		Children:   children,
	}
}

func H2_(children ...Node) *HTMLNode {
    return H2(Attr(), children...)
}




func H3(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h3",
		Children:   children,
	}
}

func H3_(children ...Node) *HTMLNode {
    return H3(Attr(), children...)
}




func H4(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h4",
		Children:   children,
	}
}

func H4_(children ...Node) *HTMLNode {
    return H4(Attr(), children...)
}




func H5(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h5",
		Children:   children,
	}
}

func H5_(children ...Node) *HTMLNode {
    return H5(Attr(), children...)
}




func H6(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "h6",
		Children:   children,
	}
}

func H6_(children ...Node) *HTMLNode {
    return H6(Attr(), children...)
}




func Head(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "head",
		Children:   children,
	}
}

func Head_(children ...Node) *HTMLNode {
    return Head(Attr(), children...)
}




func Header(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "header",
		Children:   children,
	}
}

func Header_(children ...Node) *HTMLNode {
    return Header(Attr(), children...)
}




func Hgroup(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "hgroup",
		Children:   children,
	}
}

func Hgroup_(children ...Node) *HTMLNode {
    return Hgroup(Attr(), children...)
}




func Hr(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "hr",
	}
}

func Hr_() *VoidHTMLNode {
    return Hr(Attr())
}




func Html(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "html",
		Children:   children,
	}
}

func Html_(children ...Node) *HTMLNode {
    return Html(Attr(), children...)
}




func I(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "i",
		Children:   children,
	}
}

func I_(children ...Node) *HTMLNode {
    return I(Attr(), children...)
}




func Iframe(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "iframe",
		Children:   children,
	}
}

func Iframe_(children ...Node) *HTMLNode {
    return Iframe(Attr(), children...)
}




func Img(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "img",
	}
}

func Img_() *VoidHTMLNode {
    return Img(Attr())
}




func Input(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "input",
	}
}

func Input_() *VoidHTMLNode {
    return Input(Attr())
}




func Ins(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ins",
		Children:   children,
	}
}

func Ins_(children ...Node) *HTMLNode {
    return Ins(Attr(), children...)
}




func Isindex(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "isindex",
		Children:   children,
	}
}

func Isindex_(children ...Node) *HTMLNode {
    return Isindex(Attr(), children...)
}




func Kbd(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "kbd",
		Children:   children,
	}
}

func Kbd_(children ...Node) *HTMLNode {
    return Kbd(Attr(), children...)
}




func Keygen(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "keygen",
		Children:   children,
	}
}

func Keygen_(children ...Node) *HTMLNode {
    return Keygen(Attr(), children...)
}




func Label(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "label",
		Children:   children,
	}
}

func Label_(children ...Node) *HTMLNode {
    return Label(Attr(), children...)
}




func Legend(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "legend",
		Children:   children,
	}
}

func Legend_(children ...Node) *HTMLNode {
    return Legend(Attr(), children...)
}




func Li(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "li",
		Children:   children,
	}
}

func Li_(children ...Node) *HTMLNode {
    return Li(Attr(), children...)
}




func Link(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "link",
	}
}

func Link_() *VoidHTMLNode {
    return Link(Attr())
}




func Listing(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "listing",
		Children:   children,
	}
}

func Listing_(children ...Node) *HTMLNode {
    return Listing(Attr(), children...)
}




func Main(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "main",
		Children:   children,
	}
}

func Main_(children ...Node) *HTMLNode {
    return Main(Attr(), children...)
}




func Map(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "map",
		Children:   children,
	}
}

func Map_(children ...Node) *HTMLNode {
    return Map(Attr(), children...)
}




func Mark(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "mark",
		Children:   children,
	}
}

func Mark_(children ...Node) *HTMLNode {
    return Mark(Attr(), children...)
}




func Marquee(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "marquee",
		Children:   children,
	}
}

func Marquee_(children ...Node) *HTMLNode {
    return Marquee(Attr(), children...)
}




func Menu(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "menu",
		Children:   children,
	}
}

func Menu_(children ...Node) *HTMLNode {
    return Menu(Attr(), children...)
}




func Meta(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "meta",
	}
}

func Meta_() *VoidHTMLNode {
    return Meta(Attr())
}




func Meter(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "meter",
		Children:   children,
	}
}

func Meter_(children ...Node) *HTMLNode {
    return Meter(Attr(), children...)
}




func Nav(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "nav",
		Children:   children,
	}
}

func Nav_(children ...Node) *HTMLNode {
    return Nav(Attr(), children...)
}




func Nobr(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "nobr",
		Children:   children,
	}
}

func Nobr_(children ...Node) *HTMLNode {
    return Nobr(Attr(), children...)
}




func Noframes(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "noframes",
		Children:   children,
	}
}

func Noframes_(children ...Node) *HTMLNode {
    return Noframes(Attr(), children...)
}




func Noscript(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "noscript",
		Children:   children,
	}
}

func Noscript_(children ...Node) *HTMLNode {
    return Noscript(Attr(), children...)
}




func Object(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "object",
		Children:   children,
	}
}

func Object_(children ...Node) *HTMLNode {
    return Object(Attr(), children...)
}




func Ol(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ol",
		Children:   children,
	}
}

func Ol_(children ...Node) *HTMLNode {
    return Ol(Attr(), children...)
}




func Optgroup(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "optgroup",
		Children:   children,
	}
}

func Optgroup_(children ...Node) *HTMLNode {
    return Optgroup(Attr(), children...)
}




func Option(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "option",
		Children:   children,
	}
}

func Option_(children ...Node) *HTMLNode {
    return Option(Attr(), children...)
}




func Output(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "output",
		Children:   children,
	}
}

func Output_(children ...Node) *HTMLNode {
    return Output(Attr(), children...)
}




func P(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "p",
		Children:   children,
	}
}

func P_(children ...Node) *HTMLNode {
    return P(Attr(), children...)
}




func Param(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "param",
	}
}

func Param_() *VoidHTMLNode {
    return Param(Attr())
}




func Plaintext(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "plaintext",
		Children:   children,
	}
}

func Plaintext_(children ...Node) *HTMLNode {
    return Plaintext(Attr(), children...)
}




func Pre(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "pre",
		Children:   children,
	}
}

func Pre_(children ...Node) *HTMLNode {
    return Pre(Attr(), children...)
}




func Progress(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "progress",
		Children:   children,
	}
}

func Progress_(children ...Node) *HTMLNode {
    return Progress(Attr(), children...)
}




func Q(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "q",
		Children:   children,
	}
}

func Q_(children ...Node) *HTMLNode {
    return Q(Attr(), children...)
}




func Rp(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "rp",
		Children:   children,
	}
}

func Rp_(children ...Node) *HTMLNode {
    return Rp(Attr(), children...)
}




func Rt(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "rt",
		Children:   children,
	}
}

func Rt_(children ...Node) *HTMLNode {
    return Rt(Attr(), children...)
}




func Ruby(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ruby",
		Children:   children,
	}
}

func Ruby_(children ...Node) *HTMLNode {
    return Ruby(Attr(), children...)
}




func S(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "s",
		Children:   children,
	}
}

func S_(children ...Node) *HTMLNode {
    return S(Attr(), children...)
}




func Samp(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "samp",
		Children:   children,
	}
}

func Samp_(children ...Node) *HTMLNode {
    return Samp(Attr(), children...)
}




func Script(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "script",
		Children:   children,
	}
}

func Script_(children ...Node) *HTMLNode {
    return Script(Attr(), children...)
}




func Section(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "section",
		Children:   children,
	}
}

func Section_(children ...Node) *HTMLNode {
    return Section(Attr(), children...)
}




func Select(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "select",
		Children:   children,
	}
}

func Select_(children ...Node) *HTMLNode {
    return Select(Attr(), children...)
}




func Small(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "small",
		Children:   children,
	}
}

func Small_(children ...Node) *HTMLNode {
    return Small(Attr(), children...)
}




func Source(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "source",
	}
}

func Source_() *VoidHTMLNode {
    return Source(Attr())
}




func Spacer(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "spacer",
		Children:   children,
	}
}

func Spacer_(children ...Node) *HTMLNode {
    return Spacer(Attr(), children...)
}




func Span(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "span",
		Children:   children,
	}
}

func Span_(children ...Node) *HTMLNode {
    return Span(Attr(), children...)
}




func Strike(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "strike",
		Children:   children,
	}
}

func Strike_(children ...Node) *HTMLNode {
    return Strike(Attr(), children...)
}




func Strong(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "strong",
		Children:   children,
	}
}

func Strong_(children ...Node) *HTMLNode {
    return Strong(Attr(), children...)
}




func Style(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "style",
		Children:   children,
	}
}

func Style_(children ...Node) *HTMLNode {
    return Style(Attr(), children...)
}




func Sub(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "sub",
		Children:   children,
	}
}

func Sub_(children ...Node) *HTMLNode {
    return Sub(Attr(), children...)
}




func Summary(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "summary",
		Children:   children,
	}
}

func Summary_(children ...Node) *HTMLNode {
    return Summary(Attr(), children...)
}




func Sup(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "sup",
		Children:   children,
	}
}

func Sup_(children ...Node) *HTMLNode {
    return Sup(Attr(), children...)
}




func Table(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "table",
		Children:   children,
	}
}

func Table_(children ...Node) *HTMLNode {
    return Table(Attr(), children...)
}




func Tbody(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tbody",
		Children:   children,
	}
}

func Tbody_(children ...Node) *HTMLNode {
    return Tbody(Attr(), children...)
}




func Td(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "td",
		Children:   children,
	}
}

func Td_(children ...Node) *HTMLNode {
    return Td(Attr(), children...)
}




func Textarea(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "textarea",
		Children:   children,
	}
}

func Textarea_(children ...Node) *HTMLNode {
    return Textarea(Attr(), children...)
}




func Tfoot(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tfoot",
		Children:   children,
	}
}

func Tfoot_(children ...Node) *HTMLNode {
    return Tfoot(Attr(), children...)
}




func Th(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "th",
		Children:   children,
	}
}

func Th_(children ...Node) *HTMLNode {
    return Th(Attr(), children...)
}




func Thead(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "thead",
		Children:   children,
	}
}

func Thead_(children ...Node) *HTMLNode {
    return Thead(Attr(), children...)
}




func Time(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "time",
		Children:   children,
	}
}

func Time_(children ...Node) *HTMLNode {
    return Time(Attr(), children...)
}




func Title(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "title",
		Children:   children,
	}
}

func Title_(children ...Node) *HTMLNode {
    return Title(Attr(), children...)
}




func Tr(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tr",
		Children:   children,
	}
}

func Tr_(children ...Node) *HTMLNode {
    return Tr(Attr(), children...)
}




func Track(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "track",
	}
}

func Track_() *VoidHTMLNode {
    return Track(Attr())
}




func Tt(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "tt",
		Children:   children,
	}
}

func Tt_(children ...Node) *HTMLNode {
    return Tt(Attr(), children...)
}




func U(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "u",
		Children:   children,
	}
}

func U_(children ...Node) *HTMLNode {
    return U(Attr(), children...)
}




func Ul(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "ul",
		Children:   children,
	}
}

func Ul_(children ...Node) *HTMLNode {
    return Ul(Attr(), children...)
}




func Var(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "var",
		Children:   children,
	}
}

func Var_(children ...Node) *HTMLNode {
    return Var(Attr(), children...)
}




func Video(attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "video",
		Children:   children,
	}
}

func Video_(children ...Node) *HTMLNode {
    return Video(Attr(), children...)
}




func Wbr(attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "wbr",
	}
}

func Wbr_() *VoidHTMLNode {
    return Wbr(Attr())
}



