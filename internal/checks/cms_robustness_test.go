package checks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/preflightsh/preflight/internal/config"
)

// writeFiles materializes rel->content under a fresh temp dir and returns it.
func writeFiles(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for rel, content := range files {
		full := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", rel, err)
		}
		if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
			t.Fatalf("write %s: %v", rel, err)
		}
	}
	return root
}

const renderedWithViewportAndLang = `<!doctype html>
<html dir="ltr" lang="en-US">
<head><meta name="viewport" content="width=device-width, initial-scale=1.0"></head>
<body></body></html>`

// A Craft layout whose <html lang> and viewport live in an unconventional
// partial (header.twig) the static scanner doesn't know about — so the layout
// itself carries neither tag. Mirrors the joncphillips.com false positive.
const craftLayoutNoTags = `{% extends "_partials/header.twig" %}
{% block content %}<h1>Hi</h1>{% endblock %}`

func TestViewportRenderedHTMLFallback(t *testing.T) {
	root := writeFiles(t, map[string]string{
		"templates/_layout.twig": craftLayoutNoTags,
	})

	t.Run("passes from rendered prod HTML when static scan misses it", func(t *testing.T) {
		ctx := Context{
			RootDir: root,
			Config: &config.PreflightConfig{
				Stack: "craft",
				URLs:  config.URLConfig{Production: "https://prod", Staging: "https://staging"},
			},
			PageHTMLProduction: renderedWithViewportAndLang,
			PageHTMLStaging:    renderedWithViewportAndLang,
		}
		res, _ := ViewportCheck{}.Run(ctx)
		if !res.Passed {
			t.Fatalf("viewport should pass via rendered HTML; got WARN %q", res.Message)
		}
		if !strings.Contains(res.Message, "prod: ✓") {
			t.Fatalf("expected per-env breakdown, got %q", res.Message)
		}
	})

	t.Run("still warns offline when no URL is configured", func(t *testing.T) {
		ctx := Context{
			RootDir: root,
			Config:  &config.PreflightConfig{Stack: "craft"},
		}
		res, _ := ViewportCheck{}.Run(ctx)
		if res.Passed {
			t.Fatal("viewport should warn offline when the tag is in an unscanned partial")
		}
	})
}

func TestLangRenderedHTMLFallback(t *testing.T) {
	root := writeFiles(t, map[string]string{
		"templates/_layout.twig": craftLayoutNoTags,
	})

	t.Run("passes from rendered prod HTML when static scan misses it", func(t *testing.T) {
		ctx := Context{
			RootDir: root,
			Config: &config.PreflightConfig{
				Stack: "craft",
				URLs:  config.URLConfig{Production: "https://prod"},
			},
			PageHTMLProduction: renderedWithViewportAndLang,
		}
		res, _ := LangAttributeCheck{}.Run(ctx)
		if !res.Passed {
			t.Fatalf("lang should pass via rendered HTML; got WARN %q", res.Message)
		}
		if !strings.Contains(res.Message, "prod: ✓") {
			t.Fatalf("expected per-env breakdown, got %q", res.Message)
		}
	})

	t.Run("still warns offline when no URL is configured", func(t *testing.T) {
		ctx := Context{
			RootDir: root,
			Config:  &config.PreflightConfig{Stack: "craft"},
		}
		res, _ := LangAttributeCheck{}.Run(ctx)
		if res.Passed {
			t.Fatal("lang should warn offline when the attribute is in an unscanned partial")
		}
	})
}

func TestVulnerabilitySummaryNamesEcosystem(t *testing.T) {
	out := "Found 79 security vulnerability advisories affecting 16 packages:"
	res, _ := VulnerabilityCheck{}.parseResult(fmt.Errorf("exit status 1"), out, "composer audit")

	if res.Passed {
		t.Fatal("expected WARN for vulnerabilities present")
	}
	if !strings.Contains(strings.ToLower(res.Message), "composer audit") {
		t.Fatalf("WARN message should name the ecosystem/tool; got %q", res.Message)
	}
	if strings.Contains(res.Message, "npm") {
		t.Fatalf("composer findings must not read as npm; got %q", res.Message)
	}
}

func TestHasEnvVarReferenceAcrossPlatforms(t *testing.T) {
	cases := []struct {
		name    string
		files   map[string]string
		prefix  string
		wantHit bool
	}{
		{
			name: "craft project.yaml env reference",
			files: map[string]string{
				"config/project/project.yaml": "mailer:\n  transportSettings:\n    apiKey: $AWS_SES_API_KEY\n    region: $AWS_SES_REGION\n",
			},
			prefix:  "AWS_SES_",
			wantHit: true,
		},
		{
			name: "laravel config env() call",
			files: map[string]string{
				"config/services.php": "<?php return ['mailgun' => ['secret' => env('MAILGUN_SECRET')]];",
			},
			prefix:  "MAILGUN_",
			wantHit: true,
		},
		{
			name: "fly.toml deploy manifest",
			files: map[string]string{
				"fly.toml": "[env]\n  SENDGRID_API_KEY = \"set-in-secrets\"\n",
			},
			prefix:  "SENDGRID_",
			wantHit: true,
		},
		{
			name:    "no reference anywhere",
			files:   map[string]string{"config/app.php": "<?php return [];"},
			prefix:  "RESEND_",
			wantHit: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := writeFiles(t, tc.files)
			where, ok := hasEnvVarReference(root, tc.prefix)
			if ok != tc.wantHit {
				t.Fatalf("hasEnvVarReference = %v (at %q), want %v", ok, where, tc.wantHit)
			}
		})
	}
}

func TestAWSSESPassesOnEnvReference(t *testing.T) {
	root := writeFiles(t, map[string]string{
		"config/project/project.yaml": "mailer:\n  transportType: putyourlightson\\amazonses\\mail\\AmazonSesAdapter\n  transportSettings:\n    apiKey: $AWS_SES_API_KEY\n    apiSecret: $AWS_SES_API_SECRET\n    region: $AWS_SES_REGION\n",
	})

	ctx := Context{
		RootDir: root,
		Config: &config.PreflightConfig{
			Stack:    "craft",
			Services: map[string]config.ServiceConfig{"aws_ses": {Declared: true}},
		},
	}
	res, _ := AWSSESCheck{}.Run(ctx)
	if !res.Passed {
		t.Fatalf("AWS SES should pass when configured via env reference; got WARN %q", res.Message)
	}

	// Negative: declared but no reference, no .env, no SDK code → still warns.
	bare := writeFiles(t, map[string]string{"composer.json": "{}"})
	ctx.RootDir = bare
	bareRes, _ := AWSSESCheck{}.Run(ctx)
	if bareRes.Passed {
		t.Fatal("AWS SES should warn when declared with no config evidence at all")
	}
}
