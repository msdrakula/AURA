> Source: https://portswigger.net/burp/documentation/dast/setup/what-next

DAST

# Next steps after setup

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

**Congratulations** - your instance of Burp Suite DAST is ready to use.


 You can explore the dashboard and start running scans straight away. For more information, see the [Burp Suite DAST user guide](https://portswigger.net/burp/documentation/dast/user-guide).


 However, there are some key tasks that we recommend completing before you start running a full production scanning workflow. See the sections below for more details on how to get the most out of Burp Suite DAST.


#### On this page

- [Set up your email server](https://portswigger.net/burp/documentation/dast/setup/what-next#set-up-your-email-server).
- [Set up integrations](https://portswigger.net/burp/documentation/dast/setup/what-next#set-up-integrations).
- [Set up SSO](https://portswigger.net/burp/documentation/dast/setup/what-next#set-up-sso).
- [Import sites in bulk](https://portswigger.net/burp/documentation/dast/setup/what-next#import-sites-in-bulk).

## Set up your email server

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can connect Burp Suite DAST to both internal and external email services. Setting up an email server gives you the following benefits:


- Administrators can send email invites to new users.
- Users can automatically receive end-of-scan reports.
- Administrators can automatically send password reset links to users.
- Administrators can receive system alerts, such as low disk space warnings.

#### Related pages

[Configuring your SMTP server](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server).


## Set up integrations

 You can integrate Burp Suite DAST into your existing CI/CD pipelines. For more information, see [Integrate with your CI/CD platform](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd).


 You can also connect to the following issue tracking and workflow tools:


- [Jira](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-jira).
- [GitLab](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-gitlab).
- [Trello](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-trello).
- [Slack](https://portswigger.net/burp/documentation/dast/user-guide/integrating-slack).

## Set up SSO

 Burp Suite DAST supports a number of single sign-on options. This includes LDAP or SAML-based authentication, as well as user provisioning via SCIM.


#### Related pages

- [Managing users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users).
- [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control).
- [Enabling single sign-on (SSO)](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso).

## Import sites in bulk

 Burp Suite DAST's CSV-based bulk site upload feature makes it easier to add large numbers of sites to the system. This is useful if you're looking to migrate from an alternative tool that allows you to export your existing sites.


#### Related pages

[Importing sites in bulk](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/importing-sites-in-bulk).
