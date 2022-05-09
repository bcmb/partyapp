package party

// swagger:model Party
type Party struct {
	Id       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"New year party"`
	City     string `json:"city" example:"Kyiv"`
	Address  string `json:"address" example:"Maydan nezalezhnosti 1"`
	DateTime string `json:"dateTime" example:"2022-12-31T23:00:00Z"`
}
