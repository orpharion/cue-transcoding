package excel

import (
	"cuelang.org/go/cue/ast"
	cueJson "cuelang.org/go/pkg/encoding/json"
	"encoding/json"
	"github.com/orpharion/go-excelAPI/pkg/api"
	"github.com/orpharion/go-excelAPI/pkg/object/dump"
)

// Unmarshal a workbook
func Unmarshal(src []byte) (expr ast.Expr, err error) {
	wb, err := api.CreateWorkbook(src)
	if err != nil {
		return
	}
	wb_ := dump.Workbook(&wb)
	j, err := json.Marshal(wb_)
	if err != nil {
		return
	}
	return cueJson.Unmarshal(j)
}
