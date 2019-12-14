// package headingid is an extension for the goldmark (http://github.com/yuin/goldmark).
//
// This extension enhances the automatic heading ID generation.
package headingid

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// An IDs interface is a collection of the element ids.
type IDs interface {
	// Generate generates a new element id.
	Generate(value []byte, kind ast.NodeKind) []byte

	// Put puts a given element id to the used ids table.
	Put(value []byte)
}

type ids struct {
	values map[string]bool
}

func NewIDs() IDs {
	return &ids{
		values: map[string]bool{},
	}
}

func (s *ids) Generate(value []byte, kind ast.NodeKind) []byte {
	result := []byte{}
	str := string(value)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	str = strings.Join(strings.FieldsFunc(str, f), "-")
	str = strings.ToLower(str)
	result = []byte(str)
	if len(result) == 0 {
		if kind == ast.KindHeading {
			result = []byte("heading")
		} else {
			result = []byte("id")
		}
	}
	if _, ok := s.values[util.BytesToReadOnlyString(result)]; !ok {
		s.values[util.BytesToReadOnlyString(result)] = true
		return result
	}
	for i := 1; ; i++ {
		newResult := fmt.Sprintf("%s-%d", result, i)
		if _, ok := s.values[newResult]; !ok {
			s.values[newResult] = true
			return []byte(newResult)
		}
	}
}

func (s *ids) Put(value []byte) {
	s.values[util.BytesToReadOnlyString(value)] = true
}
