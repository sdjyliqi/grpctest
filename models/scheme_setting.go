package models

//
//import "github.com/go-xorm/xorm"
//
//type ProductStrategy struct {
//	Id            int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
//	ProductId      string `json:"product_id" xorm:"not null unique VARCHAR(32)"`
//	PeriodWeight   string `json:"period_weight" xorm:"VARCHAR(1024)"`
//	PeriodConsumer string `json:"period_consumer" xorm:"VARCHAR(1024)"`
//	Status         int    `json:"status" xorm:"TINYINT(4)"`
//	Periods        int    `json:"periods" xorm:"INT(11)"`
//}
//
//
//
//
//// SaveToDB ..
//func (p *ProductStrategy) SaveToDB(db *xorm.Engine) (err error) {
//	_, err = db.Insert(p)
//	return
//}
//
