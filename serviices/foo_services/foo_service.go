package fooservices

import (
	"wagie.com/wageslavery/entity/foo_entity"
	foodb "wagie.com/wageslavery/repositories/foo_repositories"
	"wagie.com/wageslavery/viewmodels"
)

type FooService struct {
	FooRepo foo_entity.IFooRepository
}

func New() *FooService {
	return &FooService{
		FooRepo: &foodb.FooRepository{},
	}
}

func (fs *FooService) Get() ([]viewmodels.FooGet, error) {
	return fs.FooRepo.Get()
}

func (fs *FooService) Create(req viewmodels.FooGet) error {
	return fs.FooRepo.Create(req)
}
