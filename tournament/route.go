package tournament

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App, prefix string) {
	api := app.Group(prefix)

	api.Route("", func(router fiber.Router) {
		router.Post("/", Create)
		router.Get("", FindTournaments)
	})
	api.Route("/:d", func(router fiber.Router) {
		router.Delete("", DeleteTournament)
		router.Get("", FindTournamentById)
		router.Put("", UpdateTournament)
	})
}
