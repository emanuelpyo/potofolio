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
	Name          string                       `json:"name"`
	ShortName     string                       `json:"shortName"`
	Roles         map[string][]string          `json:"roles"`
	CVPath        string                       `json:"cvPath"`
	StorageKey    string                       `json:"storageKey"`
	LangKey       string                       `json:"langKey"`
	SystemPrompts map[string]string            `json:"systemPrompts"`
	Dict          map[string]map[string]string `json:"dict"`
	Projects      []projectI18n                `json:"projects"`
	Certificates  []certI18n                   `json:"certificates"`
	Graduation    []gradI18n                   `json:"graduation"`
}

type projectI18n struct {
	Index            int    `json:"index"`
	Title            string `json:"title"`
	ShortDescription Text   `json:"shortDescription"`
	Features         []Text `json:"features"`
}

type certI18n struct {
	Index int  `json:"index"`
	Title Text `json:"title"`
}

type gradI18n struct {
	Index int  `json:"index"`
	Alt   Text `json:"alt"`
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

var uiDict = map[string]map[string]string{
	"id": {
		"nav_main":               "Navigasi utama",
		"nav_home":               "Beranda",
		"nav_companies":          "Kolaborasi",
		"nav_about":              "Tentang",
		"nav_experience":         "Pengalaman",
		"nav_projects":           "Karya",
		"nav_contacts":           "Kontak",
		"theme_to_light":         "Aktifkan tema terang",
		"theme_to_dark":          "Aktifkan tema gelap",
		"lang_switch":            "Ganti bahasa",
		"lang_current":           "Bahasa Indonesia",
		"menu_open":              "Buka menu navigasi",
		"menu_close":             "Tutup menu navigasi",
		"loader_title":           "Memuat Portofolio",
		"hero_greeting":          "Halo, saya",
		"hero_view_work":         "Lihat Karya",
		"hero_download_cv":       "Unduh CV",
		"hero_cv_unavailable":    "CV belum tersedia",
		"hero_connect":           "Terhubung",
		"marquee_greeting":       "Halo, Aku Deryl",
		"marquee_role":           "Web Developer",
		"collab_title":           "KOLABORASI",
		"collab_visit":           "Kunjungi Website",
		"collab_active_intern":   "Magang Aktif",
		"about_title":            "Tentang Saya",
		"about_who":              "Siapa Saya",
		"about_approach":         "Pendekatan Saya",
		"about_details":          "Informasi Pribadi",
		"grad_eyebrow":           "Momen",
		"grad_title":             "Dokumentasi Wisuda",
		"grad_enlarge":           "Perbesar",
		"grad_view_all":          "Lihat Semua Foto",
		"grad_show_less":         "Tampilkan Lebih Sedikit",
		"grad_dialog_label":      "Pratinjau foto wisuda",
		"grad_close_preview":     "Tutup pratinjau foto",
		"grad_prev_photo":        "Foto sebelumnya",
		"grad_next_photo":        "Foto berikutnya",
		"exp_title":              "Pengalaman Kerja",
		"exp_active_intern":      "Magang aktif",
		"exp_active":             "Aktif",
		"tech_eyebrow":           "Keahlian & Alat",
		"tech_title":             "Teknologi",
		"proj_eyebrow":           "Portofolio",
		"proj_title":             "Karya Pilihan",
		"proj_empty_title":       "Karya sedang disiapkan",
		"proj_empty_desc_pre":    "Foto dan detail proyek akan ditambahkan. Data proyek tersimpan di",
		"proj_empty_desc_post":   "dan tinggal diisi.",
		"proj_no_photo":          "Foto belum ada",
		"proj_view_details":      "Lihat Detail",
		"proj_open_demo_prefix":  "Buka demo",
		"proj_close_details":     "Tutup detail proyek",
		"proj_modal_created":     "Dibuat",
		"proj_modal_tech":        "Teknologi",
		"proj_modal_features":    "Fitur Utama",
		"proj_modal_preview":     "Tampilan",
		"proj_view_preview":      "Lihat tampilan",
		"proj_live_demo":         "Live Demo",
		"proj_demo_unavailable":  "Live Demo Belum Ada",
		"proj_source_code":       "Kode Sumber",
		"cert_eyebrow":           "Sertifikasi",
		"cert_title":             "Sertifikat",
		"cert_view":              "Lihat",
		"cert_close":             "Tutup sertifikat",
		"contact_eyebrow":        "Hubungi Saya",
		"contact_title":          "Kontak",
		"contact_card_heading":   "Ada proyek atau peluang yang ingin dibicarakan?",
		"contact_card_desc":      "Terbuka untuk kolaborasi, konsultasi, maupun tawaran penuh waktu. Kirim pesan lewat salah satu kanal di samping, atau ngobrol langsung dengan asisten AI di pojok kanan bawah.",
		"contact_ask_ai":         "Tanya Asisten AI",
		"chat_open_label":        "Buka asisten AI Deryl",
		"chat_dialog_label":      "Asisten AI",
		"chat_title":             "Asisten Deryl",
		"chat_status_online":     "Online",
		"chat_clear_title":       "Mulai baru",
		"chat_clear_label":       "Mulai percakapan baru",
		"chat_close_title":       "Tutup",
		"chat_close_label":       "Tutup asisten",
		"chat_unconfigured_warn": "Asisten AI belum dikonfigurasi. Riwayat percakapan tetap tersimpan di perangkat ini.",
		"chat_welcome_text":      "Halo! Gue asisten AI Deryl. Ada yang bisa gue bantu soal portofolio, pengalaman, atau project Deryl?",
		"chat_input_placeholder": "Tanya apa saja soal Deryl...",
		"chat_input_label":       "Tulis pesan",
		"chat_send_label":        "Kirim pesan",
		"chat_unconfigured_err":  "Asisten AI belum dikonfigurasi pada server ini, jadi gue belum bisa menjawab otomatis. Silakan hubungi Deryl lewat email atau LinkedIn di bagian kontak.",
		"chat_conn_err_prefix":   "Maaf, terjadi kesalahan koneksi atau server AI.",
		"footer_rights":          "Hak cipta dilindungi.",
	},
	"en": {
		"nav_main":               "Main navigation",
		"nav_home":               "Home",
		"nav_companies":          "Companies",
		"nav_about":              "About",
		"nav_experience":         "Experience",
		"nav_projects":           "Projects",
		"nav_contacts":           "Contacts",
		"theme_to_light":         "Enable light theme",
		"theme_to_dark":          "Enable dark theme",
		"lang_switch":            "Switch language",
		"lang_current":           "English",
		"menu_open":              "Open navigation menu",
		"menu_close":             "Close navigation menu",
		"loader_title":           "Portfolio Loading",
		"hero_greeting":          "Hi, I'm",
		"hero_view_work":         "View Work",
		"hero_download_cv":       "Download CV",
		"hero_cv_unavailable":    "CV not available yet",
		"hero_connect":           "Connect",
		"marquee_greeting":       "Hello, I'm Deryl",
		"marquee_role":           "Web Developer",
		"collab_title":           "COLLABORATION",
		"collab_visit":           "Visit Website",
		"collab_active_intern":   "Active Internship",
		"about_title":            "About Me",
		"about_who":              "Who Am I",
		"about_approach":         "My Approach",
		"about_details":          "Personal Details",
		"grad_eyebrow":           "Moments",
		"grad_title":             "Graduation Documentation",
		"grad_enlarge":           "Enlarge",
		"grad_view_all":          "View All Photos",
		"grad_show_less":         "Show Less",
		"grad_dialog_label":      "Graduation photo preview",
		"grad_close_preview":     "Close photo preview",
		"grad_prev_photo":        "Previous photo",
		"grad_next_photo":        "Next photo",
		"exp_title":              "Work Experience",
		"exp_active_intern":      "Active internship",
		"exp_active":             "Active",
		"tech_eyebrow":           "Skills & Tools",
		"tech_title":             "Tech Stack",
		"proj_eyebrow":           "Portfolio",
		"proj_title":             "Selected Works",
		"proj_empty_title":       "Projects coming soon",
		"proj_empty_desc_pre":    "Project photos and details will be added soon. Project data lives in",
		"proj_empty_desc_post":   "and is ready to be filled.",
		"proj_no_photo":          "No photo yet",
		"proj_view_details":      "View Details",
		"proj_open_demo_prefix":  "Open demo",
		"proj_close_details":     "Close project details",
		"proj_modal_created":     "Created",
		"proj_modal_tech":        "Technologies",
		"proj_modal_features":    "Key Features",
		"proj_modal_preview":     "Preview",
		"proj_view_preview":      "View preview",
		"proj_live_demo":         "Live Demo",
		"proj_demo_unavailable":  "Live Demo Unavailable",
		"proj_source_code":       "Source Code",
		"cert_eyebrow":           "Certificates",
		"cert_title":             "Certificates",
		"cert_view":              "View",
		"cert_close":             "Close certificate",
		"contact_eyebrow":        "Get In Touch",
		"contact_title":          "Contact Me",
		"contact_card_heading":   "Have a project or opportunity to discuss?",
		"contact_card_desc":      "Open for collaborations, consulting, and full-time opportunities. Reach out via any channel on the right, or chat directly with the AI assistant at the bottom right.",
		"contact_ask_ai":         "Ask AI Assistant",
		"chat_open_label":        "Open Deryl's AI assistant",
		"chat_dialog_label":      "AI Assistant",
		"chat_title":             "Deryl Assistant",
		"chat_status_online":     "Online",
		"chat_clear_title":       "Start over",
		"chat_clear_label":       "Start new conversation",
		"chat_close_title":       "Close",
		"chat_close_label":       "Close assistant",
		"chat_unconfigured_warn": "AI assistant is not configured yet. Chat history is still stored on this device.",
		"chat_welcome_text":      "Hello! I am Deryl's AI assistant. How can I help you with Deryl's portfolio, experience, or projects?",
		"chat_input_placeholder": "Ask anything about Deryl...",
		"chat_input_label":       "Write a message",
		"chat_send_label":        "Send message",
		"chat_unconfigured_err":  "The AI assistant has not been configured on this server, so I cannot reply automatically yet. Please contact Deryl via email or LinkedIn in the contact section.",
		"chat_conn_err_prefix":   "Sorry, a connection or AI server error occurred.",
		"footer_rights":          "All rights reserved.",
	},
}

func buildJSConfig() template.JS {
	var projs []projectI18n
	for _, p := range projects {
		projs = append(projs, projectI18n{
			Index:            p.Index,
			Title:            p.Title,
			ShortDescription: p.ShortDescription,
			Features:         p.Features,
		})
	}

	var certs []certI18n
	for i, c := range certificates {
		certs = append(certs, certI18n{
			Index: i,
			Title: c.Title,
		})
	}

	var grads []gradI18n
	for i, g := range graduationPhotos {
		grads = append(grads, gradI18n{
			Index: i,
			Alt:   g.Alt,
		})
	}

	c := jsConfig{
		Name:          profile.Name,
		ShortName:     profile.ShortName,
		Roles:         profile.Roles,
		CVPath:        cvPath,
		StorageKey:    "deryl_chat_messages",
		LangKey:       "portfolio-language",
		SystemPrompts: systemPrompts,
		Dict:          uiDict,
		Projects:      projs,
		Certificates:  certs,
		Graduation:    grads,
	}

	b, err := json.Marshal(c)
	if err != nil {
		log.Printf("js config marshal: %v", err)
		return template.JS("{}")
	}
	return template.JS(b)
}
