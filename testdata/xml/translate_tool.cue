package translate

import (
	"encoding/xml"
	"encoding/json"
	"tool/file"
)

command: TranslateXml: {
	Input: {
		xml: "input.xml"
		json: "output.json"
	}
	XmlBytes:  {
		file.Read
		filename: Input.xml
		contents: bytes
	}
	Unmarshalled: xml.Unmarshal(XmlBytes.contents)
	Output: file.Create & {
		filename: Input.json
		contents: json.Indent(json.Marshal(Unmarshalled), "", "  ")
	}
}