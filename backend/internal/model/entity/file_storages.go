// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileStorages is the golang structure for table file_storages.
type FileStorages struct {
	Id            uint64      `json:"id"            orm:"id"             ` // ID
	ContentLength uint64      `json:"contentLength" orm:"content_length" ` // content length
	ContentMd5    string      `json:"contentMd5"    orm:"content_md5"    ` // content md5
	ContentSha1   string      `json:"contentSha1"   orm:"content_sha1"   ` // content sha1
	ContentSha256 string      `json:"contentSha256" orm:"content_sha256" ` // content sha256
	ContentSha512 string      `json:"contentSha512" orm:"content_sha512" ` // content sha512
	StorageUri    string      `json:"storageUri"    orm:"storage_uri"    ` // storage uri
	CreatedBy     uint64      `json:"createdBy"     orm:"created_by"     ` // created by
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     ` // created at
	UpdatedBy     uint64      `json:"updatedBy"     orm:"updated_by"     ` // updated by
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     ` // updated at
	DeletedBy     uint64      `json:"deletedBy"     orm:"deleted_by"     ` // deleted by
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     ` // deleted at
}
