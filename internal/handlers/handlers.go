package handlers

import (
	"ascii-art-web-stylize/internal/config"
	"ascii-art-web-stylize/internal/service"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

type Handler struct {
	cfg      *config.Config
	service  *service.AsciiService
	template *template.Template
	asciiArt string
}

func New(cfg *config.Config) *Handler {
	return &Handler{
		cfg:      cfg,
		service:  service.New(cfg),
		template: template.Must(template.ParseFiles(cfg.TemplatePath)),
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// Ensure only GET requests are accepted
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	banners, err := h.service.GetAvailableBanners()
	if err != nil {
		http.Error(w, "Failed to load banners", http.StatusInternalServerError)
		return
	}

	data := struct {
		Art     string
		Banners []string
	}{
		Art:     h.asciiArt,
		Banners: banners,
	}

	// Render the template
	if err := h.template.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Generate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	input := r.FormValue("inputText")
	banner := r.FormValue("banner")

	if input == "" {
		http.Error(w, "Text input is required", http.StatusBadRequest)
		return
	}

	latinRegex := regexp.MustCompile(`^[A-Za-z0-9\s.,!?@#\$%\^&\*\(\)\-\+=\[\]\{\}\<\>\"\\/\~\|]+$`)
	if !latinRegex.MatchString(input) {
		http.Error(w, "Error 400: Only Latin characters, numbers and allowed special characters are permitted", http.StatusBadRequest)
		return
	}

	asciiArt, err := h.service.GenerateASCIIArt(input, banner)
	if err != nil {
		http.Error(w, "Error 500: Failed to generate ASCII art", http.StatusInternalServerError)
		return
	}

	// Store the generated ASCII art
	h.asciiArt = asciiArt

	// Redirect to home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Export(w http.ResponseWriter, r *http.Request) {
	// Ensure only GET requests are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if ASCII art exists
	if h.asciiArt == "" {
		http.Error(w, "No ASCII art available for export", http.StatusBadRequest)
		return
	}

	// Determine file format from query parameters
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "txt" // Default format
	}

	var content []byte
	var fileName string

	switch format {
	case "txt":
		fileName = "ascii_art.txt"
		content = []byte(h.asciiArt)
	case "json":
		fileName = "ascii_art.json"
		content = []byte(fmt.Sprintf(`{"ascii_art": "%s"}`, h.asciiArt))
	case "xml":
		fileName = "ascii_art.xml"
		content = []byte(fmt.Sprintf(`<ascii_art>%s</ascii_art>`, h.asciiArt))
	default:
		http.Error(w, "Unsupported format", http.StatusBadRequest)
		return
	}

	// Set response headers
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))

	// Write content directly to response
	if _, err := w.Write(content); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
