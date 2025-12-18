package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var ignoreCmd = &cobra.Command{
	Use:   "ignore <check-id>",
	Short: "Add a check to the ignore list",
	Long: `Add a check ID to the ignore list in preflight.yml.
The check will be skipped in future scans.

Example:
  preflight ignore sitemap
  preflight ignore llmsTxt
  preflight ignore debug_statements`,
	Args: cobra.ExactArgs(1),
	RunE: runIgnore,
}

func init() {
	rootCmd.AddCommand(ignoreCmd)
}

func runIgnore(cmd *cobra.Command, args []string) error {
	checkID := args[0]

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	configPath := filepath.Join(cwd, "preflight.yml")

	// Read existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("preflight.yml not found. Run 'preflight init' first")
		}
		return fmt.Errorf("failed to read config: %w", err)
	}

	// Parse as generic map to preserve structure
	var cfg map[string]interface{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse preflight.yml: %w", err)
	}

	// Get or create ignore list
	var ignoreList []string
	if existing, ok := cfg["ignore"]; ok {
		if list, ok := existing.([]interface{}); ok {
			for _, item := range list {
				if s, ok := item.(string); ok {
					ignoreList = append(ignoreList, s)
				}
			}
		}
	}

	// Check if already ignored
	for _, id := range ignoreList {
		if id == checkID {
			fmt.Printf("'%s' is already in the ignore list\n", checkID)
			return nil
		}
	}

	// Add to ignore list
	ignoreList = append(ignoreList, checkID)
	cfg["ignore"] = ignoreList

	// Write back
	newData, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to serialize config: %w", err)
	}

	if err := os.WriteFile(configPath, newData, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	fmt.Printf("Added '%s' to ignore list\n", checkID)
	return nil
}

// Also add an unignore command
var unignoreCmd = &cobra.Command{
	Use:   "unignore <check-id>",
	Short: "Remove a check from the ignore list",
	Long: `Remove a check ID from the ignore list in preflight.yml.

Example:
  preflight unignore sitemap`,
	Args: cobra.ExactArgs(1),
	RunE: runUnignore,
}

func init() {
	rootCmd.AddCommand(unignoreCmd)
}

func runUnignore(cmd *cobra.Command, args []string) error {
	checkID := args[0]

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	configPath := filepath.Join(cwd, "preflight.yml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("preflight.yml not found. Run 'preflight init' first")
		}
		return fmt.Errorf("failed to read config: %w", err)
	}

	var cfg map[string]interface{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse preflight.yml: %w", err)
	}

	// Get ignore list
	var ignoreList []string
	if existing, ok := cfg["ignore"]; ok {
		if list, ok := existing.([]interface{}); ok {
			for _, item := range list {
				if s, ok := item.(string); ok {
					ignoreList = append(ignoreList, s)
				}
			}
		}
	}

	// Find and remove
	found := false
	var newList []string
	for _, id := range ignoreList {
		if id == checkID {
			found = true
		} else {
			newList = append(newList, id)
		}
	}

	if !found {
		fmt.Printf("'%s' is not in the ignore list\n", checkID)
		return nil
	}

	// Update or remove ignore key
	if len(newList) > 0 {
		cfg["ignore"] = newList
	} else {
		delete(cfg, "ignore")
	}

	newData, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to serialize config: %w", err)
	}

	if err := os.WriteFile(configPath, newData, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	fmt.Printf("Removed '%s' from ignore list\n", checkID)
	return nil
}

// Helper to list available check IDs
var listChecksCmd = &cobra.Command{
	Use:   "checks",
	Short: "List all available check and service IDs that can be ignored",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("=== Checks ===")
		fmt.Println()

		fmt.Println("SEO & Social:")
		fmt.Println("  - seoMeta")
		fmt.Println("  - canonical")
		fmt.Println("  - structured_data")
		fmt.Println("  - indexNow (opt-in)")
		fmt.Println("  - ogTwitter")
		fmt.Println("  - viewport")
		fmt.Println("  - lang")
		fmt.Println()

		fmt.Println("Security & Infrastructure:")
		fmt.Println("  - securityHeaders")
		fmt.Println("  - ssl")
		fmt.Println("  - www_redirect")
		fmt.Println("  - email_auth (opt-in)")
		fmt.Println("  - secrets")
		fmt.Println()

		fmt.Println("Environment & Health:")
		fmt.Println("  - envParity")
		fmt.Println("  - healthEndpoint")
		fmt.Println()

		fmt.Println("Code Quality & Performance:")
		fmt.Println("  - vulnerability")
		fmt.Println("  - debug_statements")
		fmt.Println("  - error_pages")
		fmt.Println("  - image_optimization")
		fmt.Println()

		fmt.Println("Legal & Compliance:")
		fmt.Println("  - legal_pages")
		fmt.Println()

		fmt.Println("Web Standard Files:")
		fmt.Println("  - favicon")
		fmt.Println("  - robotsTxt")
		fmt.Println("  - sitemap")
		fmt.Println("  - llmsTxt")
		fmt.Println("  - adsTxt (opt-in)")
		fmt.Println("  - humansTxt (opt-in)")
		fmt.Println("  - license (opt-in)")
		fmt.Println()

		fmt.Println("=== Services (with validation checks) ===")
		fmt.Println()
		fmt.Println("These services have checks that verify proper integration:")
		fmt.Println()

		fmt.Println("Payments:")
		fmt.Println("  - stripe: Verifies API keys, webhook secret, SDK initialization")
		fmt.Println()

		fmt.Println("Error Tracking & Monitoring:")
		fmt.Println("  - sentry: Verifies Sentry.init() in application code")
		fmt.Println("  - bugsnag: Verifies Bugsnag.start() initialization")
		fmt.Println("  - rollbar: Verifies Rollbar.init() initialization")
		fmt.Println("  - honeybadger: Verifies Honeybadger.configure() initialization")
		fmt.Println("  - datadog: Verifies Datadog RUM or APM initialization")
		fmt.Println("  - newrelic: Verifies New Relic browser agent or APM")
		fmt.Println("  - logrocket: Verifies LogRocket.init() initialization")
		fmt.Println()

		fmt.Println("Email:")
		fmt.Println("  - postmark: Verifies API key in env or SDK initialization")
		fmt.Println("  - sendgrid: Verifies API key in env or SDK initialization")
		fmt.Println("  - mailgun: Verifies API key in env or SDK initialization")
		fmt.Println("  - aws_ses: Verifies SES configuration or SDK initialization")
		fmt.Println("  - resend: Verifies API key in env or SDK initialization")
		fmt.Println()

		fmt.Println("Analytics:")
		fmt.Println("  - plausible: Verifies Plausible script tag in templates")
		fmt.Println("  - fathom: Verifies Fathom script tag in templates")
		fmt.Println("  - google_analytics: Verifies GA/GTM script in templates")
		fmt.Println("  - fullres: Verifies Fullres script in templates")
		fmt.Println("  - datafast: Verifies Datafa.st script in templates")
		fmt.Println("  - posthog: Verifies posthog.init() initialization")
		fmt.Println("  - mixpanel: Verifies mixpanel.init() initialization")
		fmt.Println("  - amplitude: Verifies amplitude.init() initialization")
		fmt.Println("  - segment: Verifies analytics.load() initialization")
		fmt.Println("  - hotjar: Verifies Hotjar tracking code in templates")
		fmt.Println()

		fmt.Println("Infrastructure:")
		fmt.Println("  - redis: Verifies Redis connection configuration")
		fmt.Println("  - sidekiq: Verifies Sidekiq configuration files")
		fmt.Println()

		fmt.Println("=== Services (detection only, no validation check yet) ===")
		fmt.Println()

		fmt.Println("Payments:")
		fmt.Println("  - paypal, braintree, paddle, lemonsqueezy")
		fmt.Println()

		fmt.Println("Email:")
		fmt.Println("  - mailchimp, convertkit, beehiiv, aweber, activecampaign,")
		fmt.Println("    campaignmonitor, drip, klaviyo, buttondown")
		fmt.Println()

		fmt.Println("Auth:")
		fmt.Println("  - auth0, clerk, workos, firebase, supabase")
		fmt.Println()

		fmt.Println("Communication:")
		fmt.Println("  - twilio, slack, discord, intercom, crisp")
		fmt.Println()

		fmt.Println("Infrastructure:")
		fmt.Println("  - rabbitmq, elasticsearch, convex")
		fmt.Println()

		fmt.Println("Storage & CDN:")
		fmt.Println("  - aws_s3, cloudinary, cloudflare")
		fmt.Println()

		fmt.Println("Search:")
		fmt.Println("  - algolia")
		fmt.Println()

		fmt.Println("AI:")
		fmt.Println("  - openai, anthropic, google_ai, mistral, cohere, replicate,")
		fmt.Println("    huggingface, grok, perplexity, together_ai")
		fmt.Println()

		fmt.Println("SEO:")
		fmt.Println("  - indexnow")
		fmt.Println()

		fmt.Println("Cookie Consent:")
		fmt.Println("  - cookieconsent, cookiebot, onetrust, termly, cookieyes, iubenda")
		fmt.Println()

		fmt.Println("Use 'preflight ignore <id>' to silence a check or service")
		fmt.Println("Use 'preflight unignore <id>' to re-enable it")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listChecksCmd)
}
