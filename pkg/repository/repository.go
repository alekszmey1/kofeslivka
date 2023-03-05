package repository

import (
	"awesomeProject/kofeslivka/pkg/entity"
	"database/sql"
	"fmt"

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
func (r *repository) CreateUser(ingredient *entity.RawMaterial) (int64, error) {

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
