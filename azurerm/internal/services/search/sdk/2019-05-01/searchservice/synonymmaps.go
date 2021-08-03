package searchservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
)

// SynonymMapsClient is the client that can be used to manage and query indexes and documents, as well as manage other
// resources, on a search service.
type SynonymMapsClient struct {
	BaseClient
}

// NewSynonymMapsClient creates an instance of the SynonymMapsClient client.
func NewSynonymMapsClient(searchServiceName string) SynonymMapsClient {
	return SynonymMapsClient{New(searchServiceName)}
}

// Create creates a new synonym map.
// Parameters:
// synonymMap - the definition of the synonym map to create.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client SynonymMapsClient) Create(ctx context.Context, synonymMap SynonymMap, clientRequestID *uuid.UUID) (result SynonymMap, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynonymMapsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: synonymMap,
			Constraints: []validation.Constraint{{Target: "synonymMap.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "synonymMap.Format", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "synonymMap.Synonyms", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("searchservice.SynonymMapsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, synonymMap, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SynonymMapsClient) CreatePreparer(ctx context.Context, synonymMap SynonymMap, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPath("/synonymmaps"),
		autorest.WithJSON(synonymMap),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SynonymMapsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SynonymMapsClient) CreateResponder(resp *http.Response) (result SynonymMap, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates a new synonym map or updates a synonym map if it already exists.
// Parameters:
// synonymMapName - the name of the synonym map to create or update.
// synonymMap - the definition of the synonym map to create or update.
// clientRequestID - the tracking ID sent with the request to help with debugging.
// ifMatch - defines the If-Match condition. The operation will be performed only if the ETag on the server
// matches this value.
// ifNoneMatch - defines the If-None-Match condition. The operation will be performed only if the ETag on the
// server does not match this value.
func (client SynonymMapsClient) CreateOrUpdate(ctx context.Context, synonymMapName string, synonymMap SynonymMap, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result SynonymMap, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynonymMapsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: synonymMap,
			Constraints: []validation.Constraint{{Target: "synonymMap.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "synonymMap.Format", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "synonymMap.Synonyms", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("searchservice.SynonymMapsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, synonymMapName, synonymMap, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SynonymMapsClient) CreateOrUpdatePreparer(ctx context.Context, synonymMapName string, synonymMap SynonymMap, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"synonymMapName": autorest.Encode("path", synonymMapName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/synonymmaps('{synonymMapName}')", pathParameters),
		autorest.WithJSON(synonymMap),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Prefer", "return=representation"))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SynonymMapsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SynonymMapsClient) CreateOrUpdateResponder(resp *http.Response) (result SynonymMap, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a synonym map.
// Parameters:
// synonymMapName - the name of the synonym map to delete.
// clientRequestID - the tracking ID sent with the request to help with debugging.
// ifMatch - defines the If-Match condition. The operation will be performed only if the ETag on the server
// matches this value.
// ifNoneMatch - defines the If-None-Match condition. The operation will be performed only if the ETag on the
// server does not match this value.
func (client SynonymMapsClient) Delete(ctx context.Context, synonymMapName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynonymMapsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, synonymMapName, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SynonymMapsClient) DeletePreparer(ctx context.Context, synonymMapName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"synonymMapName": autorest.Encode("path", synonymMapName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/synonymmaps('{synonymMapName}')", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SynonymMapsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SynonymMapsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves a synonym map definition.
// Parameters:
// synonymMapName - the name of the synonym map to retrieve.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client SynonymMapsClient) Get(ctx context.Context, synonymMapName string, clientRequestID *uuid.UUID) (result SynonymMap, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynonymMapsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, synonymMapName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SynonymMapsClient) GetPreparer(ctx context.Context, synonymMapName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"synonymMapName": autorest.Encode("path", synonymMapName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/synonymmaps('{synonymMapName}')", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SynonymMapsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SynonymMapsClient) GetResponder(resp *http.Response) (result SynonymMap, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all synonym maps available for a search service.
// Parameters:
// selectParameter - selects which top-level properties of the synonym maps to retrieve. Specified as a
// comma-separated list of JSON property names, or '*' for all properties. The default is all properties.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client SynonymMapsClient) List(ctx context.Context, selectParameter string, clientRequestID *uuid.UUID) (result ListSynonymMapsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynonymMapsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, selectParameter, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SynonymMapsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SynonymMapsClient) ListPreparer(ctx context.Context, selectParameter string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPath("/synonymmaps"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SynonymMapsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SynonymMapsClient) ListResponder(resp *http.Response) (result ListSynonymMapsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}