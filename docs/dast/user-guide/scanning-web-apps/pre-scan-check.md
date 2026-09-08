> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/pre-scan-check

DAST

# Performing a pre-scan check

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can perform a pre-scan check to:


-
 Make sure that Burp Suite DAST can reach your target URLs.

-
 Review and edit your [recorded login sequences](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins), so you can make sure that they log in to your sites successfully.


 Burp Suite DAST prompts you to perform a pre-scan check when you [add a new web app site](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps). You can also perform a pre-scan check on existing sites at any time.


 To perform a pre-scan check:


1.
 From the top menu, select **Sites > All sites**.

1.
 Select the site that you want to perform a pre-scan check for.

1.
 In the **Pre-scan check** menu at the top of the main pane, click **Run pre-scan check**.


 The pre-scan check runs the URL connection check, and attempts to run your recorded login sequences where applicable. The pre-scan check captures still images from the recorded login sequences for you to review.


 When the pre-scan check is complete, you can review the connection check for your URLs. Expand the **Pre-scan check** menu and go to the **Connection checks** tab to see the connection status of your target URLs. You'll see some remedial advice if the connection check fails.


#### Note

 If you want to learn more about the connection check for a particular URL, click **View response** to see the request and response messages. This information may help you to remediate an issue.


## Managing recorded logins

 If you have any recorded login sequences for the site, you'll see the **Recorded logins** tab. This enables you to review and edit your recorded logins. For more information about reviewing recorded logins, see [Reviewing a recorded login sequence](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins#reviewing-a-recorded-login).


To edit a recorded login:

1.
 Click the  pencil icon next to the recorded login sequence you want to modify.

1.
 In the **Edit recorded login** dialog, update the JSON-based script that contains the recorded actions if necessary. You can also change the label for the recorded login.

1.
 Click **Next**.

1.
 Modify the URL and confirmation text used to verify the [authentication status](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins#status-checker) during scans.

1.
 Click **Save**.
