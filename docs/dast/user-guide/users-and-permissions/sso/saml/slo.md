> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/slo

DAST

# Configuring single logout

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Burp Suite DAST provides optional support for single logout (SLO). You can configure SLO after you [configure SAML SSO](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml).


 If you enable SLO, users are automatically logged out of the identity provider when they log out of Burp Suite DAST. This prevents users from inadvertently remaining logged in to multiple applications.


 To configure single logout:


1.
 Generate a new [self-signed x509 certificate](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/managing-certificates).

1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu , select **Integrations**.

1.
 On the **SAML** tile, click **Edit**.

1.
 In **Relying trust information**, copy the **Relying party single logout URL**. Leave this page open.

1.
 Go to your identity provider's admin panel and edit the SAML settings for your Burp Suite DAST integration.

1.
 Paste the URL from your clipboard into the appropriate field.

1.
 Obtain the **Single Logout URL** from your identity provider. This may have a different name depending on your identity provider.

1.
 In Burp Suite DAST, select **Use single logout**.

1.
 In the **Identity provider single logout URL** field, enter the URL you obtained from your identity provider.

1.
 Paste your self-signed certificate in **Service provider certificate**.

1.
 Paste the private key for your certificate in **Service provider private key**.


#### Note

 Some identity providers, such as Okta, require that Burp Suite DAST signs all the single logout messages that it generates. This is necessary to verify that they come from a trusted source. In this case, you may also need to upload the certificate that you generated to your identity provider.
