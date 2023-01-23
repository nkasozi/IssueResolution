package logic

import "IssueResolution/db/db_models"

func (receiver Logic) CreatePerson(Person db_models.Person) (error, string) {
	err, ID := receiver.Db.AddToPersons(Person)
	return err, ID
}
