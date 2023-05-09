package main

import (
	"fmt"

	"github.com/anggunpermata/book-directory/config"
	"github.com/anggunpermata/book-directory/route"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.LoadConfig()
	route.InitRoutes(e)
	if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		e.Logger.Fatalf("failed to start book directory server with error:%v", err)
	}
}
