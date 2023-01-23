package db_models

type Outcomes struct {
	ID                  string
	SolutionId          string
	CoupleId            string
	WasSolutionAgreedTo bool
}
