> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/azure

DAST

# Integrating with Azure API Management

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can integrate Burp Suite DAST with Azure API Management to automatically discover APIs published in your Azure API Management instance.

#### Note

 This integration gives Burp Suite DAST read-only access to Azure API Management. No APIs are created or scanned during the setup process. Once the connection is established, discovered APIs are shown in **API finder**.


 Burp Suite DAST only discovers APIs that have a downloadable OpenAPI specification. APIs without a specification, such as SOAP-only APIs, are not discovered.


## Prerequisites

1.
 You have **Modify Settings** permission in Burp Suite DAST.

1.
 You have registered an application in Azure Active Directory and have a client ID and client secret for it.

1.
 The application has the built-in **API Management Service Reader** role assigned at the scope of your API Management instance.


## Connecting to Azure API Management

1.
 Go to **Settings  > Integrations**.

1.
 On the **Azure API Management** tile, click **Configure**.

1.
 Click **Add integration**. The **Azure API Management connection details** dialog opens.

1.
 Enter a name in the **Integration name** field.

1.
 Enter your connection details:


  - **Tenant ID** - The directory ID of the Azure Active Directory tenant where your application is registered.
  - **Client ID** - The application ID of your registered application.
  - **Client secret** - A client secret you have generated for your application.
  - **Subscription ID** - The ID of the Azure subscription containing your API Management instance.
  - **Resource group** - The name of the resource group containing your API Management instance.
  - **APIM instance name** - The name of your API Management instance.

1.
 Enter a schedule for how often you want Burp Suite DAST to check the instance for updates. The minimum interval is 15 minutes.

1.
 Click **Save**.


 Once Burp Suite DAST connects to Azure successfully, a confirmation message appears. The connection is then listed on the **Integrations** page, where you can see its status and how many APIs have been discovered.

To manage the APIs discovered by the integration:

1.
 Click **View discovered APIs** to go to **API finder**.

1.
 Click on an API to review its status.

1.
 Select multiple APIs and click **Create sites** to create an API site for each API. For more information, see [Creating sites for added APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/creating-sites).


#### Note

 Azure client secrets expire after a maximum of 24 months. When a secret expires, the integration stops working until you provide a new one. Make a note of the expiration date so you can replace the secret before it expires.


 When a scheduled update discovers changes to your APIs, you will see a red notification dot in the **API finder** tab. For more information, see [Updating your API sites](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/updating-sites)

#### Related pages

-

[Discovering APIs from integrations](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations)
