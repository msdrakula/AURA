> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations

DAST

# Discovering APIs from integrations

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can connect Burp Suite DAST to third-party platforms to automatically discover APIs deployed in your environment. Discovered APIs appear in **API finder**, where you can review them and create scan sites.

 You can do this in the following ways:

-

[Integrating with AWS](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/aws) - Use the built-in integration for Amazon API Gateway.
-

[Integrating with Azure API Management](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/azure) - Use the built-in integration for Azure API Management.
-

[Integrating with Google Apigee](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/apigee) - Use the built-in integration for Google Apigee.
-

[Coding custom integrations](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/coding-custom-integrations) - Write a script to push APIs into API finder from any other source using the GraphQL API.

 Once your integration is set up, you can review and manage the APIs that appear in API finder:

-

[Creating sites for added APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/creating-sites) - Review APIs in API finder and create scan sites for the ones you want to scan.
-

[Updating your API sites](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/updating-sites) - Keep your API sites up to date when your integration discovers new endpoints.

#### Related pages

-

[Adding APIs to DAST](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis)
