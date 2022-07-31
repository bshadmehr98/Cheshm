package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bshadmehr98/cheshm/libs"
)

type SignupController struct {
	beego.Controller
}

type UserSignupForm struct {
	FirstName string `valid:"Required;MinSize(2)"`
	LastName  string `valid:"Required;MinSize(2)"`
	Email     string `valid:"Required;Email;MaxSize(100)"`
	Password  string `valid:"Required;MinSize(8)"`
}

func (c *SignupController) Get() {
	c.Data["Form"] = &UserSignupForm{}
	c.TplName = "signup.html"
}

func (c *SignupController) Post() {
	u := UserSignupForm{}

	if err := c.ParseForm(&u); err != nil {
		c.Abort("500")
	}

	isValid, errors, err := libs.CreateValidationError(u)
	if !isValid {
		if err != nil {
			c.Abort("500")
		}
		c.Data["Errors"] = errors
	}

	c.TplName = "signup.html"
	c.Data["Form"] = &u

}
