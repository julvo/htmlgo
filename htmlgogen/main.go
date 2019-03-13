package main

import (
    "fmt"
    "os"
    "text/template"
    "path/filepath"
    "strings"
)

type ElementFunc struct {
    FuncName    string
    TagName     string
}

type VoidElementFunc ElementFunc

type AttributeFunc struct {
    FuncName    string
    AttrName    string
}

type Params struct {
    ElementFuncs        []ElementFunc
    VoidElementFuncs    []VoidElementFunc
    AttributeFuncs      []AttributeFunc
}

const templDir = "htmlgogen/templates"
// special-cases data-*, doctype
func main() {
    templPaths, err := filepath.Glob(filepath.Join(templDir, "*.go"))
    check(err)
    moreTemplPaths, err := filepath.Glob(filepath.Join(templDir, "/*/*.go"))
    check(err)
    templPaths = append(templPaths, moreTemplPaths...)

    params := NewParams()

    for _, templPath := range templPaths {
        saveAs, err := filepath.Rel(templDir, templPath)
        check(err)

        fmt.Printf("Generating %s...\n", saveAs)

        templ := template.Must(template.ParseFiles(templPath))
        err = os.MkdirAll(filepath.Dir(saveAs), 0744)
        check(err)

        f, err := os.Create(saveAs)
        check(err)
        defer f.Close()

        err = templ.Execute(f, params)
        check(err)
    }

}

func NewParams() Params {
    ps := Params{
        []ElementFunc{},
        []VoidElementFunc{},
        []AttributeFunc{},
    }

    for _, tag := range tags {
        if _, ok := selfClosingTags[tag]; ok {
            ps.VoidElementFuncs = append(ps.VoidElementFuncs, VoidElementFunc{
                                             FuncName:  GetFuncName(tag),
                                             TagName:   tag,
                                         })
        } else {
            ps.ElementFuncs = append(ps.ElementFuncs, ElementFunc{
                                             FuncName:  GetFuncName(tag),
                                             TagName:   tag,
                                         })
        }
    }
    for _, attr := range attributes {
            ps.AttributeFuncs = append(ps.AttributeFuncs, AttributeFunc{
                                             FuncName:  GetFuncName(attr),
                                             AttrName:  attr,
                                         })
    }
    return ps
}

func GetFuncName(s string) string {
    parts := strings.Split(s, "-")
    for i, p := range parts {
        parts[i] = strings.Title(p)
    }
    return strings.Join(parts, "")
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}
