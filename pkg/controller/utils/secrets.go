package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	k8sDeepHash "github.com/openshift/hive/contrib/pkg/hash"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// LoadSecretData loads a given secret key and returns it's data as a string.
func LoadSecretData(c client.Client, secretName, namespace, dataKey string) (string, error) {
	s := &corev1.Secret{}
	err := c.Get(context.TODO(), types.NamespacedName{Name: secretName, Namespace: namespace}, s)
	if err != nil {
		return "", err
	}
	retStr, ok := s.Data[dataKey]
	if !ok {
		return "", fmt.Errorf("secret %s did not contain key %s", secretName, dataKey)
	}
	return string(retStr), nil
}

// CalculateHashOfSecret returns a hash of the corev1.Secret
// The secret is a map and we can't guarantee its order of the output so we need find a way to reproducible hash
func CalculateHashOfSecret(secret interface{}) string {
	hasher := sha256.New()
	k8sDeepHash.DeepHashObject(hasher, secret)
	sum := hex.EncodeToString(hasher.Sum(nil))
	return sum
}
