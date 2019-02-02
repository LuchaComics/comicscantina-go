package controller

import (
    "fmt"
    "net/http"
    "strconv"
	"github.com/go-chi/render"
	"github.com/luchacomics/comicscantina-go/internal/model_manager"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
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
    // user := r.Context().Value("user").(*model.User)

    // Setup our variables for the paginator.
    pageString := r.FormValue("page")
    pageIndex, err := strconv.ParseUint(pageString, 10, 64)
    if err != nil {
        pageIndex = 0
    }

    organizations, _ := model_manager.OrganizationManagerInstance().AllByPageIndex(pageIndex)
    fmt.Println(organizations)

    // Iterate through each `Country` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewOrganizationListResponse(organizations)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}


func RetrieveOrganizationFunc(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("TODO: Implement - Retrieve"))
}
