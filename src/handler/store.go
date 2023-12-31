package handler

import (
	"coffee-paws/src/models"
	"coffee-paws/src/services"
	"coffee-paws/src/utils"
	"log"
	"net/http"

	openapi "github.com/coffee-paws/api"

	"github.com/labstack/echo/v4"
)

// e.POST("/stores/", PostStore)
func PostStore(c echo.Context) error {
	s := new(openapi.CreateStore)
	if err := c.Bind(s); err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusBadRequest, "bad request")
	}
	id := utils.GenerateULID()
	store := models.Store{
		ID: id,
		StoreType: s.StoreType,
		Name: s.Name,
		Address: s.Address,
		Url: s.Url,
	}

	res := services.PostStore(store)
	return c.JSON(http.StatusOK, res)
}

// e.Get("/stores/", GetStores)
func GetStores(c echo.Context) error {
	users := services.GetStores()

	return c.JSON(http.StatusOK, users)
}

// e.Get("/stores/:id", GetStores)
func GetStore(c echo.Context) error {
	id := c.Param("id")
	store := services.GetStore(id)

	return c.JSON(http.StatusOK, store)
}