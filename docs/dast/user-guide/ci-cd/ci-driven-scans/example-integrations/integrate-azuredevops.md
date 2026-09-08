> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/example-integrations/integrate-azuredevops

DAST

# Integrating a CI-driven scan with Azure DevOps

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This page contains instructions to integrate a CI-driven scan with Azure DevOps. This enables you to use Burp Scanner to run scans as a stage in your existing CI/CD pipeline, and fail builds if issue thresholds are met.

 You configure the scan by defining a set of simple parameters in a YAML file. To learn how to do this, see [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).

 These instructions have been tested with Azure DevOps version 2.387.3.

## Before you start

 You need to complete the following steps before you start:

-

 Deploy Burp Suite DAST. See [Setting up Burp Suite DAST](https://portswigger.net/burp/documentation/dast/setup).

-

 Create an API user in the **CI-driven scan initiator group**, and save the **API key**. See [Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user).

-

 Save the configuration for your CI-driven scan as a YAML file. See [Creating a configuration file for a CI-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).


## Azure DevOps agent requirements

 To integrate a CI-driven scan with Azure DevOps, your Azure DevOps agent must have Docker installed.

 You do not need to install any plugins other than the Azure DevOps defaults.

 For information on the machine specification required to run a CI-driven scan, see [System requirements for
 CI-driven scans](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/getting-started#system-requirements).

## Configuring the Azure DevOps pipeline

1.

 Navigate to your project in Azure DevOps.

1.

 In the left navigation menu, click **Pipelines > New pipeline**.

1.

 Choose the type of repository where your code is stored.

1.

 Choose the specific code repository you want to use, and authorize Azure DevOps to access it if needed.

1.

 Click **Existing Azure Pipelines YAML file**, and select the file you want to use. If you don't already have a script in any branch of the repository that you want to use, choose **Starter pipeline** to build the script for your pipeline.


### (Optional) Creating a starter pipeline YAML file

 If you create a new starter pipeline, you can use the following example script.

 The script includes steps to:

-

 Install Docker.

-

 Run the CI-driven scan container.

-

 Publish test results in JUnit format.

`trigger:
- none

pool:
    name: Azure Pipelines
    vmImage: ubuntu-latest

steps:
- task: DockerInstaller@0
    inputs:
        dockerVersion: '17.09.0-ce'

- script: |
        docker run --rm \
        -u $(id -u) -v $(Agent.BuildDirectory):$(Agent.BuildDirectory):rw -w $(Agent.BuildDirectory) \
        -e BURP_CONFIG_FILE_PATH=$(Build.SourcesDirectory)/burp_config.yml \
        -e BURP_REPORT_FILE_PATH=$(Agent.BuildDirectory)/burp_junit_report.xml \
        public.ecr.aws/portswigger/enterprise-scan-container:latest

    displayName: 'Docker Run Burp Scanner'

- task: PublishTestResults@2
    condition: always()
    inputs:
        testResultsFormat: 'JUnit'
        testResultsFiles: '**/burp_*.xml'
        searchFolder: '$(Agent.BuildDirectory)'
        failTaskOnFailedTests: false
        publishRunAttachments: true
`

 To learn more about creating and editing the configuration file, see [Creating a configuration file for a CI-driven
 scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/ci-driven-scans/create-config).

## Running the Azure DevOps pipeline

 In your project dashboard **Pipelines** area, click **Run pipeline**. This takes you to the pipeline run page, where you can monitor real-time status updates and logs for each step of the pipeline.

 Based on the build failure rules specified in the scan container configuration, the scan fails with a non-zero exit code if issues are identified.

## Viewing scan results in Azure DevOps

 To view the results of your scan:

1.

 When your scan has completed, click its most recent pipeline run to open its details page.

1.

 In the **Jobs** section, locate the job that includes the scan tasks and click to expand it. Review these detailed logs to see the initial scan results.

1.

 Go to the **Tests** tab to view any failed tests. Click a test to see its **Result Details**.


### Remediation advice and evidence

 In the **Debug** tab of the **Results Details** report for a failed test, you can find remediation advice and evidence for security issues identified by Burp Scanner. This section includes:

-

 Links to relevant parts of the Web Security Academy, providing further detail on web security vulnerabilities.

-

 Requests sent by Burp Scanner to produce the issue, as well as the response sent by the application.
