> Source: https://portswigger.net/burp/documentation/desktop/running-scans/webapp-scans/full-crawl-and-audit

Professional

# Running a full crawl and audit

-

**Last updated: ** September 3, 2026
-

**Read time: ** 5 Minutes

 Burp Scanner can crawl and audit web applications from one or more start URLs. When scanning, it uses a built-in Chromium browser to interact with the application just like a user would, automating the process of mapping and testing the accessible attack surface.


## Step 1: Configure scan type

 To run a full crawl and audit of a web application:


1. On the **Dashboard**, click **New scan** to open the scan launcher.
1. On the **Scan type** tab, select **Crawl and audit**.

 Once you have selected the type of scan you want to run, select the **Scan details** tab.


#### Related pages

 For information on how Burp Scanner crawls and audits web applications, see the [Burp Scanner](https://portswigger.net/burp/documentation/scanner) documentation.


## Step 2: Configure AI scan enhancements (optional)

 Burp can use AI to filter out false positives when it checks for access control issues. These issues are complex and hard to test for automatically. They often require a high level of manual validation.


 Burp's Broken Access Control (BAC) false positive reduction feature analyzes each broken access control issue identified by Burp Scanner. It checks to make sure it relates to a location that an unauthenticated user should not be able to see. This reduces manual verification, saving time and improving scan accuracy.


 To enable this feature for your scan, select the **BAC false positive reduction** checkbox in the ** AI enhancements** panel.


#### Note

 Burp can only analyze broken access control issues for false positives in scans that perform both authenticated and unauthenticated crawls. To configure your scan to do this:


1.

Select a scan configuration that allows for unauthenticated crawling. This can either be a custom configuration or the **Deep scan** preset. For more information, see [Step 4: Select a scan configuration](https://portswigger.net/burp/documentation/desktop/running-scans/webapp-scans/full-crawl-and-audit#step-4-select-a-scan-configuration).
1.

Specify login credentials. For more information, see [Step 5: Configure application logins](https://portswigger.net/burp/documentation/desktop/running-scans/webapp-scans/full-crawl-and-audit#step-5-configure-application-logins-optional).

### Configuring AI credit behavior

 AI-powered features require AI credits. You can configure Burp's behavior if you run out of credits mid-scan using the **Insufficient AI credit behavior** option:


-

**Pause scan and alert me** - Burp pauses the scan and prompts you to buy more credits.
-

**Continue scanning without AI enhancements** - Burp completes the scan without running any further false positive checks.

#### Note

 For more information on how AI credits work in Burp Suite, see [AI credits](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-credits).


## Step 3: Configure scan details

 In the **Scan details** tab, configure the following basic details of the scan:


1.

Enter a URL into the **URLs to scan** field. This is the URL that the scan starts from. To enter multiple URLs, place each on a new line.
1.

Select **Protocol settings**.

  - **Scan using HTTP & HTTPS**.
  - **Scan using my specified protocols**. If you select this option, make sure you include the scheme (`http:` or
 `https:`) in the **URLs to scan** field.

1.

Optionally, use the settings in the **Detailed scope configuration** section to refine the scan scope. This limits the URLs that Burp Scanner can access during the scan.
1.

If you want to isolate the scan, select **Run isolated scan**. Results from isolated scans do not appear in the **Target > Site map**, **Target > Crawl paths**, or
 **Dashboard > All issues** tabs. It can be useful to isolate a scan if you want to test scan configurations without impacting "live" scan results, for example.

 Once you have specified scan details, select the **Scan configuration** tab.


#### Related pages

[Setting scan scope in Burp Suite Professional](https://portswigger.net/burp/documentation/desktop/running-scans/setting-pro-scope) - Gives detailed information on how scan scope works in Burp Suite Professional.


## Step 4: Select a scan configuration

 Scan configurations are collections of settings that define how a scan runs. From the **Scan configuration** tab, you can do the following:


-

**Set up a new configuration** - You can either use a ready-made preset and start scanning immediately, or customize your own.
-

**Load a configuration** - Choose an existing configuration from your configuration library.

 When you've configured your settings, either click **Scan** to start the scan or select another tab to configure application logins or the resource pool.


#### Related pages

- [Configuring scans in Burp Suite Professional](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-scans) - Detailed information on how to configure your scan.
- [Scan configurations (Burp Scanner)](https://portswigger.net/burp/documentation/scanner/scan-configurations) - Reference material on preset and custom scan configurations, as well as the settings available in custom scan configurations.
- [Importing custom configurations](https://portswigger.net/burp/documentation/desktop/settings/library#importing-custom-configurations) - How to import scan configurations from other installations of Burp into your configuration library.

## Step 5: Configure application logins (optional)

 The **Application login** tab enables you to provide credentials for Burp Scanner to submit when it finds login forms. This enables it to discover and audit content that is only accessible to authenticated users.


 There are two types of login credential you can add in Burp Suite Professional:


- Username and password pairs are intended for sites that use a basic, single-step login mechanism. To manage username and password pairs, select **Use login credentials**. From here, you can add new credentials or edit your existing ones.
- Recorded login sequences are intended for sites that use more complex login mechanisms, such as Single Sign-On. You can record login sequences using the Login Recorder for Burp Suite Chrome extension, which is pre-installed in Burp's browser. To manage your recorded login sequences, select **Use recorded login sequences**. From here, you can add new sequences or edit your existing ones.

 You can only use one of the login mechanisms per site.


#### Related pages

- [Managing application logins](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins) - Gives further information on using application logins in Burp Suite Professional.
- [Authenticated scanning (Burp Scanner)](https://portswigger.net/burp/documentation/scanner/authenticated-scanning) - Gives information on how to record login sequences.

## Step 6: Select a resource pool (optional)

 A resource pool is a group of tasks that share a quota of network resources. You can configure each resource pool with its own throttling settings. These control the number of requests that can be made concurrently, or the rate at which requests can be made, or both.


 The **Resource pools** tab enables you to define the pool in which your scan will run. You can select an existing resource pool from the list, or create a new resource pool.


#### Related pages

[Resource pools](https://portswigger.net/burp/documentation/desktop/running-scans/managing-resource-pools) - Gives information on the use cases for resource pools and how to configure them.
