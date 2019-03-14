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


func Accept(values ...string) Attribute {
    return Attribute(`accept="` + strings.Join(values, " ") + `"`)
}

func AcceptCharset(values ...string) Attribute {
    return Attribute(`accept-charset="` + strings.Join(values, " ") + `"`)
}

func Accesskey(values ...string) Attribute {
    return Attribute(`accesskey="` + strings.Join(values, " ") + `"`)
}

func Action(values ...string) Attribute {
    return Attribute(`action="` + strings.Join(values, " ") + `"`)
}

func Align(values ...string) Attribute {
    return Attribute(`align="` + strings.Join(values, " ") + `"`)
}

func Alt(values ...string) Attribute {
    return Attribute(`alt="` + strings.Join(values, " ") + `"`)
}

func Async(values ...string) Attribute {
    return Attribute(`async="` + strings.Join(values, " ") + `"`)
}

func Autocomplete(values ...string) Attribute {
    return Attribute(`autocomplete="` + strings.Join(values, " ") + `"`)
}

func Autofocus(values ...string) Attribute {
    return Attribute(`autofocus="` + strings.Join(values, " ") + `"`)
}

func Autoplay(values ...string) Attribute {
    return Attribute(`autoplay="` + strings.Join(values, " ") + `"`)
}

func Bgcolor(values ...string) Attribute {
    return Attribute(`bgcolor="` + strings.Join(values, " ") + `"`)
}

func Border(values ...string) Attribute {
    return Attribute(`border="` + strings.Join(values, " ") + `"`)
}

func Charset(values ...string) Attribute {
    return Attribute(`charset="` + strings.Join(values, " ") + `"`)
}

func Checked(values ...string) Attribute {
    return Attribute(`checked="` + strings.Join(values, " ") + `"`)
}

func Cite(values ...string) Attribute {
    return Attribute(`cite="` + strings.Join(values, " ") + `"`)
}

func Class(values ...string) Attribute {
    return Attribute(`class="` + strings.Join(values, " ") + `"`)
}

func Color(values ...string) Attribute {
    return Attribute(`color="` + strings.Join(values, " ") + `"`)
}

func Cols(values ...string) Attribute {
    return Attribute(`cols="` + strings.Join(values, " ") + `"`)
}

func Colspan(values ...string) Attribute {
    return Attribute(`colspan="` + strings.Join(values, " ") + `"`)
}

func Content(values ...string) Attribute {
    return Attribute(`content="` + strings.Join(values, " ") + `"`)
}

func Contenteditable(values ...string) Attribute {
    return Attribute(`contenteditable="` + strings.Join(values, " ") + `"`)
}

func Controls(values ...string) Attribute {
    return Attribute(`controls="` + strings.Join(values, " ") + `"`)
}

func Coords(values ...string) Attribute {
    return Attribute(`coords="` + strings.Join(values, " ") + `"`)
}

func Data(values ...string) Attribute {
    return Attribute(`data="` + strings.Join(values, " ") + `"`)
}

func Datetime(values ...string) Attribute {
    return Attribute(`datetime="` + strings.Join(values, " ") + `"`)
}

func Default(values ...string) Attribute {
    return Attribute(`default="` + strings.Join(values, " ") + `"`)
}

func Defer(values ...string) Attribute {
    return Attribute(`defer="` + strings.Join(values, " ") + `"`)
}

func Dir(values ...string) Attribute {
    return Attribute(`dir="` + strings.Join(values, " ") + `"`)
}

func Dirname(values ...string) Attribute {
    return Attribute(`dirname="` + strings.Join(values, " ") + `"`)
}

func Disabled(values ...string) Attribute {
    return Attribute(`disabled="` + strings.Join(values, " ") + `"`)
}

func Download(values ...string) Attribute {
    return Attribute(`download="` + strings.Join(values, " ") + `"`)
}

func Draggable(values ...string) Attribute {
    return Attribute(`draggable="` + strings.Join(values, " ") + `"`)
}

func Dropzone(values ...string) Attribute {
    return Attribute(`dropzone="` + strings.Join(values, " ") + `"`)
}

func Enctype(values ...string) Attribute {
    return Attribute(`enctype="` + strings.Join(values, " ") + `"`)
}

func For(values ...string) Attribute {
    return Attribute(`for="` + strings.Join(values, " ") + `"`)
}

func Form(values ...string) Attribute {
    return Attribute(`form="` + strings.Join(values, " ") + `"`)
}

func Formaction(values ...string) Attribute {
    return Attribute(`formaction="` + strings.Join(values, " ") + `"`)
}

func Headers(values ...string) Attribute {
    return Attribute(`headers="` + strings.Join(values, " ") + `"`)
}

func Height(values ...string) Attribute {
    return Attribute(`height="` + strings.Join(values, " ") + `"`)
}

func Hidden(values ...string) Attribute {
    return Attribute(`hidden="` + strings.Join(values, " ") + `"`)
}

func High(values ...string) Attribute {
    return Attribute(`high="` + strings.Join(values, " ") + `"`)
}

func Href(values ...string) Attribute {
    return Attribute(`href="` + strings.Join(values, " ") + `"`)
}

func Hreflang(values ...string) Attribute {
    return Attribute(`hreflang="` + strings.Join(values, " ") + `"`)
}

func HttpEquiv(values ...string) Attribute {
    return Attribute(`http-equiv="` + strings.Join(values, " ") + `"`)
}

func Id(values ...string) Attribute {
    return Attribute(`id="` + strings.Join(values, " ") + `"`)
}

func Ismap(values ...string) Attribute {
    return Attribute(`ismap="` + strings.Join(values, " ") + `"`)
}

func Kind(values ...string) Attribute {
    return Attribute(`kind="` + strings.Join(values, " ") + `"`)
}

func Label(values ...string) Attribute {
    return Attribute(`label="` + strings.Join(values, " ") + `"`)
}

func Lang(values ...string) Attribute {
    return Attribute(`lang="` + strings.Join(values, " ") + `"`)
}

func List(values ...string) Attribute {
    return Attribute(`list="` + strings.Join(values, " ") + `"`)
}

func Loop(values ...string) Attribute {
    return Attribute(`loop="` + strings.Join(values, " ") + `"`)
}

func Low(values ...string) Attribute {
    return Attribute(`low="` + strings.Join(values, " ") + `"`)
}

func Max(values ...string) Attribute {
    return Attribute(`max="` + strings.Join(values, " ") + `"`)
}

func Maxlength(values ...string) Attribute {
    return Attribute(`maxlength="` + strings.Join(values, " ") + `"`)
}

func Media(values ...string) Attribute {
    return Attribute(`media="` + strings.Join(values, " ") + `"`)
}

func Method(values ...string) Attribute {
    return Attribute(`method="` + strings.Join(values, " ") + `"`)
}

func Min(values ...string) Attribute {
    return Attribute(`min="` + strings.Join(values, " ") + `"`)
}

func Multiple(values ...string) Attribute {
    return Attribute(`multiple="` + strings.Join(values, " ") + `"`)
}

func Muted(values ...string) Attribute {
    return Attribute(`muted="` + strings.Join(values, " ") + `"`)
}

func Name(values ...string) Attribute {
    return Attribute(`name="` + strings.Join(values, " ") + `"`)
}

func Novalidate(values ...string) Attribute {
    return Attribute(`novalidate="` + strings.Join(values, " ") + `"`)
}

func Onabort(values ...string) Attribute {
    return Attribute(`onabort="` + strings.Join(values, " ") + `"`)
}

func Onafterprint(values ...string) Attribute {
    return Attribute(`onafterprint="` + strings.Join(values, " ") + `"`)
}

func Onbeforeprint(values ...string) Attribute {
    return Attribute(`onbeforeprint="` + strings.Join(values, " ") + `"`)
}

func Onbeforeunload(values ...string) Attribute {
    return Attribute(`onbeforeunload="` + strings.Join(values, " ") + `"`)
}

func Onblur(values ...string) Attribute {
    return Attribute(`onblur="` + strings.Join(values, " ") + `"`)
}

func Oncanplay(values ...string) Attribute {
    return Attribute(`oncanplay="` + strings.Join(values, " ") + `"`)
}

func Oncanplaythrough(values ...string) Attribute {
    return Attribute(`oncanplaythrough="` + strings.Join(values, " ") + `"`)
}

func Onchange(values ...string) Attribute {
    return Attribute(`onchange="` + strings.Join(values, " ") + `"`)
}

func Onclick(values ...string) Attribute {
    return Attribute(`onclick="` + strings.Join(values, " ") + `"`)
}

func Oncontextmenu(values ...string) Attribute {
    return Attribute(`oncontextmenu="` + strings.Join(values, " ") + `"`)
}

func Oncopy(values ...string) Attribute {
    return Attribute(`oncopy="` + strings.Join(values, " ") + `"`)
}

func Oncuechange(values ...string) Attribute {
    return Attribute(`oncuechange="` + strings.Join(values, " ") + `"`)
}

func Oncut(values ...string) Attribute {
    return Attribute(`oncut="` + strings.Join(values, " ") + `"`)
}

func Ondblclick(values ...string) Attribute {
    return Attribute(`ondblclick="` + strings.Join(values, " ") + `"`)
}

func Ondrag(values ...string) Attribute {
    return Attribute(`ondrag="` + strings.Join(values, " ") + `"`)
}

func Ondragend(values ...string) Attribute {
    return Attribute(`ondragend="` + strings.Join(values, " ") + `"`)
}

func Ondragenter(values ...string) Attribute {
    return Attribute(`ondragenter="` + strings.Join(values, " ") + `"`)
}

func Ondragleave(values ...string) Attribute {
    return Attribute(`ondragleave="` + strings.Join(values, " ") + `"`)
}

func Ondragover(values ...string) Attribute {
    return Attribute(`ondragover="` + strings.Join(values, " ") + `"`)
}

func Ondragstart(values ...string) Attribute {
    return Attribute(`ondragstart="` + strings.Join(values, " ") + `"`)
}

func Ondrop(values ...string) Attribute {
    return Attribute(`ondrop="` + strings.Join(values, " ") + `"`)
}

func Ondurationchange(values ...string) Attribute {
    return Attribute(`ondurationchange="` + strings.Join(values, " ") + `"`)
}

func Onemptied(values ...string) Attribute {
    return Attribute(`onemptied="` + strings.Join(values, " ") + `"`)
}

func Onended(values ...string) Attribute {
    return Attribute(`onended="` + strings.Join(values, " ") + `"`)
}

func Onerror(values ...string) Attribute {
    return Attribute(`onerror="` + strings.Join(values, " ") + `"`)
}

func Onfocus(values ...string) Attribute {
    return Attribute(`onfocus="` + strings.Join(values, " ") + `"`)
}

func Onhashchange(values ...string) Attribute {
    return Attribute(`onhashchange="` + strings.Join(values, " ") + `"`)
}

func Oninput(values ...string) Attribute {
    return Attribute(`oninput="` + strings.Join(values, " ") + `"`)
}

func Oninvalid(values ...string) Attribute {
    return Attribute(`oninvalid="` + strings.Join(values, " ") + `"`)
}

func Onkeydown(values ...string) Attribute {
    return Attribute(`onkeydown="` + strings.Join(values, " ") + `"`)
}

func Onkeypress(values ...string) Attribute {
    return Attribute(`onkeypress="` + strings.Join(values, " ") + `"`)
}

func Onkeyup(values ...string) Attribute {
    return Attribute(`onkeyup="` + strings.Join(values, " ") + `"`)
}

func Onload(values ...string) Attribute {
    return Attribute(`onload="` + strings.Join(values, " ") + `"`)
}

func Onloadeddata(values ...string) Attribute {
    return Attribute(`onloadeddata="` + strings.Join(values, " ") + `"`)
}

func Onloadedmetadata(values ...string) Attribute {
    return Attribute(`onloadedmetadata="` + strings.Join(values, " ") + `"`)
}

func Onloadstart(values ...string) Attribute {
    return Attribute(`onloadstart="` + strings.Join(values, " ") + `"`)
}

func Onmousedown(values ...string) Attribute {
    return Attribute(`onmousedown="` + strings.Join(values, " ") + `"`)
}

func Onmousemove(values ...string) Attribute {
    return Attribute(`onmousemove="` + strings.Join(values, " ") + `"`)
}

func Onmouseout(values ...string) Attribute {
    return Attribute(`onmouseout="` + strings.Join(values, " ") + `"`)
}

func Onmouseover(values ...string) Attribute {
    return Attribute(`onmouseover="` + strings.Join(values, " ") + `"`)
}

func Onmouseup(values ...string) Attribute {
    return Attribute(`onmouseup="` + strings.Join(values, " ") + `"`)
}

func Onmousewheel(values ...string) Attribute {
    return Attribute(`onmousewheel="` + strings.Join(values, " ") + `"`)
}

func Onoffline(values ...string) Attribute {
    return Attribute(`onoffline="` + strings.Join(values, " ") + `"`)
}

func Ononline(values ...string) Attribute {
    return Attribute(`ononline="` + strings.Join(values, " ") + `"`)
}

func Onpagehide(values ...string) Attribute {
    return Attribute(`onpagehide="` + strings.Join(values, " ") + `"`)
}

func Onpageshow(values ...string) Attribute {
    return Attribute(`onpageshow="` + strings.Join(values, " ") + `"`)
}

func Onpaste(values ...string) Attribute {
    return Attribute(`onpaste="` + strings.Join(values, " ") + `"`)
}

func Onpause(values ...string) Attribute {
    return Attribute(`onpause="` + strings.Join(values, " ") + `"`)
}

func Onplay(values ...string) Attribute {
    return Attribute(`onplay="` + strings.Join(values, " ") + `"`)
}

func Onplaying(values ...string) Attribute {
    return Attribute(`onplaying="` + strings.Join(values, " ") + `"`)
}

func Onpopstate(values ...string) Attribute {
    return Attribute(`onpopstate="` + strings.Join(values, " ") + `"`)
}

func Onprogress(values ...string) Attribute {
    return Attribute(`onprogress="` + strings.Join(values, " ") + `"`)
}

func Onratechange(values ...string) Attribute {
    return Attribute(`onratechange="` + strings.Join(values, " ") + `"`)
}

func Onreset(values ...string) Attribute {
    return Attribute(`onreset="` + strings.Join(values, " ") + `"`)
}

func Onresize(values ...string) Attribute {
    return Attribute(`onresize="` + strings.Join(values, " ") + `"`)
}

func Onscroll(values ...string) Attribute {
    return Attribute(`onscroll="` + strings.Join(values, " ") + `"`)
}

func Onsearch(values ...string) Attribute {
    return Attribute(`onsearch="` + strings.Join(values, " ") + `"`)
}

func Onseeked(values ...string) Attribute {
    return Attribute(`onseeked="` + strings.Join(values, " ") + `"`)
}

func Onseeking(values ...string) Attribute {
    return Attribute(`onseeking="` + strings.Join(values, " ") + `"`)
}

func Onselect(values ...string) Attribute {
    return Attribute(`onselect="` + strings.Join(values, " ") + `"`)
}

func Onstalled(values ...string) Attribute {
    return Attribute(`onstalled="` + strings.Join(values, " ") + `"`)
}

func Onstorage(values ...string) Attribute {
    return Attribute(`onstorage="` + strings.Join(values, " ") + `"`)
}

func Onsubmit(values ...string) Attribute {
    return Attribute(`onsubmit="` + strings.Join(values, " ") + `"`)
}

func Onsuspend(values ...string) Attribute {
    return Attribute(`onsuspend="` + strings.Join(values, " ") + `"`)
}

func Ontimeupdate(values ...string) Attribute {
    return Attribute(`ontimeupdate="` + strings.Join(values, " ") + `"`)
}

func Ontoggle(values ...string) Attribute {
    return Attribute(`ontoggle="` + strings.Join(values, " ") + `"`)
}

func Onunload(values ...string) Attribute {
    return Attribute(`onunload="` + strings.Join(values, " ") + `"`)
}

func Onvolumechange(values ...string) Attribute {
    return Attribute(`onvolumechange="` + strings.Join(values, " ") + `"`)
}

func Onwaiting(values ...string) Attribute {
    return Attribute(`onwaiting="` + strings.Join(values, " ") + `"`)
}

func Onwheel(values ...string) Attribute {
    return Attribute(`onwheel="` + strings.Join(values, " ") + `"`)
}

func Open(values ...string) Attribute {
    return Attribute(`open="` + strings.Join(values, " ") + `"`)
}

func Optimum(values ...string) Attribute {
    return Attribute(`optimum="` + strings.Join(values, " ") + `"`)
}

func Pattern(values ...string) Attribute {
    return Attribute(`pattern="` + strings.Join(values, " ") + `"`)
}

func Placeholder(values ...string) Attribute {
    return Attribute(`placeholder="` + strings.Join(values, " ") + `"`)
}

func Poster(values ...string) Attribute {
    return Attribute(`poster="` + strings.Join(values, " ") + `"`)
}

func Preload(values ...string) Attribute {
    return Attribute(`preload="` + strings.Join(values, " ") + `"`)
}

func Readonly(values ...string) Attribute {
    return Attribute(`readonly="` + strings.Join(values, " ") + `"`)
}

func Rel(values ...string) Attribute {
    return Attribute(`rel="` + strings.Join(values, " ") + `"`)
}

func Required(values ...string) Attribute {
    return Attribute(`required="` + strings.Join(values, " ") + `"`)
}

func Reversed(values ...string) Attribute {
    return Attribute(`reversed="` + strings.Join(values, " ") + `"`)
}

func Rows(values ...string) Attribute {
    return Attribute(`rows="` + strings.Join(values, " ") + `"`)
}

func Rowspan(values ...string) Attribute {
    return Attribute(`rowspan="` + strings.Join(values, " ") + `"`)
}

func Sandbox(values ...string) Attribute {
    return Attribute(`sandbox="` + strings.Join(values, " ") + `"`)
}

func Scope(values ...string) Attribute {
    return Attribute(`scope="` + strings.Join(values, " ") + `"`)
}

func Selected(values ...string) Attribute {
    return Attribute(`selected="` + strings.Join(values, " ") + `"`)
}

func Shape(values ...string) Attribute {
    return Attribute(`shape="` + strings.Join(values, " ") + `"`)
}

func Size(values ...string) Attribute {
    return Attribute(`size="` + strings.Join(values, " ") + `"`)
}

func Sizes(values ...string) Attribute {
    return Attribute(`sizes="` + strings.Join(values, " ") + `"`)
}

func Span(values ...string) Attribute {
    return Attribute(`span="` + strings.Join(values, " ") + `"`)
}

func Spellcheck(values ...string) Attribute {
    return Attribute(`spellcheck="` + strings.Join(values, " ") + `"`)
}

func Src(values ...string) Attribute {
    return Attribute(`src="` + strings.Join(values, " ") + `"`)
}

func Srcdoc(values ...string) Attribute {
    return Attribute(`srcdoc="` + strings.Join(values, " ") + `"`)
}

func Srclang(values ...string) Attribute {
    return Attribute(`srclang="` + strings.Join(values, " ") + `"`)
}

func Srcset(values ...string) Attribute {
    return Attribute(`srcset="` + strings.Join(values, " ") + `"`)
}

func Start(values ...string) Attribute {
    return Attribute(`start="` + strings.Join(values, " ") + `"`)
}

func Step(values ...string) Attribute {
    return Attribute(`step="` + strings.Join(values, " ") + `"`)
}

func Style(values ...string) Attribute {
    return Attribute(`style="` + strings.Join(values, " ") + `"`)
}

func Tabindex(values ...string) Attribute {
    return Attribute(`tabindex="` + strings.Join(values, " ") + `"`)
}

func Target(values ...string) Attribute {
    return Attribute(`target="` + strings.Join(values, " ") + `"`)
}

func Title(values ...string) Attribute {
    return Attribute(`title="` + strings.Join(values, " ") + `"`)
}

func Translate(values ...string) Attribute {
    return Attribute(`translate="` + strings.Join(values, " ") + `"`)
}

func Type(values ...string) Attribute {
    return Attribute(`type="` + strings.Join(values, " ") + `"`)
}

func Usemap(values ...string) Attribute {
    return Attribute(`usemap="` + strings.Join(values, " ") + `"`)
}

func Value(values ...string) Attribute {
    return Attribute(`value="` + strings.Join(values, " ") + `"`)
}

func Width(values ...string) Attribute {
    return Attribute(`width="` + strings.Join(values, " ") + `"`)
}

func Wrap(values ...string) Attribute {
    return Attribute(`wrap="` + strings.Join(values, " ") + `"`)
}

