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


func CreateCategoryFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &serializer.CategoryDetailRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    category, _ := data.Save(r.Context())

    // Take newly created Category model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewCategoryResponse(category))
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `categoryID` parameter from the URL and
// attempt to lookup the Category model data object in the database. If
// the object was found then attach it to the context, else return an error.
func CategoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if categoryIDString := chi.URLParam(r, "categoryID"); categoryIDString != "" {
			categoryID, _ := strconv.ParseUint(categoryIDString, 10, 64)
			category, count := model_manager.CategoryManagerInstance().GetByID(categoryID)
            if count == 1 {
                ctx := context.WithValue(r.Context(), "category", category)
        		next.ServeHTTP(w, r.WithContext(ctx))
            }
		}
        render.Render(w, r, serializer.ErrNotFound)
        return
	})
}


func RetrieveCategoryFunc(w http.ResponseWriter, r *http.Request) {
    category := r.Context().Value("category").(*model.Category)

	if err := render.Render(w, r, serializer.NewCategoryResponse(category)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Middleware will extract all the URL parameters from the URL and attach the
// values to the request context.
func CategoryFiltersCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the request context.
        ctx := r.Context()

        // Attempt to extract our URL parameters.
        storeIDString := r.FormValue("store_id")
        storeID, err := strconv.ParseUint(storeIDString, 10, 64)
        if err != nil {
            storeID = 0
        }
		ctx = context.WithValue(ctx, "storeID", storeID)

        // Attach our updated context to the request.
        next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func ListCategoriesFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our URL parameter.
    pageIndex := r.Context().Value("pageIndex").(uint64)

    // Filter the Category model data based on the context of the user:
    // (1) Owners have all their categories listed.
    // (2) Employers have all their categories listed.
    // (3) All categories listed if user is staff.
    var categories []model.Category;
    categories, _ = model_manager.CategoryManagerInstance().PaginatedAll(pageIndex)

    // Iterate through each `Category` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewCategoryListResponse(categories)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}
