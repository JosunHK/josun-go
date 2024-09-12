package templUtil

import (
	"github.com/a-h/templ"
	log "github.com/sirupsen/logrus"
)

func ToJSONString(object any) string {
	JSONStr, err := templ.JSONString(object)
	if err != nil {
		JSONStr = "{}"
	}

	log.Info("JSONStr: ", JSONStr)

	return JSONStr
}
