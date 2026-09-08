> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions

DAST

# Managing users and permissions

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Burp Suite DAST uses role-based access control to manage permissions for your users. For more information, see [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control).


 You can add users to Burp Suite DAST locally, or using single sign-on (SSO).


#### Note

 Even if you want to manage user authentication with your existing SSO solution, we recommend that you create a backup local admin user.


## Managing users locally

 You can add and manage your users individually through the Burp Suite DAST dashboard. You can do this for all your users, or you can combine it with your SSO solution. For more information, see [Adding local users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/adding-local-users).


## Using SSO

 Burp Suite DAST supports authentication via the LDAP and SAML SSO standards. You can also create and remove users via SCIM. For more information, see [Enabling SSO](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso).


#### Warning

 Be cautious when creating users and assigning permissions. Users may be able to access and exploit internal systems if your infrastructure isn't sufficiently secured. To mitigate this risk, it's crucial to limit the creation of new users and assign permissions carefully.


#### Related pages

- [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control)
- [Adding local users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/adding-local-users)
- [Enabling SSO](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso)
- [Managing users locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users)
- [Managing roles locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles)
- [Managing groups locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups)
- [Restricting access to sites](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/restricting-access-to-sites)
- [Resetting your admin password](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/resetting-your-admin-password)
