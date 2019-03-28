# htmlgo
A library for writing type-safe HTML in Go

## Why?
As much as I like the simplicity and the contextual escaping of `html/template`, it doesn't provide any type safety.
Pipelines (e.g. `{{.DoesNotExist.Values}}`) can produce errors at runtime, which could have been caught during compilation.
Using nested templates with `html/template` can become hard to maintain, as
templates don't define which data they expect.
Moreover, using Go functions to produce HTML elements prevents malformed HTML.

This library is inspired by [Haskell's blaze-html](http://hackage.haskell.org/package/blaze-html).

## Status
* Support for all HTML5 tags and attributes
* Secure, contextual escaping (based on `html/template`)
* Not optimised for performance, expect it to be slower than `html/template`

## API
### Tags
Functions for all HTML tags are part of the top-level package `htmlgo`. The
function signatures are `Tagname(attrs []attributes.Attribute, children ...HTML) HTML`
To omit the first argument, i.e. to create an element without attributes, there
are functions with an underscore suffix `Tagname_(children ...HTML) HTML` to
reduce verbosity.

### Attributes
Functions to create attributes are located in the package `htmlgo/attributes`.
Use `Attr(attrs ...attributes.Attribute)` from `htmlgo` as a less verbose way to
create a slice of attributes. The function signatures are `Attributename(data
interface{}, templates ...string) Attribute`. The `data` will be placed into the
given `templates` using `html/template` and, therefore, follows the same syntax.
Note, as `templates` is a variadic argument, it can be omitted entirely, in
which case a `{{.}}` template is used. Data provided as `data` will be escaped,
whereas the `templates` itself can be used to provide values which shall not be
escaped. Again, there are convenience functions with an underscore suffix to omit the first argument,
which is `data` in this case `Attributename_(templates ...string) Attribute`.
Use the underscore-suffixed attribute functions wisely, as they will not escape
their arguments. A good rule of thumb is that you should never pass a variable
into the suffixed functions, only string literals.

The dataset attributes `data-*` can be added using `Dataset(key, value string)`.

## Example

```golang
package main

import (
    "fmt"

    . "github.com/julvo/htmlgo"
    a "github.com/julvo/htmlgo/attributes"
)

func main() {
    var numberDivs HTML
    for i := 0; i < 3; i++ {
        numberDivs += Div(Attr(a.Style_("font-family:monospace;")),
                          Text(i))
    }

    page :=
        Html5_(
            Head_(),
            Body_(
                H1_(Text("Welcome <script>")),
                numberDivs,
                Div(Attr(a.Dataset("hello", "htmlgo"))),
                Script_(JavaScript("alert('This is escaped');")),
                Script_(JavaScript("This is escaped", "alert({{.}});"))))

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
