package routes

import (
	model "Go-Proj/model"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)
type Repository struct {
	DB *gorm.DB
}
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/order/:id", r.GetOrderByID)
	api.Get("/orders", r.GetOrders)
}

func (r *Repository) GetOrderByID(context *fiber.Ctx) error {

	id := context.Params("id")
	order := &model.Order{}
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get order"})
		return nil
	}
	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(order).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the order"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "order id fetched successfully",
		"data":    order,
	})
	return nil
}


func (r *Repository) GetOrders(context *fiber.Ctx) error {
	order := &[]model.Order{}

	err := r.DB.Find(order).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get order"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "orders fetched successfully",
		"data":    order,
	})
	return nil
}