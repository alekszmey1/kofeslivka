package app

import (
	controller2 "awesomeProject/kofeslivka/pkg/controller"
	repository2 "awesomeProject/kofeslivka/pkg/repository"
	usecase2 "awesomeProject/kofeslivka/pkg/usecase"
	"net/http"
)

func Run() {
	repository := repository2.NewRepository()
	usecase := usecase2.NewUsecase(repository)
	controller := controller2.NewController(usecase)
	mux := http.NewServeMux()

	mux.HandleFunc("/create", controller.CreateUser)
	//	mux.HandleFunc("/make_friends", controller.MakeFriends)
	//	mux.HandleFunc("/delete", controller.DeleteUser)
	//	mux.HandleFunc("/get_friends", controller.GetFriends)
	//	mux.HandleFunc("/put", controller.UpdateAge)
	http.ListenAndServe("localhost:8080", mux)
}
