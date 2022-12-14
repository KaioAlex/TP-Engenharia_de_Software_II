package controllers

import (
	"net/http"
	"time"

	"ll/context"
	"ll/email"
	"ll/models"
	"ll/rand"
	"ll/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers(us models.UserService, emailer *email.Client) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
		emailer:   emailer,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        models.UserService
	emailer   *email.Client
}

// New is used to render the form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	parseURLParams(r, &form)
	u.NewView.Render(w, r, form)
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is used to process the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form SignupForm
	vd.Yield = &form
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		u.NewView.Render(w, r, vd)
		return
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		vd.SetAlert(err)
		u.NewView.Render(w, r, vd)
		return
	}
	u.emailer.Welcome(user.Name, user.Email)
	err := u.signIn(w, &user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	alert := views.Alert{
		Level:   views.AlertLvlSuccess,
		Message: "Welcome to SpotImages!",
	}
	views.RedirectAlert(w, r, "/galleries", http.StatusFound, alert)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) SetErrorInLoginView(w http.ResponseWriter, r *http.Request, err error, customAlertError string) {
	vd := views.Data{}
	if customAlertError != "" {
		vd.AlertError(customAlertError)
	} else {
		vd.SetAlert(err)
	}

	u.LoginView.Render(w, r, vd)
}

func getLoginFormFromRequest(r *http.Request) (LoginForm, error) {
	var form LoginForm
	err := parseForm(r, &form)
	return form, err
}

func getAuthenticateError(err error) string {
	customAlertError := ""
	switch err {
	case models.ErrNotFound:
		customAlertError = "Invalid email address"
	default:
		customAlertError = ""
	}
	return customAlertError
}

// Login is used to verify the provided email address and
// password and then log the user in if they are correct.
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {

	form, err := getLoginFormFromRequest(r)
	if err != nil {
		u.SetErrorInLoginView(w, r, err, "")
		return
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		customAlertError := getAuthenticateError(err)
		u.SetErrorInLoginView(w, r, err, customAlertError)
		return
	}

	err = u.signIn(w, user)
	if err != nil {
		u.SetErrorInLoginView(w, r, err, "")
		return
	}

	http.Redirect(w, r, "/galleries", http.StatusFound)
}

// Logout is used to delete a users session cookie (remember_token)
// and then will update the user resource with a new remember token.
//
// POST /logout
func (u *Users) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	user := context.User(r.Context())
	token, _ := rand.RememberToken()
	user.Remember = token
	u.us.Update(user)
	http.Redirect(w, r, "/", http.StatusFound)
}

// signIn is used to sign the given user in via cookies
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}

	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.Remember,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}
