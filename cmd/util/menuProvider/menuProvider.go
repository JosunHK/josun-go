package menuProvider

import (
	"context"

	"github.com/JosunHK/josun-go.git/cmd/database"
	i18nUtil "github.com/JosunHK/josun-go.git/cmd/util/i18n"
	sqlc "github.com/JosunHK/josun-go.git/db/generated"
	log "github.com/sirupsen/logrus"
)

func TranslMenu(ctx context.Context, rawMenu []sqlc.MenuItem) []sqlc.MenuItem {
	menu := []sqlc.MenuItem{}
	for _, item := range rawMenu {
		item.Label = i18nUtil.T(ctx, item.Label)
		menu = append(menu, item)
	}

	return menu

}

func GetMenu(ctx context.Context, key string) []sqlc.MenuItem {
	rawMenu := GetRawMenu(ctx, key)

	log.Debug("menu size: ", len(rawMenu))
	return TranslMenu(ctx, rawMenu)
}

func GetRawMenu(ctx context.Context, key string) []sqlc.MenuItem {
	DB := database.DB
	queries := sqlc.New(DB)

	result, err := queries.GetMenu(ctx, key)
	if err != nil {
		log.Error("Error getting menu: ", err)
		return []sqlc.MenuItem{}
	}

	return result
}
