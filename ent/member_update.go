// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"enttest/ent/member"
	"enttest/ent/predicate"
	"enttest/ent/schema/constants"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MemberUpdate) SetUpdatedAt(i int64) *MemberUpdate {
	mu.mutation.ResetUpdatedAt()
	mu.mutation.SetUpdatedAt(i)
	return mu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (mu *MemberUpdate) AddUpdatedAt(i int64) *MemberUpdate {
	mu.mutation.AddUpdatedAt(i)
	return mu
}

// SetUID sets the "uid" field.
func (mu *MemberUpdate) SetUID(u uint) *MemberUpdate {
	mu.mutation.ResetUID()
	mu.mutation.SetUID(u)
	return mu
}

// AddUID adds u to the "uid" field.
func (mu *MemberUpdate) AddUID(u uint) *MemberUpdate {
	mu.mutation.AddUID(u)
	return mu
}

// SetAvatar sets the "avatar" field.
func (mu *MemberUpdate) SetAvatar(s string) *MemberUpdate {
	mu.mutation.SetAvatar(s)
	return mu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableAvatar(s *string) *MemberUpdate {
	if s != nil {
		mu.SetAvatar(*s)
	}
	return mu
}

// ClearAvatar clears the value of the "avatar" field.
func (mu *MemberUpdate) ClearAvatar() *MemberUpdate {
	mu.mutation.ClearAvatar()
	return mu
}

// SetGender sets the "gender" field.
func (mu *MemberUpdate) SetGender(c constants.Gender) *MemberUpdate {
	mu.mutation.SetGender(c)
	return mu
}

// SetDescription sets the "description" field.
func (mu *MemberUpdate) SetDescription(s string) *MemberUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNickname sets the "nickname" field.
func (mu *MemberUpdate) SetNickname(s string) *MemberUpdate {
	mu.mutation.SetNickname(s)
	return mu
}

// SetAreaCode sets the "area_code" field.
func (mu *MemberUpdate) SetAreaCode(u uint8) *MemberUpdate {
	mu.mutation.ResetAreaCode()
	mu.mutation.SetAreaCode(u)
	return mu
}

// SetNillableAreaCode sets the "area_code" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableAreaCode(u *uint8) *MemberUpdate {
	if u != nil {
		mu.SetAreaCode(*u)
	}
	return mu
}

// AddAreaCode adds u to the "area_code" field.
func (mu *MemberUpdate) AddAreaCode(u uint8) *MemberUpdate {
	mu.mutation.AddAreaCode(u)
	return mu
}

// ClearAreaCode clears the value of the "area_code" field.
func (mu *MemberUpdate) ClearAreaCode() *MemberUpdate {
	mu.mutation.ClearAreaCode()
	return mu
}

// SetMobile sets the "mobile" field.
func (mu *MemberUpdate) SetMobile(u uint64) *MemberUpdate {
	mu.mutation.ResetMobile()
	mu.mutation.SetMobile(u)
	return mu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableMobile(u *uint64) *MemberUpdate {
	if u != nil {
		mu.SetMobile(*u)
	}
	return mu
}

// AddMobile adds u to the "mobile" field.
func (mu *MemberUpdate) AddMobile(u uint64) *MemberUpdate {
	mu.mutation.AddMobile(u)
	return mu
}

// ClearMobile clears the value of the "mobile" field.
func (mu *MemberUpdate) ClearMobile() *MemberUpdate {
	mu.mutation.ClearMobile()
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableEmail(s *string) *MemberUpdate {
	if s != nil {
		mu.SetEmail(*s)
	}
	return mu
}

// ClearEmail clears the value of the "email" field.
func (mu *MemberUpdate) ClearEmail() *MemberUpdate {
	mu.mutation.ClearEmail()
	return mu
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MemberUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Avatar(); ok {
		if err := member.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf("ent: validator failed for field \"avatar\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Gender(); ok {
		if err := member.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Description(); ok {
		if err := member.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf("ent: validator failed for field \"nickname\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: member.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: member.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: member.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: member.FieldUID,
		})
	}
	if value, ok := mu.mutation.AddedUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: member.FieldUID,
		})
	}
	if value, ok := mu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAvatar,
		})
	}
	if mu.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldAvatar,
		})
	}
	if value, ok := mu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: member.FieldGender,
		})
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldDescription,
		})
	}
	if value, ok := mu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNickname,
		})
	}
	if value, ok := mu.mutation.AreaCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: member.FieldAreaCode,
		})
	}
	if value, ok := mu.mutation.AddedAreaCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: member.FieldAreaCode,
		})
	}
	if mu.mutation.AreaCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: member.FieldAreaCode,
		})
	}
	if value, ok := mu.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: member.FieldMobile,
		})
	}
	if value, ok := mu.mutation.AddedMobile(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: member.FieldMobile,
		})
	}
	if mu.mutation.MobileCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: member.FieldMobile,
		})
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldEmail,
		})
	}
	if mu.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldEmail,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MemberUpdateOne) SetUpdatedAt(i int64) *MemberUpdateOne {
	muo.mutation.ResetUpdatedAt()
	muo.mutation.SetUpdatedAt(i)
	return muo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (muo *MemberUpdateOne) AddUpdatedAt(i int64) *MemberUpdateOne {
	muo.mutation.AddUpdatedAt(i)
	return muo
}

// SetUID sets the "uid" field.
func (muo *MemberUpdateOne) SetUID(u uint) *MemberUpdateOne {
	muo.mutation.ResetUID()
	muo.mutation.SetUID(u)
	return muo
}

// AddUID adds u to the "uid" field.
func (muo *MemberUpdateOne) AddUID(u uint) *MemberUpdateOne {
	muo.mutation.AddUID(u)
	return muo
}

// SetAvatar sets the "avatar" field.
func (muo *MemberUpdateOne) SetAvatar(s string) *MemberUpdateOne {
	muo.mutation.SetAvatar(s)
	return muo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableAvatar(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetAvatar(*s)
	}
	return muo
}

// ClearAvatar clears the value of the "avatar" field.
func (muo *MemberUpdateOne) ClearAvatar() *MemberUpdateOne {
	muo.mutation.ClearAvatar()
	return muo
}

// SetGender sets the "gender" field.
func (muo *MemberUpdateOne) SetGender(c constants.Gender) *MemberUpdateOne {
	muo.mutation.SetGender(c)
	return muo
}

// SetDescription sets the "description" field.
func (muo *MemberUpdateOne) SetDescription(s string) *MemberUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNickname sets the "nickname" field.
func (muo *MemberUpdateOne) SetNickname(s string) *MemberUpdateOne {
	muo.mutation.SetNickname(s)
	return muo
}

// SetAreaCode sets the "area_code" field.
func (muo *MemberUpdateOne) SetAreaCode(u uint8) *MemberUpdateOne {
	muo.mutation.ResetAreaCode()
	muo.mutation.SetAreaCode(u)
	return muo
}

// SetNillableAreaCode sets the "area_code" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableAreaCode(u *uint8) *MemberUpdateOne {
	if u != nil {
		muo.SetAreaCode(*u)
	}
	return muo
}

// AddAreaCode adds u to the "area_code" field.
func (muo *MemberUpdateOne) AddAreaCode(u uint8) *MemberUpdateOne {
	muo.mutation.AddAreaCode(u)
	return muo
}

// ClearAreaCode clears the value of the "area_code" field.
func (muo *MemberUpdateOne) ClearAreaCode() *MemberUpdateOne {
	muo.mutation.ClearAreaCode()
	return muo
}

// SetMobile sets the "mobile" field.
func (muo *MemberUpdateOne) SetMobile(u uint64) *MemberUpdateOne {
	muo.mutation.ResetMobile()
	muo.mutation.SetMobile(u)
	return muo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableMobile(u *uint64) *MemberUpdateOne {
	if u != nil {
		muo.SetMobile(*u)
	}
	return muo
}

// AddMobile adds u to the "mobile" field.
func (muo *MemberUpdateOne) AddMobile(u uint64) *MemberUpdateOne {
	muo.mutation.AddMobile(u)
	return muo
}

// ClearMobile clears the value of the "mobile" field.
func (muo *MemberUpdateOne) ClearMobile() *MemberUpdateOne {
	muo.mutation.ClearMobile()
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableEmail(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetEmail(*s)
	}
	return muo
}

// ClearEmail clears the value of the "email" field.
func (muo *MemberUpdateOne) ClearEmail() *MemberUpdateOne {
	muo.mutation.ClearEmail()
	return muo
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	var (
		err  error
		node *Member
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MemberUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Avatar(); ok {
		if err := member.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf("ent: validator failed for field \"avatar\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Gender(); ok {
		if err := member.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Description(); ok {
		if err := member.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf("ent: validator failed for field \"nickname\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: member.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Member.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: member.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: member.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: member.FieldUID,
		})
	}
	if value, ok := muo.mutation.AddedUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: member.FieldUID,
		})
	}
	if value, ok := muo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAvatar,
		})
	}
	if muo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldAvatar,
		})
	}
	if value, ok := muo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: member.FieldGender,
		})
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldDescription,
		})
	}
	if value, ok := muo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNickname,
		})
	}
	if value, ok := muo.mutation.AreaCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: member.FieldAreaCode,
		})
	}
	if value, ok := muo.mutation.AddedAreaCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: member.FieldAreaCode,
		})
	}
	if muo.mutation.AreaCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: member.FieldAreaCode,
		})
	}
	if value, ok := muo.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: member.FieldMobile,
		})
	}
	if value, ok := muo.mutation.AddedMobile(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: member.FieldMobile,
		})
	}
	if muo.mutation.MobileCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: member.FieldMobile,
		})
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldEmail,
		})
	}
	if muo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldEmail,
		})
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
