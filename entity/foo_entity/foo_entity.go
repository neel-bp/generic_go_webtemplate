package foo_entity

import "wagie.com/wageslavery/viewmodels"

type IFooRepository interface {
	Get() ([]viewmodels.FooGet, error)
	Create(req viewmodels.FooGet) error
}

type IFooService interface {
	Get() ([]viewmodels.FooGet, error)
	Create(req viewmodels.FooGet) error
}
