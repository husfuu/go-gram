package helper

import "errors"

var ErrorInvalidLogin = errors.New("invalid email or password")
var ErrorEmailAlreadyExists = errors.New("email already exists")
var ErrorUsernameAlreadyExists = errors.New("username already exists")
