package checks

import (
	"os"
	"path/filepath"
	"strings"
)

// RobotsTxtCheck verifies robots.txt exists
type RobotsTxtCheck struct{}

func (c RobotsTxtCheck) ID() string {
	return "robotsTxt"
}

func (c RobotsTxtCheck) Title() string {
	return "robots.txt is present"
}

func (c RobotsTxtCheck) Run(ctx Context) (CheckResult, error) {
	paths := []string{
		"public/robots.txt",
		"static/robots.txt",
		"robots.txt",
		"app/robots.txt",
	}

	for _, path := range paths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if content, err := os.ReadFile(fullPath); err == nil {
			// Check if it has meaningful content
			contentStr := strings.TrimSpace(string(content))
			if len(contentStr) > 0 {
				return CheckResult{
					ID:       c.ID(),
					Title:    c.Title(),
					Severity: SeverityInfo,
					Passed:   true,
					Message:  "robots.txt found at " + path,
				}, nil
			}
		}
	}

	// Also check for Next.js robots.ts/js
	nextRobotsPaths := []string{
		"app/robots.ts",
		"app/robots.js",
	}

	for _, path := range nextRobotsPaths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			return CheckResult{
				ID:       c.ID(),
				Title:    c.Title(),
				Severity: SeverityInfo,
				Passed:   true,
				Message:  "robots.txt generated via " + path,
			}, nil
		}
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "robots.txt not found",
		Suggestions: []string{
			"Add robots.txt to public/ directory",
			"Include Sitemap directive pointing to sitemap.xml",
		},
	}, nil
}

// SitemapCheck verifies sitemap.xml exists
type SitemapCheck struct{}

func (c SitemapCheck) ID() string {
	return "sitemap"
}

func (c SitemapCheck) Title() string {
	return "sitemap.xml is present"
}

func (c SitemapCheck) Run(ctx Context) (CheckResult, error) {
	paths := []string{
		"public/sitemap.xml",
		"static/sitemap.xml",
		"sitemap.xml",
	}

	for _, path := range paths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			return CheckResult{
				ID:       c.ID(),
				Title:    c.Title(),
				Severity: SeverityInfo,
				Passed:   true,
				Message:  "sitemap.xml found at " + path,
			}, nil
		}
	}

	// Check for Next.js sitemap generator
	nextSitemapPaths := []string{
		"app/sitemap.ts",
		"app/sitemap.js",
		"app/sitemap.xml/route.ts",
	}

	for _, path := range nextSitemapPaths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			return CheckResult{
				ID:       c.ID(),
				Title:    c.Title(),
				Severity: SeverityInfo,
				Passed:   true,
				Message:  "sitemap.xml generated via " + path,
			}, nil
		}
	}

	// Check for sitemap generation in package.json scripts or next-sitemap
	pkgPath := filepath.Join(ctx.RootDir, "package.json")
	if content, err := os.ReadFile(pkgPath); err == nil {
		if strings.Contains(string(content), "next-sitemap") ||
			strings.Contains(string(content), "sitemap") {
			return CheckResult{
				ID:       c.ID(),
				Title:    c.Title(),
				Severity: SeverityInfo,
				Passed:   true,
				Message:  "Sitemap generation configured via package",
			}, nil
		}
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "sitemap.xml not found",
		Suggestions: []string{
			"Add sitemap.xml to public/ directory",
			"Consider using next-sitemap or similar generator",
		},
	}, nil
}

// LLMsTxtCheck verifies llms.txt exists for AI crawlers
type LLMsTxtCheck struct{}

func (c LLMsTxtCheck) ID() string {
	return "llmsTxt"
}

func (c LLMsTxtCheck) Title() string {
	return "llms.txt is present"
}

func (c LLMsTxtCheck) Run(ctx Context) (CheckResult, error) {
	paths := []string{
		"public/llms.txt",
		"static/llms.txt",
		"llms.txt",
		"public/.well-known/llms.txt",
	}

	for _, path := range paths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			return CheckResult{
				ID:       c.ID(),
				Title:    c.Title(),
				Severity: SeverityInfo,
				Passed:   true,
				Message:  "llms.txt found at " + path,
			}, nil
		}
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "llms.txt not found",
		Suggestions: []string{
			"Add llms.txt to help AI understand your site",
			"See https://llmstxt.org for specification",
		},
	}, nil
}

// AdsTxtCheck verifies ads.txt exists (optional, for ad-supported sites)
type AdsTxtCheck struct{}

func (c AdsTxtCheck) ID() string {
	return "adsTxt"
}

func (c AdsTxtCheck) Title() string {
	return "ads.txt is present"
}

func (c AdsTxtCheck) Run(ctx Context) (CheckResult, error) {
	// Check if ads.txt check is enabled in config
	// This is optional - only matters for ad-supported sites
	if ctx.Config.Checks.AdsTxt == nil || !ctx.Config.Checks.AdsTxt.Enabled {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "ads.txt check not enabled",
		}, nil
	}

	paths := []string{
		"public/ads.txt",
		"static/ads.txt",
		"ads.txt",
	}

	for _, path := range paths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if _, err := os.Stat(fullPath); err == nil {
			return CheckResult{
				ID:       c.ID(),
				Title:    c.Title(),
				Severity: SeverityInfo,
				Passed:   true,
				Message:  "ads.txt found at " + path,
			}, nil
		}
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "ads.txt not found",
		Suggestions: []string{
			"Add ads.txt for authorized digital sellers",
			"Required if running programmatic ads",
		},
	}, nil
}
