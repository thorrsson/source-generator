package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 
type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The username for users
    users []string
}
// NewItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 instantiates a new ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1()(*ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1) {
    m := &ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetUsers gets the users property value. The username for users
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1) GetUsers()([]string) {
    return m.users
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUsers() != nil {
        err := writer.WriteCollectionOfStringValues("users", m.GetUsers())
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
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUsers sets the users property value. The username for users
func (m *ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1) SetUsers(value []string)() {
    m.users = value
}
