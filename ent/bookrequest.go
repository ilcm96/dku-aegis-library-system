// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ilcm96/dku-aegis-library/ent/bookrequest"
)

// BookRequest is the model entity for the BookRequest schema.
type BookRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// Publisher holds the value of the "publisher" field.
	Publisher string `json:"publisher,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// Approved holds the value of the "approved" field.
	Approved bookrequest.Approved `json:"approved,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookrequest.FieldID, bookrequest.FieldUserID:
			values[i] = new(sql.NullInt64)
		case bookrequest.FieldTitle, bookrequest.FieldAuthor, bookrequest.FieldPublisher, bookrequest.FieldReason, bookrequest.FieldApproved:
			values[i] = new(sql.NullString)
		case bookrequest.FieldCreatedAt, bookrequest.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookRequest fields.
func (br *BookRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookrequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			br.ID = int(value.Int64)
		case bookrequest.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				br.UserID = int(value.Int64)
			}
		case bookrequest.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				br.Title = value.String
			}
		case bookrequest.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				br.Author = value.String
			}
		case bookrequest.FieldPublisher:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publisher", values[i])
			} else if value.Valid {
				br.Publisher = value.String
			}
		case bookrequest.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				br.Reason = value.String
			}
		case bookrequest.FieldApproved:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field approved", values[i])
			} else if value.Valid {
				br.Approved = bookrequest.Approved(value.String)
			}
		case bookrequest.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				br.CreatedAt = value.Time
			}
		case bookrequest.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				br.UpdatedAt = value.Time
			}
		default:
			br.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BookRequest.
// This includes values selected through modifiers, order, etc.
func (br *BookRequest) Value(name string) (ent.Value, error) {
	return br.selectValues.Get(name)
}

// Update returns a builder for updating this BookRequest.
// Note that you need to call BookRequest.Unwrap() before calling this method if this BookRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (br *BookRequest) Update() *BookRequestUpdateOne {
	return NewBookRequestClient(br.config).UpdateOne(br)
}

// Unwrap unwraps the BookRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (br *BookRequest) Unwrap() *BookRequest {
	_tx, ok := br.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookRequest is not a transactional entity")
	}
	br.config.driver = _tx.drv
	return br
}

// String implements the fmt.Stringer.
func (br *BookRequest) String() string {
	var builder strings.Builder
	builder.WriteString("BookRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", br.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", br.UserID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(br.Title)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(br.Author)
	builder.WriteString(", ")
	builder.WriteString("publisher=")
	builder.WriteString(br.Publisher)
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(br.Reason)
	builder.WriteString(", ")
	builder.WriteString("approved=")
	builder.WriteString(fmt.Sprintf("%v", br.Approved))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(br.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(br.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BookRequests is a parsable slice of BookRequest.
type BookRequests []*BookRequest
