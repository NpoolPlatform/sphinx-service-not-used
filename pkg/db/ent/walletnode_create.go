// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sphinx/ent/coininfo"
	"sphinx/ent/walletnode"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WalletNodeCreate is the builder for creating a WalletNode entity.
type WalletNodeCreate struct {
	config
	mutation *WalletNodeMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (wnc *WalletNodeCreate) SetUUID(s string) *WalletNodeCreate {
	wnc.mutation.SetUUID(s)
	return wnc
}

// SetLocation sets the "location" field.
func (wnc *WalletNodeCreate) SetLocation(s string) *WalletNodeCreate {
	wnc.mutation.SetLocation(s)
	return wnc
}

// SetHostVendor sets the "host_vendor" field.
func (wnc *WalletNodeCreate) SetHostVendor(s string) *WalletNodeCreate {
	wnc.mutation.SetHostVendor(s)
	return wnc
}

// SetPublicIP sets the "public_ip" field.
func (wnc *WalletNodeCreate) SetPublicIP(s string) *WalletNodeCreate {
	wnc.mutation.SetPublicIP(s)
	return wnc
}

// SetLocalIP sets the "local_ip" field.
func (wnc *WalletNodeCreate) SetLocalIP(s string) *WalletNodeCreate {
	wnc.mutation.SetLocalIP(s)
	return wnc
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (wnc *WalletNodeCreate) SetCreatetimeUtc(i int) *WalletNodeCreate {
	wnc.mutation.SetCreatetimeUtc(i)
	return wnc
}

// SetLastOnlineTimeUtc sets the "last_online_time_utc" field.
func (wnc *WalletNodeCreate) SetLastOnlineTimeUtc(i int) *WalletNodeCreate {
	wnc.mutation.SetLastOnlineTimeUtc(i)
	return wnc
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (wnc *WalletNodeCreate) SetCoinID(id int) *WalletNodeCreate {
	wnc.mutation.SetCoinID(id)
	return wnc
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (wnc *WalletNodeCreate) SetCoin(c *CoinInfo) *WalletNodeCreate {
	return wnc.SetCoinID(c.ID)
}

// Mutation returns the WalletNodeMutation object of the builder.
func (wnc *WalletNodeCreate) Mutation() *WalletNodeMutation {
	return wnc.mutation
}

// Save creates the WalletNode in the database.
func (wnc *WalletNodeCreate) Save(ctx context.Context) (*WalletNode, error) {
	var (
		err  error
		node *WalletNode
	)
	if len(wnc.hooks) == 0 {
		if err = wnc.check(); err != nil {
			return nil, err
		}
		node, err = wnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wnc.check(); err != nil {
				return nil, err
			}
			wnc.mutation = mutation
			if node, err = wnc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wnc.hooks) - 1; i >= 0; i-- {
			if wnc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wnc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wnc *WalletNodeCreate) SaveX(ctx context.Context) *WalletNode {
	v, err := wnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wnc *WalletNodeCreate) Exec(ctx context.Context) error {
	_, err := wnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnc *WalletNodeCreate) ExecX(ctx context.Context) {
	if err := wnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wnc *WalletNodeCreate) check() error {
	if _, ok := wnc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := wnc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "location"`)}
	}
	if _, ok := wnc.mutation.HostVendor(); !ok {
		return &ValidationError{Name: "host_vendor", err: errors.New(`ent: missing required field "host_vendor"`)}
	}
	if _, ok := wnc.mutation.PublicIP(); !ok {
		return &ValidationError{Name: "public_ip", err: errors.New(`ent: missing required field "public_ip"`)}
	}
	if _, ok := wnc.mutation.LocalIP(); !ok {
		return &ValidationError{Name: "local_ip", err: errors.New(`ent: missing required field "local_ip"`)}
	}
	if _, ok := wnc.mutation.CreatetimeUtc(); !ok {
		return &ValidationError{Name: "createtime_utc", err: errors.New(`ent: missing required field "createtime_utc"`)}
	}
	if _, ok := wnc.mutation.LastOnlineTimeUtc(); !ok {
		return &ValidationError{Name: "last_online_time_utc", err: errors.New(`ent: missing required field "last_online_time_utc"`)}
	}
	if _, ok := wnc.mutation.CoinID(); !ok {
		return &ValidationError{Name: "coin", err: errors.New("ent: missing required edge \"coin\"")}
	}
	return nil
}

func (wnc *WalletNodeCreate) sqlSave(ctx context.Context) (*WalletNode, error) {
	_node, _spec := wnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wnc *WalletNodeCreate) createSpec() (*WalletNode, *sqlgraph.CreateSpec) {
	var (
		_node = &WalletNode{config: wnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: walletnode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: walletnode.FieldID,
			},
		}
	)
	if value, ok := wnc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := wnc.mutation.Location(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldLocation,
		})
		_node.Location = value
	}
	if value, ok := wnc.mutation.HostVendor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldHostVendor,
		})
		_node.HostVendor = value
	}
	if value, ok := wnc.mutation.PublicIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldPublicIP,
		})
		_node.PublicIP = value
	}
	if value, ok := wnc.mutation.LocalIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldLocalIP,
		})
		_node.LocalIP = value
	}
	if value, ok := wnc.mutation.CreatetimeUtc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: walletnode.FieldCreatetimeUtc,
		})
		_node.CreatetimeUtc = value
	}
	if value, ok := wnc.mutation.LastOnlineTimeUtc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: walletnode.FieldLastOnlineTimeUtc,
		})
		_node.LastOnlineTimeUtc = value
	}
	if nodes := wnc.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   walletnode.CoinTable,
			Columns: []string{walletnode.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.coin_info_wallet_nodes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WalletNodeCreateBulk is the builder for creating many WalletNode entities in bulk.
type WalletNodeCreateBulk struct {
	config
	builders []*WalletNodeCreate
}

// Save creates the WalletNode entities in the database.
func (wncb *WalletNodeCreateBulk) Save(ctx context.Context) ([]*WalletNode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wncb.builders))
	nodes := make([]*WalletNode, len(wncb.builders))
	mutators := make([]Mutator, len(wncb.builders))
	for i := range wncb.builders {
		func(i int, root context.Context) {
			builder := wncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WalletNodeMutation)
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
					_, err = mutators[i+1].Mutate(root, wncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wncb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wncb *WalletNodeCreateBulk) SaveX(ctx context.Context) []*WalletNode {
	v, err := wncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wncb *WalletNodeCreateBulk) Exec(ctx context.Context) error {
	_, err := wncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wncb *WalletNodeCreateBulk) ExecX(ctx context.Context) {
	if err := wncb.Exec(ctx); err != nil {
		panic(err)
	}
}