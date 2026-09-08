> Source: https://portswigger.net/burp/documentation/dast/user-guide/reference/team-page

DAST

# Team page

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can use the team page to [manage access](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control) for [groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups), [roles](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles), and [users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users). There is a tab for each of these permission levels:

## Users

 The **Users** tab shows a list of all the users in the team. You can use the filter buttons to filter the users by the following values:

-
 API users.

-
 SCIM users.

-
 Locked out.

-
 Enabled.

-
 Never logged in.


 To [create a new user](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users#creating-a-new-user), click the **New user** button.

 To [edit an existing user](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users#editing-users), click their name.

## Groups

 The **Groups** tab shows a list of the groups that users can be allocated to. Groups can be used to [restrict access to sites](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups#restricting-access-to-sites), or to assign roles to groups of users. You can use the filter buttons to filter the groups by the following values:

-
 Built-in.

-
 Custom.

-
 SCIM.


 To [create a new group](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups), click the **New user** button.

 To [edit an existing group](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups), click the group name. This opens the **Configure group permissions** page, where you can assign **Roles**, **Users**, and **Site restrictions** to the group.

## Roles

 The **Roles** tab shows a list of the roles that have been created. Roles can be used to [manage groups of permissions for users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles). You can use the filter buttons to filter the roles by the following values:

-
 Built-in.

-
 Custom.


 To create a new role, click the **New role** button.

 To edit the permissions assigned to a role, click the Role name.

#### Related pages

[Configuring your SMTP server](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server) - explains how to connect your SMTP server to Burp Suite DAST so that your users can receive email updates.
