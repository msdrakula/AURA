> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/example-integrations/integrate-gitlab

DAST

# Integrating a CI-driven scan with GitLab

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This page contains instructions to integrate a CI-driven scan with GitLab. This enables you to use Burp Scanner to run scans as a stage in your existing CI/CD pipeline, and fail builds if issue thresholds are met.

 You configure the scan by defining a set of simple parameters in a YAML file. To learn how to do this, see [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).

 These instructions have been tested with GitLab version 17.1.

## Before you start

 You need to complete the following steps before you start:

-

 Deploy Burp Suite DAST. See [Setting up Burp Suite DAST](https://portswigger.net/burp/documentation/dast/setup).

-

 Create an API user in the **CI-driven scan initiator group**, and save the **API key**. See [Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user).

-

 Save the YAML configuration file for your CI-driven scan. See [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).


## GitLab agent requirements

 To integrate a CI-driven scan with GitLab, your GitLab agent must have Docker installed.

 You do not need to install any plugins other than the GitLab defaults.

 For information on the machine specification required to run a CI-driven scan, see [System requirements for CI-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/getting-started#system-requirements).

## Configuring the GitLab pipeline

1.

 Navigate to your project in GitLab.

1.

 Select **Code > Repository**.

1.

 Above the file list, use the drop-down to select the branch you want to commit to.

1.

 In the drop-down to the right of this, select the file you want to use. If you don't already have a script in the branch, choose **New file**.


### (Optional) Creating a starter pipeline YAML file

 If you don't already have a YAML configuration file, you can use the following example script.

 The script uses pipeline secrets for the Enterprise URL and API key:
`pipeline:
    image: docker:latest
    stage: test
    services:
        - docker:dind
    script:
        - docker run --rm
            -u $(id -u) -v $(pwd):$(pwd):rw -w $(pwd)
            -e BURP_CONFIG_FILE_PATH=$(pwd)/burp-config.yml
            -e BURP_REPORT_FILE_PATH=$(pwd)/burp_junit_report.xml
            -e BURP_ENTERPRISE_API_KEY=$BURP_ENTERPRISE_API_KEY
            -e BURP_ENTERPRISE_SERVER_URL=$BURP_ENTERPRISE_SERVER_URL
            public.ecr.aws/portswigger/enterprise-scan-container:latest
    artifacts:
        when: always
        paths:
            - burp_junit_report.xml
        reports:
            junit: burp_junit_report.xml`

 To learn more about creating and editing the configuration file, see [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).

## Running the GitLab pipeline

 To run the GitLab pipeline:

1.

 In your project dashboard, select **Build > Pipelines**.

1.

 Select **Run pipeline**.

1.

 In the **Run for branch name or tag** field, select the branch or tag to run the pipeline for.

1.

 Enter any CI/CD variables required for the pipeline to run. You can set specific variables to have their values prefilled in the form.

1.

 Select **Run pipeline**.


 The pipeline now runs the jobs according to its configuration.

## Viewing scan results in GitLab

 When your scan has completed, you can view the results of your scan:

1.

 In your project dashboard, select **Build > Pipelines**.

1.

 Select a pipeline to open its details page.

1.

 Select the **Tests** tab.


### Remediation advice and evidence

 Click **View details** in the **Tests** tab of the report to find remediation advice and evidence for security issues identified by Burp Scanner. This section includes:

-

 Links to relevant parts of the Web Security Academy, providing further details about web security vulnerabilities.

-

 Requests sent by Burp Scanner to produce the issue, as well as the response sent by the application.
