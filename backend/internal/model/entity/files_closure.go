// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package entity

// FilesClosure is the golang structure for table files_closure.
type FilesClosure struct {
	AncestorId   uint64 `json:"ancestorId"   orm:"ancestor_id"   ` // ancestor id
	DescendantId uint64 `json:"descendantId" orm:"descendant_id" ` // descendant id
	Distance     int64  `json:"distance"     orm:"distance"      ` // distance
}
