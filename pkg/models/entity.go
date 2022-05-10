package models

type Entity struct {
	//gorm.Model
	ID         int64      `json:"id" gorm:"primaryKey"`
	Identifier string     `json:"identifier" gorm:"index:idx_identifier,unique"`
	TypeID     int        `json:"type_id"`
	Account    Account    `json:"account"`
	CreditCard CreditCard `json:"credit_card"`
}
