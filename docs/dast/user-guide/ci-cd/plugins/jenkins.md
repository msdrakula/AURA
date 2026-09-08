> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/jenkins

DAST

# Integrating Burp Suite DAST with Jenkins

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Integrating Burp Suite DAST with Jenkins is made simple thanks to our Jenkins plugin. Before beginning, you should decide which [integration type](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/integration-types) you want to configure. In most cases, we recommend the site-driven scan option.


 Integrating with Jenkins involves the following steps.


## Create an API user

 Regardless of which integration type you want to configure, you first need to create an API user.


#### Read more
[Creating an API user for the CI/CD integration](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/create-api-user)

## Download and install the plugin

 We provide a plugin for Jenkins to make the integration process as simple as possible. Please note that this requires Java 11 and Jenkins version 2.164.1 or higher.


1.
 Go to our [website](https://portswigger.net/burp/releases#driver) and download the Jenkins plugin. The download contains a file with the extension `.hpi`.

1.
 Log in to Jenkins as an administrator.

1.
 Go to **Manage Jenkins > Manage Plugins** and open the **Advanced** tab.

1.
 Under **Upload Plugin**, upload the HPI file that you just downloaded. The plugin will begin installing.

1.
 Once the upload is complete, restart Jenkins.


 When creating new build steps, you should now see two new types available for selection: **Burp site-driven scan** and **Burp scan**.


## Configure the integration

 The steps for configuring the integration differ greatly depending on whether you want to use the site-driven scan or **Burp scan** option.


#### Read more

-  [Configuring a site-driven scan (recommended)](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/jenkins/site-driven)
-  [Configuring a Burp scan](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/jenkins/burp-scan)
