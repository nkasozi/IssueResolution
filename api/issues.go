package api

func (handler ApiEndpoints) CreateIssue(request CreateIssueRequest) (err error, response *CreateIssueResponse) {
	//create the issue
	err, issueId := handler.CoreLogic.CreateIssue(request.Issue)

	//if we fail to create the issue
	//we return the error
	if err != nil {
		return err, nil
	}

	//otherwise, we return success and the new resources ID
	return nil, &CreateIssueResponse{
		ID: issueId,
	}
	return
}
