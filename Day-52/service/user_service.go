package service

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"day52/domain"
)

type UserService struct {
	repo      domain.UserRepository
	jwtSecret []byte
}

type JWTClaims struct {
	UserID string `json:"sub"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

func NewUserService(repo domain.UserRepository, secret string) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: []byte(secret),
	}
}

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password + "day52_salt"))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *UserService) GenerateJWT(user *domain.User) (string, error) {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))

	claims := JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Exp:    time.Now().Add(24 * time.Hour).Unix(),
	}
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payload := base64.RawURLEncoding.EncodeToString(claimsBytes)

	unsignedToken := header + "." + payload
	h := hmac.New(sha256.New, s.jwtSecret)
	h.Write([]byte(unsignedToken))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return unsignedToken + "." + signature, nil
}

func (s *UserService) ValidateJWT(tokenStr string) (*JWTClaims, error) {
	parts := strings.Split(tokenStr, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	unsignedToken := parts[0] + "." + parts[1]
	h := hmac.New(sha256.New, s.jwtSecret)
	h.Write([]byte(unsignedToken))
	expectedSig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	if !hmac.Equal([]byte(parts[2]), []byte(expectedSig)) {
		return nil, fmt.Errorf("invalid token signature")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	var claims JWTClaims
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, fmt.Errorf("failed to unmarshal claims: %w", err)
	}

	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("token has expired")
	}

	return &claims, nil
}

func (s *UserService) Register(req domain.RegisterRequest) (*domain.AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("email and password required")
	}

	idBytes := make([]byte, 8)
	_, _ = rand.Read(idBytes)
	userID := "usr_" + hex.EncodeToString(idBytes)

	user := &domain.User{
		ID:        userID,
		Email:     req.Email,
		Name:      req.Name,
		Password:  hashPassword(req.Password),
		Role:      "user",
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	token, err := s.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *UserService) Login(req domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, domain.ErrInvalidPassword
	}

	if user.Password != hashPassword(req.Password) {
		return nil, domain.ErrInvalidPassword
	}

	token, err := s.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *UserService) GetProfile(userID string) (*domain.User, error) {
	return s.repo.FindByID(userID)
}
