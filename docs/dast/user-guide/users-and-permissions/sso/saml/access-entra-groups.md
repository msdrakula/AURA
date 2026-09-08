> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-entra-groups

DAST

# Enabling Burp Suite DAST to access your Entra ID groups

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you're not using SCIM, you can match local Burp Suite DAST groups to your Entra ID groups. This enables you to manage these groups locally.


 To configure your groups in a way that Burp Suite DAST can recognize:


1.
 In the Entra ID portal, open the application that represents Burp Suite DAST.

1.
 Under **Set up Single Sign-on with SAML**, go to **User Attributes and Claims** and add a group claim.

1.

 Select the **Customize the name of the group claim** checkbox and enter the following values:


  -  **Name**: `Group`
  -  **Namespace**: `http://schemas.xmlsoap.org/claims`

 The next step depends on how you manage your users:


-
 If your Entra ID instance uses an on-premise installation of Active Directory, select `sAMAccountName` as the source attribute. Set each Burp Suite DAST group's **Identity provider group** value to the corresponding `sAMAccountName`.

-
 If you manage your users in Entra ID, select `Group ID` as the source attribute. Set each Burp Suite DAST group's **Identity provider group** value to the corresponding `Group ID`.


 Your Entra ID application may already send group membership under a different claim. In this case, enter that claim name in Burp Suite DAST rather than adding a claim to Entra ID. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).


## Adding your groups to Burp Suite DAST

 To grant permissions, set the **Identity provider group** value on each group that you create in Burp Suite DAST. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).
