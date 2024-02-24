package usecase

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"online-store-application/model"
	"online-store-application/model/dto/req"
	"online-store-application/repository"
	"online-store-application/utils/common"
	"time"
)

type (
	CategoriesUseCase interface {
		CreateCategories(payload req.CreateCategoriesRequest) error
	}

	categoriesUseCase struct {
		repo repository.CategoriesRepository
	}
)

func (c *categoriesUseCase) CreateCategories(payload req.CreateCategoriesRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}
	//
	//var users model.Users
	//if users.Role != model.Admin {
	//	return fmt.Errorf("only admin users can create categories")
	//}

	categories := model.Categories{
		Id:        common.GenerateID(),
		Type:      payload.Type,
		CreatedAt: time.Now(),
	}

	err = c.repo.Save(categories)
	if err != nil {
		return fmt.Errorf("failed save category: %v", err.Error())
	}

	return nil
}

func NewCategoriesUseCase(repo repository.CategoriesRepository) CategoriesUseCase {
	return &categoriesUseCase{
		repo: repo,
	}
}
