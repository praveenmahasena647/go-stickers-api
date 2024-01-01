package handlers

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/praveenmahasena647/go-stickers/cmd/app/dbs"
)

func GetAll(ctx echo.Context) error {
	var stickers, err = dbs.FetchAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"err": "Internal Serer Error",
		})
	}
	return ctx.JSON(http.StatusOK, stickers)
}

func GetbyID(ctx echo.Context) error {
	var sticker, err = dbs.FetchByID(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"err": "couldn't fetch to the ID given",
		})
	}
	return ctx.JSON(http.StatusOK, sticker)
}

func PostOne(ctx echo.Context) error {
	var b, bErr = io.ReadAll(ctx.Request().Body)
	if bErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"err": "Error during decoding Request Body",
		})
	}
	var stickers, err = dbs.InsertOne(b)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, map[string]string{
			"err": "error during inserting data in",
		})
	}
	return ctx.JSON(http.StatusCreated, stickers)
}

func EditbyID(ctx echo.Context) error {
	var b, bErr = io.ReadAll(ctx.Request().Body)
	if bErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"err": "couldnt decode the request body payload",
		})
	}
	var sticker, err = dbs.Update(ctx.Param("id"), b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{
			"err": "could not do the Alter event",
		})
	}
	return ctx.JSON(http.StatusAccepted, sticker)
}

func DeletebyID(ctx echo.Context) error {
	var stic, err = dbs.DeleteOne(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, map[string]string{
			"err": "couldnt do delete operation",
		})
	}
	return ctx.JSON(http.StatusOK, stic)
}
