package foodb

import (
	"testing"

	"wagie.com/wageslavery/resources"
	_ "wagie.com/wageslavery/testing_init"
)

func TestMain(m *testing.M) {
	resources.InitializeSQLite()
	m.Run()
}

func TestGet(t *testing.T) {
	repo := FooRepository{}
	_, err := repo.Get()
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

}
