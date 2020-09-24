package controllers

import (
	"BeegoHello01/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}
/*
	func (c *MainController) Get() {

		//获取请求数据
		userName := c.Ctx.Input.Query("user")
		password := c.Ctx.Input.Query("psd")
		//使用固定数据进行数据校验
		if userName != "zhanghuayang" || password != "425680"{
			c.Ctx.ResponseWriter.Write([]byte("数据错误"))
			return
		}
		c.Ctx.ResponseWriter.Write([]byte("数据校验成功"))


		c.Data["Website"] = "zhanghuayang"
		c.Data["Email"] = "http://www.baidu.com"
		c.TplName = "index.tpl"
	}
	//post方法请求
	func (c *MainController)Post(){

		//接收post请求的参数
		userName := c.Ctx.Request.FormValue("user")
		password := c.Ctx.Request.FormValue("psd")
		fmt.Println(userName,password)

		//数据校验
		if userName != "zhanghuayang" && password != "425680"{
			c.Ctx.WriteString("数据请求失败")
			return
		}
		c.Ctx.WriteString("校验成功")

	}

 */
func (c *MainController)  Post(){
	var person models.Person
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		c.Ctx.WriteString("接收失败")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil{
		c.Ctx.WriteString("解析失败")
		return
	}
	fmt.Println("姓名：",person.Name)
	fmt.Println(("年龄："),person.Age)
	fmt.Println("性别：",person.Sex)
	c.Ctx.WriteString("数据解析成功")
}