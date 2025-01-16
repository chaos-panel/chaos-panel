// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Labels is the golang structure for table labels.
type Labels struct {
	LabelId    uint64      `json:"labelId"    orm:"label_id"    ` // label id
	UnionId    uint64      `json:"unionId"    orm:"union_id"    ` // union id
	LabelIndex uint64      `json:"labelIndex" orm:"label_index" ` // label index
	LabelName  string      `json:"labelName"  orm:"label_name"  ` // label name
	LabelColor string      `json:"labelColor" orm:"label_color" ` // label color
	CreatedBy  uint64      `json:"createdBy"  orm:"created_by"  ` // created by
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // locked at
}
