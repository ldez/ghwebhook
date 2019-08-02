package ghwebhook

import "github.com/google/go-github/v27/github"

// ContentReferenceEvent is triggered when the body or comment of an issue or pull request includes a URL that matches a configured content reference domain.
// Only GitHub Apps can receive this event.
//
// GitHub API docs: https://developer.github.com/v3/activity/events/types/#contentreferenceevent
type ContentReferenceEvent struct {
	Action           *string            `json:"action,omitempty"`
	ContentReference *ContentReference  `json:"content_reference,omitempty"`
	Repo             *github.Repository `json:"repository,omitempty"`
	Sender           *github.User       `json:"sender,omitempty"`
	Installation     *github.App        `json:"installation,omitempty"`
}

// GetAction returns the Action field if it's non-nil, zero value otherwise.
func (c *ContentReferenceEvent) GetAction() string {
	if c == nil || c.Action == nil {
		return ""
	}
	return *c.Action
}

// GetContentReference returns the ContentReference field.
func (c *ContentReferenceEvent) GetContentReference() *ContentReference {
	if c == nil {
		return nil
	}
	return c.ContentReference
}

// GetRepo returns the Repo field.
func (c *ContentReferenceEvent) GetRepo() *github.Repository {
	if c == nil {
		return nil
	}
	return c.Repo
}

// GetSender returns the Sender field.
func (c *ContentReferenceEvent) GetSender() *github.User {
	if c == nil {
		return nil
	}
	return c.Sender
}

// GetInstallation returns the Installation field.
func (c *ContentReferenceEvent) GetInstallation() *github.App {
	if c == nil {
		return nil
	}
	return c.Installation
}

// ContentReference represents a GitHub Content Reference.
type ContentReference struct {
	ID        *int64  `json:"id,omitempty"`
	NodeID    *string `json:"node_id,omitempty"`
	Reference *string `json:"reference,omitempty"`
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (c *ContentReference) GetID() int64 {
	if c == nil || c.ID == nil {
		return 0
	}
	return *c.ID
}

// GetNodeID returns the NodeID field if it's non-nil, zero value otherwise.
func (c *ContentReference) GetNodeID() string {
	if c == nil || c.NodeID == nil {
		return ""
	}
	return *c.NodeID
}

// GetReference returns the Reference field if it's non-nil, zero value otherwise.
func (c *ContentReference) GetReference() string {
	if c == nil || c.Reference == nil {
		return ""
	}
	return *c.Reference
}

// RepositoryImportEvent is triggered when a successful, canceled, or failed repository import finishes for a GitHub organization or a personal repository.
//
// GitHub API docs: https://developer.github.com/v3/activity/events/types/#repositoryimportevent
type RepositoryImportEvent struct {
	Status       *string              `json:"status,omitempty"`
	Repo         *github.Repository   `json:"repository,omitempty"`
	Organization *github.Organization `json:"organization,omitempty"`
	Sender       *github.User         `json:"sender,omitempty"`
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (r *RepositoryImportEvent) GetStatus() string {
	if r == nil || r.Status == nil {
		return ""
	}
	return *r.Status
}

// GetRepo returns the Repo field.
func (r *RepositoryImportEvent) GetRepo() *github.Repository {
	if r == nil {
		return nil
	}
	return r.Repo
}

// GetOrganization returns the Organization field.
func (r *RepositoryImportEvent) GetOrganization() *github.Organization {
	if r == nil {
		return nil
	}
	return r.Organization
}

// GetSender returns the Sender field.
func (r *RepositoryImportEvent) GetSender() *github.User {
	if r == nil {
		return nil
	}
	return r.Sender
}

// SecurityAdvisoryEvent is triggered when a new security advisory is published, updated, or withdrawn.
// A security advisory provides information about security-related vulnerabilities in software on GitHub.
//
// GitHub API docs: https://developer.github.com/v3/activity/events/types/#securityadvisoryevent
type SecurityAdvisoryEvent struct {
	Action           *string           `json:"action,omitempty"`
	SecurityAdvisory *SecurityAdvisory `json:"security_advisory,omitempty"`
}

// GetAction returns the Action field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryEvent) GetAction() string {
	if s == nil || s.Action == nil {
		return ""
	}
	return *s.Action
}

// GetSecurityAdvisory returns the SecurityAdvisory field.
func (s *SecurityAdvisoryEvent) GetSecurityAdvisory() *SecurityAdvisory {
	if s == nil {
		return nil
	}
	return s.SecurityAdvisory
}

// SecurityAdvisory represents a GitHub Security Advisory.
type SecurityAdvisory struct {
	GhsaID          *string           `json:"ghsa_id,omitempty"`
	Summary         *string           `json:"summary,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Severity        *string           `json:"severity,omitempty"`
	Identifiers     []Identifiers     `json:"identifiers,omitempty"`
	References      []References      `json:"references,omitempty"`
	PublishedAt     *github.Timestamp `json:"published_at,omitempty"`
	UpdatedAt       *github.Timestamp `json:"updated_at,omitempty"`
	WithdrawnAt     *github.Timestamp `json:"withdrawn_at,omitempty"`
	Vulnerabilities []Vulnerabilities `json:"vulnerabilities,omitempty"`
}

// GetGhsaID returns the GhsaID field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetGhsaID() string {
	if s == nil || s.GhsaID == nil {
		return ""
	}
	return *s.GhsaID
}

// GetSummary returns the Summary field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetSummary() string {
	if s == nil || s.Summary == nil {
		return ""
	}
	return *s.Summary
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetDescription() string {
	if s == nil || s.Description == nil {
		return ""
	}
	return *s.Description
}

// GetSeverity returns the Severity field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetSeverity() string {
	if s == nil || s.Severity == nil {
		return ""
	}
	return *s.Severity
}

// GetPublishedAt returns the PublishedAt field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetPublishedAt() github.Timestamp {
	if s == nil || s.PublishedAt == nil {
		return github.Timestamp{}
	}
	return *s.PublishedAt
}

// GetUpdatedAt returns the UpdatedAt field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetUpdatedAt() github.Timestamp {
	if s == nil || s.UpdatedAt == nil {
		return github.Timestamp{}
	}
	return *s.UpdatedAt
}

// GetWithdrawnAt returns the WithdrawnAt field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetWithdrawnAt() github.Timestamp {
	if s == nil || s.WithdrawnAt == nil {
		return github.Timestamp{}
	}
	return *s.WithdrawnAt
}

// Identifiers represents a GitHub Security Advisory Identifiers.
type Identifiers struct {
	Value *string `json:"value,omitempty"`
	Type  *string `json:"type,omitempty"`
}

// GetValue returns the Value field if it's non-nil, zero value otherwise.
func (i *Identifiers) GetValue() string {
	if i == nil || i.Value == nil {
		return ""
	}
	return *i.Value
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (i *Identifiers) GetType() string {
	if i == nil || i.Type == nil {
		return ""
	}
	return *i.Type
}

// References represents a GitHub Security Advisory References.
type References struct {
	URL *string `json:"url,omitempty"`
}

// GetURL returns the URL field if it's non-nil, zero value otherwise.
func (r *References) GetURL() string {
	if r == nil || r.URL == nil {
		return ""
	}
	return *r.URL
}

// Vulnerabilities represents a GitHub Security Advisory Vulnerabilities.
type Vulnerabilities struct {
	Package                *Package             `json:"package,omitempty"`
	Severity               *string              `json:"severity,omitempty"`
	VulnerableVersionRange *string              `json:"vulnerable_version_range,omitempty"`
	FirstPatchedVersion    *FirstPatchedVersion `json:"first_patched_version,omitempty"`
}

// GetPackage returns the Package field.
func (v *Vulnerabilities) GetPackage() *Package {
	if v == nil {
		return nil
	}
	return v.Package
}

// GetSeverity returns the Severity field if it's non-nil, zero value otherwise.
func (v *Vulnerabilities) GetSeverity() string {
	if v == nil || v.Severity == nil {
		return ""
	}
	return *v.Severity
}

// GetVulnerableVersionRange returns the VulnerableVersionRange field if it's non-nil, zero value otherwise.
func (v *Vulnerabilities) GetVulnerableVersionRange() string {
	if v == nil || v.VulnerableVersionRange == nil {
		return ""
	}
	return *v.VulnerableVersionRange
}

// GetFirstPatchedVersion returns the FirstPatchedVersion field.
func (v *Vulnerabilities) GetFirstPatchedVersion() *FirstPatchedVersion {
	if v == nil {
		return nil
	}
	return v.FirstPatchedVersion
}

// Package represents a GitHub Security Advisory Package.
type Package struct {
	Ecosystem *string `json:"ecosystem,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// GetEcosystem returns the Ecosystem field if it's non-nil, zero value otherwise.
func (p *Package) GetEcosystem() string {
	if p == nil || p.Ecosystem == nil {
		return ""
	}
	return *p.Ecosystem
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (p *Package) GetName() string {
	if p == nil || p.Name == nil {
		return ""
	}
	return *p.Name
}

// FirstPatchedVersion represents a GitHub Security Advisory First Patched Version.
type FirstPatchedVersion struct {
	Identifier *string `json:"identifier,omitempty"`
}

// GetIdentifier returns the Identifier field if it's non-nil, zero value otherwise.
func (f *FirstPatchedVersion) GetIdentifier() string {
	if f == nil || f.Identifier == nil {
		return ""
	}
	return *f.Identifier
}
