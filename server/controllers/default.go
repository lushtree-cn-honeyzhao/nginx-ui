package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type RespData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
	beego.Controller
	jsonData *RespData
	// json real data
	respData map[string]any
}

func (c *BaseController) json() {
	c.checkJsonData()
	if c.respData != nil {
		c.jsonData.Data = c.respData
	}
	c.Data["json"] = c.jsonData
	c.ServeJSON()
}

func (c *BaseController) ErrorJson(error error) {
	c.setCode(-1).setMsg(error.Error()).json()
}

func (c *BaseController) checkJsonData() *BaseController {
	data := c.jsonData
	if data == nil {
		data = &RespData{
			Code: 0,
			Msg:  "success",
			Data: nil,
		}
		c.jsonData = data
	}
	return c
}

func (c *BaseController) setData(v interface{}) *BaseController {
	c.checkJsonData()
	c.jsonData.Data = v
	return c
}

func (c *BaseController) addRespData(k string, v interface{}) *BaseController {
	c.checkJsonData()
	if c.respData == nil {
		c.respData = map[string]any{}
	}
	c.respData[k] = v
	return c
}

func (c *BaseController) setCode(code int) *BaseController {
	c.checkJsonData()
	c.jsonData.Code = code
	return c
}
func (c *BaseController) setMsg(msg string) *BaseController {
	c.checkJsonData()
	c.jsonData.Msg = msg
	return c
}

func (c *BaseController) getParam(k string) string {
	return c.Ctx.Input.Param(k)
}
func (c *BaseController) getIntParam(k string) (int, error) {
	idStr := c.Ctx.Input.Param(k)
	id, err := strconv.Atoi(idStr)
	logs.Info("id", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
