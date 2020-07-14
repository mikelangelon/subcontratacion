package web

import (
	"fmt"
	"html/template"
	"mikelangelon/m/v2/internal/app/user"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginResource struct {
	UserService user.UserService
	OkCall      func(w http.ResponseWriter, r *http.Request)
}

var jwtKey = []byte("my_secret_key")
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Claims struct {
	Profile Profile `json:"profile"`
	jwt.StandardClaims
}

type Profile struct {
	Username string `json:"username"`
}

func (r LoginResource) Auth(resp http.ResponseWriter, req *http.Request) (*Profile, bool) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			resp.WriteHeader(http.StatusUnauthorized)
			return nil, false
		}
		// For any other type of error, return a bad request status
		resp.WriteHeader(http.StatusBadRequest)
		return nil, false
	}

	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			resp.WriteHeader(http.StatusUnauthorized)
			return nil, false
		}
		resp.WriteHeader(http.StatusBadRequest)
		return nil, false
	}
	if !tkn.Valid {
		resp.WriteHeader(http.StatusUnauthorized)
		return nil, false
	}
	return &claims.Profile, true
}
func (r LoginResource) Logout(resp http.ResponseWriter, req *http.Request) {
	http.SetCookie(resp, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Minute),
	})
	http.Redirect(resp, req, "/login-form", http.StatusFound)
	return
}

type loginForm struct {
	Message *string
	Profile *Profile
}

func (r LoginResource) LoginForm(resp http.ResponseWriter, req *http.Request, message *string) {
	t, err := template.ParseFiles(
		"web/template/simple-layout.html",
		"web/template/topmenu.html",
		"web/content/login-form.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(resp, loginForm{
		Message: message,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (r LoginResource) Login(resp http.ResponseWriter, req *http.Request) {
	username := req.FormValue("user")
	password := req.FormValue("password")

	err := r.UserService.Auth(username, password)
	if err != nil {
		switch err {
		case user.ErrUserNotValid:
			message := "Usuario no existe."
			r.LoginForm(resp, req, &message)
			return
		default:
			message := "Algo fue mal. Intentelo mas tarde por favor."
			r.LoginForm(resp, req, &message)
			return
		}
	}
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Profile: Profile{
			Username: username,
		},
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		message := "Algo fue mal. Intentelo mas tarde por favor."
		r.LoginForm(resp, req, &message)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(resp, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	fmt.Print("redirect to home")
	http.Redirect(resp, req, "/home", http.StatusFound)
	return
	// r.OkCall(resp, req)
}
