# cue-transcoding

Support for additional data and schema formats for CUE.

Build into a [customised CUE fork](https://github.com/orpharion/Cue/tree/cue_transcoding) of the CUE library and application.

## formats

- [xml](pkg/encoding/xml/xml.go)
- [excel](pkg/encoding/excel/excel.go)

## Usage

Install the customised cue binary, and then:

See testdata/xml - converts input.xml to output.json
```shell
cd testdata/xml
cue cmd TranslateXml
```
testdata/excel - converts input.xlsx to output.json
```shell
cd testdata/excel
cue cmd TranslateExcel
```