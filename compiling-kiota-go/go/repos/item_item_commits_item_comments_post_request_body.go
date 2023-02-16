package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCommitsItemCommentsPostRequestBody 
type ItemItemCommitsItemCommentsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contents of the comment.
    body *string
    // **Deprecated**. Use **position** parameter instead. Line number in the file to comment on.
    line *int32
    // Relative path of the file to comment on.
    path *string
    // Line index in the diff to comment on.
    position *int32
}
// NewItemItemCommitsItemCommentsPostRequestBody instantiates a new ItemItemCommitsItemCommentsPostRequestBody and sets the default values.
func NewItemItemCommitsItemCommentsPostRequestBody()(*ItemItemCommitsItemCommentsPostRequestBody) {
    m := &ItemItemCommitsItemCommentsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCommitsItemCommentsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCommitsItemCommentsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCommitsItemCommentsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCommitsItemCommentsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The contents of the comment.
func (m *ItemItemCommitsItemCommentsPostRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCommitsItemCommentsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    return res
}
// GetLine gets the line property value. **Deprecated**. Use **position** parameter instead. Line number in the file to comment on.
func (m *ItemItemCommitsItemCommentsPostRequestBody) GetLine()(*int32) {
    return m.line
}
// GetPath gets the path property value. Relative path of the file to comment on.
func (m *ItemItemCommitsItemCommentsPostRequestBody) GetPath()(*string) {
    return m.path
}
// GetPosition gets the position property value. Line index in the diff to comment on.
func (m *ItemItemCommitsItemCommentsPostRequestBody) GetPosition()(*int32) {
    return m.position
}
// Serialize serializes information the current object
func (m *ItemItemCommitsItemCommentsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("line", m.GetLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("position", m.GetPosition())
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
func (m *ItemItemCommitsItemCommentsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The contents of the comment.
func (m *ItemItemCommitsItemCommentsPostRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetLine sets the line property value. **Deprecated**. Use **position** parameter instead. Line number in the file to comment on.
func (m *ItemItemCommitsItemCommentsPostRequestBody) SetLine(value *int32)() {
    m.line = value
}
// SetPath sets the path property value. Relative path of the file to comment on.
func (m *ItemItemCommitsItemCommentsPostRequestBody) SetPath(value *string)() {
    m.path = value
}
// SetPosition sets the position property value. Line index in the diff to comment on.
func (m *ItemItemCommitsItemCommentsPostRequestBody) SetPosition(value *int32)() {
    m.position = value
}
