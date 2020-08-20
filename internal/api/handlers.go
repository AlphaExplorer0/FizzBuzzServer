package api

import (
	"errors"

	"github.com/gorilla/mux"
	"leboncointest.com/fizzbuzzserver/internal/config"
)

// API object
type API struct {
	Conf *config.Config
	Hits map[fizzbuzzParams]int
}

// Register :
func (api *API) Register(router *mux.Router, config *config.Config) error {
	if router == nil {
		return errors.New("server not initialized")
	}

	api.AddFizzBuzzHandlers(router)

	api.Conf = config
	api.Hits = make(map[fizzbuzzParams]int)

	return nil
}
