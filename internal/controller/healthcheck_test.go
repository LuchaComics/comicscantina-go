package controller

import (
    "strings"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/luchacomics/comicscantina-go/internal/controller"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
)


func Test_HealthCheckFunc(t *testing.T) {
    // Restart the database.
    dao := database.Instance()
    dao.DropAndCreateDatabase()

    // Run our unit test.
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    res := httptest.NewRecorder()
    controller.HealthCheckFunc(res, req)

    // Validate our request.
    exp := "ComicsCantina"
    act := res.Body.String()
    act_str := string(act)
    if strings.Contains(act_str, exp) == false {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}
