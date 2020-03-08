package api

import (
	"fmt"
	"src/first_work"
	"testing"
)

func TestLogin(t *testing.T) {
	var u User
	table := first_work.DB.CreateTable(u)
	fmt.Println(table)

}
