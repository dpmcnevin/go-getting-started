package main

import (
    "bytes"
    "log"
    "net/http"
    "os"
    "strconv"
    "fmt"
    "math/rand"


    "github.com/gin-gonic/gin"
    "github.com/russross/blackfriday"
)

var (
    repeat int
)

func repeatHandler(c *gin.Context) {
    var buffer bytes.Buffer
    for i := 0; i < repeat; i++ {
	// indicate which number "Hello!" is being printed
        buffer.WriteString(fmt.Sprintf("Hello! #%d\n", i+1))
    }
    c.String(http.StatusOK, buffer.String())
}

func main() {
    var err error
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    tStr := os.Getenv("REPEAT")
    repeat, err = strconv.Atoi(tStr)
    if err != nil {
        log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
        repeat = 5
    }

    router := gin.New()
    router.Use(gin.Logger())
    router.LoadHTMLGlob("templates/*.tmpl.html")
    router.Static("/static", "static")
	
    // directory listing and descriptions
    router.GET("/", func(c *gin.Context) {
        var buffer bytes.Buffer
	buffer.WriteString("Dir Listing:\n")
	buffer.WriteString("---------------\n")
	buffer.WriteString("/repeat\t\t\tPrints a message a certain amount of times.\n")
	buffer.WriteString("/randRepeat\t\tPrints a message a random number of times.\n")
	buffer.WriteString("/redirect\t\tRedirects to /repeat.\n")

	c.String(http.StatusOK, buffer.String())
    })

    router.GET("/mark", func(c *gin.Context) {
        c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
    })

    router.GET("/repeat", repeatHandler)



    router.GET("/randRepeat", func(c *gin.Context){
	
	randInt := rand.Intn(1000) // generate a random int
	
	var buffer bytes.Buffer

	// print "Hello!" randInt times
	for i := 1; i <= randInt; i++ {
		buffer.WriteString("Hello!\t")
	}
	c.String(http.StatusOK, buffer.String())
    })



    router.GET("/redirect", func(c *gin.Context){
	c.Redirect(307, "repeat")    // send a redirection code and redirect to /repeat
    })

    router.POST("/testPOST", func(c *gin.Context){
    	log.Printf("C: %#v", c)
    })
    
    router.Run(":" + port)
}
