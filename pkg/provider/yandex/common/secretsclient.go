/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package common

import (
	"context"
	"fmt"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	esv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
)

// https://github.com/external-secrets/external-secrets/issues/644
var _ esv1beta1.SecretsClient = &yandexCloudSecretsClient{}

// Implementation of v1beta1.SecretsClient.
type yandexCloudSecretsClient struct {
	secretGetter SecretGetter
	secretSetter SecretSetter
	iamToken     string
}

func (c *yandexCloudSecretsClient) GetSecret(ctx context.Context, ref esv1beta1.ExternalSecretDataRemoteRef) ([]byte, error) {
	return c.secretGetter.GetSecret(ctx, c.iamToken, ref.Key, ref.Version, ref.Property)
}

func (c *yandexCloudSecretsClient) DeleteSecret(_ context.Context, _ esv1beta1.PushRemoteRef) error {
	return fmt.Errorf("not implemented")
}

func (c *yandexCloudSecretsClient) PushSecret(_ context.Context, _ []byte, _ *apiextensionsv1.JSON, _ esv1beta1.PushRemoteRef) error {
	return fmt.Errorf("not implemented")
}

func (c *yandexCloudSecretsClient) Validate() (esv1beta1.ValidationResult, error) {
	return esv1beta1.ValidationResultReady, nil
}

func (c *yandexCloudSecretsClient) GetSecretMap(ctx context.Context, ref esv1beta1.ExternalSecretDataRemoteRef) (map[string][]byte, error) {
	return c.secretGetter.GetSecretMap(ctx, c.iamToken, ref.Key, ref.Version)
}

func (c *yandexCloudSecretsClient) GetAllSecrets(_ context.Context, _ esv1beta1.ExternalSecretFind) (map[string][]byte, error) {
	// TO be implemented
	return nil, fmt.Errorf("GetAllSecrets not supported")
}

func (c *yandexCloudSecretsClient) Close(_ context.Context) error {
	return nil
}
