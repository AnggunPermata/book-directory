package route

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "status ok")
	})
	e.GET("/healthcheck", func (c echo.Context) error {
		return c.JSON(200, map[string]bool{
			"ok":true,
		})
	})
	
}
