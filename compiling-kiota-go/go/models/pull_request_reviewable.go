package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestReviewable 
type PullRequestReviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorAssociation()(*AuthorAssociation)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetCommitId()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLinks()(PullRequestReview__linksable)
    GetNodeId()(*string)
    GetPullRequestUrl()(*string)
    GetState()(*string)
    GetSubmittedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUser()(NullableSimpleUserable)
    SetAuthorAssociation(value *AuthorAssociation)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetCommitId(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLinks(value PullRequestReview__linksable)()
    SetNodeId(value *string)()
    SetPullRequestUrl(value *string)()
    SetState(value *string)()
    SetSubmittedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUser(value NullableSimpleUserable)()
}
