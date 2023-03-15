// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"file_flow/ent/centralstoragepool"
	"file_flow/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CentralStoragePoolDelete is the builder for deleting a CentralStoragePool entity.
type CentralStoragePoolDelete struct {
	config
	hooks    []Hook
	mutation *CentralStoragePoolMutation
}

// Where appends a list predicates to the CentralStoragePoolDelete builder.
func (cspd *CentralStoragePoolDelete) Where(ps ...predicate.CentralStoragePool) *CentralStoragePoolDelete {
	cspd.mutation.Where(ps...)
	return cspd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cspd *CentralStoragePoolDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CentralStoragePoolMutation](ctx, cspd.sqlExec, cspd.mutation, cspd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cspd *CentralStoragePoolDelete) ExecX(ctx context.Context) int {
	n, err := cspd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cspd *CentralStoragePoolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(centralstoragepool.Table, sqlgraph.NewFieldSpec(centralstoragepool.FieldID, field.TypeInt))
	if ps := cspd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cspd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cspd.mutation.done = true
	return affected, err
}

// CentralStoragePoolDeleteOne is the builder for deleting a single CentralStoragePool entity.
type CentralStoragePoolDeleteOne struct {
	cspd *CentralStoragePoolDelete
}

// Where appends a list predicates to the CentralStoragePoolDelete builder.
func (cspdo *CentralStoragePoolDeleteOne) Where(ps ...predicate.CentralStoragePool) *CentralStoragePoolDeleteOne {
	cspdo.cspd.mutation.Where(ps...)
	return cspdo
}

// Exec executes the deletion query.
func (cspdo *CentralStoragePoolDeleteOne) Exec(ctx context.Context) error {
	n, err := cspdo.cspd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{centralstoragepool.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cspdo *CentralStoragePoolDeleteOne) ExecX(ctx context.Context) {
	if err := cspdo.Exec(ctx); err != nil {
		panic(err)
	}
}
