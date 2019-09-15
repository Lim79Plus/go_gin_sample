package common

import "testing"

func TestGenToken(t *testing.T) {
	InitConf()
	id := uint(1234)
	result := GenToken(id)
	t.Log("TestGenToken result:", result)
	if result == "" {
		t.Fatal("failed test", "result", result)
	}
}
