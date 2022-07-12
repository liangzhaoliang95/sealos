// Copyright © 2022 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sso

import (
	"strings"

	"github.com/labring/sealos/pkg/auth/conf"
	"github.com/labring/sealos/pkg/utils/logger"
)

type User struct {
	ID   string
	Name string
}

type Client interface {
	GetRedirectURL() (string, error)
	GetUserInfo(state string, code string) (User, error)
}

type ClientType string

const CasdoorClientType ClientType = "casdoor"

func InitSSO() Client {
	var client Client
	switch ClientType(strings.ToLower(conf.GlobalConfig.SSOType)) {
	case CasdoorClientType:
		client = NewCasdoorClient()
	default:
		logger.Warn("No valid SSO platform specified")
	}
	return client
}