// Code generated by ent, DO NOT EDIT.

package bookrequest

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the bookrequest type in the database.
	Label = "book_request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldPublisher holds the string denoting the publisher field in the database.
	FieldPublisher = "publisher"
	// FieldReason holds the string denoting the reason field in the database.
	FieldReason = "reason"
	// FieldApproved holds the string denoting the approved field in the database.
	FieldApproved = "approved"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the bookrequest in the database.
	Table = "book_requests"
)

// Columns holds all SQL columns for bookrequest fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldTitle,
	FieldAuthor,
	FieldPublisher,
	FieldReason,
	FieldApproved,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Approved defines the type for the "approved" enum field.
type Approved string

// Approved values.
const (
	ApprovedPENDING  Approved = "PENDING"
	ApprovedAPPROVED Approved = "APPROVED"
	ApprovedREJECTED Approved = "REJECTED"
)

func (a Approved) String() string {
	return string(a)
}

// ApprovedValidator is a validator for the "approved" field enum values. It is called by the builders before save.
func ApprovedValidator(a Approved) error {
	switch a {
	case ApprovedPENDING, ApprovedAPPROVED, ApprovedREJECTED:
		return nil
	default:
		return fmt.Errorf("bookrequest: invalid enum value for approved field: %q", a)
	}
}

// OrderOption defines the ordering options for the BookRequest queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByPublisher orders the results by the publisher field.
func ByPublisher(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublisher, opts...).ToFunc()
}

// ByReason orders the results by the reason field.
func ByReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReason, opts...).ToFunc()
}

// ByApproved orders the results by the approved field.
func ByApproved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApproved, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
