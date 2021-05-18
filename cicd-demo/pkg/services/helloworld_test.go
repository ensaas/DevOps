package services

import (
	"testing"
)

func TestHelloworld(t *testing.T){
	Version, _ := Output()
	if len(Version) == 0 {
		t.Error("Test failed")
		return
	}

}