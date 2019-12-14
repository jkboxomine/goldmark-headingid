goldmark-headingid
=========================

goldmark-headingid is an extension for the [goldmark](http://github.com/yuin/goldmark) 
that enhances the automatic heading ID generation.

This extension overcomes the restrictions of the goldmark's default heading ID logic, including:

* goldmark takes into account ASCII alphanumeric (one-byte) only while generating auto heading IDs, 
simply discarding extended latin characters (2 bytes) and other international characters (3 bytes).
* goldmark considers only a small set of punctuations during the auto heading ID generation. 
As a result, the IDs for headings with other punctuations become less readable and identifiable.

Installation
--------------------

```
go get github.com/jkboxomine/goldmark-headingid
```

Usage
--------------------

```go
import (
	"bytes"
	"fmt"

	"github.com/jkboxomine/goldmark-headingid"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func main() {
	var html bytes.Buffer
	ctx := parser.NewContext(parser.WithIDs(headingid.NewIDs()))
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	mdContent := []byte(`<Markdown content>`)

	if err := md.Convert(mdContent, &html, parser.WithContext(ctx)); err != nil {
		panic(err)
	}
	fmt.Println(html.String())
}
```

License
--------------------
MIT
