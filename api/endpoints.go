package api

import (
	"IssueResolution/logic"
)

type Endpoints interface {
	CreateCouple(request CreateCoupleRequest) (err error, response *CreateCoupleResponse)
	CreateIssue(request CreateIssueRequest) (err error, response *CreateIssueResponse)
	CreatePerson(request CreatePersonRequest) (err error, response *CreatePersonResponse)
	CreatePersonsBelongingToCouple(request CreatePersonsBelongingToCoupleRequest) (err error, response *CreatePersonsBelongingToCoupleResponse)
}

//goland:noinspection ALL
type ApiEndpoints struct {
	CoreLogic logic.CoreLogic
}
