// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"file_flow/ent/centralstoragepool"
	"file_flow/ent/schema"
	"file_flow/ent/share"
	"file_flow/ent/user"
	"file_flow/ent/userstoragepool"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	centralstoragepoolMixin := schema.CentralStoragePool{}.Mixin()
	centralstoragepoolMixinHooks0 := centralstoragepoolMixin[0].Hooks()
	centralstoragepool.Hooks[0] = centralstoragepoolMixinHooks0[0]
	centralstoragepoolMixinInters0 := centralstoragepoolMixin[0].Interceptors()
	centralstoragepool.Interceptors[0] = centralstoragepoolMixinInters0[0]
	centralstoragepoolFields := schema.CentralStoragePool{}.Fields()
	_ = centralstoragepoolFields
	// centralstoragepoolDescFilename is the schema descriptor for filename field.
	centralstoragepoolDescFilename := centralstoragepoolFields[0].Descriptor()
	// centralstoragepool.FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	centralstoragepool.FilenameValidator = centralstoragepoolDescFilename.Validators[0].(func(string) error)
	// centralstoragepoolDescExt is the schema descriptor for ext field.
	centralstoragepoolDescExt := centralstoragepoolFields[1].Descriptor()
	// centralstoragepool.ExtValidator is a validator for the "ext" field. It is called by the builders before save.
	centralstoragepool.ExtValidator = centralstoragepoolDescExt.Validators[0].(func(string) error)
	// centralstoragepoolDescPath is the schema descriptor for path field.
	centralstoragepoolDescPath := centralstoragepoolFields[3].Descriptor()
	// centralstoragepool.PathValidator is a validator for the "path" field. It is called by the builders before save.
	centralstoragepool.PathValidator = centralstoragepoolDescPath.Validators[0].(func(string) error)
	// centralstoragepoolDescHash is the schema descriptor for hash field.
	centralstoragepoolDescHash := centralstoragepoolFields[4].Descriptor()
	// centralstoragepool.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	centralstoragepool.HashValidator = centralstoragepoolDescHash.Validators[0].(func(string) error)
	// centralstoragepoolDescCreateAt is the schema descriptor for create_at field.
	centralstoragepoolDescCreateAt := centralstoragepoolFields[5].Descriptor()
	// centralstoragepool.DefaultCreateAt holds the default value on creation for the create_at field.
	centralstoragepool.DefaultCreateAt = centralstoragepoolDescCreateAt.Default.(time.Time)
	shareFields := schema.Share{}.Fields()
	_ = shareFields
	// shareDescCreateAt is the schema descriptor for create_at field.
	shareDescCreateAt := shareFields[3].Descriptor()
	// share.DefaultCreateAt holds the default value on creation for the create_at field.
	share.DefaultCreateAt = shareDescCreateAt.Default.(time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinInters0 := userMixin[0].Interceptors()
	user.Interceptors[0] = userMixinInters0[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[0].Descriptor()
	// user.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	user.NicknameValidator = userDescNickname.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[1].Descriptor()
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescCreateAt is the schema descriptor for create_at field.
	userDescCreateAt := userFields[4].Descriptor()
	// user.DefaultCreateAt holds the default value on creation for the create_at field.
	user.DefaultCreateAt = userDescCreateAt.Default.(time.Time)
	userstoragepoolMixin := schema.UserStoragePool{}.Mixin()
	userstoragepoolMixinHooks0 := userstoragepoolMixin[0].Hooks()
	userstoragepool.Hooks[0] = userstoragepoolMixinHooks0[0]
	userstoragepoolMixinInters0 := userstoragepoolMixin[0].Interceptors()
	userstoragepool.Interceptors[0] = userstoragepoolMixinInters0[0]
	userstoragepoolFields := schema.UserStoragePool{}.Fields()
	_ = userstoragepoolFields
	// userstoragepoolDescParentID is the schema descriptor for parent_id field.
	userstoragepoolDescParentID := userstoragepoolFields[2].Descriptor()
	// userstoragepool.DefaultParentID holds the default value on creation for the parent_id field.
	userstoragepool.DefaultParentID = userstoragepoolDescParentID.Default.(int)
	// userstoragepoolDescFilename is the schema descriptor for filename field.
	userstoragepoolDescFilename := userstoragepoolFields[3].Descriptor()
	// userstoragepool.FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	userstoragepool.FilenameValidator = userstoragepoolDescFilename.Validators[0].(func(string) error)
	// userstoragepoolDescExt is the schema descriptor for ext field.
	userstoragepoolDescExt := userstoragepoolFields[4].Descriptor()
	// userstoragepool.ExtValidator is a validator for the "ext" field. It is called by the builders before save.
	userstoragepool.ExtValidator = userstoragepoolDescExt.Validators[0].(func(string) error)
	// userstoragepoolDescCreateAt is the schema descriptor for create_at field.
	userstoragepoolDescCreateAt := userstoragepoolFields[5].Descriptor()
	// userstoragepool.DefaultCreateAt holds the default value on creation for the create_at field.
	userstoragepool.DefaultCreateAt = userstoragepoolDescCreateAt.Default.(time.Time)
}

const (
	Version = "v0.11.9"                                         // Version of ent codegen.
	Sum     = "h1:dbbCkAiPVTRBIJwoZctiSYjB7zxQIBOzVSU5H9VYIQI=" // Sum of ent codegen.
)
