package model

type Item struct {
	Id          uint   `json:"id"`
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RId         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
