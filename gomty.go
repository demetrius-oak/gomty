package gomty

import (
	"fmt"
	"io"
	"strings"

	"github.com/dave/jennifer/jen"
	"golang.org/x/net/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Options struct {
	Suffix  string
	Package string
	Name    string
}

// traverse html nodes generating code
func traverse(n *html.Node) *jen.Statement {

	var children []jen.Code
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, traverse(c))
	}

	switch n.Type {
	case html.ElementNode:
		tag := cases.Title(language.Und, cases.NoLower).String(n.Data)
		if tag == "Html" {
			tag = "HTML"
		}
		if tag == "Title" {
			tag = "TitleEl"
		}
		for _, attr := range n.Attr {
			at := cases.Title(language.Und, cases.NoLower).String(attr.Key)
			atCode := jen.Line().Qual("", at).Call(jen.Lit(attr.Val))
			children = append([]jen.Code{atCode}, children...)
		}
		return jen.Line().Qual("github.com/maragudk/gomponents/html", tag).Call(children...)

	case html.DocumentNode:
		return jen.Line().Add(children...).Line()

	case html.TextNode:
		data := strings.TrimSpace(n.Data)
		if len(data) > 0 {
			return jen.Qual("github.com/maragudk/gomponents", "Text").
				Call(
					jen.Lit(data),
				)
		}
	}
	return nil
}

// Transform html content to gomponents
func Transform(reader io.Reader, writer io.Writer, opts *Options) error {

	doc, err := html.Parse(reader)
	if err != nil {
		return err
	}

	root := jen.NewFile(opts.Package)
	root.ImportAlias("github.com/maragudk/gomponents", "g")
	root.ImportAlias("github.com/maragudk/gomponents/html", ".")
	code := traverse(doc)

	funcName := fmt.Sprintf("%s%s", opts.Name, opts.Suffix)
	root.Func().Id(funcName).
		Params(jen.Id("children").Op("...").Qual("", "g.Node")).
		Call().
		Block(code)

	err = root.Render(writer)
	if err != nil {
		return err
	}

	return nil
}
