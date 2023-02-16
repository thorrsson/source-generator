package appmanifests

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithCodeItemRequestBuilder builds and executes requests for operations under \app-manifests\{code}
type WithCodeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewWithCodeItemRequestBuilderInternal instantiates a new WithCodeItemRequestBuilder and sets the default values.
func NewWithCodeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCodeItemRequestBuilder) {
    m := &WithCodeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app-manifests/{code}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithCodeItemRequestBuilder instantiates a new WithCodeItemRequestBuilder and sets the default values.
func NewWithCodeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCodeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithCodeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Conversions the conversions property
func (m *WithCodeItemRequestBuilder) Conversions()(*ItemConversionsRequestBuilder) {
    return NewItemConversionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
