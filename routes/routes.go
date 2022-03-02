package routes

import (
	"CRUD/models"
	"fmt"
	"mux/middleware"
	"mux/sessions"
	"mux/utils"
	"net/http"

	"github.com/gorilla/mux"
)

//Initiate a mux router
func NewRouter() *mux.Router {
	//Initiate a mux router
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.CheckAuth(index_GET_handler)).Methods("GET")
	r.HandleFunc("/", middleware.CheckAuth(index_POST_handler)).Methods("POST")

	r.HandleFunc("/login", login_GET_handler).Methods("GET")
	r.HandleFunc("/login", login_POST_handler).Methods("POST")

	r.HandleFunc("/register", reg_GET_handler).Methods("GET")
	r.HandleFunc("/register", reg_POST_handler).Methods("POST")

	//Initiate a file server
	fs := http.FileServer(http.Dir("./static/"))

	//Tell the router to use this file server for all paths starting with static prefix
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))

	return r
}

func index_GET_handler(w http.ResponseWriter, r *http.Request) {
	updates, err := models.GetUpdates()
	fmt.Println(updates)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	utils.ExecuteTemplate(w, "index.html", updates)
}

func index_POST_handler(w http.ResponseWriter, r *http.Request) {
	//Get the user Id first
	session, _ := sessions.Store.Get(r, "session")
	untypedUserId := session.Values["user_id"]
	userId, ok := untypedUserId.(int64)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	r.ParseForm()
	body := r.PostForm.Get("update")
	error := models.AddUpdate(userId, body)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func login_GET_handler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func login_POST_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	user, err := models.LoginUser(username, password)

	if err != nil {
		switch err {
		case models.ErrUserNotFound:
			utils.ExecuteTemplate(w, "login.html", "User does not exist!")
		case models.ErrInvalidPassword:
			utils.ExecuteTemplate(w, "login.html", "Password is incorrect!")
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error!"))
		}
		return
	}
	userId, err := user.GetUserId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}
	session, _ := sessions.Store.Get(r, "session")
	session.Values["user_id"] = userId
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func reg_GET_handler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func reg_POST_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := models.RegisterUser(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

// func test_session_GET_handler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "session")
// 	untyped, ok := session.Values["username"]

// 	if !ok {
// 		return
// 	}

// 	username, ok := untyped.(string)
// 	if !ok {
// 		return
// 	}

// 	w.Write([]byte(username))
// }
