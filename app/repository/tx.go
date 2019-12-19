package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func SetTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, "tx", tx)
}

func GetTx(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		log.Println("GetTx: TX NOT FOUND")
		return nil, errors.New("sql.Trx not exist")
	}

	log.Println("GetTx: TX FOUND")
	return tx, nil
}

func NewTxIfNotExist(ctx context.Context, db *sql.DB) (*sql.Tx, error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		log.Println("TX NOT FOUND")
		return db.Begin()
	}

	log.Println("TX FOUND")
	return tx, nil
}

// modified from : https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/

// A Txfn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(*sql.Tx) error

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func WithTransaction(db *sql.DB, fn TxFn) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("db.Begin err")
		return
	}

	defer func() {
		if p := recover(); p != nil {
			log.Println("will rollback & repanic")
			// a panic occurred, rollback and repanic
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("will rollback")
			// something went wrong, rollback
			_ = tx.Rollback()
		} else {
			log.Println("will commit")
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
