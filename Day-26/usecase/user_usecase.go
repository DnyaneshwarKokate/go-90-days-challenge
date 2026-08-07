package usecase

import (
	"day-26/domain"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase injects UserRepository dependency into UseCase
func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (u *userUseCase) RegisterUser(name, email, role string) (*domain.User, error) {
	// 1. Business Logic: Validate required fields
	if name == "" || email == "" || role == "" {
		return nil, domain.ErrInvalidUserData
	}

	// 2. Business Logic: Check if user with email already exists
	existingUser, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	// 3. Create entity
	newUser := &domain.User{
		Name:  name,
		Email: email,
		Role:  role,
	}

	// 4. Persist via repository
	return u.userRepo.Create(newUser)
}

func (u *userUseCase) GetUserByID(id int) (*domain.User, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidUserData
	}
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) ListUsers() ([]domain.User, error) {
	return u.userRepo.GetAll()
}
