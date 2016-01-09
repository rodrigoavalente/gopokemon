package gopokemon

import (
	"encoding/json"
	"net/http"
)

type PokemonData struct {
	Name       string `json:"name"`
	PokedexID  int    `json:"pkdx_id"`
	NationalID int    `json:"national_id"`
	Weight     string `json:"weight"`
	Types      []struct {
		Name string `json:"name"`
	} `json:"types"`
}

func QueryPokemon(pokemon_id string) (PokemonData, error) {
	response, err := http.Get("http://pokeapi.co/api/v1/pokemon/" + pokemon_id + "/")

	if err != nil {
		return PokemonData{}, nil
	}

	defer response.Body.Close()
	var pokemon PokemonData
	if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
		return PokemonData{}, err
	}

	return pokemon, nil
}
