package i18n

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/eduardolat/goeasyi18n"
	log "github.com/sirupsen/logrus"
)

type Transl func(s string) string

var I18n *goeasyi18n.I18n

var T_LIST = []string{"en", "zh"}

var LOCALE_MENU = []sqlc.MenuItem{
	{
		Label: "English",
		Value: "en",
	},
	{
		Label: "中文",
		Value: "zh",
	},
}

type Item struct {
	Key     string `json:"Key"`
	Default string `json:"Default"`
	One     string `json:"One"`
	Many    string `json:"Many"`
}

func GetMenuItem(locale string) sqlc.MenuItem {
	for _, item := range LOCALE_MENU {
		if item.Value == locale {
			return item
		}
	}
	return sqlc.MenuItem{}
}

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
	return "en"
}

func GetItems(locale string) []Item {
	var items []Item
	table, err := readJSON(locale)
	if err != nil {
		log.Error("Error reading json file: ", err)
		return []Item{}
	}

	json.Unmarshal([]byte(table), &items)

	return items
}

func AddItem(locale string, item Item) error {
	items := GetItems(locale)
	for _, i := range items {
		if i.Key == item.Key {
			return fmt.Errorf("Key already exists")
		}
	}

	items = append(items, item)

	bytes, err := json.Marshal(items)
	if err != nil {
		log.Error("Error marshalling json: ", err)
		return err
	}

	path := fmt.Sprintf("./web/static/i18n/%s.json", locale)
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		log.Error("Error writing file: ", err)
		return err
	}

	return nil
}

func DeleteItem(locale, key string) error {
	items := GetItems(locale)
	var toRemove int
	for i, item := range items {
		if item.Key == key {
			toRemove = i
		}
	}

	items = append(items[:toRemove], items[toRemove+1:]...)

	bytes, err := json.Marshal(items)
	if err != nil {
		log.Error("Error marshalling json: ", err)
		return err
	}

	path := fmt.Sprintf("./web/static/i18n/%s.json", locale)
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		log.Error("Error writing file: ", err)
		return err
	}

	return nil
}
