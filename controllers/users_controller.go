package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"crud/models"
	"crud/utils"

	"github.com/gorilla/mux"
)

// GetUser obtiene un user por su ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Estructura vacia donde se gurdarán los datos
	users := models.User{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM contacts WHERE ID = ?
	db.Find(&users, id)
	// Se comprueba que exista el registro
	if users.ID > 0 {
		// Se codifican los datos a formato JSON
		j, _ := json.Marshal(users)
		// Se envian los datos
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

// GetUsers obtiene todos los contactos
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Slice (array) donde se guardaran los datos
	users := []models.Contact{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM users
	db.Find(&users)
	// Se codifican los datos a formato JSON
	j, _ := json.Marshal(users)
	// Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)
}

// StoreUser guarda un nuevo user
func StoreUser(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se gurdaran los datos del body
	user := models.User{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se decodifican los datos del body a la estructura user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	// Se guardan los datos en la DB
	err = db.Create(&user).Error
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}
	// Se codifica el nuevo registro y se devuelve
	j, _ := json.Marshal(user)
	utils.SendResponse(w, http.StatusCreated, j)
}

// UpdateUser modifica los datos de un user por su ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Estructuras donde se almacenaran los datos
	userFind := models.User{}
	userData := models.User{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se buscan los datos
	db.Find(&userFind, id)
	if userFind.ID > 0 {
		// Si existe el registro se decodifican los datos del body
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			// Sí hay algun error en los datos se devolvera un error 400
			utils.SendErr(w, http.StatusBadRequest)
			return
		}
		// Se modifican los datos
		db.Model(&userFind).Updates(userData)
		// Se codifica el registro modificado y se devuelve
		j, _ := json.Marshal(userFind)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

// DeleteUser elimina un user por ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se guardara el registo buscado
	user := models.User{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se busca el user
	db.Find(&user, id)
	if user.ID > 0 {
		// Sí existe, se borra y se envia contenido vacio
		db.Delete(user)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}
