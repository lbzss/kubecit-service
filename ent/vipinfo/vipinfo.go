// Code generated by ent, DO NOT EDIT.

package vipinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the vipinfo type in the database.
	Label = "vip_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVipType holds the string denoting the vip_type field in the database.
	FieldVipType = "vip_type"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldExpireAt holds the string denoting the expire_at field in the database.
	FieldExpireAt = "expire_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUserInfo holds the string denoting the user_info edge name in mutations.
	EdgeUserInfo = "user_info"
	// Table holds the table name of the vipinfo in the database.
	Table = "vip_infos"
	// UserInfoTable is the table that holds the user_info relation/edge.
	UserInfoTable = "vip_infos"
	// UserInfoInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInfoInverseTable = "users"
	// UserInfoColumn is the table column denoting the user_info relation/edge.
	UserInfoColumn = "user_id"
)

// Columns holds all SQL columns for vipinfo fields.
var Columns = []string{
	FieldID,
	FieldVipType,
	FieldStartAt,
	FieldExpireAt,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt func() time.Time
)

// OrderOption defines the ordering options for the VipInfo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVipType orders the results by the vip_type field.
func ByVipType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVipType, opts...).ToFunc()
}

// ByStartAt orders the results by the start_at field.
func ByStartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartAt, opts...).ToFunc()
}

// ByExpireAt orders the results by the expire_at field.
func ByExpireAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpireAt, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByUserInfoField orders the results by user_info field.
func ByUserInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserInfoStep(), sql.OrderByField(field, opts...))
	}
}
func newUserInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserInfoTable, UserInfoColumn),
	)
}