package repository

import (
	"strconv"

	"github.com/labstack/gommon/log"
)

type VarMap map[string]string

type URLStore interface {
	GetURL(path string) (int, VarMap)
	AddURL(id int, key string)
	UpdateURL(id int, key string)
	DeleteURL(id int) bool
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

func (s *URLStoreImpl) GetURL(path string) (int, VarMap) {
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

func (s *URLStoreImpl) AddURL(id int, key string) {
	data := strconv.Itoa(id)

	s.store.Add(id, key, data)
}

func (s *URLStoreImpl) UpdateURL(id int, key string) {
	s.DeleteURL(id)
	s.AddURL(id, key)
}

func (s *URLStoreImpl) DeleteURL(id int) bool {
	return s.store.Delete(id)
}

func (s *URLStoreImpl) Count() int {
	return s.store.Count()
}
