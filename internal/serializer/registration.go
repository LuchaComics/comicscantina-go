package serializer

import (
    "errors"
    "net/http"
    "github.com/luchacomics/comicscantina-go/internal/model_resource"
)

// RegistrationRequest is the request payload for User data model.
//
// NOTE: It's good practice to have well defined request and response payloads
// so you can manage the specific inputs and outputs for clients, and also gives
// you the opportunity to transform data on input or output, for example
// on request, we'd like to protect certain fields and on output perhaps
// we'd like to include a computed field based on other values that aren't
// in the data model. Also, check out this awesome blog post on struct composition:
// http://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
type RegistrationRequest struct {
    Email string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
	FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
}

func (data *RegistrationRequest) Bind(r *http.Request) error {
    // Check to see if the user exists in the database.
    _, count := model_resource.DBLookupUserByEmail(data.Email)
    if count > 0 {
        return errors.New("Email is not unique.")
    }
	return nil
}

type RegistrationResponse struct {
    TokenString string `json:"token" form:"string"`
    UserID uint64 `json:"user_id,omitempty" form:"int"`
    Email string `json:"email" form:"email"`
    FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
}

func NewRegistrationResponse(tokenString string, userID uint64, email string, firstName string, lastName string) *RegistrationResponse {
	resp := &RegistrationResponse{
        TokenString: tokenString,
        UserID: userID,
        Email: email,
        FirstName: firstName,
        LastName: lastName,
    }
	return resp
}

func (rd *RegistrationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
