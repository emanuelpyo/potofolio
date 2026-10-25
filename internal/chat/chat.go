package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"deryl-portfolio/internal/config"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	Text    string    `json:"text"`
	History []message `json:"history"`
}

type completionRequest struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	Temperature float64   `json:"temperature"`
	TopP        float64   `json:"top_p"`
	MaxTokens   int       `json:"max_tokens"`
	Stream      bool      `json:"stream"`
}

type streamChunk struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

// Handler streams a chat completion as plain text. The API key is read from the
// server environment only and never reaches the browser.
type Handler struct {
	cfg    config.Config
	client *http.Client
}

func New(cfg config.Config) *Handler {
	return &Handler{
		cfg: cfg,
		client: &http.Client{
			Transport: &http.Transport{
				ResponseHeaderTimeout: 30 * time.Second,
			},
		},
	}
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"success": false, "message": msg})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method tidak diizinkan")
		return
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Pesan tidak boleh kosong")
		return
	}
	if strings.TrimSpace(req.Text) == "" {
		writeJSONError(w, http.StatusBadRequest, "Pesan tidak boleh kosong")
		return
	}
	if h.cfg.AIAPIKey == "" {
		log.Println("AI_API_KEY is missing in environment variables")
		writeJSONError(w, http.StatusInternalServerError, "Konfigurasi API AI tidak ditemukan.")
		return
	}

	messages := append(req.History, message{Role: "user", Content: req.Text})
	body, err := json.Marshal(completionRequest{
		Model:       h.cfg.AIModel,
		Messages:    messages,
		Temperature: 0.7,
		TopP:        0.7,
		MaxTokens:   1024,
		Stream:      true,
	})
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Gagal terhubung ke server AI.")
		return
	}

	endpoint := strings.TrimRight(h.cfg.AIBaseURL, "/") + "/chat/completions"
	upstream, err := http.NewRequestWithContext(r.Context(), http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Gagal terhubung ke server AI.")
		return
	}
	upstream.Header.Set("Content-Type", "application/json")
	upstream.Header.Set("Accept", "text/event-stream")
	upstream.Header.Set("Authorization", "Bearer "+h.cfg.AIAPIKey)

	resp, err := h.client.Do(upstream)
	if err != nil {
		log.Printf("Chat API Error: %v", err)
		writeJSONError(w, http.StatusInternalServerError, "Gagal terhubung ke server AI.")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Chat upstream status: %d", resp.StatusCode)
		writeJSONError(w, http.StatusInternalServerError, "Gagal terhubung ke server AI.")
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-transform")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	flusher, _ := w.(http.Flusher)
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data:") {
			continue
		}
		payload := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if payload == "" || payload == "[DONE]" {
			if payload == "[DONE]" {
				break
			}
			continue
		}
		var chunk streamChunk
		if err := json.Unmarshal([]byte(payload), &chunk); err != nil {
			continue
		}
		if len(chunk.Choices) == 0 || chunk.Choices[0].Delta.Content == "" {
			continue
		}
		if _, err := fmt.Fprint(w, chunk.Choices[0].Delta.Content); err != nil {
			return
		}
		if flusher != nil {
			flusher.Flush()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Chat stream error: %v", err)
	}
}
