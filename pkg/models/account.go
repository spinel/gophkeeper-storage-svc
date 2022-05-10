package models

type Account struct {
	//gorm.Model
	Id       int64  `json:"id" gorm:"primaryKey"`
	EntityID int64  `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
