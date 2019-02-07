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


func CreateShipperFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &serializer.ShipperDetailRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    shipper, _ := data.Save(r.Context())

    // Take newly created Shipper model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewShipperResponse(shipper))
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `shipperID` parameter from the URL and
// attempt to lookup the Shipper model data object in the database. If
// the object was found then attach it to the context, else return an error.
func ShipperCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shipperIDString := chi.URLParam(r, "shipperID"); shipperIDString != "" {
			shipperID, _ := strconv.ParseUint(shipperIDString, 10, 64)
			shipper, count := model_manager.ShipperManagerInstance().GetByID(shipperID)
            if count == 1 {
                ctx := context.WithValue(r.Context(), "shipper", shipper)
        		next.ServeHTTP(w, r.WithContext(ctx))
                return
            }
		}
        render.Render(w, r, serializer.ErrNotFound)
        return
	})
}


func RetrieveShipperFunc(w http.ResponseWriter, r *http.Request) {
    shipper := r.Context().Value("shipper").(*model.Shipper)

	if err := render.Render(w, r, serializer.NewShipperResponse(shipper)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Middleware will extract all the URL parameters from the URL and attach the
// values to the request context.
func ShipperFiltersCtx(next http.Handler) http.Handler {
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


func ListShippersFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our URL parameter.
    pageIndex := r.Context().Value("pageIndex").(uint64)

    // Filter the Shipper model data based on the context of the user:
    // (1) Owners have all their categories listed.
    // (2) Employers have all their categories listed.
    // (3) All categories listed if user is staff.
    var categories []model.Shipper;
    categories, _ = model_manager.ShipperManagerInstance().PaginatedAll(pageIndex)

    // Iterate through each `Shipper` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewShipperListResponse(categories)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}
