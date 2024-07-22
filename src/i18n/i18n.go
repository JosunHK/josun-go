package i18n

import (
	"fmt"
	"io"
	"os"

	"github.com/eduardolat/goeasyi18n"
)

type Transl func(s string) string

func InitI18n() (*goeasyi18n.I18n, error) {
	i18n := goeasyi18n.NewI18n(goeasyi18n.Config{
		FallbackLanguageName:    "en",
		DisableConsistencyCheck: false,
	})

	enJSON, err := readJSON("en")
	if err != nil {
		return nil, fmt.Errorf("error loading en translations %v", err)
	}

	zhJSON, err := readJSON("zh")
	if err != nil {
		return nil, fmt.Errorf("error loading zh translations %v", err)
	}

	enTranslations, err := goeasyi18n.LoadFromJsonString(enJSON)
	if err != nil {
		return nil, fmt.Errorf("error loading en translations %v", err)
	}

	zhTranslations, err := goeasyi18n.LoadFromJsonString(zhJSON)
	if err != nil {
		return nil, fmt.Errorf("error loading zh translations %v", err)
	}

	i18n.AddLanguage("en", enTranslations)
	i18n.AddLanguage("zh", zhTranslations)

	return i18n, nil
}

func readJSON(locale string) (string, error) {
	path := fmt.Sprintf("./static/i18n/%s.json", locale)

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
