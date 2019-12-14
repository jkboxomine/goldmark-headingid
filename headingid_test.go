package headingid

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func TestHeadingID(t *testing.T) {
	var html bytes.Buffer
	ctx := parser.NewContext(parser.WithIDs(NewIDs()))
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	mdContent := []byte(`
# Snapshot-Release vX.Y MM/DD/YYYY

1st level heading with punctuation

## Não há quem goste de dor

2nd level heading with extended Latin

### Chaînes@de sites en i18n/

3rd level heading with extended Latin and punctuation

#### 봄 꿀밤  단 꿀밤 v1.0 (2019년-1월-1일)

4th level heading with Korean (CJK), punctuation, and extra whitespace
`)

	expect := []byte(`
<h1 id="snapshot-release-vx-y-mm-dd-yyyy">Snapshot-Release vX.Y MM/DD/YYYY</h1>
<p>1st level heading with punctuation</p>
<h2 id="não-há-quem-goste-de-dor">Não há quem goste de dor</h2>
<p>2nd level heading with extended Latin</p>
<h3 id="chaînes-de-sites-en-i18n">Chaînes@de sites en i18n/</h3>
<p>3rd level heading with extended Latin and punctuation</p>
<h4 id="봄-꿀밤-단-꿀밤-v1-0-2019년-1월-1일">봄 꿀밤  단 꿀밤 v1.0 (2019년-1월-1일)</h4>
<p>4th level heading with Korean (CJK), punctuation, and extra whitespace</p>	
`)

	if err := md.Convert(mdContent, &html, parser.WithContext(ctx)); err != nil {
		t.Fatal(err)
	}

	s := strings.TrimSpace(html.String())
	e := strings.TrimSpace(string(expect))

	if s != e {
		t.Error("got\n", s, "\n\nexpected\n", e)
	}
}
