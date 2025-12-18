package checks

import (
	"regexp"
)

// AWSS3Check verifies AWS S3 is properly set up
type AWSS3Check struct{}

func (c AWSS3Check) ID() string {
	return "aws_s3"
}

func (c AWSS3Check) Title() string {
	return "AWS S3"
}

func (c AWSS3Check) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["aws_s3"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "AWS S3 not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "AWS_") || hasEnvVar(ctx.RootDir, "S3_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "AWS S3 configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@aws-sdk/client-s3`),
		regexp.MustCompile(`aws-sdk.*S3`),
		regexp.MustCompile(`Aws\\S3`),
		regexp.MustCompile(`boto3.*s3`),
		regexp.MustCompile(`s3\.amazonaws\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "AWS S3 SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "AWS S3 is declared but SDK not found",
		Suggestions: []string{
			"Add AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY to environment",
			"Initialize AWS S3 client in your application",
		},
	}, nil
}

// CloudinaryCheck verifies Cloudinary is properly set up
type CloudinaryCheck struct{}

func (c CloudinaryCheck) ID() string {
	return "cloudinary"
}

func (c CloudinaryCheck) Title() string {
	return "Cloudinary"
}

func (c CloudinaryCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["cloudinary"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cloudinary not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "CLOUDINARY_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cloudinary configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`res\.cloudinary\.com`),
		regexp.MustCompile(`@cloudinary/`),
		regexp.MustCompile(`cloudinary\.v2`),
		regexp.MustCompile(`cloudinary\.config`),
		regexp.MustCompile(`cloudinary\.uploader`),
		regexp.MustCompile(`from\s+["']cloudinary["']`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cloudinary SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Cloudinary is declared but SDK not found",
		Suggestions: []string{
			"Add CLOUDINARY_URL or CLOUDINARY_CLOUD_NAME to environment",
			"Initialize Cloudinary SDK in your application",
		},
	}, nil
}

// CloudflareCheck verifies Cloudflare is properly set up
type CloudflareCheck struct{}

func (c CloudflareCheck) ID() string {
	return "cloudflare"
}

func (c CloudflareCheck) Title() string {
	return "Cloudflare"
}

func (c CloudflareCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["cloudflare"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cloudflare not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "CLOUDFLARE_") || hasEnvVar(ctx.RootDir, "CF_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cloudflare configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@cloudflare/`),
		regexp.MustCompile(`cdnjs\.cloudflare\.com`),
		regexp.MustCompile(`api\.cloudflare\.com`),
		regexp.MustCompile(`cloudflare\.com/client`),
		regexp.MustCompile(`wrangler\.toml`),
		regexp.MustCompile(`wrangler deploy`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cloudflare integration found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Cloudflare is declared but integration not found",
		Suggestions: []string{
			"Add CLOUDFLARE_API_TOKEN to environment",
			"Configure Cloudflare Workers or Pages if applicable",
		},
	}, nil
}
