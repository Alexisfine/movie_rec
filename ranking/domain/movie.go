package domain

import "time"

type Movie struct {
	Id          string
	Name        string
	Director    []string
	Producers   []string
	Actors      []string
	Genre       []string
	Rating      int64
	ReleaseDate time.Time
}
