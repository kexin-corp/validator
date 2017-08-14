package validator

import (
	"strconv"
)

var provMap = map[int64]string{
	11: "北京", 12: "天津", 13: "河北", 14: "山西", 15: "内蒙古",
	21: "辽宁", 22: "吉林", 23: "黑龙江",
	31: "上海", 32: "江苏", 33: "浙江", 34: "安徽", 35: "福建", 36: "江西", 37: "山东",
	41: "河南", 42: "湖北", 43: "湖南", 44: "广东", 45: "广西", 46: "海南",
	50: "重庆", 51: "四川", 52: "贵州", 53: "云南", 54: "西藏",
	61: "陕西", 62: "甘肃", 63: "青海", 64: "宁夏", 65: "新疆",
	71: "台湾", 81: "香港", 82: "澳门", 91: "国外",
}

var checksum = [...]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
var weight = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

func Parse(card string) (ok bool, province string, year, month, day int64) {
	if len(card) != 15 && len(card) != 18 {
		return
	}

	province_i, err := strconv.ParseInt(card[0:3], 10, 64)
	if err != nil {
		return
	}
	province, _ = provMap[province_i]

	var month_s, day_s string
	if len(card) == 15 {
		year, err = strconv.ParseInt(card[6:8], 10, 64)
		if err != nil {
			return
		}
		year += 1900
		month_s = card[8:10]
		day_s = card[10:12]
	} else {
		year, err = strconv.ParseInt(card[6:10], 10, 64)
		if err != nil {
			return
		}
		month_s = card[10:12]
		day_s = card[12:14]
	}
	if year < 1700 || year > 3000 {
		return
	}

	month, err = strconv.ParseInt(month_s, 10, 64)
	if err != nil {
		return
	}
	if month > 12 {
		return
	}
	day, err = strconv.ParseInt(day_s, 10, 64)
	if err != nil {
		return
	}
	if day > 31 {
		return
	}

	sum := 0
	for i := 0; i < len(card)-1; i++ {
		sum += int(card[i]-'0') * weight[i]
	}

	lastNum := card[len(card)-1]
	if lastNum == 'x' {
		lastNum = 'X'
	}

	cs := sum % 11
	if checksum[cs] != lastNum {
		return
	}

	ok = true
	return
}

func idcardCheck(v interface{}, param string) error {
	if id, ok := v.(string); ok {
		ok, _, _, _, _ = Parse(id)
		if ok {
			return nil
		}

	}

	return ErrIdCard
}
