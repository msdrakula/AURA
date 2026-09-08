> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles

DAST

# Managing roles locally

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can use roles to define a set of permissions to perform specific types of action. It's often useful to define roles that match the job functions of people within your team. This makes it easier to place users into the right groups, quickly and reliably.


 Burp Suite DAST has built-in roles, such as Administrator and Scan initiator. You can also create your own custom roles.


#### Warning

 For security reasons, be cautious when assigning permissions to user roles. Users may be able to access and exploit internal systems if your infrastructure isn't sufficiently secured.


 In addition, some permissions may escalate privileges in an unintended way. For example, the **Modify settings** or **Manage extensions** settings may enable users to create new users or upload malicious extensions.


## Creating a new role

 There are dependencies between some permissions. For example, to have permission to edit an entity, you also need permission to view the same entity. If this dependency prevents the selection of a permission, the permission is grayed out.


 To create a new role:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Roles**.

1.
 Click **New role**.

1.
 Enter a **Role name**.

1.
 Select the required permissions from the tree.

1.
 When you're finished, click **Save**.


## Modifying roles

 You can't modify the built-in roles. However, you can modify custom roles:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Roles**.

1.
 From the list, click on the role that you want to modify.

1.
 Modify the permissions for the role.

1.
 When you're finished, click **Save**.


## Deleting roles

 You can't delete the built-in roles. To delete a custom role:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Roles**.

1.
 Find the row for the role that you want to delete, and click .

1.
 At the prompt, click **Delete**.


#### Related pages

- [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control)
- [Managing users locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users)
- [Managing groups locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups)
