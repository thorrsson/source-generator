package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertSecurityAdvisory_identifiers an advisory identifier.
type DependabotAlertSecurityAdvisory_identifiers struct {
    // The type of advisory identifier.
    typeEscaped *DependabotAlertSecurityAdvisory_identifiers_type
    // The value of the advisory identifer.
    value *string
}
// NewDependabotAlertSecurityAdvisory_identifiers instantiates a new dependabotAlertSecurityAdvisory_identifiers and sets the default values.
func NewDependabotAlertSecurityAdvisory_identifiers()(*DependabotAlertSecurityAdvisory_identifiers) {
    m := &DependabotAlertSecurityAdvisory_identifiers{
    }
    return m
}
// CreateDependabotAlertSecurityAdvisory_identifiersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDependabotAlertSecurityAdvisory_identifiersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlertSecurityAdvisory_identifiers(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DependabotAlertSecurityAdvisory_identifiers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDependabotAlertSecurityAdvisory_identifiers_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*DependabotAlertSecurityAdvisory_identifiers_type))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type of advisory identifier.
func (m *DependabotAlertSecurityAdvisory_identifiers) GetType()(*DependabotAlertSecurityAdvisory_identifiers_type) {
    return m.typeEscaped
}
// GetValue gets the value property value. The value of the advisory identifer.
func (m *DependabotAlertSecurityAdvisory_identifiers) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DependabotAlertSecurityAdvisory_identifiers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// SetType sets the type property value. The type of advisory identifier.
func (m *DependabotAlertSecurityAdvisory_identifiers) SetType(value *DependabotAlertSecurityAdvisory_identifiers_type)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. The value of the advisory identifer.
func (m *DependabotAlertSecurityAdvisory_identifiers) SetValue(value *string)() {
    m.value = value
}
