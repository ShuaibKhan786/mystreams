package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func getIDfromParms(c *fiber.Ctx) *int {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("Failed to convert and id recv from parm")
		return nil
	}

	return &idInt
}
