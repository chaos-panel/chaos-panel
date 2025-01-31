// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

// RolesClosure is the golang structure for table roles_closure.
type RolesClosure struct {
	AncestorId   uint64 `json:"ancestorId"   orm:"ancestor_id"   ` // ancestor id
	DescendantId uint64 `json:"descendantId" orm:"descendant_id" ` // descendant id
	Distance     uint64 `json:"distance"     orm:"distance"      ` // distance
}
