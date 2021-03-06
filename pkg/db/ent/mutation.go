// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTransaction = "Transaction"
)

// TransactionMutation represents an operation that mutates the Transaction nodes in the graph.
type TransactionMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	name           *string
	amount         *uint64
	addamount      *uint64
	from           *string
	to             *string
	transaction_id *string
	cid            *string
	exit_code      *int64
	addexit_code   *int64
	status         *transaction.Status
	created_at     *uint32
	addcreated_at  *uint32
	updated_at     *uint32
	addupdated_at  *uint32
	deleted_at     *uint32
	adddeleted_at  *uint32
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Transaction, error)
	predicates     []predicate.Transaction
}

var _ ent.Mutation = (*TransactionMutation)(nil)

// transactionOption allows management of the mutation configuration using functional options.
type transactionOption func(*TransactionMutation)

// newTransactionMutation creates new mutation for the Transaction entity.
func newTransactionMutation(c config, op Op, opts ...transactionOption) *TransactionMutation {
	m := &TransactionMutation{
		config:        c,
		op:            op,
		typ:           TypeTransaction,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTransactionID sets the ID field of the mutation.
func withTransactionID(id uuid.UUID) transactionOption {
	return func(m *TransactionMutation) {
		var (
			err   error
			once  sync.Once
			value *Transaction
		)
		m.oldValue = func(ctx context.Context) (*Transaction, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Transaction.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTransaction sets the old Transaction of the mutation.
func withTransaction(node *Transaction) transactionOption {
	return func(m *TransactionMutation) {
		m.oldValue = func(context.Context) (*Transaction, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TransactionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TransactionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Transaction entities.
func (m *TransactionMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TransactionMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *TransactionMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TransactionMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TransactionMutation) ResetName() {
	m.name = nil
}

// SetAmount sets the "amount" field.
func (m *TransactionMutation) SetAmount(u uint64) {
	m.amount = &u
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *TransactionMutation) Amount() (r uint64, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldAmount(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds u to the "amount" field.
func (m *TransactionMutation) AddAmount(u uint64) {
	if m.addamount != nil {
		*m.addamount += u
	} else {
		m.addamount = &u
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *TransactionMutation) AddedAmount() (r uint64, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *TransactionMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetFrom sets the "from" field.
func (m *TransactionMutation) SetFrom(s string) {
	m.from = &s
}

// From returns the value of the "from" field in the mutation.
func (m *TransactionMutation) From() (r string, exists bool) {
	v := m.from
	if v == nil {
		return
	}
	return *v, true
}

// OldFrom returns the old "from" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldFrom(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldFrom is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldFrom requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFrom: %w", err)
	}
	return oldValue.From, nil
}

// ResetFrom resets all changes to the "from" field.
func (m *TransactionMutation) ResetFrom() {
	m.from = nil
}

// SetTo sets the "to" field.
func (m *TransactionMutation) SetTo(s string) {
	m.to = &s
}

// To returns the value of the "to" field in the mutation.
func (m *TransactionMutation) To() (r string, exists bool) {
	v := m.to
	if v == nil {
		return
	}
	return *v, true
}

// OldTo returns the old "to" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldTo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTo: %w", err)
	}
	return oldValue.To, nil
}

// ResetTo resets all changes to the "to" field.
func (m *TransactionMutation) ResetTo() {
	m.to = nil
}

// SetTransactionID sets the "transaction_id" field.
func (m *TransactionMutation) SetTransactionID(s string) {
	m.transaction_id = &s
}

// TransactionID returns the value of the "transaction_id" field in the mutation.
func (m *TransactionMutation) TransactionID() (r string, exists bool) {
	v := m.transaction_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTransactionID returns the old "transaction_id" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldTransactionID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTransactionID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTransactionID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTransactionID: %w", err)
	}
	return oldValue.TransactionID, nil
}

// ResetTransactionID resets all changes to the "transaction_id" field.
func (m *TransactionMutation) ResetTransactionID() {
	m.transaction_id = nil
}

// SetCid sets the "cid" field.
func (m *TransactionMutation) SetCid(s string) {
	m.cid = &s
}

// Cid returns the value of the "cid" field in the mutation.
func (m *TransactionMutation) Cid() (r string, exists bool) {
	v := m.cid
	if v == nil {
		return
	}
	return *v, true
}

// OldCid returns the old "cid" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldCid(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCid: %w", err)
	}
	return oldValue.Cid, nil
}

// ResetCid resets all changes to the "cid" field.
func (m *TransactionMutation) ResetCid() {
	m.cid = nil
}

// SetExitCode sets the "exit_code" field.
func (m *TransactionMutation) SetExitCode(i int64) {
	m.exit_code = &i
	m.addexit_code = nil
}

// ExitCode returns the value of the "exit_code" field in the mutation.
func (m *TransactionMutation) ExitCode() (r int64, exists bool) {
	v := m.exit_code
	if v == nil {
		return
	}
	return *v, true
}

// OldExitCode returns the old "exit_code" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldExitCode(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldExitCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldExitCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExitCode: %w", err)
	}
	return oldValue.ExitCode, nil
}

// AddExitCode adds i to the "exit_code" field.
func (m *TransactionMutation) AddExitCode(i int64) {
	if m.addexit_code != nil {
		*m.addexit_code += i
	} else {
		m.addexit_code = &i
	}
}

// AddedExitCode returns the value that was added to the "exit_code" field in this mutation.
func (m *TransactionMutation) AddedExitCode() (r int64, exists bool) {
	v := m.addexit_code
	if v == nil {
		return
	}
	return *v, true
}

// ResetExitCode resets all changes to the "exit_code" field.
func (m *TransactionMutation) ResetExitCode() {
	m.exit_code = nil
	m.addexit_code = nil
}

// SetStatus sets the "status" field.
func (m *TransactionMutation) SetStatus(t transaction.Status) {
	m.status = &t
}

// Status returns the value of the "status" field in the mutation.
func (m *TransactionMutation) Status() (r transaction.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldStatus(ctx context.Context) (v transaction.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *TransactionMutation) ResetStatus() {
	m.status = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TransactionMutation) SetCreatedAt(u uint32) {
	m.created_at = &u
	m.addcreated_at = nil
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TransactionMutation) CreatedAt() (r uint32, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldCreatedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// AddCreatedAt adds u to the "created_at" field.
func (m *TransactionMutation) AddCreatedAt(u uint32) {
	if m.addcreated_at != nil {
		*m.addcreated_at += u
	} else {
		m.addcreated_at = &u
	}
}

// AddedCreatedAt returns the value that was added to the "created_at" field in this mutation.
func (m *TransactionMutation) AddedCreatedAt() (r uint32, exists bool) {
	v := m.addcreated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TransactionMutation) ResetCreatedAt() {
	m.created_at = nil
	m.addcreated_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TransactionMutation) SetUpdatedAt(u uint32) {
	m.updated_at = &u
	m.addupdated_at = nil
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TransactionMutation) UpdatedAt() (r uint32, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldUpdatedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// AddUpdatedAt adds u to the "updated_at" field.
func (m *TransactionMutation) AddUpdatedAt(u uint32) {
	if m.addupdated_at != nil {
		*m.addupdated_at += u
	} else {
		m.addupdated_at = &u
	}
}

// AddedUpdatedAt returns the value that was added to the "updated_at" field in this mutation.
func (m *TransactionMutation) AddedUpdatedAt() (r uint32, exists bool) {
	v := m.addupdated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TransactionMutation) ResetUpdatedAt() {
	m.updated_at = nil
	m.addupdated_at = nil
}

// SetDeletedAt sets the "deleted_at" field.
func (m *TransactionMutation) SetDeletedAt(u uint32) {
	m.deleted_at = &u
	m.adddeleted_at = nil
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *TransactionMutation) DeletedAt() (r uint32, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldDeletedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// AddDeletedAt adds u to the "deleted_at" field.
func (m *TransactionMutation) AddDeletedAt(u uint32) {
	if m.adddeleted_at != nil {
		*m.adddeleted_at += u
	} else {
		m.adddeleted_at = &u
	}
}

// AddedDeletedAt returns the value that was added to the "deleted_at" field in this mutation.
func (m *TransactionMutation) AddedDeletedAt() (r uint32, exists bool) {
	v := m.adddeleted_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *TransactionMutation) ResetDeletedAt() {
	m.deleted_at = nil
	m.adddeleted_at = nil
}

// Where appends a list predicates to the TransactionMutation builder.
func (m *TransactionMutation) Where(ps ...predicate.Transaction) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TransactionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Transaction).
func (m *TransactionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TransactionMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m.name != nil {
		fields = append(fields, transaction.FieldName)
	}
	if m.amount != nil {
		fields = append(fields, transaction.FieldAmount)
	}
	if m.from != nil {
		fields = append(fields, transaction.FieldFrom)
	}
	if m.to != nil {
		fields = append(fields, transaction.FieldTo)
	}
	if m.transaction_id != nil {
		fields = append(fields, transaction.FieldTransactionID)
	}
	if m.cid != nil {
		fields = append(fields, transaction.FieldCid)
	}
	if m.exit_code != nil {
		fields = append(fields, transaction.FieldExitCode)
	}
	if m.status != nil {
		fields = append(fields, transaction.FieldStatus)
	}
	if m.created_at != nil {
		fields = append(fields, transaction.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, transaction.FieldUpdatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, transaction.FieldDeletedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TransactionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldName:
		return m.Name()
	case transaction.FieldAmount:
		return m.Amount()
	case transaction.FieldFrom:
		return m.From()
	case transaction.FieldTo:
		return m.To()
	case transaction.FieldTransactionID:
		return m.TransactionID()
	case transaction.FieldCid:
		return m.Cid()
	case transaction.FieldExitCode:
		return m.ExitCode()
	case transaction.FieldStatus:
		return m.Status()
	case transaction.FieldCreatedAt:
		return m.CreatedAt()
	case transaction.FieldUpdatedAt:
		return m.UpdatedAt()
	case transaction.FieldDeletedAt:
		return m.DeletedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TransactionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case transaction.FieldName:
		return m.OldName(ctx)
	case transaction.FieldAmount:
		return m.OldAmount(ctx)
	case transaction.FieldFrom:
		return m.OldFrom(ctx)
	case transaction.FieldTo:
		return m.OldTo(ctx)
	case transaction.FieldTransactionID:
		return m.OldTransactionID(ctx)
	case transaction.FieldCid:
		return m.OldCid(ctx)
	case transaction.FieldExitCode:
		return m.OldExitCode(ctx)
	case transaction.FieldStatus:
		return m.OldStatus(ctx)
	case transaction.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case transaction.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case transaction.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Transaction field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case transaction.FieldAmount:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case transaction.FieldFrom:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFrom(v)
		return nil
	case transaction.FieldTo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTo(v)
		return nil
	case transaction.FieldTransactionID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTransactionID(v)
		return nil
	case transaction.FieldCid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCid(v)
		return nil
	case transaction.FieldExitCode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExitCode(v)
		return nil
	case transaction.FieldStatus:
		v, ok := value.(transaction.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case transaction.FieldCreatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case transaction.FieldUpdatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case transaction.FieldDeletedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TransactionMutation) AddedFields() []string {
	var fields []string
	if m.addamount != nil {
		fields = append(fields, transaction.FieldAmount)
	}
	if m.addexit_code != nil {
		fields = append(fields, transaction.FieldExitCode)
	}
	if m.addcreated_at != nil {
		fields = append(fields, transaction.FieldCreatedAt)
	}
	if m.addupdated_at != nil {
		fields = append(fields, transaction.FieldUpdatedAt)
	}
	if m.adddeleted_at != nil {
		fields = append(fields, transaction.FieldDeletedAt)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TransactionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldAmount:
		return m.AddedAmount()
	case transaction.FieldExitCode:
		return m.AddedExitCode()
	case transaction.FieldCreatedAt:
		return m.AddedCreatedAt()
	case transaction.FieldUpdatedAt:
		return m.AddedUpdatedAt()
	case transaction.FieldDeletedAt:
		return m.AddedDeletedAt()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) AddField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldAmount:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	case transaction.FieldExitCode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddExitCode(v)
		return nil
	case transaction.FieldCreatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreatedAt(v)
		return nil
	case transaction.FieldUpdatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpdatedAt(v)
		return nil
	case transaction.FieldDeletedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TransactionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TransactionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TransactionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Transaction nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TransactionMutation) ResetField(name string) error {
	switch name {
	case transaction.FieldName:
		m.ResetName()
		return nil
	case transaction.FieldAmount:
		m.ResetAmount()
		return nil
	case transaction.FieldFrom:
		m.ResetFrom()
		return nil
	case transaction.FieldTo:
		m.ResetTo()
		return nil
	case transaction.FieldTransactionID:
		m.ResetTransactionID()
		return nil
	case transaction.FieldCid:
		m.ResetCid()
		return nil
	case transaction.FieldExitCode:
		m.ResetExitCode()
		return nil
	case transaction.FieldStatus:
		m.ResetStatus()
		return nil
	case transaction.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case transaction.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case transaction.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TransactionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TransactionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TransactionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TransactionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TransactionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TransactionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TransactionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Transaction unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TransactionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Transaction edge %s", name)
}
