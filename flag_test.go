package mini

import (
	"fmt"
	"testing"
)


func TestNew(t *testing.T) {
	fmt.Println(New([]string{"-n", "wang", "-a", "19", "--b"}))
}

func TestRaw(t *testing.T) {
	fmt.Println(Raw([]string{"-n", "wang", "-a", "19", "--b", "nodejs", "python"}))
}
