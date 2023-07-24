package controllers

import "github.com/gofiber/fiber/v2"

// ---------------------------------------------------------------------
//
//	THIS WILL SHOW AT THE HOMEPAGE - START
//
// ---------------------------------------------------------------------
func Home(c *fiber.Ctx) error {
	return c.SendString("Hello! Everyone, \nWelcome to the Andy Brand Project.")
}

// ---------------------------------------------------------------------
//
//	THIS WILL SHOW AT THE HOMEPAGE - END
//
// ---------------------------------------------------------------------
