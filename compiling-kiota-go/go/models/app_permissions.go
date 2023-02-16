package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppPermissions the permissions granted to the user-to-server access token.
type AppPermissions struct {
    // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
    actions *AppPermissions_actions
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
    administration *AppPermissions_administration
    // The level of permission to grant the access token for checks on code.
    checks *AppPermissions_checks
    // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
    contents *AppPermissions_contents
    // The level of permission to grant the access token for deployments and deployment statuses.
    deployments *AppPermissions_deployments
    // The level of permission to grant the access token for managing repository environments.
    environments *AppPermissions_environments
    // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
    issues *AppPermissions_issues
    // The level of permission to grant the access token for organization teams and members.
    members *AppPermissions_members
    // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
    metadata *AppPermissions_metadata
    // The level of permission to grant the access token to manage access to an organization.
    organization_administration *AppPermissions_organization_administration
    // The level of permission to grant the access token to view and manage announcement banners for an organization.
    organization_announcement_banners *AppPermissions_organization_announcement_banners
    // The level of permission to grant the access token for custom repository roles management. This property is in beta and is subject to change.
    organization_custom_roles *AppPermissions_organization_custom_roles
    // The level of permission to grant the access token to manage the post-receive hooks for an organization.
    organization_hooks *AppPermissions_organization_hooks
    // The level of permission to grant the access token for organization packages published to GitHub Packages.
    organization_packages *AppPermissions_organization_packages
    // The level of permission to grant the access token for viewing an organization's plan.
    organization_plan *AppPermissions_organization_plan
    // The level of permission to grant the access token to manage organization projects and projects beta (where available).
    organization_projects *AppPermissions_organization_projects
    // The level of permission to grant the access token to manage organization secrets.
    organization_secrets *AppPermissions_organization_secrets
    // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
    organization_self_hosted_runners *AppPermissions_organization_self_hosted_runners
    // The level of permission to grant the access token to view and manage users blocked by the organization.
    organization_user_blocking *AppPermissions_organization_user_blocking
    // The level of permission to grant the access token for packages published to GitHub Packages.
    packages *AppPermissions_packages
    // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
    pages *AppPermissions_pages
    // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
    pull_requests *AppPermissions_pull_requests
    // The level of permission to grant the access token to view and manage announcement banners for a repository.
    repository_announcement_banners *AppPermissions_repository_announcement_banners
    // The level of permission to grant the access token to manage the post-receive hooks for a repository.
    repository_hooks *AppPermissions_repository_hooks
    // The level of permission to grant the access token to manage repository projects, columns, and cards.
    repository_projects *AppPermissions_repository_projects
    // The level of permission to grant the access token to view and manage secret scanning alerts.
    secret_scanning_alerts *AppPermissions_secret_scanning_alerts
    // The level of permission to grant the access token to manage repository secrets.
    secrets *AppPermissions_secrets
    // The level of permission to grant the access token to view and manage security events like code scanning alerts.
    security_events *AppPermissions_security_events
    // The level of permission to grant the access token to manage just a single file.
    single_file *AppPermissions_single_file
    // The level of permission to grant the access token for commit statuses.
    statuses *AppPermissions_statuses
    // The level of permission to grant the access token to manage team discussions and related comments.
    team_discussions *AppPermissions_team_discussions
    // The level of permission to grant the access token to manage Dependabot alerts.
    vulnerability_alerts *AppPermissions_vulnerability_alerts
    // The level of permission to grant the access token to update GitHub Actions workflow files.
    workflows *AppPermissions_workflows
}
// NewAppPermissions instantiates a new appPermissions and sets the default values.
func NewAppPermissions()(*AppPermissions) {
    m := &AppPermissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppPermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppPermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppPermissions(), nil
}
// GetActions gets the actions property value. The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
func (m *AppPermissions) GetActions()(*AppPermissions_actions) {
    return m.actions
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppPermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdministration gets the administration property value. The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
func (m *AppPermissions) GetAdministration()(*AppPermissions_administration) {
    return m.administration
}
// GetChecks gets the checks property value. The level of permission to grant the access token for checks on code.
func (m *AppPermissions) GetChecks()(*AppPermissions_checks) {
    return m.checks
}
// GetContents gets the contents property value. The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
func (m *AppPermissions) GetContents()(*AppPermissions_contents) {
    return m.contents
}
// GetDeployments gets the deployments property value. The level of permission to grant the access token for deployments and deployment statuses.
func (m *AppPermissions) GetDeployments()(*AppPermissions_deployments) {
    return m.deployments
}
// GetEnvironments gets the environments property value. The level of permission to grant the access token for managing repository environments.
func (m *AppPermissions) GetEnvironments()(*AppPermissions_environments) {
    return m.environments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppPermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_actions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActions(val.(*AppPermissions_actions))
        }
        return nil
    }
    res["administration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_administration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdministration(val.(*AppPermissions_administration))
        }
        return nil
    }
    res["checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_checks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecks(val.(*AppPermissions_checks))
        }
        return nil
    }
    res["contents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_contents)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContents(val.(*AppPermissions_contents))
        }
        return nil
    }
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_deployments)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployments(val.(*AppPermissions_deployments))
        }
        return nil
    }
    res["environments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_environments)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironments(val.(*AppPermissions_environments))
        }
        return nil
    }
    res["issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_issues)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssues(val.(*AppPermissions_issues))
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_members)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembers(val.(*AppPermissions_members))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_metadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(*AppPermissions_metadata))
        }
        return nil
    }
    res["organization_administration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_administration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationAdministration(val.(*AppPermissions_organization_administration))
        }
        return nil
    }
    res["organization_announcement_banners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_announcement_banners)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationAnnouncementBanners(val.(*AppPermissions_organization_announcement_banners))
        }
        return nil
    }
    res["organization_custom_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_custom_roles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationCustomRoles(val.(*AppPermissions_organization_custom_roles))
        }
        return nil
    }
    res["organization_hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_hooks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationHooks(val.(*AppPermissions_organization_hooks))
        }
        return nil
    }
    res["organization_packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_packages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPackages(val.(*AppPermissions_organization_packages))
        }
        return nil
    }
    res["organization_plan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_plan)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPlan(val.(*AppPermissions_organization_plan))
        }
        return nil
    }
    res["organization_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_projects)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationProjects(val.(*AppPermissions_organization_projects))
        }
        return nil
    }
    res["organization_secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_secrets)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationSecrets(val.(*AppPermissions_organization_secrets))
        }
        return nil
    }
    res["organization_self_hosted_runners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_self_hosted_runners)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationSelfHostedRunners(val.(*AppPermissions_organization_self_hosted_runners))
        }
        return nil
    }
    res["organization_user_blocking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_user_blocking)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationUserBlocking(val.(*AppPermissions_organization_user_blocking))
        }
        return nil
    }
    res["packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_packages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackages(val.(*AppPermissions_packages))
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_pages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPages(val.(*AppPermissions_pages))
        }
        return nil
    }
    res["pull_requests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_pull_requests)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequests(val.(*AppPermissions_pull_requests))
        }
        return nil
    }
    res["repository_announcement_banners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_repository_announcement_banners)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryAnnouncementBanners(val.(*AppPermissions_repository_announcement_banners))
        }
        return nil
    }
    res["repository_hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_repository_hooks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryHooks(val.(*AppPermissions_repository_hooks))
        }
        return nil
    }
    res["repository_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_repository_projects)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryProjects(val.(*AppPermissions_repository_projects))
        }
        return nil
    }
    res["secret_scanning_alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_secret_scanning_alerts)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningAlerts(val.(*AppPermissions_secret_scanning_alerts))
        }
        return nil
    }
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_secrets)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecrets(val.(*AppPermissions_secrets))
        }
        return nil
    }
    res["security_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_security_events)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEvents(val.(*AppPermissions_security_events))
        }
        return nil
    }
    res["single_file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_single_file)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleFile(val.(*AppPermissions_single_file))
        }
        return nil
    }
    res["statuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_statuses)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatuses(val.(*AppPermissions_statuses))
        }
        return nil
    }
    res["team_discussions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_team_discussions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamDiscussions(val.(*AppPermissions_team_discussions))
        }
        return nil
    }
    res["vulnerability_alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_vulnerability_alerts)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVulnerabilityAlerts(val.(*AppPermissions_vulnerability_alerts))
        }
        return nil
    }
    res["workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_workflows)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflows(val.(*AppPermissions_workflows))
        }
        return nil
    }
    return res
}
// GetIssues gets the issues property value. The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
func (m *AppPermissions) GetIssues()(*AppPermissions_issues) {
    return m.issues
}
// GetMembers gets the members property value. The level of permission to grant the access token for organization teams and members.
func (m *AppPermissions) GetMembers()(*AppPermissions_members) {
    return m.members
}
// GetMetadata gets the metadata property value. The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
func (m *AppPermissions) GetMetadata()(*AppPermissions_metadata) {
    return m.metadata
}
// GetOrganizationAdministration gets the organization_administration property value. The level of permission to grant the access token to manage access to an organization.
func (m *AppPermissions) GetOrganizationAdministration()(*AppPermissions_organization_administration) {
    return m.organization_administration
}
// GetOrganizationAnnouncementBanners gets the organization_announcement_banners property value. The level of permission to grant the access token to view and manage announcement banners for an organization.
func (m *AppPermissions) GetOrganizationAnnouncementBanners()(*AppPermissions_organization_announcement_banners) {
    return m.organization_announcement_banners
}
// GetOrganizationCustomRoles gets the organization_custom_roles property value. The level of permission to grant the access token for custom repository roles management. This property is in beta and is subject to change.
func (m *AppPermissions) GetOrganizationCustomRoles()(*AppPermissions_organization_custom_roles) {
    return m.organization_custom_roles
}
// GetOrganizationHooks gets the organization_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for an organization.
func (m *AppPermissions) GetOrganizationHooks()(*AppPermissions_organization_hooks) {
    return m.organization_hooks
}
// GetOrganizationPackages gets the organization_packages property value. The level of permission to grant the access token for organization packages published to GitHub Packages.
func (m *AppPermissions) GetOrganizationPackages()(*AppPermissions_organization_packages) {
    return m.organization_packages
}
// GetOrganizationPlan gets the organization_plan property value. The level of permission to grant the access token for viewing an organization's plan.
func (m *AppPermissions) GetOrganizationPlan()(*AppPermissions_organization_plan) {
    return m.organization_plan
}
// GetOrganizationProjects gets the organization_projects property value. The level of permission to grant the access token to manage organization projects and projects beta (where available).
func (m *AppPermissions) GetOrganizationProjects()(*AppPermissions_organization_projects) {
    return m.organization_projects
}
// GetOrganizationSecrets gets the organization_secrets property value. The level of permission to grant the access token to manage organization secrets.
func (m *AppPermissions) GetOrganizationSecrets()(*AppPermissions_organization_secrets) {
    return m.organization_secrets
}
// GetOrganizationSelfHostedRunners gets the organization_self_hosted_runners property value. The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
func (m *AppPermissions) GetOrganizationSelfHostedRunners()(*AppPermissions_organization_self_hosted_runners) {
    return m.organization_self_hosted_runners
}
// GetOrganizationUserBlocking gets the organization_user_blocking property value. The level of permission to grant the access token to view and manage users blocked by the organization.
func (m *AppPermissions) GetOrganizationUserBlocking()(*AppPermissions_organization_user_blocking) {
    return m.organization_user_blocking
}
// GetPackages gets the packages property value. The level of permission to grant the access token for packages published to GitHub Packages.
func (m *AppPermissions) GetPackages()(*AppPermissions_packages) {
    return m.packages
}
// GetPages gets the pages property value. The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
func (m *AppPermissions) GetPages()(*AppPermissions_pages) {
    return m.pages
}
// GetPullRequests gets the pull_requests property value. The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
func (m *AppPermissions) GetPullRequests()(*AppPermissions_pull_requests) {
    return m.pull_requests
}
// GetRepositoryAnnouncementBanners gets the repository_announcement_banners property value. The level of permission to grant the access token to view and manage announcement banners for a repository.
func (m *AppPermissions) GetRepositoryAnnouncementBanners()(*AppPermissions_repository_announcement_banners) {
    return m.repository_announcement_banners
}
// GetRepositoryHooks gets the repository_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for a repository.
func (m *AppPermissions) GetRepositoryHooks()(*AppPermissions_repository_hooks) {
    return m.repository_hooks
}
// GetRepositoryProjects gets the repository_projects property value. The level of permission to grant the access token to manage repository projects, columns, and cards.
func (m *AppPermissions) GetRepositoryProjects()(*AppPermissions_repository_projects) {
    return m.repository_projects
}
// GetSecrets gets the secrets property value. The level of permission to grant the access token to manage repository secrets.
func (m *AppPermissions) GetSecrets()(*AppPermissions_secrets) {
    return m.secrets
}
// GetSecretScanningAlerts gets the secret_scanning_alerts property value. The level of permission to grant the access token to view and manage secret scanning alerts.
func (m *AppPermissions) GetSecretScanningAlerts()(*AppPermissions_secret_scanning_alerts) {
    return m.secret_scanning_alerts
}
// GetSecurityEvents gets the security_events property value. The level of permission to grant the access token to view and manage security events like code scanning alerts.
func (m *AppPermissions) GetSecurityEvents()(*AppPermissions_security_events) {
    return m.security_events
}
// GetSingleFile gets the single_file property value. The level of permission to grant the access token to manage just a single file.
func (m *AppPermissions) GetSingleFile()(*AppPermissions_single_file) {
    return m.single_file
}
// GetStatuses gets the statuses property value. The level of permission to grant the access token for commit statuses.
func (m *AppPermissions) GetStatuses()(*AppPermissions_statuses) {
    return m.statuses
}
// GetTeamDiscussions gets the team_discussions property value. The level of permission to grant the access token to manage team discussions and related comments.
func (m *AppPermissions) GetTeamDiscussions()(*AppPermissions_team_discussions) {
    return m.team_discussions
}
// GetVulnerabilityAlerts gets the vulnerability_alerts property value. The level of permission to grant the access token to manage Dependabot alerts.
func (m *AppPermissions) GetVulnerabilityAlerts()(*AppPermissions_vulnerability_alerts) {
    return m.vulnerability_alerts
}
// GetWorkflows gets the workflows property value. The level of permission to grant the access token to update GitHub Actions workflow files.
func (m *AppPermissions) GetWorkflows()(*AppPermissions_workflows) {
    return m.workflows
}
// Serialize serializes information the current object
func (m *AppPermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        cast := (*m.GetActions()).String()
        err := writer.WriteStringValue("actions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdministration() != nil {
        cast := (*m.GetAdministration()).String()
        err := writer.WriteStringValue("administration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChecks() != nil {
        cast := (*m.GetChecks()).String()
        err := writer.WriteStringValue("checks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetContents() != nil {
        cast := (*m.GetContents()).String()
        err := writer.WriteStringValue("contents", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeployments() != nil {
        cast := (*m.GetDeployments()).String()
        err := writer.WriteStringValue("deployments", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnvironments() != nil {
        cast := (*m.GetEnvironments()).String()
        err := writer.WriteStringValue("environments", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIssues() != nil {
        cast := (*m.GetIssues()).String()
        err := writer.WriteStringValue("issues", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := (*m.GetMembers()).String()
        err := writer.WriteStringValue("members", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := (*m.GetMetadata()).String()
        err := writer.WriteStringValue("metadata", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationAdministration() != nil {
        cast := (*m.GetOrganizationAdministration()).String()
        err := writer.WriteStringValue("organization_administration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationAnnouncementBanners() != nil {
        cast := (*m.GetOrganizationAnnouncementBanners()).String()
        err := writer.WriteStringValue("organization_announcement_banners", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationCustomRoles() != nil {
        cast := (*m.GetOrganizationCustomRoles()).String()
        err := writer.WriteStringValue("organization_custom_roles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationHooks() != nil {
        cast := (*m.GetOrganizationHooks()).String()
        err := writer.WriteStringValue("organization_hooks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationPackages() != nil {
        cast := (*m.GetOrganizationPackages()).String()
        err := writer.WriteStringValue("organization_packages", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationPlan() != nil {
        cast := (*m.GetOrganizationPlan()).String()
        err := writer.WriteStringValue("organization_plan", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationProjects() != nil {
        cast := (*m.GetOrganizationProjects()).String()
        err := writer.WriteStringValue("organization_projects", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationSecrets() != nil {
        cast := (*m.GetOrganizationSecrets()).String()
        err := writer.WriteStringValue("organization_secrets", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationSelfHostedRunners() != nil {
        cast := (*m.GetOrganizationSelfHostedRunners()).String()
        err := writer.WriteStringValue("organization_self_hosted_runners", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationUserBlocking() != nil {
        cast := (*m.GetOrganizationUserBlocking()).String()
        err := writer.WriteStringValue("organization_user_blocking", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPackages() != nil {
        cast := (*m.GetPackages()).String()
        err := writer.WriteStringValue("packages", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPages() != nil {
        cast := (*m.GetPages()).String()
        err := writer.WriteStringValue("pages", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPullRequests() != nil {
        cast := (*m.GetPullRequests()).String()
        err := writer.WriteStringValue("pull_requests", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryAnnouncementBanners() != nil {
        cast := (*m.GetRepositoryAnnouncementBanners()).String()
        err := writer.WriteStringValue("repository_announcement_banners", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryHooks() != nil {
        cast := (*m.GetRepositoryHooks()).String()
        err := writer.WriteStringValue("repository_hooks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryProjects() != nil {
        cast := (*m.GetRepositoryProjects()).String()
        err := writer.WriteStringValue("repository_projects", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecrets() != nil {
        cast := (*m.GetSecrets()).String()
        err := writer.WriteStringValue("secrets", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanningAlerts() != nil {
        cast := (*m.GetSecretScanningAlerts()).String()
        err := writer.WriteStringValue("secret_scanning_alerts", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityEvents() != nil {
        cast := (*m.GetSecurityEvents()).String()
        err := writer.WriteStringValue("security_events", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSingleFile() != nil {
        cast := (*m.GetSingleFile()).String()
        err := writer.WriteStringValue("single_file", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatuses() != nil {
        cast := (*m.GetStatuses()).String()
        err := writer.WriteStringValue("statuses", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTeamDiscussions() != nil {
        cast := (*m.GetTeamDiscussions()).String()
        err := writer.WriteStringValue("team_discussions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVulnerabilityAlerts() != nil {
        cast := (*m.GetVulnerabilityAlerts()).String()
        err := writer.WriteStringValue("vulnerability_alerts", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflows() != nil {
        cast := (*m.GetWorkflows()).String()
        err := writer.WriteStringValue("workflows", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
func (m *AppPermissions) SetActions(value *AppPermissions_actions)() {
    m.actions = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppPermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdministration sets the administration property value. The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
func (m *AppPermissions) SetAdministration(value *AppPermissions_administration)() {
    m.administration = value
}
// SetChecks sets the checks property value. The level of permission to grant the access token for checks on code.
func (m *AppPermissions) SetChecks(value *AppPermissions_checks)() {
    m.checks = value
}
// SetContents sets the contents property value. The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
func (m *AppPermissions) SetContents(value *AppPermissions_contents)() {
    m.contents = value
}
// SetDeployments sets the deployments property value. The level of permission to grant the access token for deployments and deployment statuses.
func (m *AppPermissions) SetDeployments(value *AppPermissions_deployments)() {
    m.deployments = value
}
// SetEnvironments sets the environments property value. The level of permission to grant the access token for managing repository environments.
func (m *AppPermissions) SetEnvironments(value *AppPermissions_environments)() {
    m.environments = value
}
// SetIssues sets the issues property value. The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
func (m *AppPermissions) SetIssues(value *AppPermissions_issues)() {
    m.issues = value
}
// SetMembers sets the members property value. The level of permission to grant the access token for organization teams and members.
func (m *AppPermissions) SetMembers(value *AppPermissions_members)() {
    m.members = value
}
// SetMetadata sets the metadata property value. The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
func (m *AppPermissions) SetMetadata(value *AppPermissions_metadata)() {
    m.metadata = value
}
// SetOrganizationAdministration sets the organization_administration property value. The level of permission to grant the access token to manage access to an organization.
func (m *AppPermissions) SetOrganizationAdministration(value *AppPermissions_organization_administration)() {
    m.organization_administration = value
}
// SetOrganizationAnnouncementBanners sets the organization_announcement_banners property value. The level of permission to grant the access token to view and manage announcement banners for an organization.
func (m *AppPermissions) SetOrganizationAnnouncementBanners(value *AppPermissions_organization_announcement_banners)() {
    m.organization_announcement_banners = value
}
// SetOrganizationCustomRoles sets the organization_custom_roles property value. The level of permission to grant the access token for custom repository roles management. This property is in beta and is subject to change.
func (m *AppPermissions) SetOrganizationCustomRoles(value *AppPermissions_organization_custom_roles)() {
    m.organization_custom_roles = value
}
// SetOrganizationHooks sets the organization_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for an organization.
func (m *AppPermissions) SetOrganizationHooks(value *AppPermissions_organization_hooks)() {
    m.organization_hooks = value
}
// SetOrganizationPackages sets the organization_packages property value. The level of permission to grant the access token for organization packages published to GitHub Packages.
func (m *AppPermissions) SetOrganizationPackages(value *AppPermissions_organization_packages)() {
    m.organization_packages = value
}
// SetOrganizationPlan sets the organization_plan property value. The level of permission to grant the access token for viewing an organization's plan.
func (m *AppPermissions) SetOrganizationPlan(value *AppPermissions_organization_plan)() {
    m.organization_plan = value
}
// SetOrganizationProjects sets the organization_projects property value. The level of permission to grant the access token to manage organization projects and projects beta (where available).
func (m *AppPermissions) SetOrganizationProjects(value *AppPermissions_organization_projects)() {
    m.organization_projects = value
}
// SetOrganizationSecrets sets the organization_secrets property value. The level of permission to grant the access token to manage organization secrets.
func (m *AppPermissions) SetOrganizationSecrets(value *AppPermissions_organization_secrets)() {
    m.organization_secrets = value
}
// SetOrganizationSelfHostedRunners sets the organization_self_hosted_runners property value. The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
func (m *AppPermissions) SetOrganizationSelfHostedRunners(value *AppPermissions_organization_self_hosted_runners)() {
    m.organization_self_hosted_runners = value
}
// SetOrganizationUserBlocking sets the organization_user_blocking property value. The level of permission to grant the access token to view and manage users blocked by the organization.
func (m *AppPermissions) SetOrganizationUserBlocking(value *AppPermissions_organization_user_blocking)() {
    m.organization_user_blocking = value
}
// SetPackages sets the packages property value. The level of permission to grant the access token for packages published to GitHub Packages.
func (m *AppPermissions) SetPackages(value *AppPermissions_packages)() {
    m.packages = value
}
// SetPages sets the pages property value. The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
func (m *AppPermissions) SetPages(value *AppPermissions_pages)() {
    m.pages = value
}
// SetPullRequests sets the pull_requests property value. The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
func (m *AppPermissions) SetPullRequests(value *AppPermissions_pull_requests)() {
    m.pull_requests = value
}
// SetRepositoryAnnouncementBanners sets the repository_announcement_banners property value. The level of permission to grant the access token to view and manage announcement banners for a repository.
func (m *AppPermissions) SetRepositoryAnnouncementBanners(value *AppPermissions_repository_announcement_banners)() {
    m.repository_announcement_banners = value
}
// SetRepositoryHooks sets the repository_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for a repository.
func (m *AppPermissions) SetRepositoryHooks(value *AppPermissions_repository_hooks)() {
    m.repository_hooks = value
}
// SetRepositoryProjects sets the repository_projects property value. The level of permission to grant the access token to manage repository projects, columns, and cards.
func (m *AppPermissions) SetRepositoryProjects(value *AppPermissions_repository_projects)() {
    m.repository_projects = value
}
// SetSecrets sets the secrets property value. The level of permission to grant the access token to manage repository secrets.
func (m *AppPermissions) SetSecrets(value *AppPermissions_secrets)() {
    m.secrets = value
}
// SetSecretScanningAlerts sets the secret_scanning_alerts property value. The level of permission to grant the access token to view and manage secret scanning alerts.
func (m *AppPermissions) SetSecretScanningAlerts(value *AppPermissions_secret_scanning_alerts)() {
    m.secret_scanning_alerts = value
}
// SetSecurityEvents sets the security_events property value. The level of permission to grant the access token to view and manage security events like code scanning alerts.
func (m *AppPermissions) SetSecurityEvents(value *AppPermissions_security_events)() {
    m.security_events = value
}
// SetSingleFile sets the single_file property value. The level of permission to grant the access token to manage just a single file.
func (m *AppPermissions) SetSingleFile(value *AppPermissions_single_file)() {
    m.single_file = value
}
// SetStatuses sets the statuses property value. The level of permission to grant the access token for commit statuses.
func (m *AppPermissions) SetStatuses(value *AppPermissions_statuses)() {
    m.statuses = value
}
// SetTeamDiscussions sets the team_discussions property value. The level of permission to grant the access token to manage team discussions and related comments.
func (m *AppPermissions) SetTeamDiscussions(value *AppPermissions_team_discussions)() {
    m.team_discussions = value
}
// SetVulnerabilityAlerts sets the vulnerability_alerts property value. The level of permission to grant the access token to manage Dependabot alerts.
func (m *AppPermissions) SetVulnerabilityAlerts(value *AppPermissions_vulnerability_alerts)() {
    m.vulnerability_alerts = value
}
// SetWorkflows sets the workflows property value. The level of permission to grant the access token to update GitHub Actions workflow files.
func (m *AppPermissions) SetWorkflows(value *AppPermissions_workflows)() {
    m.workflows = value
}
