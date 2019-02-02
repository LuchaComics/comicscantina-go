package controller

import (
    "fmt"
    "encoding/json"
    "bytes"
    "strings"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/luchacomics/comicscantina-go/internal/controller"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
)


func Test_RegisterFunc(t *testing.T) {
    // Restart the database.
    dao := database.Instance()
    dao.DropAndCreateDatabase()

    registration := &serializer.RegistrationRequest{
        Email: "test@test.com",
        Password: "123password",
        FirstName: "Bob",
        LastName: "Joe",
    }
    jsonRegistration, _ := json.Marshal(registration)
    req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonRegistration))
    req.Header.Set("Content-Type", "application/json")

    res := httptest.NewRecorder()
    controller.RegisterFunc(res, req)

    // Validate our request.
    // (1) Validate the output.
    exp := "token"
    act := res.Body.String()
    act_str := string(act)
    if strings.Contains(act_str, exp) == false {
        t.Fatalf("Expected %s got %s", exp, act)
    }

    // (2) Valide the status.
    if res.Code != 201 {
        t.Fatalf("Expected %d got %d", 201, res.Code)
    }

    // (3) For debugging purposes only.
    fmt.Println("response:", res.Body.String())

}
