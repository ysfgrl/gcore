package ghelper

import (
	"github.com/ysfgrl/gcore/gerror"
	"golang.org/x/crypto/bcrypt"
)

var Bcrypt = bc{}

type bc struct {
}

func (b *bc) Encrypt(password string) (string, *gerror.Error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", gerror.GetError(err)
	}
	return string(hashed), nil
}

func (b *bc) Verify(hashed string, password string) *gerror.Error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return gerror.GetError(err)
	}
	return nil
}
