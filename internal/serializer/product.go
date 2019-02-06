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


// ProductDetailRequest is the request payload for Product data model.
type ProductDetailRequest struct {
    Name                string `json:"name"; form:"name";`
    Description         string `json:"description,omitempty"; form:"description";`
    OrganizationID      uint64 `json:"organization_id,string,omitempty"`
    StoreID             uint64 `json:"store_id,string,omitempty"`
}

// Function will validate the input payload.
func (data *ProductDetailRequest) Bind(r *http.Request) error {
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

// Function will create Product data model from the input payload.
func (data *ProductDetailRequest) Save(ctx context.Context) (*model.Product, error) {
    // The model we will be creating.
    var product model.Product

    // Create our `User` object in our database.
    product = model.Product {
        Name:               data.Name,
        OrganizationID:     data.OrganizationID,
        StoreID:            data.StoreID,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&product)

    return &product, nil
}


//----------------------------------------------------------------------------//
//                                  DETAILS                                   //
//----------------------------------------------------------------------------//


// ProductResponse is the response payload for Product data model.
type ProductResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Name                string `json:"name,omitempty"`
    OrganizationID      uint64 `json:"organization_id,omitempty" form:"int"`
    StoreID             uint64 `json:"store_id,omitempty" form:"int"`
}

// Function will create our output payload.
func NewProductResponse(organization *model.Product) *ProductResponse {
	resp := &ProductResponse{
        ID:                 organization.ID,
        Name:               organization.Name,
        OrganizationID:     organization.OrganizationID,
        StoreID:     organization.StoreID,
    }
	return resp
}

func (rd *ProductResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Individual product list response payload.
type ProductListItemResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Name                string `json:"name,omitempty"`
    OrganizationID      uint64 `json:"organization_id,omitempty" form:"int"`
    StoreID             uint64 `json:"store_id,omitempty" form:"int"`
}

// Constructor creates a ProductListItemResponse payload from the
// Product model data.
func NewProductListItemResponse(object *model.Product) *ProductListItemResponse {
	resp := &ProductListItemResponse{
        ID: object.ID,
        Name: object.Name,
        OrganizationID: object.OrganizationID,
        StoreID: object.StoreID,
    }
	return resp
}

func (rd *ProductListItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// Constructor creates a JSON response payload from the array of Product
// model data objects.
func NewProductListResponse(products []model.Product) []render.Renderer {
	list := []render.Renderer{}
	for _, product := range products {
		list = append(list, NewProductListItemResponse(&product))
	}
	return list
}
