package checks

import (
	"os"
	"path/filepath"
	"strings"
)

type LegalPagesCheck struct{}

func (c LegalPagesCheck) ID() string {
	return "legal_pages"
}

func (c LegalPagesCheck) Title() string {
	return "Privacy & Terms pages"
}

func (c LegalPagesCheck) Run(ctx Context) (CheckResult, error) {
	hasPrivacy := false
	hasTerms := false
	var privacyPath, termsPath string

	// Common privacy policy paths/filenames
	privacyPatterns := []string{
		"privacy", "privacy-policy", "privacy_policy",
		"legal/privacy", "legal/privacy-policy",
		"pages/privacy", "pages/privacy-policy",
	}

	// Common terms paths/filenames
	termsPatterns := []string{
		"terms", "terms-of-service", "terms_of_service", "tos",
		"legal/terms", "legal/terms-of-service",
		"pages/terms", "pages/terms-of-service",
	}

	// Extensions to check
	extensions := []string{
		"", ".html", ".htm", ".php", ".md", ".mdx",
		".tsx", ".jsx", ".js", ".ts", ".vue", ".svelte",
		".erb", ".erb.html", ".html.erb",
		".blade.php", ".twig", ".njk", ".liquid",
		".astro",
	}

	// Directories to search
	searchDirs := []string{
		"",
		"app",
		"src/app",
		"src/pages",
		"pages",
		"views",
		"resources/views",
		"templates",
		"content",
		"public",
		"static",
	}

	// Check for privacy policy
	for _, dir := range searchDirs {
		if hasPrivacy {
			break
		}
		for _, pattern := range privacyPatterns {
			if hasPrivacy {
				break
			}
			for _, ext := range extensions {
				checkPath := filepath.Join(ctx.RootDir, dir, pattern+ext)
				if _, err := os.Stat(checkPath); err == nil {
					hasPrivacy = true
					privacyPath = filepath.Join(dir, pattern+ext)
					break
				}
				// Also check with /page pattern for Next.js app router
				if dir == "app" || dir == "src/app" {
					pagePath := filepath.Join(ctx.RootDir, dir, pattern, "page"+ext)
					if _, err := os.Stat(pagePath); err == nil {
						hasPrivacy = true
						privacyPath = filepath.Join(dir, pattern, "page"+ext)
						break
					}
				}
			}
		}
	}

	// Check for terms
	for _, dir := range searchDirs {
		if hasTerms {
			break
		}
		for _, pattern := range termsPatterns {
			if hasTerms {
				break
			}
			for _, ext := range extensions {
				checkPath := filepath.Join(ctx.RootDir, dir, pattern+ext)
				if _, err := os.Stat(checkPath); err == nil {
					hasTerms = true
					termsPath = filepath.Join(dir, pattern+ext)
					break
				}
				// Also check with /page pattern for Next.js app router
				if dir == "app" || dir == "src/app" {
					pagePath := filepath.Join(ctx.RootDir, dir, pattern, "page"+ext)
					if _, err := os.Stat(pagePath); err == nil {
						hasTerms = true
						termsPath = filepath.Join(dir, pattern, "page"+ext)
						break
					}
				}
			}
		}
	}

	// Also check main layout for links to privacy/terms
	if !hasPrivacy || !hasTerms {
		if ctx.Config.Checks.SEOMeta != nil && ctx.Config.Checks.SEOMeta.MainLayout != "" {
			layoutPath := filepath.Join(ctx.RootDir, ctx.Config.Checks.SEOMeta.MainLayout)
			if content, err := os.ReadFile(layoutPath); err == nil {
				contentLower := strings.ToLower(string(content))
				if !hasPrivacy && (strings.Contains(contentLower, "/privacy") || strings.Contains(contentLower, "privacy-policy")) {
					hasPrivacy = true
					privacyPath = "linked in layout"
				}
				if !hasTerms && (strings.Contains(contentLower, "/terms") || strings.Contains(contentLower, "terms-of-service")) {
					hasTerms = true
					termsPath = "linked in layout"
				}
			}
		}
	}

	if hasPrivacy && hasTerms {
		msg := "Found"
		if privacyPath != "linked in layout" {
			msg += " privacy at " + privacyPath
		} else {
			msg += " privacy link"
		}
		if termsPath != "linked in layout" {
			msg += ", terms at " + termsPath
		} else {
			msg += ", terms link"
		}
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  msg,
		}, nil
	}

	var missing []string
	if !hasPrivacy {
		missing = append(missing, "privacy policy")
	}
	if !hasTerms {
		missing = append(missing, "terms of service")
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Missing: " + strings.Join(missing, ", "),
		Suggestions: []string{
			"Add a privacy policy page (e.g., /privacy)",
			"Add terms of service page (e.g., /terms)",
		},
	}, nil
}
