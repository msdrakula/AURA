> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/scan-notifications/email-notifications

DAST

# Setting up email notifications

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Setting up email notifications can help you to keep up to date with your organization's security posture. When you configure email notifications for a site, Burp Suite DAST sends a scan summary report to your nominated users as soon as a scan of that site finishes.


#### Note

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

Your SMTP server must be connected to Burp Suite DAST in order for you to set up email notifications. For more information on how to set up your SMTP server, see [Configuring your SMTP server
 ](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server).

## Setting up email notifications when creating a new site

 To set up email notifications during the process of creating a new site:


1. Select **Sites >  Add a new site** to display the **Create a new site** page.
1. In the **Scan settings** section, select the **Notifications** tab.
1. In the **Send scan summary reports by email** section, enter an **Email** address.
1. To specify an additional email, click the  plus button and enter the required email address.

## Setting up email notifications for existing sites

 To set up email notifications for an existing site:


1. Select **Sites** to display the site tree.
1. Select the site you want to set up email notifications for.
1. Select the **Details** tab.
1. Click **Edit**.
1. In the **Scan settings** section, select the **Notifications** tab.
1. In the **Send scan summary reports by email** section, enter an **Email** address.
1. To specify an additional email, click the  plus button and enter the required email address.
1. Click **Save**.

#### Related pages

-

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted [Configuring your SMTP server](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server) - explains how to set up an SMTP server
- [Adding new sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites) - explains the process of setting up a new site in detail.
