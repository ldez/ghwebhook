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
		// https://github.com/google/go-github/blob/b642d879f3a562c418066129038d0e101929fec7/github/messages.go#L45-L102
		"Events": []string{
			"BranchProtectionRuleEvent",
			"CheckRunEvent",
			"CheckSuiteEvent",
			"CommitCommentEvent",
			"ContentReferenceEvent",
			"CreateEvent",
			"DeleteEvent",
			"DeployKeyEvent",
			"DeploymentEvent",
			"DeploymentStatusEvent",
			"DiscussionEvent",
			"ForkEvent",
			"GitHubAppAuthorizationEvent",
			"GollumEvent",
			"InstallationEvent",
			"InstallationRepositoriesEvent",
			"IssueCommentEvent",
			"IssuesEvent",
			"LabelEvent",
			"MarketplacePurchaseEvent",
			"MemberEvent",
			"MembershipEvent",
			"MetaEvent",
			"MilestoneEvent",
			"OrganizationEvent",
			"OrgBlockEvent",
			"PackageEvent",
			"PageBuildEvent",
			"PingEvent",
			"ProjectEvent",
			"ProjectCardEvent",
			"ProjectColumnEvent",
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
			"RepositoryVulnerabilityAlertEvent",
			"ReleaseEvent",
			"SecretScanningAlertEvent",
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
