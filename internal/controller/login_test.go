package controller

import (
    "fmt"
    "encoding/json"
    "bytes"
    "strings"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/luchacomics/comicscantina-go/internal/base/database"
    "github.com/luchacomics/comicscantina-go/internal/controller"
    "github.com/luchacomics/comicscantina-go/internal/serializer"
    "github.com/luchacomics/comicscantina-go/internal/model_resource"
)


func Test_LoginFunc(t *testing.T) {
    // Restart the database.
    dao := database.Instance()
    dao.DropAndCreateDatabase()

    // Setup our user object.
    model_resource.DBNewUser("test@test.com", "123password", "Bob", "Joe")

    // Run our test.
    jsonString := &serializer.LoginRequest{
        Email: "test@test.com",
        Password: "123password",
    }
    jsonBytesArray, _ := json.Marshal(jsonString)
    req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonBytesArray))
    req.Header.Set("Content-Type", "application/json")

    res := httptest.NewRecorder()
    controller.LoginFunc(res, req)

    // Validate our request.
    // (1) Validate the output.
    exp := "token"
    act := res.Body.String()
    act_str := string(act)
    if strings.Contains(act_str, exp) == false {
        t.Fatalf("Expected %s got %s", exp, act)
    }

    // (2) Valide the status.
    if res.Code != 200 {
        t.Fatalf("Expected %d got %d", 201, res.Code)
    }

    // (3) For debugging purposes only.
    fmt.Println("response:", res.Body.String())

}
