package checks

import (
	"regexp"
)

// OpenAICheck verifies OpenAI is properly set up
type OpenAICheck struct{}

func (c OpenAICheck) ID() string {
	return "openai"
}

func (c OpenAICheck) Title() string {
	return "OpenAI"
}

func (c OpenAICheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["openai"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "OpenAI not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "OPENAI_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "OpenAI API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`new OpenAI\(`),
		regexp.MustCompile(`OpenAI\(\s*\{`),
		regexp.MustCompile(`api\.openai\.com`),
		regexp.MustCompile(`from\s+["']openai["']`),
		regexp.MustCompile(`require\s*\(\s*["']openai["']\)`),
		regexp.MustCompile(`import\s+openai`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "OpenAI SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "OpenAI is declared but SDK not found",
		Suggestions: []string{
			"Add OPENAI_API_KEY to environment",
			"Initialize OpenAI client in your application",
		},
	}, nil
}

// AnthropicCheck verifies Anthropic is properly set up
type AnthropicCheck struct{}

func (c AnthropicCheck) ID() string {
	return "anthropic"
}

func (c AnthropicCheck) Title() string {
	return "Anthropic"
}

func (c AnthropicCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["anthropic"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Anthropic not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "ANTHROPIC_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Anthropic API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@anthropic-ai/sdk`),
		regexp.MustCompile(`new Anthropic\(`),
		regexp.MustCompile(`Anthropic\(\s*\{`),
		regexp.MustCompile(`api\.anthropic\.com`),
		regexp.MustCompile(`from\s+["']@anthropic-ai`),
		regexp.MustCompile(`import\s+anthropic`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Anthropic SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Anthropic is declared but SDK not found",
		Suggestions: []string{
			"Add ANTHROPIC_API_KEY to environment",
			"Initialize Anthropic client in your application",
		},
	}, nil
}

// GoogleAICheck verifies Google AI is properly set up
type GoogleAICheck struct{}

func (c GoogleAICheck) ID() string {
	return "google_ai"
}

func (c GoogleAICheck) Title() string {
	return "Google AI"
}

func (c GoogleAICheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["google_ai"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Google AI not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "GOOGLE_AI_") || hasEnvVar(ctx.RootDir, "GEMINI_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Google AI API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@google/generative-ai`),
		regexp.MustCompile(`generativelanguage\.googleapis\.com`),
		regexp.MustCompile(`GoogleGenerativeAI`),
		regexp.MustCompile(`gemini-pro`),
		regexp.MustCompile(`gemini-1\.5`),
		regexp.MustCompile(`models/gemini`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Google AI SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Google AI is declared but SDK not found",
		Suggestions: []string{
			"Add GOOGLE_AI_API_KEY or GEMINI_API_KEY to environment",
			"Initialize Google AI client in your application",
		},
	}, nil
}

// MistralCheck verifies Mistral is properly set up
type MistralCheck struct{}

func (c MistralCheck) ID() string {
	return "mistral"
}

func (c MistralCheck) Title() string {
	return "Mistral AI"
}

func (c MistralCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["mistral"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Mistral AI not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "MISTRAL_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Mistral AI API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@mistralai/`),
		regexp.MustCompile(`mistralai`),
		regexp.MustCompile(`api\.mistral\.ai`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Mistral AI SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Mistral AI is declared but SDK not found",
		Suggestions: []string{
			"Add MISTRAL_API_KEY to environment",
			"Initialize Mistral client in your application",
		},
	}, nil
}

// CohereCheck verifies Cohere is properly set up
type CohereCheck struct{}

func (c CohereCheck) ID() string {
	return "cohere"
}

func (c CohereCheck) Title() string {
	return "Cohere"
}

func (c CohereCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["cohere"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cohere not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "COHERE_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cohere API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`cohere-ai`),
		regexp.MustCompile(`api\.cohere\.ai`),
		regexp.MustCompile(`cohere\.ai`),
		regexp.MustCompile(`CohereClient`),
		regexp.MustCompile(`from\s+["']cohere["']`),
		regexp.MustCompile(`import\s+cohere`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Cohere SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Cohere is declared but SDK not found",
		Suggestions: []string{
			"Add COHERE_API_KEY to environment",
			"Initialize Cohere client in your application",
		},
	}, nil
}

// ReplicateCheck verifies Replicate is properly set up
type ReplicateCheck struct{}

func (c ReplicateCheck) ID() string {
	return "replicate"
}

func (c ReplicateCheck) Title() string {
	return "Replicate"
}

func (c ReplicateCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["replicate"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Replicate not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "REPLICATE_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Replicate API token found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`api\.replicate\.com`),
		regexp.MustCompile(`replicate\.run\(`),
		regexp.MustCompile(`replicate\.predictions`),
		regexp.MustCompile(`from\s+["']replicate["']`),
		regexp.MustCompile(`import\s+replicate`),
		regexp.MustCompile(`new Replicate\(`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Replicate SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Replicate is declared but SDK not found",
		Suggestions: []string{
			"Add REPLICATE_API_TOKEN to environment",
			"Initialize Replicate client in your application",
		},
	}, nil
}

// HuggingFaceCheck verifies Hugging Face is properly set up
type HuggingFaceCheck struct{}

func (c HuggingFaceCheck) ID() string {
	return "huggingface"
}

func (c HuggingFaceCheck) Title() string {
	return "Hugging Face"
}

func (c HuggingFaceCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["huggingface"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Hugging Face not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "HUGGINGFACE_") || hasEnvVar(ctx.RootDir, "HF_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Hugging Face API token found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@huggingface/`),
		regexp.MustCompile(`huggingface\.co`),
		regexp.MustCompile(`HfInference`),
		regexp.MustCompile(`from\s+["']@huggingface`),
		regexp.MustCompile(`from\s+transformers\s+import`),
		regexp.MustCompile(`import\s+transformers`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Hugging Face SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Hugging Face is declared but SDK not found",
		Suggestions: []string{
			"Add HUGGINGFACE_API_TOKEN or HF_TOKEN to environment",
			"Initialize Hugging Face client in your application",
		},
	}, nil
}

// GrokCheck verifies Grok (xAI) is properly set up
type GrokCheck struct{}

func (c GrokCheck) ID() string {
	return "grok"
}

func (c GrokCheck) Title() string {
	return "Grok (xAI)"
}

func (c GrokCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["grok"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Grok not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "XAI_") || hasEnvVar(ctx.RootDir, "GROK_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Grok API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`api\.x\.ai`),
		regexp.MustCompile(`xai-sdk`),
		regexp.MustCompile(`grok-`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Grok SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Grok is declared but SDK not found",
		Suggestions: []string{
			"Add XAI_API_KEY to environment",
			"Initialize Grok client in your application",
		},
	}, nil
}

// PerplexityCheck verifies Perplexity is properly set up
type PerplexityCheck struct{}

func (c PerplexityCheck) ID() string {
	return "perplexity"
}

func (c PerplexityCheck) Title() string {
	return "Perplexity"
}

func (c PerplexityCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["perplexity"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Perplexity not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "PERPLEXITY_") || hasEnvVar(ctx.RootDir, "PPLX_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Perplexity API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`api\.perplexity\.ai`),
		regexp.MustCompile(`perplexity\.ai`),
		regexp.MustCompile(`PerplexityClient`),
		regexp.MustCompile(`from\s+["']perplexity["']`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Perplexity SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Perplexity is declared but SDK not found",
		Suggestions: []string{
			"Add PERPLEXITY_API_KEY to environment",
			"Initialize Perplexity client in your application",
		},
	}, nil
}

// TogetherAICheck verifies Together AI is properly set up
type TogetherAICheck struct{}

func (c TogetherAICheck) ID() string {
	return "together_ai"
}

func (c TogetherAICheck) Title() string {
	return "Together AI"
}

func (c TogetherAICheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["together_ai"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Together AI not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "TOGETHER_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Together AI API key found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`together-ai`),
		regexp.MustCompile(`api\.together\.xyz`),
		regexp.MustCompile(`together\.ai`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Together AI SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Together AI is declared but SDK not found",
		Suggestions: []string{
			"Add TOGETHER_API_KEY to environment",
			"Initialize Together AI client in your application",
		},
	}, nil
}
