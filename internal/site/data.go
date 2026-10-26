package site

import (
	"html/template"
)

// Text holds bilingual content (Indonesian and English).
type Text struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

type Profile struct {
	Name            string
	ShortName       string
	Roles           map[string][]string
	Tagline         Text
	Location        Text
	Email           string
	GitHub          string
	LinkedIn        string
	Education       Text
	About           Text
	Approach        Text
	PersonalDetails []Detail
}

type Detail struct {
	Key   string
	Label Text
	Value Text
	Href  string
}

type Social struct {
	Href  string
	Label string
	Icon  template.HTML
}

type QuickStat struct {
	Key     string
	Message Text
	Icon    template.HTML
}

type Company struct {
	Name        string
	Role        Text
	Logo        string
	Description Text
	URL         string
}

type Experience struct {
	ID          int
	Company     string
	Role        Text
	Date        Text
	Description Text
	Skills      []string
}

type Tech struct {
	Name string
	SVG  string
}

type TechCategory struct {
	Key          string
	Title        Text
	Description  Text
	Technologies []Tech
}

type Project struct {
	Index            int
	ImagePath        string
	Title            string
	ShortDescription Text
	CreatedAt        string
	Features         []Text
	Gallery          []string
	GalleryFit       string
	Tech             []string
	GithubURL        string
	LiveDemoURL      string
	IsPrivateRepo    bool
}

type Certificate struct {
	Issuer string
	Title  Text
	Type   string
	Src    string
}

type GraduationPhoto struct {
	Src string
	Alt Text
}

type ContactLink struct {
	Label Text
	Value string
	Href  string
	Icon  template.HTML
}

var profile = Profile{
	Name:      "Stiven Deryl Imanuel Pyo",
	ShortName: "Deryl",
	Roles: map[string][]string{
		"id": {"Pengembang Web", "Dukungan IT", "Frontend Developer"},
		"en": {"Web Developer", "IT Support", "Frontend Developer"},
	},
	Tagline: Text{
		ID: "Halo! saya suka bereksperimen dengan ide-ide baru. Bagi saya, setiap baris kode adalah cara untuk menghadirkan sesuatu yang bermanfaat dan keren di dunia digital.",
		EN: "Hello! I love experimenting with new ideas. To me, every line of code is an opportunity to bring something useful and impactful to the digital world.",
	},
	Location: Text{
		ID: "Indonesia",
		EN: "Indonesia",
	},
	Email:    "deryl@lawkey.id",
	GitHub:   "https://github.com/emanuelpyo",
	LinkedIn: "https://www.linkedin.com/in/derylpyo/",
	Education: Text{
		ID: "Universitas Sanata Dharma Yogyakarta",
		EN: "Sanata Dharma University Yogyakarta",
	},
	About: Text{
		ID: "Setelah menyelesaikan kuliah, saya mulai lebih serius menikmati proses coding dan mengembangkan karier di bidang web development. Saat ini, saya bergabung dengan tim Lawkey sebagai Web Developer dan sedang menjalani magang sebagai Web Developer di PT Primadaya Plastisindo Tbk.",
		EN: "After finishing university, I pursued web development with a strong focus on craftsmanship and growth. Currently, I work as a Web Developer at Lawkey and serve as a Web Developer Intern at PT Primadaya Plastisindo Tbk.",
	},
	Approach: Text{
		ID: "Saya mengembangkan antarmuka web yang rapi dan mudah digunakan, sekaligus menjaga sistem internal tetap berjalan. Web development mengasah kemampuan UI/UX, sementara IT Support melatih pemecahan masalah teknis.",
		EN: "I craft clean, user-friendly web interfaces while keeping internal systems reliable. Web development sharpens my UI/UX skills, while IT Support grounds my technical problem-solving.",
	},
	PersonalDetails: []Detail{
		{
			Key:   "name",
			Label: Text{ID: "Nama", EN: "Name"},
			Value: Text{ID: "Stiven Deryl Imanuel Pyo", EN: "Stiven Deryl Imanuel Pyo"},
		},
		{
			Key:   "location",
			Label: Text{ID: "Lokasi", EN: "Location"},
			Value: Text{ID: "Indonesia", EN: "Indonesia"},
		},
		{
			Key:   "email",
			Label: Text{ID: "Email", EN: "Email"},
			Value: Text{ID: "deryl@lawkey.id", EN: "deryl@lawkey.id"},
			Href:  "mailto:deryl@lawkey.id",
		},
		{
			Key:   "focus",
			Label: Text{ID: "Fokus", EN: "Focus"},
			Value: Text{ID: "Pengembangan Web", EN: "Web Development"},
		},
		{
			Key:   "interests",
			Label: Text{ID: "Minat", EN: "Interests"},
			Value: Text{ID: "Machine Learning", EN: "Machine Learning"},
		},
		{
			Key:   "education",
			Label: Text{ID: "Pendidikan", EN: "Education"},
			Value: Text{ID: "Universitas Sanata Dharma Yogyakarta", EN: "Sanata Dharma University Yogyakarta"},
		},
	},
}

var socials = []Social{
	{
		Href:  profile.GitHub,
		Label: "GitHub",
		Icon:  `<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" /></svg>`,
	},
	{
		Href:  profile.LinkedIn,
		Label: "LinkedIn",
		Icon:  `<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true"><path d="M4.98 3.5c0 1.381-1.11 2.5-2.48 2.5s-2.48-1.119-2.48-2.5c0-1.38 1.11-2.5 2.48-2.5s2.48 1.12 2.48 2.5zm.02 4.5h-5v16h5v-16zm7.982 0h-4.968v16h4.969v-8.399c0-4.67 6.029-5.052 6.029 0v8.399h4.988v-10.131c0-7.88-8.922-7.593-11.018-3.714v-2.155z" /></svg>`,
	},
	{
		Href:  "mailto:" + profile.Email,
		Label: "Email",
		Icon:  `<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true"><path d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" fill="none" /></svg>`,
	},
}

var quickStats = []QuickStat{
	{
		Key: "role",
		Message: Text{
			ID: "Pengembang Web",
			EN: "Web Developer",
		},
		Icon: `<svg class="w-5 md:w-6" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24"><path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="m8 8-4 4 4 4m8 0 4-4-4-4m-2-3-4 14" /></svg>`,
	},
	{
		Key: "lawkey",
		Message: Text{
			ID: "Tim Lawkey",
			EN: "Lawkey Team",
		},
		Icon: `<svg class="w-5 md:w-6" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24"><path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 12V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5M5 12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2M5 12h14m-7 4v3m-4 0h8" /></svg>`,
	},
	{
		Key: "intern",
		Message: Text{
			ID: "Magang Pengembang Web",
			EN: "Web Developer Intern",
		},
		Icon: `<svg class="w-5 md:w-6" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24"><path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 8v4l3 3m6-3a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" /></svg>`,
	},
}

var companies = []Company{
	{
		Name: "Lawkey",
		Role: Text{
			ID: "Pengembang Web",
			EN: "Web Developer",
		},
		Logo: "/images/logo/lawkey-dashboardweb.jpeg",
		Description: Text{
			ID: "Bergabung dengan tim Lawkey sebagai Web Developer.",
			EN: "Joined the Lawkey team as a Web Developer.",
		},
		URL: "https://www.lawkey.id/#top",
	},
	{
		Name: "PT Primadaya Plastisindo Tbk",
		Role: Text{
			ID: "Magang Pengembang Web",
			EN: "Web Developer Intern",
		},
		Logo: "/images/logo/pdpsolo.jpeg",
		Description: Text{
			ID: "Menjalani magang di Departemen IT Support & Web Developer.",
			EN: "Completing an internship in the IT Support & Web Developer department.",
		},
		URL: "",
	},
}

var experiences = []Experience{
	{
		ID:      1,
		Company: "Lawkey",
		Role: Text{
			ID: "Pengembang Web",
			EN: "Web Developer",
		},
		Date: Text{
			ID: "Oktober 2026",
			EN: "October 2026",
		},
		Description: Text{
			ID: "Bergabung dengan tim Lawkey sebagai Web Developer, mengerjakan pengembangan antarmuka dan fitur web.",
			EN: "Joined the Lawkey team as a Web Developer, building web interfaces and user-facing features.",
		},
		Skills: []string{"Pengembangan Web"},
	},
	{
		ID:      2,
		Company: "PT Primadaya Plastisindo Tbk",
		Role: Text{
			ID: "Magang Pengembang Web",
			EN: "Web Developer Intern",
		},
		Date: Text{
			ID: "September 2026",
			EN: "September 2026",
		},
		Description: Text{
			ID: "Sedang menjalani magang sebagai Web Developer, membantu pengembangan sistem web internal perusahaan.",
			EN: "Serving as a Web Developer Intern, supporting the development of internal company web systems.",
		},
		Skills: []string{"Pengembangan Web"},
	},
}

var techCategories = []TechCategory{
	{
		Key: "frontend",
		Title: Text{
			ID: "Frontend",
			EN: "Frontend",
		},
		Description: Text{
			ID: "Bahasa dan pustaka untuk membangun antarmuka web yang interaktif.",
			EN: "Languages and libraries for building interactive web interfaces.",
		},
		Technologies: []Tech{
			{Name: "HTML5", SVG: "/icons/html.svg"},
			{Name: "CSS3", SVG: "/icons/css.svg"},
			{Name: "JavaScript"},
			{Name: "TypeScript", SVG: "/icons/typescript.svg"},
			{Name: "React.js", SVG: "/icons/react.svg"},
			{Name: "Tailwind CSS", SVG: "/icons/tailwindcss.svg"},
		},
	},
	{
		Key: "backend",
		Title: Text{
			ID: "Backend",
			EN: "Backend",
		},
		Description: Text{
			ID: "Teknologi sisi server dan framework.",
			EN: "Server-side technologies and frameworks.",
		},
		Technologies: []Tech{
			{Name: "PHP"},
			{Name: "Laravel", SVG: "/icons/laravel.svg"},
		},
	},
	{
		Key: "database",
		Title: Text{
			ID: "Basis Data",
			EN: "Databases & Data",
		},
		Description: Text{
			ID: "Basis data dan perangkat untuk pengolahan data.",
			EN: "Databases and data processing tools.",
		},
		Technologies: []Tech{
			{Name: "MySQL", SVG: "/icons/mysql.svg"},
			{Name: "Python"},
		},
	},
	{
		Key: "tools",
		Title: Text{
			ID: "Alat & Infrastruktur",
			EN: "Tools & Infrastructure",
		},
		Description: Text{
			ID: "Version control, editor, dan deployment.",
			EN: "Version control, tools, and deployment environments.",
		},
		Technologies: []Tech{
			{Name: "Git", SVG: "/icons/git.svg"},
			{Name: "GitHub", SVG: "/icons/github.svg"},
			{Name: "Docker", SVG: "/icons/docker.svg"},
		},
	},
}

var projects = []Project{
	{
		Index:     0,
		ImagePath: "/projects/cate-studio/1.png",
		Title:     "Cate Studio",
		ShortDescription: Text{
			ID: "Platform e-commerce untuk brand fashion yang ingin toko online eksklusif. Menyediakan katalog produk, filter, keranjang belanja, integrasi pembayaran, dan nota digital.",
			EN: "An e-commerce platform for fashion brands looking for an exclusive online storefront, featuring product catalogs, filters, shopping cart, payment integration, and digital receipts.",
		},
		Features: []Text{},
		Gallery: []string{
			"/projects/cate-studio/1.png",
			"/projects/cate-studio/2.png",
			"/projects/cate-studio/3.png",
			"/projects/cate-studio/4.png",
			"/projects/cate-studio/5.png",
			"/projects/cate-studio/6.png",
			"/projects/cate-studio/7.png",
		},
		Tech: []string{"HTML", "CSS", "JavaScript"},
	},
	{
		Index:     1,
		ImagePath: "/projects/cate-studio/wakita/w1.png",
		Title:     "Wakita.id",
		ShortDescription: Text{
			ID: "Platform komunikasi bisnis berbasis WhatsApp untuk membantu perusahaan mengelola interaksi pelanggan melalui chatbot, broadcast, otomatisasi pesan, dan integrasi sistem. Solusinya mencakup chatbot, broadcast pesan, notifikasi otomatis, pengingat, serta integrasi dengan sistem bisnis atau aplikasi milik klien, sehingga proses komunikasi lebih efisien, respons ke pelanggan lebih cepat, dan customer service lebih terstruktur.",
			EN: "A WhatsApp-based business communication platform helping companies manage customer interactions through chatbots, broadcast messaging, automated notifications, reminders, and client system integrations to streamline communication and structure customer service.",
		},
		Features: []Text{
			{ID: "Chatbot & sistem bot WhatsApp", EN: "WhatsApp chatbot & bot automation"},
			{ID: "Broadcast pesan ke pelanggan", EN: "Customer broadcast messaging"},
			{ID: "Otomatisasi notifikasi & pengingat", EN: "Automated notifications & reminders"},
			{ID: "Integrasi WhatsApp dengan sistem klien", EN: "WhatsApp integration with client systems"},
			{ID: "Dukungan customer service & pengelolaan percakapan", EN: "Customer service support & chat management"},
		},
		Gallery: []string{
			"/projects/cate-studio/wakita/login-wakita.png",
			"/projects/cate-studio/wakita/early-acc-wakita.png",
			"/projects/cate-studio/wakita/w1.png",
			"/projects/cate-studio/wakita/w2.png",
		},
		GalleryFit:  "contain",
		Tech:        []string{},
		LiveDemoURL: "https://wakita.id/",
	},
}

var certificates = []Certificate{
	{
		Issuer: "IBM",
		Title: Text{
			ID: "Completion Certificate, SkillsBuild",
			EN: "Completion Certificate, SkillsBuild",
		},
		Type: "pdf",
		Src:  "/images/sertif/skillsbuild-completion.pdf",
	},
	{
		Issuer: "IBM",
		Title: Text{
			ID: "Natural Language Processing, SkillsBuild",
			EN: "Natural Language Processing, SkillsBuild",
		},
		Type: "image",
		Src:  "/images/sertif/skillsbuild-nlp.jpg",
	},
}

var graduationPhotos = []GraduationPhoto{
	{Src: "/images/wisuda/wisuda1.jpeg", Alt: Text{ID: "Dokumentasi wisuda 1", EN: "Graduation documentation 1"}},
	{Src: "/images/wisuda/wisuda2.jpeg", Alt: Text{ID: "Dokumentasi wisuda 2", EN: "Graduation documentation 2"}},
	{Src: "/images/wisuda/wisuda3.jpeg", Alt: Text{ID: "Dokumentasi wisuda 3", EN: "Graduation documentation 3"}},
	{Src: "/images/wisuda/wisuda4.jpg", Alt: Text{ID: "Dokumentasi wisuda 4", EN: "Graduation documentation 4"}},
	{Src: "/images/wisuda/wisuda5.jpeg", Alt: Text{ID: "Dokumentasi wisuda 5", EN: "Graduation documentation 5"}},
}

var contactLinks = []ContactLink{
	{
		Label: Text{ID: "Email", EN: "Email"},
		Value: profile.Email,
		Href:  "mailto:" + profile.Email,
		Icon:  `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><rect width="20" height="16" x="2" y="4" rx="2" /><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" /></svg>`,
	},
	{
		Label: Text{ID: "GitHub", EN: "GitHub"},
		Value: "emanuelpyo",
		Href:  profile.GitHub,
		Icon:  `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" /></svg>`,
	},
	{
		Label: Text{ID: "LinkedIn", EN: "LinkedIn"},
		Value: "derylpyo",
		Href:  profile.LinkedIn,
		Icon:  `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true"><path d="M19 0h-14c-2.761 0-5 2.239-5 5v14c0 2.761 2.239 5 5 5h14c2.762 0 5-2.239 5-5v-14c0-2.761-2.238-5-5-5zm-11 19h-3v-11h3v11zm-1.5-12.268c-.966 0-1.75-.79-1.75-1.764s.784-1.764 1.75-1.764 1.75.79 1.75 1.764-.783 1.764-1.75 1.764zm13.5 12.268h-3v-5.604c0-3.368-4-3.113-4 0v5.604h-3v-11h3v1.765c1.396-2.586 7-2.777 7 2.476v6.759z" /></svg>`,
	},
}

// SystemPrompts defines the AI assistant prompt in both languages.
var systemPrompts = map[string]string{
	"id": `System Context: Kamu adalah Deryl Assistant, asisten AI pribadi untuk Stiven Deryl Imanuel Pyo (Deryl). Jawab pertanyaan pengunjung website portofolio dengan ramah, profesional, dan ringkas dalam bahasa Indonesia.

Informasi tentang Deryl:
1. Profil:
   - Nama: Stiven Deryl Imanuel Pyo
   - Peran: Web Developer dan IT Support, lulusan Informatika
   - Lokasi: Indonesia
   - Email: deryl@lawkey.id
   - LinkedIn: https://www.linkedin.com/in/derylpyo/
   - GitHub: https://github.com/emanuelpyo
   - Pendidikan: Universitas Sanata Dharma Yogyakarta
2. Pengalaman:
   - Web Developer di Lawkey (bergabung dalam tim)
   - Web Developer Intern di PT Primadaya Plastisindo Tbk (sedang menjalani magang)
3. Keahlian: HTML, CSS, JavaScript, TypeScript, React, Tailwind, PHP, Laravel, Node.js, MySQL, Python, TensorFlow.

Aturan: Jawab langsung ke intinya, jangan menambahkan informasi yang tidak ada di profil ini, dan selalu bersikap ramah.`,

	"en": `System Context: You are Deryl Assistant, the personal AI assistant for Stiven Deryl Imanuel Pyo (Deryl). Answer questions from portfolio website visitors in a friendly, professional, and concise manner in English.

Information about Deryl:
1. Profile:
   - Name: Stiven Deryl Imanuel Pyo
   - Role: Web Developer and IT Support, Computer Science graduate
   - Location: Indonesia
   - Email: deryl@lawkey.id
   - LinkedIn: https://www.linkedin.com/in/derylpyo/
   - GitHub: https://github.com/emanuelpyo
   - Education: Sanata Dharma University Yogyakarta
2. Experience:
   - Web Developer at Lawkey (team member)
   - Web Developer Intern at PT Primadaya Plastisindo Tbk (active internship)
3. Skills: HTML, CSS, JavaScript, TypeScript, React, Tailwind, PHP, Laravel, Node.js, MySQL, Python, TensorFlow.

Rules: Be direct and concise, do not invent facts outside this profile, and always remain welcoming and professional.`,
}
