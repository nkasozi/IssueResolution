package db

import "IssueResolution/db/db_models"

type InMemoryDB struct {
	Persons                  []db_models.Person
	Issues                   []db_models.Issues
	Couples                  []db_models.Couple
	PersonsBelongingToCouple []db_models.PersonsBelongingToCouple
}

func InitInMemoryDb() (db *InMemoryDB) {
	db = &InMemoryDB{
		Persons:                  []db_models.Person{},
		Issues:                   []db_models.Issues{},
		Couples:                  []db_models.Couple{},
		PersonsBelongingToCouple: []db_models.PersonsBelongingToCouple{},
	}
	return
}

func (receiver *InMemoryDB) GenerateID() string {
	return ""
}

func (receiver *InMemoryDB) AddToPersons(person db_models.Person) (error, string) {
	person.ID = receiver.GenerateID()
	receiver.Persons = append(receiver.Persons, person)
	return nil, person.ID
}

func (receiver *InMemoryDB) AddToIssues(issue db_models.Issues) (error, string) {
	issue.ID = receiver.GenerateID()
	receiver.Issues = append(receiver.Issues, issue)
	return nil, issue.ID
}

func (receiver *InMemoryDB) AddToCouples(couple db_models.Couple) (error, string) {
	couple.ID = receiver.GenerateID()
	receiver.Couples = append(receiver.Couples, couple)
	return nil, couple.ID
}

func (receiver *InMemoryDB) AddToPersonsBelongingToCouple(person db_models.PersonsBelongingToCouple) (error, string) {
	person.ID = receiver.GenerateID()
	receiver.PersonsBelongingToCouple = append(receiver.PersonsBelongingToCouple, person)
	return nil, person.ID
}
