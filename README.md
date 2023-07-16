# gomty
Transform html into gomponents


## Install

`go install github.com/demetrius-oak/gomty/cmd/gomty@latest`


## Usage

```
gomty -h
Transform html file to gomponents

Usage:
  gomty [file] [flags]

Flags:
  -h, --help             help for gomty
  -n, --name string      Gomponent name (default "index")
  -p, --package string   Package name (default "components")
  -s, --suffix string    Suffix name (default "Component")

``````


## Examples

### Local file:

```bash
gomty ./user_form.html -p forms -n User -s Form > ./forms/user.go
```

### Remote input:

```bash
curl -l https://www.myblog.com/page.html | gomty -n layout
``````

```html
<html>
  <head>
    <title>My website</title>
  </head>
  <body>
    <h1 class="title">My First Heading</h1>
    <p>My first paragraph.</p>
  </body>
</html>
```


will output:

```go
package components

import (
        g "github.com/maragudk/gomponents"
        . "github.com/maragudk/gomponents/html"
)

func LayoutComponent(children ...g.Node) {

        HTML(
                Head(
                        TitleEl(g.Text("My website"))),
                Body(
                        H1(
                                Class("title"), g.Text("My First Heading")),
                        P(g.Text("My first paragraph."))))

}
``````