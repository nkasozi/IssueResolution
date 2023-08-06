package mocks

import (
	"IssueResolution/db/db_models"
	"github.com/stretchr/testify/mock"
)

// MockCoreLogic mock core logic dependencies
type MockCoreLogic struct {
	mock.Mock
}

func (m *MockCoreLogic) CreatePerson(Person db_models.Person) (error, string) {
	args := m.Called(Person)
	return args.Error(0), args.String(1)
}

func (m *MockCoreLogic) CreateIssue(Issue db_models.Issues) (error, string) {
	args := m.Called(Issue)
	return args.Error(0), args.String(1)
}

func (m *MockCoreLogic) CreatePersonsBelongingCouple(personsBelongingToCouple db_models.PersonsBelongingToCouple) (error, string) {
	args := m.Called(personsBelongingToCouple)
	return args.Error(0), args.String(1)
}

func (m *MockCoreLogic) CreateCouple(couple db_models.Couple) (error, string) {
	args := m.Called(couple)
	return args.Error(0), args.String(1)
}
