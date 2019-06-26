package utils

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMergeJsons(t *testing.T) {
	tests := []struct {
		name        string
		jsonStr1    string
		jsonStr2    string
		expectedStr string
		expectedErr bool
	}{
		{
			name:        "Merge pull secrets 01",
			jsonStr1:    `{"auths": {"registry.svc.ci.openshift.org": {"auth": "dXNljlfjldsfSDD"}}}`,
			jsonStr2:    `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"lmohanty@redhat.com"},"quay.io":{"auth":"b3BlbnNoVkc=","email":"lmohanty@redhat.com"},"registry.connect.redhat.com":{"auth":"NjQ4ODeDZ3d1pN","email":"lmohanty@redhat.com"},"registry.redhat.io":{"auth":"NjQ4ODX1pN","email":"lmohanty@redhat.com"}}}`,
			expectedStr: `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"lmohanty@redhat.com"},"quay.io":{"auth":"b3BlbnNoVkc=","email":"lmohanty@redhat.com"},"registry.connect.redhat.com":{"auth":"NjQ4ODeDZ3d1pN","email":"lmohanty@redhat.com"},"registry.redhat.io":{"auth":"NjQ4ODX1pN","email":"lmohanty@redhat.com"},"registry.svc.ci.openshift.org":{"auth":"dXNljlfjldsfSDD"}}}`,
		},
		{
			name:        "Merge pull secrets 02",
			jsonStr1:    `{"auths": {"registry.svc.ci.openshift.org": {"auth": "dXNljlfjldsfSDD"}}}`,
			jsonStr2:    `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"lmohanty@redhat.com"}}}`,
			expectedStr: `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"lmohanty@redhat.com"},"registry.svc.ci.openshift.org":{"auth":"dXNljlfjldsfSDD"}}}`,
		},
		{
			name:        "Merege global and local same auth key but different secret",
			jsonStr1:    `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"lmohanty@redhat.com"}}}`,
			jsonStr2:    `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTlocal=","email":"lmohanty@redhat.com"}}}`,
			expectedStr: `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTlocal=","email":"lmohanty@redhat.com"}}}`,
		},
		{
			name:        "Merge of pull secrets should fail",
			jsonStr1:    `{"auths": {"registry.svc.ci.openshift.org": {"auth": "dXNljlfjldsfSDD"}}}`,
			jsonStr2:    `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"lmohanty@redhat.com"}}`,
			expectedStr: `{"auths":{"cloud.openshift.com":{"auth":"b34xVjWERckjfUyV1pMQTc=","email":"lmohanty@redhat.com"}}}`,
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tLogger := log.New()
			resultStr, err := MergeJsons(test.jsonStr1, test.jsonStr2, tLogger)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.expectedStr, resultStr)
			}
		})
	}

}
