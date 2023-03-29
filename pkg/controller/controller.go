package controller

import (
	"awesomeProject/kofeslivka/pkg/entity"
	"awesomeProject/kofeslivka/pkg/usecase"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func (c *Controller) SaveArticle(w http.ResponseWriter, r *http.Request) {
	ing := &entity.RawMaterial{}
	ing.Name = r.FormValue("name")
	ing.Protein = stringToFloat(r.FormValue("protein"))
	ing.Fats = stringToFloat(r.FormValue("fats"))
	ing.Carbohydrates = stringToFloat(r.FormValue("carbohydrates"))
	ing.Calories = stringToFloat(r.FormValue("calories"))
	fmt.Println(ing)
	id, err := c.usecase.AddIngredient(ing)
	fmt.Println(id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("креате")
	t, err := template.ParseFiles("../templates/create.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create", nil)
}

func (c *Controller) CreateDesert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("креате десерт")
	t, err := template.ParseFiles("../templates/create_desert.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create_desert", nil)
}

func stringToFloat(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return result
}
func (c *Controller) MakeCake(w http.ResponseWriter, r *http.Request) {
	des := &entity.Desert{}
	des.Name = r.FormValue("name_desert")
	des.List = r.FormValue("list")
	fmt.Println(des)
	c.usecase.MakeDesert(des)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
