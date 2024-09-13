package templUtil

import (
	"github.com/a-h/templ"
)

func ToJSONString(object any) string {
	JSONStr, err := templ.JSONString(object)
	if err != nil {
		JSONStr = "{}"
	}

	return JSONStr
}
