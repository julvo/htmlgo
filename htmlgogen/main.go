package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	templGlob := filepath.Join(
		build.Default.GOPATH, "src/github.com/julvo/htmlgo/htmlgogen/templates/*.go")
	outputDir := filepath.Join(
		build.Default.GOPATH, "src/github.com/julvo/htmlgo")

	templPaths, err := filepath.Glob(templGlob)
	check(err)

	templData := MakeTemplateData()

	for _, templPath := range templPaths {
		saveAs := filepath.Join(outputDir, filepath.Base(templPath))

		log.Printf("Generating %s...\n", saveAs)

		templ, err := template.New(filepath.Base(templPath)).
			Delims("[[", "]]").
			ParseFiles(templPath)
		check(err)

		err = os.MkdirAll(filepath.Dir(saveAs), 0744)
		check(err)

		f, err := os.Create(saveAs)
		check(err)
		defer f.Close()

		err = templ.Execute(f, templData)
		check(err)
	}
}

type Tag struct {
	Name          string
	IsSelfClosing bool
}

func (t Tag) ToPascalCase() string {
	return toPascalCase(t.Name)
}

type Attribute struct {
	Name      string
	IsBoolean bool
}

func (a Attribute) ToPascalCase() string {
	return toPascalCase(a.Name)
}

func (a Attribute) IsUnique() bool {
	for _, tag := range tags {
		if tag == "srcset" {
			fmt.Println(tag, a.Name)
		}
		if a.Name == tag {
			return false
		}
	}
	return true
}

type TemplateData struct {
	Tags       []Tag
	Attributes []Attribute
}

func MakeTemplateData() TemplateData {
	data := TemplateData{
		Tags:       []Tag{},
		Attributes: []Attribute{},
	}

	for _, tag := range tags {
		_, isSelfClosing := selfClosingTags[tag]
		data.Tags = append(data.Tags, Tag{
			Name:          tag,
			IsSelfClosing: isSelfClosing,
		})
	}
	for _, attr := range attributes {
		data.Attributes = append(data.Attributes, Attribute{
			Name: attr,
			// TODO
			IsBoolean: false,
		})
	}
	return data
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "-")
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
