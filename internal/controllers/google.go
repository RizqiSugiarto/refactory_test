package controllers

import (
	"fmt"

	"github.com/Siddheshk02/go-oauth2/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service usecase.UserService
}

func New(UserService usecase.UserService, h *fiber.App) {
	r := &UserController{Service: UserService}

	h.Get("/login", r.GoogleLogin)
	h.Get("/google_callback", r.GoogleFeedBack)
}

func (u *UserController) GoogleLogin(c *fiber.Ctx) error {
	// code := c.Query("state")
	link := u.Service.GoogleLogins()

	c.Status(fiber.StatusSeeOther)
	c.Redirect(link)
	return c.SendString("JADI GAESS")
}

func (u *UserController) GoogleFeedBack(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "MISALKAN" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	resp, err := u.Service.GoogleLogin(code)

	if err != nil {
		fmt.Println(err)
		return c.SendString("code dosn't catch correctly")
	}

	return c.SendString(resp)

}
