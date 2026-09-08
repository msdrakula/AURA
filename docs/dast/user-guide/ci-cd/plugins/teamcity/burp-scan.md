> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/teamcity/burp-scan

DAST

# Configuring a Burp scan in TeamCity

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Configuring a [Burp scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/integration-types#site-driven-scan) in TeamCity involves largely the same process as in previous versions of Burp Suite DAST. In this section, we'll provide step-by-step instructions for the full configuration process.


#### Note

 Although we continue to support the legacy "Burp scan" option, for most users, we recommend [configuring a site-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/teamcity/site-driven) instead.


## Prerequisites

-
 You have [created an API user](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/create-api-user) in Burp Suite DAST and have access to the corresponding API key

-
 You have [installed the plugin](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/teamcity#download-and-install-the-plugin) in TeamCity.

-
 You have familiarized yourself with Burp Suite DAST's [site-matching rules](https://portswigger.net/burp/documentation/dast/user-guide/reference/site-and-scan-data-settings#how-does-burp-suite-dast-decide-which-sites-to-match).


## Create the Burp scan build step in TeamCity

 The following steps are the minimum configuration requirements to integrate TeamCity with Burp Suite DAST.


1.
 Log in to TeamCity.

1.
 Open the pipeline in which you want to incorporate a scan. Alternatively, create a new dummy project if you just want to test the integration process.

1.
 If you want to scan an existing site that you have already configured in Burp Suite DAST, make sure your pipeline deploys this application to the same URL. Alternatively, if you do not want the scan to be matched with an existing site, make sure you deploy the application to a unique URL.

1.

 Add a new build step and select the runner type **Command Line**. Add a custom script that will echo the top-level URL of the running application that you want to scan and assign it to the variable `BURP_SCAN_URL` as follows:
  `echo BURP_SCAN_URL = https://application-to-scan.com`

 This step will output the target URL in its build log in the correct format for the plugin to process in the next step. If you have a more dynamic deployment process, for example, to a Docker container, you should repeat this command multiple times to output each of the relevant URLs. All of these will be aggregated and scanned.

1.
 Add another new build step, but this time select the runner type **Burp scan**.

1.
 Enter the URL of your Burp Suite DAST REST API endpoint. This is the URL that you copied after creating the API user earlier. Make sure you include the appropriate protocol and port.


  -

 If you want to be able to download scan reports, exclude the API key from this URL as follows and instead enter the API key in the dedicated field beneath. This is the recommended approach.
  `URL: https://your-enterprise-server:8080
API key: your-api-key`
  -

 If you do not want to download scan reports, you can include the API in the URL. In this case, you should leave the API key field blank. We do not recommend this approach as it is primarily to provide continued support for legacy integrations that were configured before we adjusted the input fields.
  `URL: https://your-enterprise-server:8080/api/your-api-key
API key: [blank]`

1.
 Adjust the various [optional settings](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/optional-settings) to fine-tune how the scan and its results will affect your build. For "Burp scans", you also have the options to:


  -
 Upload a custom scan definition to either customize the scan configuration for a one-time scan or [override the default configuration](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/optional-settings#overriding-the-default-scan-configurations-from-your-ci-cd-system) for the matched site.

  -
 Define rules for [ignoring issues](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/optional-settings#ignoring-issues). This is useful for setting false positives for a "Burp scan".

 Note that this option is not available for "site-driven scans" because they inherit false positive rules from their associated site in Burp Suite DAST.


1.
 Save your pipeline.


## Test your integration

 After you finish configuring the build step, it's a good idea to check whether the integration is working correctly and that your scan is able to run successfully.


1.
 Kick off a build-on-demand and look at the build log in TeamCity. You should see the scan initialize and start crawling. Throughout the scan, you can check the status by monitoring the build log. Issues that are found will also be output to the log.

1.
 In Burp Suite DAST, go to your site and open the **Scans** tab. You should see the TeamCity-initiated scan in the list.
