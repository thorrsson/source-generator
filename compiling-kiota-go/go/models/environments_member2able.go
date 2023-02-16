package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnvironmentsMember2able 
type EnvironmentsMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetNodeId()(*string)
    GetReviewers()([]EnvironmentsMember2_reviewersable)
    GetType()(*string)
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetReviewers(value []EnvironmentsMember2_reviewersable)()
    SetType(value *string)()
}
