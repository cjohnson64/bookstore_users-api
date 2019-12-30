package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/crjohnson1208/bookstore_users-api/domain/users"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/crjohnson1208/bookstore_users-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User
	
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: return bad request to the caller
		return	
	}
	if err := json.Unmarshal(bytes, &user); err!= nil{
		//TODO: handle user creation error
		return
	}
	result, saveErr := services.CreateUser()


	fmt.Println(user)
	c.String(http.StatusNotImplemented, "implement me!")
}

func GetUser(c *gin.Context) {
	
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}