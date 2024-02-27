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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := m.QueryContext(ctx, "SELECT * FROM doughnut_list")
	defer res.Close()
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

func (m *MySQLDoughnut) GetDoughnutsWithType(d_type string) (doughnuts []model.Doughnut, res_err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stmt, err := m.PrepareContext(ctx, "SELECT * FROM doughnut_list WHERE doughnut_type = ?")
	defer stmt.Close()
	if err != nil {
		res_err = err
		return
	}
	res, err := stmt.Query(d_type)
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
func (m *MySQLDoughnut) AddDoughnuts(d []model.Doughnut) error {
	//bomb for querying...
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	//prepare first
	query := "INSERT INTO doughnut_list VALUES (?, ?)"
	stmt, err := m.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	//the excute....
	for _, x := range d {
		_, err := stmt.ExecContext(ctx, x.D_name, x.D_type)
		if err != nil {
			return err
		}

	}
	return nil
}
