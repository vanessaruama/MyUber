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

func TestDriversRoute(t *testing.T) {
	router := routes.SetUpRouter()
	reader := strings.NewReader(`{
		"id": "01",
		"username": "Vanessa",
		"car": "gol",
		"cpf": "111.111.111-1"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v2/drivers", reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var ResponseBody map[string]interface{}

	json.Unmarshal([]byte(w.Body.String()), &ResponseBody)
	assert.Equal(t, "01", ResponseBody["id"])
}
