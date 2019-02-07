package serializer

import (
    "context"
    "errors"
    "net/http"
    "github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/model"
    "github.com/luchacomics/comicscantina-go/internal/model_manager"
)


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


// SupplierDetailRequest is the request payload for Supplier data model.
type SupplierDetailRequest struct {
    Name              string `json:"name"; form:"string";`
    OrganizationID    uint64 `json:"organization_id,string,omitempty"; form:"int"`
}

// Function will validate the input payload.
func (data *SupplierDetailRequest) Bind(r *http.Request) error {
    // Extract the User model data object from the request object.
    user := r.Context().Value("user").(*model.User)

    if data.Name == "" {
        return errors.New("Missing name.")
    }
    if data.OrganizationID == 0 {
        return errors.New("Missing organization_id.")
    }
    // Check to see if this organization exists?
    // Check to see if the authenticated user belongs to this organization.
    has_membership := model_manager.OrganizationManagerInstance().UserIsMemberOf(user.ID, data.OrganizationID)
    if has_membership == false {
        return errors.New("`organization_id` is invalid - either does not exist or you are not a member of it.")
    }

    // Return with no errors.
	return nil
}

// Function will create Supplier data model from the input payload.
func (data *SupplierDetailRequest) Save(ctx context.Context) (*model.Supplier, error) {
    // The model we will be creating.
    var supplier model.Supplier

    // Create our `User` object in our database.
    supplier = model.Supplier {
        Name:             data.Name,
        OrganizationID:   data.OrganizationID,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&supplier)

    return &supplier, nil
}


//----------------------------------------------------------------------------//
//                                  DETAILS                                   //
//----------------------------------------------------------------------------//


// SupplierResponse is the response payload for Supplier data model.
type SupplierResponse struct {
    ID                uint64 `json:"id,omitempty" form:"int"`
    Name              string `json:"name";`
    OrganizationID    uint64 `json:"organization_id,omitempty"; form:"int"`
}

// Function will create our output payload.
func NewSupplierResponse(supplier *model.Supplier) *SupplierResponse {
	resp := &SupplierResponse{
        ID:               supplier.ID,
        Name:             supplier.Name,
        OrganizationID:   supplier.OrganizationID,
    }
	return resp
}

func (rd *SupplierResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Individual supplier list response payload.
type SupplierListItemResponse struct {
    ID                uint64 `json:"id,omitempty" form:"int"`
    Name              string `json:"name"; form:"string";`
    OrganizationID    uint64 `json:"organization_id,omitempty"; form:"int"`
}

// Constructor creates a SupplierListItemResponse payload from the
// Supplier model data.
func NewSupplierListItemResponse(object *model.Supplier) *SupplierListItemResponse {
	resp := &SupplierListItemResponse{
        ID: object.ID,
        Name: object.Name,
        OrganizationID: object.OrganizationID,
    }
	return resp
}

func (rd *SupplierListItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// Constructor creates a JSON response payload from the array of Supplier
// model data objects.
func NewSupplierListResponse(suppliers []model.Supplier) []render.Renderer {
	list := []render.Renderer{}
	for _, supplier := range suppliers {
		list = append(list, NewSupplierListItemResponse(&supplier))
	}
	return list
}
