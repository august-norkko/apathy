package security

import "golang.org/x/crypto/bcrypt"

func Generate(pw []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func Compare(hash, pw []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		return false, err
	}
	return true, nil
}