package serializer

import (
    // "errors"
    "net/http"
    // "github.com/luchacomics/comicscantina-go/internal/model_resource"
)


type OrganizationRequest struct {
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Email string `json:"email" form:"email"`
}

func (data *OrganizationRequest) Bind(r *http.Request) error {
    // // Check to see if the user exists in the database.
    // _, count := model_resource.DBLookupOrganizationByID(data.ID)
    // if count > 0 {
    //     return errors.New("ID is not unique.")
    // }
	return nil
}

type OrganizationResponse struct {
    ID uint64 `json:"id,omitempty" form:"int"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Email string `json:"email" form:"email"`
    OwnerID uint64 `json:"owner_id,omitempty" form:"int"`
}

func NewOrganizationResponse(id uint64, name string, description string, email string, userID uint64) *OrganizationResponse {
	resp := &OrganizationResponse{
        ID: id,
        Name: name,
        Description: description,
        Email: email,
        OwnerID: userID,
    }
	return resp
}

func (rd *OrganizationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
