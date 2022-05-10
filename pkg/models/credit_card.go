package models

type CreditCard struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	EntityID   int64  `json:"-"`
	Number     string `json:"number"`
	HolderName string `json:"holder_name"`
	Expire     string `json:"expire"`
	CVC        string `json:"cvc"`
}
