package serializer

import (
    "context"
    "errors"
    "net/http"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/model"
    "github.com/luchacomics/comicscantina-go/internal/model_resource"
)

// OrganizationRequest is the request payload for Organization data model.
type OrganizationRequest struct {
    Name                string `json:"name"; form:"name";`
    Description         string `json:"description,omitempty"; form:"description";`
    Email               string `json:"email"; form:"email";`
    StreetAddress       string `json:"street_address"; form:"street_address";`
    StreetAddressExtra  string `json:"street_address_extra"; form:"street_address_extra";`
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
    if data.StreetAddress == "" {
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

// Function will create Organization data model from the input payload.
func (data *OrganizationRequest) Save(ctx context.Context) (*model.Organization, error) {
    // Extract the current user from the request context.
    user := ctx.Value("user").(*model.User)

    // The model we will be creating.
    var organization model.Organization

    // Create our `User` object in our database.
    organization = model.Organization {
        Name:               data.Name,
        Description:        data.Description,
        Email:              data.Email,
        OwnerID:            user.ID,
        // CreatedAt:    time.Now(),
        // UpdatedAt:    time.Now(),
        StreetAddress:      data.StreetAddress,
        StreetAddressExtra: data.StreetAddressExtra,
        City:               data.City,
        Province:           data.Province,
        Country:            data.Country,
        Currency:           data.Currency,
        Language:           data.Language,
        Website:            data.Website,
        Phone:              data.Phone,
        Fax:                data.Fax,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&organization)

    return &organization, nil
}

// OrganizationResponse is the response payload for Organization data model.
type OrganizationResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Name                string `json:"name,omitempty"`
    Description         string `json:"description,omitempty"`
    Email               string `json:"email" form:"email"`
    Status              uint8 `gorm:"default: 1;"`
    OwnerID             uint64 `json:"owner_id,omitempty" form:"int"`
    StreetAddress      string `json:"street_address,omitempty"`
    StreetAddressExtra string `json:"street_address_extra,omitempty"`
    City                string `json:"city,omitempty"`
    Province            string `json:"province,omitempty"`
    Country             string `json:"country,omitempty"`
    Currency            string `json:"currency,omitempty"`
    Language            string `json:"language,omitempty"`
    Website             string `json:"website,omitempty"`
    Phone               string `json:"phone,omitempty"`
    Fax                 string `json:"fax,omitempty"`
    // CreatedAt           time.Time
    // UpdatedAt           time.Time
    Facebook            string `json:"facebook,omitempty"`
    Twitter             string `json:"twitter,omitempty"`
    YouTube             string `json:"youtube,omitempty"`
    Google              string `json:"google,omitempty"`
}

// Function will create our output payload.
func NewOrganizationResponse(organization *model.Organization) *OrganizationResponse {
	resp := &OrganizationResponse{
        ID:                 organization.ID,
        Name:               organization.Name,
        Description:        organization.Description,
        Email:              organization.Email,
        OwnerID:            organization.OwnerID,
        StreetAddress:      organization.StreetAddress,
        StreetAddressExtra: organization.StreetAddressExtra,
        City:               organization.City,
        Province:           organization.Province,
        Country:            organization.Country,
        Currency:           organization.Currency,
        Language:           organization.Language,
        Website:            organization.Website,
        Phone:              organization.Phone,
        Fax:                organization.Fax,
    }
	return resp
}

func (rd *OrganizationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
