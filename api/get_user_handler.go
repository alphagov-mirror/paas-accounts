package api

import (
	"net/http"

	"github.com/alphagov/paas-accounts/database"
	"github.com/labstack/echo"
)

func GetUserHandler(db *database.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := db.GetUser(c.Param("uuid"))
		if err != nil {
			if err == database.ErrUserNotFound {
				return NotFoundError{"user not found"}
			}
			return InternalServerError{err}
		}

		return c.JSON(http.StatusOK, user)
	}
}
