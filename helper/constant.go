package helper

import (
	"errors"
	"time"
)

var ErrorInvalidLogin = errors.New("invalid email or password")
var ErrorEmailAlreadyExists = errors.New("email already exists")
var ErrorUsernameAlreadyExists = errors.New("username already exists")
var TimeNowMillis = time.Now().UnixNano() / int64(time.Millisecond)
