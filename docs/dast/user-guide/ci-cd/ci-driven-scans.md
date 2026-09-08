> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans

DAST

# Integrating CI-driven scans

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can integrate CI-driven scans into your CI/CD pipeline. This enables Burp Scanner to run from a Docker container, and report results back to your Burp Suite DAST server. CI-driven scans make it easy to scan sites and applications before they enter production.


## What are CI-driven scans?

 When a CI-driven scan is initiated, an instance of Burp Scanner is created inside a Docker container. This instance of Burp Scanner runs a local scan on a specified URL, defined by an environment variable in your pipeline script. Once the scan has finished, the instance of Burp Scanner sends the results to your Burp Suite DAST server in JUnit XML format.


#### Note

 We provide full setup walkthroughs for Jenkins, TeamCity, and GitHub Actions. However, you can use our generic setup instructions to fully integrate with any CI platform, including CircleCI, Bamboo, and Azure DevOps.


## Configuring your scan

 CI-driven scans are configured using a YAML file. This file defines:


-
 Start URLs

-  [The scope of your scan](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/setting-site-scope)
-  [The scan configuration](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations)
-  [Site login details](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication#configuring-login-details)

 You can also use the YAML file to apply custom extensions, BChecks, and BApps to your scans. For more information, see [Using custom extensions, BChecks, and BApps with CI-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config#using-custom-extensions-bchecks-and-bapps-with-ci-driven-scans).


## Viewing your scan results

 You can view your scan results in a number of ways:


-
 In your CI/CD environment

-
 By viewing the JUnit XML file directly

-
 In the web interface for Burp Suite DAST


#### Related pages

- [System requirements for CI-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/ci-cd-system-requirements)
- [Getting started with CI-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/getting-started)
- [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config)
- [Adding a configuration file to a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/add-config)
- [Integration walkthroughs](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/example-integrations)
