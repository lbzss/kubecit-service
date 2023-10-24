// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/user"
	"kubecit-service/ent/viporder"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VipOrderCreate is the builder for creating a VipOrder entity.
type VipOrderCreate struct {
	config
	mutation *VipOrderMutation
	hooks    []Hook
}

// SetBizID sets the "biz_id" field.
func (voc *VipOrderCreate) SetBizID(i int64) *VipOrderCreate {
	voc.mutation.SetBizID(i)
	return voc
}

// SetVipType sets the "vip_type" field.
func (voc *VipOrderCreate) SetVipType(i int8) *VipOrderCreate {
	voc.mutation.SetVipType(i)
	return voc
}

// SetPayType sets the "pay_type" field.
func (voc *VipOrderCreate) SetPayType(i int8) *VipOrderCreate {
	voc.mutation.SetPayType(i)
	return voc
}

// SetPayStatus sets the "pay_status" field.
func (voc *VipOrderCreate) SetPayStatus(i int8) *VipOrderCreate {
	voc.mutation.SetPayStatus(i)
	return voc
}

// SetNillablePayStatus sets the "pay_status" field if the given value is not nil.
func (voc *VipOrderCreate) SetNillablePayStatus(i *int8) *VipOrderCreate {
	if i != nil {
		voc.SetPayStatus(*i)
	}
	return voc
}

// SetCreateAt sets the "create_at" field.
func (voc *VipOrderCreate) SetCreateAt(t time.Time) *VipOrderCreate {
	voc.mutation.SetCreateAt(t)
	return voc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (voc *VipOrderCreate) SetNillableCreateAt(t *time.Time) *VipOrderCreate {
	if t != nil {
		voc.SetCreateAt(*t)
	}
	return voc
}

// SetUpdateAt sets the "update_at" field.
func (voc *VipOrderCreate) SetUpdateAt(t time.Time) *VipOrderCreate {
	voc.mutation.SetUpdateAt(t)
	return voc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (voc *VipOrderCreate) SetNillableUpdateAt(t *time.Time) *VipOrderCreate {
	if t != nil {
		voc.SetUpdateAt(*t)
	}
	return voc
}

// SetUserID sets the "user_id" field.
func (voc *VipOrderCreate) SetUserID(i int) *VipOrderCreate {
	voc.mutation.SetUserID(i)
	return voc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (voc *VipOrderCreate) SetNillableUserID(i *int) *VipOrderCreate {
	if i != nil {
		voc.SetUserID(*i)
	}
	return voc
}

// SetPrice sets the "price" field.
func (voc *VipOrderCreate) SetPrice(i int) *VipOrderCreate {
	voc.mutation.SetPrice(i)
	return voc
}

// SetUserOrderID sets the "user_order" edge to the User entity by ID.
func (voc *VipOrderCreate) SetUserOrderID(id int) *VipOrderCreate {
	voc.mutation.SetUserOrderID(id)
	return voc
}

// SetNillableUserOrderID sets the "user_order" edge to the User entity by ID if the given value is not nil.
func (voc *VipOrderCreate) SetNillableUserOrderID(id *int) *VipOrderCreate {
	if id != nil {
		voc = voc.SetUserOrderID(*id)
	}
	return voc
}

// SetUserOrder sets the "user_order" edge to the User entity.
func (voc *VipOrderCreate) SetUserOrder(u *User) *VipOrderCreate {
	return voc.SetUserOrderID(u.ID)
}

// Mutation returns the VipOrderMutation object of the builder.
func (voc *VipOrderCreate) Mutation() *VipOrderMutation {
	return voc.mutation
}

// Save creates the VipOrder in the database.
func (voc *VipOrderCreate) Save(ctx context.Context) (*VipOrder, error) {
	voc.defaults()
	return withHooks(ctx, voc.sqlSave, voc.mutation, voc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (voc *VipOrderCreate) SaveX(ctx context.Context) *VipOrder {
	v, err := voc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (voc *VipOrderCreate) Exec(ctx context.Context) error {
	_, err := voc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (voc *VipOrderCreate) ExecX(ctx context.Context) {
	if err := voc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (voc *VipOrderCreate) defaults() {
	if _, ok := voc.mutation.CreateAt(); !ok {
		v := viporder.DefaultCreateAt()
		voc.mutation.SetCreateAt(v)
	}
	if _, ok := voc.mutation.UpdateAt(); !ok {
		v := viporder.DefaultUpdateAt()
		voc.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (voc *VipOrderCreate) check() error {
	if _, ok := voc.mutation.BizID(); !ok {
		return &ValidationError{Name: "biz_id", err: errors.New(`ent: missing required field "VipOrder.biz_id"`)}
	}
	if _, ok := voc.mutation.VipType(); !ok {
		return &ValidationError{Name: "vip_type", err: errors.New(`ent: missing required field "VipOrder.vip_type"`)}
	}
	if _, ok := voc.mutation.PayType(); !ok {
		return &ValidationError{Name: "pay_type", err: errors.New(`ent: missing required field "VipOrder.pay_type"`)}
	}
	if _, ok := voc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "VipOrder.create_at"`)}
	}
	if _, ok := voc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "VipOrder.update_at"`)}
	}
	if _, ok := voc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "VipOrder.price"`)}
	}
	return nil
}

func (voc *VipOrderCreate) sqlSave(ctx context.Context) (*VipOrder, error) {
	if err := voc.check(); err != nil {
		return nil, err
	}
	_node, _spec := voc.createSpec()
	if err := sqlgraph.CreateNode(ctx, voc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	voc.mutation.id = &_node.ID
	voc.mutation.done = true
	return _node, nil
}

func (voc *VipOrderCreate) createSpec() (*VipOrder, *sqlgraph.CreateSpec) {
	var (
		_node = &VipOrder{config: voc.config}
		_spec = sqlgraph.NewCreateSpec(viporder.Table, sqlgraph.NewFieldSpec(viporder.FieldID, field.TypeInt))
	)
	if value, ok := voc.mutation.BizID(); ok {
		_spec.SetField(viporder.FieldBizID, field.TypeInt64, value)
		_node.BizID = value
	}
	if value, ok := voc.mutation.VipType(); ok {
		_spec.SetField(viporder.FieldVipType, field.TypeInt8, value)
		_node.VipType = value
	}
	if value, ok := voc.mutation.PayType(); ok {
		_spec.SetField(viporder.FieldPayType, field.TypeInt8, value)
		_node.PayType = value
	}
	if value, ok := voc.mutation.PayStatus(); ok {
		_spec.SetField(viporder.FieldPayStatus, field.TypeInt8, value)
		_node.PayStatus = value
	}
	if value, ok := voc.mutation.CreateAt(); ok {
		_spec.SetField(viporder.FieldCreateAt, field.TypeTime, value)
		_node.CreateAt = value
	}
	if value, ok := voc.mutation.UpdateAt(); ok {
		_spec.SetField(viporder.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := voc.mutation.Price(); ok {
		_spec.SetField(viporder.FieldPrice, field.TypeInt, value)
		_node.Price = value
	}
	if nodes := voc.mutation.UserOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   viporder.UserOrderTable,
			Columns: []string{viporder.UserOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VipOrderCreateBulk is the builder for creating many VipOrder entities in bulk.
type VipOrderCreateBulk struct {
	config
	builders []*VipOrderCreate
}

// Save creates the VipOrder entities in the database.
func (vocb *VipOrderCreateBulk) Save(ctx context.Context) ([]*VipOrder, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vocb.builders))
	nodes := make([]*VipOrder, len(vocb.builders))
	mutators := make([]Mutator, len(vocb.builders))
	for i := range vocb.builders {
		func(i int, root context.Context) {
			builder := vocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VipOrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, vocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vocb *VipOrderCreateBulk) SaveX(ctx context.Context) []*VipOrder {
	v, err := vocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vocb *VipOrderCreateBulk) Exec(ctx context.Context) error {
	_, err := vocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vocb *VipOrderCreateBulk) ExecX(ctx context.Context) {
	if err := vocb.Exec(ctx); err != nil {
		panic(err)
	}
}
