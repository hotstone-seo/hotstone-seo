package service

// CenterService is center related logic
type CenterService interface {
	AddMetaTag(req AddMetaTagRequest) (int64, error)
	AddTitleTag(req AddTitleTagRequest) (int64, error)
	AddCanonicalTag(req AddCanonicalTagRequest) (int64, error)
	AddScriptTag(req AddScriptTagRequest) (int64, error)
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

// AddTitleTag to add titleTag
func (*CenterServiceImpl) AddTitleTag(req AddTitleTagRequest) (lastInsertedID int64, err error) {
	return
}

// AddCanonicalTag to add canonicalTag
func (*CenterServiceImpl) AddCanonicalTag(req AddCanonicalTagRequest) (lastInsertedID int64, err error) {
	return
}

// AddScriptTag to add scriptTag
func (*CenterServiceImpl) AddScriptTag(req AddScriptTagRequest) (lastInsertedID int64, err error) {
	return
}
