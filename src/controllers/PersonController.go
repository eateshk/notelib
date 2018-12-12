package controllers

import (
	//"container/list"
	"fmt"
	"models/person"
	// "reflect"
	"github.com/gin-gonic/gin"
)

func PersonsGet(c *gin.Context) {
	c.Request.ParseForm()
	emailid := c.Request.Form.Get("emailid")
	p := new(person.Person)
	p.EmailId = "ekandpal@gmail.com"
	var result interface{}
	var code = 200
	if emailid != "" {
		fmt.Println("emailid found is - ", emailid)
		result = person.GetPersons(*p)
	} else {
		fmt.Println("incorrect request format")
		code = 400
		result = GetHttpResult(400, "emailid missing")
	}
	c.JSON(code, result)
}

// This should never be used, it should not be supported by a rest api, or fix the inputs to just indexes. Wildcard search in person could be expensive.
func PersonGet(c *gin.Context) {
	c.Request.ParseForm()
	emailid := c.Request.Form.Get("emailid")
	var result interface{}
	var code = 200
	if emailid != "" {
		fmt.Println("emailid found is - ", emailid)
		result = person.GetPerson(emailid)
	} else {
		fmt.Println("incorrect request format")
		result = GetHttpResult(400, "emailid missing")
		code = 400
	}
	c.JSON(code, result)
}

func PersonPost(c *gin.Context) {
	var json person.Person

	if nil == c.Bind(&json) {
		var status bool
		var error string
		var data interface{}
		status, error, data = person.SavePerson(json)
		fmt.Println("data returned is  - ", data)
		if status {
			c.JSON(200, GetHttpResult(200, ""))
		} else if error != "" {
			c.JSON(400, GetHttpResult(400, error))
		} else {
			c.JSON(500, GetHttpResult(500, "Some Error Occurred in personpost"))
		}
	} else {
		c.JSON(400, GetHttpResult(400, "error while binding person in personpost func"))
	}
}

func PersonUpdate(c *gin.Context) {

	var json person.Person
	if nil == c.Bind(&json) {
		if person.UpdatePerson(json) {
			c.JSON(200, GetHttpResult(200, ""))
		} else {
			c.JSON(500, GetHttpResult(200, "some error while updating person, personupdate func"))
		}
		return
	}
	c.JSON(400, GetHttpResult(400, "error while binding in personupdate func"))
}

func PersonDelete(c *gin.Context) {

	var json person.Person
	if nil == c.Bind(&json) {
		if person.DeletePerson(json) {
			fmt.Println("data to send now is 500")
			c.JSON(200, GetHttpResult(200, ""))
		} else {
			c.JSON(500, GetHttpResult(500, "some error occurred while deleting person , persondelete func"))
		}
		return
	}
	c.JSON(400, GetHttpResult(400, "error in binding person, persondelete func"))
}



func Authentication(c *gin.Context) {

	var json person.Person
	if nil == c.Bind(&json){
		_,messsage := json.Authenticate()

		c.JSON(200,GetHttpResult(200,messsage))
		return
	}
	c.JSON(400,GetHttpResult(400,"Error "))

}
