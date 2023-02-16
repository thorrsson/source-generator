package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReviewRequestRemovedIssueEvent review Request Removed Issue Event
type ReviewRequestRemovedIssueEvent struct {
    // A GitHub user.
    actor SimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit_id property
    commit_id *string
    // The commit_url property
    commit_url *string
    // The created_at property
    created_at *string
    // The event property
    event *string
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // A GitHub user.
    requested_reviewer SimpleUserable
    // Groups of organization members that gives permissions on specified repositories.
    requested_team Teamable
    // A GitHub user.
    review_requester SimpleUserable
    // The url property
    url *string
}
// NewReviewRequestRemovedIssueEvent instantiates a new ReviewRequestRemovedIssueEvent and sets the default values.
func NewReviewRequestRemovedIssueEvent()(*ReviewRequestRemovedIssueEvent) {
    m := &ReviewRequestRemovedIssueEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReviewRequestRemovedIssueEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReviewRequestRemovedIssueEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewRequestRemovedIssueEvent(), nil
}
// GetActor gets the actor property value. A GitHub user.
func (m *ReviewRequestRemovedIssueEvent) GetActor()(SimpleUserable) {
    return m.actor
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReviewRequestRemovedIssueEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommitId gets the commit_id property value. The commit_id property
func (m *ReviewRequestRemovedIssueEvent) GetCommitId()(*string) {
    return m.commit_id
}
// GetCommitUrl gets the commit_url property value. The commit_url property
func (m *ReviewRequestRemovedIssueEvent) GetCommitUrl()(*string) {
    return m.commit_url
}
// GetCreatedAt gets the created_at property value. The created_at property
func (m *ReviewRequestRemovedIssueEvent) GetCreatedAt()(*string) {
    return m.created_at
}
// GetEvent gets the event property value. The event property
func (m *ReviewRequestRemovedIssueEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReviewRequestRemovedIssueEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(SimpleUserable))
        }
        return nil
    }
    res["commit_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitId(val)
        }
        return nil
    }
    res["commit_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitUrl(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["performed_via_github_app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformedViaGithubApp(val.(NullableIntegrationable))
        }
        return nil
    }
    res["requested_reviewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedReviewer(val.(SimpleUserable))
        }
        return nil
    }
    res["requested_team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedTeam(val.(Teamable))
        }
        return nil
    }
    res["review_requester"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewRequester(val.(SimpleUserable))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *ReviewRequestRemovedIssueEvent) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
func (m *ReviewRequestRemovedIssueEvent) GetNodeId()(*string) {
    return m.node_id
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *ReviewRequestRemovedIssueEvent) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetRequestedReviewer gets the requested_reviewer property value. A GitHub user.
func (m *ReviewRequestRemovedIssueEvent) GetRequestedReviewer()(SimpleUserable) {
    return m.requested_reviewer
}
// GetRequestedTeam gets the requested_team property value. Groups of organization members that gives permissions on specified repositories.
func (m *ReviewRequestRemovedIssueEvent) GetRequestedTeam()(Teamable) {
    return m.requested_team
}
// GetReviewRequester gets the review_requester property value. A GitHub user.
func (m *ReviewRequestRemovedIssueEvent) GetReviewRequester()(SimpleUserable) {
    return m.review_requester
}
// GetUrl gets the url property value. The url property
func (m *ReviewRequestRemovedIssueEvent) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ReviewRequestRemovedIssueEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_id", m.GetCommitId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_url", m.GetCommitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("performed_via_github_app", m.GetPerformedViaGithubApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requested_reviewer", m.GetRequestedReviewer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requested_team", m.GetRequestedTeam())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("review_requester", m.GetReviewRequester())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
// SetActor sets the actor property value. A GitHub user.
func (m *ReviewRequestRemovedIssueEvent) SetActor(value SimpleUserable)() {
    m.actor = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReviewRequestRemovedIssueEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommitId sets the commit_id property value. The commit_id property
func (m *ReviewRequestRemovedIssueEvent) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetCommitUrl sets the commit_url property value. The commit_url property
func (m *ReviewRequestRemovedIssueEvent) SetCommitUrl(value *string)() {
    m.commit_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *ReviewRequestRemovedIssueEvent) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetEvent sets the event property value. The event property
func (m *ReviewRequestRemovedIssueEvent) SetEvent(value *string)() {
    m.event = value
}
// SetId sets the id property value. The id property
func (m *ReviewRequestRemovedIssueEvent) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *ReviewRequestRemovedIssueEvent) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *ReviewRequestRemovedIssueEvent) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetRequestedReviewer sets the requested_reviewer property value. A GitHub user.
func (m *ReviewRequestRemovedIssueEvent) SetRequestedReviewer(value SimpleUserable)() {
    m.requested_reviewer = value
}
// SetRequestedTeam sets the requested_team property value. Groups of organization members that gives permissions on specified repositories.
func (m *ReviewRequestRemovedIssueEvent) SetRequestedTeam(value Teamable)() {
    m.requested_team = value
}
// SetReviewRequester sets the review_requester property value. A GitHub user.
func (m *ReviewRequestRemovedIssueEvent) SetReviewRequester(value SimpleUserable)() {
    m.review_requester = value
}
// SetUrl sets the url property value. The url property
func (m *ReviewRequestRemovedIssueEvent) SetUrl(value *string)() {
    m.url = value
}
