package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepoSearchResultItem_permissions 
type RepoSearchResultItem_permissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The admin property
    admin *bool
    // The maintain property
    maintain *bool
    // The pull property
    pull *bool
    // The push property
    push *bool
    // The triage property
    triage *bool
}
// NewRepoSearchResultItem_permissions instantiates a new repoSearchResultItem_permissions and sets the default values.
func NewRepoSearchResultItem_permissions()(*RepoSearchResultItem_permissions) {
    m := &RepoSearchResultItem_permissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepoSearchResultItem_permissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepoSearchResultItem_permissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepoSearchResultItem_permissions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepoSearchResultItem_permissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdmin gets the admin property value. The admin property
func (m *RepoSearchResultItem_permissions) GetAdmin()(*bool) {
    return m.admin
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepoSearchResultItem_permissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["admin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmin(val)
        }
        return nil
    }
    res["maintain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintain(val)
        }
        return nil
    }
    res["pull"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPull(val)
        }
        return nil
    }
    res["push"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPush(val)
        }
        return nil
    }
    res["triage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTriage(val)
        }
        return nil
    }
    return res
}
// GetMaintain gets the maintain property value. The maintain property
func (m *RepoSearchResultItem_permissions) GetMaintain()(*bool) {
    return m.maintain
}
// GetPull gets the pull property value. The pull property
func (m *RepoSearchResultItem_permissions) GetPull()(*bool) {
    return m.pull
}
// GetPush gets the push property value. The push property
func (m *RepoSearchResultItem_permissions) GetPush()(*bool) {
    return m.push
}
// GetTriage gets the triage property value. The triage property
func (m *RepoSearchResultItem_permissions) GetTriage()(*bool) {
    return m.triage
}
// Serialize serializes information the current object
func (m *RepoSearchResultItem_permissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("admin", m.GetAdmin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maintain", m.GetMaintain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("pull", m.GetPull())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("push", m.GetPush())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("triage", m.GetTriage())
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
func (m *RepoSearchResultItem_permissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdmin sets the admin property value. The admin property
func (m *RepoSearchResultItem_permissions) SetAdmin(value *bool)() {
    m.admin = value
}
// SetMaintain sets the maintain property value. The maintain property
func (m *RepoSearchResultItem_permissions) SetMaintain(value *bool)() {
    m.maintain = value
}
// SetPull sets the pull property value. The pull property
func (m *RepoSearchResultItem_permissions) SetPull(value *bool)() {
    m.pull = value
}
// SetPush sets the push property value. The push property
func (m *RepoSearchResultItem_permissions) SetPush(value *bool)() {
    m.push = value
}
// SetTriage sets the triage property value. The triage property
func (m *RepoSearchResultItem_permissions) SetTriage(value *bool)() {
    m.triage = value
}
