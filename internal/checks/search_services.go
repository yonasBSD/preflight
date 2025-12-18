package checks

import (
	"regexp"
)

// AlgoliaCheck verifies Algolia is properly set up
type AlgoliaCheck struct{}

func (c AlgoliaCheck) ID() string {
	return "algolia"
}

func (c AlgoliaCheck) Title() string {
	return "Algolia"
}

func (c AlgoliaCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["algolia"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Algolia not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "ALGOLIA_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Algolia configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`algoliasearch`),
		regexp.MustCompile(`@algolia/`),
		regexp.MustCompile(`algolia\.com`),
		regexp.MustCompile(`InstantSearch`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Algolia SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Algolia is declared but SDK not found",
		Suggestions: []string{
			"Add ALGOLIA_APP_ID and ALGOLIA_API_KEY to environment",
			"Initialize Algolia search client in your application",
		},
	}, nil
}
