package logic

import (
	"IssueResolution/db"
	"IssueResolution/db/db_models"
)

type CoreLogic interface {
	CreatePerson(Person db_models.Person) (error, string)
	CreateCouple(Couple db_models.Couple) (error, string)
	CreateIssue(Issue db_models.Issues) (error, string)
	CreatePersonsBelongingCouple(personsBelongingToCouple db_models.PersonsBelongingToCouple) (error, string)
}

type Logic struct {
	Db db.DB
}
