// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/google/uuid"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (tc *TransactionCreate) SetName(s string) *TransactionCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableName(s *string) *TransactionCreate {
	if s != nil {
		tc.SetName(*s)
	}
	return tc
}

// SetAmount sets the "amount" field.
func (tc *TransactionCreate) SetAmount(u uint64) *TransactionCreate {
	tc.mutation.SetAmount(u)
	return tc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableAmount(u *uint64) *TransactionCreate {
	if u != nil {
		tc.SetAmount(*u)
	}
	return tc
}

// SetFrom sets the "from" field.
func (tc *TransactionCreate) SetFrom(s string) *TransactionCreate {
	tc.mutation.SetFrom(s)
	return tc
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableFrom(s *string) *TransactionCreate {
	if s != nil {
		tc.SetFrom(*s)
	}
	return tc
}

// SetTo sets the "to" field.
func (tc *TransactionCreate) SetTo(s string) *TransactionCreate {
	tc.mutation.SetTo(s)
	return tc
}

// SetNillableTo sets the "to" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableTo(s *string) *TransactionCreate {
	if s != nil {
		tc.SetTo(*s)
	}
	return tc
}

// SetTransactionID sets the "transaction_id" field.
func (tc *TransactionCreate) SetTransactionID(s string) *TransactionCreate {
	tc.mutation.SetTransactionID(s)
	return tc
}

// SetCid sets the "cid" field.
func (tc *TransactionCreate) SetCid(s string) *TransactionCreate {
	tc.mutation.SetCid(s)
	return tc
}

// SetNillableCid sets the "cid" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCid(s *string) *TransactionCreate {
	if s != nil {
		tc.SetCid(*s)
	}
	return tc
}

// SetExitCode sets the "exit_code" field.
func (tc *TransactionCreate) SetExitCode(i int64) *TransactionCreate {
	tc.mutation.SetExitCode(i)
	return tc
}

// SetNillableExitCode sets the "exit_code" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableExitCode(i *int64) *TransactionCreate {
	if i != nil {
		tc.SetExitCode(*i)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TransactionCreate) SetStatus(t transaction.Status) *TransactionCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TransactionCreate) SetCreatedAt(u uint32) *TransactionCreate {
	tc.mutation.SetCreatedAt(u)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCreatedAt(u *uint32) *TransactionCreate {
	if u != nil {
		tc.SetCreatedAt(*u)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TransactionCreate) SetUpdatedAt(u uint32) *TransactionCreate {
	tc.mutation.SetUpdatedAt(u)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableUpdatedAt(u *uint32) *TransactionCreate {
	if u != nil {
		tc.SetUpdatedAt(*u)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TransactionCreate) SetDeletedAt(u uint32) *TransactionCreate {
	tc.mutation.SetDeletedAt(u)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableDeletedAt(u *uint32) *TransactionCreate {
	if u != nil {
		tc.SetDeletedAt(*u)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetID(u)
	return tc
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransactionCreate) defaults() {
	if _, ok := tc.mutation.Name(); !ok {
		v := transaction.DefaultName
		tc.mutation.SetName(v)
	}
	if _, ok := tc.mutation.Amount(); !ok {
		v := transaction.DefaultAmount
		tc.mutation.SetAmount(v)
	}
	if _, ok := tc.mutation.From(); !ok {
		v := transaction.DefaultFrom
		tc.mutation.SetFrom(v)
	}
	if _, ok := tc.mutation.To(); !ok {
		v := transaction.DefaultTo
		tc.mutation.SetTo(v)
	}
	if _, ok := tc.mutation.Cid(); !ok {
		v := transaction.DefaultCid
		tc.mutation.SetCid(v)
	}
	if _, ok := tc.mutation.ExitCode(); !ok {
		v := transaction.DefaultExitCode
		tc.mutation.SetExitCode(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := transaction.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := transaction.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		v := transaction.DefaultDeletedAt()
		tc.mutation.SetDeletedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := transaction.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := transaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if _, ok := tc.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "from"`)}
	}
	if v, ok := tc.mutation.From(); ok {
		if err := transaction.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "from": %w`, err)}
		}
	}
	if _, ok := tc.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "to"`)}
	}
	if v, ok := tc.mutation.To(); ok {
		if err := transaction.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "to": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction_id", err: errors.New(`ent: missing required field "transaction_id"`)}
	}
	if v, ok := tc.mutation.TransactionID(); ok {
		if err := transaction.TransactionIDValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id", err: fmt.Errorf(`ent: validator failed for field "transaction_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Cid(); !ok {
		return &ValidationError{Name: "cid", err: errors.New(`ent: missing required field "cid"`)}
	}
	if _, ok := tc.mutation.ExitCode(); !ok {
		return &ValidationError{Name: "exit_code", err: errors.New(`ent: missing required field "exit_code"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := transaction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := tc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "deleted_at"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := tc.mutation.From(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldFrom,
		})
		_node.From = value
	}
	if value, ok := tc.mutation.To(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTo,
		})
		_node.To = value
	}
	if value, ok := tc.mutation.TransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionID,
		})
		_node.TransactionID = value
	}
	if value, ok := tc.mutation.Cid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldCid,
		})
		_node.Cid = value
	}
	if value, ok := tc.mutation.ExitCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldExitCode,
		})
		_node.ExitCode = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transaction.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Transaction.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TransactionUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (tc *TransactionCreate) OnConflict(opts ...sql.ConflictOption) *TransactionUpsertOne {
	tc.conflict = opts
	return &TransactionUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tc *TransactionCreate) OnConflictColumns(columns ...string) *TransactionUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TransactionUpsertOne{
		create: tc,
	}
}

type (
	// TransactionUpsertOne is the builder for "upsert"-ing
	//  one Transaction node.
	TransactionUpsertOne struct {
		create *TransactionCreate
	}

	// TransactionUpsert is the "OnConflict" setter.
	TransactionUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TransactionUpsert) SetName(v string) *TransactionUpsert {
	u.Set(transaction.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateName() *TransactionUpsert {
	u.SetExcluded(transaction.FieldName)
	return u
}

// SetAmount sets the "amount" field.
func (u *TransactionUpsert) SetAmount(v uint64) *TransactionUpsert {
	u.Set(transaction.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateAmount() *TransactionUpsert {
	u.SetExcluded(transaction.FieldAmount)
	return u
}

// SetFrom sets the "from" field.
func (u *TransactionUpsert) SetFrom(v string) *TransactionUpsert {
	u.Set(transaction.FieldFrom, v)
	return u
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateFrom() *TransactionUpsert {
	u.SetExcluded(transaction.FieldFrom)
	return u
}

// SetTo sets the "to" field.
func (u *TransactionUpsert) SetTo(v string) *TransactionUpsert {
	u.Set(transaction.FieldTo, v)
	return u
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateTo() *TransactionUpsert {
	u.SetExcluded(transaction.FieldTo)
	return u
}

// SetTransactionID sets the "transaction_id" field.
func (u *TransactionUpsert) SetTransactionID(v string) *TransactionUpsert {
	u.Set(transaction.FieldTransactionID, v)
	return u
}

// UpdateTransactionID sets the "transaction_id" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateTransactionID() *TransactionUpsert {
	u.SetExcluded(transaction.FieldTransactionID)
	return u
}

// SetCid sets the "cid" field.
func (u *TransactionUpsert) SetCid(v string) *TransactionUpsert {
	u.Set(transaction.FieldCid, v)
	return u
}

// UpdateCid sets the "cid" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateCid() *TransactionUpsert {
	u.SetExcluded(transaction.FieldCid)
	return u
}

// SetExitCode sets the "exit_code" field.
func (u *TransactionUpsert) SetExitCode(v int64) *TransactionUpsert {
	u.Set(transaction.FieldExitCode, v)
	return u
}

// UpdateExitCode sets the "exit_code" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateExitCode() *TransactionUpsert {
	u.SetExcluded(transaction.FieldExitCode)
	return u
}

// SetStatus sets the "status" field.
func (u *TransactionUpsert) SetStatus(v transaction.Status) *TransactionUpsert {
	u.Set(transaction.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateStatus() *TransactionUpsert {
	u.SetExcluded(transaction.FieldStatus)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TransactionUpsert) SetCreatedAt(v uint32) *TransactionUpsert {
	u.Set(transaction.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateCreatedAt() *TransactionUpsert {
	u.SetExcluded(transaction.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TransactionUpsert) SetUpdatedAt(v uint32) *TransactionUpsert {
	u.Set(transaction.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateUpdatedAt() *TransactionUpsert {
	u.SetExcluded(transaction.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TransactionUpsert) SetDeletedAt(v uint32) *TransactionUpsert {
	u.Set(transaction.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateDeletedAt() *TransactionUpsert {
	u.SetExcluded(transaction.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transaction.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TransactionUpsertOne) UpdateNewValues() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(transaction.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Transaction.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TransactionUpsertOne) Ignore() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TransactionUpsertOne) DoNothing() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TransactionCreate.OnConflict
// documentation for more info.
func (u *TransactionUpsertOne) Update(set func(*TransactionUpsert)) *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TransactionUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TransactionUpsertOne) SetName(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateName() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateName()
	})
}

// SetAmount sets the "amount" field.
func (u *TransactionUpsertOne) SetAmount(v uint64) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateAmount() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateAmount()
	})
}

// SetFrom sets the "from" field.
func (u *TransactionUpsertOne) SetFrom(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateFrom() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateFrom()
	})
}

// SetTo sets the "to" field.
func (u *TransactionUpsertOne) SetTo(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTo(v)
	})
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateTo() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTo()
	})
}

// SetTransactionID sets the "transaction_id" field.
func (u *TransactionUpsertOne) SetTransactionID(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTransactionID(v)
	})
}

// UpdateTransactionID sets the "transaction_id" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateTransactionID() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTransactionID()
	})
}

// SetCid sets the "cid" field.
func (u *TransactionUpsertOne) SetCid(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCid(v)
	})
}

// UpdateCid sets the "cid" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateCid() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCid()
	})
}

// SetExitCode sets the "exit_code" field.
func (u *TransactionUpsertOne) SetExitCode(v int64) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetExitCode(v)
	})
}

// UpdateExitCode sets the "exit_code" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateExitCode() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateExitCode()
	})
}

// SetStatus sets the "status" field.
func (u *TransactionUpsertOne) SetStatus(v transaction.Status) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateStatus() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateStatus()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TransactionUpsertOne) SetCreatedAt(v uint32) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateCreatedAt() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TransactionUpsertOne) SetUpdatedAt(v uint32) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateUpdatedAt() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TransactionUpsertOne) SetDeletedAt(v uint32) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateDeletedAt() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateDeletedAt()
	})
}

// Exec executes the query.
func (u *TransactionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TransactionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TransactionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TransactionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TransactionUpsertOne.ID is not supported by MySQL driver. Use TransactionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TransactionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
	conflict []sql.ConflictOption
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Transaction.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TransactionUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (tcb *TransactionCreateBulk) OnConflict(opts ...sql.ConflictOption) *TransactionUpsertBulk {
	tcb.conflict = opts
	return &TransactionUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcb *TransactionCreateBulk) OnConflictColumns(columns ...string) *TransactionUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TransactionUpsertBulk{
		create: tcb,
	}
}

// TransactionUpsertBulk is the builder for "upsert"-ing
// a bulk of Transaction nodes.
type TransactionUpsertBulk struct {
	create *TransactionCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transaction.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TransactionUpsertBulk) UpdateNewValues() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(transaction.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TransactionUpsertBulk) Ignore() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TransactionUpsertBulk) DoNothing() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TransactionCreateBulk.OnConflict
// documentation for more info.
func (u *TransactionUpsertBulk) Update(set func(*TransactionUpsert)) *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TransactionUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TransactionUpsertBulk) SetName(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateName() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateName()
	})
}

// SetAmount sets the "amount" field.
func (u *TransactionUpsertBulk) SetAmount(v uint64) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateAmount() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateAmount()
	})
}

// SetFrom sets the "from" field.
func (u *TransactionUpsertBulk) SetFrom(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateFrom() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateFrom()
	})
}

// SetTo sets the "to" field.
func (u *TransactionUpsertBulk) SetTo(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTo(v)
	})
}

// UpdateTo sets the "to" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateTo() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTo()
	})
}

// SetTransactionID sets the "transaction_id" field.
func (u *TransactionUpsertBulk) SetTransactionID(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTransactionID(v)
	})
}

// UpdateTransactionID sets the "transaction_id" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateTransactionID() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTransactionID()
	})
}

// SetCid sets the "cid" field.
func (u *TransactionUpsertBulk) SetCid(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCid(v)
	})
}

// UpdateCid sets the "cid" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateCid() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCid()
	})
}

// SetExitCode sets the "exit_code" field.
func (u *TransactionUpsertBulk) SetExitCode(v int64) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetExitCode(v)
	})
}

// UpdateExitCode sets the "exit_code" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateExitCode() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateExitCode()
	})
}

// SetStatus sets the "status" field.
func (u *TransactionUpsertBulk) SetStatus(v transaction.Status) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateStatus() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateStatus()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TransactionUpsertBulk) SetCreatedAt(v uint32) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateCreatedAt() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TransactionUpsertBulk) SetUpdatedAt(v uint32) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateUpdatedAt() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TransactionUpsertBulk) SetDeletedAt(v uint32) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateDeletedAt() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateDeletedAt()
	})
}

// Exec executes the query.
func (u *TransactionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TransactionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TransactionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TransactionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
