> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control

DAST

# Role-based access control

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Burp Suite DAST uses [role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control). Once you've added your users, you can manage their permissions using roles and groups:


-
 A [user](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users) represents a person who has access to Burp Suite DAST via the web interface, or a system that has access via one of the APIs.

-
 A [role](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles) is a set of permissions to perform specific actions, such as scheduling and deleting scans. You assign roles to groups of users.

-
 A [group](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups) is a collection of users with an assigned set of roles.


 You can also restrict groups to [certain sites.](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/restricting-access-to-sites)

 You can configure groups in two different ways:

-
 Through the Burp Suite DAST web interface.

-
 Using SCIM. For more information, see [Managing SCIM users and groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/managing-scim).


## Vertical segregation of permissions

 You can use the [roles](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles) assigned to a group to provide vertical segregation of permissions. This means that different categories of users can perform different types of actions. For example, you can allow some users to initiate scans, and you can limit others so that they can only view scan results.


## Horizontal segregation of permissions

 You can restrict users' access to specific sites. This allows for horizontal segregation of permissions, meaning users can only perform their permitted actions on data related to their sites.


#### Related pages

- [Managing users and permissions](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions)
- [Adding local users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/adding-local-users)
- [Enabling single sign-on](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso)
- [Managing users locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users)
- [Managing roles locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles)
- [Managing groups locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups)
