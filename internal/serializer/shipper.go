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


// ShipperDetailRequest is the request payload for Shipper data model.
type ShipperDetailRequest struct {
    Name              string `json:"name"; form:"string";`
    OrganizationID    uint64 `json:"organization_id,string,omitempty"; form:"int"`
}

// Function will validate the input payload.
func (data *ShipperDetailRequest) Bind(r *http.Request) error {
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

// Function will create Shipper data model from the input payload.
func (data *ShipperDetailRequest) Save(ctx context.Context) (*model.Shipper, error) {
    // The model we will be creating.
    var shipper model.Shipper

    // Create our `User` object in our database.
    shipper = model.Shipper {
        Name:             data.Name,
        OrganizationID:   data.OrganizationID,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&shipper)

    return &shipper, nil
}


//----------------------------------------------------------------------------//
//                                  DETAILS                                   //
//----------------------------------------------------------------------------//


// ShipperResponse is the response payload for Shipper data model.
type ShipperResponse struct {
    ID                uint64 `json:"id,omitempty" form:"int"`
    Name              string `json:"name";`
    OrganizationID    uint64 `json:"organization_id,omitempty"; form:"int"`
}

// Function will create our output payload.
func NewShipperResponse(shipper *model.Shipper) *ShipperResponse {
	resp := &ShipperResponse{
        ID:               shipper.ID,
        Name:             shipper.Name,
        OrganizationID:   shipper.OrganizationID,
    }
	return resp
}

func (rd *ShipperResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Individual shipper list response payload.
type ShipperListItemResponse struct {
    ID                uint64 `json:"id,omitempty" form:"int"`
    Name              string `json:"name"; form:"string";`
    OrganizationID    uint64 `json:"organization_id,omitempty"; form:"int"`
}

// Constructor creates a ShipperListItemResponse payload from the
// Shipper model data.
func NewShipperListItemResponse(object *model.Shipper) *ShipperListItemResponse {
	resp := &ShipperListItemResponse{
        ID: object.ID,
        Name: object.Name,
        OrganizationID: object.OrganizationID,
    }
	return resp
}

func (rd *ShipperListItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// Constructor creates a JSON response payload from the array of Shipper
// model data objects.
func NewShipperListResponse(shippers []model.Shipper) []render.Renderer {
	list := []render.Renderer{}
	for _, shipper := range shippers {
		list = append(list, NewShipperListItemResponse(&shipper))
	}
	return list
}
