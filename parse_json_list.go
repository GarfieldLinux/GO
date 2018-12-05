package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateParams struct {
	Username     string `json:"username"`
	Guests       Guests `json:"guests"`
	RoomType     string `json:"roomType"`
	CheckinDate  string `json:"checkinDate"`
	CheckoutDate string `json:"checkoutDate"`
}

type Guests struct {
	Person []Person `json:"person"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		var f CreateParams
		if err := c.BindJSON(&f); err != nil {
			return
		}
		c.IndentedJSON(http.StatusOK, f)
	})
	r.Run(":4000")
}


//for test
run the server

go run /tmp/test.go

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /                         --> main.main.func1 (1 handlers)
[GIN-debug] Listening and serving HTTP on :4000
access from curl

curl 0:4000 -X POST -d '{"username":"foo","guests":{"person":[{"firstname":"foobar","lastname":"barfoo"},{"firstname":"foofoo","lastname":"barbar"}]}}'

{
    "username": "foo",
    "guests": {
        "person": [
            {
                "firstname": "foobar",
                "lastname": "barfoo"
            },
            {
                "firstname": "foofoo",
                "lastname": "barbar"
            }
        ]
    },
    "roomType": "",
    "checkinDate": "",
    "checkoutDate": ""
}
