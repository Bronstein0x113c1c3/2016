package db

import "day21/model"

type IDoughnutDatabase interface {
	GetDoughnuts() ([]model.Doughnut, error)
	GetDoughnutsWithType(d_type string) ([]model.Doughnut, error)
	AddDoughnuts(d []model.Doughnut) error
}
