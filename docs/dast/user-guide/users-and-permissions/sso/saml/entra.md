> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/entra

DAST

# Configuring SAML SSO with Entra ID

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This section explains how to configure SAML SSO using Entra ID (formerly Azure AD) as your identity provider. You may also need to refer to the Entra ID documentation.


## Before you start

 Make sure your web server URL includes protocol and port information. For more information, see [Configuring your web server](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/configure-web-server).


#### Note

 The relying party trust information is dependent on your web server URL.


## Step 1: Configure your Entra ID Enterprise Application

 To add Burp Suite DAST to your Entra ID Enterprise Applications:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu , select **Integrations**.

1.
 On the **SAML** tile, click **Configure**. Notice that you can copy both the **Relying party trust identifier** and the **Relying party service URL**.

1.
 In Entra ID, go to **Basic SAML Configuration**.

1.
 Paste the **Relying party service URL** into the **Reply URL (Assertion Consumer Service URL)** field.

1.
 Paste the **Relying party trust identifier** into the **Identifier (Entity ID)** field.


## Step 2: Import key details from Entra ID

 To configure Burp Suite DAST, you need to import some key details from Entra ID (formerly Azure AD):


1.
 In Entra ID, go to the **SAML Signing Certificate** page.

1.
 Download the **Federation Metadata XML** file.

1.
 In Burp Suite DAST, make sure that you're still on the **SAML** page.

1.
 In **Company details**, enter your company name.

1.
 In **SAML configuration**, click **Import metadata**.

1.
 Click **Choose file** and select the Federation metadata XML file.

1.
 In the **Group membership attribute name** field, enter `http://schemas.xmlsoap.org/claims/Group`. This must match the name and namespace of the group claim that you set in Entra ID. Importing the metadata file does not set this field for you.

1.
 Click **Save**.


## Step 3: Test your configuration

 Once the connection is successfully established, we recommend that you test your configuration by logging in to Burp Suite DAST. If the configuration was successful, you will see a message that you have logged in, but you don't yet have permission to do anything.


## Managing groups

 You can now configure how you manage your groups:


-
 You can push the groups from your identity provider using SCIM. For more information, see [Configuring SCIM](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim).

-
 Alternatively, you can duplicate your Entra ID groups in Burp Suite DAST, and manage them locally. For more information, see [Enabling Burp Suite DAST to access your Entra ID groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-entra-groups)
