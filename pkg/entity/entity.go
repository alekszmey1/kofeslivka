package entity

import (
	"log"
	"strconv"
	"strings"
)

type RawMaterial struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	Protein       float64 `json:"protein"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
	Calories      float64 `json:"calories"`
}

type IngredientGram struct {
	Name string  `json:"name"`
	Gram float64 `json:"gram"`
}
type Desert struct {
	Name string `json:"name"`
	List string `json:"list"`
}

func (*IngredientGram) MakeIngGram(s string) (*IngredientGram, error) {
	ing := &IngredientGram{}
	ss := strings.Split(s, " ")
	gram, err := strconv.ParseFloat(ss[len(ss)-1], 10)
	if err != nil {
		log.Println("не получилось конвертировать вес в число")
	}
	ss = ss[:len(ss)-1]
	s2 := strings.Join(ss, " ")
	ing.Name = s2
	ing.Gram = gram
	return ing, err
}
