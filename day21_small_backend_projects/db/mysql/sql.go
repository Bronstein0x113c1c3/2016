package mysql

import (
	"context"
	"database/sql"
	"day21/model"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// this shit for the SQL doughnut...
type MySQLDoughnut struct {
	*sql.DB
}

func Init_db(host string, port int, username string, password string, db_name string) (MySQLDoughnut, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, db_name))
	if err != nil {
		return MySQLDoughnut{}, err
	}
	return MySQLDoughnut{db}, nil
}

func (m *MySQLDoughnut) GetDoughnuts() (doughnuts []model.Doughnut, res_err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := m.QueryContext(ctx, "SELECT * FROM doughnut_list")
	if err != nil {
		res_err = err
		return
	}
	for res.Next() {
		var d_name, d_type string
		if err := res.Scan(&d_name, &d_type); err != nil {
			res_err = err
			return
		}
		doughnuts = append(doughnuts, model.NewDoughnut(d_name, d_type))
	}
	return
}
