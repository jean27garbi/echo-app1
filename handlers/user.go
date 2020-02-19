package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/jean27garbi/echo-app1/models"
	"github.com/labstack/echo/v4"
)

// ListUsers - return all users
func ListUsers(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.JSON(http.StatusOK, "team:"+team+", member:"+member)
}

// SaveUser - create a new user
func SaveUser(c echo.Context) error {

	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

// SaveUserImage - create a new user  avatar
func SaveUserImage(c echo.Context) error {

	newUser := models.User{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
	}
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+newUser.Username+"</b>")
}

// GetUser - return a specific user by id
func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// UpdateUser - update a specific user by id
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// DeleteUser - delete a specific user by id
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
