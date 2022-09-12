package model

type Item struct {
	ID          uint   `json:"id" faker:"-"`
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name" faker:"name"`
	Sale        int    `json:"sale" faker:"boundary_start=0, boundary_end=100"`
	Size        string `json:"size" faker:"len=5"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand" faker:"word"`
	Status      int    `json:"status" faker:"boundary_start=100, boundary_end=550"`
}
