package checks

import (
	"regexp"
)

// RabbitMQCheck verifies RabbitMQ is properly set up
type RabbitMQCheck struct{}

func (c RabbitMQCheck) ID() string {
	return "rabbitmq"
}

func (c RabbitMQCheck) Title() string {
	return "RabbitMQ"
}

func (c RabbitMQCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["rabbitmq"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "RabbitMQ not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "RABBITMQ_") || hasEnvVar(ctx.RootDir, "AMQP_") || hasEnvVar(ctx.RootDir, "CLOUDAMQP_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "RabbitMQ configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`amqp://`),
		regexp.MustCompile(`amqps://`),
		regexp.MustCompile(`amqplib`),
		regexp.MustCompile(`bunny`),
		regexp.MustCompile(`pika`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "RabbitMQ connection found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "RabbitMQ is declared but connection not found",
		Suggestions: []string{
			"Add RABBITMQ_URL or AMQP_URL to environment",
		},
	}, nil
}

// ElasticsearchCheck verifies Elasticsearch is properly set up
type ElasticsearchCheck struct{}

func (c ElasticsearchCheck) ID() string {
	return "elasticsearch"
}

func (c ElasticsearchCheck) Title() string {
	return "Elasticsearch"
}

func (c ElasticsearchCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["elasticsearch"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Elasticsearch not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "ELASTICSEARCH_") || hasEnvVar(ctx.RootDir, "ELASTIC_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Elasticsearch configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@elastic/elasticsearch`),
		regexp.MustCompile(`elasticsearch-py`),
		regexp.MustCompile(`Elasticsearch::Client`),
		regexp.MustCompile(`elastic\.co`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Elasticsearch client found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Elasticsearch is declared but client not found",
		Suggestions: []string{
			"Add ELASTICSEARCH_URL to environment",
			"Initialize Elasticsearch client in your application",
		},
	}, nil
}

// ConvexCheck verifies Convex is properly set up
type ConvexCheck struct{}

func (c ConvexCheck) ID() string {
	return "convex"
}

func (c ConvexCheck) Title() string {
	return "Convex"
}

func (c ConvexCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["convex"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Convex not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "CONVEX_") || hasEnvVar(ctx.RootDir, "NEXT_PUBLIC_CONVEX") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Convex configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`convex/_generated`),
		regexp.MustCompile(`ConvexProvider`),
		regexp.MustCompile(`convex\.dev`),
		regexp.MustCompile(`@convex/`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Convex initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Convex is declared but initialization not found",
		Suggestions: []string{
			"Add CONVEX_URL to environment",
			"Wrap your app with ConvexProvider",
		},
	}, nil
}
