package customer

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/observability-in-deep/lawyers-api/src/model"
)

func get(c *fiber.Ctx) error {
	customerCPF := c.Params("customerCPF")
	customer, err := Get(c.UserContext(), customerCPF)
	if err != nil {
		if err == fiber.ErrNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())

		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(customer)
}

func post(c *fiber.Ctx) error {

	body := new(models.Customers)

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	customer, err := Create(c.UserContext(), body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(customer)
}

func update(c *fiber.Ctx) error {

	body := new(models.Customers)
	if error := c.BodyParser(body); error != nil {
		return fiber.NewError(fiber.StatusBadRequest, error.Error())
	}
	customer, err := Update(c.UserContext(), c.Params("customerId"), body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(customer)
}

func delete(c *fiber.Ctx) error {
	err := Delete(c.UserContext(), c.Params("customerId"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func Register(app *fiber.App) {
	app.Get("/customer/:customerCPF", get)
	app.Post("/customer", post)
	app.Put("/customer/:customerId", update)
	app.Delete("/customer/:customerId", delete)
}
