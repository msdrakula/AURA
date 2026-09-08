> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions

DAST

# Creating local groups for SAML or LDAP

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you're not using [SCIM](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim), you can match local Burp Suite DAST groups to your identity provider groups.


 Burp Suite DAST matches these groups using an **Identity provider group** value that you set on each local group. This means you don't need to rename groups in your identity provider to match the ones in Burp Suite DAST.


#### Note

 You can add local users to the local groups in Burp Suite DAST. However, you won't be able to see any users that are managed by SAML or LDAP.


 To create local groups for SAML or LDAP in Burp Suite DAST:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Groups**.

1.
 Click **New group**.

1.
 Enter a **Group name**. This name is for your own reference, and does not affect group matching.

1.
 In the **Identity provider group** field, enter the value that your identity provider sends for the group.

1.
 Assign roles to your groups as required. If you do not assign any roles, users can log in but they can't access any functionality.

1.
 Apply [site restrictions](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/restricting-access-to-sites) for each group as necessary. This limits the sites that users in each group can access.


 For SAML, the value that you enter is the one your identity provider sends in the group membership attribute. For LDAP, it is the group's common name (CN), not its full distinguished name. If you manage your users in Entra ID, it may be the group's `Group ID` rather than its name. For more information, see [Enabling Burp Suite DAST to access your Entra ID groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-entra-groups).


 To edit a group's **Identity provider group** value, you need the **Edit groups** permission. Without it, you can view the value but not change it.


 Users can now log in to Burp Suite DAST using their existing credentials. For SAML SSO, users need to click the link on the login page to authenticate themselves via your identity provider.


## How group matching works

 Your identity provider sends a set of groups when a user signs in. Burp Suite DAST adds the user to every local group whose **Identity provider group** value matches one of them. Burp Suite DAST matches the value exactly. It does not trim spaces or change the value in any other way.


 You can set the same value on more than one local group. In this case, users in that identity provider group join all of the local groups that share the value.


 Burp Suite DAST only matches groups that have an **Identity provider group** value. Groups that you create without one are not matched until you set a value.


 Renaming a local group does not change the identity provider group that it matches. To match a different identity provider group, edit the **Identity provider group** value.


#### Note

 A predictable **Identity provider group** value is easier to target. Anyone who controls the group values your identity provider sends could gain access to that group. To protect a group with sensitive permissions, such as **Administrators**, set its value to something that isn't easy to guess.


## Setting the SAML group membership attribute

 By default, Burp Suite DAST reads group membership from a SAML attribute named `http://schemas.xmlsoap.org/claims/Group`. Your identity provider may send group membership under a different attribute. In this case, enter that attribute name in Burp Suite DAST rather than changing your identity provider:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu , select **Integrations**.

1.
 On the **SAML** tile, click **Configure**.

1.
 In **SAML configuration**, enter the attribute name in the **Group membership attribute name** field.

1.
 Click **Save**.


 You must enter a value in the **Group membership attribute name** field. To change it, you need the **Modify settings** permission.


#### Related pages

- [Managing users locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users)
- [Managing roles locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles)
- [Managing groups locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups)
- [Restricting access to sites](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/restricting-access-to-sites)
- [Resetting your admin password](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/resetting-your-admin-password)
