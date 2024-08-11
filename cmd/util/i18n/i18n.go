package i18n

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/eduardolat/goeasyi18n"
	log "github.com/sirupsen/logrus"
)

type Transl func(s string) string

var I18n *goeasyi18n.I18n

var T_LIST = []string{"en", "zh"}

func T(ctx context.Context, key string) string {
	locale := getLocaleFromCookie(ctx)
	res := I18n.T(locale, key)
	if res == "" {
		log.Error("Translation not found for key: ", key)
		return key
	}

	log.Info("Translated into ", res)
	return res
}

func InitI18n() error {
	I18n = goeasyi18n.NewI18n(goeasyi18n.Config{
		FallbackLanguageName:    "en",
		DisableConsistencyCheck: false,
	})

	for _, t := range T_LIST {
		translJSON, err := readJSON(t)
		if err != nil {
			return fmt.Errorf("error loading %v translations %v", t, err)
		}

		transl, err := goeasyi18n.LoadFromJsonString(translJSON)
		if err != nil {
			return fmt.Errorf("error loading %v translations %v", t, err)
		}

		I18n.AddLanguage(t, transl)
	}

	return nil
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

func getLocaleFromCookie(c context.Context) string {
	return "zh"
}
