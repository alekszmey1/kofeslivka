package repository

import (
	"awesomeProject/kofeslivka/pkg/entity"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const s = "mysql"
const root = "root:root@tcp(127.0.0.1:3307)/kofeslivka"

type repository struct {
}

// Функция создания хранилища
func NewRepository() *repository {
	return &repository{}
}

// Метод добавления в хранилище новую сущность
func (r *repository) AddIngredient(ingredient *entity.RawMaterial) (int64, error) {
	db, err := sql.Open(s, root)
	if err != nil {
		fmt.Println("не получилось открыть базу данных")
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into rawmaterial (name ,protein, fats, carbohydrates, calories) values (?,?,?,?,?)",
		ingredient.Name, ingredient.Protein, ingredient.Fats, ingredient.Carbohydrates, ingredient.Calories)
	if err != nil {
		panic(err)
	}
	ingredient.Id, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return ingredient.Id, nil
}
func (r *repository) MakeDesert(des *entity.Desert) (string, error) {
	db, err := sql.Open(s, root)
	if err != nil {
		fmt.Println("не получилось открыть базу данных")
		panic(err)
	}
	defer db.Close()
	listIng := makeSliceIng(des.List)
	desert := &entity.RawMaterial{}
	desert.Name = des.Name
	for _, ing := range listIng {
		insert, err := db.Query(fmt.Sprintf("SELECT * FROM `rawmaterial` WHERE `name` LIKE %s", ing.Name))
		if err != nil {
			log.Printf("не получилось получить данные ингредиента")
		}
		defer insert.Close()
		u := entity.RawMaterial{}
		for insert.Next() {
			err := insert.Scan(&u)
			calcCalories(u, *ing, *desert)
			fmt.Println(u)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		fmt.Println(desert)
	}
	return "Десерт добавлен", nil

}
func makeSliceIng(s string) []*entity.IngredientGram {
	var listIng []*entity.IngredientGram
	ss := strings.Split(s, ";")
	for _, s2 := range ss {
		ing, err := stringToIng(s2)
		if err != nil {
			log.Printf("не получилось конвертировать строку в игредиент")
		}
		listIng = append(listIng, ing)
	}
	return listIng
}
func stringToIng(s string) (*entity.IngredientGram, error) {
	e := &entity.IngredientGram{}
	ss := strings.Split(s, " ")
	gram, err := strconv.ParseFloat(ss[len(ss)-1], 10)
	if err != nil {
		log.Println("не получилось обработать грамм")
	}
	ss = ss[:len(ss)-1]
	s2 := strings.Join(ss, " ")
	e.Gram = gram
	e.Name = s2
	return e, err
}
func calcCalories(u entity.RawMaterial, ing entity.IngredientGram, desert entity.RawMaterial) entity.RawMaterial {
	desert.Fats = desert.Fats + u.Fats/100*ing.Gram
	desert.Protein = desert.Protein + u.Protein/100*ing.Gram
	desert.Carbohydrates = desert.Carbohydrates + u.Carbohydrates/100*ing.Gram
	desert.Calories = desert.Calories + u.Calories/100*ing.Gram
	return desert
}
