package service

import (
	"context"

	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	"go.uber.org/dig"
)

// UserService contain logic for User Controller
// @mock
type UserService interface {
	FindOne(ctx context.Context, id int64) (*repository.User, error)
	Find(ctx context.Context, paginationParam repository.PaginationParam) ([]*repository.User, error)
	Insert(ctx context.Context, user repository.User) (lastInsertID int64, err error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user repository.User) error
	FindOneByEmail(ctx context.Context, email string) (*repository.User, error)
}

// UserServiceImpl is implementation of UserService
type UserServiceImpl struct {
	dig.In
	UserRepo repository.UserRepo
	AuditTrailService
	HistoryService
	dbtxn.Transactional
}

// NewUserService return new instance of UserService
// @constructor
func NewUserService(impl UserServiceImpl) UserService {
	return &impl
}

// FindOne user
func (r *UserServiceImpl) FindOne(ctx context.Context, id int64) (user *repository.User, err error) {
	return r.UserRepo.FindOne(ctx, id)
}

// Find user
func (r *UserServiceImpl) Find(ctx context.Context, paginationParam repository.PaginationParam) (list []*repository.User, err error) {
	return r.UserRepo.Find(ctx, paginationParam)
}

// Insert user
func (r *UserServiceImpl) Insert(ctx context.Context, user repository.User) (newUserID int64, err error) {
	defer r.CommitMe(&ctx)()
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
func (r *UserServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	defer r.CommitMe(&ctx)()
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
func (r *UserServiceImpl) Update(ctx context.Context, user repository.User) (err error) {
	defer r.CommitMe(&ctx)()
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

// FindOneByEmail user
func (r *UserServiceImpl) FindOneByEmail(ctx context.Context, email string) (user *repository.User, err error) {
	return r.UserRepo.FindUserByEmail(ctx, email)
}
