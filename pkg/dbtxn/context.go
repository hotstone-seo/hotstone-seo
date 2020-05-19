package dbtxn

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
)

const (
	// ContextKey for transactional context
	ContextKey key = iota
)

type key int

// Context of transaction
type Context struct {
	tx  sq.BaseRunner
	err error
}

// WithContext with transaction
func WithContext(parent context.Context) context.Context {
	return context.WithValue(parent, ContextKey, &Context{})
}

func get(ctx context.Context) *Context {
	if err, ok := ctx.Value(ContextKey).(*Context); ok {
		return err
	}
	return nil
}

// SetError to set error in ctx
func SetError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if txo := get(ctx); txo != nil {
		txo.err = err
		return nil
	}
	return errors.New("Context have no TXO")
}

// DB return transaction in context if available
func DB(ctx context.Context, defaultDB sq.BaseRunner) sq.BaseRunner {
	if c := get(ctx); c != nil {
		if c.tx != nil {
			return c.tx
		}
	}
	return defaultDB
}

// Error of transaction
func Error(ctx context.Context) error {
	if txo := get(ctx); txo != nil {
		if txo.err != nil {
			return txo.err
		}
	}
	return nil
}
