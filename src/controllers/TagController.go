package controllers

import (
	//"container/list"
	"fmt"
	"github.com/gin-gonic/gin"
	"models/tagmap"
	"strings"
	//"io/ioutil"
	//"encoding/json"
)

func TagPost(c *gin.Context) {
	c.Request.ParseForm()

	var json tagmap.TagMap
	var k = c.Bind(&json)
	if nil != k {
		fmt.Println("failed to bind", k)
		c.JSON(400, "incorrect json format")
	} else {
		var status bool
		var error string
		var data interface{}
		fmt.Println("bound successfully to a tag - value -- %+v\n", json)
		fmt.Printf("%+v\n", json)
		status, error, data = tagmap.SaveTagMap(json)
		fmt.Println("tagpost : data returned is  - ", data)
		if status {
			c.JSON(200, GetHttpResult(200, ""))
		} else if error != "" {
			c.JSON(400, GetHttpResult(400, error))
		} else {
			c.JSON(500, GetHttpResult(500, "Some Error Occurred in tagpost"))
		}
		//c.JSON(400, GetHttpResult(400, "error while binding person in personpost func"))
	}
	// mapped to a model TagMap -- now save this object to mongodb like in person
}

func TagGet(c *gin.Context) {
	fmt.Println("here")
	c.Request.ParseForm()

	emailid := c.Request.Form.Get("emailId")
	tagList := strings.Split(c.Request.Form.Get("tagList"), ",")

	var result interface{}
	var code = 200
	var status bool
	var message string

	if emailid != "" {
		fmt.Println("emailid found is - ", emailid)
		status, message, result = tagmap.GetTagData(emailid, tagList)
		fmt.Println("status/message : ", status, message)
	} else {
		fmt.Println("get tags - incorrect request format")
		code = 400
		result = GetHttpResult(400, "get tags - emailid missing")
	}

	c.JSON(code, result)
}
