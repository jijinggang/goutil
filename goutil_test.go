package goutil

import (
	"reflect"
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

type Addr struct {
	Ip   string
	Port int
}

type Config struct {
	Addr  Addr
	Title string
	Open  bool
	Speed float64
	Ext   []int
}

var expected = Config{
	Addr{"localhost", 80},
	"Test Title",
	false,
	1.5,
	[]int{1, 2, 3},
}

var testString = `{
	"Addr":{"Ip":"localhost", "Port":80},
	"Title":"Test Title",
	"Open":false,
	"Speed":1.5,
	"Ext":[1,2,3]
}`

func Test_ParseJsonString(t *testing.T) {
	var actual Config
	ParseJsonString(&actual, testString)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Actual %v expected %v", actual, expected)
	}
}

func Test_NewMap(t *testing.T) {
	m1 := NewMap("name", "golang", "level", 1)
	m2 := NewMap("m1", m1, "other", 0)
	AssertEqualStr(t, m1["name"].(string), "golang")
	AssertEqualInt(t, m1["level"].(int), 1)
	AssertEqualStr(t, m2["m1"].(map[string]interface{})["name"].(string), "golang")
	AssertEqualInt(t, m2["m1"].(map[string]interface{})["level"].(int), 1)
}
