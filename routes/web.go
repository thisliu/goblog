package routes

import (
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	pc := new(controllers.PagesController)
	ac := new(controllers.ArticlesController)
	auc := new(controllers.AuthController)
	uc := new(controllers.UserController)
	cc := new(controllers.CategoriesController)

	// 静态页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", ac.Index).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 文章相关页面
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 登陆相关
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")

	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	// 用户相关
	r.HandleFunc("/users/{id:[0-9]+}", uc.Show).Methods("GET").Name("users.show")

	// 标签相关
	r.HandleFunc("/categories/create", middlewares.Auth(cc.Create)).Methods("GET").Name("categories.create")
	r.HandleFunc("/categories/{id:[0-9]+}/edit", middlewares.Auth(cc.Edit)).Methods("GET").Name("categories.edit")
	r.HandleFunc("/categories/{id:[0-9]+}/articles", middlewares.Auth(cc.Articles)).Methods("GET").Name("categories.articles")
	r.HandleFunc("/categories/{id:[0-9]+}", middlewares.Auth(cc.Show)).Methods("GET").Name("categories.show")
	r.HandleFunc("/categories/{id:[0-9]+}/update", middlewares.Auth(cc.Update)).Methods("POST").Name("categories.update")
	r.HandleFunc("/categories/{id:[0-9]+}/delete", middlewares.Auth(cc.Delete)).Methods("POST").Name("categories.delete")
	r.HandleFunc("/categories", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")

	// 开始会话
	r.Use(middlewares.StartSession)

	// 中间件：强制内容类型为 HTML
	// r.Use(middlewares.ForceHTML)
}
