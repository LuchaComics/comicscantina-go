package controller

import (
	"net/http"
	"github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/model_resource"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
    "github.com/luchacomics/comicscantina-go/internal/base/service"
)


func ProfileRetrieveFunc(w http.ResponseWriter, r *http.Request) {
	// Get the 'user_id' value from the JWT in the middleware and if zero is
	// returned then there the JWT token was not properly set up.
	user_id := service.GetUserIDFromContext(r.Context())
	if user_id == 0 {
		http.Error(w, "User ID not inputted.", http.StatusUnauthorized)
		return
	}

    // Lookup the user ID in the database.
    user, _ := model_resource.DBLookupUserByID(user_id)

    // Take our data and serialize it back into a response object to hand
    // back to the user.
    render.Status(r, http.StatusOK)
	render.Render(w, r, serializer.NewProfileResponse(
        user.ID,
        user.Email,
        user.FirstName,
        user.LastName,
    ))
}
