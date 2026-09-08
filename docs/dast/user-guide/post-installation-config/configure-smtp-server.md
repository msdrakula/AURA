> Source: https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server

DAST

# Configuring your SMTP server

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you configure a connection between Burp Suite DAST and your SMTP server, you can enable some useful features:


-
 Administrators can send email invites to new users.

-
 Users can automatically receive end-of-scan reports.

-
 Administrators can automatically send password reset links to users.

-
 Administrators can receive system alerts, such as low disk space warnings.


#### Note

 Before you can connect to your SMTP server, you need to configure your [web server settings](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/configure-web-server).


## Using an HTTP proxy server to connect to the SMTP server

 If you use an HTTP proxy connection for communication with `portswigger.net`, you can also use the proxy for connecting to an external SMTP server. This is only supported if you use an unauthenticated proxy.


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu , select **Network**.

1.
 Scroll down to **HTTP proxy server**.

1.
 Make sure that **Use proxy to connect to email server** is selected.


## Connecting to an external email service

 To configure your email settings, you need to know the SMTP server hostname and port for your email service provider. For example, Gmail uses `smtp.gmail.com` and port `587`. Ports 25 and 465 are also commonly used.


 You can find these settings in the documentation of your email service provider.


1.
 From the settings menu , select **Email**.

1.
 Select **Connect to an SMTP server**.

1.
 Enter the **SMTP server hostname**.

1.
 Enter the **Port** that the server uses for sending emails.

1.

 In the **From address** field, enter the email address that you want Burp Suite DAST to use when sending emails.


#### Note

 To reduce the risk of messages being blocked by spam filters, use an email address registered to your email service provider. For example, if you use the Gmail SMTP server, use a Gmail email address.

1.
 Select **Authenticated**.

1.
 Enter the **Username** and **Password** for the email account.

1.
 Select **Use TLS**.

1.
 When you are happy with your settings, click **Save**.


## Connecting to an internal email service

 You can use the same procedure to connect to an internal email service. However, it may not be necessary to select **Authenticated**, or **Use TLS**.


## Verifying your SMTP server connection

 When you save your configuration, the **Send a test email** field appears. Now you can check that the link in the auto-generated email works correctly:


1.
 To send a test email to the email address of the admin user, click **Send**.

1.
 Open the email and click the **Check Email Link** button.

1.
 Make sure that the link takes you to the email settings page in Burp Suite DAST.


 If the link doesn't work, do the following steps:


1.
 From the settings menu , select **Network**.

1.
 Make sure that the **Web server URL** is correct.

1.
 Send another test email and confirm that the link works.
