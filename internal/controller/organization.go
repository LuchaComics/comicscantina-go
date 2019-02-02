package controller

import (
    "net/http"
	"github.com/go-chi/render"
	// "github.com/luchacomics/comicscantina-go/internal/model"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
	// "github.com/luchacomics/comicscantina-data/internal/pkg/database"
    // "github.com/luchacomics/comicscantina-go/internal/base/service"
)


func CreateOrganizationFunc(w http.ResponseWriter, r *http.Request) {
    // // Take the user POST data and serialize it.
    data := &serializer.OrganizationRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database.
    organization, _ := data.Save(r.Context())

    // Take our data and serialize it back into a response object to hand
    // back to the organization.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewOrganizationResponse(organization))
}


func ListOrganizationsFunc(w http.ResponseWriter, r *http.Request) {
    // // Extract the current user from the request context.
    // user := ctx.Value("user").(*model.User)


    // // Setup our variables for the paginator.
    // pageString := r.FormValue("page")
    // pageIndex, err := strconv.ParseUint(pageString, 10, 64)
    // if err != nil {
    //     pageIndex = 0
    // }
    //
	// // Get our list of Country objects.
	// countries, _ := model_resource.DBGetPaginatedCountryList(pageIndex)
    //
    // // Iterate through each `Country` object and render our specific view.
    // if err := render.RenderList(w, r, serializer.NewCountryListResponse(countries)); err != nil {
	// 	render.Render(w, r, serializer.ErrRender(err))
	// 	return
	// }



    w.Write([]byte("TODO: Implement - List"))
}


func RetrieveOrganizationFunc(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("TODO: Implement - Retrieve"))
}
