package api

import (
	"fmt"
	. "future/model"
	"log"
	"net/http"
	"strings"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
    b, err := GenerateRandomBytes(s)
    return base64.URLEncoding.EncodeToString(b), err
}

func IndexApi(c *gin.Context) {
	for i :=0 ;i<100;i++{
		token, err := GenerateRandomString(20)
		if err != nil {
			// Serve an appropriately vague error to the
			// user, but log the details internally.
			fmt.Println(err)
		}

		c.String(http.StatusOK, token+"\n")
	}
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("uid")
	lastName := c.Request.FormValue("token")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg":       msg,
		"firstName": firstName,
		"lastName":  lastName,
	})
}

//test redis key查询有效时间
func GetRedisValTimeAPI(c *gin.Context) {
	key := c.Param("key")
	val := GetRedisVTime(key)
	var output=""
	if key == "rjc"{
		newval := strings.Split(val,"~")
		for _,v := range newval {
			output+=v+"\n"
		}
		c.String(http.StatusOK,output)
		return
	}
	c.String(http.StatusOK,val)
}

//test redis key 的值
func GetRedisValAPI(c *gin.Context) {
	key := c.Param("key")
	val := GetRedisV(key)
	var output=""
	if key == "rjc"{
		newval := strings.Split(val,"~")
		for _,v := range newval {
			output+=v+"\n"
		}
		c.String(http.StatusOK,output)
		return
	}
	c.String(http.StatusOK,val)
}

func SetRedisValAPI(c *gin.Context){
	key := c.Query("key")
	val := c.Query("val")
	if err:=SetRedisV(key,val,60);!err {
		c.String(http.StatusOK,"false")
		return
	}
	c.String(http.StatusOK,"true")
}
