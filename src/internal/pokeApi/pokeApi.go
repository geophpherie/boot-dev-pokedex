package pokeApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jbeyer16/boot-dev-pokedex/src/internal/pokeCache"
)

const baseUrl = "https://pokeapi.co/api/v2/"

func GetLocationAreas(url string, cache *pokeCache.Cache) (LocationAreaResponse, error) {
	if url == "" {
		url = baseUrl + "/location-area/"
	}

	cachedResponseBody, cacheHit := cache.Get(url)
	if cacheHit {
		fmt.Println("Cache Hit!")
		results := LocationAreaResponse{}
		err := json.Unmarshal(cachedResponseBody, &results)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error unmarshaling cache")
		}
		return results, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, errors.New("HTTP error")
	}
	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, errors.New("error reaching api")
	}
	decoder := json.NewDecoder(resp.Body)
	results := LocationAreaResponse{}
	err = decoder.Decode(&results)
	if err != nil {
		return LocationAreaResponse{}, errors.New("error decoding api response")
	}

	cachedbody, err := json.Marshal(results)
	if err != nil {
		return results, errors.New("cannot cache results")
	}
	// update cache with response body
	cache.Add(url, cachedbody)
	return results, nil

}

func GetLocationAreaDetail(areaName string, cache *pokeCache.Cache) (LocationAreaDetailedResponse, error) {
	url := baseUrl + "/location-area/" + areaName

	cachedResponseBody, cacheHit := cache.Get(url)
	if cacheHit {
		fmt.Println("Cache Hit!")
		results := LocationAreaDetailedResponse{}
		err := json.Unmarshal(cachedResponseBody, &results)
		if err != nil {
			return LocationAreaDetailedResponse{}, errors.New("error unmarshaling cache")
		}
		return results, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaDetailedResponse{}, errors.New("HTTP error")
	}
	if resp.StatusCode != http.StatusOK {
		return LocationAreaDetailedResponse{}, errors.New("error reaching api")
	}
	decoder := json.NewDecoder(resp.Body)
	results := LocationAreaDetailedResponse{}
	err = decoder.Decode(&results)
	if err != nil {
		return LocationAreaDetailedResponse{}, errors.New("error decoding api response")
	}

	cachedbody, err := json.Marshal(results)
	if err != nil {
		return results, errors.New("cannot cache results")
	}
	// update cache with response body
	cache.Add(url, cachedbody)
	return results, nil
}

func GetPokemonDetail(pokemonName string, cache *pokeCache.Cache) (PokemonResponse, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	cachedResponseBody, cacheHit := cache.Get(url)
	if cacheHit {
		fmt.Println("Cache Hit!")
		results := PokemonResponse{}
		err := json.Unmarshal(cachedResponseBody, &results)
		if err != nil {
			return PokemonResponse{}, errors.New("error unmarshaling cache")
		}
		return results, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return PokemonResponse{}, errors.New("HTTP error")
	}
	if resp.StatusCode != http.StatusOK {
		return PokemonResponse{}, errors.New("error reaching api")
	}
	decoder := json.NewDecoder(resp.Body)
	results := PokemonResponse{}
	err = decoder.Decode(&results)
	if err != nil {
		return PokemonResponse{}, errors.New("error decoding api response")
	}

	cachedbody, err := json.Marshal(results)
	if err != nil {
		return results, errors.New("cannot cache results")
	}
	// update cache with response body
	cache.Add(url, cachedbody)
	return results, nil
}
