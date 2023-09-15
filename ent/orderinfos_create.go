// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/orderinfos"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderInfosCreate is the builder for creating a OrderInfos entity.
type OrderInfosCreate struct {
	config
	mutation *OrderInfosMutation
	hooks    []Hook
}

// SetOrderID sets the "order_id" field.
func (oic *OrderInfosCreate) SetOrderID(i int32) *OrderInfosCreate {
	oic.mutation.SetOrderID(i)
	return oic
}

// SetProductID sets the "product_id" field.
func (oic *OrderInfosCreate) SetProductID(i int32) *OrderInfosCreate {
	oic.mutation.SetProductID(i)
	return oic
}

// SetProductName sets the "product_name" field.
func (oic *OrderInfosCreate) SetProductName(s string) *OrderInfosCreate {
	oic.mutation.SetProductName(s)
	return oic
}

// SetProductPrice sets the "product_price" field.
func (oic *OrderInfosCreate) SetProductPrice(i int32) *OrderInfosCreate {
	oic.mutation.SetProductPrice(i)
	return oic
}

// SetProductDescribe sets the "product_describe" field.
func (oic *OrderInfosCreate) SetProductDescribe(s string) *OrderInfosCreate {
	oic.mutation.SetProductDescribe(s)
	return oic
}

// SetCreateTime sets the "create_time" field.
func (oic *OrderInfosCreate) SetCreateTime(t time.Time) *OrderInfosCreate {
	oic.mutation.SetCreateTime(t)
	return oic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oic *OrderInfosCreate) SetNillableCreateTime(t *time.Time) *OrderInfosCreate {
	if t != nil {
		oic.SetCreateTime(*t)
	}
	return oic
}

// SetUpdateTime sets the "update_time" field.
func (oic *OrderInfosCreate) SetUpdateTime(t time.Time) *OrderInfosCreate {
	oic.mutation.SetUpdateTime(t)
	return oic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oic *OrderInfosCreate) SetNillableUpdateTime(t *time.Time) *OrderInfosCreate {
	if t != nil {
		oic.SetUpdateTime(*t)
	}
	return oic
}

// Mutation returns the OrderInfosMutation object of the builder.
func (oic *OrderInfosCreate) Mutation() *OrderInfosMutation {
	return oic.mutation
}

// Save creates the OrderInfos in the database.
func (oic *OrderInfosCreate) Save(ctx context.Context) (*OrderInfos, error) {
	oic.defaults()
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderInfosCreate) SaveX(ctx context.Context) *OrderInfos {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderInfosCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderInfosCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderInfosCreate) defaults() {
	if _, ok := oic.mutation.CreateTime(); !ok {
		v := orderinfos.DefaultCreateTime()
		oic.mutation.SetCreateTime(v)
	}
	if _, ok := oic.mutation.UpdateTime(); !ok {
		v := orderinfos.DefaultUpdateTime()
		oic.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderInfosCreate) check() error {
	if _, ok := oic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderInfos.order_id"`)}
	}
	if _, ok := oic.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "OrderInfos.product_id"`)}
	}
	if _, ok := oic.mutation.ProductName(); !ok {
		return &ValidationError{Name: "product_name", err: errors.New(`ent: missing required field "OrderInfos.product_name"`)}
	}
	if _, ok := oic.mutation.ProductPrice(); !ok {
		return &ValidationError{Name: "product_price", err: errors.New(`ent: missing required field "OrderInfos.product_price"`)}
	}
	if _, ok := oic.mutation.ProductDescribe(); !ok {
		return &ValidationError{Name: "product_describe", err: errors.New(`ent: missing required field "OrderInfos.product_describe"`)}
	}
	if _, ok := oic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "OrderInfos.create_time"`)}
	}
	if _, ok := oic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "OrderInfos.update_time"`)}
	}
	return nil
}

func (oic *OrderInfosCreate) sqlSave(ctx context.Context) (*OrderInfos, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrderInfosCreate) createSpec() (*OrderInfos, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderInfos{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(orderinfos.Table, sqlgraph.NewFieldSpec(orderinfos.FieldID, field.TypeInt))
	)
	if value, ok := oic.mutation.OrderID(); ok {
		_spec.SetField(orderinfos.FieldOrderID, field.TypeInt32, value)
		_node.OrderID = value
	}
	if value, ok := oic.mutation.ProductID(); ok {
		_spec.SetField(orderinfos.FieldProductID, field.TypeInt32, value)
		_node.ProductID = value
	}
	if value, ok := oic.mutation.ProductName(); ok {
		_spec.SetField(orderinfos.FieldProductName, field.TypeString, value)
		_node.ProductName = value
	}
	if value, ok := oic.mutation.ProductPrice(); ok {
		_spec.SetField(orderinfos.FieldProductPrice, field.TypeInt32, value)
		_node.ProductPrice = value
	}
	if value, ok := oic.mutation.ProductDescribe(); ok {
		_spec.SetField(orderinfos.FieldProductDescribe, field.TypeString, value)
		_node.ProductDescribe = value
	}
	if value, ok := oic.mutation.CreateTime(); ok {
		_spec.SetField(orderinfos.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := oic.mutation.UpdateTime(); ok {
		_spec.SetField(orderinfos.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// OrderInfosCreateBulk is the builder for creating many OrderInfos entities in bulk.
type OrderInfosCreateBulk struct {
	config
	builders []*OrderInfosCreate
}

// Save creates the OrderInfos entities in the database.
func (oicb *OrderInfosCreateBulk) Save(ctx context.Context) ([]*OrderInfos, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderInfos, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderInfosMutation)
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
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderInfosCreateBulk) SaveX(ctx context.Context) []*OrderInfos {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderInfosCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderInfosCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
