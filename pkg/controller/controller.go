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
	/*db, err := sql.Open(s, root)
	if err != nil {
		fmt.Println("не получилось открыть базу данных")
		panic(err)
	}
	defer db.Close()
	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var post Articles
		err = res.Scan(&post.id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)*/
	//fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Title, post.id))

	t.ExecuteTemplate(w, "index", nil)
}

/*
	func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			user := &entity.RawMaterial{}
			err := json.NewDecoder(r.Body).Decode(user)
			if err != nil {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusBadRequest)
			}
			id, err := c.usecase.CreateUser(user)
			if err != nil {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusBadRequest)
			}
			result := map[string]int64{"id": id}
			response, err := json.Marshal(result)
			if err != nil {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusCreated)
			w.Write(response)
		}
	}
*/

func (c *Controller) SaveArticle(w http.ResponseWriter, r *http.Request) {
	ing := &entity.RawMaterial{}
	ing.Name = r.FormValue("name")
	ing.Protein = stringToFloat(r.FormValue("protein"))
	ing.Fats = stringToFloat(r.FormValue("fats"))
	ing.Carbohydrates = stringToFloat(r.FormValue("carbohydrates"))
	ing.Calories = stringToFloat(r.FormValue("calories"))
	fmt.Println(ing)
	id, err := c.usecase.CreateUser(ing)
	/*if ing.Name == "" || ing.Protein == 0 || ing.Fats == 0 || ing.Carbohydrates == 0 || ing.Calories == 0 {
		fmt.Fprintf(w, "не все данные заполнены")
	} else {*/

	fmt.Println(id)
	if err != nil {
		panic(err)
	}
	/*db, err := sql.Open(s, root)
	if err != nil {
		fmt.Println("не получилось открыть базу данных")
		panic(err)
	}
	defer db.Close()
	insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles`(`titke`,`anons`,"+
		"`full_text`) 	VALUES ('%s','%s','%s')", title, anons, fullText))
	if err != nil {
		panic(err)
	}
	defer insert.Close()*/
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

func stringToFloat(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return result
}
