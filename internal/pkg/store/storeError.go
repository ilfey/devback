package store

import (
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type StoreErrorType string

const (
	StoreNotFound      StoreErrorType = "StoreNotFound"
	StoreAlreadyExists StoreErrorType = "StoreAlreadyExists"
	StoreUnknown       StoreErrorType = "StoreUnknown"
)

type StoreError interface {
	error
	Type() StoreErrorType
}

type StoreErrorImpl struct {
	Err error
}

func NewError(err error) StoreError {
	return &StoreErrorImpl{
		Err: err,
	}
}

func NewErrorAndLog(err error, logger *logrus.Entry) StoreError {
	e := &StoreErrorImpl{
		Err: err,
	}
	e.LogError(logger)

	return e
}

func (e *StoreErrorImpl) LogError(logger *logrus.Entry) {
	log := logger.WithFields(logrus.Fields{
		"type": e.Type(),
	})

	pgErr, ok := e.Err.(*pgconn.PgError)
	if ok {
		log.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
		return
	}

	if e.Type() == StoreUnknown {
		log.Errorf("[ UNKNOWN ] %v", e.Err)
		return
	}

	log.Error(e.Err)
}

func (e *StoreErrorImpl) Type() StoreErrorType {
	if e.IsNotFoundError() {
		return StoreNotFound
	}

	if pgErr, ok := e.Err.(*pgconn.PgError); ok {
		if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return StoreAlreadyExists
		}
	}

	return StoreUnknown
}

func (e *StoreErrorImpl) IsNotFoundError() bool {
	return e.Err == pgx.ErrNoRows
}

func (e *StoreErrorImpl) Error() string {
	return e.Err.Error()
}
