package TestHomeWork

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
