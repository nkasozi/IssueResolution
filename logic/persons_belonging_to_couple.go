package logic

import (
	"IssueResolution/db/db_models"
)

func (receiver Logic) CreatePersonsBelongingCouple(personsBelongingToCouple db_models.PersonsBelongingToCouple) (error, string) {
	err, ID := receiver.Db.AddToPersonsBelongingToCouple(personsBelongingToCouple)
	return err, ID
}
