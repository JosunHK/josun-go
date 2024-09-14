package i18nUtil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	i18nStruct "github.com/JosunHK/josun-go.git/cmd/struct/i18n"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	"github.com/eduardolat/goeasyi18n"
	log "github.com/sirupsen/logrus"
)

type Transl func(s string) string

type NumTransl func(num int) string

// just up to 20 no other use cases anyways
var ZH_NUM_MAP = map[int]string{0: "零", 1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六", 7: "七", 8: "八", 9: "九", 10: "十", 11: "十一", 12: "十二", 13: "十三", 14: "十四", 15: "十五", 16: "十六", 17: "十七", 18: "十八", 19: "十九", 20: "二十"}

var NUM_TRANSL_MAP = map[string]NumTransl{
	"en": defaultNumTransl,
	"zh": ChineseNumTransl,
}

const LOCALE_SETTING_ID = "locale"

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

func GetMenuItem(locale string) sqlc.MenuItem {
	for _, item := range LOCALE_MENU {
		if item.Value == locale {
			return item
		}
	}
	return sqlc.MenuItem{}
}

func T(ctx context.Context, key string) string {
	locale := getLocaleFromContext(ctx)
	res := I18n.T(locale, key)
	if res == "" {
		log.Error("Translation not found for key: ", key)
		return key
	}

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

func getLocaleFromContext(c context.Context) string {
	localeAny := c.Value(LOCALE_SETTING_ID)
	locale, ok := localeAny.(string)
	if !ok {
		return "en"
	}

	return locale
}

func GetItems(locale string) []i18nStruct.Item {
	var items []i18nStruct.Item
	table, err := readJSON(locale)
	if err != nil {
		log.Error("Error reading json file: ", err)
		return []i18nStruct.Item{}
	}

	json.Unmarshal([]byte(table), &items)

	return items
}

func AddItem(locale string, item i18nStruct.Item) error {
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

func TN(ctx context.Context, num int) string {
	locale := getLocaleFromContext(ctx)
	fn := NUM_TRANSL_MAP[locale]
	if fn == nil {
		return defaultNumTransl(num)
	}
	return fn(num)
}

func ChineseNumTransl(num int) string {
	numCharSet := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	unitCharSet := []string{"", "十", "百", "千", "萬", "億", "兆"}
	n := []int{}
	for ; num > 0; num /= 10 {
		n = append(n, num%10)
		log.Debug("n: ", n)
	}

	res := ""
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			if i != 0 && len(n) > i+1 && n[i+1] != 0 {
				res = numCharSet[0] + res
			}
		} else {
			if i == 1 {
				res = unitCharSet[i] + res
			} else {
				res = numCharSet[n[i]] + unitCharSet[i] + res
			}
		}
	}

	return res
}

func defaultNumTransl(num int) string {
	return fmt.Sprintf("%d", num)
}
