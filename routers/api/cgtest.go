package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"suyuan/g"
	"suyuan/models"
	"suyuan/util"
)

type auth struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     []int  `json:"role_id"`
}

func AddUsers(c *gin.Context) {

	appG := g.Gin{C: c}
	var reqInfo auth
	err := c.BindJSON(&reqInfo)

	if err != nil {
		appG.Response(http.StatusBadRequest, g.INVALID_PARAMS, nil)
		return
	}

	valid := validation.Validation{}
	valid.MaxSize(reqInfo.Username, 100, "username").Message("最长为100字符")
	valid.MaxSize(reqInfo.Password, 100, "password").Message("最长为100字符")

	if valid.HasErrors() {
		g.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, g.ERROR_ADD_FAIL, valid.Errors)
		return
	}

	check, _ := models.CheckUserUsername(reqInfo.Username)
	if !check {
		menu := map[string]interface{}{
			"username": reqInfo.Username,
			"password": util.EncodeMD5(reqInfo.Password),
			"role_id":  reqInfo.Role,
		}
		if id, err := models.AddUser(menu); err == nil {
			log.Println(id, err)
		} else {
			log.Println(0, err)
		}

	}
}
