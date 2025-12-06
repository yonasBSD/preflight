package checks

import (
	"os"
	"path/filepath"
	"strings"
)

type LicenseCheck struct{}

func (c LicenseCheck) ID() string {
	return "license"
}

func (c LicenseCheck) Title() string {
	return "LICENSE file is present"
}

func (c LicenseCheck) Run(ctx Context) (CheckResult, error) {
	paths := []string{
		"LICENSE",
		"LICENSE.md",
		"LICENSE.txt",
		"LICENCE",
		"LICENCE.md",
		"license",
		"license.md",
		"license.txt",
	}

	for _, path := range paths {
		fullPath := filepath.Join(ctx.RootDir, path)
		if content, err := os.ReadFile(fullPath); err == nil {
			contentStr := strings.TrimSpace(string(content))
			if len(contentStr) > 0 {
				// Try to detect license type
				licenseType := detectLicenseType(contentStr)
				message := "LICENSE file found"
				if licenseType != "" {
					message = licenseType + " license found"
				}
				return CheckResult{
					ID:       c.ID(),
					Title:    c.Title(),
					Severity: SeverityInfo,
					Passed:   true,
					Message:  message,
				}, nil
			}
		}
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "No LICENSE file found",
		Suggestions: []string{
			"Add a LICENSE file to your project",
			"Choose a license at https://choosealicense.com",
		},
	}, nil
}

func detectLicenseType(content string) string {
	contentLower := strings.ToLower(content)

	if strings.Contains(contentLower, "mit license") ||
		strings.Contains(contentLower, "permission is hereby granted, free of charge") {
		return "MIT"
	}

	if strings.Contains(contentLower, "apache license") &&
		strings.Contains(contentLower, "version 2.0") {
		return "Apache 2.0"
	}

	if strings.Contains(contentLower, "gnu general public license") {
		if strings.Contains(contentLower, "version 3") {
			return "GPL-3.0"
		}
		if strings.Contains(contentLower, "version 2") {
			return "GPL-2.0"
		}
		return "GPL"
	}

	if strings.Contains(contentLower, "bsd") {
		if strings.Contains(contentLower, "3-clause") || strings.Contains(contentLower, "three-clause") {
			return "BSD-3-Clause"
		}
		if strings.Contains(contentLower, "2-clause") || strings.Contains(contentLower, "two-clause") {
			return "BSD-2-Clause"
		}
		return "BSD"
	}

	if strings.Contains(contentLower, "isc license") {
		return "ISC"
	}

	if strings.Contains(contentLower, "mozilla public license") {
		return "MPL-2.0"
	}

	if strings.Contains(contentLower, "unlicense") ||
		strings.Contains(contentLower, "this is free and unencumbered") {
		return "Unlicense"
	}

	if strings.Contains(contentLower, "creative commons") {
		return "Creative Commons"
	}

	if strings.Contains(contentLower, "proprietary") ||
		strings.Contains(contentLower, "all rights reserved") {
		return "Proprietary"
	}

	return ""
}
