package i18nAPI

import (
	"fmt"
	"slices"

	"github.com/JosunHK/josun-go.git/cmd/util"
	"github.com/JosunHK/josun-go.git/cmd/util/i18n"
	"github.com/JosunHK/josun-go.git/web/templates/contents"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var decoder = schema.NewDecoder()

func GetItems(c echo.Context) error {
	locale := c.Param("locale")
	items := i18n.GetItems(locale)
	slices.Reverse(items)
	return util.HTML(c, contents.I18nTableItems((items)))
}

func AddItems(c echo.Context) error {
	locale := c.Param("locale")
	err := c.Request().ParseForm()
	if err != nil {
		log.Error(err)
		return err
	}

	var item i18n.Item

	err = decoder.Decode(&item, c.Request().PostForm)
	if err != nil {
		log.Error(err)
		return err
	}

	if !isItemValid(item, i18n.GetItems(locale)) {
		err := fmt.Errorf("Invalid item")
		log.Error(err)
		return err
	}

	if err := i18n.AddItem(locale, item); err != nil {
		log.Error(err)
		return err
	}

	return GetItems(c)
}

// I'll add this when I want
func DeleteItems(c echo.Context) error {
	locale := c.Param("locale")
	err := c.Request().ParseForm()
	if err != nil {
		log.Error(err)
		return err
	}

	var item i18n.Item

	err = decoder.Decode(&item, c.Request().PostForm)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := i18n.DeleteItem(locale, item.Key); err != nil {
		log.Error(err)
		return err
	}

	return GetItems(c)

}

func isItemValid(item i18n.Item, items []i18n.Item) bool {
	if item.Key == "" || item.Default == "" {
		return false
	}

	for _, item := range items {
		if item.One != "" && item.Many == "" {
			return false
		}
	}

	return true
}
