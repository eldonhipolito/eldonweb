package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/eldonhipolito/eldonweb/models"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}

}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpArticleList = models.GetAllArticles()
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	models.SetAllArticles(tmpArticleList)
}
