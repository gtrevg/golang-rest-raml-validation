package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gocraft/web"
	"github.com/xeipuuv/gojsonschema"
)

type Context struct {
	schemaDir string
}

var schemaDir string

func main() {

	schemaDir = getSchemaDir()

	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).     // Use some included middleware
		Middleware(web.ShowErrorsMiddleware). // ...
		Middleware(web.StaticMiddleware("www")).
		Post("/api/v0/keys", (*Context).createKeyValue)

	http.ListenAndServe("localhost:3000", router)
}

// ===============================================
// HELPER FUNCTIONS
// ===============================================

// getSchemaDir gets a full file path to the schemas directory
func getSchemaDir() string {
	fpdir := filepath.Dir(os.Args[0])
	dir, err := filepath.Abs(fpdir)
	if err != nil {
		println("Couldn't get absolute path of $PWD: " + fpdir)
		println("Exiting service.")
		panic(err)
	}
	return filepath.Join(dir, "www/schemas")
}

// validateRequestData takes in a schema path and the request
// and will do the legwork of determining if the post data is valid
func validateRequestData(schemaPath string, r *web.Request) (
	document map[string]interface{},
	result *gojsonschema.Result,
	err error,
) {
	err = json.NewDecoder(r.Body).Decode(&document)
	if err == nil && schemaPath != "" {
		schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)
		documentLoader := gojsonschema.NewGoLoader(document)

		result, err = gojsonschema.Validate(schemaLoader, documentLoader)
	}

	return document, result, err
}

// ===============================================
// HANDLERS
// ===============================================

// creatKeyValue creates a new book in the collection
func (c *Context) createKeyValue(w web.ResponseWriter, r *web.Request) {

	document, result, err := validateRequestData(
		"file://"+schemaDir+"/keyvalue.post.body.json",
		r,
	)

	// lazy output & not setting headers

	if err != nil || !result.Valid() {
		w.Header()
		fmt.Fprintf(w, "The document is not valid. see errors :\n")

		if result != nil {
			for _, desc := range result.Errors() {
				fmt.Fprintf(w, "- %s\n", desc)
			}
		} else {
			fmt.Fprint(w, "The document wasn't valid JSON.")
		}
		return

	}

	fmt.Fprint(w, "success", "\n")
	fmt.Fprint(w, document, "\n")
	fmt.Fprint(w, result, "\n")

}
