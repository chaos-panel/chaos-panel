// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-14 17:54:17
// =================================================================================

package entity

// TSequence is the golang structure for table t_sequence.
type TSequence struct {
	Id       uint64 `json:"id"       orm:"id"       ` // ID
	Type     string `json:"type"     orm:"type"     ` // type
	Sequence uint64 `json:"sequence" orm:"sequence" ` // sequence
}
