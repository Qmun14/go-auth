package services

import (
	"testing"

	"github.com/Qmun14/jwtAuth/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateVerifyEmail(t *testing.T) {
	token := utils.RandomString(32)
	userEmail := "mamunramdhan@gmail.com"
	userName := "Qmun14"

	err := CreateVerifyEmail(1, token, userEmail, userName)
	assert.NoError(t, err)

}
