package entity

type RawMaterial struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	Protein       float64 `json:"protein"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
	Calories      float64 `json:"calories"`
}
