> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis

DAST

# Scanning APIs

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This section explains how to scan APIs in Burp Suite DAST. For information on how to add APIs, see [Adding APIs to DAST](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis).


#### Note

 You can add as many APIs as you like to Burp Suite DAST but for scans to work correctly, you need to configure your network and firewall settings.
 For more information, see
 [Configuring network and firewall settings for a site](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/configuring-network-and-firewall-settings).


## Adding API definitions

 You can add API definitions by uploading a file or providing a URL. The supported formats are:


-
 Postman Collection

-
 OpenAPI definition file in JSON or YAML format

-
 SOAP WSDL

-
 GraphQL (URL only, using introspection)


 For Postman Collections, you can also upload a Postman environment file to automatically merge environment variables with your collection. This removes the need to manually merge variables and speeds up your setup process.


 For GraphQL APIs, make sure introspection is switched on. For more information, see [GraphQL definition requirements](https://portswigger.net/burp/documentation/scanner/api-scanning-reqs#graphql-definition-requirements).


#### Note

 We fully support OpenAPI 3.1 and provisionally support OpenAPI 3.2.


## Choosing how to add APIs

 You can add API sites individually or in bulk:

-

**Add a single API** - Create one API site at a time by uploading a file or providing a URL. Use this approach when you need to configure each API individually or when onboarding a small number of APIs.
-

**Bulk upload APIs** - Create multiple API sites in one operation. Use this approach when onboarding large numbers of APIs with shared configuration settings. This speeds up the process and helps ensure consistency across your API estate.

 For more information, see:

-

[Adding a single API](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/adding-single-api)
-

[Bulk uploading APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/bulk-uploading-apis)

## Managing authentication for API sites

 When you add an API definition, Burp Suite DAST automatically detects authentication schemes. You don't have to provide credentials immediately.

 To add authentication credentials after creating a site:

1.

Go to **Sites** and select your API site.
1.

Select the **Details** tab and click  **Edit**.
1.

Under **API definition**, select the **Authentication** tab. Add any credentials that are shown as missing.
1.

Click **Save**.

## Optional settings for your API

 When you add a new API site, you can configure the following additional settings:

-

Scan configuration
-

Connections
-

Headers and cookies
-

Extensions
-

Scanning pool
-

Burp AI and automation
-

Notifications

For more information on configuring the optional settings for your API,
 see [Configuring site settings](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings).

#### Related pages

-

[Discovering APIs from integrations](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations)
-

[Adding a single API](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/adding-single-api)
-

[Bulk uploading APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/bulk-uploading-apis)
-

[Configuring API authentication](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication)
-

[Viewing and configuring API endpoints](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configure-endpoints)
