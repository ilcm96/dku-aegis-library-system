// Code generated by ent, DO NOT EDIT.

package book

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ilcm96/dku-aegis-library/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAuthor, v))
}

// Publisher applies equality check predicate on the "publisher" field. It's identical to PublisherEQ.
func Publisher(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublisher, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldQuantity, v))
}

// Borrow applies equality check predicate on the "borrow" field. It's identical to BorrowEQ.
func Borrow(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldBorrow, v))
}

// Cover applies equality check predicate on the "cover" field. It's identical to CoverEQ.
func Cover(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCover, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCategory, v))
}

// Isbn applies equality check predicate on the "isbn" field. It's identical to IsbnEQ.
func Isbn(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldTitle, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldAuthor, v))
}

// PublisherEQ applies the EQ predicate on the "publisher" field.
func PublisherEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublisher, v))
}

// PublisherNEQ applies the NEQ predicate on the "publisher" field.
func PublisherNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldPublisher, v))
}

// PublisherIn applies the In predicate on the "publisher" field.
func PublisherIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldPublisher, vs...))
}

// PublisherNotIn applies the NotIn predicate on the "publisher" field.
func PublisherNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldPublisher, vs...))
}

// PublisherGT applies the GT predicate on the "publisher" field.
func PublisherGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldPublisher, v))
}

// PublisherGTE applies the GTE predicate on the "publisher" field.
func PublisherGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldPublisher, v))
}

// PublisherLT applies the LT predicate on the "publisher" field.
func PublisherLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldPublisher, v))
}

// PublisherLTE applies the LTE predicate on the "publisher" field.
func PublisherLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldPublisher, v))
}

// PublisherContains applies the Contains predicate on the "publisher" field.
func PublisherContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldPublisher, v))
}

// PublisherHasPrefix applies the HasPrefix predicate on the "publisher" field.
func PublisherHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldPublisher, v))
}

// PublisherHasSuffix applies the HasSuffix predicate on the "publisher" field.
func PublisherHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldPublisher, v))
}

// PublisherEqualFold applies the EqualFold predicate on the "publisher" field.
func PublisherEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldPublisher, v))
}

// PublisherContainsFold applies the ContainsFold predicate on the "publisher" field.
func PublisherContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldPublisher, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldQuantity, v))
}

// BorrowEQ applies the EQ predicate on the "borrow" field.
func BorrowEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldBorrow, v))
}

// BorrowNEQ applies the NEQ predicate on the "borrow" field.
func BorrowNEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldBorrow, v))
}

// BorrowIn applies the In predicate on the "borrow" field.
func BorrowIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldBorrow, vs...))
}

// BorrowNotIn applies the NotIn predicate on the "borrow" field.
func BorrowNotIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldBorrow, vs...))
}

// BorrowGT applies the GT predicate on the "borrow" field.
func BorrowGT(v int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldBorrow, v))
}

// BorrowGTE applies the GTE predicate on the "borrow" field.
func BorrowGTE(v int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldBorrow, v))
}

// BorrowLT applies the LT predicate on the "borrow" field.
func BorrowLT(v int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldBorrow, v))
}

// BorrowLTE applies the LTE predicate on the "borrow" field.
func BorrowLTE(v int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldBorrow, v))
}

// CoverEQ applies the EQ predicate on the "cover" field.
func CoverEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCover, v))
}

// CoverNEQ applies the NEQ predicate on the "cover" field.
func CoverNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCover, v))
}

// CoverIn applies the In predicate on the "cover" field.
func CoverIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCover, vs...))
}

// CoverNotIn applies the NotIn predicate on the "cover" field.
func CoverNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCover, vs...))
}

// CoverGT applies the GT predicate on the "cover" field.
func CoverGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCover, v))
}

// CoverGTE applies the GTE predicate on the "cover" field.
func CoverGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCover, v))
}

// CoverLT applies the LT predicate on the "cover" field.
func CoverLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCover, v))
}

// CoverLTE applies the LTE predicate on the "cover" field.
func CoverLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCover, v))
}

// CoverContains applies the Contains predicate on the "cover" field.
func CoverContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldCover, v))
}

// CoverHasPrefix applies the HasPrefix predicate on the "cover" field.
func CoverHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldCover, v))
}

// CoverHasSuffix applies the HasSuffix predicate on the "cover" field.
func CoverHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldCover, v))
}

// CoverEqualFold applies the EqualFold predicate on the "cover" field.
func CoverEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldCover, v))
}

// CoverContainsFold applies the ContainsFold predicate on the "cover" field.
func CoverContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldCover, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldCategory, v))
}

// IsbnEQ applies the EQ predicate on the "isbn" field.
func IsbnEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn, v))
}

// IsbnNEQ applies the NEQ predicate on the "isbn" field.
func IsbnNEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldIsbn, v))
}

// IsbnIn applies the In predicate on the "isbn" field.
func IsbnIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldIsbn, vs...))
}

// IsbnNotIn applies the NotIn predicate on the "isbn" field.
func IsbnNotIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldIsbn, vs...))
}

// IsbnGT applies the GT predicate on the "isbn" field.
func IsbnGT(v int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldIsbn, v))
}

// IsbnGTE applies the GTE predicate on the "isbn" field.
func IsbnGTE(v int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldIsbn, v))
}

// IsbnLT applies the LT predicate on the "isbn" field.
func IsbnLT(v int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldIsbn, v))
}

// IsbnLTE applies the LTE predicate on the "isbn" field.
func IsbnLTE(v int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldIsbn, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(sql.NotPredicates(p))
}
