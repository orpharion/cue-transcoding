package xml

import (
	"bytes"
	"cuelang.org/go/cue/ast"
	pkgJson "cuelang.org/go/pkg/encoding/json"
	"github.com/basgys/goxml2json"
)

// Unmarshal parses the XML-encoded data, without a reference schema.
func Unmarshal(src []byte) (ast.Expr, error) {
	j, err := xml2json.Convert(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	return pkgJson.Unmarshal(j.Bytes())
}
