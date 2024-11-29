package filedb

import (
	"demo/models"
	"encoding/json"
	"os"

	"math/rand"
)

type UserFileDb struct {
	FileName string
}

func NewUser(filename string) *UserFileDb {
	return &UserFileDb{FileName: filename}
}
func (f *UserFileDb) CreateUser(user *models.User) (*models.User, error) {
	user.Id = uint(rand.Intn(10000))
	_, err := writeUserToFile(f.FileName, user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func writeUserToFile(fineName string, user *models.User) (int, error) {
	f, err := os.OpenFile(fineName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	bytes, err := json.Marshal(user)
	if err != nil {
		return 0, err
	}
	return f.Write(bytes)
}
