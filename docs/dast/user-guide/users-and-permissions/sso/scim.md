> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim

DAST

# Configuring SCIM

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Burp Suite DAST allows you to integrate SCIM in order to simplify the process of provisioning and decommissioning users from a central identity provider (IdP). We have fully tested SCIM integrations with the following IdPs:


-

 Okta

-

 OneLogin

-

 Entra ID


 SCIM is typically integrated in conjunction with SAML. This means you're able to create, update, and delete users and groups via SCIM and leave SAML exclusively for handling authentication. This also provides greater transparency because it enables you to view key details about your users and groups directly from Burp Suite DAST.


 You can also use SCIM to push your users, then assign groups from Burp Suite DAST.


#### Note

 SCIM is not currently supported for self-hosted Kubernetes instances of Burp Suite DAST.


#### Related pages

-  [Integrating SCIM using Okta](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/okta)
-  [Integrating SCIM using OneLogin](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/onelogin)
-  [Integrating SCIM using Entra ID](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/entra-id)
-  [Managing SCIM users and groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/managing-scim)
