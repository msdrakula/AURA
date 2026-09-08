> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/managing-scim

DAST

# Managing SCIM users and groups

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Users and groups that are pushed to Burp Suite DAST via SCIM are labeled as such throughout the web UI. You manage these users and groups in a slightly different way to local users that were created directly in Burp Suite DAST.


 You can also combine SAML with SCIM. This provides greater transparency because it enables you to view key details about your users and groups from Burp Suite DAST.


## Assigning permissions to SCIM users

 Just like local users, SCIM users inherit their permissions from the groups they belong to. They can be members of both SCIM groups, and groups that you have created in Burp Suite DAST.


## Assigning permissions to SCIM groups

 Before you can use SCIM to configure user groups, you need to set up an LDAP or SAML connection for single sign-on (SSO):


-

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted [Configuring LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/ldap).

-  [Configuring SAML](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml).


 You cannot change the membership of a SCIM group from within Burp Suite DAST. However, you can assign roles and site restrictions to them just like you would for a local group.


 You can also set the **Identity provider group** value on a SCIM group. Your identity provider controls a SCIM group's name and membership, but not this value. Burp Suite DAST keeps the value when your identity provider synchronizes the group again. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).


## Removing a SCIM user

 You cannot remove or disable a SCIM user from within Burp Suite DAST. Instead, you need to remove their assignment in your identity provider's admin console.


#### Note for Okta users

 Even if you delete a user in Okta, this user will be disabled but still visible in Burp Suite DAST. This is due to the way Okta sends data about deleted users to connected applications. In order to completely remove these users, you need to remove your SCIM integration from the Burp Suite DAST settings.


#### Related pages

- [Managing users locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users)
- [Managing roles locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles)
- [Managing groups locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups)
