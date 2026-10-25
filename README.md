# Deryl Portfolio

Personal portfolio of **Stiven Deryl Imanuel Pyo**, Web Developer and IT Support.
Originally built with Next.js + TypeScript + Tailwind CSS v4 + `motion`, now
migrated to **Go** (`net/http` + `html/template`) with vanilla JavaScript for the
browser interactions. The visual design, content, layout, animations, routes, and
behavior are preserved.

## Tech stack

- **Server**: Go standard library (`net/http`, `html/template`), no Go dependencies
- **Styling**: Tailwind CSS v4 (CSS build only, via the Tailwind CLI)
- **Client**: vanilla JavaScript (`public/js/app.js`)
- **Chat**: streaming `POST /api/chat` proxied to an OpenAI-compatible endpoint, key read server-side only

## Requirements

- Go 1.25+ (developed against 1.27)
- Node.js only to compile Tailwind CSS (optional at deploy time, since the
  compiled stylesheet is committed to `public/css/app.css`)

## Run locally

```bash
# 1. Compile the stylesheet (only needed when templates/classes change)
npm run build:css

# 2. Run the server
npm run dev            # go run ./cmd/server
# or a production binary
npm run build && npm run start
```

Open http://localhost:3000.

## Project structure

```
cmd/server/main.go        wiring: routes, static files, server
internal/config/          environment + a tiny stdlib .env loader
internal/site/            content structs, template render, index handler
  templates/              layout.html + section partials (html/template)
internal/chat/            streaming chat proxy handler
styles/globals.css        Tailwind source (design tokens, light + dark themes)
public/                   static assets served at /
  css/app.css             compiled Tailwind output
  js/app.js               client interactions
  fonts/                  self-hosted Poppins (400/600/700/900)
docs/PRD.md               the original product spec and decision log
legacy/                   the pre-Next static site, kept for reference
```

## Routes

| Route | Method | Purpose |
|---|---|---|
| `/` | GET | the portfolio (all sections in one page) |
| `/api/chat` | POST | streaming chat proxy (`{text, history}` -> plain-text stream) |
| `/css`, `/js`, `/fonts`, `/images`, `/icons`, `/projects`, `/cv` | GET | static assets |
| any other path | GET | styled 404 page |

## Section order

Hero, Companies, About, Graduation (wisuda), Experience, Tech Stack, Projects,
Certificates, Contact.

---

## How to fill in content

Content lives in Go, separated from presentation, in `internal/site/data.go`.

### 1. Add a CV (Download CV button)

Drop your PDF at:

```
public/cv/StivenDerylImanuelPyo_Resume_AppleAcademy.pdf
```

The hero checks for that exact file with a HEAD request. While it is missing the
button shows an inactive "Segera" state and never errors. To change the path,
update the `cvPath` constant in `internal/site/handler.go`.

### 2. Add projects

Open `internal/site/data.go` and add objects to the `projects` slice. Put
screenshots in `public/projects/` and reference them as `/projects/<file>`.

The UI is empty-safe:

- No projects: the section shows an honest "Karya sedang disiapkan" state.
- Missing `ImagePath`: a neutral placeholder block, never a broken image.
- Missing `LiveDemoURL` / `GithubURL`: disabled buttons, never dead links.

### 3. Edit experience, skills, companies

- `internal/site/data.go` — `experiences`, `techCategories`, `companies`,
  `profile`, `graduationPhotos`, `certificates`, `contactLinks`, `socials`.
- Tech icons live in `public/icons/`.

After editing templates or Go files that contain class names, re-run
`npm run build:css` so Tailwind picks up any new utilities.

---

## Chat configuration

The chat UI (launcher + full-screen overlay) is fully functional: it opens,
sends, streams, shows a typing indicator, persists history in `localStorage`,
renders markdown, and has an honest error state.

Copy `.env.example` to `.env` and set:

| Variable | Required | Purpose |
|---|---|---|
| `AI_API_KEY` | Yes | Server-side key for an OpenAI-compatible endpoint |
| `AI_BASE_URL` | No | Defaults to `https://integrate.api.nvidia.com/v1` |
| `AI_MODEL` | No | Defaults to `meta/llama-3.1-8b-instruct` |
| `NEXT_PUBLIC_BASE_URL` | No | Canonical/OG base URL |
| `PORT` | No | HTTP port, defaults to `3000` |
| `STATIC_DIR` | No | Static asset directory, defaults to `public` |

The key is read only in `internal/chat/chat.go` on the server and is never sent
to the browser. Without `AI_API_KEY`, the route returns a clear
"Konfigurasi API AI tidak ditemukan." message and the UI tells the visitor the
assistant is not configured instead of faking a reply.

Restart the server after editing `.env`.

---

## Theming

Both light and dark themes ship and both work. Dark is the default for a
developer portfolio; a toggle in the navbar switches to light and the choice is
remembered in `localStorage`. Tokens live in `styles/globals.css`.

## Accessibility

- Semantic landmarks and heading order.
- Keyboard operable: nav, theme toggle, mobile menu, gallery, lightbox, project
  modal, certificate modal, and chat (all modals close with `Escape`).
- Visible focus indicators (never removed).
- `prefers-reduced-motion` collapses animation to the resting state.
- Text contrast meets WCAG AA in both themes.

## Testing

```bash
npm test        # go test ./...
npm run vet     # go vet ./...
```

## Deployment

The app is a single Go binary plus the `public/` directory and `internal/site/templates`.

1. Build the stylesheet: `npm run build:css`
2. Build the binary: `go build -o bin/portfolio ./cmd/server`
3. Run it with the `public/` directory and `internal/` templates present, or
   ship the repo and run `go run ./cmd/server`:

```bash
AI_API_KEY=... AI_BASE_URL=... AI_MODEL=... PORT=3000 ./bin/portfolio
```

Set `NEXT_PUBLIC_BASE_URL` to the public origin so canonical/OG URLs are correct.
Any host that runs a Go binary works (a VM, a container, Fly.io, Railway, Render,
a Go buildpack platform). The previous Next.js hosting is no longer required.

## Notes

- The pre-Next static site is preserved under `legacy/` and was not deleted.
- Real assets only. No stock or generated images, no invented data.
- See README history / `docs/PRD.md` for the design rationale.
