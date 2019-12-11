package struct2struct

import (
	"github.com/howood/struct2struct/internal/testdata/testa"
	"github.com/howood/struct2struct/internal/testdata/testb"
	"reflect"
	"testing"
)

var boolVal = true

type testFromStruct struct {
	UserID      int32                    `json:"userId,omitempty" teststruct:"UserID"`
	UserName    string                   `json:"userName,omitempty" teststruct:"Username"`
	Email       string                   `json:"email,omitempty" teststruct:"Email"`
	GroupID     int64                    `json:"groupID,omitempty" teststruct:"GrpID"`
	OrgID       int64                    `json:"orgID,omitempty" teststruct:"OrgID"`
	ParamA      int                      `json:"paramA,omitempty" teststruct:"ParamA"`
	ParamB      int32                    `json:"paramB,omitempty" teststruct:"ParamB"`
	ParamC      int64                    `json:"paramC,omitempty" teststruct:"ParamC"`
	ParamD      float32                  `json:"paramD,omitempty" teststruct:"ParamD"`
	ParamE      float64                  `json:"paramE,omitempty" teststruct:"ParamE"`
	ParamF      uint                     `json:"paramE,omitempty" teststruct:"ParamF"`
	ParamG      uint32                   `json:"paramE,omitempty" teststruct:"ParamG"`
	ParamH      uint64                   `json:"paramE,omitempty" teststruct:"ParamH"`
	Flag        *bool                    `json:"flag,omitempty" teststruct:"ConvertFlag"`
	FlagT       bool                     `json:"flagT,omitempty" teststruct:"ConvertFlagT"`
	FlagErr     *bool                    `json:"flagErr,omitempty" teststruct:"ConvertFlagErr"`
	InterF      interface{}              `json:"interF,omitempty" teststruct:"InterF"`
	MapData     map[string][]string      `json:"mapData,omitempty" teststruct:"MapData"`
	MapDataP    *map[string][]string     `json:"mapDataP,omitempty" teststruct:"MapDataP"`
	MapDataPt   *map[string][]string     `json:"mapDataPt,omitempty" teststruct:"MapDataPt"`
	MapData1    map[string]string        `json:"mapData1,omitempty" teststruct:"MapData1"`
	MapData2    map[string]int           `json:"mapData2,omitempty" teststruct:"MapData2"`
	MapData3    map[string]int32         `json:"mapData3,omitempty" teststruct:"MapData3"`
	MapData4    map[string]int64         `json:"mapData4,omitempty" teststruct:"MapData4"`
	MapData5    map[string]float32       `json:"mapData5,omitempty" teststruct:"MapData5"`
	MapData6    map[string]float64       `json:"mapData6,omitempty" teststruct:"MapData6"`
	MapData7    map[string]uint          `json:"mapData6,omitempty" teststruct:"MapData7"`
	MapData8    map[string]uint32        `json:"mapData6,omitempty" teststruct:"MapData8"`
	MapData9    map[string]uint64        `json:"mapData6,omitempty" teststruct:"MapData9"`
	MapData10   map[string]bool          `json:"mapData6,omitempty" teststruct:"MapData10"`
	MapData1p   map[string]*string       `json:"mapData1,omitempty" teststruct:"MapData1p"`
	MapData2p   map[string]*int          `json:"mapData2,omitempty" teststruct:"MapData2p"`
	MapData3p   map[string]*int32        `json:"mapData3,omitempty" teststruct:"MapData3p"`
	MapData4p   map[string]*int64        `json:"mapData4,omitempty" teststruct:"MapData4p"`
	MapData5p   map[string]*float32      `json:"mapData5,omitempty" teststruct:"MapData5p"`
	MapData6p   map[string]*float64      `json:"mapData6,omitempty" teststruct:"MapData6p"`
	MapData7p   map[string]*uint         `json:"mapData6,omitempty" teststruct:"MapData7p"`
	MapData8p   map[string]*uint32       `json:"mapData6,omitempty" teststruct:"MapData8p"`
	MapData9p   map[string]*uint64       `json:"mapData6,omitempty" teststruct:"MapData9p"`
	MapData10p  map[string]*bool         `json:"mapData6,omitempty" teststruct:"MapData10p"`
	MapData1a   map[string][]string      `json:"mapData1,omitempty" teststruct:"MapData1a"`
	MapData2a   map[string][]int         `json:"mapData2,omitempty" teststruct:"MapData2a"`
	MapData3a   map[string][]int32       `json:"mapData3,omitempty" teststruct:"MapData3a"`
	MapData4a   map[string][]int64       `json:"mapData4,omitempty" teststruct:"MapData4a"`
	MapData5a   map[string][]float32     `json:"mapData5,omitempty" teststruct:"MapData5a"`
	MapData6a   map[string][]float64     `json:"mapData6,omitempty" teststruct:"MapData6a"`
	MapData7a   map[string][]uint        `json:"mapData6,omitempty" teststruct:"MapData7a"`
	MapData8a   map[string][]uint32      `json:"mapData6,omitempty" teststruct:"MapData8a"`
	MapData9a   map[string][]uint64      `json:"mapData6,omitempty" teststruct:"MapData9a"`
	MapData10a  map[string][]bool        `json:"mapData6,omitempty" teststruct:"MapData10a"`
	MapData1ap  map[string][]*string     `json:"mapData1,omitempty" teststruct:"MapData1ap"`
	MapData2ap  map[string][]*int        `json:"mapData2,omitempty" teststruct:"MapData2ap"`
	MapData3ap  map[string][]*int32      `json:"mapData3,omitempty" teststruct:"MapData3ap"`
	MapData4ap  map[string][]*int64      `json:"mapData4,omitempty" teststruct:"MapData4ap"`
	MapData5ap  map[string][]*float32    `json:"mapData5,omitempty" teststruct:"MapData5ap"`
	MapData6ap  map[string][]*float64    `json:"mapData6,omitempty" teststruct:"MapData6ap"`
	MapData7ap  map[string][]*uint       `json:"mapData6,omitempty" teststruct:"MapData7ap"`
	MapData8ap  map[string][]*uint32     `json:"mapData6,omitempty" teststruct:"MapData8ap"`
	MapData9ap  map[string][]*uint64     `json:"mapData6,omitempty" teststruct:"MapData9ap"`
	MapData10ap map[string][]*bool       `json:"mapData6,omitempty" teststruct:"MapData10ap"`
	MapDataI    map[string]interface{}   `json:"mapDataI,omitempty" teststruct:"MapDataI"`
	MapDataIa   map[string][]interface{} `json:"mapDataIa,omitempty" teststruct:"MapDataIa"`
	ArrString   []string                 `json:"arrString,omitempty" teststruct:"ArrString"`
	ArrStringP  []*string                `json:"arrStringP,omitempty" teststruct:"ArrStringP"`
	SubS        *SubFromStruct           `json:"subS,omitempty" teststruct:"SubS"`
	Different   testa.DifferentStruct    `json:"different,omitempty" teststruct:"Different"`
	DifferentP  *testa.DifferentStruct   `json:"differentP,omitempty" teststruct:"DifferentP"`
	DifferentPT *testa.DifferentStruct   `json:"differentPT,omitempty" teststruct:"DifferentPT"`
	DifferentTP testa.DifferentStruct    `json:"differentTP,omitempty" teststruct:"DifferentTP"`
}

type testToStruct struct {
	UserID         *userId                   `json:"userId,omitempty"`
	Username       *string                   `json:"username,omitempty"`
	Email          string                    `json:"email,omitempty"`
	GrpID          *grpId                    `json:"grpID,omitempty"`
	OrgID          *orgId                    `json:"orgID,omitempty"`
	ParamA         *int                      `json:"paramA,omitempty"`
	ParamB         *int32                    `json:"paramB,omitempty"`
	ParamC         *int64                    `json:"paramC,omitempty"`
	ParamD         *float32                  `json:"paramD,omitempty"`
	ParamE         *float64                  `json:"paramE,omitempty"`
	ParamF         *uint                     `json:"paramE,omitempty"`
	ParamG         *uint32                   `json:"paramE,omitempty"`
	ParamH         *uint64                   `json:"paramE,omitempty"`
	ConvertFlag    *bool                     `json:"convertFlag,omitempty"`
	ConvertFlagT   *bool                     `json:"convertFlagT,omitempty"`
	ConvertFlagErr *bool                     `json:"convertFlagErr,omitempty"`
	InterF         interface{}               `json:"interF,omitempty"`
	MapData        map[string][]string       `json:"mapData,omitempty"`
	MapDataP       *map[string][]string      `json:"mapDataP,omitempty"`
	MapDataPt      map[string][]string       `json:"mapDataPt,omitempty"`
	MapData1       *map[string]string        `json:"mapData1,omitempty"`
	MapData2       *map[string]int           `json:"mapData2,omitempty"`
	MapData3       *map[string]int32         `json:"mapData3,omitempty"`
	MapData4       *map[string]int64         `json:"mapData4,omitempty"`
	MapData5       *map[string]float32       `json:"mapData5,omitempty"`
	MapData6       *map[string]float64       `json:"mapData6,omitempty"`
	MapData7       *map[string]uint          `json:"mapData7,omitempty"`
	MapData8       *map[string]uint32        `json:"mapData8,omitempty"`
	MapData9       *map[string]uint64        `json:"mapData9,omitempty"`
	MapData10      *map[string]bool          `json:"mapData10,omitempty"`
	MapData1p      *map[string]*string       `json:"mapData1p,omitempty"`
	MapData2p      *map[string]*int          `json:"mapData2p,omitempty"`
	MapData3p      *map[string]*int32        `json:"mapData3p,omitempty"`
	MapData4p      *map[string]*int64        `json:"mapData4p,omitempty"`
	MapData5p      *map[string]*float32      `json:"mapData5p,omitempty"`
	MapData6p      *map[string]*float64      `json:"mapData6p,omitempty"`
	MapData7p      *map[string]*uint         `json:"mapData7p,omitempty"`
	MapData8p      *map[string]*uint32       `json:"mapData8p,omitempty"`
	MapData9p      *map[string]*uint64       `json:"mapData9p,omitempty"`
	MapData10p     *map[string]*bool         `json:"mapData10p,omitempty"`
	MapData1a      *map[string][]string      `json:"mapData1a,omitempty"`
	MapData2a      *map[string][]int         `json:"mapData2a,omitempty"`
	MapData3a      *map[string][]int32       `json:"mapData3a,omitempty"`
	MapData4a      *map[string][]int64       `json:"mapData4a,omitempty"`
	MapData5a      *map[string][]float32     `json:"mapData5a,omitempty"`
	MapData6a      *map[string][]float64     `json:"mapData6a,omitempty"`
	MapData7a      *map[string][]uint        `json:"mapData7a,omitempty"`
	MapData8a      *map[string][]uint32      `json:"mapData8a,omitempty"`
	MapData9a      *map[string][]uint64      `json:"mapData9a,omitempty"`
	MapData10a     *map[string][]bool        `json:"mapData10a,omitempty"`
	MapData1ap     *map[string][]*string     `json:"mapData1ap,omitempty"`
	MapData2ap     *map[string][]*int        `json:"mapData2ap,omitempty"`
	MapData3ap     *map[string][]*int32      `json:"mapData3ap,omitempty"`
	MapData4ap     *map[string][]*int64      `json:"mapData4ap,omitempty"`
	MapData5ap     *map[string][]*float32    `json:"mapData5ap,omitempty"`
	MapData6ap     *map[string][]*float64    `json:"mapData6ap,omitempty"`
	MapData7ap     *map[string][]*uint       `json:"mapData7ap,omitempty"`
	MapData8ap     *map[string][]*uint32     `json:"mapData8ap,omitempty"`
	MapData9ap     *map[string][]*uint64     `json:"mapData9ap,omitempty"`
	MapData10ap    *map[string][]*bool       `json:"mapData10ap,omitempty"`
	MapDataI       *map[string]interface{}   `json:"mapDataI,omitempty"`
	MapDataIa      *map[string][]interface{} `json:"mapDataIa,omitempty"`
	ArrString      []string                  `json:"arrString,omitempty"`
	ArrStringP     []*string                 `json:"arrStringP,omitempty"`
	SubS           *SubToStruct2             `json:"subS,omitempty"`
	Different      testb.DifferentStruct     `json:"different,omitempty"`
	DifferentP     *testb.DifferentStruct    `json:"differentP,omitempty"`
	DifferentPT    testb.DifferentStruct     `json:"differentP,omitempty"`
	DifferentTP    *testb.DifferentStruct    `json:"differentP,omitempty"`
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
type SubToStruct2 struct {
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
	int_1 := 1
	var int_32 int32 = 1
	var int_64 int64 = 1
	var uint_2 uint = 2
	var uint_32 uint32 = 2
	var uint_64 uint64 = 2
	var float_32 float32 = 1.1
	var float_64 float64 = 1.1
	bool_t := true
	bool_f := false
	testfrom := testFromStruct{
		UserID:   2222222,
		UserName: "aaaaaaa",
		Email:    "",
		GroupID:  333333,
		OrgID:    444444,
		ParamA:   1,
		ParamB:   1,
		ParamC:   1,
		ParamD:   1.1,
		ParamE:   1.1,
		ParamF:   2,
		ParamG:   2,
		ParamH:   2,
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
		MapData7: map[string]uint{
			"aaa": 111,
			"bbb": 333,
		},
		MapData8: map[string]uint32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData9: map[string]uint64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData10: map[string]bool{
			"aaa": true,
			"bbb": false,
		},
		MapData1p: map[string]*string{
			"aaa": &str1,
			"bbb": &str2,
		},
		MapData2p: map[string]*int{
			"aaa": &int_1,
			"bbb": &int_1,
		},
		MapData3p: map[string]*int32{
			"aaa": &int_32,
			"bbb": &int_32,
		},
		MapData4p: map[string]*int64{
			"aaa": &int_64,
			"bbb": &int_64,
		},
		MapData5p: map[string]*float32{
			"aaa": &float_32,
			"bbb": &float_32,
		},
		MapData6p: map[string]*float64{
			"aaa": &float_64,
			"bbb": &float_64,
		},
		MapData7p: map[string]*uint{
			"aaa": &uint_2,
			"bbb": &uint_2,
		},
		MapData8p: map[string]*uint32{
			"aaa": &uint_32,
			"bbb": &uint_32,
		},
		MapData9p: map[string]*uint64{
			"aaa": &uint_64,
			"bbb": &uint_64,
		},
		MapData10p: map[string]*bool{
			"aaa": &bool_t,
			"bbb": &bool_f,
		},
		MapData1a: map[string][]string{
			"aaa": []string{"111"},
			"bbb": []string{"333"},
		},
		MapData2a: map[string][]int{
			"aaa": []int{111},
			"bbb": []int{333},
		},
		MapData3a: map[string][]int32{
			"aaa": []int32{111},
			"bbb": []int32{333},
		},
		MapData4a: map[string][]int64{
			"aaa": []int64{111},
			"bbb": []int64{333},
		},
		MapData5a: map[string][]float32{
			"aaa": []float32{111},
			"bbb": []float32{333},
		},
		MapData6a: map[string][]float64{
			"aaa": []float64{111},
			"bbb": []float64{333},
		},
		MapData7a: map[string][]uint{
			"aaa": []uint{111},
			"bbb": []uint{333},
		},
		MapData8a: map[string][]uint32{
			"aaa": []uint32{111},
			"bbb": []uint32{333},
		},
		MapData9a: map[string][]uint64{
			"aaa": []uint64{111},
			"bbb": []uint64{333},
		},
		MapData10a: map[string][]bool{
			"aaa": []bool{true},
			"bbb": []bool{false},
		},
		MapData1ap: map[string][]*string{
			"aaa": []*string{&str1},
			"bbb": []*string{&str2},
		},
		MapData2ap: map[string][]*int{
			"aaa": []*int{&int_1},
			"bbb": []*int{&int_1},
		},
		MapData3ap: map[string][]*int32{
			"aaa": []*int32{&int_32},
			"bbb": []*int32{&int_32},
		},
		MapData4ap: map[string][]*int64{
			"aaa": []*int64{&int_64},
			"bbb": []*int64{&int_64},
		},
		MapData5ap: map[string][]*float32{
			"aaa": []*float32{&float_32},
			"bbb": []*float32{&float_32},
		},
		MapData6ap: map[string][]*float64{
			"aaa": []*float64{&float_64},
			"bbb": []*float64{&float_64},
		},
		MapData7ap: map[string][]*uint{
			"aaa": []*uint{&uint_2},
			"bbb": []*uint{&uint_2},
		},
		MapData8ap: map[string][]*uint32{
			"aaa": []*uint32{&uint_32},
			"bbb": []*uint32{&uint_32},
		},
		MapData9ap: map[string][]*uint64{
			"aaa": []*uint64{&uint_64},
			"bbb": []*uint64{&uint_64},
		},
		MapData10ap: map[string][]*bool{
			"aaa": []*bool{&bool_t},
			"bbb": []*bool{&bool_f},
		},
		MapDataI: map[string]interface{}{
			"aaa": "111",
			"bbb": "333",
		},
		MapDataIa: map[string][]interface{}{
			"aaa": []interface{}{"111"},
			"bbb": []interface{}{"333"},
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
		Different: testa.DifferentStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Sub: &testa.DifferentStructASub{
				UserID:   2222222,
				UserName: "aaaaaaa",
			},
		},
		DifferentP: &testa.DifferentStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Sub: &testa.DifferentStructASub{
				UserID:   2222222,
				UserName: "aaaaaaa",
			},
		},
		DifferentPT: &testa.DifferentStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Sub: &testa.DifferentStructASub{
				UserID:   2222222,
				UserName: "aaaaaaa",
			},
		},
		DifferentTP: testa.DifferentStruct{
			UserID:   2222222,
			UserName: "aaaaaaa",
			Sub: &testa.DifferentStructASub{
				UserID:   2222222,
				UserName: "aaaaaaa",
			},
		},
	}
	uid := userId(2222222)
	gid := grpId(333333)
	oid := orgId(444444)
	aaaaaaa := "aaaaaaa"
	testto := testToStruct{
		UserID:       &uid,
		Username:     &aaaaaaa,
		Email:        "",
		GrpID:        &gid,
		OrgID:        &oid,
		ParamA:       &int_1,
		ParamB:       &int_32,
		ParamC:       &int_64,
		ParamD:       &float_32,
		ParamE:       &float_64,
		ParamF:       &uint_2,
		ParamG:       &uint_32,
		ParamH:       &uint_64,
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
		MapData1: &map[string]string{
			"aaa": "111",
			"bbb": "333",
		},
		MapData2: &map[string]int{
			"aaa": 111,
			"bbb": 333,
		},
		MapData3: &map[string]int32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData4: &map[string]int64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData5: &map[string]float32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData6: &map[string]float64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData7: &map[string]uint{
			"aaa": 111,
			"bbb": 333,
		},
		MapData8: &map[string]uint32{
			"aaa": 111,
			"bbb": 333,
		},
		MapData9: &map[string]uint64{
			"aaa": 111,
			"bbb": 333,
		},
		MapData10: &map[string]bool{
			"aaa": true,
			"bbb": false,
		},
		MapData1p: &map[string]*string{
			"aaa": &str1,
			"bbb": &str2,
		},
		MapData2p: &map[string]*int{
			"aaa": &int_1,
			"bbb": &int_1,
		},
		MapData3p: &map[string]*int32{
			"aaa": &int_32,
			"bbb": &int_32,
		},
		MapData4p: &map[string]*int64{
			"aaa": &int_64,
			"bbb": &int_64,
		},
		MapData5p: &map[string]*float32{
			"aaa": &float_32,
			"bbb": &float_32,
		},
		MapData6p: &map[string]*float64{
			"aaa": &float_64,
			"bbb": &float_64,
		},
		MapData7p: &map[string]*uint{
			"aaa": &uint_2,
			"bbb": &uint_2,
		},
		MapData8p: &map[string]*uint32{
			"aaa": &uint_32,
			"bbb": &uint_32,
		},
		MapData9p: &map[string]*uint64{
			"aaa": &uint_64,
			"bbb": &uint_64,
		},
		MapData10p: &map[string]*bool{
			"aaa": &bool_t,
			"bbb": &bool_f,
		},
		MapData1a: &map[string][]string{
			"aaa": []string{"111"},
			"bbb": []string{"333"},
		},
		MapData2a: &map[string][]int{
			"aaa": []int{111},
			"bbb": []int{333},
		},
		MapData3a: &map[string][]int32{
			"aaa": []int32{111},
			"bbb": []int32{333},
		},
		MapData4a: &map[string][]int64{
			"aaa": []int64{111},
			"bbb": []int64{333},
		},
		MapData5a: &map[string][]float32{
			"aaa": []float32{111},
			"bbb": []float32{333},
		},
		MapData6a: &map[string][]float64{
			"aaa": []float64{111},
			"bbb": []float64{333},
		},
		MapData7a: &map[string][]uint{
			"aaa": []uint{111},
			"bbb": []uint{333},
		},
		MapData8a: &map[string][]uint32{
			"aaa": []uint32{111},
			"bbb": []uint32{333},
		},
		MapData9a: &map[string][]uint64{
			"aaa": []uint64{111},
			"bbb": []uint64{333},
		},
		MapData10a: &map[string][]bool{
			"aaa": []bool{true},
			"bbb": []bool{false},
		},
		MapData1ap: &map[string][]*string{
			"aaa": []*string{&str1},
			"bbb": []*string{&str2},
		},
		MapData2ap: &map[string][]*int{
			"aaa": []*int{&int_1},
			"bbb": []*int{&int_1},
		},
		MapData3ap: &map[string][]*int32{
			"aaa": []*int32{&int_32},
			"bbb": []*int32{&int_32},
		},
		MapData4ap: &map[string][]*int64{
			"aaa": []*int64{&int_64},
			"bbb": []*int64{&int_64},
		},
		MapData5ap: &map[string][]*float32{
			"aaa": []*float32{&float_32},
			"bbb": []*float32{&float_32},
		},
		MapData6ap: &map[string][]*float64{
			"aaa": []*float64{&float_64},
			"bbb": []*float64{&float_64},
		},
		MapData7ap: &map[string][]*uint{
			"aaa": []*uint{&uint_2},
			"bbb": []*uint{&uint_2},
		},
		MapData8ap: &map[string][]*uint32{
			"aaa": []*uint32{&uint_32},
			"bbb": []*uint32{&uint_32},
		},
		MapData9ap: &map[string][]*uint64{
			"aaa": []*uint64{&uint_64},
			"bbb": []*uint64{&uint_64},
		},
		MapData10ap: &map[string][]*bool{
			"aaa": []*bool{&bool_t},
			"bbb": []*bool{&bool_f},
		},
		MapDataI: &map[string]interface{}{
			"aaa": "111",
			"bbb": "333",
		},
		MapDataIa: &map[string][]interface{}{
			"aaa": []interface{}{"111"},
			"bbb": []interface{}{"333"},
		},
		ArrString: []string{
			"111", "222",
		},
		ArrStringP: []*string{
			&str1, &str2,
		},
		SubS: &SubToStruct2{
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
	if reflect.DeepEqual(result.(testToStruct), testto) == false {
		aaa := result.(testToStruct)
		t.Logf("%#v", aaa.Different)
		t.Logf("%#v", testto.Different)
		t.Logf("%#v", aaa.DifferentP)
		t.Logf("%#v", testto.DifferentP)
		t.Logf("%#v", aaa.DifferentPT)
		t.Logf("%#v", testto.DifferentPT)
		t.Fatal("failed Convert")
	}
	resultstruct2 := testFromStruct{}
	result2, err := ConvertStructToStruct(&testto, &resultstruct2, "", "teststruct")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if reflect.DeepEqual(result2.(testFromStruct), testfrom) == false {
		aaa := result2.(testFromStruct)
		t.Logf("%#v", aaa.Different)
		t.Logf("%#v", testfrom.Different)
		t.Logf("%#v", aaa.DifferentP)
		t.Logf("%#v", testfrom.DifferentP)
		//		t.Fatal("failed Convert")
	}
	resultstruct3 := testToStruct{}
	_, err = ConvertStructToStruct(testfrom, resultstruct3, "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}

	resultstring := "eeee"
	_, err = ConvertStructToStruct(testfrom, &resultstring, "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
	_, err = ConvertStructToStruct(testfrom, resultstring, "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
	_, err = ConvertStructToStruct(resultstring, testfrom, "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
	_, err = ConvertStructToStruct(&testfrom, "", "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
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
	if reflect.DeepEqual(result, testresult) == false {
		t.Fatal("failed Merge")
	}
	testTo2 := testToStruct{}
	_, err = MergeStructToStruct(testfrom, testTo2, "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
	_, err = MergeStructToStruct(&testfrom, "", "teststruct", "")
	if err == nil {
		t.Fatal("Not err")
	} else {
		t.Logf("failed test %#v", err)
	}
	t.Log("success Merge")
}
