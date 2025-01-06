package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	model := map[string]interface{}{
		// https://github.com/google/go-github/blob/cd880c603cb76925a9bcf3325f2877902ab0a0eb/github/messages.go#L49-L117
		"Events": []string{
			"BranchProtectionConfigurationEvent",
			"BranchProtectionRuleEvent",
			"CheckRunEvent",
			"CheckSuiteEvent",
			"CodeScanningAlertEvent",
			"CommitCommentEvent",
			"ContentReferenceEvent",
			"CreateEvent",
			"CustomPropertyEvent",
			"CustomPropertyValuesEvent",
			"DeleteEvent",
			"DependabotAlertEvent",
			"DeployKeyEvent",
			"DeploymentEvent",
			"DeploymentReviewEvent",
			"DeploymentStatusEvent",
			"DeploymentProtectionRuleEvent",
			"DiscussionEvent",
			"DiscussionCommentEvent",
			"ForkEvent",
			"GitHubAppAuthorizationEvent",
			"GollumEvent",
			"InstallationEvent",
			"InstallationRepositoriesEvent",
			"InstallationTargetEvent",
			"IssueCommentEvent",
			"IssuesEvent",
			"LabelEvent",
			"MarketplacePurchaseEvent",
			"MemberEvent",
			"MembershipEvent",
			"MergeGroupEvent",
			"MetaEvent",
			"MilestoneEvent",
			"OrganizationEvent",
			"OrgBlockEvent",
			"PackageEvent",
			"PageBuildEvent",
			"PersonalAccessTokenRequestEvent",
			"PingEvent",
			"ProjectV2Event",
			"ProjectV2ItemEvent",
			"PublicEvent",
			"PullRequestEvent",
			"PullRequestReviewEvent",
			"PullRequestReviewCommentEvent",
			"PullRequestReviewThreadEvent",
			"PullRequestTargetEvent",
			"PushEvent",
			"RepositoryEvent",
			"RepositoryDispatchEvent",
			"RepositoryImportEvent",
			"RepositoryRulesetEvent",
			"RepositoryVulnerabilityAlertEvent",
			"ReleaseEvent",
			"SecretScanningAlertEvent",
			"SecretScanningAlertLocationEvent",
			"SecurityAdvisoryEvent",
			"SecurityAndAnalysisEvent",
			"SponsorshipEvent",
			"StarEvent",
			"StatusEvent",
			"TeamEvent",
			"TeamAddEvent",
			"UserEvent",
			"WatchEvent",
			"WorkflowDispatchEvent",
			"WorkflowJobEvent",
			"WorkflowRunEvent",
		},
	}

	err := genFile("./internal/webhook.go.tmpl", "./webhook_gen.go", model)
	if err != nil {
		log.Fatal(err)
	}

	err = genFile("./internal/event_handlers.go.tmpl", "./event_handlers_gen.go", model)
	if err != nil {
		log.Fatal(err)
	}
}

func genFile(tmplFilename, outputFilename string, model interface{}) error {
	base := template.New(filepath.Base(tmplFilename))
	tmpl := template.Must(base.ParseFiles(tmplFilename))

	b := &bytes.Buffer{}
	err := tmpl.Execute(b, model)
	if err != nil {
		return err
	}

	// gofmt
	source, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	return os.WriteFile(outputFilename, source, 0o644)
}
