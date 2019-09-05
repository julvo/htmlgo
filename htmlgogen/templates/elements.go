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

[[range .Attributes]]
func (a Attributes) [[.ToPascalCase]](value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a.Slice = append(a.Slice, Attribute{
		Name: "[[.Name]]",
		Data: value,
		Templ: templ,
	})
	return a
}
[[end]]

[[range .Attributes]]
func (a Attributes) [[.ToPascalCase]]_(templ string) Attributes {
	a.Slice = append(a.Slice, Attribute{
		Name: "[[.Name]]",
		Data: nil,
		Templ: templ,
	})
	return a
}
[[end]]

[[range .Attributes]]
[[if .IsUnique]]
func [[.ToPascalCase]](value interface{}, templates ...string) Attributes {
	templ := "{{.}}"
	if len(templates) > 0 {
		templ = templates[0]
	}
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "[[.Name]]",
				Data: value,
				Templ: templ,
			},
		},
	}
	return a
}
[[end]]
[[end]]

[[range .Attributes]]
[[if .IsUnique]]
func [[.ToPascalCase]]_(templ string) Attributes {
	a := Attributes{
		Slice: []Attribute{
			Attribute{
				Name: "[[.Name]]",
				Data: nil,
				Templ: templ,
			},
		},
	}
	return a
}
[[end]]
[[end]]

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

[[ range .Tags ]]
[[if .IsSelfClosing]]
func [[.ToPascalCase]](attributes Attributes) *VoidHTMLNode {
	return &VoidHTMLNode{
		Attributes: attributes,
		Tag:        "[[.Name]]",
	}
}

func [[.ToPascalCase]]_() *VoidHTMLNode {
    return [[.ToPascalCase]](Attr())
}

[[else]]
func [[.ToPascalCase]](attributes Attributes, children ...Node) *HTMLNode {
	return &HTMLNode{
		Attributes: attributes,
		Tag:        "[[.Name]]",
		Children:   children,
	}
}

func [[.ToPascalCase]]_(children ...Node) *HTMLNode {
    return [[.ToPascalCase]](Attr(), children...)
}

[[ end ]]
[[ end ]]
