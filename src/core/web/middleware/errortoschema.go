package core

import (
	. "../../helpers"
	"encoding/json"
	"encoding/xml"
)

type ErrorResponse struct {
	XMLName xml.Name `xml:"Error" json:"-"`
	Message string   `xml:"Message" json:"error"`
}

func ErrorToSchema(err ErrorResponse, output string) []byte {
	if output == "application/json" {
		b, err := json.Marshal(err)
		HandleError(err)
		return b
	}

	b, error := xml.MarshalIndent(err, "  ", "    ")
	HandleError(error)

	return []byte(xml.Header + string(b))
}
