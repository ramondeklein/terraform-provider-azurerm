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

// SkillsetsClient is the client that can be used to manage and query indexes and documents, as well as manage other
// resources, on a search service.
type SkillsetsClient struct {
	BaseClient
}

// NewSkillsetsClient creates an instance of the SkillsetsClient client.
func NewSkillsetsClient(searchServiceName string) SkillsetsClient {
	return SkillsetsClient{New(searchServiceName)}
}

// Create creates a new skillset in a search service.
// Parameters:
// skillset - the skillset containing one or more skills to create in a search service.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client SkillsetsClient) Create(ctx context.Context, skillset Skillset, clientRequestID *uuid.UUID) (result Skillset, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SkillsetsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skillset,
			Constraints: []validation.Constraint{{Target: "skillset.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "skillset.Description", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "skillset.Skills", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("searchservice.SkillsetsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, skillset, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SkillsetsClient) CreatePreparer(ctx context.Context, skillset Skillset, clientRequestID *uuid.UUID) (*http.Request, error) {
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
		autorest.WithPath("/skillsets"),
		autorest.WithJSON(skillset),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SkillsetsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SkillsetsClient) CreateResponder(resp *http.Response) (result Skillset, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates a new skillset in a search service or updates the skillset if it already exists.
// Parameters:
// skillsetName - the name of the skillset to create or update.
// skillset - the skillset containing one or more skills to create or update in a search service.
// clientRequestID - the tracking ID sent with the request to help with debugging.
// ifMatch - defines the If-Match condition. The operation will be performed only if the ETag on the server
// matches this value.
// ifNoneMatch - defines the If-None-Match condition. The operation will be performed only if the ETag on the
// server does not match this value.
func (client SkillsetsClient) CreateOrUpdate(ctx context.Context, skillsetName string, skillset Skillset, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result Skillset, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SkillsetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skillset,
			Constraints: []validation.Constraint{{Target: "skillset.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "skillset.Description", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "skillset.Skills", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("searchservice.SkillsetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, skillsetName, skillset, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SkillsetsClient) CreateOrUpdatePreparer(ctx context.Context, skillsetName string, skillset Skillset, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"skillsetName": autorest.Encode("path", skillsetName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/skillsets('{skillsetName}')", pathParameters),
		autorest.WithJSON(skillset),
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
func (client SkillsetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SkillsetsClient) CreateOrUpdateResponder(resp *http.Response) (result Skillset, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a skillset in a search service.
// Parameters:
// skillsetName - the name of the skillset to delete.
// clientRequestID - the tracking ID sent with the request to help with debugging.
// ifMatch - defines the If-Match condition. The operation will be performed only if the ETag on the server
// matches this value.
// ifNoneMatch - defines the If-None-Match condition. The operation will be performed only if the ETag on the
// server does not match this value.
func (client SkillsetsClient) Delete(ctx context.Context, skillsetName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SkillsetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, skillsetName, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SkillsetsClient) DeletePreparer(ctx context.Context, skillsetName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"skillsetName": autorest.Encode("path", skillsetName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/skillsets('{skillsetName}')", pathParameters),
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
func (client SkillsetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SkillsetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves a skillset in a search service.
// Parameters:
// skillsetName - the name of the skillset to retrieve.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client SkillsetsClient) Get(ctx context.Context, skillsetName string, clientRequestID *uuid.UUID) (result Skillset, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SkillsetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, skillsetName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SkillsetsClient) GetPreparer(ctx context.Context, skillsetName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"skillsetName": autorest.Encode("path", skillsetName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/skillsets('{skillsetName}')", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SkillsetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SkillsetsClient) GetResponder(resp *http.Response) (result Skillset, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all skillsets in a search service.
// Parameters:
// selectParameter - selects which top-level properties of the skillsets to retrieve. Specified as a
// comma-separated list of JSON property names, or '*' for all properties. The default is all properties.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client SkillsetsClient) List(ctx context.Context, selectParameter string, clientRequestID *uuid.UUID) (result ListSkillsetsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SkillsetsClient.List")
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
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.SkillsetsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SkillsetsClient) ListPreparer(ctx context.Context, selectParameter string, clientRequestID *uuid.UUID) (*http.Request, error) {
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
		autorest.WithPath("/skillsets"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SkillsetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SkillsetsClient) ListResponder(resp *http.Response) (result ListSkillsetsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}