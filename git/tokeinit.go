package git

import (
	"errors"
	"os"
)

// get Initial  token function to get token from .bashrc .zshrc or env
// return string and error
func GetInitialToken() (string, error) {
	token := os.Getenv("github")
	if token == "" {
		return nil, errors.New("[!] github token not found in env / .bashrc / .zshrc etc")
	}
	return token, nil
}
