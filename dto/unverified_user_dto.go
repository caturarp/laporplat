package dto

type VerifyRequest struct {
	Email string `binding:"required,email" json:"email"`
	Code  string `json:"code"`
}

type VerifyResponse struct {
	Code string `json:"code"`
}

type DeleteUnverifiedUserRequest struct {
	Email string
}

func TransformInfoToDeleteRequest(input string) DeleteUnverifiedUserRequest {
	return DeleteUnverifiedUserRequest{
		Email: input,
	}
}
