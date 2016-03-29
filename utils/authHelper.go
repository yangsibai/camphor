package utils

import (
	"net/http"
)

const Authed string = "authed"

func IsLogin(r *http.Request) (bool, error) {
	session, err := store.Get(r, "user-auth")
	if err != nil {
		return false, err
	}
	if authed, ok := session.Values[Authed].(bool); ok {
		return authed, nil
	}
	return false, nil
}

func Login(w http.ResponseWriter, r *http.Request) (success bool, err error) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	if email == Config.Auth.Email && password == Config.Auth.Password {
		session, err := store.Get(r, "user-auth")
		if err != nil {
			return false, err
		}
		session.Values["authed"] = true
		session.Save(r, w)
		return true, nil
	}
	return false, nil
}
