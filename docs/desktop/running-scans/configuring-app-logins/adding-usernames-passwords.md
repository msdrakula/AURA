> Source: https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins/adding-usernames-passwords

Professional

# Adding usernames and passwords

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 If your target uses a basic username and password-based login system, you can specify login credentials for Burp Scanner to use when scanning the site. This enables Burp Scanner to log in to the target application and access content that only authenticated users can usually see.


#### Note

Adding a username and password works well for simple login forms with only two input fields. However, if your target uses a more complex login mechanism then you should use recorded login sequences instead.

You cannot use both credential types on a single scan.

 You can manage login credentials from the **Application login** tab of the scan launcher. From here, you can:


- Add new sets of credentials to the scan.
- Edit existing sets of credentials.
- Import sets of credentials from the configuration library.

## Adding login credentials

 To specify username and password login credentials when configuring a scan:


1. From the scan launcher's **Application login** tab, make sure that **Use login credentials (username & password)** is selected.
1. Click **New** to display the **New Login Credentials** dialog.
1. Enter a unique **Label** to identify this set of login credentials.
1. Enter the **Username** and **Password**.
1. Click **OK**.

 Burp Suite adds the specified credentials to the list. You can specify more than one set of login credentials for each scan.


## Editing existing login credentials

 To edit an existing credential set, select it and click **Edit**. You can edit the following details:


- **Label**.
- **Username**.
- **Password**.

 To delete an existing credential set, select it and click **Delete**.


## Testing login functions

 There are some additional options relating to authenticated scanning in the **Testing login functions** section of the crawl configuration.


 From here, you can configure:


- Whether Burp Scanner attempts to self-register a new user on the target website before performing the crawl.
- Whether Burp Scanner uses invalid credentials to deliberately trigger login failures.

#### Related pages

- [Crawl settings](https://portswigger.net/burp/documentation/scanner/scan-configurations/crawl-settings) - Gives more information on the crawl settings available.
- [Using custom scan configurations](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-scans) - Explains how to configure scans in Burp Suite Professional.
- [Login credentials](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/login-credentials) - Explains when and how Burp Scanner uses login credentials during the crawl process.
