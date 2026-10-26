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
		`id="lang-switcher"`,
		`id="lang-toggle"`,
		`data-lang-choice="id"`,
		`data-lang-choice="en"`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("rendered page is missing %q", want)
		}
	}
}

func TestJSConfigContainsBilingualDict(t *testing.T) {
	h, err := New("http://localhost:3000")
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	body := string(h.render)
	if !strings.Contains(body, `"langKey":"portfolio-language"`) {
		t.Errorf("rendered page missing langKey in JSConfig")
	}
	if !strings.Contains(body, `"hero_view_work"`) {
		t.Errorf("rendered page missing hero_view_work dictionary key")
	}
}

func TestIndonesianIsDefaultInHTML(t *testing.T) {
	h, err := New("http://localhost:3000")
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	body := string(h.render)
	for _, want := range []string{
		`<html lang="id"`,
		"Halo, saya",
		"Tentang Saya",
		"Siapa Saya",
		"Pendekatan Saya",
		"Informasi Pribadi",
		"Pengalaman Kerja",
		"Keahlian &amp; Alat",
		"Teknologi",
		"Karya Pilihan",
		"Lihat Detail",
		"Hubungi Saya",
		"Kontak",
		"Unduh CV",
		"Terhubung",
		"Memuat Portofolio",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("HTML output does not have Indonesian default for %q", want)
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
