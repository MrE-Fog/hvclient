/*
Copyright (c) 2019-2021 GMO GlobalSign Pte. Ltd.

Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at

https://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains settings from an HVClient configuration file.
type Config struct {
	// URL is the URL (including port) to the HVCA server.
	URL string `json:"url"`

	// APIKey is the client-specific API key used to login.
	APIKey string `json:"api_key"`

	// APISecret is the client-specific API secret used to login.
	APISecret string `json:"api_secret"`

	// CertFile is the path of the client certificate file.
	CertFile string `json:"cert_file"`

	// KeyFile is the path of the client key file.
	KeyFile string `json:"key_file"`

	// KeyPassphrase is the passphrase for the client key. If the key is not
	// encrypted, this should be set to the emptry string.
	KeyPassphrase string `json:"key_passphrase"`

	// If InsecureSkipVerify is true, TLS accepts any certificate
	// presented by the server and any host name in that certificate.
	// In this mode, TLS is susceptible to man-in-the-middle attacks.
	// This should be used only for testing.
	InsecureSkipVerify bool `json:"insecure_skip_verify"`

	// ExtraHeaders contains custom HTTP request headers to be passed to the
	// HVCA server with each request.
	ExtraHeaders map[string]string `json:"extra_headers,omitempty"`

	// Timeout is the maximum time in seconds for an HVCA API request.
	Timeout int `json:"timeout"`
}

// NewFromFile creates a new Config object from a configuration file.
func NewFromFile(filename string) (*Config, error) {
	var data, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var newConfig *Config
	err = json.Unmarshal(data, &newConfig)
	if err != nil {
		return nil, err
	}

	return newConfig, nil
}
