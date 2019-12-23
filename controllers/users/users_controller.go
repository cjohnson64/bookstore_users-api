package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/crjohnson1208/bookstore_users-api/domain/users"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func CreateUser(c *gin.Context) {
	var user users.User
	
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return	
	}
	if err := json.Unmarshall(bytes, &user); err!= nil{
		//TODO: Handle Json Error
		return
	}
	fmt.Println(user)
	c.String(http.StatusNotImplemented, "implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
