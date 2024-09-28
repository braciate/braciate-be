package assetsHandler

import (
	assetsService "github.com/braciate/braciate-be/internal/api/assets/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AssetsHandler struct {
	assetsService assetsService.AssetsServiceItf
	log           *logrus.Logger
	validate      *validator.Validate
}

func New(log *logrus.Logger, AssetsService assetsService.AssetsServiceItf, validate *validator.Validate) *AssetsHandler {
	return &AssetsHandler{
		assetsService: AssetsService,
		log:           log,
		validate:      validate,
	}
}

func (h *AssetsHandler) Start(srv fiber.Router) {
	Assets := srv.Group("/assets")
	Assets.Post("/create", h.CreateAssetsHandler)
	Assets.Get("/get/:id", h.GetAllAssetsByNomination)
	Assets.Delete("delete/:id", h.DeleteAssets)

}
