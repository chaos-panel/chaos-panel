// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CasbinRules is the golang structure for table casbin_rules.
type CasbinRules struct {
	Id        uint64      `json:"id"        orm:"id"         ` // ID
	PType     string      `json:"pType"     orm:"p_type"     ` //
	V0        string      `json:"v0"        orm:"v0"         ` //
	V1        string      `json:"v1"        orm:"v1"         ` //
	V2        string      `json:"v2"        orm:"v2"         ` //
	V3        string      `json:"v3"        orm:"v3"         ` //
	V4        string      `json:"v4"        orm:"v4"         ` //
	V5        string      `json:"v5"        orm:"v5"         ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
}
