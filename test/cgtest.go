package main

import (
	"github.com/Anderson-Lu/gofasion/gofasion"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"io/ioutil"
	"log"
	"suyuan/util"
)

//r.POST("/cgtest", test.Cgtest)
//r.GET("/cgtest", test.Gettest)
//r.PUT("/cgtest/:id/:name", test.PutTest)
//r.DELETE("/cgtest/:id", test.DeleteTest)

func Cgtest(c *gin.Context){


	dataByte, _ := ioutil.ReadAll(c.Request.Body)

	fsion := gofasion.NewFasion(string(dataByte))

	username := fsion.Get("username").ValueStr()
	password := fsion.Get("password").ValueInt()
	//password := com.StrTo(fsion.Get("password").ValueStr()).MustInt()
	log.Println("username==",username)
	log.Println("password==",password)
}

func Gettest(c *gin.Context){

	ids := c.Query("id") //string
	log.Println("ids==",ids)
	id := com.StrTo(c.Query("id")).MustInt()
	log.Println("id==",id)

}

func PutTest(c *gin.Context){

	id := c.Param("id")
	name := c.Param("name")
	log.Println("id==",id," name==",name)
}

func DeleteTest(c *gin.Context){
	id := c.Param("id")
	log.Println("id==",id)
}

func main() {

	aaa := []string{"aa","bb","cc","aa","dd"}
	bbb :=util.RemoveRepByMap(aaa)
	log.Println("bbb==",bbb)
}