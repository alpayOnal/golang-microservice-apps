package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"micro_apps/micro_app/models"
	"micro_apps/micro_app/repository/kafka"
	"micro_apps/micro_app/repository/mongodb"
	"micro_apps/micro_app/utils/validation"
)

// ItemHandler  represent the httphandler for item
type ItemHandler struct {
}

func NewItemHandler(e *echo.Echo) {
	handler := &ItemHandler{}

	e.GET("/items/:id", handler.GetItem)
	e.GET("/items", handler.GetItems)

	e.POST("/items", handler.AddItem)
}

func (i *ItemHandler) AddItem(c echo.Context) error {

	item := models.NewItem()
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		log.Error("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = validation.Check(item)
	if err != nil {
		errors := strings.Split(err.Error(), ";")
		response := models.ErrorResponse{Messages: errors}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	kafka.NewItemRepository().Store(item)
	log.Printf("this is your item %#v", item)
	return c.String(http.StatusOK, "We got your Item!!!")
}

func (i *ItemHandler) GetItem(c echo.Context) error {
	id := c.Param("id")
	item, err := mongodb.GetItemRepository().ItemById(id)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, item)
}

func (i *ItemHandler) GetItems(c echo.Context) error {
	itemList, err := mongodb.GetItemRepository().Items()
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	//resp := json.Marshal(itemList)
	return c.JSON(http.StatusOK, itemList)
}
