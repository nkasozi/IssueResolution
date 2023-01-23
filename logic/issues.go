package logic

import (
	"IssueResolution/db/db_models"
)

func (receiver Logic) CreateIssue(Issue db_models.Issues) (error, string) {
	err, ID := receiver.Db.AddToIssues(Issue)
	return err, ID
}
