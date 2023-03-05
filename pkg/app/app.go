package app

import (
	controller2 "awesomeProject/kofeslivka/pkg/controller"
	repository2 "awesomeProject/kofeslivka/pkg/repository"
	usecase2 "awesomeProject/kofeslivka/pkg/usecase"
	"fmt"
	"net/http"
)

func Run() {
	repository := repository2.NewRepository()
	usecase := usecase2.NewUsecase(repository)
	controller := controller2.NewController(usecase)
	//mux := http.NewServeMux()
	fmt.Println("запуск")
	//mux.HandleFunc("/create", controller.CreateUser)
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/create", controller.Create)
	http.HandleFunc("/save_article", controller.SaveArticle)

	//	mux.HandleFunc("/make_friends", controller.MakeFriends)
	//	mux.HandleFunc("/delete", controller.DeleteUser)
	//	mux.HandleFunc("/get_friends", controller.GetFriends)
	//	mux.HandleFunc("/put", controller.UpdateAge)
	http.ListenAndServe(":8181", nil)
}

/*
func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", saveArticle)
	http.ListenAndServe(":8181", nil)
}*/
