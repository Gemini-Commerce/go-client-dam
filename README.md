# Go API client for dam

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import dam "github.com/Gemini-Commerce/go-client-dam"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `dam.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), dam.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `dam.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), dam.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `dam.ContextOperationServerIndices` and `dam.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), dam.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dam.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://dam.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BasicOperationsAPI* | [**BatchUploadAssets**](docs/BasicOperationsAPI.md#batchuploadassets) | **Post** /dam.Dam/BatchUploadAssets | Batch Upload Assets
*BasicOperationsAPI* | [**BatchUploadAssets_0**](docs/BasicOperationsAPI.md#batchuploadassets_0) | **Post** /dam/batch_upload_assets | Batch Upload Assets
*BasicOperationsAPI* | [**CreateAsset**](docs/BasicOperationsAPI.md#createasset) | **Post** /dam.Dam/CreateAsset | Create Asset
*BasicOperationsAPI* | [**CreateAsset_0**](docs/BasicOperationsAPI.md#createasset_0) | **Post** /dam/create_asset | Create Asset
*BasicOperationsAPI* | [**GetAssetByCode**](docs/BasicOperationsAPI.md#getassetbycode) | **Post** /dam.Dam/GetAssetByCode | Get Asset By Code
*BasicOperationsAPI* | [**GetAssetByCode_0**](docs/BasicOperationsAPI.md#getassetbycode_0) | **Post** /dam/get_asset_by_code | Get Asset By Code
*BasicOperationsAPI* | [**ListAssets**](docs/BasicOperationsAPI.md#listassets) | **Post** /dam.Dam/ListAssets | List Assets
*BasicOperationsAPI* | [**ListAssetsByCodes**](docs/BasicOperationsAPI.md#listassetsbycodes) | **Post** /dam.Dam/ListAssetsByCodes | List Assets By Codes
*BasicOperationsAPI* | [**ListAssetsByCodes_0**](docs/BasicOperationsAPI.md#listassetsbycodes_0) | **Post** /dam/list_assets_by_codes | List Assets By Codes
*BasicOperationsAPI* | [**ListAssetsByIds**](docs/BasicOperationsAPI.md#listassetsbyids) | **Post** /dam.Dam/ListAssetsByIds | List Assets By Ids
*BasicOperationsAPI* | [**ListAssetsByIds_0**](docs/BasicOperationsAPI.md#listassetsbyids_0) | **Post** /dam/list_assets_by_ids | List Assets By Ids
*BasicOperationsAPI* | [**ListAssets_0**](docs/BasicOperationsAPI.md#listassets_0) | **Post** /dam/list_assets | List Assets
*BasicOperationsAPI* | [**UpdateAsset**](docs/BasicOperationsAPI.md#updateasset) | **Post** /dam.Dam/UpdateAsset | Update Asset
*BasicOperationsAPI* | [**UpdateAsset_0**](docs/BasicOperationsAPI.md#updateasset_0) | **Post** /dam/update_asset | Update Asset


## Documentation For Models

 - [AssetMetadata](docs/AssetMetadata.md)
 - [AssetOriginTypes](docs/AssetOriginTypes.md)
 - [BatchUploadAssetsRequestFiles](docs/BatchUploadAssetsRequestFiles.md)
 - [DamAsset](docs/DamAsset.md)
 - [DamAssetOrigin](docs/DamAssetOrigin.md)
 - [DamAssetType](docs/DamAssetType.md)
 - [DamBatchUploadAssetsRequest](docs/DamBatchUploadAssetsRequest.md)
 - [DamBatchUploadAssetsResponse](docs/DamBatchUploadAssetsResponse.md)
 - [DamCreateAssetRequest](docs/DamCreateAssetRequest.md)
 - [DamGetAssetByCodeRequest](docs/DamGetAssetByCodeRequest.md)
 - [DamListAssetsByCodesRequest](docs/DamListAssetsByCodesRequest.md)
 - [DamListAssetsByCodesResponse](docs/DamListAssetsByCodesResponse.md)
 - [DamListAssetsByIdsRequest](docs/DamListAssetsByIdsRequest.md)
 - [DamListAssetsByIdsResponse](docs/DamListAssetsByIdsResponse.md)
 - [DamListAssetsRequest](docs/DamListAssetsRequest.md)
 - [DamListAssetsResponse](docs/DamListAssetsResponse.md)
 - [DamUpdateAssetRequest](docs/DamUpdateAssetRequest.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [UpdateAssetRequestPayload](docs/UpdateAssetRequestPayload.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

