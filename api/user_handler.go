package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kumarvinayak490/hotel-reservation/db"
)


type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler{
	return &UserHandler{
		userStore: userStore,
	}
}


func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error{
	ctx := context.Background()
	id:=c.Params("id")
	user,err :=h.userStore.GetUserByID(ctx, id)
	if err !=nil{
		return err
	}
	return c.JSON(user)
}