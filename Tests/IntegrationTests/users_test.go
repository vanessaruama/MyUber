package integrationtests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanessaruama/myuber/routes"
)

func TestUsersRoute(t *testing.T) {
	router := routes.SetUpRouter()
	reader := strings.NewReader(`{
		"id": "01",
		"username": "Vanessa",
		"email": "vanessa@teste.com",
		"cpf": "111.111.111-1"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v2/users", reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var ResponseBody map[string]interface{}

	json.Unmarshal([]byte(w.Body.String()), &ResponseBody)
	assert.Equal(t, "Vanessa", ResponseBody["username"])
}
