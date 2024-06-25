package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	// do stuff here
	app := iris.New()

	tmpl := iris.HTML("./views", ".html")
	tmpl.AddFunc("greet", func(s string) string {
		return "Greeting" + s + "!"
	})
	tmpl.Reload(true)

	app.RegisterView(tmpl)

	// register assets
	app.HandleDir("/assets", iris.Dir("./assets"))

	app.Get("/", func(ctx iris.Context) {

		data := iris.Map{
			"Title":   "Hmac",
			"Message": "Hello hmac",
		}

		ctx.ViewLayout("layout/index.html")

		if err := ctx.View("hash/hmac.html", data); err != nil {
			ctx.HTML("<h3>%s</h3>", err.Error())
			return
		}

	})

	app.Listen(":8877")
}
