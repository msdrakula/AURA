> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations

DAST

# Defining the scan configuration for a site

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 When you create a new site, the **Scan settings > Scan configuration** tab enables you to specify one or more configurations to use to scan the site.


 You must select a scan configuration in order to be able to save a web app site. However, you can save an API site without selecting a scan configuration. You can select either a preset scan mode or define a custom configuration:


- Burp Scanner's preset scan modes are predefined collections of scan settings. They offer a quick way to adjust how the scan balances speed and coverage.
- You can use a custom scan configuration to fine-tune Burp Scanner's behavior to meet your needs.

 Once you have created your site, you can [track issues over time](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/track-issues) and monitor trends. Keep in mind that if you later decide to change the scan configuration for your site, it can skew the trend data.


#### Note

 You can also set scan configurations for folders and subfolders. These configurations are applied to all the subfolders and sites in the folder. You can then fine-tune the scan configurations for individual subfolders and sites. For more information, see [Defining the scan configuration for a folder](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/folder-configs).


#### Related pages

- [Using preset scan modes](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/preset-modes).
- [Using custom scan configurations](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/custom-configs).
- [Defining the scan configuration for a folder](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/folder-configs).
- [Burp Scanner built-in configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations/burp-scanner-built-in-configs) - reference information on Burp Scanner's built-in scan configurations.
- [Custom scan configuration settings (Burp Suite DAST)](https://portswigger.net/burp/documentation/dast/user-guide/reference/custom-scan-config-settings).
