package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type fizzbuzzParams struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

// AddFizzBuzzHandlers add fizzbuzz endpoints
func (api *API) AddFizzBuzzHandlers(router *mux.Router) {
	router.HandleFunc("/fizzbuzz/v1/produce", api.produce).Methods("GET")
	router.HandleFunc("/fizzbuzz/v1/stats", api.stats).Methods("GET")
}

func (api *API) produce(w http.ResponseWriter, r *http.Request) {

	par := r.URL.Query()

	// Check params
	params, err := validateQueryParams(par)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	var b strings.Builder
	var ctr, inc int
	if params.Limit < 0 {
		ctr, inc = -1, -1
	} else {
		ctr, inc = 1, 1
	}

	for abs(ctr) <= abs(params.Limit) {
		if ctr%(params.Int1*params.Int2) == 0 {
			fmt.Fprintf(&b, "%s ", params.Str1+params.Str2)
		} else if ctr%params.Int1 == 0 {
			fmt.Fprintf(&b, "%s ", params.Str1)
		} else if ctr%params.Int2 == 0 {
			fmt.Fprintf(&b, "%s ", params.Str2)
		} else {
			fmt.Fprintf(&b, "%d ", ctr)
		}
		ctr += inc
	}

	s := b.String()
	if len(s) > 0 {
		s = s[:len(s)-1] // removes trailing " "
	}

	if _, ok := api.Hits[*params]; ok {
		api.Hits[*params]++
	} else {
		api.Hits[*params] = 1
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func (api *API) stats(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if len(api.Hits) == 0 {
		io.WriteString(w, "There has been no fizzbuzz request so far")
		return
	}

	var pMax fizzbuzzParams
	maxHits := 0
	for key, value := range api.Hits {
		if value > maxHits {
			maxHits = value
			pMax = key
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Most frequent fizzbuzz request - ")
	fmt.Fprintf(&b, "Int1: %d, Int2: %d, Limit: %d, Str1: %s, Str2: %s ", pMax.Int1, pMax.Int2, pMax.Limit, pMax.Str1, pMax.Str2)
	fmt.Fprintf(&b, "Number of hits: %d", maxHits)
	w.Write([]byte(b.String()))
}

func validateQueryParams(values url.Values) (*fizzbuzzParams, error) {

	ret := &fizzbuzzParams{}
	if param, err := getIntParam(values.Get("int1"), "int1"); err == nil {
		ret.Int1 = param
	} else {
		return nil, err
	}

	if param, err := getIntParam(values.Get("int2"), "int2"); err == nil {
		ret.Int2 = param
	} else {
		return nil, err
	}

	if param, err := getIntParam(values.Get("limit"), "limit"); err == nil {
		ret.Limit = param
	} else {
		return nil, err
	}

	if values.Get("str1") == "" {
		return nil, errors.New("empty str1")
	}
	ret.Str1 = values.Get("str1")

	if values.Get("str2") == "" {
		return nil, errors.New("empty str2")
	}
	ret.Str2 = values.Get("str2")

	return ret, nil
}

func getIntParam(val, paramName string) (int, error) {

	if val == "" {
		return -1, errors.New("empty " + paramName)
	}
	if s, err := strconv.Atoi(val); err == nil {
		if s == 0 {
			return -1, errors.New(paramName + " should not be 0")
		}
		return s, nil
	} else {
		return -1, err
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
