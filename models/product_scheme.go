package models

import (
	"github.com/go-xorm/xorm"
	"time"
)

type ProductScheme struct {
	Id             int       `xorm:"'id' pk autoincr"`
	ProductId      string    ` xorm:"not null unique VARCHAR(32)"`
	Catagory       string    ` xorm:"VARCHAR(32)"`
	PeriodWeight   string    ` xorm:"VARCHAR(1024)"`
	PeriodConsumer string    ` xorm:"VARCHAR(1024)"`
	Status         int       ` xorm:"TINYINT(4)"`
	Periods        int       ` xorm:"INT(11)"`
	LastModified   time.Time ` xorm:"DATETIME"`
	OptionVer      string    ` xorm:"VARCHAR(64)"`
}

// SaveToDB ..
func (w *ProductScheme) SaveToDB(db *xorm.Engine) (err error) {
	_, err = db.Insert(w)
	return
}
