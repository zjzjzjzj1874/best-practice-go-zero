package helper

import "regexp"

var myExp = regexp.MustCompile(`(?P<province>[^省]+自治区|.*?省|.*?行政区|.*?市)(?P<city>[^市]+自治州|.*?地区|.*?行政单位|.+盟|市辖区|.*?市|.*?县)(?P<county>[^县]+县|.+区|.+市|.+旗|.+海域|.+岛)?(?P<town>[^区]+区|.+镇)?(?P<village>.*)`)

// GetArea GetArea
func GetArea(raw string) (ret []string) {
	res := myExp.FindStringSubmatch(raw)
	if len(res) >= 4 {
		ret = res[1:4]
	}
	return
}
