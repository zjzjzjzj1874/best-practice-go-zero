package helper

import (
	"fmt"
	"testing"
)

//func TestGetArea(t *testing.T) {
//	type args struct {
//		raw string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantRet []string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotRet := GetArea(tt.args.raw); !reflect.DeepEqual(gotRet, tt.wantRet) {
//				t.Errorf("GetArea() = %v, want %v", gotRet, tt.wantRet)
//			}
//		})
//	}
//}

func TestGetArea(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "青海省西宁市城中区南关街138号"},
		{name: "青海省西宁市城中区南关街138号"},
		{name: "广西壮族自治区南宁市城中区南关街138号"},
		{name: "黑龙江省齐齐哈尔市市中区南关街138号"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := GetArea(tt.name)
			fmt.Printf("%+v", ret)

		})
	}
}
