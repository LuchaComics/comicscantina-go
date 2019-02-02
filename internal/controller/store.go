package controller

import (
    "context"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-chi/chi"
	"github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/model"
	"github.com/luchacomics/comicscantina-go/internal/model_manager"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
)


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


func ListStoresFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our URL parameter.
    pageIndex := r.Context().Value("pageIndex").(uint64)

    stores, _ := model_manager.StoreManagerInstance().AllByPageIndex(pageIndex)
    fmt.Println(stores)

    // Iterate through each `Country` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewStoreListResponse(stores)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


func CreateStoreFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &serializer.StoreDetailRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    store, _ := data.Save(r.Context())

    // Take newly created Store model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewStoreResponse(store))
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `storeID` parameter from the URL and
// attempt to lookup the Store model data object in the database. If
// the object was found then attach it to the context, else return an error.
func StoreCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if storeIDString := chi.URLParam(r, "storeID"); storeIDString != "" {
			storeID, _ := strconv.ParseUint(storeIDString, 10, 64)
			store, count := model_manager.StoreManagerInstance().GetByID(storeID)
            if count == 1 {
                ctx := context.WithValue(r.Context(), "store", store)
        		next.ServeHTTP(w, r.WithContext(ctx))
            }
		}
        render.Render(w, r, serializer.ErrNotFound)
        return
	})
}


func RetrieveStoreFunc(w http.ResponseWriter, r *http.Request) {
    store := r.Context().Value("store").(*model.Store)

	if err := render.Render(w, r, serializer.NewStoreResponse(store)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}
