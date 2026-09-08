> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/example-integrations/integrate-github

DAST

# Integrating a CI-driven scan with GitHub Actions

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can integrate a CI-driven scan with GitHub Actions. This enables you to use Burp Scanner to run scans as a stage in your existing CI/CD pipeline, and fail builds that meet your issue threshold.


 To learn how to do this, see the readme file for our GitHub Action (opens in a new tab):
  [GitHub Action for CI-driven scans](https://github.com/PortSwigger/ci-driven-scan-github-action)

 You can configure your scan using a configuration file. This enables you to use application logins, and custom scan configurations. To learn more, see [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).

#### Related pages

-  [Adding a configuration file to a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/add-config)
-  [System requirements for CI-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/getting-started#system-requirements)
