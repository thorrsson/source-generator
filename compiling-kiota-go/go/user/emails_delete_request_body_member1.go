package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailsDeleteRequestBodyMember1 deletes one or more email addresses from your GitHub account. Must contain at least one email address. **Note:** Alternatively, you can pass a single email address or an `array` of emails addresses directly, but we recommend that you pass an object using the `emails` key.
type EmailsDeleteRequestBodyMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Email addresses associated with the GitHub user account.
    emails []string
}
// NewEmailsDeleteRequestBodyMember1 instantiates a new EmailsDeleteRequestBodyMember1 and sets the default values.
func NewEmailsDeleteRequestBodyMember1()(*EmailsDeleteRequestBodyMember1) {
    m := &EmailsDeleteRequestBodyMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEmailsDeleteRequestBodyMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailsDeleteRequestBodyMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailsDeleteRequestBodyMember1(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmailsDeleteRequestBodyMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmails gets the emails property value. Email addresses associated with the GitHub user account.
func (m *EmailsDeleteRequestBodyMember1) GetEmails()([]string) {
    return m.emails
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailsDeleteRequestBodyMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmails(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EmailsDeleteRequestBodyMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmails() != nil {
        err := writer.WriteCollectionOfStringValues("emails", m.GetEmails())
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
func (m *EmailsDeleteRequestBodyMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmails sets the emails property value. Email addresses associated with the GitHub user account.
func (m *EmailsDeleteRequestBodyMember1) SetEmails(value []string)() {
    m.emails = value
}
