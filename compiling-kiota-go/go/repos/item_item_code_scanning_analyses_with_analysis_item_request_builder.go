package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\analyses\{analysis_id}
type ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteQueryParameters deletes a specified code scanning analysis from a repository. Forprivate repositories, you must use an access token with the `repo` scope. For public repositories,you must use an access token with `public_repo` scope.GitHub Apps must have the `security_events` write permission to use this endpoint.You can delete one analysis at a time.To delete a series of analyses, start with the most recent analysis and work backwards.Conceptually, the process is similar to the undo function in a text editor.When you list the analyses for a repository,one or more will be identified as deletable in the response:```"deletable": true```An analysis is deletable when it's the most recent in a set of analyses.Typically, a repository will have multiple sets of analysesfor each enabled code scanning tool,where a set is determined by a unique combination of analysis values:* `ref`* `tool`* `category`If you attempt to delete an analysis that is not the most recent in a set,you'll get a 400 response with the message:```Analysis specified is not deletable.```The response from a successful `DELETE` operation provides you withtwo alternative URLs for deleting the next analysis in the set:`next_analysis_url` and `confirm_delete_url`.Use the `next_analysis_url` URL if you want to avoid accidentally deleting the final analysisin a set. This is a useful option if you want to preserve at least one analysisfor the specified tool in your repository.Use the `confirm_delete_url` URL if you are content to remove all analyses for a tool.When you delete the last analysis in a set, the value of `next_analysis_url` and `confirm_delete_url`in the 200 response is `null`.As an example of the deletion process,let's imagine that you added a workflow that configured a particular code scanning toolto analyze the code in a repository. This tool has added 15 analyses:10 on the default branch, and another 5 on a topic branch.You therefore have two separate sets of analyses for this tool.You've now decided that you want to remove all of the analyses for the tool.To do this you must make 15 separate deletion requests.To start, you must find an analysis that's identified as deletable.Each set of analyses always has one that's identified as deletable.Having found the deletable analysis for one of the two sets,delete this analysis and then continue deleting the next analysis in the set until they're all deleted.Then repeat the process for the second set.The procedure therefore consists of a nested loop:**Outer loop**:* List the analyses for the repository, filtered by tool.* Parse this list to find a deletable analysis. If found:  **Inner loop**:  * Delete the identified analysis.  * Parse the response for the value of `confirm_delete_url` and, if found, use this in the next iteration.The above process assumes that you want to remove all trace of the tool's analyses from the GitHub user interface, for the specified repository, and it therefore uses the `confirm_delete_url` value. Alternatively, you could use the `next_analysis_url` value, which would leave the last analysis in each set undeleted to avoid removing a tool's analysis entirely.
type ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteQueryParameters struct {
    // Allow deletion if the specified analysis is the last in a set. If you attempt to delete the final analysis in a set without setting this parameter to `true`, you'll get a 400 response with the message: `Analysis is last of its type and deletion may result in the loss of historical alert data. Please specify confirm_delete.`
    Confirm_delete *string
}
// ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteQueryParameters
}
// ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal instantiates a new WithAnalysis_ItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    m := &ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/code-scanning/analyses/{analysis_id}{?confirm_delete*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder instantiates a new WithAnalysis_ItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a specified code scanning analysis from a repository. Forprivate repositories, you must use an access token with the `repo` scope. For public repositories,you must use an access token with `public_repo` scope.GitHub Apps must have the `security_events` write permission to use this endpoint.You can delete one analysis at a time.To delete a series of analyses, start with the most recent analysis and work backwards.Conceptually, the process is similar to the undo function in a text editor.When you list the analyses for a repository,one or more will be identified as deletable in the response:```"deletable": true```An analysis is deletable when it's the most recent in a set of analyses.Typically, a repository will have multiple sets of analysesfor each enabled code scanning tool,where a set is determined by a unique combination of analysis values:* `ref`* `tool`* `category`If you attempt to delete an analysis that is not the most recent in a set,you'll get a 400 response with the message:```Analysis specified is not deletable.```The response from a successful `DELETE` operation provides you withtwo alternative URLs for deleting the next analysis in the set:`next_analysis_url` and `confirm_delete_url`.Use the `next_analysis_url` URL if you want to avoid accidentally deleting the final analysisin a set. This is a useful option if you want to preserve at least one analysisfor the specified tool in your repository.Use the `confirm_delete_url` URL if you are content to remove all analyses for a tool.When you delete the last analysis in a set, the value of `next_analysis_url` and `confirm_delete_url`in the 200 response is `null`.As an example of the deletion process,let's imagine that you added a workflow that configured a particular code scanning toolto analyze the code in a repository. This tool has added 15 analyses:10 on the default branch, and another 5 on a topic branch.You therefore have two separate sets of analyses for this tool.You've now decided that you want to remove all of the analyses for the tool.To do this you must make 15 separate deletion requests.To start, you must find an analysis that's identified as deletable.Each set of analyses always has one that's identified as deletable.Having found the deletable analysis for one of the two sets,delete this analysis and then continue deleting the next analysis in the set until they're all deleted.Then repeat the process for the second set.The procedure therefore consists of a nested loop:**Outer loop**:* List the analyses for the repository, filtered by tool.* Parse this list to find a deletable analysis. If found:  **Inner loop**:  * Delete the identified analysis.  * Parse the response for the value of `confirm_delete_url` and, if found, use this in the next iteration.The above process assumes that you want to remove all trace of the tool's analyses from the GitHub user interface, for the specified repository, and it therefore uses the `confirm_delete_url` value. Alternatively, you could use the `next_analysis_url` value, which would leave the last analysis in each set undeleted to avoid removing a tool's analysis entirely.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/code-scanning#delete-a-code-scanning-analysis-from-a-repository
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningAnalysisDeletionable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodeScanningAnalysisDeletion503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodeScanningAnalysisDeletionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningAnalysisDeletionable), nil
}
// Get gets a specified code scanning analysis for a repository.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.GitHub Apps must have the `security_events` read permission to use this endpoint.The default JSON response contains fields that describe the analysis.This includes the Git reference and commit SHA to which the analysis relates,the datetime of the analysis, the name of the code scanning tool,and the number of alerts.The `rules_count` field in the default response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.If you use the Accept header `application/sarif+json`,the response contains the analysis data that was uploaded.This is formatted as[SARIF version 2.1.0](https://docs.oasis-open.org/sarif/sarif/v2.1.0/cs01/sarif-v2.1.0-cs01.html).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/code-scanning#get-a-code-scanning-analysis-for-a-repository
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningAnalysisable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodeScanningAnalysis503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodeScanningAnalysisFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningAnalysisable), nil
}
// ToDeleteRequestInformation deletes a specified code scanning analysis from a repository. Forprivate repositories, you must use an access token with the `repo` scope. For public repositories,you must use an access token with `public_repo` scope.GitHub Apps must have the `security_events` write permission to use this endpoint.You can delete one analysis at a time.To delete a series of analyses, start with the most recent analysis and work backwards.Conceptually, the process is similar to the undo function in a text editor.When you list the analyses for a repository,one or more will be identified as deletable in the response:```"deletable": true```An analysis is deletable when it's the most recent in a set of analyses.Typically, a repository will have multiple sets of analysesfor each enabled code scanning tool,where a set is determined by a unique combination of analysis values:* `ref`* `tool`* `category`If you attempt to delete an analysis that is not the most recent in a set,you'll get a 400 response with the message:```Analysis specified is not deletable.```The response from a successful `DELETE` operation provides you withtwo alternative URLs for deleting the next analysis in the set:`next_analysis_url` and `confirm_delete_url`.Use the `next_analysis_url` URL if you want to avoid accidentally deleting the final analysisin a set. This is a useful option if you want to preserve at least one analysisfor the specified tool in your repository.Use the `confirm_delete_url` URL if you are content to remove all analyses for a tool.When you delete the last analysis in a set, the value of `next_analysis_url` and `confirm_delete_url`in the 200 response is `null`.As an example of the deletion process,let's imagine that you added a workflow that configured a particular code scanning toolto analyze the code in a repository. This tool has added 15 analyses:10 on the default branch, and another 5 on a topic branch.You therefore have two separate sets of analyses for this tool.You've now decided that you want to remove all of the analyses for the tool.To do this you must make 15 separate deletion requests.To start, you must find an analysis that's identified as deletable.Each set of analyses always has one that's identified as deletable.Having found the deletable analysis for one of the two sets,delete this analysis and then continue deleting the next analysis in the set until they're all deleted.Then repeat the process for the second set.The procedure therefore consists of a nested loop:**Outer loop**:* List the analyses for the repository, filtered by tool.* Parse this list to find a deletable analysis. If found:  **Inner loop**:  * Delete the identified analysis.  * Parse the response for the value of `confirm_delete_url` and, if found, use this in the next iteration.The above process assumes that you want to remove all trace of the tool's analyses from the GitHub user interface, for the specified repository, and it therefore uses the `confirm_delete_url` value. Alternatively, you could use the `next_analysis_url` value, which would leave the last analysis in each set undeleted to avoid removing a tool's analysis entirely.
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation gets a specified code scanning analysis for a repository.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.GitHub Apps must have the `security_events` read permission to use this endpoint.The default JSON response contains fields that describe the analysis.This includes the Git reference and commit SHA to which the analysis relates,the datetime of the analysis, the name of the code scanning tool,and the number of alerts.The `rules_count` field in the default response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.If you use the Accept header `application/sarif+json`,the response contains the analysis data that was uploaded.This is formatted as[SARIF version 2.1.0](https://docs.oasis-open.org/sarif/sarif/v2.1.0/cs01/sarif-v2.1.0-cs01.html).
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
