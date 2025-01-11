// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-11 20:17:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure for table logs.
type Logs struct {
	Id         uint64      `json:"id"         orm:"id"          ` // id
	TenantId   uint64      `json:"tenantId"   orm:"tenant_id"   ` // tenant id
	ClientType string      `json:"clientType" orm:"client_type" ` // client type
	ClientInfo string      `json:"clientInfo" orm:"client_info" ` // client info
	Remark     string      `json:"remark"     orm:"remark"      ` // remark
	Log        string      `json:"log"        orm:"log"         ` // log
	CreatedBy  uint64      `json:"createdBy"  orm:"created_by"  ` // created by
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // locked at
}
