# htmlgo
A library for writing typesafe HTML in Go

## Why?
As much as I like the simplicity of `html/template`, it doesn't provide any typesafety.
Pipelines (e.g. `{{.DoesNotExist.Values}}`) can produce errors at runtime, which could have been caught during compilation.
Moreover, using Go functions to produce HTML elements prevents malformed HTML.

This library is inspired by [Haskell's blaze-html](http://hackage.haskell.org/package/blaze-html).

## Example

```golang
package main

import (
    "fmt"

    h "github.com/julvo/htmlgo"
    a "github.com/julvo/htmlgo/attributes"
)

func main() {
    var numberDivs h.HTML
    for i := 0; i < 3; i++ {
        numberDivs += h.Div(h.Attr(a.StyleRaw("font-family:monospace;")),
                            h.Text(i))
    }

    page :=
        h.Html5_(
            h.Head_(),
            h.Body_(
                h.H1_(h.Text("Welcome <script>")),
                numberDivs,
                h.Div(h.Attr(a.Dataset("hello", "htmlgo"))),
                h.Script_(h.JavaScript("alert('This is escaped');")),
                h.Script_(h.JavaScript("This is escaped", "alert({{.}});"))))

    fmt.Println(page)
}

```
will output:

```html
<!DOCTYPE HTML>
<html>
  <head>
  </head>
  <body>
    <h1>
      Welcome &lt;script&gt;
    </h1>
    <div style="font-family:monospace;">
      0
    </div>
    <div style="font-family:monospace;">
      1
    </div>
    <div style="font-family:monospace;">
      2
    </div>
    <div data-hello="htmlgo">
    </div>
    <script>
      "alert('This is escaped');"
    </script>
    <script>
      alert("This is escaped");
    </script>
  </body>
</html>
```
