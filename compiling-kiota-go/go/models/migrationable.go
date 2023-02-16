package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Migrationable 
type Migrationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchiveUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExcludeAttachments()(*bool)
    GetExcludeGitData()(*bool)
    GetExcludeMetadata()(*bool)
    GetExcludeOwnerProjects()(*bool)
    GetExcludeReleases()(*bool)
    GetGuid()(*string)
    GetId()(*int32)
    GetLockRepositories()(*bool)
    GetNodeId()(*string)
    GetOrgMetadataOnly()(*bool)
    GetOwner()(NullableSimpleUserable)
    GetRepositories()([]Repositoryable)
    GetState()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetArchiveUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExcludeAttachments(value *bool)()
    SetExcludeGitData(value *bool)()
    SetExcludeMetadata(value *bool)()
    SetExcludeOwnerProjects(value *bool)()
    SetExcludeReleases(value *bool)()
    SetGuid(value *string)()
    SetId(value *int32)()
    SetLockRepositories(value *bool)()
    SetNodeId(value *string)()
    SetOrgMetadataOnly(value *bool)()
    SetOwner(value NullableSimpleUserable)()
    SetRepositories(value []Repositoryable)()
    SetState(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
