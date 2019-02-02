package middleware

import (
    "strconv"
    "context"
	"net/http"
)

// Middleware used to extract the `page` paramter from the URL and save it
// in the context.
func PaginationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Setup our variables for the paginator.
        pageString := r.FormValue("page")
        pageIndex, err := strconv.ParseUint(pageString, 10, 64)
        if err != nil {
            pageIndex = 0
        }

        // Attach the 'page' parameter value to our context to be used.
		ctx := context.WithValue(r.Context(), "pageIndex", pageIndex)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
