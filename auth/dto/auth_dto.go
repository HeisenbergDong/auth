package dto

import (
	"auth/utils"
)

var AuthorityIdVerify = utils.Rules{
	"AuthorityId": {utils.NotEmpty()},
}

type AuthInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type UpdateAuthDto struct {
	AuthorityId string       `json:"authorityId"`
	AuthInfos []AuthInfo     `json:"authInfos"`
}