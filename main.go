package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success": true,
		"albums":  albums,
	})
}

func getAlbum(c *fiber.Ctx) error {

	strID := c.Params("id")
	id, error := strconv.Atoi(strID)

	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Infalid ID",
		})
	}

	for _, album := range albums {
		if album.ID == id {
			return c.JSON(&fiber.Map{
				"success": true,
				"album":   album,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
		"success": false,
		"message": "Album not found",
	})

}

func postAlbum(c *fiber.Ctx) error {
	var newAlbum Album

	if err := c.BodyParser(&newAlbum); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request Body",
		})
	}

	newAlbum.ID = len(albums) + 1

	albums = append(albums, newAlbum)

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"album":   newAlbum,
	})
}

func main() {
	app := fiber.New()

	app.Get("/albums", getAlbums)
	app.Get("/albums/:id", getAlbum)
	app.Post("/albums", postAlbum)

	log.Fatal(app.Listen(":8080"))
}
