package db_models

type Couple struct {
	ID   string `valid:"required"`
	Name string `valid:"required"`
}
