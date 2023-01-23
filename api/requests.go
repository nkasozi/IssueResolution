package api

import "IssueResolution/db/db_models"

type CreatePersonRequest struct {
	Person db_models.Person
}

type CreateCoupleRequest struct {
	Couple db_models.Couple
}

type CreatePersonsBelongingToCoupleRequest struct {
	CoupleId  string
	PersonIds []string
}

type CreateIssueRequest struct {
	Issue db_models.Issues
}
