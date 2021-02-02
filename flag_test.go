package mini

import (
	"fmt"
	"testing"
)


func TestFlag(t *testing.T) {
	fmt.Println(New([]string{"-n", "wang", "-a", "19"}).GetFlags())
}
