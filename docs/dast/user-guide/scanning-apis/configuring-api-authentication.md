> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication

DAST

# Configuring API authentication

-

**Last updated: ** September 3, 2026
-

**Read time: ** 5 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

## Host credentials vs API authentication

 When working with API definitions, you may need to provide two different types of credentials:

-

**Host credentials** - Optional credentials used to access the API definition file if it's hosted on a server that requires authentication. These are only used when fetching the definition file from a URL.
-

**API authentication** - Credentials used to authenticate requests to the API endpoints during scanning. These are configured as described on this page.

 This page explains how to configure API authentication for scanning your API endpoints.

## Supported authentication types

 Burp Suite DAST supports the following authentication:

-  **Basic** - Enter a username and password.

-  **Bearer auth** - Adds an access token that's sent in the Authorization header.

-  **API key / Custom token** - Adds an API key or access token in a custom location.

-  **OAuth 2.0 client credentials** - Automatically obtains and refreshes access tokens using OAuth 2.0.


 If you upload a Postman file that contains these authentication types, Burp Suite DAST extracts the credentials automatically. You can see the extracted information in the
 **Authentication** tab.

 You can also use dynamic tokens. These tokens have a limited lifespan. Burp Suite DAST enables you to fetch refreshed tokens automatically.

#### Note

 For security reasons, API definitions should include authentication schemes but not the associated credentials. For example, a definition can define that a particular API key is needed, but it must not provide the API key.


 This means that you need to add credentials for any detected schemes manually. Schemes that have been detected but not yet populated with credentials have a red notification dot next to them. To add a credential to a scheme, click its  pencil icon.


## Authentication status overview

 You do not have to supply every credential before you save a site. Burp Suite DAST warns you as you create or edit a site that has authentication schemes without credentials. You can add the missing credentials at any time.

## Adding basic authentication

 To add basic authentication:

1.
 Select the **Authentication** tab.

1.
 Click **Add API authentication** to display the **Add authentication** dialog.

1.
 Select **Basic**.

1.
 Enter the **Label**, **Username**, and **Password**.

1.
 Click **Save**.


## Adding Bearer token authentication

 Bearer auth adds an access token that is sent in the Authorization header.

#### Note

 Burp Suite DAST supports OAuth 2.0 client credentials directly. You do not need to configure a dynamic Bearer auth token to fetch OAuth tokens, because Burp Suite DAST obtains and refreshes them for you. For more information, see [Adding OAuth 2.0 client credentials authentication](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication#adding-oauth-2-0-client-credentials-authentication).


 To add Bearer token authentication:

1.
 Select the **Authentication** tab.

1.
 Click **Add API authentication** to display the **Add authentication** dialog.

1.
 Select **Bearer auth**.

1.

For **Fixed** tokens:

  -
 Set the **Token type** to **Fixed**.

  -
 Enter the **Label** and **Token**.

  -
 Click **Save**.


1.

For **Dynamic** tokens:

  -
 Set the **Token type** to **Dynamic**.

  -
 Enter a **Label** for the token (for example, "Auth Token").

  -
 Configure the token request. For more information, see [Configuring a dynamic token request](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication#configuring-a-dynamic-token-request).


## Adding API key / custom token authentication

 To add credentials for an API key, or add a token in a custom location:

1.
 Select the **Authentication** tab.

1.
 Click **Add API authentication** to display the **Add authentication** dialog.

1.
 Select **API key / Custom token**.

1.

For **Fixed** tokens:

  -
 Set the **Token type** to **Fixed**.

  -
 Enter the **Label**, **Location**, and **Key / token**.

  -
 Click **Save**.


1.

For **Dynamic** tokens:

  -
 Set the **Token type** to **Dynamic**.

  -
 Enter a **Label** for the token (for example, "Auth Token"), and choose a location from the **Add to** drop-down list to specify where the token should be added (for example, Header, Query Parameter, or Cookie).

  -
 Configure the token request. For more information, see [Configuring a dynamic token request](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication#configuring-a-dynamic-token-request).


## Adding OAuth 2.0 client credentials authentication

 To add OAuth 2.0 client credentials authentication:

1.
 Select the **Authentication** tab.

1.
 Click **Add API authentication** to display the **Add authentication** dialog.

1.
 Select **OAuth 2.0 client credentials**.

1.
 Enter a **Label** for the authentication.

1.
 Enter your **Token URL**. This is the OAuth 2.0 token endpoint.

1.
 Enter your **Client ID**.

1.
 Enter your **Client secret**.

1.

Expand **Advanced settings** to configure optional settings:

  -  **Scope** - Enter a space-separated list of OAuth 2.0 scopes if required by your API.

  -  **Authentication method** - Select how credentials are sent. Choose **POST (credentials in body)** or **HTTP Basic**.

  -  **Refresh interval (seconds)** - Set how often to refresh the token. The default is 3300 seconds (55 minutes). You can set a value between 60 and 86400 seconds.


1.
 Click **Save**.


## Configuring a dynamic token request

 Dynamic tokens have a limited lifespan. Burp Suite DAST requests a new token from your authentication service at an interval you set, so your scans always use a valid token.

 Bearer auth and API key / custom token authentication share the same token request settings. To configure the request:

1.
 Enter the **Authentication service URL** (for example, `https://api.example.com/auth`) and **Method** (for example, `POST` or `GET`) for the request to retrieve the token.

1.
 If necessary, expand **Additional headers** and enter the **Name** and **Value** for each header (for example, `Content-Type: application/x-www-form-urlencoded`). Click  to add more headers.

1.

In the **Body** field, if necessary, enter the body of the request. For example, for a JSON request:

`{
   "username": "user123",
   "password": "securePassword"
}`

For `application/x-www-form-urlencoded` content type, you must set the `Content-Type` header in **Additional headers**, then enter the body accordingly:

`username=user123&password=securePassword`
1.
 Enter a value for how often the token should be refreshed in the **Re-fetch every** field (for example, every 30 minutes). This ensures that the scanner always uses a valid token for API requests.

1.

In the **Token location** field, enter the location in the response body where the token will be located. For **JSON** responses, use the field name or dot separated names. For **XML** responses, use XPath. Leave blank to use the entire response body as the token.

For example, if the response body contains:

`{
   "token": "abc123xyz",
   "expires_in": 3600
}`

You would enter `token` as the token location.
1.
 Click **Save**.


## Editing or deleting authentication methods

 To edit an existing authentication method, click its  pencil icon.

 To delete an existing authentication method, click its  trash icon.

#### Note

 In order to modify authentication details for an API site after the site has been saved, you need `Edit site application logins` permission. This includes changing the specification upload method between a URL and a local file. Note that admin users have this permission by default.


 If you have `View site application logins` permission but not `Edit site application login details` permission, you can see details of the authentication methods used in the specification and their credentials. However, you can't edit any of them, add new authentication, amend the selection of endpoints to scan, or change the API definition file or URL.


## Scanning with incomplete authentication

 You will see a warning if you create scans that include sites with incomplete authentication. You can still run your scan, but areas that require authentication may not be tested.

#### Related pages

-  [Managing scheduled scans](https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans/manage-scheduled-scans) - explains how to schedule scans for your new site.

-  [Defining scan configuration for a site](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations) - explains how to create and work with scan configurations.

-  [Configuring site settings](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings) - explains the optional scan settings you can configure for a site.

-  [Configuring your environment network and firewall settings](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/network-firewall-config).

-  [Burp Scanner built-in configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations/burp-scanner-built-in-configs) - reference information on Burp Scanner's built-in scan configurations.
