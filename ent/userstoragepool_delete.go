// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"file_flow/ent/predicate"
	"file_flow/ent/userstoragepool"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserStoragePoolDelete is the builder for deleting a UserStoragePool entity.
type UserStoragePoolDelete struct {
	config
	hooks    []Hook
	mutation *UserStoragePoolMutation
}

// Where appends a list predicates to the UserStoragePoolDelete builder.
func (uspd *UserStoragePoolDelete) Where(ps ...predicate.UserStoragePool) *UserStoragePoolDelete {
	uspd.mutation.Where(ps...)
	return uspd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uspd *UserStoragePoolDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserStoragePoolMutation](ctx, uspd.sqlExec, uspd.mutation, uspd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uspd *UserStoragePoolDelete) ExecX(ctx context.Context) int {
	n, err := uspd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uspd *UserStoragePoolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userstoragepool.Table, sqlgraph.NewFieldSpec(userstoragepool.FieldID, field.TypeInt))
	if ps := uspd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uspd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uspd.mutation.done = true
	return affected, err
}

// UserStoragePoolDeleteOne is the builder for deleting a single UserStoragePool entity.
type UserStoragePoolDeleteOne struct {
	uspd *UserStoragePoolDelete
}

// Where appends a list predicates to the UserStoragePoolDelete builder.
func (uspdo *UserStoragePoolDeleteOne) Where(ps ...predicate.UserStoragePool) *UserStoragePoolDeleteOne {
	uspdo.uspd.mutation.Where(ps...)
	return uspdo
}

// Exec executes the deletion query.
func (uspdo *UserStoragePoolDeleteOne) Exec(ctx context.Context) error {
	n, err := uspdo.uspd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userstoragepool.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uspdo *UserStoragePoolDeleteOne) ExecX(ctx context.Context) {
	if err := uspdo.Exec(ctx); err != nil {
		panic(err)
	}
}
