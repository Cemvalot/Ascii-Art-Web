package service

import (
	"ascii-art-web-stylize/internal/config"
	"fmt"
	"os"
	"strings"
)

// AsciiService handles ASCII art generation with different banner styles
type AsciiService struct {
	cfg          *config.Config
	validBanners map[string]bool // Map of supported banner styles
}

func New(cfg *config.Config) *AsciiService {
	return &AsciiService{
		cfg: cfg,
		validBanners: map[string]bool{
			"shadow":     true,
			"standard":   true,
			"thinkertoy": true,
		},
	}
}

func (s *AsciiService) DecToLine(dec int) int {
	if dec >= 32 && dec <= 126 {
		return (dec - 32) * 9
	}
	return 0
}

func (s *AsciiService) PrintLinesFromIndices(lines []string, index int) string {
	if index < 0 || index >= len(lines) {
		return ""
	}
	return lines[index]
}

func (s *AsciiService) GenerateASCIIArt(input, banner string) (string, error) {
	if !s.validBanners[banner] {
		return "", fmt.Errorf("invalid banner: %s", banner)
	}

	bannerFile := fmt.Sprintf("%s/%s.txt", s.cfg.BannerPath, banner)
	content, err := os.ReadFile(bannerFile)
	if err != nil {
		return "", fmt.Errorf("failed to load banner file: %v", err)
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	var result strings.Builder

	for _, word := range strings.Split(input, "\n") {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		charIndices := make([]int, len(word))
		for i, char := range word {
			charIndices[i] = s.DecToLine(int(char))
		}

		for line := 1; line < 9; line++ {
			for i := 0; i < len(word); i++ {
				result.WriteString(s.PrintLinesFromIndices(lines, charIndices[i]+line))
			}
			result.WriteString("\n")
		}
		result.WriteString("\n")
	}

	return result.String(), nil
}

// GetAvailableBanners returns a list of all available banner styles from the banner directory
func (s *AsciiService) GetAvailableBanners() ([]string, error) {
	files, err := os.ReadDir(s.cfg.BannerPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read banners directory: %v", err)
	}

	var banners []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			banners = append(banners, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}

	if len(banners) == 0 {
		return nil, fmt.Errorf("no banner files found")
	}
	return banners, nil
}
