package testb

type DifferentStruct struct {
	UserID   int32                `json:"userId,omitempty"`
	UserName string               `json:"userName,omitempty"`
	Sub      *DifferentStructBSub `json:"userName,omitempty"`
}

type DifferentStruct2 struct {
	UserID   int32                `json:"userId,omitempty"`
	UserName string               `json:"userName,omitempty"`
	Sub      *DifferentStructASub `json:"userName,omitempty"`
}

type DifferentStructBSub struct {
	UserID   int32  `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type DifferentStructASub struct {
	UserID   int32  `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
}
