package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable 
type ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassPullRequestAllowances()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable)
    GetDismissalRestrictions()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable)
    GetDismissStaleReviews()(*bool)
    GetRequireCodeOwnerReviews()(*bool)
    GetRequiredApprovingReviewCount()(*int32)
    GetRequireLastPushApproval()(*bool)
    SetBypassPullRequestAllowances(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_bypass_pull_request_allowancesable)()
    SetDismissalRestrictions(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviews_dismissal_restrictionsable)()
    SetDismissStaleReviews(value *bool)()
    SetRequireCodeOwnerReviews(value *bool)()
    SetRequiredApprovingReviewCount(value *int32)()
    SetRequireLastPushApproval(value *bool)()
}
