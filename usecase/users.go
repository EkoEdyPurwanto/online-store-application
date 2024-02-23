package usecase

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"online-store-application/model"
	"online-store-application/model/dto/req"
	"online-store-application/repository"
	"online-store-application/utils/common"
	"online-store-application/utils/security"
	"time"
)

type (
	UsersUseCase interface {
		Register(payload req.RegisterRequest) error
	}

	usersUseCase struct {
		repo repository.UsersRepository
		walletsUC WalletsUseCase
	}
)

func (u *usersUseCase) Register(payload req.RegisterRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	users := model.Users{
		Id:          common.GenerateID(),
		Username:    payload.UserName,
		Password:    hashPassword,
		Email:       payload.Identifier.Email,
		PhoneNumber: payload.Identifier.PhoneNumber,
		UserStatus:  model.Active,
		Role:        payload.Role,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err = u.repo.Save(users)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}

	wallets := model.Wallets{
		Id:           common.GenerateID(),
		UserId:       users.Id,
		RekeningUser: common.GenerateRandomRekeningNumber(10),
		Balance:      0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Time{},
	}

	err = u.walletsUC.CreateWallet(wallets)
	if err != nil {
		return fmt.Errorf("failed create wallet: %v", err.Error())
	}

	return nil
}

// Constructor
func NewUsersUseCase(repo repository.UsersRepository, walletsUC WalletsUseCase) UsersUseCase {
	return &usersUseCase{
		repo: repo,
		walletsUC: walletsUC,
	}
}
