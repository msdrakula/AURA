> Source: https://portswigger.net/burp/documentation/dast/user-guide

DAST

# Burp Suite DAST user guide

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Welcome to the Burp Suite DAST user guide. This guide explains how to set up users, sites, and scans so that you can get your scanning workflow up and running. It also gives a comprehensive overview of Burp Suite DAST's other features, including CI/CD integration, issue tracking, and more.

#### Note

 This guide assumes that you have already set your Burp Suite DAST up and have access to the UI. If your organization has not yet set Burp Suite DAST up, see [Setting up Burp Suite DAST](https://portswigger.net/burp/documentation/dast/setup).


## Getting started with Burp Suite DAST

 Before you can get scanning, there are a few things you'll need to configure. Follow the steps below to get up and running:

### Step 1: Set up your users

 Managing permissions in this way makes it easy for you to give users the access they need. For example, you could set up separate roles for your security, IT infrastructure, and management teams, each with their own combination of permissions.

#### Related pages

- [Managing users](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/managing-users).
- [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control).

### Step 2: Add the sites you want to scan

 To scan a web app or an API, you need to first set it up as a Site in Burp Suite DAST.

 You can configure a wide range of settings to determine how each of your sites should be scanned, including:

-

Which of the site's URLs should be scanned and which (if any) should be excluded from scans.
-

The login mechanisms Burp Scanner should use to access your site.
-

Whether Burp Suite DAST should send any automated notifications when scanning the site.

#### Related pages

[Adding new sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites)

### Step 3: Set up a scan configuration

 In Burp Suite DAST, a scan configuration is a set of predefined settings that determine how scans should be performed on a particular site. For example, a scan configuration can specify the maximum crawl depth of the crawl, or what types of issues to report.

 You can either select a predefined scan configuration or create your own for each of your sites.

#### Related pages

[Defining the scan configuration](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations)

### Step 4: Schedule your scans

 Scheduling regular scans is the best way to see changes in your security posture and identify areas for improvement. Scans that run at set intervals with the same configuration are easier to compare than one-off scans. They help you to see how changes to your sites affect the issues you find.

 You can set up unlimited sites and run unlimited scans in Burp Suite DAST at no extra cost.

#### Related pages

[Managing scheduled scans](https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans/manage-scheduled-scans)

### Step 5: View scan results

 Burp Suite DAST makes it easy for you to track your scanning progress over time. You can also view details of individual issues, and raise tickets in third-party issue tracking systems if you have set up the relevant integrations.

#### Related pages

[Working with scan results](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results)

## What else can I do with Burp Suite DAST?

 Burp Suite DAST offers a wide range of additional features, enabling you to:

-

[Use Burp AI to further investigate issues.](https://portswigger.net/burp/documentation/dast/user-guide/using-ai/burp-ai)
-

[Integrate scans easily with any CI/CD platform.](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd)
-

[Generate scan reports by email.](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server)
-

[Export scan results to other reporting tools.](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/generate-reports#export-issue-data)
-

[Generate reports for compliance standards.](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/generate-reports#download-a-compliance-report)
-

[Use extensions to customize scans.](https://portswigger.net/burp/documentation/dast/user-guide/extensions)
-

[Integrate with issue tracking platforms, such as Jira, GitLab, and Trello.](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms)
-

[Integrate with Slack.](https://portswigger.net/burp/documentation/dast/user-guide/integrating-slack)
-

[Integrate with Splunk.](https://portswigger.net/burp/documentation/dast/user-guide/integrating-splunk)
-

[Integrate with third-party applications via the GraphQL API.](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation)
