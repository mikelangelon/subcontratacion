package main

import (
	"context"
	"fmt"
	"html/template"
	"mikelangelon/m/v2/api/web"
	"mikelangelon/m/v2/internal/app/budgetrequest"
	"mikelangelon/m/v2/internal/app/company"
	"mikelangelon/m/v2/internal/app/user"
	"mikelangelon/m/v2/internal/pkg/db"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	db := setupMongo()

	services := setupServiceLayer(db)

	server := api(services)

	fmt.Println("Listening in 8080...")
	http.ListenAndServe(":8080", server)
}

type services struct {
	User    user.UserService
	Company company.CompanyService
}

type dbRepos struct {
	UserDB    db.UserRepo
	CompanyDB db.CompanysRepo
}

func setupServiceLayer(dbRepos dbRepos) services {
	return services{
		User:    user.New(dbRepos.UserDB),
		Company: company.New(dbRepos.CompanyDB),
	}
}

func setupMongo() dbRepos {
	database := initMongoDB()
	return dbRepos{
		UserDB:    db.NewUserRepo(database),
		CompanyDB: db.NewCompanyRepo(database),
	}
}

func initMongoDB() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client.Database("subcon")
}

func api(services services) *http.ServeMux {
	// #### Define resources ####
	budgetService := budgetrequest.New()
	homeR := web.HomeResource{
		CompanyService: services.Company,
		BudgetService:  budgetService,
	}
	// companyR := web.CompanyResource{
	// 	CompanyService: companyService,
	// 	OkCall:         homeR.Home,
	// }
	registrationR := web.RegistrationResource{
		CompanyService: services.Company,
		UserService:    services.User,
		OkCall:         homeR.Home,
	}
	budgetRequestR := web.BudgetRequestResource{
		BudgetService: budgetService,
		OkCall:        homeR.Home,
	}

	loginR := web.LoginResource{
		UserService: services.User,
		OkCall:      homeR.Home,
	}

	profileR := web.ProfileResource{}
	// #### Define paths ####
	s := http.NewServeMux()
	s.HandleFunc("/", homeR.Home)
	s.HandleFunc("/logout", loginR.Logout)
	s.HandleFunc("/login", loginR.Login)
	s.HandleFunc("/secret", serveFWithAuth(serveStatic))
	s.HandleFunc("/login-form", func(writer http.ResponseWriter, request *http.Request) {
		loginR.LoginForm(writer, request, nil)
	})
	s.HandleFunc("/register-form", serveRegisterForm)
	s.HandleFunc("/budget-request-form", serveBudgetRequestForm)
	s.HandleFunc("/budget-request-view", budgetRequestR.View)
	s.HandleFunc("/budget-request-search", budgetRequestR.Search)
	s.HandleFunc("/offer-form", serveOfferForm)
	s.HandleFunc("/register", registrationR.Register)
	s.HandleFunc("/budgetrequest", budgetRequestR.Create)
	s.HandleFunc("/offer", budgetRequestR.Create)
	s.HandleFunc("/profile", profileR.Profile)

	s.HandleFunc("/ayuda", serveHelp)
	s.HandleFunc("/contacto", serveContact)

	// #### Define static resources ####
	fs := http.FileServer(http.Dir("web/assets/"))
	s.Handle("/static/", http.StripPrefix("/static/", fs))
	s.HandleFunc("/home", homeR.Home)
	return s
}

type serveFunc func(http.ResponseWriter, *http.Request)

func serveFWithAuth(f serveFunc) serveFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, valid := web.LoginResource{}.Auth(w, r)
		if !valid {
			serveRegisterForm(w, r)
			return
		}
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}
func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/example.html")
	if err != nil {
		fmt.Println(err)
	}
	items := struct {
		Country string
		City    string
	}{
		Country: "Australia",
		City:    "Paris",
	}
	t.Execute(w, items)
}

func serveRegisterForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"web/template/simple-layout.html",
		"web/template/topmenu.html",
		"web/content/registration-form.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func serveBudgetRequestForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"web/template/layout.html",
		"web/template/topmenu.html",
		"web/template/leftmenu.html",
		"web/template/banner.html",
		"web/content/budgetrequest-form.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func serveBudgetRequestSearch(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"web/template/layout.html",
		"web/template/topmenu.html",
		"web/template/leftmenu.html",
		"web/template/banner.html",
		"web/content/budgetrequest-search.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func serveHelp(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"web/template/layout.html",
		"web/template/topmenu.html",
		"web/template/leftmenu.html",
		"web/template/banner.html",
		"web/content/help.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func serveContact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"web/template/layout.html",
		"web/template/topmenu.html",
		"web/template/leftmenu.html",
		"web/template/banner.html",
		"web/content/contact.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func serveOfferForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/offerform.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func basicFiles(content string) []string {
	return []string{
		"web/template/topmenu.html",
		content,
	}
}
