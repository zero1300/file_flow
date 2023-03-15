// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"file_flow/ent/centralstoragepool"
	"file_flow/ent/predicate"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CentralStoragePoolUpdate is the builder for updating CentralStoragePool entities.
type CentralStoragePoolUpdate struct {
	config
	hooks    []Hook
	mutation *CentralStoragePoolMutation
}

// Where appends a list predicates to the CentralStoragePoolUpdate builder.
func (cspu *CentralStoragePoolUpdate) Where(ps ...predicate.CentralStoragePool) *CentralStoragePoolUpdate {
	cspu.mutation.Where(ps...)
	return cspu
}

// SetDeleteAt sets the "delete_at" field.
func (cspu *CentralStoragePoolUpdate) SetDeleteAt(t time.Time) *CentralStoragePoolUpdate {
	cspu.mutation.SetDeleteAt(t)
	return cspu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cspu *CentralStoragePoolUpdate) SetNillableDeleteAt(t *time.Time) *CentralStoragePoolUpdate {
	if t != nil {
		cspu.SetDeleteAt(*t)
	}
	return cspu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (cspu *CentralStoragePoolUpdate) ClearDeleteAt() *CentralStoragePoolUpdate {
	cspu.mutation.ClearDeleteAt()
	return cspu
}

// SetFilename sets the "filename" field.
func (cspu *CentralStoragePoolUpdate) SetFilename(s string) *CentralStoragePoolUpdate {
	cspu.mutation.SetFilename(s)
	return cspu
}

// SetExt sets the "ext" field.
func (cspu *CentralStoragePoolUpdate) SetExt(s string) *CentralStoragePoolUpdate {
	cspu.mutation.SetExt(s)
	return cspu
}

// SetSize sets the "size" field.
func (cspu *CentralStoragePoolUpdate) SetSize(f float64) *CentralStoragePoolUpdate {
	cspu.mutation.ResetSize()
	cspu.mutation.SetSize(f)
	return cspu
}

// AddSize adds f to the "size" field.
func (cspu *CentralStoragePoolUpdate) AddSize(f float64) *CentralStoragePoolUpdate {
	cspu.mutation.AddSize(f)
	return cspu
}

// SetPath sets the "path" field.
func (cspu *CentralStoragePoolUpdate) SetPath(s string) *CentralStoragePoolUpdate {
	cspu.mutation.SetPath(s)
	return cspu
}

// SetHash sets the "hash" field.
func (cspu *CentralStoragePoolUpdate) SetHash(s string) *CentralStoragePoolUpdate {
	cspu.mutation.SetHash(s)
	return cspu
}

// Mutation returns the CentralStoragePoolMutation object of the builder.
func (cspu *CentralStoragePoolUpdate) Mutation() *CentralStoragePoolMutation {
	return cspu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cspu *CentralStoragePoolUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CentralStoragePoolMutation](ctx, cspu.sqlSave, cspu.mutation, cspu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cspu *CentralStoragePoolUpdate) SaveX(ctx context.Context) int {
	affected, err := cspu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cspu *CentralStoragePoolUpdate) Exec(ctx context.Context) error {
	_, err := cspu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cspu *CentralStoragePoolUpdate) ExecX(ctx context.Context) {
	if err := cspu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cspu *CentralStoragePoolUpdate) check() error {
	if v, ok := cspu.mutation.Filename(); ok {
		if err := centralstoragepool.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.filename": %w`, err)}
		}
	}
	if v, ok := cspu.mutation.Ext(); ok {
		if err := centralstoragepool.ExtValidator(v); err != nil {
			return &ValidationError{Name: "ext", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.ext": %w`, err)}
		}
	}
	if v, ok := cspu.mutation.Path(); ok {
		if err := centralstoragepool.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.path": %w`, err)}
		}
	}
	if v, ok := cspu.mutation.Hash(); ok {
		if err := centralstoragepool.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.hash": %w`, err)}
		}
	}
	return nil
}

func (cspu *CentralStoragePoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cspu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(centralstoragepool.Table, centralstoragepool.Columns, sqlgraph.NewFieldSpec(centralstoragepool.FieldID, field.TypeInt))
	if ps := cspu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cspu.mutation.DeleteAt(); ok {
		_spec.SetField(centralstoragepool.FieldDeleteAt, field.TypeTime, value)
	}
	if cspu.mutation.DeleteAtCleared() {
		_spec.ClearField(centralstoragepool.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := cspu.mutation.Filename(); ok {
		_spec.SetField(centralstoragepool.FieldFilename, field.TypeString, value)
	}
	if value, ok := cspu.mutation.Ext(); ok {
		_spec.SetField(centralstoragepool.FieldExt, field.TypeString, value)
	}
	if value, ok := cspu.mutation.Size(); ok {
		_spec.SetField(centralstoragepool.FieldSize, field.TypeFloat64, value)
	}
	if value, ok := cspu.mutation.AddedSize(); ok {
		_spec.AddField(centralstoragepool.FieldSize, field.TypeFloat64, value)
	}
	if value, ok := cspu.mutation.Path(); ok {
		_spec.SetField(centralstoragepool.FieldPath, field.TypeString, value)
	}
	if value, ok := cspu.mutation.Hash(); ok {
		_spec.SetField(centralstoragepool.FieldHash, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cspu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{centralstoragepool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cspu.mutation.done = true
	return n, nil
}

// CentralStoragePoolUpdateOne is the builder for updating a single CentralStoragePool entity.
type CentralStoragePoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CentralStoragePoolMutation
}

// SetDeleteAt sets the "delete_at" field.
func (cspuo *CentralStoragePoolUpdateOne) SetDeleteAt(t time.Time) *CentralStoragePoolUpdateOne {
	cspuo.mutation.SetDeleteAt(t)
	return cspuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cspuo *CentralStoragePoolUpdateOne) SetNillableDeleteAt(t *time.Time) *CentralStoragePoolUpdateOne {
	if t != nil {
		cspuo.SetDeleteAt(*t)
	}
	return cspuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (cspuo *CentralStoragePoolUpdateOne) ClearDeleteAt() *CentralStoragePoolUpdateOne {
	cspuo.mutation.ClearDeleteAt()
	return cspuo
}

// SetFilename sets the "filename" field.
func (cspuo *CentralStoragePoolUpdateOne) SetFilename(s string) *CentralStoragePoolUpdateOne {
	cspuo.mutation.SetFilename(s)
	return cspuo
}

// SetExt sets the "ext" field.
func (cspuo *CentralStoragePoolUpdateOne) SetExt(s string) *CentralStoragePoolUpdateOne {
	cspuo.mutation.SetExt(s)
	return cspuo
}

// SetSize sets the "size" field.
func (cspuo *CentralStoragePoolUpdateOne) SetSize(f float64) *CentralStoragePoolUpdateOne {
	cspuo.mutation.ResetSize()
	cspuo.mutation.SetSize(f)
	return cspuo
}

// AddSize adds f to the "size" field.
func (cspuo *CentralStoragePoolUpdateOne) AddSize(f float64) *CentralStoragePoolUpdateOne {
	cspuo.mutation.AddSize(f)
	return cspuo
}

// SetPath sets the "path" field.
func (cspuo *CentralStoragePoolUpdateOne) SetPath(s string) *CentralStoragePoolUpdateOne {
	cspuo.mutation.SetPath(s)
	return cspuo
}

// SetHash sets the "hash" field.
func (cspuo *CentralStoragePoolUpdateOne) SetHash(s string) *CentralStoragePoolUpdateOne {
	cspuo.mutation.SetHash(s)
	return cspuo
}

// Mutation returns the CentralStoragePoolMutation object of the builder.
func (cspuo *CentralStoragePoolUpdateOne) Mutation() *CentralStoragePoolMutation {
	return cspuo.mutation
}

// Where appends a list predicates to the CentralStoragePoolUpdate builder.
func (cspuo *CentralStoragePoolUpdateOne) Where(ps ...predicate.CentralStoragePool) *CentralStoragePoolUpdateOne {
	cspuo.mutation.Where(ps...)
	return cspuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cspuo *CentralStoragePoolUpdateOne) Select(field string, fields ...string) *CentralStoragePoolUpdateOne {
	cspuo.fields = append([]string{field}, fields...)
	return cspuo
}

// Save executes the query and returns the updated CentralStoragePool entity.
func (cspuo *CentralStoragePoolUpdateOne) Save(ctx context.Context) (*CentralStoragePool, error) {
	return withHooks[*CentralStoragePool, CentralStoragePoolMutation](ctx, cspuo.sqlSave, cspuo.mutation, cspuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cspuo *CentralStoragePoolUpdateOne) SaveX(ctx context.Context) *CentralStoragePool {
	node, err := cspuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cspuo *CentralStoragePoolUpdateOne) Exec(ctx context.Context) error {
	_, err := cspuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cspuo *CentralStoragePoolUpdateOne) ExecX(ctx context.Context) {
	if err := cspuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cspuo *CentralStoragePoolUpdateOne) check() error {
	if v, ok := cspuo.mutation.Filename(); ok {
		if err := centralstoragepool.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.filename": %w`, err)}
		}
	}
	if v, ok := cspuo.mutation.Ext(); ok {
		if err := centralstoragepool.ExtValidator(v); err != nil {
			return &ValidationError{Name: "ext", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.ext": %w`, err)}
		}
	}
	if v, ok := cspuo.mutation.Path(); ok {
		if err := centralstoragepool.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.path": %w`, err)}
		}
	}
	if v, ok := cspuo.mutation.Hash(); ok {
		if err := centralstoragepool.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "CentralStoragePool.hash": %w`, err)}
		}
	}
	return nil
}

func (cspuo *CentralStoragePoolUpdateOne) sqlSave(ctx context.Context) (_node *CentralStoragePool, err error) {
	if err := cspuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(centralstoragepool.Table, centralstoragepool.Columns, sqlgraph.NewFieldSpec(centralstoragepool.FieldID, field.TypeInt))
	id, ok := cspuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CentralStoragePool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cspuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, centralstoragepool.FieldID)
		for _, f := range fields {
			if !centralstoragepool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != centralstoragepool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cspuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cspuo.mutation.DeleteAt(); ok {
		_spec.SetField(centralstoragepool.FieldDeleteAt, field.TypeTime, value)
	}
	if cspuo.mutation.DeleteAtCleared() {
		_spec.ClearField(centralstoragepool.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := cspuo.mutation.Filename(); ok {
		_spec.SetField(centralstoragepool.FieldFilename, field.TypeString, value)
	}
	if value, ok := cspuo.mutation.Ext(); ok {
		_spec.SetField(centralstoragepool.FieldExt, field.TypeString, value)
	}
	if value, ok := cspuo.mutation.Size(); ok {
		_spec.SetField(centralstoragepool.FieldSize, field.TypeFloat64, value)
	}
	if value, ok := cspuo.mutation.AddedSize(); ok {
		_spec.AddField(centralstoragepool.FieldSize, field.TypeFloat64, value)
	}
	if value, ok := cspuo.mutation.Path(); ok {
		_spec.SetField(centralstoragepool.FieldPath, field.TypeString, value)
	}
	if value, ok := cspuo.mutation.Hash(); ok {
		_spec.SetField(centralstoragepool.FieldHash, field.TypeString, value)
	}
	_node = &CentralStoragePool{config: cspuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cspuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{centralstoragepool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cspuo.mutation.done = true
	return _node, nil
}
