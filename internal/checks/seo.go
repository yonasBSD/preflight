package checks

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type SEOMetadataCheck struct{}

func (c SEOMetadataCheck) ID() string {
	return "seoMeta"
}

func (c SEOMetadataCheck) Title() string {
	return "SEO metadata"
}

func (c SEOMetadataCheck) Run(ctx Context) (CheckResult, error) {
	cfg := ctx.Config.Checks.SEOMeta

	// Get configured layout or auto-detect
	var configuredLayout string
	if cfg != nil {
		configuredLayout = cfg.MainLayout
	}
	layoutFile := getLayoutFile(ctx.RootDir, ctx.Config.Stack, configuredLayout)

	if layoutFile == "" {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "No layout file found, skipping",
		}, nil
	}

	layoutPath := filepath.Join(ctx.RootDir, layoutFile)
	content, err := os.ReadFile(layoutPath)
	if err != nil {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityWarn,
			Passed:   false,
			Message:  "Could not read layout file: " + cfg.MainLayout,
			Suggestions: []string{
				"Check that the mainLayout path is correct in preflight.yml",
			},
		}, nil
	}

	contentStr := string(content)

	// Required SEO elements
	checks := map[string]*regexp.Regexp{
		"title":          regexp.MustCompile(`<title[^>]*>`),
		"description":    regexp.MustCompile(`<meta[^>]+name=["']description["'][^>]*>`),
		"og:title":       regexp.MustCompile(`<meta[^>]+property=["']og:title["'][^>]*>`),
		"og:description": regexp.MustCompile(`<meta[^>]+property=["']og:description["'][^>]*>`),
	}

	var missing []string
	for name, pattern := range checks {
		if !pattern.MatchString(contentStr) {
			// Check for alternate patterns (some frameworks use different formats)
			if !checkAlternatePatterns(contentStr, name) {
				missing = append(missing, name)
			}
		}
	}

	if len(missing) == 0 {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "All required SEO metadata present",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Missing SEO metadata: " + strings.Join(missing, ", "),
		Suggestions: []string{
			"Add missing meta tags to your layout",
			"Consider using a SEO component or helper",
		},
	}, nil
}

// getLayoutFile returns the configured layout or auto-detects one based on stack
func getLayoutFile(rootDir string, stack string, configuredLayout string) string {
	// Use configured layout if set
	if configuredLayout != "" {
		return configuredLayout
	}

	// Auto-detect based on stack
	layoutsByStack := map[string][]string{
		"next": {
			"app/layout.tsx", "app/layout.js", "app/layout.jsx",
			"src/app/layout.tsx", "src/app/layout.js",
			"pages/_app.tsx", "pages/_app.js", "pages/_document.tsx", "pages/_document.js",
		},
		"react": {
			"index.html", "public/index.html", "src/index.html",
		},
		"vite": {
			"index.html", "src/index.html",
		},
		"vue": {
			"index.html", "public/index.html", "src/App.vue",
		},
		"svelte": {
			"src/app.html", "index.html",
		},
		"angular": {
			"src/index.html",
		},
		"rails": {
			"app/views/layouts/application.html.erb",
			"app/views/layouts/base.html.erb",
		},
		"laravel": {
			"resources/views/layouts/app.blade.php",
			"resources/views/layouts/main.blade.php",
		},
		"django": {
			"templates/base.html",
			"templates/layout.html",
		},
		"craft": {
			"templates/_layout.twig",
			"templates/_layouts/main.twig",
			"templates/_layouts/base.twig",
			"templates/_base.twig",
		},
		"hugo": {
			"layouts/_default/baseof.html",
			"layouts/_default/base.html",
		},
		"jekyll": {
			"_layouts/default.html",
			"_layouts/base.html",
		},
		"gatsby": {
			"src/components/layout.js",
			"src/components/Layout.js",
			"src/components/layout.tsx",
		},
		"astro": {
			"src/layouts/Layout.astro",
			"src/layouts/Base.astro",
			"src/layouts/BaseLayout.astro",
		},
		"eleventy": {
			"_includes/base.njk",
			"_includes/layout.njk",
		},
		"php": {
			"templates/layout.php",
			"includes/header.php",
			"layout.php",
		},
		"node": {
			"views/layout.ejs",
			"views/layout.pug",
			"views/layouts/main.hbs",
		},
	}

	// Try stack-specific layouts first
	if layouts, ok := layoutsByStack[stack]; ok {
		for _, layout := range layouts {
			if _, err := os.Stat(filepath.Join(rootDir, layout)); err == nil {
				return layout
			}
		}
	}

	// Fallback: try common layouts for any stack
	commonLayouts := []string{
		"app/layout.tsx", "app/layout.js",
		"index.html", "public/index.html",
		"templates/_layout.twig",
		"app/views/layouts/application.html.erb",
	}
	for _, layout := range commonLayouts {
		if _, err := os.Stat(filepath.Join(rootDir, layout)); err == nil {
			return layout
		}
	}

	return ""
}

func checkAlternatePatterns(content, name string) bool {
	alternates := map[string][]*regexp.Regexp{
		"title": {
			regexp.MustCompile(`\btitle\s*[:=]`),  // JSX/React
			regexp.MustCompile(`<Title>`),         // Next.js Head
		},
		"description": {
			regexp.MustCompile(`name:\s*["']description["']`),
			regexp.MustCompile(`<meta\s+name="description"`),
		},
		"og:title": {
			regexp.MustCompile(`property:\s*["']og:title["']`),
			regexp.MustCompile(`openGraph.*title`),
		},
		"og:description": {
			regexp.MustCompile(`property:\s*["']og:description["']`),
			regexp.MustCompile(`openGraph.*description`),
		},
	}

	if patterns, ok := alternates[name]; ok {
		for _, pattern := range patterns {
			if pattern.MatchString(content) {
				return true
			}
		}
	}

	return false
}
