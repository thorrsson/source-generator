package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RenamedIssueEventable 
type RenamedIssueEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(SimpleUserable)
    GetCommitId()(*string)
    GetCommitUrl()(*string)
    GetCreatedAt()(*string)
    GetEvent()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetRename()(RenamedIssueEvent_renameable)
    GetUrl()(*string)
    SetActor(value SimpleUserable)()
    SetCommitId(value *string)()
    SetCommitUrl(value *string)()
    SetCreatedAt(value *string)()
    SetEvent(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetRename(value RenamedIssueEvent_renameable)()
    SetUrl(value *string)()
}
