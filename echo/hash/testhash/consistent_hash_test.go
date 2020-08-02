package testhash

import (
	"fmt"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	cm := New(3, nil)
	cm.AddNodes("A", "B", "C")
	fmt.Println(cm.Get("laileilae"))
	cm.DeleteNode("C")
	fmt.Println(cm.Get("laileilae"))
	cm.AddNode("C")
	fmt.Println(cm.Get("laileilae"))
}
