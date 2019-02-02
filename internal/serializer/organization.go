package serializer

import (
    "errors"
    "net/http"
    "github.com/luchacomics/comicscantina-go/internal/model_resource"
)

// OrganizationRequest is the request payload for Organization data model.
type OrganizationRequest struct {
    Name                string `json:"name"; form:"name";`
    Description         string `json:"description,omitempty"; form:"description";`
    Email               string `json:"email"; form:"email";`
    StreetAdddress      string `json:"street_address"; form:"street_address";`
    StreetAdddressExtra string `json:"street_address_extra"; form:"street_address_extra";`
    City                string `json:"city"; form:"city";`
    Province            string `json:"province"; form:"province";`
    Country             string `json:"country"; form:"country";`
    Currency            string `json:"currency"; form:"currency";`
    Language            string `json:"language"; form:"language";`
    Website             string `json:"website"; form:"website";`
    Phone               string `json:"phone"; form:"phone";`
    Fax                 string `json:"fax"; form:"fax";`
}

// Function will validate the input payload.
func (data *OrganizationRequest) Bind(r *http.Request) error {
    if data.Name == "" {
        return errors.New("Missing name.")
    }
    _, count := model_resource.DBLookupOrganizationByName(data.Name)
    if count > 0 {
        return errors.New("Name is not unique.")
    }
    if data.Email == "" {
        return errors.New("Missing email.")
    }
    if data.StreetAdddress == "" {
        return errors.New("Missing street address.")
    }
    if data.City == "" {
        return errors.New("Missing city.")
    }
    if data.Province == "" {
        return errors.New("Missing province.")
    }
    if data.Country == "" {
        return errors.New("Missing country.")
    }

    // Return with no errors.
	return nil
}

// OrganizationResponse is the response payload for Organization data model.
type OrganizationResponse struct {
    ID uint64 `json:"id,omitempty" form:"int"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Email string `json:"email" form:"email"`
    OwnerID uint64 `json:"owner_id,omitempty" form:"int"`
}

// Function will create our output payload.
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
