//struct and methods

package service

import (
	"context"
	"errors"
	"log"
	"project/models"
	"project/repo"
	"regexp"
)

type UserServiceInterface interface {
	//user methods
	Login(authdetails models.Authentication) (validToken string, err error)
	ListUsers() (users []models.UserList, err error)
	CreateUser(ctx context.Context, user models.CreateUser) (createduser models.CreateUser, err error)
	GetUser(ctx context.Context, variable interface{}) (user models.User, err error)
	//UpdateUser(ctx context.Context, email string, user models.User) (updateUser models.UpdateUser, err error)
	//GetUserByEmail(email string) (user models.User, err error)
	DeleteUser(uid int) (selectuser models.DeleteUser, bookReportLists []models.BookReportList, err error)
	// login(authdetails models.Authentication) (err error)
	UpdateUserWithID(ctx context.Context, uid int, user models.User) (updateUser models.UpdateUser, err error)
	GetUserByID(uid int) (user models.GetUser, err error)
}

type userService struct {
	repo     repo.UserRepoInterface
	gentoken AuthTokenInterface
}

func InitUserService(r repo.UserRepoInterface, at AuthTokenInterface) UserServiceInterface {

	return &userService{
		repo:     r,
		gentoken: at,
	}
}

func (us *userService) ListUsers() (users []models.UserList, err error) {
	//val, ok := ctx.Value("ClaimsToVerify").(*models.Claims)
	users, err = us.repo.ListUser()
	if err != nil {
		return
	}
	return
}

func (us *userService) Login(authdetails models.Authentication) (validToken string, err error) {
	user, err := us.repo.GetUserByEmail(authdetails.Email)

	if user.Password == authdetails.Password {
		validToken, err = us.gentoken.GenerateToken(user.User_ID, user.Email, user.Role)
		return
	} else {
		log.Println(err)
		err = errors.New("login failed. please check credantials")
		return
	}
}

func (us *userService) CreateUser(ctx context.Context, user models.CreateUser) (createduser models.CreateUser, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
	cratedByEmail := val.Email
	updatedByEmail := val.Email

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isvalid := emailRegex.MatchString(user.Email)

	if isvalid {
		if user.Role == "user" || user.Role == "admin" || user.Role == "superadmin" {
			createduser, err = us.repo.CreateUser(cratedByEmail, updatedByEmail, user)
			if createduser.Role == "3" {
				createduser.Role = "user"
			} else if createduser.Role == "2" {
				createduser.Role = "admin"
			} else if createduser.Role == "1" {
				createduser.Role = "superadmin"
			}
			if err != nil {
				return
			}
		} else {
			err = errors.New("role must be superadmin,admin,user")
		}

	} else {
		err = errors.New("email address format is incorrect")
	}

	return
}

func (us *userService) GetUserByID(uid int) (user models.GetUser, err error) {
	user, err = us.repo.GetUserByID(uid)
	if user.Role == "3" {
		user.Role = "user"
	} else if user.Role == "2" {
		user.Role = "admin"
	} else if user.Role == "1" {
		user.Role = "superadmin"
	}
	if err != nil {
		return
	}
	return
}

func (us *userService) GetUser(ctx context.Context, variable interface{}) (user models.User, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)

	if val.Email == variable || val.UserID == variable || val.Role == "admin" || val.Role == "superadmin" {
		user, err = us.repo.GetUser(variable)
		if err != nil {
			return
		}
	} else {
		err = errors.New("you are unauthorized person")
		return
	}

	return
}

func (us *userService) UpdateUserWithID(ctx context.Context, uid int, user models.User) (updateUser models.UpdateUser, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
	updatedBy := val.Email
	if val.UserID == uid || val.Role == "admin" || val.Role == "superadmin" {
		updateUser, err = us.repo.UpdateUserWithID(uid, user, updatedBy)
		if err != nil {
			return
		}
	} else {
		err = errors.New("you are unauthorized person")
		return
	}
	return
}

// func (us *userService) UpdateUser(ctx context.Context, email string, user models.User) (updateUser models.UpdateUser, err error) {
// 	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
// 	if val.Email == email || val.Role == "admin" || val.Role == "superadmin" {
// 		updateUser, err = us.repo.UpdateUser(email, user)
// 		if err != nil {
// 			return
// 		}
// 	} else {
// 		err = errors.New("you are unauthorized person")
// 		return
// 	}
// 	return
// }

// func (us *userService) GetUserByEmail(email string) (user models.User, err error) {
// 	user, err = us.repo.GetUserByEmail(email)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func (us *userService) DeleteUser(uid int) (selectuser models.DeleteUser, bookReportLists []models.BookReportList, err error) {
	selectuser, bookReportLists, err = us.repo.DeleteUser(uid)
	if err != nil {
		return
	}
	return
}
