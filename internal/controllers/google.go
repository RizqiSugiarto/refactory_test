package controllers

import (
	"github.com/Siddheshk02/go-oauth2/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service usecase.UserService
}

func New(UserService usecase.UserService, h *fiber.App) {
	r := &UserController{Service: UserService}

	h.Get("/login", r.GoogleLogin)
}

func (u *UserController) GoogleLogin(c *fiber.Ctx) error {
	code := c.Query("state")
	link, err := u.Service.GoogleLogin(code)

	if err != nil {
		return c.SendString(err.Error())
	}

	c.Status(fiber.StatusSeeOther)
	c.Redirect(link)
	return c.SendString("JADI GAESS")
}
