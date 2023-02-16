package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentBranchPolicy details of a deployment branch policy.
type DeploymentBranchPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the branch policy.
    id *int32
    // The name pattern that branches must match in order to deploy to the environment.
    name *string
    // The node_id property
    node_id *string
}
// NewDeploymentBranchPolicy instantiates a new deploymentBranchPolicy and sets the default values.
func NewDeploymentBranchPolicy()(*DeploymentBranchPolicy) {
    m := &DeploymentBranchPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentBranchPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentBranchPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentBranchPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentBranchPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentBranchPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    return res
}
// GetId gets the id property value. The unique identifier of the branch policy.
func (m *DeploymentBranchPolicy) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name pattern that branches must match in order to deploy to the environment.
func (m *DeploymentBranchPolicy) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
func (m *DeploymentBranchPolicy) GetNodeId()(*string) {
    return m.node_id
}
// Serialize serializes information the current object
func (m *DeploymentBranchPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentBranchPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier of the branch policy.
func (m *DeploymentBranchPolicy) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name pattern that branches must match in order to deploy to the environment.
func (m *DeploymentBranchPolicy) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *DeploymentBranchPolicy) SetNodeId(value *string)() {
    m.node_id = value
}
