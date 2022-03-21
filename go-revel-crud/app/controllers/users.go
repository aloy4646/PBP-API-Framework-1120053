package controllers

import (
	"cobaRevel/go-revel-crud/app/db"
	"cobaRevel/go-revel-crud/app/entities"
	"cobaRevel/go-revel-crud/app/models"
	"log"
	"strconv"

	"net/http"

	"github.com/revel/revel"
)

type Users struct {
	App
}

func (c Users) GetAllUsers() revel.Result {

	db := db.Connect()
	defer db.Close()

	query := "SELECT id,name,age,address,email,password from users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		var errorResult entities.Response
		errorResult.Success = false
		errorResult.Status = 401
		errorResult.Message = "Query Error"
		return c.RenderJSON(errorResult)
	}

	var user models.User
	var users []models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
			log.Println(err.Error())
			var errorResult entities.Response
			errorResult.Success = false
			errorResult.Status = 402
			errorResult.Message = "Could not get the users"
			return c.RenderJSON(errorResult)
		} else {
			users = append(users, user)
		}
	}

	var result = entities.UsersResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Success",
		Data:    users,
	}
	return c.RenderJSON(result)
}

func (c Users) InsertNewUser() revel.Result {
	db := db.Connect()
	defer db.Close()

	name := c.Params.Form.Get("name")
	age, _ := strconv.Atoi(c.Params.Form.Get("age"))
	address := c.Params.Form.Get("address")
	email := c.Params.Form.Get("email")
	password := c.Params.Form.Get("password")

	resultQuery, errQuery := db.Exec("INSERT INTO users (name, age, address, email, password) VALUES (?,?,?,?,?)",
		name,
		age,
		address,
		email,
		password,
	)

	if errQuery != nil {
		log.Println(errQuery)
		var errorResult entities.Response
		errorResult.Success = false
		errorResult.Status = 401
		errorResult.Message = "Query Error"
		return c.RenderJSON(errorResult)
	}

	id, _ := resultQuery.LastInsertId()
	var user models.User = models.User{Id: int(id), Name: name, Age: age, Address: address, Email: email, Password: password}

	var result = entities.UserResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Success",
		Data:    user,
	}
	return c.RenderJSON(result)
}

func (c Users) UpdateUser() revel.Result {
	db := db.Connect()
	defer db.Close()

	id, _ := strconv.Atoi(c.Params.Form.Get("id"))
	name := c.Params.Form.Get("name")
	age, _ := strconv.Atoi(c.Params.Form.Get("age"))
	address := c.Params.Form.Get("address")
	email := c.Params.Form.Get("email")
	password := c.Params.Form.Get("password")

	resultQuery, errQuery := db.Exec("UPDATE users SET name=?, age=?, address=?, email=?, password=? WHERE id=?",
		name,
		age,
		address,
		email,
		password,
		id,
	)

	if errQuery != nil {
		log.Println(errQuery)
		var errorResult entities.Response
		errorResult.Success = false
		errorResult.Status = 401
		errorResult.Message = "Query Error"
		return c.RenderJSON(errorResult)
	}

	rowsAffected, _ := resultQuery.RowsAffected()
	return responseFromRowsAffected(c, rowsAffected)
}

func (c Users) DeleteUser() revel.Result {
	db := db.Connect()
	defer db.Close()

	id := c.Params.Route.Get("id")

	resultQuery, err := db.Exec("DELETE FROM users WHERE id=?",
		id,
	)

	if err != nil {
		log.Println(err)
		var errorResult entities.Response
		errorResult.Success = false
		errorResult.Status = 401
		errorResult.Message = "Query Error"
		return c.RenderJSON(errorResult)
	}

	rowsAffected, _ := resultQuery.RowsAffected()
	return responseFromRowsAffected(c, rowsAffected)
}

func responseFromRowsAffected(c Users, rowsAffected int64) revel.Result {
	var result entities.Response
	if rowsAffected > 0 {
		result.Success = true
		result.Status = http.StatusOK
		result.Message = "Success, " + strconv.FormatInt(rowsAffected, 10) + " rows affected"
	} else {
		result.Success = false
		result.Status = 407
		result.Message = "Failed, 0 rows affected"
	}

	return c.RenderJSON(result)
}
