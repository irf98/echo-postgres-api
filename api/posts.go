package api

import (
	"awesomeProject/api/models"
	"awesomeProject/api/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreatePost(c echo.Context) error {
	user := models.User{}
	post := models.Post{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sess, _ := session.Get("session", c)
	if sess.Values["user"] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Login before you post!")
	}

	err = services.GetUserByEmail(&user, sess.Values["user"].(string))

	err = services.CreatePost(&post, user, data["body"], data["tag"])
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, post)
}

func GetUserPosts(c echo.Context) error {
	var posts []models.Post

	id := c.Param("id")

	err := services.GetUserPosts(&posts, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exists.")
	}

	return c.JSON(http.StatusOK, posts)
}

func GetPostsByTag(c echo.Context) error {
	var posts []models.Post
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = services.GetPostsByTag(&posts, data["tag"])
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func UpVotePost(c echo.Context) error {
	post := models.Post{}

	id := c.Param("id")

	sess, _ := session.Get("session", c)
	if sess.Values["user"] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Login before you vote!")
	}

	err := services.UpVotePost(&post, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Successfully up voted the post.")
}

func DownVotePost(c echo.Context) error {
	post := models.Post{}

	id := c.Param("id")

	sess, _ := session.Get("session", c)
	if sess.Values["user"] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Login before you vote!")
	}

	err := services.DownVotePost(&post, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Successfully down voted the post.")
}

func DeletePost(c echo.Context) error {
	post := models.Post{}
	user := models.User{}

	id := c.Param("id")

	sess, _ := session.Get("session", c)
	if sess.Values["user"] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Login before you delete!")
	}

	err := services.GetUserByEmail(&user, sess.Values["user"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = services.GetPostById(&post, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if post.FKUser != user.ID {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not the author of this post.")
	}

	err = services.DeletePost(&post, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Post successfully deleted.")
}
