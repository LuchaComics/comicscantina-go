package controller

import (
    "context"
    "net/http"
    "strconv"
    // "github.com/go-chi/chi"
	"github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/model"
	"github.com/luchacomics/comicscantina-go/internal/model_manager"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
)


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


func CreateProductFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &serializer.ProductDetailRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    product, _ := data.Save(r.Context())

    // Take newly created Product model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, serializer.NewProductResponse(product))
}


// //----------------------------------------------------------------------------//
// //                                 RETRIEVE                                   //
// //----------------------------------------------------------------------------//
//
//
// // Middleware will extract the `productID` parameter from the URL and
// // attempt to lookup the Product model data object in the database. If
// // the object was found then attach it to the context, else return an error.
// func ProductCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if productIDString := chi.URLParam(r, "productID"); productIDString != "" {
// 			productID, _ := strconv.ParseUint(productIDString, 10, 64)
// 			product, count := model_manager.ProductManagerInstance().GetByID(productID)
//             if count == 1 {
//                 ctx := context.WithValue(r.Context(), "product", product)
//         		next.ServeHTTP(w, r.WithContext(ctx))
//             }
// 		}
//         render.Render(w, r, serializer.ErrNotFound)
//         return
// 	})
// }
//
//
// func RetrieveProductFunc(w http.ResponseWriter, r *http.Request) {
//     product := r.Context().Value("product").(*model.Product)
//
// 	if err := render.Render(w, r, serializer.NewProductResponse(product)); err != nil {
// 		render.Render(w, r, serializer.ErrRender(err))
// 		return
// 	}
// }


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


// Middleware will extract all the URL parameters from the URL and attach the
// values to the request context.
func ProductFiltersCtx(next http.Handler) http.Handler {
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


func ListProductsFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our URL parameter.
    pageIndex := r.Context().Value("pageIndex").(uint64)
    user := r.Context().Value("user").(*model.User)
    storeID := r.Context().Value("storeID").(uint64)

    // Filter the Product model data based on the context of the user:
    // (1) Owners have all their products listed.
    // (2) Employers have all their products listed.
    // (3) All products listed if user is staff.
    var products []model.Product;
    if user.OrganizationID != 0 {
        products, _ = model_manager.ProductManagerInstance().PaginatedFilterBy(
            user.OrganizationID,
            storeID,
            pageIndex,
        )
    } else if user.EmployerID != 0 {
        products, _ = model_manager.ProductManagerInstance().PaginatedFilterBy(
            user.EmployerID,
            storeID,
            pageIndex,
        )
    } else if user.GroupID == 2 {
        products, _ = model_manager.ProductManagerInstance().PaginatedFilterBy(
            0, // Skip filtering by organization_id.
            storeID,
            pageIndex,
        )
    }

    // Iterate through each `Product` object and render our specific view.
    if err := render.RenderList(w, r, serializer.NewProductListResponse(products)); err != nil {
		render.Render(w, r, serializer.ErrRender(err))
		return
	}
}
