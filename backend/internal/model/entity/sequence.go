// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

// Sequence is the golang structure for table sequence.
type Sequence struct {
	Id       uint64 `json:"id"       orm:"id"       ` // ID
	Type     string `json:"type"     orm:"type"     ` // type
	Sequence uint64 `json:"sequence" orm:"sequence" ` // sequence
}
