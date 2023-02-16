package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MembershipsOrgsItemWithOrgPatchRequestBody 
type MembershipsOrgsItemWithOrgPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewMembershipsOrgsItemWithOrgPatchRequestBody instantiates a new MembershipsOrgsItemWithOrgPatchRequestBody and sets the default values.
func NewMembershipsOrgsItemWithOrgPatchRequestBody()(*MembershipsOrgsItemWithOrgPatchRequestBody) {
    m := &MembershipsOrgsItemWithOrgPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMembershipsOrgsItemWithOrgPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMembershipsOrgsItemWithOrgPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMembershipsOrgsItemWithOrgPatchRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MembershipsOrgsItemWithOrgPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MembershipsOrgsItemWithOrgPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *MembershipsOrgsItemWithOrgPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MembershipsOrgsItemWithOrgPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
