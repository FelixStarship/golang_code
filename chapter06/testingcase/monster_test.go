package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	res := monster.Store()
	if res != nil {
		t.Fatalf("monster.Store() err=%v", res)
	}
	t.Log("测试通过")
}

func TestReStore(t *testing.T) {
	monster := &Monster{}
	res := monster.ReStore()
	if res != nil {
		t.Fatalf("monster.ReStore() err=%v", res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() err=%v", monster.Name)
	}
	t.Logf("monster.ReStore()测试通过 Monster=%v", *monster)
}
