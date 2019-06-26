package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesCheckSums(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr1 string
	}{
		{
			name:     "check checksum 01",
			jsonStr1: `{"auths":{"cloud.okd.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"abc@xyz.com"},"quay.io":{"auth":"b3BlbnNoVkc=","email":"abc@xyz.com"},"registry.connect.redhat.com":{"auth":"NjQ4ODeDZ3d1pN","email":"abc@xyz.com"},"registry.redhat.io":{"auth":"NjQ4ODX1pN","email":"abc@xyz.com"}}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultHash1 := GetHashOfPullSecret(test.jsonStr1)
			for i := 0; i < 10; i++ {
				resultHash2 := GetHashOfPullSecret(test.jsonStr1)
				assert.Equal(t, resultHash1, resultHash2)
			}
		})
	}
}
