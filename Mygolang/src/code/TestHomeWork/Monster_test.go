package TestHomeWork


//测试文件必须以XXXXXX_test.go命名
//测试的方法必须是Test_xxxxxx()来命名


import (
	"testing"
)

func TestMonster_Store(t *testing.T) {

	m := Monster{Name: "slike", Age: 16, Skill: "eat"}
	m.Store()

}

func TestName(t *testing.T) {

	m := Monster{}
	m1 := m.ReStore()
	t.Logf("%v",*m1)

}
