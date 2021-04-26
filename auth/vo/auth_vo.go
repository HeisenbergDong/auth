package vo

import "auth/auth/dto"

type PolicyPathVo struct {
	Paths []dto.AuthInfo `json:"paths"`
}