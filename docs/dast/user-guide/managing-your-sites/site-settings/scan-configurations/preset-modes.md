> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/preset-modes

DAST

# Using preset scan modes

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Preset scan modes are predefined collections of scan settings. They offer a quick way to adjust how the scan balances speed and coverage.


#### Note

 Preset scan modes are only available for web app sites.


 To select a preset scan mode:


1. From the top menu, select **Sites**.
1. Select the site from the list.
1. Select the **Details** tab and click **Edit**.
1. From the **Scan settings** panel, click the **Scan configuration** tab.
1. Make sure that **Use a preset scan mode** is selected.
1. Click one of the available options.

#### Note

 We recommend keeping a consistent scan configuration for each site you add. Changing the scan configuration can affect issue trends over time and cause Burp Suite DAST to give inaccurate time estimates while scanning.


 If you want to scan a site you have already added with a new configuration, we recommend adding the site again with the new configuration selected.
