// Code generated by ent, DO NOT EDIT.

package distributororder

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the distributororder type in the database.
	Label = "distributor_order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderId holds the string denoting the orderid field in the database.
	FieldOrderId = "order_id"
	// FieldBatchId holds the string denoting the batchid field in the database.
	FieldBatchId = "batch_id"
	// FieldProjectId holds the string denoting the projectid field in the database.
	FieldProjectId = "project_id"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldCreateAt holds the string denoting the createat field in the database.
	FieldCreateAt = "create_at"
	// FieldPTaskId holds the string denoting the ptaskid field in the database.
	FieldPTaskId = "p_task_id"
	// FieldProof holds the string denoting the proof field in the database.
	FieldProof = "proof"
	// Table holds the table name of the distributororder in the database.
	Table = "distributor_orders"
)

// Columns holds all SQL columns for distributororder fields.
var Columns = []string{
	FieldID,
	FieldOrderId,
	FieldBatchId,
	FieldProjectId,
	FieldAccount,
	FieldAmount,
	FieldIndex,
	FieldCreateAt,
	FieldPTaskId,
	FieldProof,
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

// OrderOption defines the ordering options for the DistributorOrder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrderId orders the results by the orderId field.
func ByOrderId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderId, opts...).ToFunc()
}

// ByBatchId orders the results by the batchId field.
func ByBatchId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBatchId, opts...).ToFunc()
}

// ByProjectId orders the results by the projectId field.
func ByProjectId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectId, opts...).ToFunc()
}

// ByAccount orders the results by the account field.
func ByAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccount, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByIndex orders the results by the index field.
func ByIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndex, opts...).ToFunc()
}

// ByCreateAt orders the results by the createAt field.
func ByCreateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateAt, opts...).ToFunc()
}

// ByPTaskId orders the results by the pTaskId field.
func ByPTaskId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPTaskId, opts...).ToFunc()
}
