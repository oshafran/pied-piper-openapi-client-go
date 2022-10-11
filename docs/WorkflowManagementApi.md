# \WorkflowManagementApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflow**](WorkflowManagementApi.md#CreateWorkflow) | **Post** /workflow | 
[**DeleteWorkflow**](WorkflowManagementApi.md#DeleteWorkflow) | **Delete** /workflow | 
[**GetWorkflows**](WorkflowManagementApi.md#GetWorkflows) | **Get** /workflow | 
[**SaveWorkflow**](WorkflowManagementApi.md#SaveWorkflow) | **Put** /workflow | 



## CreateWorkflow

> map[string]interface{} CreateWorkflow(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | Request to create workflow with given user context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowManagementApi.CreateWorkflow(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowManagementApi.CreateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowManagementApi.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Request to create workflow with given user context | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflow

> DeleteWorkflow(ctx).Id(id).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Workflow id (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Request to delete the workflow (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowManagementApi.DeleteWorkflow(context.Background()).Id(id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowManagementApi.DeleteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Workflow id | 
 **body** | **map[string]interface{}** | Request to delete the workflow | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflows

> []map[string]interface{} GetWorkflows(ctx).Type_(type_).Id(id).GroupId(groupId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | Workflow type (optional)
    id := "id_example" // string | Workflow id (optional)
    groupId := "groupId_example" // string | group id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowManagementApi.GetWorkflows(context.Background()).Type_(type_).Id(id).GroupId(groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowManagementApi.GetWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflows`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowManagementApi.GetWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Workflow type | 
 **id** | **string** | Workflow id | 
 **groupId** | **string** | group id | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveWorkflow

> SaveWorkflow(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | Request to save already created workflow with given user context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowManagementApi.SaveWorkflow(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowManagementApi.SaveWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Request to save already created workflow with given user context | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

