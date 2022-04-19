package api

import (
	"awesomeProject/api/models"
	"awesomeProject/api/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = services.GetUserByEmail(&user, data["email"])
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User already exists")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	err = services.CreateUser(&user, data["email"], password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id":         user.ID,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt,
	})
}

func Login(c echo.Context) error {
	user := models.User{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = services.GetUserByEmail(&user, data["email"])
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exists.")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect password.")
	}

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	var users []models.User

	err := services.GetUsers(&users)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	user := models.User{}

	id := c.Param("id")

	err := services.GetUserById(&user, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User does not exists.")
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserByEmail(c echo.Context) error {
	user := models.User{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = services.GetUserByEmail(&user, data["email"])
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User does not exists.")
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserEmail(c echo.Context) error {
	user := models.User{}
	var data = make(map[string]string)

	id := c.Param("id")

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = services.UpdateUserEmail(&user, id, data["email"])
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Email successfully updated.")
}

func UpdateUserPassword(c echo.Context) error {
	user := models.User{}
	var data = make(map[string]string)

	id := c.Param("id")

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	err = services.UpdateUserPassword(&user, id, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Password successfully updated.")
}

func DeleteUser(c echo.Context) error {
	user := models.User{}

	id := c.Param("id")

	err := services.GetUserById(&user, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exists.")
	}

	err = services.DeleteUser(&user, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User successfully deleted.")
}

func DeleteUserPermanently(c echo.Context) error {
	user := models.User{}

	id := c.Param("id")

	err := services.GetDeletedUser(&user, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exists.")
	}

	err = services.DeleteUserPermanently(&user, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User permanently deleted.")
}
