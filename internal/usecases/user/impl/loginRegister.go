package impl

import (
	"base/internal/inbound/model"
	dbModel "base/internal/repository/db/model"
	uModel "base/internal/usecases/model"
	"context"
	"errors"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Register registers a new user in the system
func (u *usecase) Register(ctx context.Context, username string, email string, password string) (users model.Users, err error) {
	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.Users{}, err
	}

	user := dbModel.User{
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		CreatedBy: username,
		// CreatedDate: time.now(),
		UpdatedBy: username,
		// UpdatedDate: time.now(),
	}

	// Save user in the repository
	_, err = u.User.CreateUser(ctx, user)
	if err != nil {
		return model.Users{}, err
	}

	tokenStr, err := u.Ujwt.GenerateToken(ctx, &uModel.DataJwt{
		Id:       strconv.Itoa(user.ID),
		Username: user.Username,
		Email:    user.Email,
	})

	users = model.Users{
		Username: username,
		Email:    email,
		UserId:   user.ID,
		Jwt:      &tokenStr,
	}
	return
}

func (u *usecase) Login(ctx context.Context, usernameOrEmail string, password string) (user model.Users, err error) {
	// Check if usernameOrEmail is email or username
	var userData dbModel.User
	if u.isEmail(usernameOrEmail) {
		userData, err = u.User.GetUserByEmail(ctx, usernameOrEmail)
	} else {
		userData, err = u.User.GetUserByUsername(ctx, usernameOrEmail)
	}

	if err != nil {
		return model.Users{}, errors.New("invalid credentials")
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return model.Users{}, errors.New("invalid credentials")
	}

	user = model.Users{
		Username: userData.Username,
		Email:    userData.Email,
		UserId:   userData.ID,
	}

	return user, nil
}

// Helper function to check if a string is an email address
func (u *usecase) isEmail(input string) bool {
	// Simple email check
	if strings.Contains(input, "@") && strings.Contains(input, ".") {
		return true
	}
	return false
}
