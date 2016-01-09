package gopokemon

import (
	"net/http"
)

type PokemonData struct {
	Name       string `json:"name"`
	NationalID int    `json:"national_id"`
}
