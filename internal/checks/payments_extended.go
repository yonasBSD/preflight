package checks

import (
	"regexp"
)

// PayPalCheck verifies PayPal is properly set up
type PayPalCheck struct{}

func (c PayPalCheck) ID() string {
	return "paypal"
}

func (c PayPalCheck) Title() string {
	return "PayPal"
}

func (c PayPalCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["paypal"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "PayPal not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "PAYPAL_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "PayPal configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`paypal\.com/sdk`),
		regexp.MustCompile(`@paypal/`),
		regexp.MustCompile(`paypal-js`),
		regexp.MustCompile(`PayPalButtons`),
		regexp.MustCompile(`paypalobjects\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "PayPal SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "PayPal is declared but SDK not found",
		Suggestions: []string{
			"Add PayPal SDK script or @paypal/react-paypal-js",
			"Configure PAYPAL_CLIENT_ID in environment",
		},
	}, nil
}

// BraintreeCheck verifies Braintree is properly set up
type BraintreeCheck struct{}

func (c BraintreeCheck) ID() string {
	return "braintree"
}

func (c BraintreeCheck) Title() string {
	return "Braintree"
}

func (c BraintreeCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["braintree"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Braintree not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "BRAINTREE_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Braintree configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`braintree\.BraintreeGateway`),
		regexp.MustCompile(`Braintree\\Gateway`),
		regexp.MustCompile(`Braintree::`),
		regexp.MustCompile(`braintreepayments`),
		regexp.MustCompile(`braintree-web`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Braintree SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Braintree is declared but SDK not found",
		Suggestions: []string{
			"Initialize Braintree gateway in your application",
			"Configure BRAINTREE_MERCHANT_ID, BRAINTREE_PUBLIC_KEY, BRAINTREE_PRIVATE_KEY",
		},
	}, nil
}

// PaddleCheck verifies Paddle is properly set up
type PaddleCheck struct{}

func (c PaddleCheck) ID() string {
	return "paddle"
}

func (c PaddleCheck) Title() string {
	return "Paddle"
}

func (c PaddleCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["paddle"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Paddle not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "PADDLE_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Paddle configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`cdn\.paddle\.com`),
		regexp.MustCompile(`Paddle\.Setup`),
		regexp.MustCompile(`Paddle\.Checkout`),
		regexp.MustCompile(`@paddle/paddle-js`),
		regexp.MustCompile(`paddle-node`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Paddle SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Paddle is declared but SDK not found",
		Suggestions: []string{
			"Add Paddle.js script to your checkout page",
			"Configure PADDLE_VENDOR_ID in environment",
		},
	}, nil
}

// LemonSqueezyCheck verifies LemonSqueezy is properly set up
type LemonSqueezyCheck struct{}

func (c LemonSqueezyCheck) ID() string {
	return "lemonsqueezy"
}

func (c LemonSqueezyCheck) Title() string {
	return "LemonSqueezy"
}

func (c LemonSqueezyCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["lemonsqueezy"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "LemonSqueezy not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "LEMONSQUEEZY_") || hasEnvVar(ctx.RootDir, "LEMON_SQUEEZY_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "LemonSqueezy configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@lemonsqueezy/`),
		regexp.MustCompile(`lemonsqueezy\.com`),
		regexp.MustCompile(`LemonSqueezy`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "LemonSqueezy SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "LemonSqueezy is declared but SDK not found",
		Suggestions: []string{
			"Add @lemonsqueezy/lemonsqueezy.js to your project",
			"Configure LEMONSQUEEZY_API_KEY in environment",
		},
	}, nil
}
