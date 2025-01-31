// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

// ResourcesClosure is the golang structure for table resources_closure.
type ResourcesClosure struct {
	AncestorId   uint64 `json:"ancestorId"   orm:"ancestor_id"   ` // ancestor id
	DescendantId uint64 `json:"descendantId" orm:"descendant_id" ` // descendant id
	Distance     int64  `json:"distance"     orm:"distance"      ` // distance
}
