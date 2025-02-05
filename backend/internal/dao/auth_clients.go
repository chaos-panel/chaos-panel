// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"github.com/chaos-plus/chaos-plus/internal/dao/internal"
)

// internalAuthClientsDao is an internal type for wrapping the internal DAO implementation.
type internalAuthClientsDao = *internal.AuthClientsDao

// authClientsDao is the data access object for the table chaosplus_auth_clients.
// You can define custom methods on it to extend its functionality as needed.
type authClientsDao struct {
	internalAuthClientsDao
}

var (
	// AuthClients is a globally accessible object for table chaosplus_auth_clients operations.
	AuthClients = authClientsDao{
		internal.NewAuthClientsDao(),
	}
)

// Add your custom methods and functionality below.
