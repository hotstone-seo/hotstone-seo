package service

// CenterService is center related logic
type CenterService interface {
	AddMetaTag(req AddMetaTagRequest) (int64, error)
}

// CenterServiceImpl implementation of CenterService
type CenterServiceImpl struct {
}

// NewCenterService return new instance of CenterService
func NewCenterService() CenterService {
	return &CenterServiceImpl{}
}

// AddMetaTag to add metaTag
func (*CenterServiceImpl) AddMetaTag(req AddMetaTagRequest) (lastInsertedID int64, err error) {
	return
}
