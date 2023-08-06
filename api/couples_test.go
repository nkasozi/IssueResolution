package api

import (
	"IssueResolution/db/db_models"
	"IssueResolution/logic/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApiEndpoints_CreateCouple_Success(t *testing.T) {
	//Given - setup
	mockCoreLogic := new(mocks.MockCoreLogic)
	handler := ApiEndpoints{CoreLogic: mockCoreLogic}
	request := CreateCoupleRequest{Couple: db_models.Couple{
		ID:   "001",
		Name: "John",
	}}
	expectedId := "12345"
	expected := &CreateCoupleResponse{
		ID: expectedId,
	}
	mockCoreLogic.On("CreateCouple", request.Couple).Return(nil, expectedId)

	//when - act
	err, response := handler.CreateCouple(request)

	//then -assert
	assert.Nil(t, err)
	assert.Equal(t, expected, response)
	mockCoreLogic.AssertExpectations(t)
}

func TestApiEndpoints_CreateCouple_Failure(t *testing.T) {
	//Given - setup
	mockCoreLogic := new(mocks.MockCoreLogic)
	handler := ApiEndpoints{CoreLogic: mockCoreLogic}
	request := CreateCoupleRequest{Couple: db_models.Couple{
		ID:   "001",
		Name: "John",
	}}
	expected := errors.New("some random error")
	mockCoreLogic.On("CreateCouple", request.Couple).Return(expected, "")

	//when - act
	err, response := handler.CreateCouple(request)

	//then -assert
	assert.Error(t, err)
	assert.Nil(t, response)
	mockCoreLogic.AssertExpectations(t)
}
