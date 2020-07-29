package model

//Base base model
type Base struct {
	CreatedAt string `json:"createdAt" gorm:"-"`
	UpdatedAt string `json:"updatedAt" gorm:"-"`
	DeletedAt string `json:"deletedAt" gorm:"-"`
}
