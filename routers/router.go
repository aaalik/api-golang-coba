package routers

import (
	conV1 "github.com/aaalik/coba-golang/controller/v1"
	"github.com/kataras/iris/v12"
)

func SetupRouter() {
	app := iris.Default()

	v1 := app.Party("/v1")
	{
		itemsAPI := v1.Party("/items")
		{
			itemsAPI.Get("/", conV1.List)
			itemsAPI.Get("/{id}", conV1.Item)
			itemsAPI.Post("/", conV1.SaveItem)
			itemsAPI.Patch("/{id}", conV1.EditItem)
			itemsAPI.Delete("/{id}", conV1.DeleteItem)
		}
	}

	app.Listen(":8080")
}
