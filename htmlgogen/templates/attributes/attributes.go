package attributes

import "strings"

type Attribute string

// Build a slice of type []Attribute for cosmetic purposes
func Attr(attrs ...Attribute) []Attribute {
    return attrs
}

// Begin of manually implemented attributes

func Dataset(key, value string) Attribute {
    return Attribute("data-" + key + `="` + value + `"`)
}

// Begin of generated attributes

[[ range .AttributeFuncs ]]
func [[.FuncName]](values ...string) Attribute {
    return Attribute(`[[.AttrName]]="` + strings.Join(values, " ") + `"`)
}
[[ end ]]
