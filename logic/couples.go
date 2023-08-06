package logic

import (
	"IssueResolution/db/db_models"
	"github.com/asaskevich/govalidator"
)

func (receiver Logic) CreateCouple(Couple db_models.Couple) (error, string) {
	isValid, err := govalidator.ValidateStruct(Couple)
	if !isValid {
		return err, ""
	}
	err, coupleID := receiver.Db.AddToCouples(Couple)
	return err, coupleID
}
