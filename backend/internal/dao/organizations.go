// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"github.com/chaos-plus/chaos-plus/internal/dao/internal"
)

// internalOrganizationsDao is an internal type for wrapping the internal DAO implementation.
type internalOrganizationsDao = *internal.OrganizationsDao

// organizationsDao is the data access object for the table chaosplus_organizations.
// You can define custom methods on it to extend its functionality as needed.
type organizationsDao struct {
	internalOrganizationsDao
}

var (
	// Organizations is a globally accessible object for table chaosplus_organizations operations.
	Organizations = organizationsDao{
		internal.NewOrganizationsDao(),
	}
)

// Add your custom methods and functionality below.
