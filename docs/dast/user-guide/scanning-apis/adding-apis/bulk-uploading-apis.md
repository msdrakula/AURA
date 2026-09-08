> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/bulk-uploading-apis

DAST

# Bulk uploading APIs

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can create multiple API sites in one operation. Each file or URL creates its own site. Bulk uploading speeds up the process of onboarding large numbers of APIs and enables consistent configuration across your API estate.

 You can bulk import APIs in two ways:

-

From the **Sites** menu, select **Bulk upload APIs** to import API definition files or URLs.
-

From the **API finder** menu, select one or more APIs and click **Create sites**. For more information, see [Creating sites for added APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/creating-sites).

## Before you start

 Make sure you have the API definition files ready or know the URLs where they are hosted. You can upload a mix of Postman Collections, OpenAPI definitions, and SOAP WSDL files.

#### Note

 You must have permission to scan all URLs included in the scope of the imported sites.


## How to bulk upload APIs

 To bulk upload your APIs, complete the following steps:

### Update behavior

 Burp Suite DAST only shows this step when you create sites from the **API finder** menu.

1.

Choose how API updates are applied:

  -

**Always review changes before updating** - The notification dot is shown on the **API finder** menu when changes to the API definitions are detected. You can review the changes and update your sites manually.
  -

**Automatically keep APIs up to date** - New versions of API definitions are applied to your sites automatically.

1.

Click **Next**.

### Import type

 Burp Suite DAST only shows this step when you create sites from the **Sites** menu.

1.

Choose how to add your API definitions:

  -

**Upload files** - Select API definition files from your local machine. Click **Upload more files** to add additional files.
  -

**Host URLs** - Paste a list of links to the URLs of your API definition files.

1.

If you want to remove any of the uploaded files, select the tick boxes and click  **Delete**.
1.

Click **Next**.

### Select methods

 Burp Suite DAST only shows this step when you create sites from the **API finder** menu.

1.

Choose the request methods to enable for this import. These also apply endpoints discovered in future automatic updates:

  -

**GET**
  -

**POST**
  -

**PUT**
  -

**DELETE**
  -

**Other** - Includes PATCH, HEAD, OPTIONS, and any non-standard methods.

1.

Click **Next**.

### Choose folder

1.

Select a destination folder for the imported sites.
1.

Click **Next**.

### Network

1.

Configure the network settings that apply to all imported sites:

  -

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted **Scanning pool** - Select which scanning pool to use.
  -

**Upstream proxy servers** (optional) - [Add upstream proxy server rules](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/upstream-proxy-servers) if needed.
  -

**Headers and cookies** (optional) - [Add custom headers and cookies](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/headers-and-cookies) to requests.

1.

Click **Next**.

### Review upload

 This step only appears when you create sites from the **Sites** menu.

1.

Review the sites that will be created. Make any necessary changes.
1.

To remove one or more sites, select their tick boxes and click  **Delete**.
1.

If any sites failed parsing, download the log to view more details about the errors.
1.

Click **Next**.

### Authentication

 You can view detected authentication types for your sites, or add them. Any authentication you add manually applies to all sites in this bulk upload.

To add authentication credentials:

1.

Click **Add more**.
1.

Select **Basic**, **Bearer auth**, **API key / Custom token**, or **OAuth 2.0 client credentials**.
1.

Enter the details for the chosen authentication method.
1.

Click **Save**.
1.

To edit authentication credentials, click the  pencil icon.
1.

Click **Next**.

### Scan configuration

1.

Choose the scan configuration:

  -

**Default configuration** - Uses the default scan configuration for API sites. If the parent folder has scan configurations applied, your sites inherit those configurations instead.
  -

**Use a custom configuration** - Select a custom scan configuration.

1.

Under **Extensions**, select the BChecks and custom extensions you want to use. Click **Select all BChecks** to use all available checks.
1.

Click **Import**.

 Burp Suite DAST creates a separate site for each API definition and adds them to the site tree in the selected folder.

## After bulk upload

 After importing your APIs, you can:

-

Add or update authentication credentials for individual sites.
-

Fine-tune scan configurations for specific sites.
-

Schedule scans for your imported sites.

 For more information, see [Configuring site settings](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings).

#### Related pages

-

[Scanning APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis)
-

[Adding a single API](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/adding-single-api)
-

[Creating sites for added APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/creating-sites)
-

[Configuring API authentication](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication)
