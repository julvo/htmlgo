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

func DatasetRaw(key, value string) Attribute {
    return Attribute{
        Data: map[string]string{},
        Templ: `{{define "Dataset"}}data-`+key+`="`+value+`"{{end}}`,
        Name: "Dataset",
    }
}

// Begin of generated attributes


func Accept(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Accept" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Accept"}}accept="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Accept"}}accept="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AcceptRaw(values ...string) Attribute {
    return Accept(nil, values...)
}


func AcceptCharset(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "AcceptCharset" }
    if len(templs) == 0 {
        attr.Templ = `{{define "AcceptCharset"}}accept-charset="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "AcceptCharset"}}accept-charset="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AcceptCharsetRaw(values ...string) Attribute {
    return AcceptCharset(nil, values...)
}


func Accesskey(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Accesskey" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Accesskey"}}accesskey="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Accesskey"}}accesskey="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AccesskeyRaw(values ...string) Attribute {
    return Accesskey(nil, values...)
}


func Action(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Action" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Action"}}action="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Action"}}action="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ActionRaw(values ...string) Attribute {
    return Action(nil, values...)
}


func Align(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Align" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Align"}}align="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Align"}}align="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AlignRaw(values ...string) Attribute {
    return Align(nil, values...)
}


func Alt(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Alt" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Alt"}}alt="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Alt"}}alt="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AltRaw(values ...string) Attribute {
    return Alt(nil, values...)
}


func AriaLabel(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "AriaLabel" }
    if len(templs) == 0 {
        attr.Templ = `{{define "AriaLabel"}}aria-label="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "AriaLabel"}}aria-label="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AriaLabelRaw(values ...string) Attribute {
    return AriaLabel(nil, values...)
}


func Async(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Async" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Async"}}async="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Async"}}async="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AsyncRaw(values ...string) Attribute {
    return Async(nil, values...)
}


func Autocomplete(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Autocomplete" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Autocomplete"}}autocomplete="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Autocomplete"}}autocomplete="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AutocompleteRaw(values ...string) Attribute {
    return Autocomplete(nil, values...)
}


func Autofocus(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Autofocus" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Autofocus"}}autofocus="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Autofocus"}}autofocus="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AutofocusRaw(values ...string) Attribute {
    return Autofocus(nil, values...)
}


func Autoplay(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Autoplay" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Autoplay"}}autoplay="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Autoplay"}}autoplay="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func AutoplayRaw(values ...string) Attribute {
    return Autoplay(nil, values...)
}


func Bgcolor(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Bgcolor" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Bgcolor"}}bgcolor="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Bgcolor"}}bgcolor="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func BgcolorRaw(values ...string) Attribute {
    return Bgcolor(nil, values...)
}


func Border(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Border" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Border"}}border="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Border"}}border="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func BorderRaw(values ...string) Attribute {
    return Border(nil, values...)
}


func Charset(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Charset" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Charset"}}charset="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Charset"}}charset="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func CharsetRaw(values ...string) Attribute {
    return Charset(nil, values...)
}


func Checked(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Checked" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Checked"}}checked="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Checked"}}checked="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func CheckedRaw(values ...string) Attribute {
    return Checked(nil, values...)
}


func Cite(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Cite" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Cite"}}cite="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Cite"}}cite="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func CiteRaw(values ...string) Attribute {
    return Cite(nil, values...)
}


func Class(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Class" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Class"}}class="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Class"}}class="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ClassRaw(values ...string) Attribute {
    return Class(nil, values...)
}


func Color(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Color" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Color"}}color="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Color"}}color="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ColorRaw(values ...string) Attribute {
    return Color(nil, values...)
}


func Cols(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Cols" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Cols"}}cols="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Cols"}}cols="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ColsRaw(values ...string) Attribute {
    return Cols(nil, values...)
}


func Colspan(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Colspan" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Colspan"}}colspan="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Colspan"}}colspan="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ColspanRaw(values ...string) Attribute {
    return Colspan(nil, values...)
}


func Content(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Content" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Content"}}content="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Content"}}content="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ContentRaw(values ...string) Attribute {
    return Content(nil, values...)
}


func Contenteditable(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Contenteditable" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Contenteditable"}}contenteditable="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Contenteditable"}}contenteditable="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ContenteditableRaw(values ...string) Attribute {
    return Contenteditable(nil, values...)
}


func Controls(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Controls" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Controls"}}controls="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Controls"}}controls="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ControlsRaw(values ...string) Attribute {
    return Controls(nil, values...)
}


func Coords(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Coords" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Coords"}}coords="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Coords"}}coords="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func CoordsRaw(values ...string) Attribute {
    return Coords(nil, values...)
}


func Data(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Data" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Data"}}data="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Data"}}data="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DataRaw(values ...string) Attribute {
    return Data(nil, values...)
}


func Datetime(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Datetime" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Datetime"}}datetime="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Datetime"}}datetime="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DatetimeRaw(values ...string) Attribute {
    return Datetime(nil, values...)
}


func Default(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Default" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Default"}}default="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Default"}}default="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DefaultRaw(values ...string) Attribute {
    return Default(nil, values...)
}


func Defer(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Defer" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Defer"}}defer="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Defer"}}defer="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DeferRaw(values ...string) Attribute {
    return Defer(nil, values...)
}


func Dir(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Dir" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Dir"}}dir="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Dir"}}dir="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DirRaw(values ...string) Attribute {
    return Dir(nil, values...)
}


func Dirname(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Dirname" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Dirname"}}dirname="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Dirname"}}dirname="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DirnameRaw(values ...string) Attribute {
    return Dirname(nil, values...)
}


func Disabled(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Disabled" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Disabled"}}disabled="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Disabled"}}disabled="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DisabledRaw(values ...string) Attribute {
    return Disabled(nil, values...)
}


func Download(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Download" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Download"}}download="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Download"}}download="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DownloadRaw(values ...string) Attribute {
    return Download(nil, values...)
}


func Draggable(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Draggable" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Draggable"}}draggable="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Draggable"}}draggable="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DraggableRaw(values ...string) Attribute {
    return Draggable(nil, values...)
}


func Dropzone(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Dropzone" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Dropzone"}}dropzone="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Dropzone"}}dropzone="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func DropzoneRaw(values ...string) Attribute {
    return Dropzone(nil, values...)
}


func Enctype(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Enctype" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Enctype"}}enctype="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Enctype"}}enctype="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func EnctypeRaw(values ...string) Attribute {
    return Enctype(nil, values...)
}


func For(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "For" }
    if len(templs) == 0 {
        attr.Templ = `{{define "For"}}for="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "For"}}for="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ForRaw(values ...string) Attribute {
    return For(nil, values...)
}


func Form(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Form" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Form"}}form="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Form"}}form="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func FormRaw(values ...string) Attribute {
    return Form(nil, values...)
}


func Formaction(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Formaction" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Formaction"}}formaction="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Formaction"}}formaction="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func FormactionRaw(values ...string) Attribute {
    return Formaction(nil, values...)
}


func Headers(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Headers" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Headers"}}headers="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Headers"}}headers="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func HeadersRaw(values ...string) Attribute {
    return Headers(nil, values...)
}


func Height(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Height" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Height"}}height="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Height"}}height="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func HeightRaw(values ...string) Attribute {
    return Height(nil, values...)
}


func Hidden(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Hidden" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Hidden"}}hidden="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Hidden"}}hidden="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func HiddenRaw(values ...string) Attribute {
    return Hidden(nil, values...)
}


func High(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "High" }
    if len(templs) == 0 {
        attr.Templ = `{{define "High"}}high="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "High"}}high="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func HighRaw(values ...string) Attribute {
    return High(nil, values...)
}


func Href(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Href" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Href"}}href="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Href"}}href="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func HrefRaw(values ...string) Attribute {
    return Href(nil, values...)
}


func Hreflang(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Hreflang" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Hreflang"}}hreflang="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Hreflang"}}hreflang="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func HreflangRaw(values ...string) Attribute {
    return Hreflang(nil, values...)
}


func HttpEquiv(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "HttpEquiv" }
    if len(templs) == 0 {
        attr.Templ = `{{define "HttpEquiv"}}http-equiv="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "HttpEquiv"}}http-equiv="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func HttpEquivRaw(values ...string) Attribute {
    return HttpEquiv(nil, values...)
}


func Id(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Id" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Id"}}id="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Id"}}id="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func IdRaw(values ...string) Attribute {
    return Id(nil, values...)
}


func InitialScale(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "InitialScale" }
    if len(templs) == 0 {
        attr.Templ = `{{define "InitialScale"}}initial-scale="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "InitialScale"}}initial-scale="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func InitialScaleRaw(values ...string) Attribute {
    return InitialScale(nil, values...)
}


func Ismap(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ismap" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ismap"}}ismap="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ismap"}}ismap="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func IsmapRaw(values ...string) Attribute {
    return Ismap(nil, values...)
}


func Kind(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Kind" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Kind"}}kind="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Kind"}}kind="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func KindRaw(values ...string) Attribute {
    return Kind(nil, values...)
}


func Label(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Label" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Label"}}label="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Label"}}label="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func LabelRaw(values ...string) Attribute {
    return Label(nil, values...)
}


func Lang(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Lang" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Lang"}}lang="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Lang"}}lang="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func LangRaw(values ...string) Attribute {
    return Lang(nil, values...)
}


func List(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "List" }
    if len(templs) == 0 {
        attr.Templ = `{{define "List"}}list="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "List"}}list="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ListRaw(values ...string) Attribute {
    return List(nil, values...)
}


func Loop(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Loop" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Loop"}}loop="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Loop"}}loop="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func LoopRaw(values ...string) Attribute {
    return Loop(nil, values...)
}


func Low(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Low" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Low"}}low="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Low"}}low="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func LowRaw(values ...string) Attribute {
    return Low(nil, values...)
}


func Max(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Max" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Max"}}max="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Max"}}max="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func MaxRaw(values ...string) Attribute {
    return Max(nil, values...)
}


func Maxlength(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Maxlength" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Maxlength"}}maxlength="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Maxlength"}}maxlength="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func MaxlengthRaw(values ...string) Attribute {
    return Maxlength(nil, values...)
}


func Media(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Media" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Media"}}media="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Media"}}media="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func MediaRaw(values ...string) Attribute {
    return Media(nil, values...)
}


func Method(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Method" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Method"}}method="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Method"}}method="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func MethodRaw(values ...string) Attribute {
    return Method(nil, values...)
}


func Min(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Min" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Min"}}min="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Min"}}min="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func MinRaw(values ...string) Attribute {
    return Min(nil, values...)
}


func Multiple(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Multiple" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Multiple"}}multiple="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Multiple"}}multiple="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func MultipleRaw(values ...string) Attribute {
    return Multiple(nil, values...)
}


func Muted(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Muted" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Muted"}}muted="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Muted"}}muted="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func MutedRaw(values ...string) Attribute {
    return Muted(nil, values...)
}


func Name(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Name" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Name"}}name="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Name"}}name="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func NameRaw(values ...string) Attribute {
    return Name(nil, values...)
}


func Novalidate(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Novalidate" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Novalidate"}}novalidate="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Novalidate"}}novalidate="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func NovalidateRaw(values ...string) Attribute {
    return Novalidate(nil, values...)
}


func Onabort(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onabort" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onabort"}}onabort="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onabort"}}onabort="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnabortRaw(values ...string) Attribute {
    return Onabort(nil, values...)
}


func Onafterprint(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onafterprint" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onafterprint"}}onafterprint="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onafterprint"}}onafterprint="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnafterprintRaw(values ...string) Attribute {
    return Onafterprint(nil, values...)
}


func Onbeforeprint(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onbeforeprint" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onbeforeprint"}}onbeforeprint="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onbeforeprint"}}onbeforeprint="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnbeforeprintRaw(values ...string) Attribute {
    return Onbeforeprint(nil, values...)
}


func Onbeforeunload(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onbeforeunload" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onbeforeunload"}}onbeforeunload="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onbeforeunload"}}onbeforeunload="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnbeforeunloadRaw(values ...string) Attribute {
    return Onbeforeunload(nil, values...)
}


func Onblur(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onblur" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onblur"}}onblur="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onblur"}}onblur="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnblurRaw(values ...string) Attribute {
    return Onblur(nil, values...)
}


func Oncanplay(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oncanplay" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oncanplay"}}oncanplay="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oncanplay"}}oncanplay="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OncanplayRaw(values ...string) Attribute {
    return Oncanplay(nil, values...)
}


func Oncanplaythrough(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oncanplaythrough" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oncanplaythrough"}}oncanplaythrough="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oncanplaythrough"}}oncanplaythrough="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OncanplaythroughRaw(values ...string) Attribute {
    return Oncanplaythrough(nil, values...)
}


func Onchange(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onchange" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onchange"}}onchange="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onchange"}}onchange="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnchangeRaw(values ...string) Attribute {
    return Onchange(nil, values...)
}


func Onclick(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onclick" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onclick"}}onclick="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onclick"}}onclick="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnclickRaw(values ...string) Attribute {
    return Onclick(nil, values...)
}


func Oncontextmenu(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oncontextmenu" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oncontextmenu"}}oncontextmenu="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oncontextmenu"}}oncontextmenu="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OncontextmenuRaw(values ...string) Attribute {
    return Oncontextmenu(nil, values...)
}


func Oncopy(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oncopy" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oncopy"}}oncopy="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oncopy"}}oncopy="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OncopyRaw(values ...string) Attribute {
    return Oncopy(nil, values...)
}


func Oncuechange(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oncuechange" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oncuechange"}}oncuechange="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oncuechange"}}oncuechange="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OncuechangeRaw(values ...string) Attribute {
    return Oncuechange(nil, values...)
}


func Oncut(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oncut" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oncut"}}oncut="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oncut"}}oncut="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OncutRaw(values ...string) Attribute {
    return Oncut(nil, values...)
}


func Ondblclick(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondblclick" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondblclick"}}ondblclick="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondblclick"}}ondblclick="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndblclickRaw(values ...string) Attribute {
    return Ondblclick(nil, values...)
}


func Ondrag(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondrag" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondrag"}}ondrag="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondrag"}}ondrag="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndragRaw(values ...string) Attribute {
    return Ondrag(nil, values...)
}


func Ondragend(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondragend" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondragend"}}ondragend="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondragend"}}ondragend="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndragendRaw(values ...string) Attribute {
    return Ondragend(nil, values...)
}


func Ondragenter(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondragenter" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondragenter"}}ondragenter="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondragenter"}}ondragenter="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndragenterRaw(values ...string) Attribute {
    return Ondragenter(nil, values...)
}


func Ondragleave(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondragleave" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondragleave"}}ondragleave="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondragleave"}}ondragleave="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndragleaveRaw(values ...string) Attribute {
    return Ondragleave(nil, values...)
}


func Ondragover(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondragover" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondragover"}}ondragover="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondragover"}}ondragover="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndragoverRaw(values ...string) Attribute {
    return Ondragover(nil, values...)
}


func Ondragstart(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondragstart" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondragstart"}}ondragstart="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondragstart"}}ondragstart="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndragstartRaw(values ...string) Attribute {
    return Ondragstart(nil, values...)
}


func Ondrop(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondrop" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondrop"}}ondrop="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondrop"}}ondrop="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndropRaw(values ...string) Attribute {
    return Ondrop(nil, values...)
}


func Ondurationchange(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ondurationchange" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ondurationchange"}}ondurationchange="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ondurationchange"}}ondurationchange="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OndurationchangeRaw(values ...string) Attribute {
    return Ondurationchange(nil, values...)
}


func Onemptied(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onemptied" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onemptied"}}onemptied="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onemptied"}}onemptied="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnemptiedRaw(values ...string) Attribute {
    return Onemptied(nil, values...)
}


func Onended(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onended" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onended"}}onended="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onended"}}onended="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnendedRaw(values ...string) Attribute {
    return Onended(nil, values...)
}


func Onerror(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onerror" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onerror"}}onerror="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onerror"}}onerror="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnerrorRaw(values ...string) Attribute {
    return Onerror(nil, values...)
}


func Onfocus(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onfocus" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onfocus"}}onfocus="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onfocus"}}onfocus="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnfocusRaw(values ...string) Attribute {
    return Onfocus(nil, values...)
}


func Onhashchange(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onhashchange" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onhashchange"}}onhashchange="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onhashchange"}}onhashchange="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnhashchangeRaw(values ...string) Attribute {
    return Onhashchange(nil, values...)
}


func Oninput(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oninput" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oninput"}}oninput="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oninput"}}oninput="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OninputRaw(values ...string) Attribute {
    return Oninput(nil, values...)
}


func Oninvalid(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Oninvalid" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Oninvalid"}}oninvalid="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Oninvalid"}}oninvalid="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OninvalidRaw(values ...string) Attribute {
    return Oninvalid(nil, values...)
}


func Onkeydown(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onkeydown" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onkeydown"}}onkeydown="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onkeydown"}}onkeydown="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnkeydownRaw(values ...string) Attribute {
    return Onkeydown(nil, values...)
}


func Onkeypress(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onkeypress" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onkeypress"}}onkeypress="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onkeypress"}}onkeypress="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnkeypressRaw(values ...string) Attribute {
    return Onkeypress(nil, values...)
}


func Onkeyup(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onkeyup" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onkeyup"}}onkeyup="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onkeyup"}}onkeyup="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnkeyupRaw(values ...string) Attribute {
    return Onkeyup(nil, values...)
}


func Onload(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onload" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onload"}}onload="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onload"}}onload="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnloadRaw(values ...string) Attribute {
    return Onload(nil, values...)
}


func Onloadeddata(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onloadeddata" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onloadeddata"}}onloadeddata="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onloadeddata"}}onloadeddata="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnloadeddataRaw(values ...string) Attribute {
    return Onloadeddata(nil, values...)
}


func Onloadedmetadata(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onloadedmetadata" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onloadedmetadata"}}onloadedmetadata="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onloadedmetadata"}}onloadedmetadata="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnloadedmetadataRaw(values ...string) Attribute {
    return Onloadedmetadata(nil, values...)
}


func Onloadstart(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onloadstart" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onloadstart"}}onloadstart="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onloadstart"}}onloadstart="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnloadstartRaw(values ...string) Attribute {
    return Onloadstart(nil, values...)
}


func Onmousedown(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onmousedown" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onmousedown"}}onmousedown="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onmousedown"}}onmousedown="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnmousedownRaw(values ...string) Attribute {
    return Onmousedown(nil, values...)
}


func Onmousemove(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onmousemove" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onmousemove"}}onmousemove="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onmousemove"}}onmousemove="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnmousemoveRaw(values ...string) Attribute {
    return Onmousemove(nil, values...)
}


func Onmouseout(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onmouseout" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onmouseout"}}onmouseout="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onmouseout"}}onmouseout="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnmouseoutRaw(values ...string) Attribute {
    return Onmouseout(nil, values...)
}


func Onmouseover(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onmouseover" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onmouseover"}}onmouseover="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onmouseover"}}onmouseover="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnmouseoverRaw(values ...string) Attribute {
    return Onmouseover(nil, values...)
}


func Onmouseup(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onmouseup" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onmouseup"}}onmouseup="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onmouseup"}}onmouseup="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnmouseupRaw(values ...string) Attribute {
    return Onmouseup(nil, values...)
}


func Onmousewheel(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onmousewheel" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onmousewheel"}}onmousewheel="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onmousewheel"}}onmousewheel="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnmousewheelRaw(values ...string) Attribute {
    return Onmousewheel(nil, values...)
}


func Onoffline(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onoffline" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onoffline"}}onoffline="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onoffline"}}onoffline="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnofflineRaw(values ...string) Attribute {
    return Onoffline(nil, values...)
}


func Ononline(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ononline" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ononline"}}ononline="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ononline"}}ononline="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnonlineRaw(values ...string) Attribute {
    return Ononline(nil, values...)
}


func Onpagehide(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onpagehide" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onpagehide"}}onpagehide="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onpagehide"}}onpagehide="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnpagehideRaw(values ...string) Attribute {
    return Onpagehide(nil, values...)
}


func Onpageshow(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onpageshow" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onpageshow"}}onpageshow="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onpageshow"}}onpageshow="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnpageshowRaw(values ...string) Attribute {
    return Onpageshow(nil, values...)
}


func Onpaste(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onpaste" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onpaste"}}onpaste="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onpaste"}}onpaste="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnpasteRaw(values ...string) Attribute {
    return Onpaste(nil, values...)
}


func Onpause(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onpause" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onpause"}}onpause="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onpause"}}onpause="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnpauseRaw(values ...string) Attribute {
    return Onpause(nil, values...)
}


func Onplay(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onplay" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onplay"}}onplay="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onplay"}}onplay="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnplayRaw(values ...string) Attribute {
    return Onplay(nil, values...)
}


func Onplaying(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onplaying" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onplaying"}}onplaying="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onplaying"}}onplaying="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnplayingRaw(values ...string) Attribute {
    return Onplaying(nil, values...)
}


func Onpopstate(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onpopstate" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onpopstate"}}onpopstate="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onpopstate"}}onpopstate="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnpopstateRaw(values ...string) Attribute {
    return Onpopstate(nil, values...)
}


func Onprogress(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onprogress" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onprogress"}}onprogress="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onprogress"}}onprogress="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnprogressRaw(values ...string) Attribute {
    return Onprogress(nil, values...)
}


func Onratechange(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onratechange" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onratechange"}}onratechange="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onratechange"}}onratechange="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnratechangeRaw(values ...string) Attribute {
    return Onratechange(nil, values...)
}


func Onreset(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onreset" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onreset"}}onreset="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onreset"}}onreset="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnresetRaw(values ...string) Attribute {
    return Onreset(nil, values...)
}


func Onresize(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onresize" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onresize"}}onresize="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onresize"}}onresize="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnresizeRaw(values ...string) Attribute {
    return Onresize(nil, values...)
}


func Onscroll(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onscroll" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onscroll"}}onscroll="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onscroll"}}onscroll="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnscrollRaw(values ...string) Attribute {
    return Onscroll(nil, values...)
}


func Onsearch(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onsearch" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onsearch"}}onsearch="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onsearch"}}onsearch="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnsearchRaw(values ...string) Attribute {
    return Onsearch(nil, values...)
}


func Onseeked(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onseeked" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onseeked"}}onseeked="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onseeked"}}onseeked="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnseekedRaw(values ...string) Attribute {
    return Onseeked(nil, values...)
}


func Onseeking(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onseeking" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onseeking"}}onseeking="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onseeking"}}onseeking="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnseekingRaw(values ...string) Attribute {
    return Onseeking(nil, values...)
}


func Onselect(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onselect" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onselect"}}onselect="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onselect"}}onselect="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnselectRaw(values ...string) Attribute {
    return Onselect(nil, values...)
}


func Onstalled(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onstalled" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onstalled"}}onstalled="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onstalled"}}onstalled="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnstalledRaw(values ...string) Attribute {
    return Onstalled(nil, values...)
}


func Onstorage(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onstorage" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onstorage"}}onstorage="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onstorage"}}onstorage="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnstorageRaw(values ...string) Attribute {
    return Onstorage(nil, values...)
}


func Onsubmit(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onsubmit" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onsubmit"}}onsubmit="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onsubmit"}}onsubmit="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnsubmitRaw(values ...string) Attribute {
    return Onsubmit(nil, values...)
}


func Onsuspend(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onsuspend" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onsuspend"}}onsuspend="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onsuspend"}}onsuspend="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnsuspendRaw(values ...string) Attribute {
    return Onsuspend(nil, values...)
}


func Ontimeupdate(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ontimeupdate" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ontimeupdate"}}ontimeupdate="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ontimeupdate"}}ontimeupdate="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OntimeupdateRaw(values ...string) Attribute {
    return Ontimeupdate(nil, values...)
}


func Ontoggle(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Ontoggle" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Ontoggle"}}ontoggle="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Ontoggle"}}ontoggle="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OntoggleRaw(values ...string) Attribute {
    return Ontoggle(nil, values...)
}


func Onunload(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onunload" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onunload"}}onunload="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onunload"}}onunload="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnunloadRaw(values ...string) Attribute {
    return Onunload(nil, values...)
}


func Onvolumechange(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onvolumechange" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onvolumechange"}}onvolumechange="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onvolumechange"}}onvolumechange="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnvolumechangeRaw(values ...string) Attribute {
    return Onvolumechange(nil, values...)
}


func Onwaiting(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onwaiting" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onwaiting"}}onwaiting="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onwaiting"}}onwaiting="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnwaitingRaw(values ...string) Attribute {
    return Onwaiting(nil, values...)
}


func Onwheel(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Onwheel" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Onwheel"}}onwheel="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Onwheel"}}onwheel="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OnwheelRaw(values ...string) Attribute {
    return Onwheel(nil, values...)
}


func Open(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Open" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Open"}}open="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Open"}}open="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OpenRaw(values ...string) Attribute {
    return Open(nil, values...)
}


func Optimum(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Optimum" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Optimum"}}optimum="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Optimum"}}optimum="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func OptimumRaw(values ...string) Attribute {
    return Optimum(nil, values...)
}


func Pattern(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Pattern" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Pattern"}}pattern="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Pattern"}}pattern="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func PatternRaw(values ...string) Attribute {
    return Pattern(nil, values...)
}


func Placeholder(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Placeholder" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Placeholder"}}placeholder="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Placeholder"}}placeholder="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func PlaceholderRaw(values ...string) Attribute {
    return Placeholder(nil, values...)
}


func Poster(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Poster" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Poster"}}poster="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Poster"}}poster="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func PosterRaw(values ...string) Attribute {
    return Poster(nil, values...)
}


func Preload(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Preload" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Preload"}}preload="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Preload"}}preload="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func PreloadRaw(values ...string) Attribute {
    return Preload(nil, values...)
}


func Readonly(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Readonly" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Readonly"}}readonly="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Readonly"}}readonly="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ReadonlyRaw(values ...string) Attribute {
    return Readonly(nil, values...)
}


func Rel(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Rel" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Rel"}}rel="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Rel"}}rel="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func RelRaw(values ...string) Attribute {
    return Rel(nil, values...)
}


func Required(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Required" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Required"}}required="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Required"}}required="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func RequiredRaw(values ...string) Attribute {
    return Required(nil, values...)
}


func Reversed(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Reversed" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Reversed"}}reversed="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Reversed"}}reversed="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ReversedRaw(values ...string) Attribute {
    return Reversed(nil, values...)
}


func Role(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Role" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Role"}}role="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Role"}}role="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func RoleRaw(values ...string) Attribute {
    return Role(nil, values...)
}


func Rows(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Rows" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Rows"}}rows="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Rows"}}rows="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func RowsRaw(values ...string) Attribute {
    return Rows(nil, values...)
}


func Rowspan(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Rowspan" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Rowspan"}}rowspan="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Rowspan"}}rowspan="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func RowspanRaw(values ...string) Attribute {
    return Rowspan(nil, values...)
}


func Sandbox(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Sandbox" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Sandbox"}}sandbox="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Sandbox"}}sandbox="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SandboxRaw(values ...string) Attribute {
    return Sandbox(nil, values...)
}


func Scope(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Scope" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Scope"}}scope="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Scope"}}scope="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ScopeRaw(values ...string) Attribute {
    return Scope(nil, values...)
}


func Selected(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Selected" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Selected"}}selected="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Selected"}}selected="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SelectedRaw(values ...string) Attribute {
    return Selected(nil, values...)
}


func Shape(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Shape" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Shape"}}shape="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Shape"}}shape="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ShapeRaw(values ...string) Attribute {
    return Shape(nil, values...)
}


func Size(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Size" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Size"}}size="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Size"}}size="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SizeRaw(values ...string) Attribute {
    return Size(nil, values...)
}


func Sizes(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Sizes" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Sizes"}}sizes="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Sizes"}}sizes="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SizesRaw(values ...string) Attribute {
    return Sizes(nil, values...)
}


func Span(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Span" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Span"}}span="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Span"}}span="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SpanRaw(values ...string) Attribute {
    return Span(nil, values...)
}


func Spellcheck(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Spellcheck" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Spellcheck"}}spellcheck="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Spellcheck"}}spellcheck="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SpellcheckRaw(values ...string) Attribute {
    return Spellcheck(nil, values...)
}


func Src(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Src" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Src"}}src="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Src"}}src="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SrcRaw(values ...string) Attribute {
    return Src(nil, values...)
}


func Srcdoc(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Srcdoc" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Srcdoc"}}srcdoc="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Srcdoc"}}srcdoc="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SrcdocRaw(values ...string) Attribute {
    return Srcdoc(nil, values...)
}


func Srclang(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Srclang" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Srclang"}}srclang="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Srclang"}}srclang="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SrclangRaw(values ...string) Attribute {
    return Srclang(nil, values...)
}


func Srcset(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Srcset" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Srcset"}}srcset="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Srcset"}}srcset="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func SrcsetRaw(values ...string) Attribute {
    return Srcset(nil, values...)
}


func Start(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Start" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Start"}}start="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Start"}}start="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func StartRaw(values ...string) Attribute {
    return Start(nil, values...)
}


func Step(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Step" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Step"}}step="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Step"}}step="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func StepRaw(values ...string) Attribute {
    return Step(nil, values...)
}


func Style(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Style" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Style"}}style="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Style"}}style="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func StyleRaw(values ...string) Attribute {
    return Style(nil, values...)
}


func Tabindex(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Tabindex" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Tabindex"}}tabindex="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Tabindex"}}tabindex="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func TabindexRaw(values ...string) Attribute {
    return Tabindex(nil, values...)
}


func Target(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Target" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Target"}}target="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Target"}}target="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func TargetRaw(values ...string) Attribute {
    return Target(nil, values...)
}


func Title(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Title" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Title"}}title="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Title"}}title="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func TitleRaw(values ...string) Attribute {
    return Title(nil, values...)
}


func Translate(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Translate" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Translate"}}translate="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Translate"}}translate="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func TranslateRaw(values ...string) Attribute {
    return Translate(nil, values...)
}


func Type(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Type" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Type"}}type="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Type"}}type="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func TypeRaw(values ...string) Attribute {
    return Type(nil, values...)
}


func Usemap(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Usemap" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Usemap"}}usemap="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Usemap"}}usemap="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func UsemapRaw(values ...string) Attribute {
    return Usemap(nil, values...)
}


func Value(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Value" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Value"}}value="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Value"}}value="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func ValueRaw(values ...string) Attribute {
    return Value(nil, values...)
}


func Width(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Width" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Width"}}width="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Width"}}width="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func WidthRaw(values ...string) Attribute {
    return Width(nil, values...)
}


func Wrap(data interface{}, templs ...string) Attribute {
    attr := Attribute{ Data: data, Name: "Wrap" }
    if len(templs) == 0 {
        attr.Templ = `{{define "Wrap"}}wrap="{{.}}"{{end}}`
    } else {
        attr.Templ = `{{define "Wrap"}}wrap="` + strings.Join(templs, " ") + `"{{end}}`
    }
    return attr
}

func WrapRaw(values ...string) Attribute {
    return Wrap(nil, values...)
}

