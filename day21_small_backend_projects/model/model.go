package model

type Doughnut struct {
	D_name string `json:"name"`
	D_type string `json:"type"`
}

func NewDoughnut(d_name, d_type string) Doughnut {
	return Doughnut{
		D_name: d_name,
		D_type: d_type,
	}
}
