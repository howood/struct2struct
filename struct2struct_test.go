package struct2struct

import (
	"reflect"
	"testing"
)

type testFromStruct struct {
	UserId     int32                `json:"userId,omitempty" teststruct:"UserId"`
	UserName   string               `json:"userName,omitempty" teststruct:"Username"`
	Email      string               `json:"email,omitempty" teststruct:"Email"`
	GroupID    int64                `json:"groupID,omitempty" teststruct:"GrpId"`
	OrgID      int64                `json:"orgID,omitempty" teststruct:"OrgId"`
	Flag       *bool                `json:"flag,omitempty" teststruct:"ConvertFlag"`
	FlagT      bool                 `json:"flagT,omitempty" teststruct:"ConvertFlagT"`
	MapData    map[string][]string  `json:"mapData,omitempty" teststruct:"MapData"`
	MapDataP   *map[string][]string `json:"mapDataP,omitempty" teststruct:"MapDataP"`
	MapDataPt  *map[string][]string `json:"mapDataPt,omitempty" teststruct:"MapDataPt"`
	ArrString  []string             `json:"arrString,omitempty" teststruct:"ArrString"`
	ArrStringP []*string            `json:"arrStringP,omitempty" teststruct:"ArrStringP"`
	SubS       *SubFromStruct       `json:"subS,omitempty" teststruct:"SubS"`
}

type testToStruct struct {
	UserId       *userId              `json:"userId,omitempty"`
	Username     string               `json:"username,omitempty"`
	Email        string               `json:"email,omitempty"`
	GrpId        *grpId               `json:"grpId,omitempty"`
	OrgId        *orgId               `json:"orgId,omitempty"`
	ConvertFlag  *bool                `json:"convertFlag,omitempty"`
	ConvertFlagT *bool                `json:"convertFlagT,omitempty"`
	MapData      map[string][]string  `json:"mapData,omitempty"`
	MapDataP     *map[string][]string `json:"mapDataP,omitempty"`
	MapDataPt    map[string][]string  `json:"mapDataPt,omitempty"`
	ArrString    []string             `json:"arrString,omitempty"`
	ArrStringP   []*string            `json:"arrStringP,omitempty"`
	SubS         *SubToStruct         `json:"subS,omitempty"`
}

type SubFromStruct struct {
	UserId     int32                `json:"userId,omitempty" teststruct:"UserId"`
	UserName   string               `json:"userName,omitempty" teststruct:"Username"`
	Email      string               `json:"email,omitempty" teststruct:"Email"`
	GroupID    int64                `json:"groupID,omitempty" teststruct:"GrpId"`
	OrgID      int64                `json:"orgID,omitempty" teststruct:"OrgId"`
	Flag       *bool                `json:"flag,omitempty" teststruct:"ConvertFlag"`
	FlagT      bool                 `json:"flagT,omitempty" teststruct:"ConvertFlagT"`
	MapData    map[string][]string  `json:"mapData,omitempty" teststruct:"MapData"`
	MapDataP   *map[string][]string `json:"mapDataP,omitempty" teststruct:"MapDataP"`
	MapDataPt  *map[string][]string `json:"mapDataPt,omitempty" teststruct:"MapDataPt"`
	ArrString  []string             `json:"arrString,omitempty" teststruct:"ArrString"`
	ArrStringP []*string            `json:"arrStringP,omitempty" teststruct:"ArrStringP"`
}

type SubToStruct struct {
	UserId     int32                `json:"userId,omitempty"`
	UserName   string               `json:"userName,omitempty"`
	Email      string               `json:"email,omitempty"`
	GroupID    int64                `json:"groupID,omitempty"`
	OrgID      int64                `json:"orgID,omitempty"`
	Flag       *bool                `json:"flag,omitempty"`
	FlagT      bool                 `json:"flagT,omitempty"`
	MapData    map[string][]string  `json:"mapData,omitempty"`
	MapDataP   *map[string][]string `json:"mapDataP,omitempty"`
	MapDataPt  *map[string][]string `json:"mapDataPt,omitempty"`
	ArrString  []string             `json:"arrString,omitempty"`
	ArrStringP []*string            `json:"arrStringP,omitempty"`
}

type userId int32
type orgId int64
type grpId int64

func Test_Convert(t *testing.T) {
	str1 := "111"
	str2 := "222"
	testfrom := testFromStruct{
		UserId:   2222222,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  333333,
		OrgID:    444444,
		Flag:     &BOOLVAL,
		FlagT:    true,
		MapData: map[string][]string{
			"aaa": []string{"111", "222"},
			"bbb": []string{"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": []string{"111", "222"},
			"ddd": []string{"333", "444"},
		},
		MapDataPt: &map[string][]string{
			"eee": []string{"111", "222"},
			"fff": []string{"333", "444"},
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubFromStruct{
			UserId:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &BOOLVAL,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": []string{"111", "222"},
				"bbb": []string{"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": []string{"111", "222"},
				"ddd": []string{"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": []string{"111", "222"},
				"fff": []string{"333", "444"},
			},
			ArrString: []string{
				"111", "222",
			},
			ArrStringP: []*string{
				&str1, &str2,
			},
		},
	}
	uid := userId(2222222)
	gid := grpId(333333)
	oid := orgId(444444)
	testto := testToStruct{
		UserId:       &uid,
		Username:     "aaaaaaa",
		Email:        "bbbbb",
		GrpId:        &gid,
		OrgId:        &oid,
		ConvertFlag:  &BOOLVAL,
		ConvertFlagT: &BOOLVAL,
		MapData: map[string][]string{
			"aaa": []string{"111", "222"},
			"bbb": []string{"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": []string{"111", "222"},
			"ddd": []string{"333", "444"},
		},
		MapDataPt: map[string][]string{
			"eee": []string{"111", "222"},
			"fff": []string{"333", "444"},
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubToStruct{
			UserId:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &BOOLVAL,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": []string{"111", "222"},
				"bbb": []string{"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": []string{"111", "222"},
				"ddd": []string{"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": []string{"111", "222"},
				"fff": []string{"333", "444"},
			},
			ArrString: []string{
				"111", "222",
			},
			ArrStringP: []*string{
				&str1, &str2,
			},
		},
	}
	resultstruct := testToStruct{}
	result := ConvertStructToStruct(&testfrom, &resultstruct, "teststruct", "").(testToStruct)
	t.Log(result)
	t.Log(result.SubS)
	if reflect.DeepEqual(result, testto) == false {
		t.Fatal("failed Convert")
	}
	resultstruct2 := testFromStruct{}
	result2 := ConvertStructToStruct(&testto, &resultstruct2, "", "teststruct").(testFromStruct)
	t.Log(result2)
	t.Log(result2.SubS)
	if reflect.DeepEqual(result2, testfrom) == false {
		t.Fatal("failed Convert")
	}
	t.Log("success Convert")
}

func Test_MergeStructToStruct(t *testing.T) {
	str1 := "111"
	str2 := "222"
	testfrom := testFromStruct{
		UserId:  333,
		GroupID: 7777,
		OrgID:   8888,
		Flag:    &BOOLVAL,
		FlagT:   true,
		MapData: map[string][]string{
			"aaa": []string{"111", "222"},
			"bbb": []string{"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": []string{"111", "222"},
			"ddd": []string{"333", "444"},
		},
		MapDataPt: &map[string][]string{
			"eee": []string{"111", "222"},
			"fff": []string{"333", "444"},
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubFromStruct{
			UserId:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &BOOLVAL,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": []string{"111", "222"},
				"bbb": []string{"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": []string{"111", "222"},
				"ddd": []string{"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": []string{"111", "222"},
				"fff": []string{"333", "444"},
			},
			ArrString: []string{
				"111", "222",
			},
			ArrStringP: []*string{
				&str1, &str2,
			},
		},
	}
	testTo := testFromStruct{
		UserId:   4444,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  333,
		OrgID:    444,
		ArrString: []string{
			"6666", "6655",
		},
		MapData: map[string][]string{
			"vvvv": []string{"111666", "222666"},
			"ssss": []string{"333666", "444666"},
		},
		MapDataP: &map[string][]string{
			"gggg": []string{"111666", "222666"},
			"wwww": []string{"333666", "444666"},
		},
		MapDataPt: &map[string][]string{
			"ooooo": []string{"1116666", "22266666"},
			"ppppp": []string{"3336666", "444666666"},
		},
	}
	testresult := testFromStruct{
		UserId:   333,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  7777,
		OrgID:    8888,
		Flag:     &BOOLVAL,
		FlagT:    true,
		MapData: map[string][]string{
			"aaa": []string{"111", "222"},
			"bbb": []string{"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": []string{"111", "222"},
			"ddd": []string{"333", "444"},
		},
		MapDataPt: &map[string][]string{
			"eee": []string{"111", "222"},
			"fff": []string{"333", "444"},
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubFromStruct{
			UserId:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &BOOLVAL,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": []string{"111", "222"},
				"bbb": []string{"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": []string{"111", "222"},
				"ddd": []string{"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": []string{"111", "222"},
				"fff": []string{"333", "444"},
			},
			ArrString: []string{
				"111", "222",
			},
			ArrStringP: []*string{
				&str1, &str2,
			},
		},
	}
	result := MergeStructToStruct(&testfrom, &testTo, "", "")
	t.Log(result)
	if reflect.DeepEqual(result, testresult) == false {
		t.Fatal("failed Convert")
	}
	t.Log("success Convert")
}
