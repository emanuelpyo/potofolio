package site

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexRendersContent(t *testing.T) {
	h, err := New("http://localhost:3000")
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	body := string(h.render)

	for _, want := range []string{
		"Stiven Deryl Imanuel Pyo",
		"/css/app.css",
		"/js/app.js",
		`id="home"`,
		`id="projects"`,
		`id="contacts"`,
		"Cate Studio",
		"Wakita.id",
		"/images/profile.jpeg",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("rendered page is missing %q", want)
		}
	}
}

func TestUnknownPathRenders404(t *testing.T) {
	h, err := New("http://localhost:3000")
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/nope", nil))
	if rec.Code != 404 {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Halaman tidak ditemukan") {
		t.Fatalf("unexpected 404 body: %s", rec.Body.String())
	}
}

func TestRootRenders(t *testing.T) {
	h, err := New("http://localhost:3000")
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
