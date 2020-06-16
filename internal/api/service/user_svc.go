package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
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
		UserRepo repository.UserRepo
		AuditTrailService
		HistoryService
		dbtxn.Transactional
	}
)

// NewUserSvc return new instance of UserService
// @ctor
func NewUserSvc(impl UserSvcImpl) UserSvc {
	return &impl
}

// FindOne user
func (r *UserSvcImpl) FindOne(ctx context.Context, id int64) (user *repository.User, err error) {
	return r.UserRepo.FindOne(ctx, id)
}

// Find user
func (r *UserSvcImpl) Find(ctx context.Context) (list []*repository.User, err error) {
	return r.UserRepo.Find(ctx)
}

// Insert user
func (r *UserSvcImpl) Insert(ctx context.Context, user repository.User) (newUserID int64, err error) {
	defer r.BeginTxn(&ctx)()
	if newUserID, err = r.UserRepo.Insert(ctx, user); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	newUser, err := r.UserRepo.FindOne(ctx, newUserID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "users", newUserID, repository.Insert, nil, newUser); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return newUserID, nil
}

// Delete user
func (r *UserSvcImpl) Delete(ctx context.Context, id int64) (err error) {
	defer r.BeginTxn(&ctx)()
	oldUser, err := r.UserRepo.FindOne(ctx, id)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.HistoryService.RecordHistory(ctx, "users", id, oldUser); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.UserRepo.Delete(ctx, id); err != nil {
		r.CancelMe(ctx, err)
		return
	}

	if _, err = r.AuditTrailService.RecordChanges(ctx, "users", id, repository.Delete, oldUser, nil); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}

// Update user
func (r *UserSvcImpl) Update(ctx context.Context, user repository.User) (err error) {
	defer r.BeginTxn(&ctx)()
	oldUser, err := r.UserRepo.FindOne(ctx, user.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if err = r.UserRepo.Update(ctx, user); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	newUser, err := r.UserRepo.FindOne(ctx, user.ID)
	if err != nil {
		r.CancelMe(ctx, err)
		return
	}
	if _, err = r.AuditTrailService.RecordChanges(ctx, "users", user.ID, repository.Update, oldUser, newUser); err != nil {
		r.CancelMe(ctx, err)
		return
	}
	return nil
}
