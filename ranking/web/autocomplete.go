package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ranking/domain"
	"ranking/service"
	"time"
)

type AutoCompleteHandler struct {
	svc service.AutoCompleteSvc
}

func NewAutoCompleteHandler(svc service.AutoCompleteSvc) *AutoCompleteHandler {
	return &AutoCompleteHandler{
		svc: svc,
	}
}

func (m *AutoCompleteHandler) AutoComplete(w http.ResponseWriter, r *http.Request) {
	type AutoCompleteReq struct {
		Keyword string `json:"keyword"`
	}
	var req AutoCompleteReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	movies, err := m.svc.CompleteSearchBar(context.Background(), req.Keyword)
	if err != nil {
		fmt.Printf("an error occured in autocomplete {}", err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "server internal error encountered, please try again one or two minutes later or contact admin")
		return
	}

	res := make([]MovieVo, len(movies))
	for i := 0; i < len(movies); i++ {
		res[i] = toMovieVo(movies[i])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func toMovieVo(movie domain.Movie) MovieVo {
	return MovieVo{
		Id:          movie.Id,
		Name:        movie.Name,
		Directors:   movie.Director,
		Producers:   movie.Producers,
		Actors:      movie.Actors,
		Genre:       movie.Genre,
		Rating:      movie.Rating,
		ReleaseDate: movie.ReleaseDate.Format(time.DateOnly),
	}
}
