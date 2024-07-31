package i18n

import (
	"fmt"
	"io"
	"os"

	"github.com/eduardolat/goeasyi18n"
)

type Transl func(s string) string

var T_LIST = []string{"en", "zh"}

func InitI18n() (*goeasyi18n.I18n, error) {
	i18n := goeasyi18n.NewI18n(goeasyi18n.Config{
		FallbackLanguageName:    "en",
		DisableConsistencyCheck: false,
	})

	for _, t := range T_LIST {
		translJSON, err := readJSON(t)
		if err != nil {
			return nil, fmt.Errorf("error loading %v translations %v", t, err)
		}

		transl, err := goeasyi18n.LoadFromJsonString(translJSON)
		if err != nil {
			return nil, fmt.Errorf("error loading %v translations %v", t, err)
		}

		i18n.AddLanguage(t, transl)
	}

	return i18n, nil
}

func readJSON(locale string) (string, error) {
	path := fmt.Sprintf("./web/static/i18n/%s.json", locale)

	jsonFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return "", nil
	}

	return string(byteValue), nil
}
