package main

import (
	"net/http"
	"runtime"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
    "github.com/go-chi/render"
    "github.com/luchacomics/comicscantina-go/internal/controller"
	_ "github.com/luchacomics/comicscantina-go/internal/base/database"
	"github.com/luchacomics/comicscantina-go/internal/base/service"
	cc_middleware "github.com/luchacomics/comicscantina-go/internal/base/middleware"
)

// Initialize our applications shared functions.
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // Use all CPU cores
}

// Entry point into our web service.
func main() {
	r := chi.NewRouter()

	// Load up our middleware.
    r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

    // Load up our API endpoints.
    r.Get("/", controller.HealthCheckFunc)
	r.Post("/api/v1/register", controller.RegisterFunc)
    r.Post("/api/v1/login", controller.LoginFunc)

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(service.GetJWTTokenAuthority()))

		// Handle valid / invalid tokens. In the following API endpoints, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

        // This is the comics cantina authenticated user middleware which will
		// lookup the verified JWT token and attach as a context to the request.
		r.Use(cc_middleware.ProfileCtx)

		// API endpoints.
		r.Get("/api/v1/profile", controller.ProfileRetrieveFunc)
		r.Get("/api/v1/organizations", controller.ListOrganizationsFunc)
		r.Post("/api/v1/organizations", controller.CreateOrganizationFunc)
		r.With(controller.OrganizationCtx).Get("/api/v1/organization/{organizationID}", controller.RetrieveOrganizationFunc)
	})

	http.ListenAndServe(":8080", r)
}
