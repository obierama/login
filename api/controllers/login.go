package controllers

import (
	"encoding/json"
	"io/ioutil"
	"login/api/model"
	"login/config/middlewares"
	"login/config/respon"
	"net/http"
)

func (a *App) Login(w http.ResponseWriter, r *http.Request) {

	user := &model.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	usr, err := user.GetUser(a.DB, "email = ?", user.Email)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return

	}

	err = model.CheckPasswordHash(user.Password, usr.Password)
	if err != nil {
		respon.FAILED(w, http.StatusBadRequest, "Login failed, check your password", nil)
		return
	}

	token, err := middlewares.EncodeAuthToken(usr.ID)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["token"] = token
	resp["data"] = usr
	respon.JSON(w, http.StatusOK, resp)
	return

}

// func (a *App) GetDetail(w http.ResponseWriter, r *http.Request) {
// 	resp = map[string]interface{}{"status": true, "message": "Succes", "code": 200}
// 	user := &model.User{}
// 	userId := r.Context().Value("userID").(float64)
// 	id := int(userId)
// 	users, err := user.GetUserInt(a.DB, "id = ?", int(id))
// 	if err != nil {
// 		respon.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	resp["data"] = users
// 	respon.JSON(w, http.StatusOK, resp)
// 	return
// }
