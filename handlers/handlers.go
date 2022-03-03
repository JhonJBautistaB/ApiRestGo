package handlers

import (
	"ApiRestGo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Se obtienen todos los usuarios\n")
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Se obtienen un usuario\n")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	//response := models.CreateDefaultResponse(w)
	//user, err := models.GetUser(userId)

	if user, err := models.GetUser(userId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		//response.NotFound(err.Error())
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
	// output, _ := json.Marshal(&response)
	// fmt.Fprintf(w, string(output))
	//response.Send()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Se crea un usuairo\n")
	user := models.Users{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuairo\n")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se Elimina un usuairo\n")
}
