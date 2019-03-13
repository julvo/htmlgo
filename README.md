# htmlgo
A library for writing typesafe HTML in Go

## Why?
As much as I like the simplicity of `html/template`, it doesn't provide any typesafety.
Pipelines (e.g. `{{.DoesNotExist.Values}}`) can produce errors at runtime, which could have been caught during compilation.
Moreover, using Go functions to produce HTML elements prevents malformed HTML.

This library is inspired by [Haskell's blaze-html](http://hackage.haskell.org/package/blaze-html).

## Example

```
package main

import (
    "fmt"

    h "github.com/julvo/htmlgo"
    a "github.com/julvo/htmlgo/attributes"
)

func main() {
    doc := h.DoctypeHtml5 +
           h.Html(a.Attr(),
              h.Head(a.Attr()),
              h.Body(a.Attr(),
                h.Div(a.Attr(a.Class("columns")),
                    h.Div(a.Attr(a.Class("column", "is-narrow"), a.Data_("url", "https://hello.world")),
                        h.Text("Hello World"),
                    ),
                  h.Div(a.Attr(a.Class("column")),
                        h.Text(`<script>alert("Text is escaped")</script>`),
                    ),
                  h.Div(a.Attr(a.Class("column"))),
                ),
              ),
           )
    fmt.Println(doc)
}

```
will output:

```
<!DOCTYPE HTML>
<html>
  <head>
  </head>
  <body>
    <div class="columns">
      <div class="column is-narrow" data-url="https://hello.world">
        Hello World
      </div>
      <div class="column">
        &lt;script&gt;alert(&#34;Text is escaped&#34;)&lt;/script&gt;
      </div>
      <div class="column">
      </div>
    </div>
  </body>
</html>
```
