package struct2struct

import (
	"reflect"
	"testing"
)

var boolVal = true

type testFromStruct struct {
	UserID     int32                `json:"userId,omitempty" teststruct:"UserID"`
	UserName   string               `json:"userName,omitempty" teststruct:"Username"`
	Email      string               `json:"email,omitempty" teststruct:"Email"`
	GroupID    int64                `json:"groupID,omitempty" teststruct:"GrpID"`
	OrgID      int64                `json:"orgID,omitempty" teststruct:"OrgID"`
	ParamA     int                  `json:"paramA,omitempty"`
	ParamB     int32                `json:"paramB,omitempty"`
	ParamC     int64                `json:"paramC,omitempty"`
	ParamD     float32              `json:"paramD,omitempty"`
	ParamE     float64              `json:"paramE,omitempty"`
	Flag       *bool                `json:"flag,omitempty" teststruct:"ConvertFlag"`
	FlagT      bool                 `json:"flagT,omitempty" teststruct:"ConvertFlagT"`
	MapData    map[string][]string  `json:"mapData,omitempty" teststruct:"MapData"`
	MapDataP   *map[string][]string `json:"mapDataP,omitempty" teststruct:"MapDataP"`
	MapDataPt  *map[string][]string `json:"mapDataPt,omitempty" teststruct:"MapDataPt"`
	MapData1   map[string]string    `json:"mapData1,omitempty" teststruct:"MapData1"`
	MapData2   map[string]int       `json:"mapData2,omitempty" teststruct:"MapData2"`
	MapData3   map[string]int32     `json:"mapData3,omitempty" teststruct:"MapData3"`
	MapData4   map[string]int64     `json:"mapData4,omitempty" teststruct:"MapData4"`
	MapData5   map[string]float32   `json:"mapData5,omitempty" teststruct:"MapData5"`
	MapData6   map[string]float64   `json:"mapData6,omitempty" teststruct:"MapData6"`
	ArrString  []string             `json:"arrString,omitempty" teststruct:"ArrString"`
	ArrStringP []*string            `json:"arrStringP,omitempty" teststruct:"ArrStringP"`
	SubS       *SubFromStruct       `json:"subS,omitempty" teststruct:"SubS"`
}

type testToStruct struct {
	UserID       *userId              `json:"userId,omitempty"`
	Username     string               `json:"username,omitempty"`
	Email        string               `json:"email,omitempty"`
	GrpID        *grpId               `json:"grpID,omitempty"`
	OrgID        *orgId               `json:"orgID,omitempty"`
	ParamA       int                  `json:"paramA,omitempty"`
	ParamB       int32                `json:"paramB,omitempty"`
	ParamC       int64                `json:"paramC,omitempty"`
	ParamD       float32              `json:"paramD,omitempty"`
	ParamE       float64              `json:"paramE,omitempty"`
	ConvertFlag  *bool                `json:"convertFlag,omitempty"`
	ConvertFlagT *bool                `json:"convertFlagT,omitempty"`
	MapData      map[string][]string  `json:"mapData,omitempty"`
	MapDataP     *map[string][]string `json:"mapDataP,omitempty"`
	MapDataPt    map[string][]string  `json:"mapDataPt,omitempty"`
	MapData1     map[string]string    `json:"mapData1,omitempty"`
	MapData2     map[string]int       `json:"mapData2,omitempty"`
	MapData3     map[string]int32     `json:"mapData3,omitempty"`
	MapData4     map[string]int64     `json:"mapData4,omitempty"`
	MapData5     map[string]float32   `json:"mapData5,omitempty"`
	MapData6     map[string]float64   `json:"mapData6,omitempty"`
	ArrString    []string             `json:"arrString,omitempty"`
	ArrStringP   []*string            `json:"arrStringP,omitempty"`
	SubS         *SubToStruct         `json:"subS,omitempty"`
}

type SubFromStruct struct {
	UserID     int32                `json:"userId,omitempty" teststruct:"UserID"`
	UserName   string               `json:"userName,omitempty" teststruct:"Username"`
	Email      string               `json:"email,omitempty" teststruct:"Email"`
	GroupID    int64                `json:"groupID,omitempty" teststruct:"GrpID"`
	OrgID      int64                `json:"orgID,omitempty" teststruct:"OrgID"`
	Flag       *bool                `json:"flag,omitempty" teststruct:"ConvertFlag"`
	FlagT      bool                 `json:"flagT,omitempty" teststruct:"ConvertFlagT"`
	MapData    map[string][]string  `json:"mapData,omitempty" teststruct:"MapData"`
	MapDataP   *map[string][]string `json:"mapDataP,omitempty" teststruct:"MapDataP"`
	MapDataPt  *map[string][]string `json:"mapDataPt,omitempty" teststruct:"MapDataPt"`
	ArrString  []string             `json:"arrString,omitempty" teststruct:"ArrString"`
	ArrStringP []*string            `json:"arrStringP,omitempty" teststruct:"ArrStringP"`
}

type SubToStruct struct {
	UserID     int32                `json:"userId,omitempty"`
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
		UserID:   2222222,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  333333,
		OrgID:    444444,
		ParamA:   1,
		ParamB:   1,
		ParamC:   1,
		ParamD:   1.1,
		ParamE:   1.1,
		Flag:     &boolVal,
		FlagT:    true,
		MapData: map[string][]string{
			"aaa": {"111", "222"},
			"bbb": {"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": {"111", "222"},
			"ddd": {"333", "444"},
		},
		MapDataPt: &map[string][]string{
			"eee": {"111", "222"},
			"fff": {"333", "444"},
		},
		MapData1: map[string]string{
			"aaa": "111",
			"bbb": "333",
		},
		MapData2: map[string]int{
			"aaa": 111,
			"bbb": 333,
		},
		MapData3: map[string]int32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData4: map[string]int64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData5: map[string]float32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData6: map[string]float64{
			"aaa": 111,
			"bbb": 333,
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubFromStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &boolVal,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": {"111", "222"},
				"bbb": {"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": {"111", "222"},
				"ddd": {"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": {"111", "222"},
				"fff": {"333", "444"},
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
		UserID:       &uid,
		Username:     "aaaaaaa",
		Email:        "bbbbb",
		GrpID:        &gid,
		OrgID:        &oid,
		ParamA:       1,
		ParamB:       1,
		ParamC:       1,
		ParamD:       1.1,
		ParamE:       1.1,
		ConvertFlag:  &boolVal,
		ConvertFlagT: &boolVal,
		MapData: map[string][]string{
			"aaa": {"111", "222"},
			"bbb": {"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": {"111", "222"},
			"ddd": {"333", "444"},
		},
		MapDataPt: map[string][]string{
			"eee": {"111", "222"},
			"fff": {"333", "444"},
		},
		MapData1: map[string]string{
			"aaa": "111",
			"bbb": "333",
		},
		MapData2: map[string]int{
			"aaa": 111,
			"bbb": 333,
		},
		MapData3: map[string]int32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData4: map[string]int64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData5: map[string]float32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData6: map[string]float64{
			"aaa": 111,
			"bbb": 333,
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubToStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &boolVal,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": {"111", "222"},
				"bbb": {"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": {"111", "222"},
				"ddd": {"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": {"111", "222"},
				"fff": {"333", "444"},
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
	result, err := ConvertStructToStruct(&testfrom, &resultstruct, "teststruct", "")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(result.(testToStruct))
	t.Log(result.(testToStruct).SubS)
	if reflect.DeepEqual(result.(testToStruct), testto) == false {
		t.Fatal("failed Convert")
	}
	resultstruct2 := testFromStruct{}
	result2, err := ConvertStructToStruct(&testto, &resultstruct2, "", "teststruct")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(result2.(testFromStruct))
	t.Log(result2.(testFromStruct).SubS)
	if reflect.DeepEqual(result2.(testFromStruct), testfrom) == false {
		t.Fatal("failed Convert")
	}
	resultstruct3 := testToStruct{}
	result3, err := ConvertStructToStruct(testfrom, resultstruct3, "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
	t.Log(result3)
	t.Log(resultstruct3)
	t.Log("success Convert")
}

func Test_MergeStructToStruct(t *testing.T) {
	str1 := "111"
	str2 := "222"
	testfrom := testFromStruct{
		UserID:  333,
		GroupID: 7777,
		OrgID:   8888,
		ParamA:  1,
		ParamB:  1,
		ParamC:  1,
		ParamD:  1.1,
		ParamE:  1.1,
		Flag:    &boolVal,
		FlagT:   true,
		MapData: map[string][]string{
			"aaa": {"111", "222"},
			"bbb": {"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": {"111", "222"},
			"ddd": {"333", "444"},
		},
		MapDataPt: &map[string][]string{
			"eee": {"111", "222"},
			"fff": {"333", "444"},
		},
		MapData1: map[string]string{
			"aaa": "111",
			"bbb": "333",
		},
		MapData2: map[string]int{
			"aaa": 111,
			"bbb": 333,
		},
		MapData3: map[string]int32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData4: map[string]int64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData5: map[string]float32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData6: map[string]float64{
			"aaa": 111,
			"bbb": 333,
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubFromStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &boolVal,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": {"111", "222"},
				"bbb": {"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": {"111", "222"},
				"ddd": {"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": {"111", "222"},
				"fff": {"333", "444"},
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
		UserID:   4444,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  333,
		OrgID:    444,
		ArrString: []string{
			"6666", "6655",
		},
		MapData: map[string][]string{
			"vvvv": {"111666", "222666"},
			"ssss": {"333666", "444666"},
		},
		MapDataP: &map[string][]string{
			"gggg": {"111666", "222666"},
			"wwww": {"333666", "444666"},
		},
		MapDataPt: &map[string][]string{
			"ooooo": {"1116666", "22266666"},
			"ppppp": {"3336666", "444666666"},
		},
	}
	testresult := testFromStruct{
		UserID:   333,
		UserName: "aaaaaaa",
		Email:    "bbbbb",
		GroupID:  7777,
		OrgID:    8888,
		Flag:     &boolVal,
		FlagT:    true,
		ParamA:   1,
		ParamB:   1,
		ParamC:   1,
		ParamD:   1.1,
		ParamE:   1.1,
		MapData: map[string][]string{
			"aaa": {"111", "222"},
			"bbb": {"333", "444"},
		},
		MapDataP: &map[string][]string{
			"ccc": {"111", "222"},
			"ddd": {"333", "444"},
		},
		MapDataPt: &map[string][]string{
			"eee": {"111", "222"},
			"fff": {"333", "444"},
		},
		MapData1: map[string]string{
			"aaa": "111",
			"bbb": "333",
		},
		MapData2: map[string]int{
			"aaa": 111,
			"bbb": 333,
		},
		MapData3: map[string]int32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData4: map[string]int64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData5: map[string]float32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData6: map[string]float64{
			"aaa": 111,
			"bbb": 333,
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubFromStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Email:    "bbbbb",
			GroupID:  333333,
			OrgID:    444444,
			Flag:     &boolVal,
			FlagT:    true,
			MapData: map[string][]string{
				"aaa": {"111", "222"},
				"bbb": {"333", "444"},
			},
			MapDataP: &map[string][]string{
				"ccc": {"111", "222"},
				"ddd": {"333", "444"},
			},
			MapDataPt: &map[string][]string{
				"eee": {"111", "222"},
				"fff": {"333", "444"},
			},
			ArrString: []string{
				"111", "222",
			},
			ArrStringP: []*string{
				&str1, &str2,
			},
		},
	}
	result, err := MergeStructToStruct(&testfrom, &testTo, "", "")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(result)
	if reflect.DeepEqual(result, testresult) == false {
		t.Fatal("failed Merge")
	}
	testTo2 := testToStruct{}
	result2, err := MergeStructToStruct(testfrom, testTo2, "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
	t.Log(result2)
	t.Log(testTo2)
	t.Log("success Merge")
}
