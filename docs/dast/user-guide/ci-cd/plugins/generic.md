> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/generic

DAST

# Integrating with other CI/CD platforms

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Although we provide plugins for [Jenkins](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/jenkins) and [TeamCity](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/teamcity), you can integrate Burp Suite DAST with most other CI/CD platforms using our generic, platform-agnostic driver. The functionality is similar, but instead of using a UI to control your settings, you have to pass them into a build step as command line parameters.


#### Note

 The CI/CD driver allows you to configure both [site-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/integration-types#site-driven-scan) and [Burp scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/integration-types#burp-scan). This is determined by whether the `--site-id` parameter is present in the command that triggers the scan. If you're unsure which option is right for you, please refer to the following page for more information: [Integration types](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/integration-types)

#### In this section

-  [Configuring a site-driven scan with the generic CI/CD driver](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/generic/site-driven)
-  [Configuring a Burp scan with the generic CI/CD driver](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/generic/burp-scan)
-  [Parameter reference for the generic CI/CD driver](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/generic/parameter-reference)
-  [Optional settings](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/optional-settings)

  -  [Overriding the default scan configurations from your CI/CD system](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/optional-settings#overriding-the-default-scan-configurations-from-your-ci-cd-system)
  -  [Ignoring issues](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/optional-settings#ignoring-issues)
