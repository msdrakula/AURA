> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/usernames-passwords

DAST

# Adding usernames and passwords for a web app

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If your web app uses a basic username and password-based login system, you can specify login credentials for Burp Scanner to use when scanning the site. Specifying a valid username and password enables Burp Scanner to log in to the web app and audit content that only authenticated users can usually see.


#### Note

 We recommend using a recorded login sequence, even for sites that use basic username and password authentication. Recorded logins enable the pre-scan check to confirm that authentication is working, and provide troubleshooting.


 Recorded logins also support our status checker, which can make scans faster and deeper.


### Specifying username and password details when adding a new web app site

 To specify username and password login credentials during the process of adding a new web app site:


1.
 On the top menu, select **Sites > Add a new site** to display the **Create a new site** page.

1.
 In the **Scan settings** section, select **Authentication > Application logins**.

1.
 Make sure that **Usernames and passwords** is selected, and click **Add login credentials**.

1.
 In the dialog box, enter a unique **Label** to identify this set of login credentials.

1.
 Enter the **Username** and **Password**.

1.
 Click **Save**.


### Specifying username and password details for an existing web app

 To specify username and password login credentials for an existing web app site:


1.
 On the top menu, select **Sites** to display the site tree.

1.
 Select the site you want to set up notifications for.

1.
 Select the **Details** tab and click **Edit**.

1.
 In the **Scan settings** section, select **Authentication > Application logins**.

1.
 Make sure that **Usernames and passwords** is selected, and click **Add login credentials**.

1.
 In the dialog box, enter a unique **Label** to identify this set of login credentials.

1.
 Enter the **Username** and **Password**.

1.
 Click **Save** to close the dialog box.

1.
 Click **Save**.


 To specify an additional set of credentials, click the plus button  and repeat steps 6 to 9.


 To delete a set of credentials, click the trash icon .


#### Related pages

- [Adding new web app sites](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps).
- [Adding recorded login sequences](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins).
