package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemLabelsItemWithNamePatchRequestBody 
type ItemItemLabelsItemWithNamePatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading `#`.
    color *string
    // A short description of the label. Must be 100 characters or fewer.
    description *string
    // The new name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing `:strawberry:` will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png ":strawberry:"). For a full list of available emoji and codes, see "[Emoji cheat sheet](https://github.com/ikatyang/emoji-cheat-sheet)."
    new_name *string
}
// NewItemItemLabelsItemWithNamePatchRequestBody instantiates a new ItemItemLabelsItemWithNamePatchRequestBody and sets the default values.
func NewItemItemLabelsItemWithNamePatchRequestBody()(*ItemItemLabelsItemWithNamePatchRequestBody) {
    m := &ItemItemLabelsItemWithNamePatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemLabelsItemWithNamePatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemLabelsItemWithNamePatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemLabelsItemWithNamePatchRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemLabelsItemWithNamePatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading `#`.
func (m *ItemItemLabelsItemWithNamePatchRequestBody) GetColor()(*string) {
    return m.color
}
// GetDescription gets the description property value. A short description of the label. Must be 100 characters or fewer.
func (m *ItemItemLabelsItemWithNamePatchRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemLabelsItemWithNamePatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["new_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewName(val)
        }
        return nil
    }
    return res
}
// GetNewName gets the new_name property value. The new name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing `:strawberry:` will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png ":strawberry:"). For a full list of available emoji and codes, see "[Emoji cheat sheet](https://github.com/ikatyang/emoji-cheat-sheet)."
func (m *ItemItemLabelsItemWithNamePatchRequestBody) GetNewName()(*string) {
    return m.new_name
}
// Serialize serializes information the current object
func (m *ItemItemLabelsItemWithNamePatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("new_name", m.GetNewName())
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
func (m *ItemItemLabelsItemWithNamePatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading `#`.
func (m *ItemItemLabelsItemWithNamePatchRequestBody) SetColor(value *string)() {
    m.color = value
}
// SetDescription sets the description property value. A short description of the label. Must be 100 characters or fewer.
func (m *ItemItemLabelsItemWithNamePatchRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetNewName sets the new_name property value. The new name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing `:strawberry:` will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png ":strawberry:"). For a full list of available emoji and codes, see "[Emoji cheat sheet](https://github.com/ikatyang/emoji-cheat-sheet)."
func (m *ItemItemLabelsItemWithNamePatchRequestBody) SetNewName(value *string)() {
    m.new_name = value
}
