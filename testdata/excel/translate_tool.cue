package translate

import (
	"encoding/excel"
	"encoding/json"
	"tool/file"
)

command: TranslateExcel: {
	Input: {
		excel: "input.xlsx"
		json: "output.json"
	}
	ExcelBytes:  {
		file.Read
		filename: Input.excel
		contents: bytes
	}
	Unmarshalled: excel.Unmarshal(ExcelBytes.contents)
	Output: file.Create & {
		filename: Input.json
		contents: json.Indent(json.Marshal(Unmarshalled), "", "  ")
	}
}