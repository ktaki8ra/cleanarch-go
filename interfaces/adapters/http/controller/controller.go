package controller

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/ktaki8ra/cleanarch-go/interfaces/config"

    "errors"
    "log/slog"
    "net/http"
    "strconv"
)

type Controller struct {
    DIModules config.DIModules
}
func New(diModules config.DIModules) *Controller {
    return &Controller{
        DIModules: diModules,
    }
}

func (c *Controller) Run(httpConfig config.HttpConfig) error {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/user/create", UserCreateController(c.DIModules))

    if err := e.Start(":" + strconv.Itoa(httpConfig.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
        slog.Error("failed to start server", "error", err)
        return err
    }
    return nil
}
