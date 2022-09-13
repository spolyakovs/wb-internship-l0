package model

type Delivery struct {
	ID      uint   `json:"id" faker:"-"`
	Name    string `json:"name" faker:"name"`
	Phone   string `json:"phone" faker:"phone_number"`
	Zip     string `json:"zip" faker:"len=16"`
	City    string `json:"city" faker:"oneof: Paris, London, New York, Moscow"`
	Address string `json:"address" faker:"oneof: Address test, Address test 2"`
	Region  string `json:"region" faker:"oneof: Region 1, Region 2"`
	Email   string `json:"email" faker:"email"`
}
