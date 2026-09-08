> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/onelogin

DAST

# Integrating SCIM using OneLogin

-

**Last updated: ** September 3, 2026
-

**Read time: ** 4 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 In this section, we'll guide you through the process of integrating SCIM with Burp Suite DAST using OneLogin as your identity provider (IdP).

## Prerequisites

-
 You already have your users set up in OneLogin.

-
 If you want to use SCIM in conjunction with SAML, you have already created a `SCIM Provisioner with SAML (SCIM v2 Core)` application in OneLogin and have [completed the SAML configuration](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml).


## Get your SCIM URL and API token

 You need to obtain the SCIM URL and API token. OneLogin uses these to communicate with Burp Suite DAST. The process is different for obtaining these for Cloud and self-hosted instances.

1.

Log in to Burp Suite DAST as an administrator.
1.

From the settings menu, select **Integrations**.
1.

On the SCIM tile, click **Configure**.
1.

Get your SCIM URL:

  -

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud Your SCIM URL is automatically generated and displayed on screen for you to copy.
  -

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted The base URL takes the following format:

`https://<host>:<port>/scim/v2`

The host is usually the same domain name or IP address as in the Burp Suite DAST web server URL. However, this may differ depending on your network infrastructure. Enter the port that you want to use for the SCIM URL. This should be a different port than the one you use for the web server URL so that you can configure separate firewall rules for this connection.

1.

Get your API key:

  -

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud Click **Generate API token**.
  -

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted Click **Save & generate API token**.

1.

When prompted, copy and save the new API token somewhere secure.

#### Note

 If you lose your API token, you can generate a new one by clicking **Regenerate API token** in the upper-right corner of the SCIM settings page.


## Upload a TLS certificate

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 For production use, we strongly recommend enabling TLS on the connection by uploading a `PKCS#12` certificate. Note that this must have the `.p12` file extension - certificates in `.pfx` format are not supported.

1.
 From the settings menu , select **Integrations**.

1.
 On the SCIM tile, click **Edit**.

1.
 Under **Configure SCIM**, activate the **Use TLS** toggle.

1.
 When prompted, upload your certificate and enter the certificate password.

1.
 Click **Save**.


## Configure the connection in OneLogin

 Once you've got your SCIM URL and generated an API token in Burp Suite DAST, you can use this information to configure the connection from OneLogin.

### Enter the connection details

1.
 In OneLogin, select the application that you created for Burp Suite DAST.

1.
 From the left-hand navigation menu, select **Configuration** and scroll down to the **API Connection** section.

1.
 In the **SCIM Base URL** field, enter your SCIM URL.

1.

 In the **Custom Headers** field, add the following header:
  `Content-Type: application/scim+json`
1.
 In the **SCIM Bearer Token** field, enter the API token that you copied from Burp Suite DAST.

1.

 In the **SCIM JSON Template** field, paste the following template:
  `{
    "schemas": [
        "urn:ietf:params:scim:schemas:core:2.0:User"
    ],
    "userName": "{$parameters.scimusername}",
    "name": {
        "givenName": "{$user.firstname}",
        "familyName": "{$user.lastname}"
    },
    "emails": [
        {
            "primary": true,
            "value": "{$user.email}",
            "type": "work"
        }
    ],
    "displayName": "{$user.display_name}"
}`
1.
 At the top of the **API Connection** section, under **API Status**, click the button to enable the connection.

1.
 Save your changes.


### Configure the parameters

1.
 In OneLogin, select the application that you created for Burp Suite DAST.

1.
 From the left-hand navigation menu, select **Parameters**.

1.
 Click the entry for the `NameID` parameter.

1.
 In the dialog that opens, change the value of this parameter to `Username`.

1.
 Save your changes.


## Enable SCIM provisioning

 Once you've successfully configured the SCIM connection between OneLogin and Burp Suite DAST, you can enable SCIM provisioning so that you can sync your users and groups.

1.
 In OneLogin, select the application that you created for Burp Suite DAST.

1.
 From the left-hand navigation menu, select **Provisioning**.

1.
 Under **Workflow**, select the **Enable provisioning** checkbox.

1.
 Configure the rest of the settings on this page however you like. We recommend choosing the option to delete users when they are deleted in OneLogin or their app access is removed. Otherwise, redundant users will still be visible in Burp Suite DAST.

1.
 Save your changes.


## Push your OneLogin users to Burp Suite DAST

 Once you have successfully configured the OneLogin integration, you can push your users so that they are available in Burp Suite DAST. To do this, just assign your users and roles to the application that you created in OneLogin.

 After a while, these users will be available in Burp Suite DAST. Any changes you make to these users in OneLogin will automatically be synced. However, note that users will not have access to any functionality unless they are assigned to a group with the relevant roles in Burp Suite DAST.

#### Note

 You can push users from OneLogin, but not groups. When using OneLogin as your IdP, you need to create and manage all of your group assignments directly in Burp Suite DAST.


### Troubleshooting provisioning issues in OneLogin

 To check that all of your users were provisioned successfully:

1.
 In OneLogin, select the application that you created for Burp Suite DAST.

1.
 From the left-hand navigation menu, select **Users**.

1.
 In the **Provisioning State** column, check for any users with the status `Failed`.


 If you find any users that were not provisioned successfully, click the name of the user and in the dialog that opens, click **Reset Login**. This will clear the user's current provisioning state and re-attempt to provision the user.
