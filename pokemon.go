package gopokemon

import (
	"encoding/json"
	"net/http"
)

type PokemonData struct {
	Name       string `json:"name"`
	PokedexID  int    `json:"pkdx_id"`
	NationalID int    `json:"national_id"`
	Species    string `json:"species"`
	Weight     string `json:"weight"`
	Types      []struct {
		Name string `json:"name"`
	} `json:"types"`
	Evolutions []struct {
		Level       int    `json:"level"`
		Method      string `json:"method"`
		To          string `json:"to"`
		ResourceURI string `json:"resource_uri"`
	}
	Descriptions []struct {
		Name        string `json:"name"`
		ResourceURI string `json:"resource_uri"`
	}
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
