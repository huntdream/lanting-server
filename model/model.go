package model

// Base model
type Base struct {
	CreatedAt string `json:"createdAt" `
	UpdatedAt string `json:"updatedAt" `
}

type Ids struct {
	Ids []int64 `json:"ids"`
}
