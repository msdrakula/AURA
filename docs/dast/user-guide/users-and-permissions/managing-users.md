> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users

DAST

# Managing users locally

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 A user is a person who has access to Burp Suite DAST through the web interface, or a system that has access via one of the APIs.


 This section explains how to create users locally in Burp Suite DAST. Alternatively, you can use SSO and SCIM to manage your users. For more information, see [Managing SCIM users and groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/managing-scim).


 Burp Suite DAST uses role-based access control. For more information, see [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control).


## Viewing users

1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **All users**.

1.
 To filter the list of users, click the filter buttons.


## Creating a new user

 You can either create local users directly in Burp Suite DAST, or configure a SCIM integration to push users from your existing identity provider (IdP). For more information, see:


-  [Adding local users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/adding-local-users)
-  [Enabling single sign-on (SSO)](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso)

## Editing users

 You can edit most of the user details. However, you can't change a user's login type.


 To edit an existing user:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **All users**.

1.
 Click the user in the list and edit their details.

1.
 When you're finished, scroll down and click **Save**.


#### Note

 If you have enabled a SCIM integration, you need to manage any SCIM users and groups using your identity provider's administration console.


## Suspending a user temporarily

 You can temporarily suspend a user. You may want to do this if they are on extended leave, for example:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **All users**.

1.
 Click the user in the list.

1.
 To temporarily suspend the user, deselect **Enabled**.

1.
 Scroll down and click **Save**.


## Deleting a user

 To delete a user:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **All users**.

1.
 Find the user in the list.

1.
 In the right-hand column, click .

1.
 At the prompt, click **Delete**.


#### Related pages

- [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control)
- [Adding local users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/adding-local-users)
- [Managing roles locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles)
- [Managing groups locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups)
