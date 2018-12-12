package main

import (
	"controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"fmt"
	"net/http"
	"reflect"
	"utils/logger"
)

func main() {
	r := gin.Default() // r for router

	fmt.Println("Router type is : ", reflect.TypeOf(r))
	//for session store creation and usage
	store := sessions.NewCookieStore([]byte("secret"))

	r.Use(sessions.Sessions("portal_login", store))

	r.Static("/assets", "./../static")
	r.LoadHTMLGlob("./../templates/*")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./../static/images/favicon.ico")
	r.StaticFile("/home", "./../static/clean_index.html")
	r.StaticFile("/l", "./../static/mylogin.html")

	/*
		//Loading template below
		r.LoadHTMLGlob("./../templates/*")
		//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
		r.GET("/templatetest", func(c *gin.Context) {
			logger.Log("debug", "first html renderer")
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Welcome to your tag library, Eatesh",
			})
		})
	*/

	/*
		// Loading using ember build.
		r.Static("/assets", "./../public/assets")
		r.StaticFile("/", "./../public")
	*/

	/*
		// test group v0
		v0 := r.Group("/")
		{
			v0.GET("/", controllers.RootStatusGet)
		}
	*/

	// GENERIC APIS BELOW, WHICH DON'T REQUIRE A VERSION.

	genericApis := r.Group("/")
	{

		genericApis.GET("/", root)

		genericApis.GET("/login", loginPage)

		genericApis.POST("/login", login)

		genericApis.GET("/logout", logout)

	}

	// versioned apis below
	v1 := r.Group("/v1")
	{
		logger.Log("debug", "This is test error log")
		v1.GET("/", controllers.V1StatusGet)
		v1.GET("/person", controllers.PersonGet)
		v1.GET("/persons", controllers.PersonsGet)
		v1.POST("/person", controllers.PersonPost)
		v1.PATCH("/person", controllers.PersonUpdate)
		v1.DELETE("/person", controllers.PersonDelete)
		v1.POST("/person/authenticate",controllers.Authentication)

		/*
			type - post
			content-type - application/json
			payload -  {"Tags" : ["1", "two", "3"], "EmailId" : "kandpaleatesh@fiberlink.com", "TagData" : "This is my testtag", "UpdateTime" : "forty", "Others": {"hell" : "bell", "heaven" : "tank"}}
			url - http://localhost:3000/v1/tag
		*/
		v1.POST("/tag", controllers.TagPost)

		//http://localhost:3000/v1/tag?emailId=kandpaleatesh@fiberlink.com&tagList=three,two
		v1.GET("/tag", controllers.TagGet)

		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/", func(c *gin.Context) {
			// You need to call ParseForm() on the request to receive url and form params first
			session := sessions.Default(c)
			k := session.Get("answer")
			message := "Helloes, this is just a test v2 api, its not implemented yet !!"

			if k == 44 {
				message = "You're now logged in it seems."
			} else {
				message = "You're not logged in yet."
				session.Set("answer", 44)
			}
			session.Save()

			c.Request.ParseForm()
			fmt.Println("Router type is : ", reflect.TypeOf(session))

			c.String(200, message)
		})
		// v2.POST("/login", controllers.LoginEndpoint)
		//v2.POST("/submit", submitEndpoint)
		//v2.POST("/read", readEndpoint)
	}

	// Listen and server on 0.0.0.0:3000
	r.Run(":3000")
} // main ends here.

func isLoggedIn(c *gin.Context) string {
	var s sessions.Session
	k := c.Keys["github.com/gin-gonic/contrib/sessions"]
	s = k.(sessions.Session)
	value := s.Get("emailId")
	emailId, exists := value.(string)
	fmt.Println("checking if login, emailid found is : ", emailId, exists)
	return emailId

	/*
	 //required during logout.
	s.Delete(c)
	s.Clear()
	err := s.Save()
	fmt.Println("err in delete>save : ", err)

	fmt.Println(s.Get("emailId"))
	c.Keys["github.com/gin-gonic/contrib/sessions"].(sessions.Session).Delete(c)
	k = c.Keys["github.com/gin-gonic/contrib/sessions"]
	s = k.(sessions.Session)
	fmt.Println(s.Get("answer"))
	*/
	/*	session := sessions.Default(c)
		fmt.Println("Session type is : ", reflect.TypeOf(session))
		fmt.Println("context keys are : ", reflect.TypeOf(c.Keys))
			keys := make([]string, 0, len(c.Keys))
			for k := range c.Keys {
				keys = append(keys, k)
				fmt.Println("key is : ", k)
			}
	*/
	//fmt.Println("type, value --", reflect.TypeOf(keys[0]), keys[0])
	//fmt.Println(c.Keys[keys[0]])
	//fmt.Println("Session name is : ", reflect.TypeOf(session.name))
	//return sessions.Default(c) != nil
}

func root(c *gin.Context) {
	fmt.Println(sessions.Default(c))
	if isLoggedIn(c) == "" {
		c.Redirect(http.StatusMovedPermanently, "http://eateshkandpal.com/home")
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"message" : "You have logged in successfully, one can't reach this page if not logged in but this is all we've built till now after login :)",
			"title": "Build something amazing guys. This is all you've built till now :P. ",
		})
		// c.String(200, "Whoa ! you're logged in bro!")
	}
}

func loginPage(c *gin.Context) {
	if isLoggedIn(c) == "" {
		c.Redirect(http.StatusMovedPermanently, "http://eateshkandpal.com/l")
	} else {
		c.Redirect(http.StatusMovedPermanently, "http://eateshkandpal.com/")
	}
}


func login(c *gin.Context) {
	if 	isLoggedIn(c) == "" {
		// do login process here.
		emailId := c.PostForm("emailid")
		password := c.PostForm("password")
		// Create a new module for checking credentials later.
		if emailId == "testing@testing.com" && password == "welcomepeople" {
			// go to a welcome page
			// save session
			session := sessions.Default(c)
			session.Set("emailId", "testing@testing.com")
			session.Save()
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"message" : "You have logged in successfully, one can't reach this page if not logged in but this is all we've built till now after login :)",
				"title": "Build something amazing guys. This is all you've built till now :P. ",
			})
		} else {
			// go to login page again, make it a template so that error could be shown.
			c.HTML(http.StatusOK, "loginpage.tmpl", gin.H{
				"message": "Put correct login details here, or i'll give you a suckerpunch !! ;)",
			})

		}
		fmt.Println("email/exists : " , emailId)
		fmt.Println("password/exists : " , password)

	} else {
		// the dudette is logged in already. Take him to welcome page directly.
		fmt.Println("This should never happen, Alert! Because GET:login won't show form in case guy is logged in.")
		c.Redirect(http.StatusMovedPermanently, "http://eateshkandpal.com/welcome") // this ain't created yet.
	}
}

func logout(c *gin.Context) {

	if sessions.Default(c) != nil {
		sessions.Default(c).Clear()
		sessions.Default(c).Save()
		fmt.Println("Deleted/saved session data !")
	}
	c.String(200, "Logged out successfully")
}
