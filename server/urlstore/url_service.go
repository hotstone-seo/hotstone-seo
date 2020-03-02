package urlstore

import (
	"context"
	"strconv"

	log "github.com/sirupsen/logrus"

	"go.uber.org/dig"
)

// URLService contain logic of url [mock]
type URLService interface {
	FullSync() error
	Sync() error
	Match(url string) (int, map[string]string)
	DumpTree() string
	Get(path string, pvalues []string) (data interface{}, pnames []string)
	Delete(id int64) bool
	Insert(id int64, key string)
	Update(id int64, key string)
	Count() int
}

// NewURLService return new instance of URLService [constructor]
func NewURLService(svc URLSyncRepo) URLService {
	return &URLServiceImpl{
		URLSyncRepo:   svc,
		Store:         NewStore(),
		LatestVersion: 0,
	}
}

// URLServiceImpl is implementation of URLService
type URLServiceImpl struct {
	dig.In
	URLSyncRepo
	Store
	LatestVersion int
}

// FullSync to sync from url-sync data to in-memory url-store from beginning
func (s *URLServiceImpl) FullSync() error {

	list, err := s.Find(context.Background())
	if err != nil {
		return err
	}

	if len(list) == 0 {
		return nil
	}

	s.Store = NewStore()
	s.setStore(list)

	oldestURLSync := list[len(list)-1]
	s.LatestVersion = int(oldestURLSync.Version)

	return nil
}

// Sync to  from url-sync data to in-memory url-store based on diff
func (s *URLServiceImpl) Sync() error {
	ctx := context.Background()

	LatestVersionSync, err := s.GetLatestVersion(ctx)
	if err != nil {
		return err
	}

	if s.LatestVersion == int(LatestVersionSync) {
		return nil
	}

	if s.LatestVersion != 0 && LatestVersionSync == 0 {
		s.Store = NewStore()
		s.LatestVersion = int(LatestVersionSync)
		return nil
	}

	if s.LatestVersion > int(LatestVersionSync) {
		return s.FullSync()
	}

	if s.LatestVersion < int(LatestVersionSync) {
		listDiffURLSync, err := s.GetListDiff(ctx, int64(s.LatestVersion))
		if err != nil {
			return err
		}
		s.setStore(listDiffURLSync)

		oldestURLSync := listDiffURLSync[len(listDiffURLSync)-1]
		s.LatestVersion = int(oldestURLSync.Version)
	}

	return nil
}

// Match return rule id and parameter map
func (s *URLServiceImpl) Match(url string) (int, map[string]string) {
	maxParams := 256
	pvalues := make([]string, maxParams)
	varValue := map[string]string{}

	data, pnames := s.Store.Get(url, pvalues)
	// fmt.Printf("[DATA:%s][PNAMES:%+v]", data, pnames)
	if data == nil {
		return -1, varValue
	}

	if len(pnames) > 0 {
		for i, name := range pnames {
			varValue[name] = pvalues[i]
		}
	}

	idStr, ok := data.(string)
	if !ok {
		log.Warnf("[GetURL] Failed to cast data to string. data=%+v", data)
		return -1, varValue
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warnf("[GetURL] Failed to convert string data to int. idStr=%+v", idStr)
		return -1, varValue
	}

	return id, varValue
}

// Insert to store
func (s *URLServiceImpl) Insert(id int64, key string) {
	data := strconv.FormatInt(id, 10)
	s.Store.Add(id, key, data)
}

// Update store
func (s *URLServiceImpl) Update(id int64, key string) {
	s.Delete(id)
	s.Insert(id, key)
}

// DumpTree to debug purpose
func (s *URLServiceImpl) DumpTree() string {
	return s.Store.String()
}

func (s *URLServiceImpl) setStore(listURLSync []*URLSync) {
	for _, sync := range listURLSync {
		switch sync.Operation {
		case "INSERT":
			s.Insert(sync.RuleID, *sync.LatestURLPattern)
		case "UPDATE":
			s.Update(sync.RuleID, *sync.LatestURLPattern)
		case "DELETE":
			s.Store.Delete(sync.RuleID)
		}
	}
}
