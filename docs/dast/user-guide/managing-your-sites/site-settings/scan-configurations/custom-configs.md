> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/custom-configs

DAST

# Using custom scan configurations

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 In addition to the Burp Suite DAST preset scan modes, you can create and import custom configurations. This section explains how to create and assign custom configurations to a site.


#### Note

 We recommend keeping a consistent scan configuration for each site you add. Changing the scan configuration can affect issue trends over time and cause Burp Suite DAST to give inaccurate time estimates while scanning.


 If you want to scan a site you have already added with a new configuration, we recommend adding the site again with the new configuration selected.


## Assign a custom scan configuration to a site

 To select a custom scan configuration for a pre-existing site:


1.
 From the top menu, select **Sites**.

1.
 Select the site from the list.

1.
 Select the **Details** tab and click **Edit**.

1.
 Under **Scan settings**, select the **Scan configuration** tab.

1.
 To display a list of scan configurations that are currently assigned to the site, select **Use a custom configuration**.

1.

 To add a scan configuration to your site, you have two choices:


  -
 Select a configuration from the drop-down box to add it to the list.

  -
 To create a new custom scan configuration, click **Create custom configuration**.


#### Related pages

-  [Scan configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations/custom-scan-configurations).

-  [Burp Scanner built-in configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations/burp-scanner-built-in-configs) - reference information on Burp Scanner's built-in scan configurations.

-  [Custom scan configuration settings (Burp Suite DAST)](https://portswigger.net/burp/documentation/dast/user-guide/reference/custom-scan-config-settings).


## Create a custom scan configuration

 To create a custom scan configuration:


1.
 From the settings menu , select **Scan configurations**.

1.
 On the **Scan configuration** page, click **New configuration**.

1.

 Add a name for the configuration:


  -
 Click the **New Scan Configuration** title bar.

  -
 Enter a name.

  -
 Click **OK**.


1.
 Expand each scan configuration menu and change the settings as required.

1.
 When you're happy with your changes, click **Save**.


 You can now select your new scan configuration from the configuration library when you create a new site.


#### Related pages

-  [Scan configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations/custom-scan-configurations).

-  [Custom scan configuration settings (Burp Suite DAST)](https://portswigger.net/burp/documentation/dast/user-guide/reference/custom-scan-config-settings)

## Exporting scan configurations

 You can export your scan configurations from Burp Suite DAST or Burp Suite Professional. This enables you to:


-
 Share your scan configurations with other users in your organization.

-
 Share scan configurations between Burp Suite DAST and Burp Suite Professional.

-
 Use your scan configuration in a CI-driven scan.


 To export a scan configuration from Burp Suite DAST:


1.
 From the settings menu , select **Scan configurations**.

1.
 To download your chosen scan configuration, click the download icon  in the right-hand column.


#### Related pages

 For more information on exporting configuration files from the desktop editions for Burp, see the [Configuration library](https://portswigger.net/burp/documentation/desktop/settings/library) page.


## Importing scan configurations

 You can import scan configurations from other installations of Burp Suite DAST, or Burp Suite Professional.


 To import a scan configuration:


1.
 Export the scan configurations from Burp Suite DAST, or Burp Suite Professional.

1.
 From the settings menu , select **Scan configurations**.

1.
 Click **Import** to display the open file dialog.

1.
 Select the configuration file that you want to import.
