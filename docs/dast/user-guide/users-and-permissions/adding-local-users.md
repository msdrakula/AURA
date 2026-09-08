> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/adding-local-users

DAST

# Adding local users

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can create users and edit permissions directly from the Burp Suite DAST dashboard. If you have a small number of users, you may want to add all of them locally.


 Alternatively, you may want to create local users in addition to those managed by your single sign-on (SSO) solution. Even if you use SSO, we recommend that you create a backup local admin user.


#### Note

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted Before you create new users, we recommend [connecting your SMTP server](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server). This enables Burp Suite DAST to automatically send email invites to newly created users.


 To create a new local user in Burp Suite DAST:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the **Team** menu, select **Add a new user**.

1.
 In the **User credentials** section, enter the details for the new user.

1.
 Under **Choose a login type**, select **Password**.

1.
 To allow the user to log in immediately, select **Enabled**.

1.
 Select the groups that you want the user to belong to.

1.
 When you're finished, scroll down and click **Save**.


 The new user appears in the list of users. The user should automatically receive an email invite to complete the registration process and obtain their password.


 Alternatively, you can copy the link when prompted and email it to the user manually.


## Creating an API user

 If you need to create a user to enable integration with other software, refer to [Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user).


#### Related pages

- [Enabling SSO](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso)
- [Managing users locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users)
- [Managing roles locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-roles)
- [Managing groups locally](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-groups)
