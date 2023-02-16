package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchRestrictionPolicy_teamsable 
type BranchRestrictionPolicy_teamsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetMembersUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetParent()(*string)
    GetPermission()(*string)
    GetPrivacy()(*string)
    GetRepositoriesUrl()(*string)
    GetSlug()(*string)
    GetUrl()(*string)
    SetDescription(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetMembersUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetParent(value *string)()
    SetPermission(value *string)()
    SetPrivacy(value *string)()
    SetRepositoriesUrl(value *string)()
    SetSlug(value *string)()
    SetUrl(value *string)()
}
