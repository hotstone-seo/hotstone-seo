package service

import (
	"context"
	"strconv"

	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

type (
	// StructuredDataService manages instances of Structured Data
	// @mock
	StructuredDataService interface {
		FindByRule(ctx context.Context, ruleID int64) ([]*repository.StructuredData, error)
		repository.StructuredDataRepo
	}
	// StructuredDataServiceImpl is an impolementation of StructuredDataService
	StructuredDataServiceImpl struct {
		dig.In
		repository.StructuredDataRepo
		AuditTrail AuditTrailSvc
	}
)

// NewStructuredDataService returns nrw instance of StructuredDataService
// @ctor
func NewStructuredDataService(impl StructuredDataServiceImpl) StructuredDataService {
	return &impl
}

// FindByRule returns list of structured data based on rule ID
func (s *StructuredDataServiceImpl) FindByRule(ctx context.Context, ruleID int64) ([]*repository.StructuredData, error) {
	return s.Find(ctx, dbkit.Equal("rule_id", strconv.FormatInt(ruleID, 10)))
}

func (s *StructuredDataServiceImpl) Insert(ctx context.Context, strData repository.StructuredData) (newID int64, err error) {
	if strData.Data == nil {
		strData.Data = map[string]interface{}{}
	}
	if newID, err = s.StructuredDataRepo.Insert(ctx, strData); err != nil {
		return
	}
	s.AuditTrail.RecordInsert(ctx, "structured data", newID, strData)
	return newID, nil
}

func (s *StructuredDataServiceImpl) Update(ctx context.Context, strData repository.StructuredData) (err error) {
	var prevStrData *repository.StructuredData
	if prevStrData, err = s.StructuredDataRepo.FindOne(ctx, strData.ID); err != nil {
		return
	}
	if strData.Data == nil {
		strData.Data = make(map[string]interface{}, 0)
	}
	if err = s.StructuredDataRepo.Update(ctx, strData); err != nil {
		return
	}

	s.AuditTrail.RecordUpdate(ctx, "structured data", strData.ID, prevStrData, strData)
	return nil
}

func (s *StructuredDataServiceImpl) Delete(ctx context.Context, id int64) (err error) {
	var strData *repository.StructuredData
	if strData, err = s.StructuredDataRepo.FindOne(ctx, id); err != nil {
		return
	}
	if err = s.StructuredDataRepo.Delete(ctx, id); err != nil {
		return
	}
	s.AuditTrail.RecordDelete(ctx, "structured data", id, strData)
	return nil
}
