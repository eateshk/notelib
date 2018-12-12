package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"

)

func V1StatusGet(c *gin.Context) {
	/**TODO
		- call mongodb with current environment config, and check if we can connect.
		- check if there is any more dependencies to be checked.
	*/
	fmt.Println("Someone asked for v1 api status")
	c.JSON(200, GetHttpResult(200, "All Is Well :)"))
}


func RootStatusGet(c *gin.Context){
	fmt.Println("Someone asked for server status")
        c.JSON(200, GetHttpResult(200, "Hello, Nice to have you here. To do something you should explore some versioned api."))
}

