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


func CreateSupplierFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &serializer.SupplierDetailRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    supplier, _ := data.Save(r.Context())

    // Take newly created Supplier model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewSupplierResponse(supplier))
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `supplierID` parameter from the URL and
// attempt to lookup the Supplier model data object in the database. If
// the object was found then attach it to the context, else return an error.
func SupplierCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if supplierIDString := chi.URLParam(r, "supplierID"); supplierIDString != "" {
			supplierID, _ := strconv.ParseUint(supplierIDString, 10, 64)
			supplier, count := model_manager.SupplierManagerInstance().GetByID(supplierID)
            if count == 1 {
                ctx := context.WithValue(r.Context(), "supplier", supplier)
        		next.ServeHTTP(w, r.WithContext(ctx))
                return
            }
		}
        render.Render(w, r, serializer.ErrNotFound)
        return
	})
}


func RetrieveSupplierFunc(w http.ResponseWriter, r *http.Request) {
    supplier := r.Context().Value("supplier").(*model.Supplier)

	if err := render.Render(w, r, serializer.NewSupplierResponse(supplier)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Middleware will extract all the URL parameters from the URL and attach the
// values to the request context.
func SupplierFiltersCtx(next http.Handler) http.Handler {
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


func ListSuppliersFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our URL parameter.
    pageIndex := r.Context().Value("pageIndex").(uint64)

    // Filter the Supplier model data based on the context of the user:
    // (1) Owners have all their categories listed.
    // (2) Employers have all their categories listed.
    // (3) All categories listed if user is staff.
    var categories []model.Supplier;
    categories, _ = model_manager.SupplierManagerInstance().PaginatedAll(pageIndex)

    // Iterate through each `Supplier` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewSupplierListResponse(categories)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}
