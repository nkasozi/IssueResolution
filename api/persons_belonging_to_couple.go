package api

import "IssueResolution/db/db_models"

func (handler ApiEndpoints) CreatePersonsBelongingToCouple(request CreatePersonsBelongingToCoupleRequest) (err error, response *CreatePersonsBelongingToCoupleResponse) {
	//create the person
	resourceID := ""

	for _, personId := range request.PersonIds {

		personBelongingToCouple := db_models.PersonsBelongingToCouple{
			CoupleId: request.CoupleId,
			PersonId: personId,
		}

		err, resourceID = handler.CoreLogic.CreatePersonsBelongingCouple(personBelongingToCouple)

		//if we fail to create the person
		//we return the error
		if err != nil {
			return err, nil
		}
	}

	//otherwise, we return success and the new resources ID
	return nil, &CreatePersonsBelongingToCoupleResponse{
		ID: resourceID,
	}
	return
}
