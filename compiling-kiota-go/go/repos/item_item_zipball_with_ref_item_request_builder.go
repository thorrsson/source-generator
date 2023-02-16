package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemZipballWithRefItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\zipball\{ref}
type ItemItemZipballWithRefItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemZipballWithRefItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemZipballWithRefItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemZipballWithRefItemRequestBuilderInternal instantiates a new WithRefItemRequestBuilder and sets the default values.
func NewItemItemZipballWithRefItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemZipballWithRefItemRequestBuilder) {
    m := &ItemItemZipballWithRefItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/zipball/{ref}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemZipballWithRefItemRequestBuilder instantiates a new WithRefItemRequestBuilder and sets the default values.
func NewItemItemZipballWithRefItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemZipballWithRefItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemZipballWithRefItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download a zip archive for a repository. If you omit `:ref`, the repository’s default branch (usually`main`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to usethe `Location` header to make a second `GET` request.**Note**: For private repositories, these links are temporary and expire after five minutes. If the repository is empty, you will receive a 404 when you follow the redirect.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/repos#download-a-repository-archive
func (m *ItemItemZipballWithRefItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemZipballWithRefItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation gets a redirect URL to download a zip archive for a repository. If you omit `:ref`, the repository’s default branch (usually`main`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to usethe `Location` header to make a second `GET` request.**Note**: For private repositories, these links are temporary and expire after five minutes. If the repository is empty, you will receive a 404 when you follow the redirect.
func (m *ItemItemZipballWithRefItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemZipballWithRefItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
