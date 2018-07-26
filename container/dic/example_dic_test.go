package dic_test

import (
	"fmt"
	"testing"

	"github.com/recallsong/go-utils/container/dic"
)

func TestDic(t *testing.T) {
	dic := dic.Dic{
		"int":   123,
		"str":   "str_val",
		"dic":   dic.New().Set("key", "str_val"),
		"map":   map[string]interface{}{"key": 123},
		"slice": []int{1, 2, 3},
	}
	if val := dic.GetInt("int", 0); val != 123 {
		t.Error(`dic.GetInt(123,0) != 123`)
	}
	if val := dic.GetString("str", "error"); val != "str_val" {
		t.Error(`dic.GetString("str", "error"); val != "str_val"`)
	}
	{
		d, err := dic.GetDic("dic")
		if err != nil {
			t.Error(`dic.GetDic("dic")`, err.Error())
		}
		if d.JSON() != `{"key":"str_val"}` {
			t.Error(`d.String() != "{\"key\":\"str_val\"}`)
		}
	}
	{
		d, err := dic.GetDic("map")
		if err != nil {
			t.Error(`dic.GetDic("map")`, err.Error())
		}
		if d.JSON() != `{"key":123}` {
			t.Error(`d.String() != "{\"key\":123}`)
		}
	}
	{
		d, err := dic.GetDic("nil")
		if err != nil {
			t.Error(`dic.GetDic("nil")`, err.Error())
		}
		if d != nil {
			t.Error(`d != nil`)
		}
	}
}

func ExampleDic() {
	dic := dic.Dic{
		"int":   123,
		"str":   "hello",
		"dic":   dic.New().Set("key", "str_val"),
		"map":   map[string]interface{}{"key": 123},
		"slice": []int{1, 2, 3},
	}
	dic["other"] = "other_values"
	dic.Set("other1", "value1").Set("other2", "value2")
	fmt.Println(dic)
	d, _ := dic.GetDic("dic")
	fmt.Println(d)
	m, _ := dic.GetDic("map")
	fmt.Println(m)
	s := dic.Get("slice")
	fmt.Println(s)
	size := dic.Size()
	fmt.Println(size)
	keys := dic.Keys()
	fmt.Println(len(keys))
	// Output:
	// {"dic":{"key":"str_val"},"int":123,"map":{"key":123},"other":"other_values","other1":"value1","other2":"value2","slice":[1,2,3],"str":"hello"}
	// {"key":"str_val"}
	// {"key":123}
	// [1 2 3]
	// 8
	// 8
}
