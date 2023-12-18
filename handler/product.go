package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abgeo/fx-workshop/model"
	"github.com/abgeo/fx-workshop/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler struct {
	logger      *zap.Logger
	productRepo *repository.ProductRepository
}

func NewProductHandler() (*ProductHandler, error) {
	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	productRepo, err := repository.NewProductRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize product repository: %w", err)
	}

	return &ProductHandler{
		logger:      log,
		productRepo: productRepo,
	}, nil
}

func (handler *ProductHandler) Create(ctx *gin.Context) {
	entity := model.Product{}

	if err := ctx.ShouldBind(&entity); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	handler.logger.Info(
		"Creating new product",
		zap.String("code", entity.Code),
	)

	if err := handler.productRepo.Create(&entity); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, entity)
}

func (handler *ProductHandler) GetAll(ctx *gin.Context) {
	websites, err := handler.productRepo.FindAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, websites)
}

func (handler *ProductHandler) GetSingle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})

		return
	}

	website, err := handler.productRepo.FindByID(uint(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, website)
}
