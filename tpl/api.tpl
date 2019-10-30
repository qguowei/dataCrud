package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"newjxc/{{.SevPackageName}}"
	"newjxc/utils"
	"strconv"
)

func Get{{.ModelName}}(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusOK, utils.JsonError(utils.CodeInvalidData, "参数错误"))
        return
    }
	in := {{.SevPackageName}}.{{.ModelName}}InfoInput{
		CusSN:  c.MustGet("cus_sn").(string),
		ID:    int32(id),
	}
	info, err := {{.SevPackageName}}.{{.ModelName}}Info(&in)
	if err != nil {
		code, msg := getApiCode(err)
		c.JSON(http.StatusOK, utils.JsonError(code, msg))
		return
	}
	c.JSON(http.StatusOK, utils.JsonSuccess(info))
}

func Get{{.ModelName}}List(c *gin.Context) {
	where := utils.ParseKeywords(c, map[string][]string{
		"like": {"user_name", "number"},
		"eq":   {"dep_id"},
	})
	page, limit := utils.ParseLimit(c, 10)
	in := {{.SevPackageName}}.{{.ModelName}}ListInput{
		CusSN: c.MustGet("cus_sn").(string),
		Page:  page,
		Limit: limit,
		Where: where,
	}
	data, err := {{.SevPackageName}}.{{.ModelName}}List(&in)
	if err != nil {
		code, msg := getApiCode(err)
		c.JSON(http.StatusOK, utils.JsonError(code, msg))
		return
	}
	c.JSON(http.StatusOK, utils.JsonSuccess(data))
}

func Add{{.ModelName}}(c *gin.Context) {
	var in {{.SevPackageName}}.{{.ModelName}}CreateInput
	if err := c.BindJSON(&in); err != nil {
	    log.Println("Add{{.ModelName}} is err:",err)
		c.JSON(http.StatusOK, utils.JsonError(utils.CodeInvalidData, "参数错误"))
		return
	}
	in.CusSN = c.MustGet("cus_sn").(string)
	info, err := {{.SevPackageName}}.{{.ModelName}}Create(&in)
	if err != nil {
		code, msg := getApiCode(err)
		c.JSON(http.StatusOK, utils.JsonError(code, msg))
		return
	}
	c.JSON(http.StatusOK, utils.JsonSuccess(info))
}

func Edit{{.ModelName}}(c *gin.Context) {
	var in {{.SevPackageName}}.{{.ModelName}}UpdateInput
	if err := c.BindJSON(&in); err != nil {
	    log.Println("Edit{{.ModelName}} is err:",err)
		c.JSON(http.StatusOK, utils.JsonError(utils.CodeInvalidData, "参数错误"))
		return
	}
	in.CusSN = c.MustGet("cus_sn").(string)
	info, err := {{.SevPackageName}}.{{.ModelName}}Update(&in)
	if err != nil {
		code, msg := getApiCode(err)
		c.JSON(http.StatusOK, utils.JsonError(code, msg))
		return
	}
	c.JSON(http.StatusOK, utils.JsonSuccess(info))
}

func Del{{.ModelName}}(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusOK, utils.JsonError(utils.CodeInvalidData, "参数错误"))
        return
    }
	cusSn := c.MustGet("cus_sn").(string)
	_, err = {{.SevPackageName}}.{{.ModelName}}Delete(&{{.SevPackageName}}.{{.ModelName}}InfoInput{
		CusSN:  cusSn,
		ID:    int32(id),
	})
	if err != nil {
		code, msg := getApiCode(err)
		c.JSON(http.StatusOK, utils.JsonError(code, msg))
		return
	}
	c.JSON(http.StatusOK, utils.JsonSuccess("ok"))
}
