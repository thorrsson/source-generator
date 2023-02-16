package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks 
type ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The `id` of the GitHub App.
    app_id *int32
    // Set to `true` to enable automatic creation of CheckSuite events upon pushes to the repository, or `false` to disable them.
    setting *bool
}
// NewItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks instantiates a new ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks and sets the default values.
func NewItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks()(*ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) {
    m := &ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppId gets the app_id property value. The `id` of the GitHub App.
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) GetAppId()(*int32) {
    return m.app_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["app_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
        }
        return nil
    }
    return res
}
// GetSetting gets the setting property value. Set to `true` to enable automatic creation of CheckSuite events upon pushes to the repository, or `false` to disable them.
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) GetSetting()(*bool) {
    return m.setting
}
// Serialize serializes information the current object
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("app_id", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppId sets the app_id property value. The `id` of the GitHub App.
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) SetAppId(value *int32)() {
    m.app_id = value
}
// SetSetting sets the setting property value. Set to `true` to enable automatic creation of CheckSuite events upon pushes to the repository, or `false` to disable them.
func (m *ItemItemCheckSuitesPreferencesPatchRequestBody_auto_trigger_checks) SetSetting(value *bool)() {
    m.setting = value
}
