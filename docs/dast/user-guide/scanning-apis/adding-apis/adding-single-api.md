> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/adding-single-api

DAST

# Adding a single API

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This page explains how to create a site for a single API. You can add an API definition by uploading a file or providing a URL.

## Adding API definitions by uploading a file

 When you upload an API definition file, Burp Suite DAST uses that version for every scan until you upload a new one. You can't upload GraphQL API definition files.

 To add an API definition by uploading a file:

1.

Go to **Sites** and select **Add a new site**.
1.

Select **API** from the **Site type** panel.
1.

Enter a unique **Site name**.
1.

To add the site to an existing folder, select from the **Site folder** drop-down menu.
 Leave the field blank to create the site at the top level of the site tree.
1.

Click **Select file** and select the definition file to upload.
1.

If you upload a Postman Collection, you are given the option to upload a Postman environment file. Burp Suite DAST merges the environment variables with your Postman Collection. Click **Add environment file** and select your environment file.

## Adding API definitions by providing a URL

When you provide a URL, Burp Suite DAST uses the latest version of the API definition for every scan. If the definition file is on a server that requires authentication, you can add host credentials. You can validate that the URL is reachable and the file can be parsed.

 To add an API definition by providing a URL:

1.

Go to **Sites** and select **Add a new site**.
1.

Select **API** from the **Site type** panel.
1.

Enter a unique **Site name**.
1.

To add the site to an existing folder, select from the **Site folder** drop-down.
 Leave the field blank to create the site at the top level of the site tree.
1.
 Click **Host URL** and enter the URL where your API definition file is hosted.

1.

(Optional) If your API definition file is on a server that requires authentication, click **Add host credentials** and enter the credentials.
1.
 Click **Validate** to confirm that Burp Suite DAST can reach the URL and parse the API definition.
 If successful, the **Endpoints** tab lists the parsed endpoints.


#### Note

 If you provide a URL for a Postman Collection, Burp Suite DAST also extracts the credentials for detected authentication schemes. They are shown in the **Authentication** tab.


 If you manually edit the credentials for your site, they may be overwritten if you later click **Validate**.


## Saving and scanning your site

 Once you've finished creating your site, click **Save**.

 Burp Suite DAST adds the new API to the site tree and prompts you to schedule a scan.

#### Related pages

-

[Scanning APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis)
-

[Creating sites for added APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/creating-sites) - If you use integrations to discover APIs, you can create sites directly from API finder.
-

[Configuring API authentication](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configuring-api-authentication)
-

[Viewing and configuring API endpoints](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/configure-endpoints)
