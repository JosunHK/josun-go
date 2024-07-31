package cfg

import (
	"github.com/JosunHK/josun-go.git/cmd/i18n"
	"github.com/eduardolat/goeasyi18n"
)

type Cfg struct {
	I18n *goeasyi18n.I18n
}

func CfgInit() (Cfg, error) {
	i18n, err := i18n.InitI18n()
	if err != nil {
		return Cfg{}, err
	}

	cfg := Cfg{
		I18n: i18n,
	}

	return cfg, nil
}
