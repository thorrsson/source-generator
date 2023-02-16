package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LinkWithType hypermedia Link with Type
type LinkWithType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The href property
    href *string
    // The type property
    typeEscaped *string
}
// NewLinkWithType instantiates a new linkWithType and sets the default values.
func NewLinkWithType()(*LinkWithType) {
    m := &LinkWithType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLinkWithTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLinkWithTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLinkWithType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LinkWithType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LinkWithType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["href"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHref(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetHref gets the href property value. The href property
func (m *LinkWithType) GetHref()(*string) {
    return m.href
}
// GetType gets the type property value. The type property
func (m *LinkWithType) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *LinkWithType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("href", m.GetHref())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *LinkWithType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHref sets the href property value. The href property
func (m *LinkWithType) SetHref(value *string)() {
    m.href = value
}
// SetType sets the type property value. The type property
func (m *LinkWithType) SetType(value *string)() {
    m.typeEscaped = value
}
