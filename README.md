# Hawkeye SDK for Go

A lightweight Go client for interacting with the Hawkeye Claims API. The SDK wraps the platform's REST endpoints with idiomatic Go types, request helpers, and ergonomic option builders so you can create claims, upload documents, and record log-trail entries from your own applications.

## Features

- 🔐 Token-based authentication with production and QA (dev) environments.
- 📁 Claims service with client-side validation for required fields and typed response models.
- 📄 Document upload helpers with optional metadata (category, visibility, notes).
- 📝 Log trail entry creation with customizable timestamps.
- 🏢 Insurance companies lookup with search and filtering capabilities.
- 🧪 Comprehensive unit tests using `httptest` for safe local development.

## Installation

Make sure you are using Go 1.22 or newer. Then add the module to your project:

```bash
go get github.com/Hawkeye-Claims/hawkeyesdk@latest
```

The module path is `github.com/Hawkeye-Claims/hawkeyesdk`, so you can import it directly:

```go
import "github.com/Hawkeye-Claims/hawkeyesdk/pkg/hawkeyesdk"
```

## Quick start

```go
package main

import (
    "context"
    "log"

    "github.com/Hawkeye-Claims/hawkeyesdk/pkg/hawkeyesdk"
)

func main() {
    client := hawkeyesdk.NewHawkeyeClient("<your-api-token>")

    // You can switch to the Hawkeye QA environment when you need it:
    // client := hawkeyesdk.NewHawkeyeClient("<your-api-token>", hawkeyesdk.WithEnvironment(hawkeyesdk.DEV))

    claim := hawkeyesdk.ClaimPost{
        RenterName:     "Test Renter",
        InsCompaniesID: "Hawkeye Insurance",
        DateOfLoss:     "2024-01-01",
        VehMake:        "Ford",
        VehModel:       "F150",
        VehColor:       "Blue",
        VehVIN:         "VIN123",
    }

    resp, err := client.Claims.CreateClaim(context.Background(), claim)
    if err != nil {
        log.Fatalf("create claim failed: %v", err)
    }

    log.Printf("Created claim #%d: %s", resp.Filenumber, resp.Message)
}
```

### Authentication & environments

- **Authentication:** Pass your Hawkeye API token to `NewHawkeyeClient`. The SDK automatically injects the token in the `Authorization` header for every request.
- **Environments:** Production is the default (`https://hawkeye.g2it.co/api`). To target QA, pass `hawkeyesdk.WithEnvironment(hawkeyesdk.DEV)` when constructing the client. You can also override `ClientSettings.BaseUrl` or `ClientSettings.HTTPClient` after creation if you need full control (for example, to inject custom transports or mock servers).

## Services overview

### Claims

The `ClaimsService` exposes helpers to create, update, and fetch claims.

```go
resp, err := client.Claims.CreateClaim(ctx, claim)
claim, err := client.Claims.GetSingleClaim(ctx, filenumber)
claims, err := client.Claims.GetClaims(ctx, hawkeyesdk.WithIncludeInactive(true))
```

When creating a claim, the SDK validates the required fields before sending the request. You can discover the required list at runtime:

```go
required := hawkeyesdk.ClaimPostRequiredFields()
// []string{"RenterName", "InsCompaniesID", "DateOfLoss", "VehMake", "VehModel", "VehColor", "VehVIN"}
```

If validation fails, the call returns an error that wraps the missing fields so you can surface friendly messages to your users.

### Insurance companies

Query the list of insurance companies available in the Hawkeye system:

```go
companies, err := client.InsCompanies.GetInsuranceCompanies(ctx)

// Search with a query and specify limit (max 20, defaults to 5)
companies, err := client.InsCompanies.GetInsuranceCompanies(
    ctx,
    hawkeyesdk.WithQueryParameters("State Farm", 10),
)
```

The service returns an array of `InsCompany` structs containing the company ID, name, and optional probability score (used for search ranking). Use the company ID when creating or updating claims via the `InsCompaniesID` field.

### Document files

Upload links to supporting documents with metadata describing the file:

```go
resp, err := client.DocFiles.UploadFile(
    claimID,
    "https://storage.example.com/myfile.pdf",
    hawkeyesdk.WithCategory(hawkeyesdk.POLICY),
    hawkeyesdk.WithVisibleToClient(true),
    hawkeyesdk.WithNotes("Uploaded automatically from CRM"),
)
```

Categories are strongly typed via the `DocType` enumeration and automatically converted to the strings the API expects. When the upload succeeds, you receive a simple `ApiResponse` with the API-provided message.

### Log trails

Capture activity history against a claim:

```go
resp, err := client.LogTrails.CreateLogTrail(ctx, claimID, "Contacted insured", hawkeyesdk.WithDate("04/12/2024"))
```

If you omit the date, the SDK defaults to the current date in `MM/DD/YYYY` format.

## Error handling

Non-2xx responses are translated into an `*hawkeyesdk.APIError` that includes the HTTP status code and any message returned by Hawkeye. You can type-assert to access the structured fields:

```go
if err != nil {
    var apiErr *hawkeyesdk.APIError
    if errors.As(err, &apiErr) {
        log.Printf("hawkeye error %d: %s", apiErr.StatusCode, apiErr.Message)
    }
}
```

For serialization errors, network failures, or validation issues, the SDK returns wrapped Go errors so callers keep full context.

## Running tests

The repository ships with unit tests that exercise the HTTP clients using `httptest` servers. Run them locally with:

```bash
go test ./...
```

## Contributing

1. Fork the repository and create a feature branch.
2. Install Go 1.22 or newer.
3. Run `go test ./...` before opening a pull request.
4. Describe the context of your change clearly—especially any new Hawkeye endpoints or models.

Issues and pull requests are welcome!

## Project structure

```
go.mod
pkg/
  hawkeyesdk/
    claims.go          // claim management client & validation helpers
    docfiles.go        // document upload client
    logtrails.go       // log trail entry client
    inscompanies.go    // insurance companies lookup client
    models.go          // shared response/request models and enums
    errors.go          // API error translation helpers
    client.go          // root client wiring for all services
    *_test.go          // unit tests using httptest servers
```

The SDK lives under `pkg/hawkeyesdk`, which is the public package consumers import. Tests mirror the service files to keep behavior well covered.

## License

Released under the [MIT License](./LICENSE).
