package dbtxn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/dig"
)

// Transactional database
type Transactional struct {
	dig.In
	*sql.DB
}

// CommitMe to create begin transaction and return commit function to be deffered
func (t *Transactional) CommitMe(ctx *context.Context) func() error {
	var (
		tx  *sql.Tx
		err error
	)
	*ctx = WithContext(*ctx)
	if tx, err = t.DB.BeginTx(*ctx, nil); err != nil {
		return func() error {
			if r := recover(); r != nil {
				SetError(*ctx, fmt.Errorf("%v", r))
			}
			return err
		}
	}
	set(*ctx, tx)
	return func() error {
		if r := recover(); r != nil {
			SetError(*ctx, fmt.Errorf("%v", r))
			return tx.Rollback()
		}
		if err := Error(*ctx); err != nil {
			return tx.Rollback()
		}
		return tx.Commit()
	}
}

// CancelMe is store error to context to trigger the rollback mechanism
func (t *Transactional) CancelMe(ctx context.Context, err error) error {
	return SetError(ctx, err)
}

func set(ctx context.Context, tx sq.BaseRunner) error {
	if txo := get(ctx); txo != nil {
		txo.tx = tx
		return nil
	}
	return errors.New("Context have no transaction")
}
