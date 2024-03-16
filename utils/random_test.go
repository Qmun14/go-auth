package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	randStr1 := RandomString(32)
	randStr2 := RandomString(32)

	fmt.Printf("The rand value is = %s \n", randStr1)
	fmt.Printf("The rand value is = %s \n", randStr2)

	assert.NotEmpty(t, randStr1)
	assert.NotEmpty(t, randStr2)

	assert.NotEqual(t, randStr1, randStr2)
}
func TestRandomEmail(t *testing.T) {
	randEmail1 := RandomEmail()
	randEmail2 := RandomEmail()

	fmt.Printf("The rand value is = %s \n", randEmail1)
	fmt.Printf("The rand value is = %s \n", randEmail2)

	assert.NotEmpty(t, randEmail1)
	assert.NotEmpty(t, randEmail2)

	assert.NotEqual(t, randEmail1, randEmail2)
}
