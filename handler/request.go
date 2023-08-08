package handler

import "fmt"

func errParamsRequire(name, typ string) error {
	return fmt.Errorf("param: %s (type:  %s) is required!", name, typ)

}

// Create Opening

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamsRequire("role", "string")
	}
	if r.Company == "" {
		return errParamsRequire("company", "string")
	}
	if r.Location == "" {
		return errParamsRequire("location", "string")
	}
	if r.Link == "" {
		return errParamsRequire("link", "string")
	}
	if r.Remote == nil {
		return errParamsRequire("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamsRequire("salary", "int64")
	}
	return nil
}

//Update Opening

type UpdateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// if any field is provided, validation is trusthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary < 0 {
		return nil
	}
	// If none of the filds were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
