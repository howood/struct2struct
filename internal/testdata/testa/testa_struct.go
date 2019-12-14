package testa

type DifferentStruct struct {
	UserID   int32                `json:"userId,omitempty" teststruct:"UserID"`
	UserName string               `json:"userName,omitempty" teststruct:"UserName"`
	Sub      *DifferentStructASub `json:"userName,omitempty" teststruct:"Sub"`
}

type DifferentStructASub struct {
	UserID   int32  `json:"userId,omitempty" teststruct:"UserID"`
	UserName string `json:"userName,omitempty" teststruct:"UserName"`
}
