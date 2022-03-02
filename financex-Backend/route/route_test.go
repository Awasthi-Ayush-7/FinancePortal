package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	dbop "kivancaydogmus.com/apps/userApp/dbOp"
)

// Test - 1
func Test_getUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	expectedLat := `[{"PersonID":51,"UserName":"person1","Password":"I'll do it myself","Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjkyNDM0ODMsInVzZXJfaWQiOjAsInVzZXJfbmFtZSI6InBlcnNvbjEifQ.E-3Mkm7WulyKdJ40OiwFsFttabnTTRdobGVygaj0rm0"},{"PersonID":52,"UserName":"sample user name","Password":"sample password","Token":""}]`

	if rr.Body.String() != expectedLat {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expectedLat)
	}
}

// Test - 2
func Test_addUser(t *testing.T) {
	samplePerson := dbop.Person{Id: 1, UserName: "username", Password: "password"}
	bytePerson, _ := json.Marshal(samplePerson)

	req, err := http.NewRequest("POST", "/signUp", bytes.NewReader(bytePerson))

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	//rr.Body = bytes.NewBuffer(bytePerson)
	handler := http.HandlerFunc(addUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	req.Header.Set("Token", samplePerson.Token)

}

// Test - 3
func Test_login(t *testing.T) {
	samplePerson := dbop.Person{Id: 1, UserName: "username", Password: "password"}
	bytePerson, _ := json.Marshal(samplePerson)
	req, err := http.NewRequest("POST", "/signIn", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}
	req.Header.Set("Token", samplePerson.Token)
}

// Test - 4
func Test_deleteMe(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if myMap, ok := r.Context().Value("props").(jwt.MapClaims); !ok {
			t.Errorf("props not in request %q", myMap)
		} else {
			username := myMap["user_name"]
			if v, e := username.(string); !e {
				t.Errorf("this map does not have user_name %q", v)
			} else {
				if id := dbop.DeleteMe(v); id == 0 {
					t.Errorf("an error occured during the delete the user %q", v)
				}
			}
		}
	})
	handler := middleware.MiddleWare(testHandler)
	handler.ServeHTTP(rr, req)
}
