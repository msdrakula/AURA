> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/adfs

DAST

# Configuring SAML SSO with ADFS

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This section explains how to configure SAML SSO using Active Directory Federation Services (ADFS) as your identity provider. You may also need to refer to the ADFS documentation.


## Before you start

 Make sure your web server URL includes protocol and port information. For more information, see [Configuring your web server](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/configure-web-server).


#### Note

 The relying party trust information is dependent on your web server URL.


## Step 1: Add Burp Suite DAST to your trusted applications

 To add Burp Suite DAST to your trusted applications:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu , select **Integrations**.

1.
 On the SAML tile, click **Configure**. Notice that you can copy both the **Relying party trust identifier** and the **Relying party service URL**.

1.
 In ADFS, run the **Add Relying Party Trust** wizard.

1.
 Paste the Relying party service URL into the **Relying party SAML 2.0 SSO Service URL** field.

1.
 Paste the Relying party trust identifier into the **Relying party trust identifier** field.


## Step 2: Obtain key details from ADFS

 To configure Burp Suite DAST, you need to obtain the following key details from ADFS:


-
 The **Entity ID**. This is the URL that is sent as the `Issuer` value in SAML responses.

-
 The **SSO URL**. Burp Suite DAST sends users to this URL when they choose to log in using SAML.

-
 The **token-signing certificate**. Burp Suite DAST uses this to verify that the SAML response was genuinely issued by ADFS.


 For more information on how to find these, see the ADFS documentation.


## Step 3: Enter the key details in Burp Suite DAST

 To enter the key details in Burp Suite DAST:


1.
 In Burp Suite DAST, make sure that you're still on the **SAML** page.

1.
 In **Company details**, enter your company name.

1.
 Enter the key details in the relevant fields.

1.
 In the **Group membership attribute name** field, enter `http://schemas.xmlsoap.org/claims/Group`. This must match the claim type that you configured in ADFS.

1.
 Click **Save**.


## Step 4: Test your configuration

 Once the connection is successfully established, we recommend that you test your configuration by logging in to Burp Suite DAST. If the configuration was successful, you will see a message that you have logged in, but you don't yet have permission to do anything.


## Managing groups

 You can now configure how you manage your groups:


-
 You can push the groups from your identity provider using SCIM. For more information, see [Configuring SCIM](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim).

-
 Alternatively, you can duplicate your ADFS groups in Burp Suite DAST, and manage them locally. For more information, see [Enabling Burp Suite DAST to access your ADFS groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-adfs-groups)
