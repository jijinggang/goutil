package file

import (
	"testing"
)

var datas = [][]string{
	[]string{"C:\\a Bc\\xxx.txt", "C:/a Bc"},
	[]string{"/opt/name/ss.sh", "/opt/name"},
}

func AssertEqualStr(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("Expect %s, but %s", a, b)
	}
}

func AssertEqualInt(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("Expect %d, but %d", a, b)
	}
}
func Test_GetPath(t *testing.T) {
	for _, data := range datas {
		a := GetPath(data[0])
		b := data[1]
		AssertEqualStr(t, a, b)
	}
}
