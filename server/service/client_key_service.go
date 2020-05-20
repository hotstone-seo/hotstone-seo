package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/hotstone-seo/hotstone-seo/pkg/dbtxn"
	"github.com/hotstone-seo/hotstone-seo/server/repository"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// ClientKeyService contain logic for ClientKeyController
// @mock
type ClientKeyService interface {
	repository.ClientKeyRepo
	IsValidClientKey(ctx context.Context, clientKey string) bool
}

// ClientKeyServiceImpl is implementation of ClientKeyService
type ClientKeyServiceImpl struct {
	dig.In
	repository.ClientKeyRepo
	dbtxn.Transactional
	AuditTrailService AuditTrailService
	HistoryService    HistoryService
}

// NewClientKeyService return new instance of ClientKeyService
// @constructor
func NewClientKeyService(impl ClientKeyServiceImpl) ClientKeyService {
	return &impl
}

// Insert client key
func (s *ClientKeyServiceImpl) Insert(ctx context.Context, data repository.ClientKey) (newData repository.ClientKey, err error) {
	newPrefix, newKey, newKeyHashed := generateClientKey()
	data.Prefix = newPrefix
	data.Key = newKeyHashed
	newData, err = s.ClientKeyRepo.Insert(ctx, data)
	newData.Key = newKey
	if err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"client_keys",
			newData.ID,
			repository.Insert,
			nil,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return
}

// Update client key
func (s *ClientKeyServiceImpl) Update(ctx context.Context, data repository.ClientKey) (err error) {
	var oldData *repository.ClientKey
	if oldData, err = s.ClientKeyRepo.FindOne(ctx, dbkit.Equal("id", data.ID)); err != nil {
		return
	}
	if err = s.ClientKeyRepo.Update(ctx, data); err != nil {
		return
	}
	go func() {
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"client_keys",
			data.ID,
			repository.Update,
			oldData,
			data,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

// Delete client key
func (s *ClientKeyServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var oldData *repository.ClientKey
	if oldData, err = s.ClientKeyRepo.FindOne(ctx, dbkit.Equal("id", id)); err != nil {
		return
	}
	if err = s.ClientKeyRepo.Delete(ctx, id); err != nil {
		s.CancelMe(ctx, err)
		return
	}
	go func() {
		if _, histErr := s.HistoryService.RecordHistory(
			ctx,
			"client_keys",
			id,
			oldData,
		); histErr != nil {
			log.Error(histErr)
		}
		if _, auditErr := s.AuditTrailService.RecordChanges(
			ctx,
			"client_keys",
			id,
			repository.Delete,
			oldData,
			nil,
		); auditErr != nil {
			log.Error(auditErr)
		}
	}()
	return nil
}

// IsValidClientKey check validity of client key
func (s *ClientKeyServiceImpl) IsValidClientKey(ctx context.Context, clientKey string) bool {
	prefix, _, hashKey, err := extractClientKey(clientKey)
	if err != nil {
		return false
	}

	cKey, err := s.FindOne(ctx, dbkit.Equal("prefix", prefix), dbkit.Equal("key", hashKey))
	if err != nil || cKey == nil {
		return false
	}

	return true
}

func generateClientKey() (prefix, key, keyHashed string) {
	newPrefix := uniuri.NewLen(7)
	newKey := uniuri.NewLen(32)
	return newPrefix, newKey, fmt.Sprintf("%x", sha256.Sum256([]byte(newKey)))
}

func extractClientKey(clientKey string) (prefix, key, hashKey string, err error) {
	errNotValidKey := errors.New("Not valid key")
	k := strings.Split(clientKey, ".")
	if len(k) != 2 {
		return "", "", "", errNotValidKey
	}
	prefix = k[0]
	key = k[1]
	if len(prefix) != 7 && len(key) != 32 {
		return "", "", "", errNotValidKey
	}
	hashKey = fmt.Sprintf("%x", sha256.Sum256([]byte(key)))
	return
}
