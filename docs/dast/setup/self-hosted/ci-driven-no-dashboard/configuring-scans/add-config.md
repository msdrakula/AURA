> Source: https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/configuring-scans/add-config

DAST

# Adding a configuration file to a CI-driven scan with no dashboard

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 This section explains how to add a configuration file to the container for your scans. This enables you to use more advanced features, such as application logins or built-in scan configurations.

 To learn how to create a configuration file and download a template, see [Creating a configuration file for a CI-driven scan with no dashboard](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/configuring-scans/create-config).

 The instructions on this page are suitable for all CI platforms. For reference, we've also provided some more specific examples of how this process looks when integrating with the following platforms:

-  [Example Jenkins integration](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/example-integrations/integrate-jenkins)
-  [Example TeamCity integration](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/example-integrations/integrate-teamcity)
-  [Example GitHub Actions integration](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/example-integrations/integrate-github)

 To add a configuration file:

1.
 Create your configuration file, see [Creating a configuration file for a CI-driven scan with no dashboard](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/configuring-scans/create-config).

1.
 Save the configuration file as `burp_config.yml` in the root of the working directory.

1.

 Use the following command to run a scan:
  `docker run --rm --pull=always \
    -u $(id -u) \
    -v $(pwd):$(pwd) \
    -w $(pwd) \
    public.ecr.aws/portswigger/enterprise-scan-container:latest`

#### Note

 The above command mounts your current directory into the scan container, and sets it as the working directory for the container.


 The scan container looks for the configuration file `burp_config.yml` in the root of its working directory.


**Next step - **Example integrations for CI-driven scans with no dashboard  [CONTINUE](https://portswigger.net/burp/documentation/dast/setup/self-hosted/ci-driven-no-dashboard/example-integrations)
