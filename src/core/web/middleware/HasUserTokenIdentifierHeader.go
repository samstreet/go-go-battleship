package middleware

import (
	"../../helpers"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type ErrorResponse struct {
	XMLName xml.Name `xml:"Error" json:"-"`
	Message string   `xml:"Message" json:"error"`
}

func HasUserTokenIdentifierHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		identifier := r.Header.Get("X-USER-IDENTIFIER")
		if identifier == "" {
			body := ErrorToSchema(ErrorResponse{Message: "Missing X-USER-IDENTIFIER Header"}, r.Header.Get("Accept"))
			w.Header().Set("Content-Type", r.Header.Get("Accept"))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func ErrorToSchema(err ErrorResponse, output string) []byte {
	if output == "application/json" {
		b, err := json.Marshal(err)
		helpers.HandleError(err)
		return b
	}

	b, error := xml.MarshalIndent(err, "  ", "    ")
	helpers.HandleError(error)

	return []byte(xml.Header + string(b))
}
