package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/hotstone-seo/hotstone-server/app/repository"
	"github.com/hotstone-seo/hotstone-server/app/service"
	"github.com/hotstone-seo/hotstone-server/typical"
)

func init() {
	typical.Context.Constructors.Append(repository.NewRuleRepo)
	typical.Context.Constructors.Append(repository.NewTxIfNotExist)
	typical.Context.Constructors.Append(repository.NewURLStoreSyncRepo)
	typical.Context.Constructors.Append(service.NewRuleService)
	typical.Context.Constructors.Append(service.NewUrlStore)
}