package checks

import (
	"os"
	"path/filepath"
	"regexp"
)

type FaviconCheck struct{}

func (c FaviconCheck) ID() string {
	return "favicon"
}

func (c FaviconCheck) Title() string {
	return "Favicon and app icons present"
}

func (c FaviconCheck) Run(ctx Context) (CheckResult, error) {
	var found []string
	var missing []string

	// Check for common favicon locations
	faviconPaths := []string{
		"public/favicon.ico",
		"public/favicon.png",
		"public/favicon.svg",
		"favicon.ico",
		"app/favicon.ico",
		"app/icon.png",
		"app/icon.svg",
		"static/favicon.ico",
	}

	hasFavicon := false
	for _, path := range faviconPaths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			hasFavicon = true
			found = append(found, path)
			break
		}
	}

	if !hasFavicon {
		missing = append(missing, "favicon")
	}

	// Check for Apple Touch Icon
	appleTouchPaths := []string{
		"public/apple-touch-icon.png",
		"public/apple-icon.png",
		"app/apple-icon.png",
		"app/apple-touch-icon.png",
		"static/apple-touch-icon.png",
	}

	hasAppleIcon := false
	for _, path := range appleTouchPaths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			hasAppleIcon = true
			found = append(found, path)
			break
		}
	}

	// Also check HTML for apple-touch-icon link
	if !hasAppleIcon {
		cfg := ctx.Config.Checks.SEOMeta
		if cfg != nil && cfg.MainLayout != "" {
			layoutPath := filepath.Join(ctx.RootDir, cfg.MainLayout)
			if content, err := os.ReadFile(layoutPath); err == nil {
				if regexp.MustCompile(`(?i)apple-touch-icon`).Match(content) {
					hasAppleIcon = true
					found = append(found, "apple-touch-icon (in HTML)")
				}
			}
		}
	}

	if !hasAppleIcon {
		missing = append(missing, "apple-touch-icon")
	}

	// Check for web app manifest
	manifestPaths := []string{
		"public/manifest.json",
		"public/site.webmanifest",
		"app/manifest.json",
		"app/manifest.ts",
		"app/manifest.js",
		"static/manifest.json",
	}

	hasManifest := false
	for _, path := range manifestPaths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			hasManifest = true
			found = append(found, path)
			break
		}
	}

	if !hasManifest {
		missing = append(missing, "web manifest")
	}

	// Determine result
	if len(missing) == 0 {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "All icons and manifest present",
		}, nil
	}

	if hasFavicon && len(missing) <= 2 {
		// Has favicon but missing apple icon or manifest - just warn
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityWarn,
			Passed:   false,
			Message:  "Missing: " + joinStrings(missing, ", "),
			Suggestions: []string{
				"Add apple-touch-icon.png (180x180px) for iOS",
				"Add manifest.json for PWA support",
			},
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityError,
		Passed:   false,
		Message:  "Missing favicon",
		Suggestions: []string{
			"Add favicon.ico or favicon.png to public/",
			"Use https://realfavicongenerator.net for complete icon set",
		},
	}, nil
}

func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}
