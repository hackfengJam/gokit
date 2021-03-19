package isql_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hackfengJam/gokit/isql"
)

const (
	DateFormatDateStandard = "2006-01-02"
)

type Order struct {
	ID         uint64    `db:"id" json:"id"`
	OrderID    uint64    `db:"id" json:"order_id"`
	BuyerID    uint64    `db:"buyer_id" json:"buyer_id"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}

var emptyOrder = Order{}

func (o *Order) TableName() string {
	return fmt.Sprintf("order_%s_%d", o.CreateTime.Format(DateFormatDateStandard), o.BuyerID)
}

func TestNewSqlHandler(t *testing.T) {
	sqlFormat := `
SELECT order_id FROM %s WHERE buyer_id=?`

	now := time.Now()
	timeList := []time.Time{
		now.AddDate(0, -1, 0),
		now,
		now.AddDate(0, 1, 0),
	}

	buyerID := uint64(10000010)
	args := []interface{}{buyerID}
	// sqlHandler
	sqlHandler := isql.NewSQLHandler()
	for _, tt := range timeList {
		o := Order{
			BuyerID:    buyerID,
			CreateTime: tt,
		}
		tableName := o.TableName()
		sqlHandler.Append(fmt.Sprintf(sqlFormat, tableName), args...)
	}

	sql, params := sqlHandler.Build(" UNION ALL ")
	sql = fmt.Sprintf(`SELECT * FROM ( %s ) as tmp;`, sql)

	t.Log(sql, params)
}
