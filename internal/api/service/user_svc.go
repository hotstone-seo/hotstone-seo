package service

import (
	"context"
	"database/sql"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

type (
	// UserSvc contain logic for User Controller
	// @mock
	UserSvc interface {
		FindOne(ctx context.Context, id int64) (*repository.User, error)
		Find(ctx context.Context) ([]*repository.User, error)
		Insert(ctx context.Context, user repository.User) (lastInsertID int64, err error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, user repository.User) error
	}
	// UserSvcImpl is implementation of UserService
	UserSvcImpl struct {
		dig.In
		UserRepo   repository.UserRepo
		AuditTrail AuditTrailSvc
		dbtxn.Transactional
	}
)

// NewUserSvc return new instance of UserService
// @ctor
func NewUserSvc(impl UserSvcImpl) UserSvc {
	return &impl
}

// FindOne user
func (r *UserSvcImpl) FindOne(ctx context.Context, id int64) (*repository.User, error) {
	users, err := r.UserRepo.Find(ctx, dbkit.Equal(repository.UserTable.ID, id))
	if err != nil {
		return nil, err
	}
	if len(users) < 1 {
		return nil, sql.ErrNoRows
	}
	return users[0], nil
}

// Find user
func (r *UserSvcImpl) Find(ctx context.Context) (list []*repository.User, err error) {
	return r.UserRepo.Find(ctx)
}

// Insert user
func (r *UserSvcImpl) Insert(ctx context.Context, user repository.User) (int64, error) {
	id, err := r.UserRepo.Insert(ctx, user)
	if err != nil {
		return -1, err
	}
	user.ID = id

	r.AuditTrail.RecordInsert(ctx, "users", id, user)
	return id, nil
}

// Delete user
func (r *UserSvcImpl) Delete(ctx context.Context, id int64) error {
	defer r.BeginTxn(&ctx)()
	users, _ := r.UserRepo.Find(ctx, dbkit.Equal(repository.UserTable.ID, id))
	if len(users) < 1 {
		return nil
	}

	if err := r.UserRepo.Delete(ctx, id); err != nil {
		r.CancelMe(ctx, err)
		return err
	}

	r.AuditTrail.RecordDelete(ctx, "users", id, users[0])
	return nil
}

// Update user
func (r *UserSvcImpl) Update(ctx context.Context, user repository.User) error {
	defer r.BeginTxn(&ctx)()
	oldUser, err := r.UserRepo.Find(ctx, dbkit.Equal(repository.UserTable.ID, user.ID))
	if err != nil {
		return err
	}
	if err := r.UserRepo.Update(ctx, user); err != nil {
		return err
	}

	r.AuditTrail.RecordUpdate(ctx, "users", user.ID, oldUser, user)
	return nil
}
