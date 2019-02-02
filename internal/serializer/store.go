package serializer

import (
    // "context"
    // "errors"
    "net/http"
    "github.com/go-chi/render"
    // "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/model"
    // "github.com/luchacomics/comicscantina-go/internal/model_manager"
)

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
