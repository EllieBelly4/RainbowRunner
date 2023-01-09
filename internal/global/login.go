package global

// TODO move all of this somewhere permanent, I just need this logic for testing

import (
	"RainbowRunner/internal/types"
	"errors"
	"math/rand"
	"time"
)

type loginRequest struct {
	loginName     string
	oneTimePass   uint32
	timeRequested time.Time
}

var loginRequests = make(map[uint32]loginRequest)

func GenerateOneTimeKey() uint32 {
	var key = rand.Uint32()

	if _, ok := loginRequests[key]; ok {
		return GenerateOneTimeKey()
	}

	return key
}

func GetAccountFromOneTimeKey(oneTimeKey uint32) *string {
	if _, ok := loginRequests[oneTimeKey]; !ok {
		return nil
	}

	return types.Pointer(loginRequests[oneTimeKey].loginName)
}

func AddLoginRequest(oneTimeKey uint32, loginName string) error {
	_, ok := loginRequests[oneTimeKey]

	if ok {
		return errors.New("login request with that key already exists")
	}

	loginRequests[oneTimeKey] = loginRequest{
		loginName:     loginName,
		oneTimePass:   oneTimeKey,
		timeRequested: time.Now(),
	}

	return nil
}
