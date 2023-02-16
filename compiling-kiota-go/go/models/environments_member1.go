package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnvironmentsMember1 
type EnvironmentsMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // The type property
    typeEscaped *string
    // The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
    wait_timer *int32
}
// NewEnvironmentsMember1 instantiates a new environmentsMember1 and sets the default values.
func NewEnvironmentsMember1()(*EnvironmentsMember1) {
    m := &EnvironmentsMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironmentsMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnvironmentsMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironmentsMember1(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnvironmentsMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnvironmentsMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
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
    res["wait_timer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitTimer(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *EnvironmentsMember1) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
func (m *EnvironmentsMember1) GetNodeId()(*string) {
    return m.node_id
}
// GetType gets the type property value. The type property
func (m *EnvironmentsMember1) GetType()(*string) {
    return m.typeEscaped
}
// GetWaitTimer gets the wait_timer property value. The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
func (m *EnvironmentsMember1) GetWaitTimer()(*int32) {
    return m.wait_timer
}
// Serialize serializes information the current object
func (m *EnvironmentsMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
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
        err := writer.WriteInt32Value("wait_timer", m.GetWaitTimer())
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
func (m *EnvironmentsMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *EnvironmentsMember1) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *EnvironmentsMember1) SetNodeId(value *string)() {
    m.node_id = value
}
// SetType sets the type property value. The type property
func (m *EnvironmentsMember1) SetType(value *string)() {
    m.typeEscaped = value
}
// SetWaitTimer sets the wait_timer property value. The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
func (m *EnvironmentsMember1) SetWaitTimer(value *int32)() {
    m.wait_timer = value
}
