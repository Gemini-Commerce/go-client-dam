# \BasicOperationsAPI

All URIs are relative to *https://dam.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchUploadAssets**](BasicOperationsAPI.md#BatchUploadAssets) | **Post** /dam.Dam/BatchUploadAssets | Batch Upload Assets
[**BatchUploadAssets_0**](BasicOperationsAPI.md#BatchUploadAssets_0) | **Post** /dam/batch_upload_assets | Batch Upload Assets
[**CreateAsset**](BasicOperationsAPI.md#CreateAsset) | **Post** /dam.Dam/CreateAsset | Create Asset
[**CreateAsset_0**](BasicOperationsAPI.md#CreateAsset_0) | **Post** /dam/create_asset | Create Asset
[**GetAssetByCode**](BasicOperationsAPI.md#GetAssetByCode) | **Post** /dam.Dam/GetAssetByCode | Get Asset By Code
[**GetAssetByCode_0**](BasicOperationsAPI.md#GetAssetByCode_0) | **Post** /dam/get_asset_by_code | Get Asset By Code
[**ListAssets**](BasicOperationsAPI.md#ListAssets) | **Post** /dam.Dam/ListAssets | List Assets
[**ListAssetsByCodes**](BasicOperationsAPI.md#ListAssetsByCodes) | **Post** /dam.Dam/ListAssetsByCodes | List Assets By Codes
[**ListAssetsByCodes_0**](BasicOperationsAPI.md#ListAssetsByCodes_0) | **Post** /dam/list_assets_by_codes | List Assets By Codes
[**ListAssetsByIds**](BasicOperationsAPI.md#ListAssetsByIds) | **Post** /dam.Dam/ListAssetsByIds | List Assets By Ids
[**ListAssetsByIds_0**](BasicOperationsAPI.md#ListAssetsByIds_0) | **Post** /dam/list_assets_by_ids | List Assets By Ids
[**ListAssets_0**](BasicOperationsAPI.md#ListAssets_0) | **Post** /dam/list_assets | List Assets
[**UpdateAsset**](BasicOperationsAPI.md#UpdateAsset) | **Post** /dam.Dam/UpdateAsset | Update Asset
[**UpdateAsset_0**](BasicOperationsAPI.md#UpdateAsset_0) | **Post** /dam/update_asset | Update Asset



## BatchUploadAssets

> DamBatchUploadAssetsResponse BatchUploadAssets(ctx).Body(body).Execute()

Batch Upload Assets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamBatchUploadAssetsRequest("TenantId_example", []openapiclient.BatchUploadAssetsRequestFiles{*openapiclient.NewBatchUploadAssetsRequestFiles("Filename_example", "Size_example", "MimeType_example")}) // DamBatchUploadAssetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.BatchUploadAssets(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.BatchUploadAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchUploadAssets`: DamBatchUploadAssetsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.BatchUploadAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUploadAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamBatchUploadAssetsRequest**](DamBatchUploadAssetsRequest.md) |  | 

### Return type

[**DamBatchUploadAssetsResponse**](DamBatchUploadAssetsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUploadAssets_0

> DamBatchUploadAssetsResponse BatchUploadAssets_0(ctx).Body(body).Execute()

Batch Upload Assets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamBatchUploadAssetsRequest("TenantId_example", []openapiclient.BatchUploadAssetsRequestFiles{*openapiclient.NewBatchUploadAssetsRequestFiles("Filename_example", "Size_example", "MimeType_example")}) // DamBatchUploadAssetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.BatchUploadAssets_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.BatchUploadAssets_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchUploadAssets_0`: DamBatchUploadAssetsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.BatchUploadAssets_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUploadAssets_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamBatchUploadAssetsRequest**](DamBatchUploadAssetsRequest.md) |  | 

### Return type

[**DamBatchUploadAssetsResponse**](DamBatchUploadAssetsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAsset

> DamAsset CreateAsset(ctx).Body(body).Execute()

Create Asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamCreateAssetRequest("TenantId_example", openapiclient.damAssetType("UNKNOWN"), "Code_example", *openapiclient.NewDamAssetOrigin()) // DamCreateAssetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.CreateAsset(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.CreateAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAsset`: DamAsset
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.CreateAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamCreateAssetRequest**](DamCreateAssetRequest.md) |  | 

### Return type

[**DamAsset**](DamAsset.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAsset_0

> DamAsset CreateAsset_0(ctx).Body(body).Execute()

Create Asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamCreateAssetRequest("TenantId_example", openapiclient.damAssetType("UNKNOWN"), "Code_example", *openapiclient.NewDamAssetOrigin()) // DamCreateAssetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.CreateAsset_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.CreateAsset_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAsset_0`: DamAsset
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.CreateAsset_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAsset_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamCreateAssetRequest**](DamCreateAssetRequest.md) |  | 

### Return type

[**DamAsset**](DamAsset.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetByCode

> DamAsset GetAssetByCode(ctx).Body(body).Execute()

Get Asset By Code

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamGetAssetByCodeRequest("TenantId_example", "Code_example") // DamGetAssetByCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.GetAssetByCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.GetAssetByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetByCode`: DamAsset
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.GetAssetByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamGetAssetByCodeRequest**](DamGetAssetByCodeRequest.md) |  | 

### Return type

[**DamAsset**](DamAsset.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetByCode_0

> DamAsset GetAssetByCode_0(ctx).Body(body).Execute()

Get Asset By Code

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamGetAssetByCodeRequest("TenantId_example", "Code_example") // DamGetAssetByCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.GetAssetByCode_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.GetAssetByCode_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetByCode_0`: DamAsset
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.GetAssetByCode_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetByCode_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamGetAssetByCodeRequest**](DamGetAssetByCodeRequest.md) |  | 

### Return type

[**DamAsset**](DamAsset.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssets

> DamListAssetsResponse ListAssets(ctx).Body(body).Execute()

List Assets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamListAssetsRequest("TenantId_example") // DamListAssetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.ListAssets(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.ListAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssets`: DamListAssetsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.ListAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamListAssetsRequest**](DamListAssetsRequest.md) |  | 

### Return type

[**DamListAssetsResponse**](DamListAssetsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetsByCodes

> DamListAssetsByCodesResponse ListAssetsByCodes(ctx).Body(body).Execute()

List Assets By Codes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamListAssetsByCodesRequest("TenantId_example", []string{"Codes_example"}) // DamListAssetsByCodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.ListAssetsByCodes(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.ListAssetsByCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetsByCodes`: DamListAssetsByCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.ListAssetsByCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsByCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamListAssetsByCodesRequest**](DamListAssetsByCodesRequest.md) |  | 

### Return type

[**DamListAssetsByCodesResponse**](DamListAssetsByCodesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetsByCodes_0

> DamListAssetsByCodesResponse ListAssetsByCodes_0(ctx).Body(body).Execute()

List Assets By Codes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamListAssetsByCodesRequest("TenantId_example", []string{"Codes_example"}) // DamListAssetsByCodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.ListAssetsByCodes_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.ListAssetsByCodes_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetsByCodes_0`: DamListAssetsByCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.ListAssetsByCodes_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsByCodes_4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamListAssetsByCodesRequest**](DamListAssetsByCodesRequest.md) |  | 

### Return type

[**DamListAssetsByCodesResponse**](DamListAssetsByCodesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetsByIds

> DamListAssetsByIdsResponse ListAssetsByIds(ctx).Body(body).Execute()

List Assets By Ids

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamListAssetsByIdsRequest("TenantId_example", []string{"Ids_example"}) // DamListAssetsByIdsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.ListAssetsByIds(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.ListAssetsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetsByIds`: DamListAssetsByIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.ListAssetsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamListAssetsByIdsRequest**](DamListAssetsByIdsRequest.md) |  | 

### Return type

[**DamListAssetsByIdsResponse**](DamListAssetsByIdsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetsByIds_0

> DamListAssetsByIdsResponse ListAssetsByIds_0(ctx).Body(body).Execute()

List Assets By Ids

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamListAssetsByIdsRequest("TenantId_example", []string{"Ids_example"}) // DamListAssetsByIdsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.ListAssetsByIds_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.ListAssetsByIds_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetsByIds_0`: DamListAssetsByIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.ListAssetsByIds_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsByIds_5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamListAssetsByIdsRequest**](DamListAssetsByIdsRequest.md) |  | 

### Return type

[**DamListAssetsByIdsResponse**](DamListAssetsByIdsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssets_0

> DamListAssetsResponse ListAssets_0(ctx).Body(body).Execute()

List Assets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamListAssetsRequest("TenantId_example") // DamListAssetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.ListAssets_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.ListAssets_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssets_0`: DamListAssetsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.ListAssets_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssets_6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamListAssetsRequest**](DamListAssetsRequest.md) |  | 

### Return type

[**DamListAssetsResponse**](DamListAssetsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAsset

> DamAsset UpdateAsset(ctx).Body(body).Execute()

Update Asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamUpdateAssetRequest("TenantId_example", "Id_example") // DamUpdateAssetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UpdateAsset(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UpdateAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAsset`: DamAsset
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UpdateAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamUpdateAssetRequest**](DamUpdateAssetRequest.md) |  | 

### Return type

[**DamAsset**](DamAsset.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAsset_0

> DamAsset UpdateAsset_0(ctx).Body(body).Execute()

Update Asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-dam"
)

func main() {
	body := *openapiclient.NewDamUpdateAssetRequest("TenantId_example", "Id_example") // DamUpdateAssetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UpdateAsset_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UpdateAsset_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAsset_0`: DamAsset
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UpdateAsset_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsset_7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DamUpdateAssetRequest**](DamUpdateAssetRequest.md) |  | 

### Return type

[**DamAsset**](DamAsset.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

