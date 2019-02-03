package controller

import (
    "context"
    "net/http"
    "strconv"
    "github.com/go-chi/chi"
	"github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/model"
	"github.com/luchacomics/comicscantina-go/internal/model_manager"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
)


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


func CreateOrganizationFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &serializer.OrganizationRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    organization, _ := data.Save(r.Context())

    // Take newly created Organization model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewOrganizationDetailResponse(organization))
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//

// PUBLIC

func ListPublicOrganizationsFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our URL parameter.
    pageIndex := r.Context().Value("pageIndex").(uint64)

    organizations, _ := model_manager.OrganizationManagerInstance().FilterActiveStatusByPageIndex(pageIndex)

    // Iterate through each `Country` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewPublicOrganizationListResponse(organizations)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}

// PROTECTED

func ListOrganizationsFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our URL parameter.
    pageIndex := r.Context().Value("pageIndex").(uint64)

    organizations, _ := model_manager.OrganizationManagerInstance().AllByPageIndex(pageIndex)

    // Iterate through each `Country` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewOrganizationListResponse(organizations)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `organizationID` parameter from the URL and
// attempt to lookup the Organization model data object in the database. If
// the object was found then attach it to the context, else return an error.
func OrganizationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if organizationIDString := chi.URLParam(r, "organizationID"); organizationIDString != "" {
			organizationID, _ := strconv.ParseUint(organizationIDString, 10, 64)
			organization, count := model_manager.OrganizationManagerInstance().GetByID(organizationID)
            if count == 1 {
                ctx := context.WithValue(r.Context(), "organization", organization)
        		next.ServeHTTP(w, r.WithContext(ctx))
            }
		}
        render.Render(w, r, serializer.ErrNotFound)
        return
	})
}


func RetrieveOrganizationFunc(w http.ResponseWriter, r *http.Request) {
    organization := r.Context().Value("organization").(*model.Organization)

	if err := render.Render(w, r, serializer.NewOrganizationDetailResponse(organization)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}
