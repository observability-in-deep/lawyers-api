package lawyers

import "github.com/gofiber/fiber/v2"

func get(c *fiber.Ctx) error {
	lawyerOab := c.Params("lawyerOab")
	lawyer, err := Get(c.UserContext(), lawyerOab)
	if err != nil {
		if err == fiber.ErrNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())

		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(lawyer)
}

func Register(app *fiber.App) {
	app.Get("/lawyers/:LawyerOab", get)

}
