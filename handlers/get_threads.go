package handlers

import (
	"encoding/json"
	"github.com/karur4n/kujibiki-server-go/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (si *ServerInterfaceImpl) GetThreads(ctx echo.Context) error {
	threads, err := models.FindAllThreads()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(threads)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(jsonData))
}