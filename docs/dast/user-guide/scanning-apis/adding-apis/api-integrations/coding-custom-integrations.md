> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/coding-custom-integrations

DAST

# Coding custom integrations with GraphQL API

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If Burp Suite DAST doesn't have a built-in integration for the platform where your APIs are stored, you can use the GraphQL API to push API definitions into API finder from any source. You can use this to discover OpenAPI and SOAP WSDL definitions, and Postman Collections.

#### Note

To help you get started, you can access some example scripts in the [API finder examples](https://github.com/PortSwigger/dast-api-finder-examples) repository on GitHub.

## When to use a custom integration

 Burp Suite DAST provides built-in connectors for Amazon API Gateway, Azure API management, and Google Apigee. If your APIs are managed on either of these platforms, you can connect directly to API finder without any scripting.

 Use a custom integration if your APIs are stored somewhere that doesn't have a built-in connector. For example:

- API definitions stored in Git repositories alongside application code.
- Internal service catalogs or custom-built API registries.
- Any other source that your team maintains separately from an API management platform.

## Creating a custom integration

 To create an integration using the GraphQL API:

1.

Create an API user. For more information, see [Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user).
1.

Add the API user to the **API Uploaders** group. For more information, see [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control).
1.

Write a script using the GraphQL API, to push APIs into API finder. For more information, see [Getting started with the GraphQL API](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/graphql-api/getting-started).
1.

Use API finder to review APIs, and create sites for them. For more information, see [Creating sites for added APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/creating-sites).
1.

When your APIs change, run your script again with the same `unique_id`. This tells Burp Suite DAST to update the API, instead of creating a new one.

#### Related pages

-

[Discovering APIs from integrations](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations)
-

[Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user)
-

[GraphQL API](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/graphql-api)
-

[dast-api-finder-examples on GitHub](https://github.com/PortSwigger/dast-api-finder-examples)
