package main

import (
	"fmt"
	"third/rrf"
)

type Cook struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	a := rrf.Default()
	g := a.Group("/abc")
	a.GET("/book", GetBook)

	a.POST("/cook", PostCook)
	g.GET("/book", middle, GetBook)
	a.Run("8080")

}
func middle(c *rrf.Context) {
	fmt.Println("this is middle")
}
func GetBook(c *rrf.Context) {
	query := c.Query("id")
	c.Write("get book id=" + query)

}
func PostCook(c *rrf.Context) {
	b := &Cook{}

	err := c.BindJson(b)
	if err != nil {
		fmt.Println("json error:", err)
	}
	fmt.Println(b.Name, b.Id)

}
