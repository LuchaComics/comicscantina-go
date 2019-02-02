package controller

import (
    // "context"
    "fmt"
    "net/http"
    // "strconv"
    // "github.com/go-chi/chi"
	"github.com/go-chi/render"
    // "github.com/luchacomics/comicscantina-go/internal/model"
	"github.com/luchacomics/comicscantina-go/internal/model_manager"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
)


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
