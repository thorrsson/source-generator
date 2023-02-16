package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistSimple403Error 
type GistSimple403Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The block property
    block GistSimple403Error_blockable
    // The documentation_url property
    documentation_url *string
    // The message property
    message *string
}
// NewGistSimple403Error instantiates a new GistSimple403Error and sets the default values.
func NewGistSimple403Error()(*GistSimple403Error) {
    m := &GistSimple403Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistSimple403ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGistSimple403ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistSimple403Error(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GistSimple403Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlock gets the block property value. The block property
func (m *GistSimple403Error) GetBlock()(GistSimple403Error_blockable) {
    return m.block
}
// GetDocumentationUrl gets the documentation_url property value. The documentation_url property
func (m *GistSimple403Error) GetDocumentationUrl()(*string) {
    return m.documentation_url
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GistSimple403Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["block"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistSimple403Error_blockFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlock(val.(GistSimple403Error_blockable))
        }
        return nil
    }
    res["documentation_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *GistSimple403Error) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *GistSimple403Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("block", m.GetBlock())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentation_url", m.GetDocumentationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *GistSimple403Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlock sets the block property value. The block property
func (m *GistSimple403Error) SetBlock(value GistSimple403Error_blockable)() {
    m.block = value
}
// SetDocumentationUrl sets the documentation_url property value. The documentation_url property
func (m *GistSimple403Error) SetDocumentationUrl(value *string)() {
    m.documentation_url = value
}
// SetMessage sets the message property value. The message property
func (m *GistSimple403Error) SetMessage(value *string)() {
    m.message = value
}
