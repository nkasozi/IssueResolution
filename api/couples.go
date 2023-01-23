package api

func (handler ApiEndpoints) CreateCouple(request CreateCoupleRequest) (err error, response *CreateCoupleResponse) {
	//create the couple
	err, coupleId := handler.CoreLogic.CreateCouple(request.Couple)

	//if we fail to create the couple
	//we return the error
	if err != nil {
		return err, nil
	}

	//otherwise, we return success and the new resources ID
	return nil, &CreateCoupleResponse{
		ID: coupleId,
	}
	return
}
