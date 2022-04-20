package services

import (
	"awesomeProject/api/models"
	"awesomeProject/db"
	"github.com/rs/xid"
)

func CreateUser(user *models.User, e string, p []byte) error {
	user.ID = xid.New().String()
	user.Email = e
	user.Password = p

	res := db.DB().Create(user)

	return res.Error
}

func GetUsers(users *[]models.User) error {
	res := db.DB().Select([]string{"id", "email", "password"}).Find(users)

	return res.Error
}

func GetUserById(user *models.User, id string) error {
	res := db.DB().First(user, id)

	return res.Error
}

func GetDeletedUser(user *models.User, id string) error {
	res := db.DB().Unscoped().Where("id = ?", id).First(user)

	return res.Error
}

func GetUserByEmail(user *models.User, email string) error {
	res := db.DB().Where("email = ?", email).First(user)

	return res.Error
}

func UpdateUserEmail(user *models.User, id string, e string) error {
	res := db.DB().Raw("UPDATE users SET email = ? WHERE id = ?;", e, id).Scan(user)

	return res.Error
}

func UpdateUserPassword(user *models.User, id string, p []byte) error {
	res := db.DB().Raw("UPDATE users SET password = ? WHERE id = ?;", p, id).Scan(user)

	return res.Error
}

func DeleteUser(user *models.User, id string) error {
	res := db.DB().Delete(user, id)

	return res.Error
}

func DeleteUserPermanently(user *models.User, id string) error {
	res := db.DB().Unscoped().Delete(user, id)

	return res.Error
}
