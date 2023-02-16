package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublicUserable 
type PublicUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetBio()(*string)
    GetBlog()(*string)
    GetCollaborators()(*int32)
    GetCompany()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDiskUsage()(*int32)
    GetEmail()(*string)
    GetEventsUrl()(*string)
    GetFollowers()(*int32)
    GetFollowersUrl()(*string)
    GetFollowing()(*int32)
    GetFollowingUrl()(*string)
    GetGistsUrl()(*string)
    GetGravatarId()(*string)
    GetHireable()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLocation()(*string)
    GetLogin()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetOrganizationsUrl()(*string)
    GetOwnedPrivateRepos()(*int32)
    GetPlan()(PublicUser_planable)
    GetPrivateGists()(*int32)
    GetPublicGists()(*int32)
    GetPublicRepos()(*int32)
    GetReceivedEventsUrl()(*string)
    GetReposUrl()(*string)
    GetSiteAdmin()(*bool)
    GetStarredUrl()(*string)
    GetSubscriptionsUrl()(*string)
    GetSuspendedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTotalPrivateRepos()(*int32)
    GetTwitterUsername()(*string)
    GetType()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetBio(value *string)()
    SetBlog(value *string)()
    SetCollaborators(value *int32)()
    SetCompany(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDiskUsage(value *int32)()
    SetEmail(value *string)()
    SetEventsUrl(value *string)()
    SetFollowers(value *int32)()
    SetFollowersUrl(value *string)()
    SetFollowing(value *int32)()
    SetFollowingUrl(value *string)()
    SetGistsUrl(value *string)()
    SetGravatarId(value *string)()
    SetHireable(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLocation(value *string)()
    SetLogin(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOrganizationsUrl(value *string)()
    SetOwnedPrivateRepos(value *int32)()
    SetPlan(value PublicUser_planable)()
    SetPrivateGists(value *int32)()
    SetPublicGists(value *int32)()
    SetPublicRepos(value *int32)()
    SetReceivedEventsUrl(value *string)()
    SetReposUrl(value *string)()
    SetSiteAdmin(value *bool)()
    SetStarredUrl(value *string)()
    SetSubscriptionsUrl(value *string)()
    SetSuspendedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTotalPrivateRepos(value *int32)()
    SetTwitterUsername(value *string)()
    SetType(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
