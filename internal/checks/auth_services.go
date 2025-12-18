package checks

import (
	"regexp"
)

// Auth0Check verifies Auth0 is properly set up
type Auth0Check struct{}

func (c Auth0Check) ID() string {
	return "auth0"
}

func (c Auth0Check) Title() string {
	return "Auth0"
}

func (c Auth0Check) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["auth0"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Auth0 not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "AUTH0_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Auth0 configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@auth0/`),
		regexp.MustCompile(`auth0\.com`),
		regexp.MustCompile(`Auth0Provider`),
		regexp.MustCompile(`createAuth0Client`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Auth0 SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Auth0 is declared but SDK not found",
		Suggestions: []string{
			"Add AUTH0_DOMAIN and AUTH0_CLIENT_ID to environment",
			"Initialize Auth0 SDK in your application",
		},
	}, nil
}

// ClerkCheck verifies Clerk is properly set up
type ClerkCheck struct{}

func (c ClerkCheck) ID() string {
	return "clerk"
}

func (c ClerkCheck) Title() string {
	return "Clerk"
}

func (c ClerkCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["clerk"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Clerk not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "CLERK_") || hasEnvVar(ctx.RootDir, "NEXT_PUBLIC_CLERK") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Clerk configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@clerk/`),
		regexp.MustCompile(`ClerkProvider`),
		regexp.MustCompile(`clerk\.com`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Clerk SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Clerk is declared but SDK not found",
		Suggestions: []string{
			"Add CLERK_SECRET_KEY and NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY",
			"Wrap your app with ClerkProvider",
		},
	}, nil
}

// WorkOSCheck verifies WorkOS is properly set up
type WorkOSCheck struct{}

func (c WorkOSCheck) ID() string {
	return "workos"
}

func (c WorkOSCheck) Title() string {
	return "WorkOS"
}

func (c WorkOSCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["workos"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "WorkOS not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "WORKOS_") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "WorkOS configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@workos-inc/`),
		regexp.MustCompile(`workos\.com`),
		regexp.MustCompile(`WorkOS`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "WorkOS SDK initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "WorkOS is declared but SDK not found",
		Suggestions: []string{
			"Add WORKOS_API_KEY and WORKOS_CLIENT_ID to environment",
		},
	}, nil
}

// FirebaseCheck verifies Firebase is properly set up
type FirebaseCheck struct{}

func (c FirebaseCheck) ID() string {
	return "firebase"
}

func (c FirebaseCheck) Title() string {
	return "Firebase"
}

func (c FirebaseCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["firebase"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Firebase not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "FIREBASE_") || hasEnvVar(ctx.RootDir, "NEXT_PUBLIC_FIREBASE") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Firebase configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`firebase/app`),
		regexp.MustCompile(`from\s+["']firebase`),
		regexp.MustCompile(`@firebase/`),
		regexp.MustCompile(`firebaseConfig`),
		regexp.MustCompile(`firebase\.google\.com`),
		regexp.MustCompile(`firebase\.initializeApp`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Firebase initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Firebase is declared but initialization not found",
		Suggestions: []string{
			"Add Firebase config to your environment",
			"Initialize Firebase with initializeApp()",
		},
	}, nil
}

// SupabaseCheck verifies Supabase is properly set up
type SupabaseCheck struct{}

func (c SupabaseCheck) ID() string {
	return "supabase"
}

func (c SupabaseCheck) Title() string {
	return "Supabase"
}

func (c SupabaseCheck) Run(ctx Context) (CheckResult, error) {
	service, declared := ctx.Config.Services["supabase"]
	if !declared || !service.Declared {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Supabase not declared, skipping",
		}, nil
	}

	if hasEnvVar(ctx.RootDir, "SUPABASE_") || hasEnvVar(ctx.RootDir, "NEXT_PUBLIC_SUPABASE") {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Supabase configuration found in environment",
		}, nil
	}

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`@supabase/`),
		regexp.MustCompile(`supabase\.co`),
		regexp.MustCompile(`supabase\.createClient`),
		regexp.MustCompile(`createClient\s*\([^)]*supabase`),
		regexp.MustCompile(`from\s+["']@supabase`),
	}

	found := searchForPatterns(ctx.RootDir, ctx.Config.Stack, patterns)

	if found {
		return CheckResult{
			ID:       c.ID(),
			Title:    c.Title(),
			Severity: SeverityInfo,
			Passed:   true,
			Message:  "Supabase initialization found",
		}, nil
	}

	return CheckResult{
		ID:       c.ID(),
		Title:    c.Title(),
		Severity: SeverityWarn,
		Passed:   false,
		Message:  "Supabase is declared but initialization not found",
		Suggestions: []string{
			"Add SUPABASE_URL and SUPABASE_ANON_KEY to environment",
			"Initialize Supabase client with createClient()",
		},
	}, nil
}
