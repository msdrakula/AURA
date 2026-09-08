> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication

DAST

# Configuring authentication for web apps

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Adding authentication credentials for web app sites enables Burp Scanner to discover and audit content that is only accessible to authenticated users.


 You can add the following types of authentication credentials:


-
 Site login details

-
 Platform authentication details


#### Note

 This page explains how to configure web app authentication. For information on how to configure API authentication, see [Adding new APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis).


## Configuring login details

 Adding login credentials for a web app site enables Burp Scanner to discover and audit content that is only accessible to authenticated users.


 There are two types of login credential that you can add in Burp Suite DAST:


- Username and password pairs are intended for web apps that use a basic, single-step login mechanism.
- Recorded login sequences are intended for web apps that use more complex login mechanisms, such as Single Sign-On, TOTP MFA, or WebAuthn. You can create recorded login sequences manually, or record them using Burp AI.

 You can only use one of the available login mechanisms per site.


#### Related pages

- [Adding usernames and passwords](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/usernames-passwords)
- [Adding recorded login sequences](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins)
- [Managing steps in a recorded login](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/managing-recorded-login-steps)
- [Configuring TOTP MFA](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/totp-mfa)
- [WebAuthn passkeys in recorded logins](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/webauthn)

## Configuring platform authentication details

 Adding credentials for NTLM and HTTP Basic authentication enables Burp Scanner to automatically authenticate to destination web servers at the platform level.


#### Related pages

[Configuring platform authentication](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/platform-authentication)
