package db

import "IssueResolution/db/db_models"

type DB interface {
	AddToPersons(person db_models.Person) (error, string)
	AddToIssues(issue db_models.Issues) (error, string)
	AddToCouples(couple db_models.Couple) (error, string)
	AddToPersonsBelongingToCouple(person db_models.PersonsBelongingToCouple) (error, string)
}
