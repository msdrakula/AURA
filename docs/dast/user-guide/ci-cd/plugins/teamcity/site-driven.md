> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/teamcity/site-driven

DAST

# Configuring a site-driven scan in TeamCity

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 To integrate Burp Suite DAST with TeamCity, we recommend using the [site-driven scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/integration-types#site-driven-scan) option. In this section, we'll provide step-by-step instructions on the full configuration process.


## Prerequisites

-
 You have [created an API user](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/create-api-user) in Burp Suite DAST and have access to the corresponding API key

-
 You have [installed the plugin](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/teamcity#download-and-install-the-plugin) in TeamCity.

-
 You have finished [setting up the site](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites) that you want to scan in Burp Suite DAST. We recommend running a couple of scans from the web UI to make sure that you're happy with the scan configuration and scanner behavior before starting the CI/CD integration.


## Whitelist your TeamCity URL

 Site-driven scans interact with your DAST server via the GraphQL API. In order to support this behavior, you need to whitelist your TeamCity URL so that TeamCity can make the necessary cross-origin requests for retrieving your site tree and creating new scans.


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu, select **Network**.

1.
 On the network settings page, scroll down to the **Allowed Origins for GraphQL API** section.

1.

 In the provided field, enter your TeamCity URL, including the protocol and port. For example:
  `https://your-teamcity-domain.com:8080`
1.
 Save your entries.


#### Read more
[Whitelisting an application for CORS](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/whitelist-application-cors)

## Create the site-driven scan build step in TeamCity

1.
 Log in to TeamCity.

1.
 Open the pipeline in which you want to incorporate a scan. Alternatively, create a new dummy project if you just want to test the integration process.

1.
 Make sure your pipeline deploys the application that you want to scan to the same URL as the corresponding site in Burp Suite DAST.

1.
 Add a new build step and select the runner type **Burp site-driven scan**.

1.

 Enter the URL of your DAST server. This is the URL that you normally use to access Burp Suite DAST. Make sure you include the appropriate protocol and port. By default, this will be something like:
  `https://your-enterprise-server.com:8080`
1.
 Enter the API key that you generated when creating the API user earlier. If you've lost this, you need to generate a new API key or create a new API user from the Burp Suite DAST web UI.

1.
 Once you have entered both of these values, your site tree will automatically be fetched from Burp Suite DAST. From the drop-down menu, select the site that you want to scan.

1.
 Adjust the various [optional settings](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/optional-settings) to fine-tune how the scan and its results will affect your build.

1.
 Save your pipeline.


## Test your integration

 After you finish configuring the build step, it's a good idea to check whether the integration is working correctly and that your scan is able to run successfully.


1.
 Kick off a build-on-demand and look at the build log in TeamCity. You should see the scan initialize and start crawling. Throughout the scan, you can check the status by monitoring the build log.

1.
 In Burp Suite DAST, go to your site and open the **Scans** tab. You should see the TeamCity-initiated scan in the list.
