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


// CategoryDetailRequest is the request payload for Category data model.
type CategoryDetailRequest struct {
    Name              string `json:"name"; form:"string";`
    ShortDescription  string `json:"short_description"; form:"string";`
    LongDescription   string `json:"long_description"; form:"string";`
    OrganizationID    uint64 `json:"organization_id,string,omitempty"`
}

// Function will validate the input payload.
func (data *CategoryDetailRequest) Bind(r *http.Request) error {
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

// Function will create Category data model from the input payload.
func (data *CategoryDetailRequest) Save(ctx context.Context) (*model.Category, error) {
    // The model we will be creating.
    var category model.Category

    // Create our `User` object in our database.
    category = model.Category {
        Name:             data.Name,
        ShortDescription: data.ShortDescription,
        LongDescription:  data.LongDescription,
        OrganizationID:   data.OrganizationID,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&category)

    return &category, nil
}


//----------------------------------------------------------------------------//
//                                  DETAILS                                   //
//----------------------------------------------------------------------------//


// CategoryResponse is the response payload for Category data model.
type CategoryResponse struct {
    ID                uint64 `json:"id,omitempty" form:"int"`
    Name              string `json:"name"; form:"string";`
    ShortDescription  string `json:"short_description"; form:"string";`
    LongDescription   string `json:"long_description"; form:"string";`
    OrganizationID    uint64 `json:"organization_id,string,omitempty"`
}

// Function will create our output payload.
func NewCategoryResponse(category *model.Category) *CategoryResponse {
	resp := &CategoryResponse{
        ID:               category.ID,
        Name:             category.Name,
        ShortDescription: category.ShortDescription,
        LongDescription:  category.LongDescription,
        OrganizationID:   category.OrganizationID,
    }
	return resp
}

func (rd *CategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Individual category list response payload.
type CategoryListItemResponse struct {
    ID                uint64 `json:"id,omitempty" form:"int"`
    Name              string `json:"name"; form:"string";`
    ShortDescription  string `json:"short_description"; form:"string";`
    LongDescription   string `json:"long_description"; form:"string";`
    OrganizationID    uint64 `json:"organization_id,string,omitempty"`
}

// Constructor creates a CategoryListItemResponse payload from the
// Category model data.
func NewCategoryListItemResponse(object *model.Category) *CategoryListItemResponse {
	resp := &CategoryListItemResponse{
        ID: object.ID,
        Name: object.Name,
        ShortDescription: object.ShortDescription,
        LongDescription: object.LongDescription,
        OrganizationID: object.OrganizationID,
    }
	return resp
}

func (rd *CategoryListItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// Constructor creates a JSON response payload from the array of Category
// model data objects.
func NewCategoryListResponse(categorys []model.Category) []render.Renderer {
	list := []render.Renderer{}
	for _, category := range categorys {
		list = append(list, NewCategoryListItemResponse(&category))
	}
	return list
}
