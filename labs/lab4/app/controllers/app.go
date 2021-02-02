package controllers

import (
	"fmt"
	"lab4/app/stores"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) View() revel.Result {
	website := stores.Load()
	fmt.Println(website)

	return c.Render(website)
}

func (c App) Edit(title string) revel.Result {
	website := stores.Load()

	if c.Request.Method == "POST" {
		c.Validation.Required(title)

		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(App.Edit)
		}

		website.Title = title
		if err := website.Save(); err != nil {
			c.Flash.Error(err.Error())
			c.FlashParams()
			return c.Redirect(App.Edit)
		}

		return c.Redirect(App.View)
	}

	return c.Render(website)
}

func (c App) Save(body string) revel.Result {
	website := stores.Load()

	if c.Request.Method == "POST" {
		c.Validation.Required(body)

		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(App.Save)
		}

		website.Body = body
		if err := website.Save(); err != nil {
			c.Flash.Error(err.Error())
			c.FlashParams()
			return c.Redirect(App.Edit)
		}
		return c.Redirect(App.View)
	}

	return c.Render(website)
}
