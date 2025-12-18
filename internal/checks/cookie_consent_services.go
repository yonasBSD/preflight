package checks

import (
	"regexp"
)

// CookieConsentJSCheck verifies CookieConsent JS library is properly set up
type CookieConsentJSCheck struct{}

func (c CookieConsentJSCheck) ID() string {
	return "cookieconsent"
}

func (c CookieConsentJSCheck) Title() string {
	return "CookieConsent"
}

func (c CookieConsentJSCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["cookieconsent"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cookie Consent not declared, skipping",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`cookieconsent`),
		regexp.MustCompile(`CookieConsent`),
		regexp.MustCompile(`cdn\.jsdelivr\.net.*cookieconsent`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cookie Consent script found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Cookie Consent is declared but script not found",
		Suggestions: []string{
			"Add Cookie Consent script to your templates",
		},
	}, nil
}

// CookiebotCheck verifies Cookiebot is properly set up
type CookiebotCheck struct{}

func (c CookiebotCheck) ID() string {
	return "cookiebot"
}

func (c CookiebotCheck) Title() string {
	return "Cookiebot"
}

func (c CookiebotCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["cookiebot"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cookiebot not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "COOKIEBOT_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cookiebot configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`consent\.cookiebot\.com`),
		regexp.MustCompile(`Cookiebot`),
		regexp.MustCompile(`cookiebot`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cookiebot script found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Cookiebot is declared but script not found",
		Suggestions: []string{
			"Add Cookiebot script to your templates",
			"Add COOKIEBOT_CBID to environment",
		},
	}, nil
}

// OneTrustCheck verifies OneTrust is properly set up
type OneTrustCheck struct{}

func (c OneTrustCheck) ID() string {
	return "onetrust"
}

func (c OneTrustCheck) Title() string {
	return "OneTrust"
}

func (c OneTrustCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["onetrust"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "OneTrust not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "ONETRUST_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "OneTrust configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`cdn\.cookielaw\.org`),
		regexp.MustCompile(`onetrust`),
		regexp.MustCompile(`OneTrust`),
		regexp.MustCompile(`optanon`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "OneTrust script found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "OneTrust is declared but script not found",
		Suggestions: []string{
			"Add OneTrust script to your templates",
		},
	}, nil
}

// TermlyCheck verifies Termly is properly set up
type TermlyCheck struct{}

func (c TermlyCheck) ID() string {
	return "termly"
}

func (c TermlyCheck) Title() string {
	return "Termly"
}

func (c TermlyCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["termly"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Termly not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "TERMLY_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Termly configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`app\.termly\.io`),
		regexp.MustCompile(`termly`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Termly script found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Termly is declared but script not found",
		Suggestions: []string{
			"Add Termly consent banner script to your templates",
		},
	}, nil
}

// CookieYesCheck verifies CookieYes is properly set up
type CookieYesCheck struct{}

func (c CookieYesCheck) ID() string {
	return "cookieyes"
}

func (c CookieYesCheck) Title() string {
	return "CookieYes"
}

func (c CookieYesCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["cookieyes"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "CookieYes not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "COOKIEYES_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "CookieYes configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`cdn-cookieyes\.com`),
		regexp.MustCompile(`cookieyes`),
		regexp.MustCompile(`CookieYes`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "CookieYes script found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "CookieYes is declared but script not found",
		Suggestions: []string{
			"Add CookieYes script to your templates",
		},
	}, nil
}

// IubendaCheck verifies Iubenda is properly set up
type IubendaCheck struct{}

func (c IubendaCheck) ID() string {
	return "iubenda"
}

func (c IubendaCheck) Title() string {
	return "Iubenda"
}

func (c IubendaCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["iubenda"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Iubenda not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "IUBENDA_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Iubenda configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`cdn\.iubenda\.com`),
		regexp.MustCompile(`iubenda`),
		regexp.MustCompile(`_iub`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Iubenda script found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Iubenda is declared but script not found",
		Suggestions: []string{
			"Add Iubenda cookie banner script to your templates",
		},
	}, nil
}
