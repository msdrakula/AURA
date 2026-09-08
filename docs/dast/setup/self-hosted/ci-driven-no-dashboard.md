> Source: https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard

DAST

# Setup guide: CI-driven scans with no dashboard

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You can integrate Burp Scanner into your CI/CD pipeline really easily, without the need to set up a DAST server. This is ideal if you only want to run CI-driven scans, and you don't need to use the features of Burp Suite DAST's dashboard.


 This option enables you to run Burp Scanner from a Docker container in your CI/CD platform. You can view the results of your scans directly in your CI/CD platform, where they're saved as a JUnit or Burp XML file.


 Configuring your scans is straightforward. You can use a configuration file to define:


-
 Start URLs, and the [scope of your scan](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/setting-site-scope)
-  [The scan configuration](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations)
-  [Site login details](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication#configuring-login-details)

 The configuration file is in YAML format, and includes comments to make it easy to use.


#### Related pages

-  [Running a basic CI-driven scan with no dashboard](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/basic-scan)
-  [Configuring CI-driven scans with no dashboard](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/configuring-scans)
-  [Example integrations for CI-driven scans with no dashboard](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/example-integrations)

**Next step - **Running your first CI-driven scan with no dashboard  [CONTINUE](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/basic-scan)
