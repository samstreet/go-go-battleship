package core

import (
	"encoding/json"
	"encoding/xml"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func TransformDTOToSchema(dto interface{}, output string) []byte {
	if output == "application/json" {
		b, err := json.Marshal(dto)
		HandleError(err)
		return b
	}

	b, err := xml.MarshalIndent(dto, "  ", "    ")
	HandleError(err)

	return []byte(xml.Header + string(b))
}
