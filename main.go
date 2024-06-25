package main

import (
	"github.com/kataras/iris/v12"
)

type HashRequest struct {
	Content string `form:"content"`
	Digest  string `form:"digest"`
}

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

	app.Post("/hash", func(ctx iris.Context) {
		var req HashRequest
		if err := ctx.ReadForm(&req); err != nil {
			app.Logger().Error(err.Error())
			return
		}

		// call service

		data := iris.Map{
			"Title":     "Hash",
			"MD5":       "",
			"Sha1":      "",
			"Sha256":    "",
			"Sha224":    "",
			"Sha512":    "",
			"Sha384":    "",
			"Sha3":      "",
			"RIPEMD160": "",
		}

		ctx.ViewLayout("layout/index.html")

		if err := ctx.View("hash/hmac.html", data); err != nil {
			ctx.HTML("<h3>%s</h3>", err.Error())
			return
		}

	})

	app.Get("/hash", func(ctx iris.Context) {

		data := iris.Map{
			"Title":     "Hash",
			"MD5":       "",
			"Sha1":      "",
			"Sha256":    "",
			"Sha224":    "",
			"Sha512":    "",
			"Sha384":    "",
			"Sha3":      "",
			"RIPEMD160": "",
		}

		ctx.ViewLayout("layout/index.html")

		if err := ctx.View("hash/hmac.html", data); err != nil {
			ctx.HTML("<h3>%s</h3>", err.Error())
			return
		}

	})

	app.Listen(":8877")
}
