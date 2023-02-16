package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependencyGraphDiff_vulnerabilitiesable 
type DependencyGraphDiff_vulnerabilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvisoryGhsaId()(*string)
    GetAdvisorySummary()(*string)
    GetAdvisoryUrl()(*string)
    GetSeverity()(*string)
    SetAdvisoryGhsaId(value *string)()
    SetAdvisorySummary(value *string)()
    SetAdvisoryUrl(value *string)()
    SetSeverity(value *string)()
}
