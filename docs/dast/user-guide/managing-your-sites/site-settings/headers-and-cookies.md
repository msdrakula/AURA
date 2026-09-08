> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/headers-and-cookies

DAST

# Adding headers and cookies

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Custom headers and cookies enable you to authenticate or add required headers for your web applications. You can add custom headers and cookies to sites, folders, or subfolders.

## Adding custom headers

To add a custom header:

1. Select a site or folder, and then select the **Details** tab.
1. Click **Edit**.
1. Under **Scan settings**, go to the **Headers and cookies** tab.
1. Click **Add a header or cookie**.
1. Select **Add header**.
1. **Name** - Enter the header name (for example: `Connection` or `Authorization`).
1. **Value** - Enter the header value. For example:


  - For Connection: `keep-alive`
  - For Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`
  - For User-Agent: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)`

1. **Scope Prefix** - Enter `https://api.example.com/` or leave empty for all requests.
1. Click **Finish** to close the dialog box, then click **Save**.

#### Note

 You can also use a custom scan configuration to customize your User-Agent.


## Adding session cookies

To add a session cookie:

1. Select a site or folder, and then select the **Details** tab.
1. Click **Edit**.
1. Under **Scan settings**, go to the **Headers and cookies** tab.
1. Click **Add a header or cookie**.
1. Select **Add cookie**.
1. Complete the fields:


  - **Name**: `session` (the cookie name)
  - **Value**: `abc123def456ghi789` (anything that comes after `session` in the cookie)
  - **Scope Prefix**: `https://example.com/secure/` (for secure areas only)

1. Click **Finish** to close the dialog box, then click **Save**.

## Understanding scope prefix

Scope prefix controls where your header or cookie is applied:

- **With a prefix**: Header/cookie only applies to URLs starting with that exact prefix
- **Empty**: Header/cookie applies to all requests within your site scope

### Examples:

- `https://example.com/admin/` - Only admin section
- `https://api.example.com/v1/` - Only API v1 endpoints
- Empty - All requests in your site scope

 To add additional headers or cookies, click the add icon .

 To edit a header or cookie from the list, click the edit icon .

 To delete a header or cookie, click the trash icon .

#### Related pages

- [Setting the site scope](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/setting-site-scope)
- [Configuring site settings](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings)
