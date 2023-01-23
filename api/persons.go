package api

func (handler ApiEndpoints) CreatePerson(request CreatePersonRequest) (err error, response *CreatePersonResponse) {
	//create the person
	err, personId := handler.CoreLogic.CreatePerson(request.Person)

	//if we fail to create the person
	//we return the error
	if err != nil {
		return err, nil
	}

	//otherwise, we return success and the new resources ID
	return nil, &CreatePersonResponse{
		ID: personId,
	}
}
