package checks

import (
	"regexp"
)

// TwilioCheck verifies Twilio is properly set up
type TwilioCheck struct{}

func (c TwilioCheck) ID() string {
	return "twilio"
}

func (c TwilioCheck) Title() string {
	return "Twilio"
}

func (c TwilioCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["twilio"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Twilio not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "TWILIO_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Twilio configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@twilio/`),
		regexp.MustCompile(`Twilio\.Rest`),
		regexp.MustCompile(`twilio\.com`),
		regexp.MustCompile(`new Twilio\(`),
		regexp.MustCompile(`from\s+["']twilio["']`),
		regexp.MustCompile(`require\s*\(\s*["']twilio["']\)`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Twilio SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Twilio is declared but SDK not found",
		Suggestions: []string{
			"Add TWILIO_ACCOUNT_SID and TWILIO_AUTH_TOKEN to environment",
		},
	}, nil
}

// SlackCheck verifies Slack is properly set up
type SlackCheck struct{}

func (c SlackCheck) ID() string {
	return "slack"
}

func (c SlackCheck) Title() string {
	return "Slack"
}

func (c SlackCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["slack"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Slack not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "SLACK_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Slack configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@slack/`),
		regexp.MustCompile(`slack-ruby`),
		regexp.MustCompile(`hooks\.slack\.com`),
		regexp.MustCompile(`api\.slack\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Slack integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Slack is declared but integration not found",
		Suggestions: []string{
			"Add SLACK_WEBHOOK_URL or SLACK_TOKEN to environment",
		},
	}, nil
}

// DiscordCheck verifies Discord is properly set up
type DiscordCheck struct{}

func (c DiscordCheck) ID() string {
	return "discord"
}

func (c DiscordCheck) Title() string {
	return "Discord"
}

func (c DiscordCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["discord"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Discord not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "DISCORD_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Discord configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`discord\.js`),
		regexp.MustCompile(`discord\.py`),
		regexp.MustCompile(`discordrb`),
		regexp.MustCompile(`discord\.com/api`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Discord SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Discord is declared but SDK not found",
		Suggestions: []string{
			"Add DISCORD_TOKEN or DISCORD_WEBHOOK_URL to environment",
		},
	}, nil
}

// IntercomCheck verifies Intercom is properly set up
type IntercomCheck struct{}

func (c IntercomCheck) ID() string {
	return "intercom"
}

func (c IntercomCheck) Title() string {
	return "Intercom"
}

func (c IntercomCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["intercom"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Intercom not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "INTERCOM_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Intercom configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`widget\.intercom\.io`),
		regexp.MustCompile(`Intercom\(`),
		regexp.MustCompile(`intercomSettings`),
		regexp.MustCompile(`@intercom/`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Intercom widget found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Intercom is declared but widget not found",
		Suggestions: []string{
			"Add Intercom widget script to your templates",
			"Add INTERCOM_APP_ID to environment",
		},
	}, nil
}

// CrispCheck verifies Crisp is properly set up
type CrispCheck struct{}

func (c CrispCheck) ID() string {
	return "crisp"
}

func (c CrispCheck) Title() string {
	return "Crisp"
}

func (c CrispCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["crisp"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Crisp not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "CRISP_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Crisp configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`client\.crisp\.chat`),
		regexp.MustCompile(`CRISP_WEBSITE_ID`),
		regexp.MustCompile(`\$crisp`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Crisp widget found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Crisp is declared but widget not found",
		Suggestions: []string{
			"Add Crisp chat widget script to your templates",
			"Add CRISP_WEBSITE_ID to environment",
		},
	}, nil
}
