package checks

import (
	"regexp"
)

// BugsnagCheck verifies Bugsnag is properly set up
type BugsnagCheck struct{}

func (c BugsnagCheck) ID() string {
	return "bugsnag"
}

func (c BugsnagCheck) Title() string {
	return "Bugsnag"
}

func (c BugsnagCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["bugsnag"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Bugsnag not declared, skipping",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`Bugsnag\.start`),
		regexp.MustCompile(`bugsnag\.notify`),
		regexp.MustCompile(`@bugsnag/`),
		regexp.MustCompile(`bugsnag-js`),
		regexp.MustCompile(`Bugsnag\.configure`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Bugsnag initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Bugsnag is declared but initialization not found",
		Suggestions: []string{
			"Add Bugsnag.start() to your application entry point",
			"Check Bugsnag docs for your framework",
		},
	}, nil
}

// RollbarCheck verifies Rollbar is properly set up
type RollbarCheck struct{}

func (c RollbarCheck) ID() string {
	return "rollbar"
}

func (c RollbarCheck) Title() string {
	return "Rollbar"
}

func (c RollbarCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["rollbar"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Rollbar not declared, skipping",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`Rollbar\.init`),
		regexp.MustCompile(`Rollbar\.configure`),
		regexp.MustCompile(`rollbar\.com`),
		regexp.MustCompile(`@rollbar/`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Rollbar initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Rollbar is declared but initialization not found",
		Suggestions: []string{
			"Add Rollbar.init() with your access token",
			"Check Rollbar docs for your framework",
		},
	}, nil
}

// HoneybadgerCheck verifies Honeybadger is properly set up
type HoneybadgerCheck struct{}

func (c HoneybadgerCheck) ID() string {
	return "honeybadger"
}

func (c HoneybadgerCheck) Title() string {
	return "Honeybadger"
}

func (c HoneybadgerCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["honeybadger"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Honeybadger not declared, skipping",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`Honeybadger\.configure`),
		regexp.MustCompile(`Honeybadger\.notify`),
		regexp.MustCompile(`@honeybadger-io/`),
		regexp.MustCompile(`honeybadger-js`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Honeybadger initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Honeybadger is declared but initialization not found",
		Suggestions: []string{
			"Add Honeybadger.configure() with your API key",
			"Check Honeybadger docs for your framework",
		},
	}, nil
}

// DatadogCheck verifies Datadog is properly set up
type DatadogCheck struct{}

func (c DatadogCheck) ID() string {
	return "datadog"
}

func (c DatadogCheck) Title() string {
	return "Datadog"
}

func (c DatadogCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["datadog"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Datadog not declared, skipping",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`datadogRum\.init`),
		regexp.MustCompile(`DD_RUM`),
		regexp.MustCompile(`dd-trace`),
		regexp.MustCompile(`@datadog/`),
		regexp.MustCompile(`datadoghq\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Datadog initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Datadog is declared but initialization not found",
		Suggestions: []string{
			"Add Datadog RUM or APM initialization",
			"Check Datadog docs for your framework",
		},
	}, nil
}

// NewRelicCheck verifies New Relic is properly set up
type NewRelicCheck struct{}

func (c NewRelicCheck) ID() string {
	return "newrelic"
}

func (c NewRelicCheck) Title() string {
	return "New Relic"
}

func (c NewRelicCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["newrelic"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "New Relic not declared, skipping",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`newrelic`),
		regexp.MustCompile(`@newrelic/`),
		regexp.MustCompile(`NREUM`),
		regexp.MustCompile(`nr-data\.net`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "New Relic initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "New Relic is declared but initialization not found",
		Suggestions: []string{
			"Add New Relic browser agent or APM",
			"Check New Relic docs for your framework",
		},
	}, nil
}

// LogRocketCheck verifies LogRocket is properly set up
type LogRocketCheck struct{}

func (c LogRocketCheck) ID() string {
	return "logrocket"
}

func (c LogRocketCheck) Title() string {
	return "LogRocket"
}

func (c LogRocketCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["logrocket"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "LogRocket not declared, skipping",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`LogRocket\.init`),
		regexp.MustCompile(`logrocket`),
		regexp.MustCompile(`cdn\.logrocket\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "LogRocket initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "LogRocket is declared but initialization not found",
		Suggestions: []string{
			"Add LogRocket.init() with your app ID",
			"Check LogRocket docs for your framework",
		},
	}, nil
}
