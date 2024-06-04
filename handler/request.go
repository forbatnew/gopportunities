package handler

import "fmt"

func errParameterIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s(type: %s) is required", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Role == "" || r.Company == "" || r.Location == "" || r.Link == "" || r.Remote == nil || r.Salary <= 0 {
		return fmt.Errorf("malformed request body")
	}
	if r.Role == "" {
		return errParameterIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParameterIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParameterIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParameterIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParameterIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParameterIsRequired("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Id       int64  `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	//If any field is provided, validation is truly
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
