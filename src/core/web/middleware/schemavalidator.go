package core

import (
	. "../../helpers"
	"encoding/json"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"net/http"
)

func SchemaValidator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		schemaLoader := gojsonschema.NewReferenceLoader("file://src/schema.json")

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		HandleError(err)

		var body interface{}
		err = json.Unmarshal(b, &body)
		HandleError(err)

		loader := gojsonschema.NewGoLoader(body)
		result, err := gojsonschema.Validate(schemaLoader, loader)
		HandleError(err)

		if !result.Valid() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}
