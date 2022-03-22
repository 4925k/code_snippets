package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqBody struct {
	Client string `json:"client"`
	Token  string `json:"token"`
}

func main() {
	r := gin.Default()
	r.Use(audit)
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))
	r.POST("/content", content)
	r.GET("/ping/pong", pong)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func pong(c *gin.Context) {
	//fmt.Println(c.FullPath())
	//fmt.Println(c.Request.RequestURI)
	c.JSON(200, gin.H{
		"message": "pong",
	})
	fmt.Println("asdsad")
}

func content(c *gin.Context) {
	var newReq ReqBody
	if err := c.BindJSON(&newReq); err != nil {
		return
	}

	c.IndentedJSON(http.StatusAccepted, newReq.Token)

}

func audit(c *gin.Context) {
	var newReq ReqBody
	if err := c.BindJSON(&newReq); err != nil {
		return
	}

	content, _ := json.Marshal(newReq)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(content))

	fmt.Printf("endpoint:%s, client:%s, token:%s\n", c.Request.RequestURI, newReq.Client, newReq.Token)

	// file := fmt.Sprintf("./audit/%s.audit", newReq.Client)
	// log := fmt.Sprintf("%s req from %s\n", c.Request.RequestURI, newReq.Token)
	// os.WriteFile(file, []byte(log), 0666)
	c.Next()
}
