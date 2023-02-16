package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeOfConductable 
type CodeOfConductable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetHtmlUrl()(*string)
    GetKey()(*string)
    GetName()(*string)
    GetUrl()(*string)
    SetBody(value *string)()
    SetHtmlUrl(value *string)()
    SetKey(value *string)()
    SetName(value *string)()
    SetUrl(value *string)()
}
