# \LightConeWebSocketAPIApi

All URIs are relative to *http://uat2.loopring.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWsKeyV2**](LightConeWebSocketAPIApi.md#GetWsKeyV2) | **Get** /v2/ws/key | 
[**GetWsKeyV3V3**](LightConeWebSocketAPIApi.md#GetWsKeyV3V3) | **Get** /v3/ws/key | 
[**WsV3**](LightConeWebSocketAPIApi.md#WsV3) | **Get** /v3/ws | 



## GetWsKeyV2

> map[string]interface{} GetWsKeyV2(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LightConeWebSocketAPIApi.GetWsKeyV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LightConeWebSocketAPIApi.GetWsKeyV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWsKeyV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LightConeWebSocketAPIApi.GetWsKeyV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWsKeyV2Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWsKeyV3V3

> map[string]interface{} GetWsKeyV3V3(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LightConeWebSocketAPIApi.GetWsKeyV3V3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LightConeWebSocketAPIApi.GetWsKeyV3V3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWsKeyV3V3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LightConeWebSocketAPIApi.GetWsKeyV3V3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWsKeyV3V3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WsV3

> map[string]interface{} WsV3(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LightConeWebSocketAPIApi.WsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LightConeWebSocketAPIApi.WsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WsV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LightConeWebSocketAPIApi.WsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWsV3Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

