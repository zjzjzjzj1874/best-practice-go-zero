package file

import (
	"encoding/json"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("#json test", func(t *testing.T) {
		var testJson struct {
			Age   int    `json:"age"`
			Hello string `json:"hello"`
			Name  string `json:"name"`
		}
		if err := json.Unmarshal([]byte(TestJson), &testJson); err != nil {
			t.Fatal(err)
		}

		t.Log(testJson)
	})

	t.Run("#json slice", func(t *testing.T) {
		var testSlice []struct {
			Age   int    `json:"age"`
			Hello string `json:"hello"`
			Name  string `json:"name"`
		}
		if err := json.Unmarshal([]byte(TestSlice), &testSlice); err != nil {
			t.Fatal(err)
		}

		t.Log(testSlice)
	})
}
