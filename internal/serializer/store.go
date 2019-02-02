package serializer

import (
    "context"
    "errors"
    "net/http"
    "github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/model"
    // "github.com/luchacomics/comicscantina-go/internal/model_manager"
)


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


// StoreDetailRequest is the request payload for Store data model.
type StoreDetailRequest struct {
    Name                string `json:"name"; form:"name";`
    Description         string `json:"description,omitempty"; form:"description";`
    OrganizationID      uint64 `json:"organization_id,string,omitempty"`
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
func (data *StoreDetailRequest) Bind(r *http.Request) error {
    if data.Name == "" {
        return errors.New("Missing name.")
    }
    if data.OrganizationID == 0 {
        return errors.New("Missing organization_id.")
    }
    // TODO: Check to see if this organization exists?
    // TODO: Check to see if the authenticated user belongs to this organization.
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

// Function will create Store data model from the input payload.
func (data *StoreDetailRequest) Save(ctx context.Context) (*model.Store, error) {
    // The model we will be creating.
    var store model.Store

    // Create our `User` object in our database.
    store = model.Store {
        Name:               data.Name,
        Description:        data.Description,
        Email:              data.Email,
        // OrganizationID:     data.OrganizationID,
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
    db.Create(&store)

    return &store, nil
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Individual store list response payload.
type StoreListItemResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Name                string `json:"name,omitempty"`
    Description         string `json:"description,omitempty"`
    OrganizationID      uint64 `json:"organization_id,omitempty" form:"int"`
}

// Constructor creates a StoreListItemResponse payload from the
// Store model data.
func NewStoreListItemResponse(object *model.Store) *StoreListItemResponse {
	resp := &StoreListItemResponse{
        ID: object.ID,
        Name: object.Name,
        Description: object.Description,
        OrganizationID: object.OrganizationID,
    }
	return resp
}

func (rd *StoreListItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// Constructor creates a JSON response payload from the array of Store
// model data objects.
func NewStoreListResponse(stores []model.Store) []render.Renderer {
	list := []render.Renderer{}
	for _, store := range stores {
		list = append(list, NewStoreListItemResponse(&store))
	}
	return list
}


//----------------------------------------------------------------------------//
//                                  DETAILS                                   //
//----------------------------------------------------------------------------//


// StoreResponse is the response payload for Store data model.
type StoreResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Name                string `json:"name,omitempty"`
    Description         string `json:"description,omitempty"`
    Email               string `json:"email" form:"email"`
    Status              uint8 `json:"status" form:"int"`
    OrganizationID      uint64 `json:"organization_id,omitempty" form:"int"`
    StreetAddress       string `json:"street_address,omitempty"`
    StreetAddressExtra  string `json:"street_address_extra,omitempty"`
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
func NewStoreResponse(organization *model.Store) *StoreResponse {
	resp := &StoreResponse{
        ID:                 organization.ID,
        Name:               organization.Name,
        Description:        organization.Description,
        Email:              organization.Email,
        Status:             organization.Status,
        OrganizationID:     organization.OrganizationID,
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

func (rd *StoreResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
