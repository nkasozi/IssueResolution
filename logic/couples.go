package logic

import (
	"IssueResolution/db/db_models"
)

func (receiver Logic) CreateCouple(Couple db_models.Couple) (error, string) {
	err, coupleID := receiver.Db.AddToCouples(Couple)
	return err, coupleID
}
