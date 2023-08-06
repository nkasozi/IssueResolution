package logic

import (
	"IssueResolution/db/db_models"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockDb struct {
	mock.Mock
}

func (m *mockDb) AddToPersons(person db_models.Person) (error, string) {
	//TODO implement me
	panic("implement me")
}

func (m *mockDb) AddToIssues(issue db_models.Issues) (error, string) {
	//TODO implement me
	panic("implement me")
}

func (m *mockDb) AddToPersonsBelongingToCouple(person db_models.PersonsBelongingToCouple) (error, string) {
	//TODO implement me
	panic("implement me")
}

func (m *mockDb) AddToCouples(couple db_models.Couple) (error, string) {
	args := m.Called(couple)
	return args.Error(0), args.String(1)
}

func TestLogic_CreateCouple_Success(t *testing.T) {
	//Given (may do mocking dependencies)
	myMockDb := &mockDb{}
	logic := Logic{Db: myMockDb}
	request := db_models.Couple{
		ID:   "001",
		Name: "John",
	}
	expected := "12345"
	myMockDb.On("AddToCouples", request).Return(nil, expected)

	//When
	err, coupleId := logic.CreateCouple(request)

	//Then
	assert.NoError(t, err)
	assert.Equal(t, expected, coupleId)
	myMockDb.AssertExpectations(t)
}

func TestLogic_CreateCouple_Failure_Validation(t *testing.T) {
	//Given (may do mocking dependencies)
	myMockDb := &mockDb{}
	logic := Logic{Db: myMockDb}
	request := db_models.Couple{
		ID:   "",
		Name: "John",
	}

	//When
	err, _ := logic.CreateCouple(request)

	//Then
	assert.NotNil(t, err)
}

func TestLogic_CreateCouple_Failure_Database(t *testing.T) {
	//Given (may do mocking dependencies)
	myMockDb := &mockDb{}
	logic := Logic{Db: myMockDb}
	request := db_models.Couple{
		ID:   "001",
		Name: "John",
	}
	expected := errors.New("some random error")
	myMockDb.On("AddToCouples", request).Return(expected, "")

	//When
	err, _ := logic.CreateCouple(request)

	//Then
	assert.NotNil(t, err)
	myMockDb.AssertExpectations(t)
}
