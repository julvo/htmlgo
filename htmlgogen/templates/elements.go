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
[[range .Attributes]]
	[[.ToPascalCase]] string
[[end]]
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
[[range .Attributes]]
		a.[[.ToPascalCase]] = strings.Replace(a.[[.ToPascalCase]], "{{.", "{{."+placeholder, -1)
[[end]]

		for k, v := range a.templData {
			vals[placeholder+k] = v
		}
	}

	switch {
[[range .Attributes]]
    case a.[[.ToPascalCase]] != "":
		templ.WriteString(` [[.Name]]="` + a.[[.ToPascalCase]] + `"`)
[[end]]
	case a.DisabledBoolean:
		templ.WriteString(" disabled")
	}
	for k, v := range a.Dataset {
		templ.WriteString(" data-" + k + `="` + v + `"`)
	}

}

// Begin of generated elements

[[ range .Tags ]]
[[if .IsSelfClosing]]
func [[.ToPascalCase]](attributes Attr) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "[[.Name]]",
	}
}

func [[.ToPascalCase]]_() *VoidHTMLNode {
    return [[.ToPascalCase]](Attr{})
}

[[else]]
func [[.ToPascalCase]](attributes Attr, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "[[.Name]]",
		Children:   children,
	}
}

func [[.ToPascalCase]]_(children ...Node) *HTMLNode {
    return [[.ToPascalCase]](Attr{}, children...)
}

[[ end ]]
[[ end ]]