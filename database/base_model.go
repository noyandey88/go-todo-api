package database

type BaseModel struct {
	ID        uint  `json:"id" gorm:"primaryKey"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}
