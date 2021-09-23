package controllers

import (
	"encoding/json"
	"io/ioutil"
	"login/api/model"
	"login/config/respon"
	"net/http"
)

var (
	resp = map[string]interface{}{"status": true, "message": "Succes", "code": 200}
)

func (a *App) Register(w http.ResponseWriter, r *http.Request) {

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

	err = user.Validate("")
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userCreated, err := user.SaveUser(a.DB)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = userCreated
	respon.JSON(w, http.StatusCreated, resp)
	return

}
