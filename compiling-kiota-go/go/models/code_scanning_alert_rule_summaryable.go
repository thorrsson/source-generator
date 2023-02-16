package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAlertRuleSummaryable 
type CodeScanningAlertRuleSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetId()(*string)
    GetName()(*string)
    GetSeverity()(*CodeScanningAlertRuleSummary_severity)
    GetTags()([]string)
    SetDescription(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetSeverity(value *CodeScanningAlertRuleSummary_severity)()
    SetTags(value []string)()
}
