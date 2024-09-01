package categoriesHandler

import (
	categoriesService "github.com/braciate/braciate-be/internal/api/categories/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	categoryService categoriesService.CategoriesServiceItf
	log             *logrus.Logger
	validate        *validator.Validate
}

func New(log *logrus.Logger, categoryService categoriesService.CategoriesServiceItf, validate *validator.Validate) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
		log:             log,
		validate:        validate,
	}
}

func (h *CategoryHandler) Start(srv fiber.Router) {
	categorries := srv.Group("/categories")
	categorries.Post("/create", h.CreateCategoryHandler)
}
