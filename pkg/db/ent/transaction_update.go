// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/google/uuid"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetAmountUint64 sets the "amount_uint64" field.
func (tu *TransactionUpdate) SetAmountUint64(u uint64) *TransactionUpdate {
	tu.mutation.ResetAmountUint64()
	tu.mutation.SetAmountUint64(u)
	return tu
}

// AddAmountUint64 adds u to the "amount_uint64" field.
func (tu *TransactionUpdate) AddAmountUint64(u uint64) *TransactionUpdate {
	tu.mutation.AddAmountUint64(u)
	return tu
}

// SetAmountFloat64 sets the "amount_float64" field.
func (tu *TransactionUpdate) SetAmountFloat64(f float64) *TransactionUpdate {
	tu.mutation.ResetAmountFloat64()
	tu.mutation.SetAmountFloat64(f)
	return tu
}

// AddAmountFloat64 adds f to the "amount_float64" field.
func (tu *TransactionUpdate) AddAmountFloat64(f float64) *TransactionUpdate {
	tu.mutation.AddAmountFloat64(f)
	return tu
}

// SetAddressFrom sets the "address_from" field.
func (tu *TransactionUpdate) SetAddressFrom(s string) *TransactionUpdate {
	tu.mutation.SetAddressFrom(s)
	return tu
}

// SetAddressTo sets the "address_to" field.
func (tu *TransactionUpdate) SetAddressTo(s string) *TransactionUpdate {
	tu.mutation.SetAddressTo(s)
	return tu
}

// SetNeedManualReview sets the "need_manual_review" field.
func (tu *TransactionUpdate) SetNeedManualReview(b bool) *TransactionUpdate {
	tu.mutation.SetNeedManualReview(b)
	return tu
}

// SetNillableNeedManualReview sets the "need_manual_review" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableNeedManualReview(b *bool) *TransactionUpdate {
	if b != nil {
		tu.SetNeedManualReview(*b)
	}
	return tu
}

// SetType sets the "type" field.
func (tu *TransactionUpdate) SetType(t transaction.Type) *TransactionUpdate {
	tu.mutation.SetType(t)
	return tu
}

// SetTransactionIDInsite sets the "transaction_id_insite" field.
func (tu *TransactionUpdate) SetTransactionIDInsite(s string) *TransactionUpdate {
	tu.mutation.SetTransactionIDInsite(s)
	return tu
}

// SetTransactionIDChain sets the "transaction_id_chain" field.
func (tu *TransactionUpdate) SetTransactionIDChain(s string) *TransactionUpdate {
	tu.mutation.SetTransactionIDChain(s)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TransactionUpdate) SetStatus(t transaction.Status) *TransactionUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetMutex sets the "mutex" field.
func (tu *TransactionUpdate) SetMutex(b bool) *TransactionUpdate {
	tu.mutation.SetMutex(b)
	return tu
}

// SetNillableMutex sets the "mutex" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableMutex(b *bool) *TransactionUpdate {
	if b != nil {
		tu.SetMutex(*b)
	}
	return tu
}

// SetSignatureUser sets the "signature_user" field.
func (tu *TransactionUpdate) SetSignatureUser(s string) *TransactionUpdate {
	tu.mutation.SetSignatureUser(s)
	return tu
}

// SetSignaturePlatform sets the "signature_platform" field.
func (tu *TransactionUpdate) SetSignaturePlatform(s string) *TransactionUpdate {
	tu.mutation.SetSignaturePlatform(s)
	return tu
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (tu *TransactionUpdate) SetCreatetimeUtc(i int64) *TransactionUpdate {
	tu.mutation.ResetCreatetimeUtc()
	tu.mutation.SetCreatetimeUtc(i)
	return tu
}

// AddCreatetimeUtc adds i to the "createtime_utc" field.
func (tu *TransactionUpdate) AddCreatetimeUtc(i int64) *TransactionUpdate {
	tu.mutation.AddCreatetimeUtc(i)
	return tu
}

// SetUpdatetimeUtc sets the "updatetime_utc" field.
func (tu *TransactionUpdate) SetUpdatetimeUtc(i int64) *TransactionUpdate {
	tu.mutation.ResetUpdatetimeUtc()
	tu.mutation.SetUpdatetimeUtc(i)
	return tu
}

// AddUpdatetimeUtc adds i to the "updatetime_utc" field.
func (tu *TransactionUpdate) AddUpdatetimeUtc(i int64) *TransactionUpdate {
	tu.mutation.AddUpdatetimeUtc(i)
	return tu
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (tu *TransactionUpdate) SetCoinID(id uuid.UUID) *TransactionUpdate {
	tu.mutation.SetCoinID(id)
	return tu
}

// SetNillableCoinID sets the "coin" edge to the CoinInfo entity by ID if the given value is not nil.
func (tu *TransactionUpdate) SetNillableCoinID(id *uuid.UUID) *TransactionUpdate {
	if id != nil {
		tu = tu.SetCoinID(*id)
	}
	return tu
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (tu *TransactionUpdate) SetCoin(c *CoinInfo) *TransactionUpdate {
	return tu.SetCoinID(c.ID)
}

// AddReviewIDs adds the "review" edge to the Review entity by IDs.
func (tu *TransactionUpdate) AddReviewIDs(ids ...int32) *TransactionUpdate {
	tu.mutation.AddReviewIDs(ids...)
	return tu
}

// AddReview adds the "review" edges to the Review entity.
func (tu *TransactionUpdate) AddReview(r ...*Review) *TransactionUpdate {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddReviewIDs(ids...)
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// ClearCoin clears the "coin" edge to the CoinInfo entity.
func (tu *TransactionUpdate) ClearCoin() *TransactionUpdate {
	tu.mutation.ClearCoin()
	return tu
}

// ClearReview clears all "review" edges to the Review entity.
func (tu *TransactionUpdate) ClearReview() *TransactionUpdate {
	tu.mutation.ClearReview()
	return tu
}

// RemoveReviewIDs removes the "review" edge to Review entities by IDs.
func (tu *TransactionUpdate) RemoveReviewIDs(ids ...int32) *TransactionUpdate {
	tu.mutation.RemoveReviewIDs(ids...)
	return tu
}

// RemoveReview removes "review" edges to Review entities.
func (tu *TransactionUpdate) RemoveReview(r ...*Review) *TransactionUpdate {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveReviewIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TransactionUpdate) check() error {
	if v, ok := tu.mutation.AddressFrom(); ok {
		if err := transaction.AddressFromValidator(v); err != nil {
			return &ValidationError{Name: "address_from", err: fmt.Errorf("ent: validator failed for field \"address_from\": %w", err)}
		}
	}
	if v, ok := tu.mutation.AddressTo(); ok {
		if err := transaction.AddressToValidator(v); err != nil {
			return &ValidationError{Name: "address_to", err: fmt.Errorf("ent: validator failed for field \"address_to\": %w", err)}
		}
	}
	if v, ok := tu.mutation.GetType(); ok {
		if err := transaction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := tu.mutation.TransactionIDInsite(); ok {
		if err := transaction.TransactionIDInsiteValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id_insite", err: fmt.Errorf("ent: validator failed for field \"transaction_id_insite\": %w", err)}
		}
	}
	if v, ok := tu.mutation.TransactionIDChain(); ok {
		if err := transaction.TransactionIDChainValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id_chain", err: fmt.Errorf("ent: validator failed for field \"transaction_id_chain\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Status(); ok {
		if err := transaction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := tu.mutation.SignatureUser(); ok {
		if err := transaction.SignatureUserValidator(v); err != nil {
			return &ValidationError{Name: "signature_user", err: fmt.Errorf("ent: validator failed for field \"signature_user\": %w", err)}
		}
	}
	if v, ok := tu.mutation.SignaturePlatform(); ok {
		if err := transaction.SignaturePlatformValidator(v); err != nil {
			return &ValidationError{Name: "signature_platform", err: fmt.Errorf("ent: validator failed for field \"signature_platform\": %w", err)}
		}
	}
	return nil
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: transaction.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.AmountUint64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmountUint64,
		})
	}
	if value, ok := tu.mutation.AddedAmountUint64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmountUint64,
		})
	}
	if value, ok := tu.mutation.AmountFloat64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transaction.FieldAmountFloat64,
		})
	}
	if value, ok := tu.mutation.AddedAmountFloat64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transaction.FieldAmountFloat64,
		})
	}
	if value, ok := tu.mutation.AddressFrom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldAddressFrom,
		})
	}
	if value, ok := tu.mutation.AddressTo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldAddressTo,
		})
	}
	if value, ok := tu.mutation.NeedManualReview(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: transaction.FieldNeedManualReview,
		})
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transaction.FieldType,
		})
	}
	if value, ok := tu.mutation.TransactionIDInsite(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionIDInsite,
		})
	}
	if value, ok := tu.mutation.TransactionIDChain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionIDChain,
		})
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transaction.FieldStatus,
		})
	}
	if value, ok := tu.mutation.Mutex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: transaction.FieldMutex,
		})
	}
	if value, ok := tu.mutation.SignatureUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldSignatureUser,
		})
	}
	if value, ok := tu.mutation.SignaturePlatform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldSignaturePlatform,
		})
	}
	if value, ok := tu.mutation.CreatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldCreatetimeUtc,
		})
	}
	if value, ok := tu.mutation.AddedCreatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldCreatetimeUtc,
		})
	}
	if value, ok := tu.mutation.UpdatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldUpdatetimeUtc,
		})
	}
	if value, ok := tu.mutation.AddedUpdatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldUpdatetimeUtc,
		})
	}
	if tu.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.CoinTable,
			Columns: []string{transaction.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.CoinTable,
			Columns: []string{transaction.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   transaction.ReviewTable,
			Columns: transaction.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReviewIDs(); len(nodes) > 0 && !tu.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   transaction.ReviewTable,
			Columns: transaction.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   transaction.ReviewTable,
			Columns: transaction.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetAmountUint64 sets the "amount_uint64" field.
func (tuo *TransactionUpdateOne) SetAmountUint64(u uint64) *TransactionUpdateOne {
	tuo.mutation.ResetAmountUint64()
	tuo.mutation.SetAmountUint64(u)
	return tuo
}

// AddAmountUint64 adds u to the "amount_uint64" field.
func (tuo *TransactionUpdateOne) AddAmountUint64(u uint64) *TransactionUpdateOne {
	tuo.mutation.AddAmountUint64(u)
	return tuo
}

// SetAmountFloat64 sets the "amount_float64" field.
func (tuo *TransactionUpdateOne) SetAmountFloat64(f float64) *TransactionUpdateOne {
	tuo.mutation.ResetAmountFloat64()
	tuo.mutation.SetAmountFloat64(f)
	return tuo
}

// AddAmountFloat64 adds f to the "amount_float64" field.
func (tuo *TransactionUpdateOne) AddAmountFloat64(f float64) *TransactionUpdateOne {
	tuo.mutation.AddAmountFloat64(f)
	return tuo
}

// SetAddressFrom sets the "address_from" field.
func (tuo *TransactionUpdateOne) SetAddressFrom(s string) *TransactionUpdateOne {
	tuo.mutation.SetAddressFrom(s)
	return tuo
}

// SetAddressTo sets the "address_to" field.
func (tuo *TransactionUpdateOne) SetAddressTo(s string) *TransactionUpdateOne {
	tuo.mutation.SetAddressTo(s)
	return tuo
}

// SetNeedManualReview sets the "need_manual_review" field.
func (tuo *TransactionUpdateOne) SetNeedManualReview(b bool) *TransactionUpdateOne {
	tuo.mutation.SetNeedManualReview(b)
	return tuo
}

// SetNillableNeedManualReview sets the "need_manual_review" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableNeedManualReview(b *bool) *TransactionUpdateOne {
	if b != nil {
		tuo.SetNeedManualReview(*b)
	}
	return tuo
}

// SetType sets the "type" field.
func (tuo *TransactionUpdateOne) SetType(t transaction.Type) *TransactionUpdateOne {
	tuo.mutation.SetType(t)
	return tuo
}

// SetTransactionIDInsite sets the "transaction_id_insite" field.
func (tuo *TransactionUpdateOne) SetTransactionIDInsite(s string) *TransactionUpdateOne {
	tuo.mutation.SetTransactionIDInsite(s)
	return tuo
}

// SetTransactionIDChain sets the "transaction_id_chain" field.
func (tuo *TransactionUpdateOne) SetTransactionIDChain(s string) *TransactionUpdateOne {
	tuo.mutation.SetTransactionIDChain(s)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TransactionUpdateOne) SetStatus(t transaction.Status) *TransactionUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetMutex sets the "mutex" field.
func (tuo *TransactionUpdateOne) SetMutex(b bool) *TransactionUpdateOne {
	tuo.mutation.SetMutex(b)
	return tuo
}

// SetNillableMutex sets the "mutex" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableMutex(b *bool) *TransactionUpdateOne {
	if b != nil {
		tuo.SetMutex(*b)
	}
	return tuo
}

// SetSignatureUser sets the "signature_user" field.
func (tuo *TransactionUpdateOne) SetSignatureUser(s string) *TransactionUpdateOne {
	tuo.mutation.SetSignatureUser(s)
	return tuo
}

// SetSignaturePlatform sets the "signature_platform" field.
func (tuo *TransactionUpdateOne) SetSignaturePlatform(s string) *TransactionUpdateOne {
	tuo.mutation.SetSignaturePlatform(s)
	return tuo
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (tuo *TransactionUpdateOne) SetCreatetimeUtc(i int64) *TransactionUpdateOne {
	tuo.mutation.ResetCreatetimeUtc()
	tuo.mutation.SetCreatetimeUtc(i)
	return tuo
}

// AddCreatetimeUtc adds i to the "createtime_utc" field.
func (tuo *TransactionUpdateOne) AddCreatetimeUtc(i int64) *TransactionUpdateOne {
	tuo.mutation.AddCreatetimeUtc(i)
	return tuo
}

// SetUpdatetimeUtc sets the "updatetime_utc" field.
func (tuo *TransactionUpdateOne) SetUpdatetimeUtc(i int64) *TransactionUpdateOne {
	tuo.mutation.ResetUpdatetimeUtc()
	tuo.mutation.SetUpdatetimeUtc(i)
	return tuo
}

// AddUpdatetimeUtc adds i to the "updatetime_utc" field.
func (tuo *TransactionUpdateOne) AddUpdatetimeUtc(i int64) *TransactionUpdateOne {
	tuo.mutation.AddUpdatetimeUtc(i)
	return tuo
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (tuo *TransactionUpdateOne) SetCoinID(id uuid.UUID) *TransactionUpdateOne {
	tuo.mutation.SetCoinID(id)
	return tuo
}

// SetNillableCoinID sets the "coin" edge to the CoinInfo entity by ID if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableCoinID(id *uuid.UUID) *TransactionUpdateOne {
	if id != nil {
		tuo = tuo.SetCoinID(*id)
	}
	return tuo
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (tuo *TransactionUpdateOne) SetCoin(c *CoinInfo) *TransactionUpdateOne {
	return tuo.SetCoinID(c.ID)
}

// AddReviewIDs adds the "review" edge to the Review entity by IDs.
func (tuo *TransactionUpdateOne) AddReviewIDs(ids ...int32) *TransactionUpdateOne {
	tuo.mutation.AddReviewIDs(ids...)
	return tuo
}

// AddReview adds the "review" edges to the Review entity.
func (tuo *TransactionUpdateOne) AddReview(r ...*Review) *TransactionUpdateOne {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddReviewIDs(ids...)
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// ClearCoin clears the "coin" edge to the CoinInfo entity.
func (tuo *TransactionUpdateOne) ClearCoin() *TransactionUpdateOne {
	tuo.mutation.ClearCoin()
	return tuo
}

// ClearReview clears all "review" edges to the Review entity.
func (tuo *TransactionUpdateOne) ClearReview() *TransactionUpdateOne {
	tuo.mutation.ClearReview()
	return tuo
}

// RemoveReviewIDs removes the "review" edge to Review entities by IDs.
func (tuo *TransactionUpdateOne) RemoveReviewIDs(ids ...int32) *TransactionUpdateOne {
	tuo.mutation.RemoveReviewIDs(ids...)
	return tuo
}

// RemoveReview removes "review" edges to Review entities.
func (tuo *TransactionUpdateOne) RemoveReview(r ...*Review) *TransactionUpdateOne {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveReviewIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TransactionUpdateOne) check() error {
	if v, ok := tuo.mutation.AddressFrom(); ok {
		if err := transaction.AddressFromValidator(v); err != nil {
			return &ValidationError{Name: "address_from", err: fmt.Errorf("ent: validator failed for field \"address_from\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.AddressTo(); ok {
		if err := transaction.AddressToValidator(v); err != nil {
			return &ValidationError{Name: "address_to", err: fmt.Errorf("ent: validator failed for field \"address_to\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.GetType(); ok {
		if err := transaction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.TransactionIDInsite(); ok {
		if err := transaction.TransactionIDInsiteValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id_insite", err: fmt.Errorf("ent: validator failed for field \"transaction_id_insite\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.TransactionIDChain(); ok {
		if err := transaction.TransactionIDChainValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id_chain", err: fmt.Errorf("ent: validator failed for field \"transaction_id_chain\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Status(); ok {
		if err := transaction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.SignatureUser(); ok {
		if err := transaction.SignatureUserValidator(v); err != nil {
			return &ValidationError{Name: "signature_user", err: fmt.Errorf("ent: validator failed for field \"signature_user\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.SignaturePlatform(); ok {
		if err := transaction.SignaturePlatformValidator(v); err != nil {
			return &ValidationError{Name: "signature_platform", err: fmt.Errorf("ent: validator failed for field \"signature_platform\": %w", err)}
		}
	}
	return nil
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: transaction.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Transaction.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.AmountUint64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmountUint64,
		})
	}
	if value, ok := tuo.mutation.AddedAmountUint64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmountUint64,
		})
	}
	if value, ok := tuo.mutation.AmountFloat64(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transaction.FieldAmountFloat64,
		})
	}
	if value, ok := tuo.mutation.AddedAmountFloat64(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transaction.FieldAmountFloat64,
		})
	}
	if value, ok := tuo.mutation.AddressFrom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldAddressFrom,
		})
	}
	if value, ok := tuo.mutation.AddressTo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldAddressTo,
		})
	}
	if value, ok := tuo.mutation.NeedManualReview(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: transaction.FieldNeedManualReview,
		})
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transaction.FieldType,
		})
	}
	if value, ok := tuo.mutation.TransactionIDInsite(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionIDInsite,
		})
	}
	if value, ok := tuo.mutation.TransactionIDChain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionIDChain,
		})
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transaction.FieldStatus,
		})
	}
	if value, ok := tuo.mutation.Mutex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: transaction.FieldMutex,
		})
	}
	if value, ok := tuo.mutation.SignatureUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldSignatureUser,
		})
	}
	if value, ok := tuo.mutation.SignaturePlatform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldSignaturePlatform,
		})
	}
	if value, ok := tuo.mutation.CreatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldCreatetimeUtc,
		})
	}
	if value, ok := tuo.mutation.AddedCreatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldCreatetimeUtc,
		})
	}
	if value, ok := tuo.mutation.UpdatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldUpdatetimeUtc,
		})
	}
	if value, ok := tuo.mutation.AddedUpdatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldUpdatetimeUtc,
		})
	}
	if tuo.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.CoinTable,
			Columns: []string{transaction.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.CoinTable,
			Columns: []string{transaction.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   transaction.ReviewTable,
			Columns: transaction.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReviewIDs(); len(nodes) > 0 && !tuo.mutation.ReviewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   transaction.ReviewTable,
			Columns: transaction.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   transaction.ReviewTable,
			Columns: transaction.ReviewPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
