package controllers

import (
	"go-fiber-api/services"
	"go-fiber-api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MessageResponse struct {
	Message string
}

func GetHealthCheck(ctx *fiber.Ctx) error {
	msg := MessageResponse{
		Message: "we up",
	}

	return ctx.JSON(msg)
}

func GetServiceCheck(ctx *fiber.Ctx) error {
	res := MessageResponse{}
	res.Message = services.SayHello()

	return ctx.JSON(res)
}

func GetSW(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"sw": "star wars endpoint up",
	})
}

func GetPeople(ctx *fiber.Ctx) error {
	res := services.GetPeople()

	return ctx.JSON(res)
}

func GetPersonRandom(ctx *fiber.Ctx) error {
	res := services.GetPerson(strconv.Itoa(utils.CreateRandomNumber()))

	return ctx.JSON(res)
}

func GetPerson(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	res := services.GetPerson(id)

	return ctx.JSON(res)
}
