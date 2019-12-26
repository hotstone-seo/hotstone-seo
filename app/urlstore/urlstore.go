package urlstore

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

type VarMap map[string]string

// URLStore [nomock]
type URLStore interface {
	Get(path string) (int, VarMap)
	Add(id int, key string)
	Update(id int, key string)
	Delete(id int) bool
	Count() int
}

// URLStoreImpl is implementation of URLStore
type URLStoreImpl struct {
	store URLStoreTree
}

// InitURLStore return new instance of URLStore
func InitURLStore() URLStore {
	return &URLStoreImpl{store: newURLStoreTree()}
}

func (s *URLStoreImpl) Get(path string) (int, VarMap) {
	maxParams := 256

	pvalues := make([]string, maxParams)

	varValue := VarMap{}

	data, pnames := s.store.Get(path, pvalues)
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

func (s *URLStoreImpl) Add(id int, key string) {
	data := strconv.Itoa(id)

	s.store.Add(id, key, data)
}

func (s *URLStoreImpl) Update(id int, key string) {
	s.Delete(id)
	s.Add(id, key)
}

func (s *URLStoreImpl) Delete(id int) bool {
	return s.store.Delete(id)
}

func (s *URLStoreImpl) Count() int {
	return s.store.Count()
}
