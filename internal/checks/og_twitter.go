package checks

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type OGTwitterCheck struct{}

func (c OGTwitterCheck) ID() string {
	return "ogTwitter"
}

func (c OGTwitterCheck) Title() string {
	return "OG & Twitter cards configured"
}

func (c OGTwitterCheck) Run(ctx Context) (CheckResult, error) {
	cfg := ctx.Config.Checks.SEOMeta
	if cfg == nil || cfg.MainLayout == "" {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Check not configured (set checks.seoMeta.mainLayout)",
		}, nil
	}

	layoutPath := filepath.Join(ctx.RootDir, cfg.MainLayout)
	content, err := os.ReadFile(layoutPath)
	if err != nil {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityWarn,
			Passed:   false,
			Message:  "Could not read layout file: " + cfg.MainLayout,
		}, nil
	}

	contentStr := string(content)

	// OG and Twitter card elements
	checks := map[string]*regexp.Regexp{
		"og:image":      regexp.MustCompile(`(?i)<meta[^>]+property=["']og:image["'][^>]*>`),
		"og:url":        regexp.MustCompile(`(?i)<meta[^>]+property=["']og:url["'][^>]*>`),
		"og:type":       regexp.MustCompile(`(?i)<meta[^>]+property=["']og:type["'][^>]*>`),
		"twitter:card":  regexp.MustCompile(`(?i)<meta[^>]+name=["']twitter:card["'][^>]*>`),
		"twitter:image": regexp.MustCompile(`(?i)<meta[^>]+name=["']twitter:image["'][^>]*>`),
	}

	// Alternate patterns for Next.js/React metadata API
	alternates := map[string][]*regexp.Regexp{
		"og:image": {
			regexp.MustCompile(`(?i)openGraph.*images`),
			regexp.MustCompile(`(?i)og:image`),
			regexp.MustCompile(`(?i)opengraph-image\.(png|jpg|jpeg|svg)`),
		},
		"og:url": {
			regexp.MustCompile(`(?i)openGraph.*url`),
			regexp.MustCompile(`(?i)metadataBase`),
		},
		"og:type": {
			regexp.MustCompile(`(?i)openGraph.*type`),
		},
		"twitter:card": {
			regexp.MustCompile(`(?i)twitter.*card`),
			regexp.MustCompile(`(?i)twitter-image\.(png|jpg|jpeg|svg)`),
		},
		"twitter:image": {
			regexp.MustCompile(`(?i)twitter.*images`),
			regexp.MustCompile(`(?i)twitter-image\.(png|jpg|jpeg|svg)`),
		},
	}

	var missing []string
	var found []string

	for name, pattern := range checks {
		matched := pattern.MatchString(contentStr)

		// Try alternate patterns
		if !matched {
			if alts, ok := alternates[name]; ok {
				for _, alt := range alts {
					if alt.MatchString(contentStr) {
						matched = true
						break
					}
				}
			}
		}

		if matched {
			found = append(found, name)
		} else {
			missing = append(missing, name)
		}
	}

	// Also check for opengraph-image and twitter-image files in app directory
	ogImageFiles := []string{
		"app/opengraph-image.png",
		"app/opengraph-image.jpg",
		"app/twitter-image.png",
		"app/twitter-image.jpg",
		"public/og-image.png",
		"public/og-image.jpg",
		"public/og.png",
		"public/twitter-image.png",
	}

	for _, imgPath := range ogImageFiles {
		fullPath := filepath.Join(ctx.RootDir, imgPath)
		if _, err := os.Stat(fullPath); err == nil {
			if strings.Contains(imgPath, "opengraph") || strings.Contains(imgPath, "og") {
				// Remove og:image from missing if found
				missing = removeFromSlice(missing, "og:image")
				if !contains(found, "og:image") {
					found = append(found, "og:image (file)")
				}
			}
			if strings.Contains(imgPath, "twitter") {
				missing = removeFromSlice(missing, "twitter:image")
				if !contains(found, "twitter:image") {
					found = append(found, "twitter:image (file)")
				}
			}
		}
	}

	if len(missing) == 0 {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "OG and Twitter card metadata configured",
		}, nil
	}

	// Warn if missing image tags specifically
	severity := SeverityWarn
	if len(missing) <= 2 && !contains(missing, "og:image") {
		severity = SeverityWarn
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: severity,
		Passed:   false,
		Message:  "Missing: " + strings.Join(missing, ", "),
		Suggestions: []string{
			"Add og:image for rich social media previews",
			"Add twitter:card for Twitter/X previews",
			"Consider using 1200x630px images for best results",
		},
	}, nil
}

func removeFromSlice(slice []string, item string) []string {
	var result []string
	for _, s := range slice {
		if s != item {
			result = append(result, s)
		}
	}
	return result
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
