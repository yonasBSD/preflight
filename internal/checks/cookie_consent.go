package checks

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/preflightsh/preflight/internal/config"
)

type CookieConsentCheck struct{}

func (c CookieConsentCheck) ID() string {
	return "cookie_consent"
}

func (c CookieConsentCheck) Title() string {
	return "Cookie consent"
}

func (c CookieConsentCheck) Run(ctx Context) (CheckResult, error) {
	// Check if any cookie consent service is declared
	consentServices := []string{
		"cookieconsent", "cookiebot", "onetrust",
		"termly", "cookieyes", "iubenda",
	}

	for _, svc := range consentServices {
		if ctx.Config.Services[svc].Declared {
			return CheckResult{
				ID:       c.ID(),
				Title:    c.Title(),
				Severity: SeverityInfo,
				Passed:   true,
				Message:  "Using " + formatConsentService(svc),
			}, nil
		}
	}

	// Also scan layout/templates for cookie consent patterns
	found, service := detectCookieConsentInCode(ctx.RootDir, ctx.Config)
	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Found " + service,
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "No cookie consent solution detected",
		Suggestions: []string{
			"Add a cookie consent banner (GDPR/CCPA compliance)",
			"Consider CookieConsent, Cookiebot, OneTrust, or similar",
		},
	}, nil
}

func formatConsentService(svc string) string {
	names := map[string]string{
		"cookieconsent": "CookieConsent",
		"cookiebot":     "Cookiebot",
		"onetrust":      "OneTrust",
		"termly":        "Termly",
		"cookieyes":     "CookieYes",
		"iubenda":       "Iubenda",
	}
	if name, ok := names[svc]; ok {
		return name
	}
	return svc
}

func detectCookieConsentInCode(rootDir string, cfg *config.PreflightConfig) (bool, string) {
	patterns := map[string]*regexp.Regexp{
		"CookieConsent": regexp.MustCompile(`(?i)cookieconsent|cookie-consent|CookieConsent\.run`),
		"Cookiebot":     regexp.MustCompile(`(?i)cookiebot|Cookiebot`),
		"OneTrust":      regexp.MustCompile(`(?i)onetrust|optanon|cdn\.cookielaw\.org`),
		"Termly":        regexp.MustCompile(`(?i)termly\.io|Termly`),
		"CookieYes":     regexp.MustCompile(`(?i)cookieyes|cdn-cookieyes\.com`),
		"Iubenda":       regexp.MustCompile(`(?i)iubenda|_iub\.csConfiguration`),
		"Klaro":         regexp.MustCompile(`(?i)klaro\.js|KlaroConfig`),
		"Tarteaucitron": regexp.MustCompile(`(?i)tarteaucitron`),
	}

	// Check main layout if configured
	if cfg != nil && cfg.Checks.SEOMeta != nil && cfg.Checks.SEOMeta.MainLayout != "" {
		layoutPath := filepath.Join(rootDir, cfg.Checks.SEOMeta.MainLayout)
		if content, err := os.ReadFile(layoutPath); err == nil {
			for name, pattern := range patterns {
				if pattern.Match(content) {
					return true, name
				}
			}
		}
	}

	// Check common layout/template files
	layoutFiles := []string{
		"app/layout.tsx", "app/layout.jsx", "app/layout.js",
		"src/app/layout.tsx", "src/app/layout.jsx",
		"pages/_app.tsx", "pages/_app.jsx", "pages/_app.js",
		"pages/_document.tsx", "pages/_document.jsx",
		"index.html", "public/index.html",
		"app/views/layouts/application.html.erb",
		"resources/views/layouts/app.blade.php",
		"templates/_layout.twig", "templates/base.twig",
	}

	for _, file := range layoutFiles {
		fullPath := filepath.Join(rootDir, file)
		if content, err := os.ReadFile(fullPath); err == nil {
			for name, pattern := range patterns {
				if pattern.Match(content) {
					return true, name
				}
			}
		}
	}

	// Check package.json for cookie consent packages
	pkgPath := filepath.Join(rootDir, "package.json")
	if content, err := os.ReadFile(pkgPath); err == nil {
		contentLower := strings.ToLower(string(content))
		if strings.Contains(contentLower, "cookieconsent") ||
			strings.Contains(contentLower, "cookie-consent") ||
			strings.Contains(contentLower, "js-cookie") {
			return true, "cookie consent package"
		}
	}

	return false, ""
}
