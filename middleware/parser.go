package customMiddleware

import (
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Body() {
	println("foo")
}

func Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		next(c)
		return nil
	}
}
