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
	SubS       *SubStruct           `json:"subS,omitempty" teststruct:"SubS"`
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
	SubS         *SubStruct           `json:"subS,omitempty"`
}

type SubStruct struct {
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
		SubS: &SubStruct{
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
		SubS: &SubStruct{
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
	result := Convert(&testfrom, &resultstruct, "teststruct", "").(testToStruct)
	t.Log(result)
	t.Log(result.SubS)
	if reflect.DeepEqual(result, testto) == false {
		t.Fatal("failed Convert")
	}
	resultstruct2 := testFromStruct{}
	result2 := Convert(&testto, &resultstruct2, "", "teststruct").(testFromStruct)
	t.Log(result2)
	t.Log(result2.SubS)
	if reflect.DeepEqual(result2, testfrom) == false {
		t.Fatal("failed Convert")
	}
	t.Log("success Convert")
}

func Test_MergeStructToStruct(t *testing.T) {
	testfrom := testFromStruct{
		UserId:  333,
		GroupID: 7777,
		OrgID:   8888,
		Flag:    &BOOLVAL,
	}
	testTo := testFromStruct{
		UserId:   4444,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  333,
		OrgID:    444,
	}
	testresult := testFromStruct{
		UserId:   333,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  7777,
		OrgID:    8888,
		Flag:     &BOOLVAL,
	}
	result := Merge(&testfrom, &testTo, "", "")
	if reflect.DeepEqual(result, testresult) == false {
		t.Fatal("failed Convert")
	}
	t.Log("success Convert")
}
