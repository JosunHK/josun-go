package dummy

import (
	"strconv"

	"math/rand"

	"github.com/JosunHK/josun-go.git/cmd/util"
	dummyTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/dummy"
	"github.com/labstack/echo/v4"
)

func Odometer(c echo.Context) error {
	num := strconv.Itoa(rand.Intn(100000))
	return util.HTML(c, dummyTemplates.Update(num))
}

// func Odometer(c echo.Context) error {
// 	num := strconv.Itoa(rand.Intn(100000))
// 	return c.String(http.StatusOK, num)
// }
