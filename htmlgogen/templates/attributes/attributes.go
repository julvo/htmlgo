package attributes

import "strings"

type Attribute struct {
    Templ       string
    Data        interface{}
    Name        string
}

// Begin of manually implemented attributes

func Dataset(key, value string) Attribute {
    return Attribute{
        Data: map[string]string{
            "key":      key,
            "value":    value,
        },
        Templ: `{{define "Dataset"}}data-{{.key}}="{{.value}}"{{end}}`,
        Name: "Dataset",
    }
}

func Dataset_(key, value string) Attribute {
		key_ := strings.Replace(key, "-", "_", -1)
    return Attribute{
        Data: map[string]string{},
        Templ: `{{define "Dataset_`+key_+`"}}data-`+key+`="`+value+`"{{end}}`,
        Name: "Dataset_"+key_,
    }
}

// Begin of generated attributes
[[ range .AttributeFuncs ]]

func [[.FuncName]](data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "[[.FuncName]]" }
    if len(templs) == 0 {
        attr.Templ = `{{define "[[.FuncName]]"}}[[.AttrName]]="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "[[.FuncName]]"}}[[.AttrName]]="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func [[.FuncName]]_(values ...string) Attribute {
    return [[.FuncName]](nil, values...)
}
[[ end ]]
