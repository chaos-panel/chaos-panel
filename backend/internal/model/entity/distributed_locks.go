// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-18 00:02:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DistributedLocks is the golang structure for table distributed_locks.
type DistributedLocks struct {
	LockKey   string      `json:"lockKey"   orm:"lock_key"   ` // lock key
	HostInfo  string      `json:"hostInfo"  orm:"host_info"  ` // host info
	ExpiredAt *gtime.Time `json:"expiredAt" orm:"expired_at" ` // expired at
	CreatedBy string      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // locked at
	UpdatedBy string      `json:"updatedBy" orm:"updated_by" ` // updated by
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // time updated
}
