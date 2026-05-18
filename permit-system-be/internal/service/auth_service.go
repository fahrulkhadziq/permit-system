package service

import (
	"errors"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/model"
	"permit-license/internal/repository"
	"time"
)

type AuthService struct {
	RepoAuth    repository.AuthRepository
	RepoUser    repository.UserRepository
	RefreshRepo repository.RefreshTokenRepository
}

func (s *AuthService) Register(req dto.RegisterRequest) error {
	hash, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := model.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		UnitID:   &req.UnitID,
	}

	err = s.RepoAuth.CreateUser(&user)
	if err != nil {
		return err
	}

	userRole := model.UserRole{
		UserID: user.ID.String(),
		RoleID: req.RoleID,
	}
	return s.RepoUser.CreateUserRole(&userRole)
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.RepoUser.FindUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword(
		user.Password,
		req.Password,
	)
	if err != nil {
		return nil, err
	}

	userRole, err := s.RepoUser.GetUserRole(user.ID.String())
	if err != nil {
		return nil, err
	}

	accessToken, err :=
		helper.GenerateAccessToken(
			user.ID.String(),
			userRole.Role.Name,
			*user.UnitID,
		)

	if err != nil {
		return nil, err
	}

	refreshToken, err :=
		helper.GenerateRefreshToken(
			user.ID.String(),
		)

	if err != nil {
		return nil, err
	}

	refresh := model.RefreshToken{
		UserID: user.ID.String(),
		Token:  refreshToken,
		ExpiredAt: time.Now().
			Add(time.Hour * 24 * 7),
	}

	err = s.RefreshRepo.Create(&refresh)
	if err != nil {
		return nil, err
	}

	response := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.LoginUserResponse{
			ID:       user.ID.String(),
			FullName: user.FullName,
			Email:    user.Email,
			Role:     userRole.Role.Name,
			UnitID:   user.UnitID,
		},
	}

	return &response, nil
}

func (s *AuthService) RefreshToken(token string) (string, error) {

	data, err :=
		s.RefreshRepo.FindByToken(token)

	if err != nil {
		return "", err
	}

	if data.IsRevoked {

		return "", errors.New(
			"token revoked",
		)
	}

	if time.Now().After(data.ExpiredAt) {

		return "", errors.New(
			"token expired",
		)
	}

	userRole, err :=
		s.RepoUser.GetUserRole(
			data.User.ID.String(),
		)

	if err != nil {
		return "", err
	}

	return helper.GenerateAccessToken(
		data.User.ID.String(),
		userRole.Role.Name,
		*data.User.UnitID,
	)
}

func (s *AuthService) Logout(refreshToken string) error {

	return s.RefreshRepo.Revoke(refreshToken)
}
