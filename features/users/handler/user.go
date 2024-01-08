package handler

import (
	"gomongo/features/users"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type UserHandler struct {
	s users.Service
}

func New(s users.Service) users.Handler {
	return &UserHandler{
		s: s,
	}
}

func (ad *UserHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputData = new(UserReq)
		if err := c.Bind(&inputData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}
		var inputProcess = new(users.Contoh)
		inputProcess.Name = inputData.Name
		inputProcess.Age = inputData.Age
		inputProcess.Email = inputData.Email

		result, err := ad.s.AddUser(*inputProcess)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"massage": "failed to add user",
			})
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"massage": result,
		})
	}
}
