package web

type MovieVo struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Directors   []string `json:"directors"`
	Producers   []string `json:"producers"`
	Actors      []string `json:"actors"`
	Genre       []string `json:"genre"`
	Rating      int64    `json:"rating"`
	ReleaseDate string   `json:"release_date"`
}
