package handlers

import (
	"ApiRestGo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtienen todos los usuarios\n")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Se obtienen un usuario\n")

	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])

	response := models.GetDefaultResponse(w)
	user, err := models.GetUser(userId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFound(err.Error())
	} else {
		response.Data = user
	}
	// output, _ := json.Marshal(&response)
	// fmt.Fprintf(w, string(output))
	response.Send()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuairo\n")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuairo\n")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se Elimina un usuairo\n")
}
