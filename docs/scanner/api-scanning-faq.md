> Source: https://portswigger.net/burp/documentation/scanner/api-scanning-faq

DASTProfessional

# API scanning FAQs

-

**Last updated: ** September 3, 2026
-

**Read time: ** 17 Minutes

 This page answers common questions about scanning APIs with Burp Suite Professional and Burp Suite DAST.

## Format support

 Which format you use, and how you supply it, both affect what Burp can parse.

#### Which API definition formats can Burp scan?

| **Format**  | **Supported**
| OpenAPI 3.0 and 3.1 | Yes (`JSON` or `YAML`)
| OpenAPI 3.2 | Provisional
| Swagger 2.0 | Yes (`JSON` or `YAML`)
| Postman Collections | Yes (`v2.1.0` only)
| SOAP WSDL | Yes (1.1 and 1.2)
| GraphQL | Yes (through introspection)

 Support for OpenAPI 3.2 is provisional. Common features for this format work, but Burp may not fully use some constructs introduced in 3.2.


 GraphQL works differently from the other formats. Burp scans the live endpoint through introspection rather than you uploading a file. For more information on how this process works, see [GraphQL](https://portswigger.net/burp/documentation/scanner/api-scanning-faq#graphql).


#### Should I supply the definition as a file or a URL?

 Both Burp Suite Professional and Burp Suite DAST accept either a file upload or a URL. The best option for you depends on how often your definition changes:


-

**File upload** - If you upload a file, Burp uses that exact file every time the scan runs. This suits a stable definition, or a case in which Burp can't reach the definition URL from where it runs.
-

**URL** - If you supply a URL, Burp fetches the latest definition every time the scan runs. This suits a definition that changes often, for example one served by a CI pipeline, when you want scans to track the latest version.

 If you have a very large definition, upload it as a file. Fetching a large definition from a URL can cause issues.


#### Note

 When you supply a URL, point Burp at the raw definition rather than a documentation page. A common mistake is to use the human-readable API documentation page (Swagger UI, ReDoc, or similar) instead of the underlying `JSON` or `YAML`. The documentation page renders as an HTML page in your browser, while the raw definition usually lives at a path such as `/openapi.json`, `/swagger.json`, or `/v3/api-docs`.


 To test the link you plan to supply, open the URL in your browser. If you see a rendered documentation page rather than raw `JSON` or `YAML`, find the underlying definition URL and use that instead.


#### Which request body content types does Burp support?

 Burp generates request bodies from the schemas in your OpenAPI definition. It supports the content types it can serialize:

| **Content type**  | **Supported**
| `application/json` | Yes
| `application/x-www-form-urlencoded` | Yes
| `application/xml` | No
| `text/plain` | No

 Burp treats custom subtypes such as `application/vnd.api+json` or `application/hal+json` as their base type, which is `JSON` in these examples. Burp does not support fully proprietary content types, and skips operations involving these types.


 These limits apply only to bodies that Burp generates from an OpenAPI schema. Postman Collections work differently, because they carry complete request bodies that Burp sends as it finds them. For the content types a Postman Collection can use, see [Why does my collection fail with a "Body mode not understood" error?](https://portswigger.net/burp/documentation/scanner/api-scanning-faq#why-does-my-collection-fail-with-a-body-mode-not-understood-error)

#### Does Burp resolve external references?

 No. Burp does not resolve external `$ref`s in OpenAPI, WSDL, or XSD files. Following external references would risk server-side request forgery (SSRF), make scans depend on resources outside the definition, and make parsing less predictable.


 If your definition is split across multiple files, bundle it into one self-contained file before you upload it. Most OpenAPI and WSDL tools have a bundle or flatten option.


#### Does Burp use example values from my definition?

 Yes. Burp reads example values from your definition and uses them to seed its requests. If a parameter has no example, Burp generates one. For the full rules on how Burp turns parameters into requests, including how it treats optional and enumerated parameters in OpenAPI and SOAP, see [Requirements for API scanning](https://portswigger.net/burp/documentation/scanner/api-scanning-reqs#openapi-endpoints).


## Parser errors and unsupported features

 Burp reports several types of schema problem as an "unsupported features" error.

#### Why does my definition fail to parse?

 Burp rejects definitions it can't parse. This can be because:


-

Your `YAML` file is larger than 3 MB. Convert the file to `JSON`, which has no size limit, or split it.
-

Your file has syntax errors. Validate it with an OpenAPI or Swagger validator before you upload it.
-

Your definition uses external `$ref`s. Burp blocks external references. Bundle the definition into one file.

#### Why are some operations missing from my scan coverage?

 Some schema patterns cause Burp to drop operations from coverage. Dropped operations show in the scan event log as "Skipping location in API definition".


 Three schema patterns account for most dropped operations: an object header parameter, a composition schema on a header parameter, and a bare primitive request body. The answers that follow cover each one.


#### Why does Burp skip operations with an object header parameter?

 The HTTP protocol only supports plain text values in headers. If you declare a header as `type: object`, or point it at an object schema with `$ref`, then Burp can't serialize it and skips the operation.


For example, Burp skips this operation: `parameters:
  - name: X-Context
    in: header
    schema:
      $ref: '#/components/schemas/ContextObject'`

Burp scans this operation: `parameters:
  - name: X-Context
    in: header
    schema:
      type: string`

 If your application expects a JSON-serialized header at runtime, declare `type: string` in the definition. Serialization is the responsibility of the application rather than the schema.


#### Why does Burp skip operations with a composition schema on a header?

 A header schema must be a primitive type: string, integer, number, or boolean. Burp skips the whole operation for anything else, including a composition schema (`anyOf`, `oneOf`, or `allOf`).


 If your API uses Python, you may hit this without writing the schema yourself. A FastAPI or Pydantic `Optional[str]` annotation generates `anyOf: [{type: string}, {type: "null"}]`, which OpenAPI 3.1 reads as a string that can also be null. An HTTP header can't be null. It is either present with a string value, or absent. Declare `type: string` instead.


For example, Burp skips this operation: `parameters:
  - name: X-Api-Key
    in: header
    required: false
    schema:
      anyOf:
        - type: string
        - type: "null"`

Burp scans this operation: `parameters:
  - name: X-Api-Key
    in: header
    required: false
    schema:
      type: string`

#### Why does Burp skip operations with a bare primitive request body?

 A request body must be `type: object` with `properties`. A bare `type: string` body fails with an "unsupported features" error.


Burp skips this operation: `- in: body
  name: cardNumber
  schema:
    type: string`

Burp scans this operation: `- in: body
  name: cardNumber
  schema:
    type: object
    properties:
      cardNumber:
        type: string`

#### What do the "unsupported feature" categories in the log mean?

 If you read the scan log or the API parsing details, you may see one of these category names. Burp Scanner needs a primitive type (string, integer, number, or boolean) in certain positions, such as HTTP headers and query parameters. These categories all mean that your schema puts something more complex in one of those positions:


-

`OBJECT_TYPE` - an object, or a `$ref` to an object schema.
-

`COMPOSITION_SCHEMA` - `anyOf`, `oneOf`, or `allOf`.
-

`COMPLEX_JSON` - a JSON object or array.
-

`MULTIPART_STRUCTURE` - multipart form data.
-

`FORM_STRUCTURE` - a URL-encoded form structure.

 To fix any of these, express the parameter as a primitive type in the schema, even if your application accepts something more complex at runtime.


 The other categories point to different issues:


-

`XML_STRUCTURE` - XML that appears outside a SOAP body. Burp only handles XML inside SOAP definitions.
-

`RECURSIVE_SCHEMA` - a schema that references itself, or another schema, deeper than Burp's recursion limit. Burp drops the fragment rather than expand it indefinitely. Reduce the nesting depth to fix it.
-

`UNKNOWN_TYPE` - a `type` value Burp doesn't recognize, usually a typo or a non-standard value. Check it against the permitted OpenAPI type values.

#### Are there limits on schema nesting?

 Yes. Burp caps schema recursion at four levels per schema. Beyond that, Burp marks the fragment as `RECURSIVE_SCHEMA` and skips it rather than expanding it indefinitely. Burp also caps GraphQL introspection traversal at a fixed depth.


## URLs and base paths

 How Burp builds request URLs depends on the definition format. Mismatches are a common cause of failed requests.

#### My scan runs, but every request returns 404. What went wrong?

 In Swagger 2.0, Burp builds each URL as `scheme + host + basePath + path`. If your `paths` entries already include the `basePath` prefix, the URL is built with two prefixes, meaning that every request returns 404.


For example, this results in a duplicated prefix: `basePath: "/v1/collections/cards"
paths:
  /v1/collections/cards/query:   # already starts with basePath`

You can fix it by setting `basePath` to `/`: `basePath: "/"
paths:
  /v1/collections/cards/query:`

Or by removing the prefix from each path: `basePath: "/v1/collections/cards"
paths:
  /query:`

 This affects Swagger 2.0 only. OpenAPI 3.x uses `servers` instead of `host` and `basePath`.


#### For a file upload, do the server URLs need to be absolute?

 Yes. Files you upload must contain absolute server URLs. When Burp fetches the definition from a URL instead, it can resolve relative server entries against that URL.


## API authentication

 Both products support basic authentication, Bearer tokens, and API keys or custom tokens, each with fixed or dynamic credentials. Burp Suite DAST also has a dedicated OAuth 2.0 Client Credentials option.

#### What authentication does Burp support for API scans?

| **Method**  | **How it works**
| Basic authentication | Burp sends a static username and password.
| Bearer token (fixed) | Burp sends a long-lived or pre-generated token that you provide.
| Bearer token (dynamic) | Burp fetches a token from your token endpoint at a set interval and uses the latest one.
| API key / custom token (fixed) | Burp sends a key or token that you provide in a header, query parameter, or cookie.
| API key / custom token (dynamic) | Burp fetches the key or token from your token endpoint at a set interval.
| OAuth 2.0 Client Credentials (Burp Suite DAST only) | Burp obtains an access token using the `client_credentials` grant and refreshes it at a set interval. Burp Suite Professional has no dedicated option for this grant, so use dynamic Bearer-token authentication instead.
| Other OAuth 2.0 flows | Use dynamic Bearer-token authentication against your token endpoint.

 With dynamic authentication, Burp calls your token endpoint at intervals you set, extracts the token from the response (using JSON dot-notation or XPath), and uses it on every scan request until the next refresh. Burp refreshes on the interval rather than in response to a 401, so you should set the interval shorter than the token's lifetime. About 80% of token lifetime gives you a safe margin, for example 48 minutes for a 60-minute token.


 For the exact fields and steps in each product, see [Configuring API authentication](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication) for Burp Suite DAST, or [Configuring authentication for API scans](https://portswigger.net/burp/documentation/desktop/running-scans/api-scans/authentication) for Burp Suite Professional.


#### Note

 Configuring your identity provider correctly is your responsibility. PortSwigger Support can help with Burp's authentication behavior, but can't debug your identity provider setup.


#### How do I authenticate with the OAuth 2.0 password grant (ROPC)?

 Use dynamic Bearer-token authentication. In the Resource Owner Password Credentials (ROPC) flow, the token endpoint expects a `client_secret` as well as a username and password. Burp Suite DAST's OAuth 2.0 Client Credentials option has no fields for a username and password, and Burp Suite Professional has no Client Credentials option, so both products use the dynamic Bearer token for this flow.


For example, suppose your token request looks like this: `curl --location 'https://auth.example.com/connect/token' \
     --header 'Content-Type: application/x-www-form-urlencoded' \
     --data-urlencode 'grant_type=password' \
     --data-urlencode 'client_id=my-client' \
     --data-urlencode 'client_secret=s3cret' \
     --data-urlencode 'username=alice' \
     --data-urlencode 'password=hunter2'`

To set up the dynamic Bearer token:

1.

Set **Authentication service URL** to `https://auth.example.com/connect/token`.
1.

Set **Method** to `POST`.
1.

Add a header of `Content-Type: application/x-www-form-urlencoded`.
1.

Set **Body** to the following parameters, as a single line: `grant_type=password
&client_id=my-client
&client_secret=s3cret
&username=alice
&password=hunter2`
1.

Set **Token location** to `access_token`, for a response such as: `{
  "access_token": "...",
  "token_type": "Bearer",
  "expires_in": 3600
}`

If the token is nested, use the full dot-path, for example `data.access_token`.
1.

Set a refresh interval shorter than the token's lifetime.

 Burp calls the token endpoint on the interval you set, extracts the access token, and adds an `Authorization` header with the value `Bearer <token>` to every scan request. To adapt this for a custom OAuth-style flow, change the request body and headers to match your token endpoint.


#### How do I authenticate with the OAuth 2.0 Authorization Code flow?

 The Authorization Code flow involves a browser redirect and user consent, and Burp can't automate that interactive step directly. This also applies to the PKCE variants used by single-page apps, and to the pattern used by Microsoft Entra ID, Azure AD, and MSAL.


 The recommended approach is to sign in interactively once, then have Burp exchange the refresh token you capture for a fresh access token on a schedule.


To set up the sign-in:

1.

Sign in to the application with a test account that has the scopes your API needs.
1.

Capture the refresh token returned alongside the access token. Depending on your identity provider, it may come from a callback URL, a token response in the browser's network tab, or a known storage location such as MSAL session storage.
1.

Confirm the refresh token lasts long enough to cover your scan. Many identity providers issue refresh tokens valid for days or weeks. Some issue them only when you request the `offline_access` scope.

Your identity provider's refresh request looks something like this: `curl --location \
     'https://login.microsoftonline.com/{tenant-id}/oauth2/v2.0/token' \
     --header 'Content-Type: application/x-www-form-urlencoded' \
     --data-urlencode 'grant_type=refresh_token' \
     --data-urlencode 'client_id=my-spa-client-id' \
     --data-urlencode 'refresh_token=<captured-token>' \
     --data-urlencode 'scope=api://my-api/.default offline_access'`

To set up the dynamic Bearer token:

1.

Set **Authentication service URL** to your identity provider's token endpoint, for example `https://login.microsoftonline.com/{tenant-id}/oauth2/v2.0/token`.
1.

Set **Method** to `POST`.
1.

Add a header of `Content-Type: application/x-www-form-urlencoded`.
1.

Set **Body** to the following parameters, as a single line: `grant_type=refresh_token
&client_id=my-spa-client-id
&refresh_token=<captured-token>
&scope=api://my-api/.default offline_access`
1.

Set **Token location** to `access_token`.
1.

Set a refresh interval shorter than the access token's lifetime.

Watch for these issues:

-

Some identity providers, including Microsoft Entra ID in certain configurations, return a new refresh token in each response, invalidating the old one. In this case, the refresh-token-in-body approach stops working after the first refresh. To mitigate this, disable rotation for the test client, use a long-lived service-account credential instead, or test the affected endpoints manually.
-

Dynamic authentication fails if the refresh token expires partway through a long scan. Choose a test client whose refresh-token lifetime exceeds your longest expected scan, or keep scans short enough to finish within one refresh-token window.
-

If your sign-in enforces multi-factor authentication (MFA), you may not be able to automate it. Use a test account exempt from MFA, or a service principal designed for non-interactive use.

 If none of these approaches work for your environment, obtain a fresh long-lived bearer token manually before each scan and configure it as a fixed Bearer token.


#### Burp detected authentication from my definition. Can I still add my own?

 Yes. When your definition declares authentication that Burp can use, such as OpenAPI `securitySchemes` or Postman `auth` blocks, Burp labels it **Detected**. You can add your own authentication for all other endpoints. Your manual authentication applies only to endpoints that a **Detected** method doesn't already cover, so the two can't conflict.


#### My file loads, but my security scheme isn't showing as "Detected"

 If a `security` block (top-level or per-operation) references a scheme name that isn't defined under `components.securitySchemes`, Burp ignores that reference. The definition still parses, but the scheme doesn't appear as **Detected** authentication for the affected endpoints. This is easy to miss, because Burp doesn't report an error.


For example, this definition references `cookieSession`, but never defines it. The file loads, but Burp drops the scheme: `security:
  - cookieSession: []
paths:
  /users:
    get: ...
# no components.securitySchemes block`

This definition defines the scheme correctly, so Burp picks it up as **Detected** authentication: `security:
  - cookieSession: []
components:
  securitySchemes:
    cookieSession:
      type: apiKey
      in: cookie
      name: SESSION`

 If authentication you expect to see under **Detected** isn't showing up, check that every name referenced in a `security` entry has a matching definition under `components.securitySchemes`.


## Postman collections

 A Postman Collection carries complete requests, so Burp sends the bodies it finds rather than generating them from a schema.

#### Which Postman version does Burp support?

 Burp supports Postman Collections exported in `v2.1.0`. Burp rejects `v1` and `v2.0`, so re-export in `v2.1.0` if you need to.


#### Can Burp use my Postman variables and environment files?

 Burp supports variables (the `{{variable}}` syntax) and substitutes them from the collection. If a variable has no value, Burp generates one. Burp does not run Postman scripts.


 Environment support differs by product:


-

DAST When you upload a Postman Collection, you can also upload a Postman environment file. Burp Suite DAST merges the environment variables with your collection, so you don't have to merge them by hand. For more information, see [Adding a single API](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/adding-single-api).
-

Professional Burp Suite Professional doesn't accept a separate environment file. Define the values you need as collection variables in Postman's **Variables** tab, or edit the exported `JSON`, before you upload the collection.

#### Why does my collection fail with a "Body mode not understood" error?

 Postman supports several body modes, some of which Burp can't scan:

| **Body mode**  | **Supported**
| `raw` (any content type) | Yes
| `urlencoded` | Yes
| `formdata` | No
| `file` | No
| `graphql` | No. Scan GraphQL directly against the endpoint (see [GraphQL](https://portswigger.net/burp/documentation/scanner/api-scanning-faq#graphql)).

 For most form payloads, change the request body in Postman from `form-data` to `x-www-form-urlencoded` before you export.


 Support depends on the body mode rather than the content type. In `raw` mode, Burp sends the body exactly as the collection provides it, so content types that Burp can't generate from an OpenAPI schema still work here. For example, a `text/plain` body fails in an OpenAPI definition but scans in a Postman Collection that uses `raw` mode.


#### Can I scan file-upload endpoints from a Postman collection?

 No. Burp's Postman parser doesn't support file-upload endpoints, whether they use Postman's `file` body mode or `formdata` with a file part. In this case, Burp displays a parser error and doesn't list the affected requests as endpoints. To scan a file-upload endpoint you can either:


-

Record a browser session in Burp that performs the upload, then scan from that session. Burp can replay and audit the multipart requests it captured.
-

Run a Crawl and Audit against the running application. If the upload form is reachable from a crawled page, Burp can discover and scan it without an API definition for that endpoint.

## SOAP and WSDL

 Burp scans a SOAP service from the WSDL file that describes it.

#### Which SOAP versions does Burp support?

 Burp supports SOAP 1.1 and 1.2, and determines the content type for each endpoint from the WSDL. It supports both Document and RPC binding styles.


#### How do I authenticate SOAP scans?

 A WSDL doesn't carry authentication details, so Burp has nothing to detect. Configure any authentication you need manually in the scan settings.


#### Can Burp use a multi-file WSDL?

 Burp doesn't follow external `<xsd:import>` or `<wsdl:import>` references. Bundle the WSDL into one self-contained file before you upload it.


## GraphQL

#### How does Burp scan GraphQL APIs?

 Burp doesn't use a schema file for GraphQL. It scans through introspection: you point Burp at the GraphQL endpoint, and it queries `__schema` to discover the available queries and mutations. This affects what you need to set up, and what Burp covers:


-

Your target must have introspection enabled. If your production environment disables it, use a staging environment that allows it.
-

Burp needs to find only the single GraphQL endpoint, because every GraphQL operation uses the same one.
-

Burp scans queries and mutations, but doesn't scan subscriptions.

 For the full GraphQL scan process, see [Requirements for API scanning](https://portswigger.net/burp/documentation/scanner/api-scanning-reqs#graphql-definition-requirements).


#### What if Burp can't find the GraphQL endpoint when crawling?

 When you select **Test common GraphQL endpoints**, Burp tries a list of standard endpoint suffixes, such as `/graphql` and common variants. If your endpoint is non-standard, you may need to add it manually.


## Large or complex definitions

 Very large or deeply cross-referenced definitions can slow the parser or make it fail.

#### My definition is large, or the upload seems to hang. What can I do?

 This usually comes from one of two patterns:


-

Your `YAML` file is over the 3 MB size limit. See [Why does my definition fail to parse?](https://portswigger.net/burp/documentation/scanner/api-scanning-faq#why-does-my-definition-fail-to-parse)
-

Your definition has deeply circular `$ref`s, which can stall the parser. OpenAPI files generated for Microsoft Dynamics 365 are one example of this, because every entity references every other entity through navigation properties. With hundreds of interlinked schemas, Burp can exhaust available memory before it finishes parsing.

 For a circular-`$ref` definition, pre-process the file so that real request and response fields stay intact but cross-entity navigation references become a generic `{"type": "object"}`. This keeps your path coverage and request-body fields, and simplifies only the links between entities. Several open-source OpenAPI tools can do this.


#### Can I scan a definition generated by FastAPI or Spring?

 Many frameworks serve the live OpenAPI definition on a well-known path. FastAPI, for example, serves it at `/openapi.json`. If you don't have a file, point Burp at that URL, or download the `JSON` and upload it. Both options work in Burp Suite Professional and Burp Suite DAST.


 If Burp can't reach the definition URL, run a Crawl and Audit against the application instead, then re-scan any API endpoints Burp discovers in the site map.


## Coverage and reporting

 After a scan, the event log and scan summary show what Burp covered and what it skipped.

#### How do I tell which operations Burp scanned and which it skipped?

 Check the scan event log. The scan summary shows the count of endpoints scanned. The level of detail in the log depends on the product and the definition format:


-

Professional Burp records each skipped operation with a reason, for example "Skipping location in API definition because it contains a Header parameter...". If coverage is lower than you expect, search the log for "Skipping".
-

DAST For OpenAPI definitions, Burp records each skipped operation in the same way. For Postman Collections, SOAP WSDLs, and GraphQL APIs, Burp reports a single event once parsing finishes, telling you how many operations it skipped but not which ones.

 Composition-schema headers are a common cause of operations disappearing. If Burp Suite DAST tells you that operations were skipped but not which ones, check your definition against the patterns in [Parser errors and unsupported features](https://portswigger.net/burp/documentation/scanner/api-scanning-faq#parser-errors-and-unsupported-features), or upload the same definition to Burp Suite Professional to get the per-operation detail.


#### Does Burp Scanner test for API-specific vulnerabilities like the OWASP API Top 10?

 Burp Scanner parses your API definition and runs its standard web vulnerability checks against the discovered endpoints. It doesn't include dedicated checks for API-specific issue classes such as broken object-level authorization, mass assignment, or API rate-limiting. If you need OWASP API Top 10 coverage specifically, supplement scanning with manual testing.


#### Can Burp scan Server-Sent Events (SSE) or streaming-response endpoints?

 No. Burp Scanner can't scan Server-Sent Events (SSE) or streaming-response endpoints, including those used by many AI chat and real-time data APIs. Burp classifies these as `STREAMING_RESPONSE`, a scan error type rather than a supported response mode, because the scanner expects a complete request-response cycle and SSE responses are long-lived and chunked.


 If your application exposes both SSE and non-SSE endpoints, Burp scans the non-SSE ones normally and skips the SSE ones with an error. To test SSE endpoints, use Repeater or Intruder manually.


## Troubleshooting checklist

 When an API scan doesn't behave as you expect, work through this checklist:

-

Check that Burp supports your definition format and version.
-

Check that your file is under the size limit, and convert it to `JSON` if it's large.
-

Bundle your definition into one file, with no external `$ref`s.
-

Make your server and host URLs absolute for file uploads, and reachable from Burp.
-

Give every header parameter a primitive schema, rather than `type: object` or a composition schema.
-

Give every body schema `type: object` with `properties`, rather than a bare primitive.
-

For Swagger 2.0, remove any `basePath` prefix that your paths repeat.
-

For Postman, export the collection as `v2.1.0` with no `formdata`, `file`, or `graphql` body modes.
-

Configure authentication manually for SOAP, and either add credentials to a Detected method or add your own for other formats.
-

Check the scan event log for skipped operations.

 If you've worked through the list and the scan still isn't right, contact support with your API definition file (or a redacted version), the scan event log, and a short description of what you expected and what happened.

#### Related pages

-  [Requirements for API scanning](https://portswigger.net/burp/documentation/scanner/api-scanning-reqs) - the criteria a definition must meet before Burp can scan it.

-  DAST [Adding new API definitions](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis) - scanning APIs in Burp Suite DAST.

-  Professional [Scanning APIs](https://portswigger.net/burp/documentation/desktop/running-scans/api-scans) - API-only scans in Burp Suite Professional.

-  DAST [Configuring API authentication](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication) - authentication fields and steps in Burp Suite DAST.

-  Professional [Configuring authentication for API scans](https://portswigger.net/burp/documentation/desktop/running-scans/api-scans/authentication) - authentication fields and steps in Burp Suite Professional.

-  [GraphQL API vulnerabilities](https://portswigger.net/web-security/graphql) - testing GraphQL APIs, in the Web Security Academy.
