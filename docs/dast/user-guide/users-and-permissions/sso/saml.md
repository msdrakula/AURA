> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml

DAST

# Configuring SAML single sign-on for Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Burp Suite DAST supports SAML-based single sign-on (SSO). This allows users to log in with their existing credentials.


#### Note

 You can also [integrate SCIM](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim) in combination with SAML. This enables you to create, update, and delete users and groups via SCIM and use SAML exclusively for authentication.


 Combining SCIM and SAML enables you to view key details about your users and groups from Burp Suite DAST.


 Before you configure Burp Suite DAST to use SAML, you need to enable HTTPS on your web server. Refer to [Configuring your web server](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/configure-web-server) and follow the instructions to enable TLS.


 Make sure your web server URL includes protocol and port information. The relying party trust information is dependent on your web server URL.


 We've fully tested SAML integration with the following identity providers:


- Active Directory Federation Services (ADFS)
- Okta
- Entra ID (formerly Azure AD)

#### Related pages

-  [Configuring SAML SSO with ADFS](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/adfs).

-  [Enabling Burp Suite DAST to access your ADFS groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-adfs-groups).

-  [Configuring SAML SSO with Okta](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/okta).

-  [Enabling Burp Suite DAST to access your Okta groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-okta-groups).

-  [Configuring SAML SSO with Entra ID](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/entra).

-  [Enabling Burp Suite DAST to access your Entra ID groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-entra-groups).

-  [Configuring single logout](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/slo)
