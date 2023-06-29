package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

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

const API_URL string = "http://127.0.0.1:4000/pokemon.json"
const MAX_COUNT int = 20

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

func limit(pokemon Pokemon) Pokemon {
	all := len(pokemon)
	cap := cap(pokemon)
	if all == 0 {
		return pokemon
	}
	if all > MAX_COUNT {
		return pokemon[:MAX_COUNT]
	} else if all == 1 && cap > 1 {
		return pokemon[:all+1]
	}
	return pokemon
}

func filter(pokemon Pokemon, filter string) Pokemon {
	var filtered_pokemons Pokemon
	for _, poke := range pokemon {
		if strings.Contains(strings.ToLower(poke.Name), filter) {
			filtered_pokemons = append(filtered_pokemons, poke)
		}
	}
	return filtered_pokemons
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		pokemon := get_pokemon()
		limited := limit(pokemon)
		return c.Render("index", fiber.Map{
			"pokemons": limited,
		})
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		fil := c.Query("name")
		pokemon := get_pokemon()
		filtered := filter(pokemon, fil)
		limited := limit(filtered)

		return c.Render("partial", fiber.Map{
			"pokemons": limited,
		})
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
