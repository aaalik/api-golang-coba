package v1

import (
	"strconv"

	"github.com/aaalik/coba-golang/model"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/v12"
)

func List(ctx iris.Context) {
	items := model.GetItem()

	ctx.JSON(context.Map{"response": items})
}

func Item(ctx iris.Context) {
	id := ctx.Params().Get("id")

	tmpId, err := strconv.Atoi(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	items, err := model.GetSingleItem(tmpId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	ctx.JSON(context.Map{"response": items})
}

func SaveItem(ctx iris.Context) {
	name := ctx.PostValue("name")
	price := ctx.PostValue("price")

	intPrice, err := strconv.Atoi(price)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	items, err := model.SaveItem(name, intPrice)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(context.Map{"response": items})
}

func EditItem(ctx iris.Context) {
	id := ctx.Params().Get("id")
	name := ctx.PostValue("name")
	price := ctx.PostValue("price")

	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	tmpItems, err := model.GetSingleItem(intId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	if name == "" {
		name = tmpItems.Name
	}

	if price == "" {
		price = strconv.Itoa(tmpItems.Price)
	}

	intPrice, err := strconv.Atoi(price)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	items, err := model.UpdateItem(intId, name, intPrice)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	ctx.JSON(context.Map{"response": items})
}

func DeleteItem(ctx iris.Context) {
	id := ctx.Params().Get("id")

	tmpId, err := strconv.Atoi(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	items, err := model.DeleteItem(tmpId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(context.Map{"response": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
	ctx.JSON(context.Map{"response": items})
}
