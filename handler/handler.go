package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Number     int       `form:"number" binding:"gte=6"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func StartPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if err := c.ShouldBind(&person); err == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	} else {
		log.Fatal(err)
	}

	c.JSON(200, person)
}
