package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Pokemon []struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Type           []string `json:"type"`
	Hp             int      `json:"hp"`
	Attack         int      `json:"attack"`
	Defense        int      `json:"defense"`
	SpecialAttack  int      `json:"special_attack"`
	SpecialDefense int      `json:"special_defense"`
	Speed          int      `json:"speed"`
}

const API_URL string = "https://jsonplaceholder.typicode.com/users"

func get_pokemon() Pokemon {
	resp, err := http.Get(API_URL)
	if err != nil {
		log.Fatal("Cannot fetch API")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Cannot read body")
	}

	var result Pokemon
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("Cannot unmarshal JSON")
	}
	return result
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/assets", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		pokemons := get_pokemon()
		return c.Render("index", fiber.Map{
			"title": pokemons,
		})
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
