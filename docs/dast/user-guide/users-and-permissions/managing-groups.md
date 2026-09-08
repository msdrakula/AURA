> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups

DAST

# Managing groups locally

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Groups allow you to map users to their relevant roles. This enables you to assign the permissions for a chosen set of roles to all the users in the group. Users in the group inherit the permissions that are defined in the assigned roles, subject to any restrictions on sites.


 Each user can belong to multiple groups. They inherit the roles and permissions from all the groups that they belong to.


 You can also use groups to restrict users to certain parts of the site tree.


 This section explains how to manage groups locally in Burp Suite DAST. Alternatively, you can use SSO or SCIM to manage groups. For more information, see [Managing SCIM users and groups](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/scim/managing-scim).


## Creating a new group

 To create a new group:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Groups**.

1.
 Click **New group**.

1.
 Enter a **Group name**.

1.
 If you use SAML or LDAP single sign-on, fill in the **Identity provider group** field. Enter the value that your identity provider sends for the group. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).

1.
 In the **Roles** tab, select the roles that you want to assign to the group.

1.
 In the **Users** tab, select the users that you want to assign to the group.

1.
 Click **Save**.


## Editing a group

 You can edit a group as follows:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Groups**.

1.
 From the list, click the group that you want to edit.

1.
 Edit the **Identity provider group** value if you need to change the identity provider that matches this group.

1.
 Use the **Roles** and **Users** tabs to edit the group settings.

1.
 When you're finished, click **Save**.


## Restricting access to sites

 You can use groups to restrict user access to certain sites. For further information, refer to [Restricting access to sites](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/restricting-access-to-sites).


## Deleting a group

 You can only delete custom groups. You can't delete built-in groups. To delete a group:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Groups**.

1.
 Find the row for the group that you want to delete, and click .

1.
 At the prompt, click **Delete**.


#### Related pages

- [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control)
- [Managing users locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users)
- [Managing roles locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles)
