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
	// r.Get("/api/v1/public-organizations", controller.ListPublicOrganizationsFunc)
	// r.With(controller.PublicOrganizationCtx).Get("/api/v1/public-organization/{organizationID}", controller.GetPublicOrganization)

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(service.GetJWTTokenAuthority()))

		// Handle valid / invalid tokens. In the following API endpoints, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		// API endpoints.
		r.Get("/api/v1/profile", controller.ProfileRetrieveFunc)
		r.Get("/api/v1/organizations", controller.ListOrganizationsFunc)
		r.Post("/api/v1/organizations", controller.CreateOrganizationFunc)

		// TODO:
		// | /organizations      |
		// | /organization/<ID>/ |
		// | /store              |
		// | /store/<ID>         |
		// | /products           |
		// | /product/<ID>       |
	})

	http.ListenAndServe(":8080", r)
}
