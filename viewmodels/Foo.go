package viewmodels

type FooGet struct {
	Id    int64  `json:"id"`
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}
