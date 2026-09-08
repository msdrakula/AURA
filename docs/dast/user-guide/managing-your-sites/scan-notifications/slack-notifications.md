> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/scan-notifications/slack-notifications

DAST

# Setting up Slack notifications

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Setting up Slack notifications can help you to keep up to date with your organization's security posture. When you configure Slack notifications for a site, Burp Suite DAST sends a notification to your nominated channels as soon as a scan of that site starts, fails, or finishes.


#### Note

You must integrate Burp Suite DAST with Slack in order to receive Slack notifications. For more information on how to integrate Slack, see [Integrating Burp Suite DAST with Slack](https://portswigger.net/burp/documentation/dast/user-guide/integrating-slack).

## Setting up Slack notifications when creating a new site

 To set up Slack notifications when creating a new site:


1. On the top menu, select **Sites >  Add a new site** to display the **Create a new site** page.
1. In the **Scan settings** section, select the **Notifications** tab.
1. In the **Send scan notification to Slack** section, enter a **Slack channel**.
1. To specify an additional channel, click the  plus button and enter the required channel name.

## Setting up Slack notifications for existing sites

 To set up Slack notifications for an existing site:


1. On the top menu, select **Sites** to display the site tree.
1. Select the site you want to set up notifications for.
1. Select the **Details** tab and click **Edit**.
1. In the **Scan settings** section, select the **Notifications** tab.
1. In the **Send scan notifications to Slack** section, enter a **Slack channel**.
1. To specify an additional channel, click the  plus button and enter the required channel name.
1. Click **Save**.

#### Related pages

- [Integrating Burp Suite DAST with Slack](https://portswigger.net/burp/documentation/dast/user-guide/integrating-slack).
- [Adding new sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites) - explains the process of setting up a new site in detail.
