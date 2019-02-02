package controller

import (
    "fmt"
    "context"
    "strings"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/base/service"
    "github.com/luchacomics/comicscantina-go/internal/controller"
    "github.com/luchacomics/comicscantina-go/internal/model_manager"
)


func Test_ProfileFunc(t *testing.T) {
    // Restart the database.
    dao := database.Instance()
    dao.DropAndCreateDatabase()

    // Create our user account and generate our JWT token.
    user, _ := model_manager.UserManagerInstance().Create("bart@mikasoftware.com", "123password", "Bart", "Mika")
    token := service.GenerateJWTToken(user.ID)
    bearer_token := "Bearer "+token

    // Create our request object.
    req, err := http.NewRequest("GET", "/api/v1/profile", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Attach our authentication.
    req.Header.Add("Authorization", bearer_token)
    ctx := service.NewContextWithJWTToken(token) // Required by `JWT` middleware.
    ctx = context.WithValue(ctx, "user", user)   // Required by `profile_middleware.go`.
    req = req.WithContext(ctx)

    // Create our response object.
    res := httptest.NewRecorder()

    // RUN OUR TEST.
    controller.ProfileRetrieveFunc(res, req)

    // Validate our request.
    exp := "bart@mikasoftware.com"
    act := res.Body.String()
    act_str := string(act)
    if strings.Contains(act_str, exp) == false {
        t.Fatalf("Expected %s got %s", exp, act)
    }

    // For debugging purposes only.
    fmt.Println("response:", res.Body.String())
}
