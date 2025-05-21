package models

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"success message"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJpc3N1ZXIiLCJhbGciOiJIUzI1NiIsInN1YiI6InVzZXIiLCJhdWQiOlsiYXV0aWVuY2UiXSwiaWF0IjoxNzQ1NDY3MjgzLCJleHBpcmF0aW9uIjoxNzQ1NDY3MjgzLCJqdGkiOiJpZCJ9.dxgM6uH2F8ZglV_xcPhjCRnOSJBYq9oeS1TDLkLg_eg"`
}

type VerifyResponse struct {
	Verify bool `json:"verify" example:"true"`
}