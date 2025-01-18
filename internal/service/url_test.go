package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase62Encode(t *testing.T) {
	input := int64(1737213280)
	exp := "1tZAX2"

	got := base62Enc(input)

	assert.Equal(t, exp, got)
}

func TestBase62Decode(t *testing.T) {
	input := "1tZAX2"
	exp := int64(1737213280)

	got, err := base62Dec(input)

	assert.NoError(t, err)
	assert.Equal(t, exp, got)

	inputWrong := ""
	_, err = base62Dec(inputWrong)
	assert.Equal(t, "empty string", err.Error())
}
