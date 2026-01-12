package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListEnrollments(t *testing.T) {
	rec := httptest.NewRecorder()
	ListEnrollments(rec, httptest.NewRequest("GET", "/api/enrollments", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
