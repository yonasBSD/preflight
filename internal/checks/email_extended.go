package checks

import (
	"regexp"
)

// MailchimpCheck verifies Mailchimp is properly set up
type MailchimpCheck struct{}

func (c MailchimpCheck) ID() string {
	return "mailchimp"
}

func (c MailchimpCheck) Title() string {
	return "Mailchimp"
}

func (c MailchimpCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["mailchimp"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Mailchimp not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "MAILCHIMP_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Mailchimp API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@mailchimp/`),
		regexp.MustCompile(`mailchimp\.com`),
		regexp.MustCompile(`list-manage\.com`),
		regexp.MustCompile(`mc4wp`),
		regexp.MustCompile(`mailchimp-for-wp`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Mailchimp integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Mailchimp is declared but integration not found",
		Suggestions: []string{
			"Add MAILCHIMP_API_KEY to your environment",
			"Install @mailchimp/mailchimp_marketing SDK",
		},
	}, nil
}

// ConvertKitCheck verifies ConvertKit/Kit is properly set up
type ConvertKitCheck struct{}

func (c ConvertKitCheck) ID() string {
	return "convertkit"
}

func (c ConvertKitCheck) Title() string {
	return "Kit (ConvertKit)"
}

func (c ConvertKitCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["convertkit"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Kit not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "CONVERTKIT_") || hasEnvVar(ctx.RootDir, "KIT_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Kit API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`convertkit\.com`),
		regexp.MustCompile(`app\.kit\.com`),
		regexp.MustCompile(`@convertkit/`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Kit integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Kit is declared but integration not found",
		Suggestions: []string{
			"Add CONVERTKIT_API_KEY to your environment",
			"Add Kit form embed code to your templates",
		},
	}, nil
}

// BeehiivCheck verifies Beehiiv is properly set up
type BeehiivCheck struct{}

func (c BeehiivCheck) ID() string {
	return "beehiiv"
}

func (c BeehiivCheck) Title() string {
	return "Beehiiv"
}

func (c BeehiivCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["beehiiv"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Beehiiv not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "BEEHIIV_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Beehiiv API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`beehiiv\.com`),
		regexp.MustCompile(`embeds\.beehiiv\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Beehiiv integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Beehiiv is declared but integration not found",
		Suggestions: []string{
			"Add BEEHIIV_API_KEY to your environment",
			"Add Beehiiv embed code to your templates",
		},
	}, nil
}

// AWeberCheck verifies AWeber is properly set up
type AWeberCheck struct{}

func (c AWeberCheck) ID() string {
	return "aweber"
}

func (c AWeberCheck) Title() string {
	return "AWeber"
}

func (c AWeberCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["aweber"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "AWeber not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "AWEBER_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "AWeber configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`aweber\.com`),
		regexp.MustCompile(`forms\.aweber\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "AWeber integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "AWeber is declared but integration not found",
		Suggestions: []string{
			"Add AWeber form embed code to your templates",
		},
	}, nil
}

// ActiveCampaignCheck verifies ActiveCampaign is properly set up
type ActiveCampaignCheck struct{}

func (c ActiveCampaignCheck) ID() string {
	return "activecampaign"
}

func (c ActiveCampaignCheck) Title() string {
	return "ActiveCampaign"
}

func (c ActiveCampaignCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["activecampaign"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "ActiveCampaign not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "ACTIVECAMPAIGN_") || hasEnvVar(ctx.RootDir, "AC_API") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "ActiveCampaign configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`activecampaign\.com`),
		regexp.MustCompile(`trackcmp\.net`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "ActiveCampaign integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "ActiveCampaign is declared but integration not found",
		Suggestions: []string{
			"Add ACTIVECAMPAIGN_API_KEY and ACTIVECAMPAIGN_URL to environment",
		},
	}, nil
}

// CampaignMonitorCheck verifies Campaign Monitor is properly set up
type CampaignMonitorCheck struct{}

func (c CampaignMonitorCheck) ID() string {
	return "campaignmonitor"
}

func (c CampaignMonitorCheck) Title() string {
	return "Campaign Monitor"
}

func (c CampaignMonitorCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["campaignmonitor"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Campaign Monitor not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "CAMPAIGNMONITOR_") || hasEnvVar(ctx.RootDir, "CREATESEND_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Campaign Monitor configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`campaignmonitor\.com`),
		regexp.MustCompile(`createsend\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Campaign Monitor integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Campaign Monitor is declared but integration not found",
		Suggestions: []string{
			"Add Campaign Monitor API key to environment",
		},
	}, nil
}

// DripCheck verifies Drip is properly set up
type DripCheck struct{}

func (c DripCheck) ID() string {
	return "drip"
}

func (c DripCheck) Title() string {
	return "Drip"
}

func (c DripCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["drip"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Drip not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "DRIP_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Drip configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`getdrip\.com`),
		regexp.MustCompile(`tag\.getdrip\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Drip integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Drip is declared but integration not found",
		Suggestions: []string{
			"Add Drip tracking script to your templates",
			"Add DRIP_API_KEY to environment",
		},
	}, nil
}

// KlaviyoCheck verifies Klaviyo is properly set up
type KlaviyoCheck struct{}

func (c KlaviyoCheck) ID() string {
	return "klaviyo"
}

func (c KlaviyoCheck) Title() string {
	return "Klaviyo"
}

func (c KlaviyoCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["klaviyo"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Klaviyo not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "KLAVIYO_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Klaviyo configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`klaviyo\.com`),
		regexp.MustCompile(`static\.klaviyo\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Klaviyo integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Klaviyo is declared but integration not found",
		Suggestions: []string{
			"Add Klaviyo tracking script to your templates",
			"Add KLAVIYO_API_KEY to environment",
		},
	}, nil
}

// ButtondownCheck verifies Buttondown is properly set up
type ButtondownCheck struct{}

func (c ButtondownCheck) ID() string {
	return "buttondown"
}

func (c ButtondownCheck) Title() string {
	return "Buttondown"
}

func (c ButtondownCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["buttondown"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Buttondown not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "BUTTONDOWN_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Buttondown configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`buttondown\.email`),
		regexp.MustCompile(`buttondown\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Buttondown integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Buttondown is declared but integration not found",
		Suggestions: []string{
			"Add Buttondown form embed to your templates",
			"Add BUTTONDOWN_API_KEY to environment",
		},
	}, nil
}
