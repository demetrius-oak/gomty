package gomty

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {

	h := `
	<html>
		<head>
			<title>My website</title>
		</head>
		<body>		
			<h1 class="title">My First Heading</h1>			
			<p>My first paragraph.</p>		
		</body>
	</html>`

	code := `package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func IndexComponent(children ...g.Node) {

	HTML(
		Head(
			TitleEl(g.Text("My website"))),
		Body(
			H1(
				Class("title"), g.Text("My First Heading")),
			P(g.Text("My first paragraph."))))

}
`
	r := strings.NewReader(h)
	buf := new(bytes.Buffer)

	err := Transform(r, buf, &Options{
		Suffix:  "Component",
		Package: "components",
		Name:    "Index",
	})
	assert.NoError(t, err)
	assert.Equal(t, buf.String(), code)

}
