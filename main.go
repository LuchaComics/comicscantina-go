package main

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"os"
	"os/signal"
	"time"
	"github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/render"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/valve"
	"github.com/luchacomics/comicscantina-go/internal/base/config"
    "github.com/luchacomics/comicscantina-go/internal/controller"
	_ "github.com/luchacomics/comicscantina-go/internal/base/database"
	"github.com/luchacomics/comicscantina-go/internal/base/service"
	cc_mw "github.com/luchacomics/comicscantina-go/internal/base/middleware"
)

// Initialize our applications shared functions.
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // Use all CPU cores
}

// Entry point into our web service.
func main() {
	// Our graceful valve shut-off package to manage code preemption and
	// shutdown signaling.
	valv := valve.New()
	baseCtx := valv.Context()
	r := chi.NewRouter()

    //--------------------------------//
	// Load up our global middleware. //
	//--------------------------------//
    r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

    //------------------------------------------------------------------------//
    // Load up our non-protected API endpoints. The following API endpoints   //
	// can be accessed regardless of whether a JWT token was provided or not. //
	//------------------------------------------------------------------------//
	r.Get("/", controller.HealthCheckFunc)
	r.Get("/api/v1/public/version", controller.HealthCheckFunc)
	r.Post("/api/v1/public/register", controller.RegisterFunc)
    r.Post("/api/v1/public/login", controller.LoginFunc)
    r.With(cc_mw.PaginationCtx).Get("/api/v1/public/organizations", controller.ListPublicOrganizationsFunc)

    //------------------------------------------------------------------------//
	// Load up our protected API endpoints. The following API endpoints can   //
	// only be accessed with submission of a JWT token in the header.         //
	//------------------------------------------------------------------------//
	r.Group(func(r chi.Router) {
		//--------------------------------------------------------------------//
		//                             Middleware                             //
		//--------------------------------------------------------------------//
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(service.GetJWTTokenAuthority()))

		// Handle valid / invalid tokens. In the following API endpoints, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

        // This is the comics cantina authenticated user middleware which will
		// lookup the verified JWT token and attach as a context to the request.
		r.Use(cc_mw.ProfileCtx)

		//--------------------------------------------------------------------//
		//                           API endpoints                            //
		//--------------------------------------------------------------------//

		// User
		r.Get("/api/v1/profile", controller.ProfileRetrieveFunc)

		// Organizations
		r.With(cc_mw.PaginationCtx).With(cc_mw.StaffCtx).Get("/api/v1/organizations", controller.ListOrganizationsFunc)
		r.Post("/api/v1/organizations", controller.CreateOrganizationFunc)
		r.With(controller.OrganizationCtx).Get("/api/v1/organization/{organizationID}", controller.RetrieveOrganizationFunc)

		// Store
		r.With(cc_mw.PaginationCtx).Get("/api/v1/stores", controller.ListStoresFunc)
		r.Post("/api/v1/stores", controller.CreateStoreFunc)
		r.With(controller.StoreCtx).Get("/api/v1/store/{storeID}", controller.RetrieveStoreFunc)

		// Product
		r.With(cc_mw.PaginationCtx).With(controller.ProductFiltersCtx).Get("/api/v1/products", controller.ListProductsFunc)
		r.Post("/api/v1/products", controller.CreateProductFunc)
		//TODO: IMPLEMENT.

		// Receipts
		//TODO: IMPLEMENT.
	})

    //------------------------------------------------------------------------//
	//                         HTTP Running Server                            //
	//------------------------------------------------------------------------//
	// Get our server address.
    address := config.GetSettingsVariableAddress()

    // Integrate our server with our base context.
	srv := http.Server{Addr: address, Handler: chi.ServerBaseContext(baseCtx, r)}

    // The following code was taken from the following repo:
	// https://github.com/go-chi/chi/blob/0c5e7abb4e562fa14dd2548cb57b28f979a7dcd9/_examples/graceful/main.go#L88
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Println("shutting down..")

			// first valv
			valv.Shutdown(20 * time.Second)

			// create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// start http shutdown
			srv.Shutdown(ctx)

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()
	srv.ListenAndServe()

    // // Start our web-server.
	// http.ListenAndServe(":8080", r)
}
