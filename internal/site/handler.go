package site

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed templates/*.html
var templateFS embed.FS

const cvPath = "/cv/StivenDerylImanuelPyo_Resume_AppleAcademy.pdf"

type Meta struct {
	Title         string
	Description   string
	OGDescription string
	Keywords      string
	BaseURL       string
}

type page struct {
	Meta            Meta
	Profile         Profile
	Socials         []Social
	QuickStats      []QuickStat
	Companies       []Company
	Experiences     []Experience
	TechCategories  []TechCategory
	Projects        []Project
	Certificates    []Certificate
	GraduationPhoto []GraduationPhoto
	ContactLinks    []ContactLink
	CVPath          string
	WelcomeTime     string
	JSConfig        template.JS
}

type jsConfig struct {
	Roles      []string `json:"roles"`
	Name       string   `json:"name"`
	ShortName  string   `json:"shortName"`
	Location   string   `json:"location"`
	Email      string   `json:"email"`
	LinkedIn   string   `json:"linkedin"`
	GitHub     string   `json:"github"`
	Education  string   `json:"education"`
	CVPath     string   `json:"cvPath"`
	Welcome    string   `json:"welcome"`
	StorageKey string   `json:"storageKey"`
}

// Handler renders the single-page portfolio.
type Handler struct {
	tmpl   *template.Template
	render []byte
}

// New parses the templates and pre-renders the page once at startup.
func New(baseURL string) (*Handler, error) {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"mul":  func(a, b int) int { return a * b },
		"add":  func(a, b int) int { return a + b },
		"sub":  func(a, b int) int { return a - b },
		"pad2": func(n int) string { return fmt.Sprintf("%02d", n) },
	}).ParseFS(templateFS, "templates/*.html")
	if err != nil {
		return nil, err
	}

	data := page{
		Meta: Meta{
			Title:         "Stiven Deryl Imanuel Pyo | Portfolio",
			Description:   "Portfolio pribadi Stiven Deryl Imanuel Pyo, Web Developer dan IT Support. Lulusan Informatika yang membangun antarmuka web dan merawat sistem di baliknya.",
			OGDescription: "Portfolio pribadi Stiven Deryl Imanuel Pyo, Web Developer dan IT Support.",
			Keywords:      "Stiven Deryl Imanuel Pyo, Deryl Pyo, Portfolio, Web Developer, IT Support, Frontend, Laravel, React",
			BaseURL:       baseURL,
		},
		Profile:         profile,
		Socials:         socials,
		QuickStats:      quickStats,
		Companies:       companies,
		Experiences:     experiences,
		TechCategories:  techCategories,
		Projects:        projects,
		Certificates:    certificates,
		GraduationPhoto: graduationPhotos,
		ContactLinks:    contactLinks,
		CVPath:          cvPath,
		WelcomeTime:     time.Unix(0, 0).Format("15:04"),
		JSConfig:        buildJSConfig(),
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "index", data); err != nil {
		return nil, err
	}
	return &Handler{tmpl: tmpl, render: buf.Bytes()}, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(notFound))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(h.render)
}

const notFound = `<!DOCTYPE html><html lang="id" class="dark"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>404 | Stiven Deryl Imanuel Pyo</title><link rel="stylesheet" href="/css/app.css"></head><body class="antialiased"><main class="min-h-screen flex flex-col items-center justify-center gap-4 px-6 text-center"><h1 class="text-6xl font-black tracking-tighter text-text-primary">404</h1><p class="text-text-secondary font-medium">Halaman tidak ditemukan.</p><a href="/" class="text-sm font-bold bg-text-primary text-background px-6 py-3 rounded-xl hover:-translate-y-1 transition-all duration-300">Kembali ke Beranda</a></main></body></html>`

func buildJSConfig() template.JS {
	c := jsConfig{
		Roles:      profile.Roles,
		Name:       profile.Name,
		ShortName:  profile.ShortName,
		Location:   profile.Location,
		Email:      profile.Email,
		LinkedIn:   profile.LinkedIn,
		GitHub:     profile.GitHub,
		Education:  profile.Education,
		CVPath:     cvPath,
		Welcome:    "Halo! Gue asisten AI " + profile.ShortName + ". Ada yang bisa gue bantu soal portofolio, pengalaman, atau project " + profile.ShortName + "?",
		StorageKey: "deryl_chat_messages",
	}
	b, err := json.Marshal(c)
	if err != nil {
		log.Printf("js config marshal: %v", err)
		return template.JS("{}")
	}
	return template.JS(b)
}
