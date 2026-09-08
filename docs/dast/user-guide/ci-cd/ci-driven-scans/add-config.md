> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/add-config

DAST

# Adding a configuration file to a CI-driven scan

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This section explains how to add a configuration file to the container for a CI-driven scan. The configuration file enables you to use more advanced features, such as application logins or custom scan configurations.


 To learn how to create a configuration file and download a template, see [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).


 The instructions on this page are suitable for all platforms. For examples for specific platforms, please see:


-  [Jenkins](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/example-integrations/integrate-jenkins)
-  [TeamCity](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/example-integrations/integrate-teamcity)
-  [GitHub Actions](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/example-integrations/integrate-github)

 To add a configuration file to your CI-driven scan:


1.
 Create your configuration file, see [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).

1.
 Save the configuration file as `burp_config.yml` in the root of the working directory.

1.

 Use the following command to run a scan:
  `docker run --rm --pull=always \
-u $(id -u) -v $(pwd):$(pwd) -w $(pwd) \
public.ecr.aws/portswigger/enterprise-scan-container:latest`

#### Note

 The above command mounts your current directory into the scan container, and sets it as the working directory for the container.



 The scan container looks for the configuration file `burp_config.yml` in the root of its working directory.
