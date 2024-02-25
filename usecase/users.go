package usecase

import (
	"errors"
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
		Login(payload req.LoginRequest) (string, error)
		GetById(payload string) (model.Users, error)
	}

	usersUseCase struct {
		repo      repository.UsersRepository
		walletsUC WalletsUseCase
	}
)

// Constructor
func NewUsersUseCase(repo repository.UsersRepository, walletsUC WalletsUseCase) UsersUseCase {
	return &usersUseCase{
		repo:      repo,
		walletsUC: walletsUC,
	}
}

func (u *usersUseCase) Register(payload req.RegisterRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
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
	}

	err = u.walletsUC.CreateWallet(wallets)
	if err != nil {
		return fmt.Errorf("failed create wallet: %v", err.Error())
	}

	return nil
}

func (u *usersUseCase) Login(payload req.LoginRequest) (string, error) {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return "", err
	}

	var identifier string

	if payload.Identifier.UserName != "" {
		identifier = payload.Identifier.UserName
	} else if payload.Identifier.Email != "" {
		identifier = payload.Identifier.Email
	} else if payload.Identifier.PhoneNumber != "" {
		identifier = payload.Identifier.PhoneNumber
	}

	if identifier == "" {
		return "", errors.New("invalid payload")
	}

	users, err := u.repo.FindUserByIdentifier(identifier)
	if err != nil {
		return "", err
	}

	// Validasi Password
	err = security.VerifyPassword(users.Password, payload.Password)
	if err != nil {
		return "", fmt.Errorf("unauthorized: invalid credential")
	}

	// Validasi active or not
	if users.UserStatus != "active" {
		return "", fmt.Errorf("your account is inactive")
	}

	// Generate Token
	token, err := security.GenerateJwtToken(users)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *usersUseCase) GetById(payload string) (model.Users, error) {
	byId, err := u.repo.FindById(payload)
	if err != nil {
		return model.Users{}, fmt.Errorf("user not found")
	}
	return byId, nil
}