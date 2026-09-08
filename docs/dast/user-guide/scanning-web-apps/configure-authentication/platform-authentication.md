> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/platform-authentication

DAST

# Configuring platform authentication

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can add authentication credentials for HTTP Basic and NTLM authentication. Configuring platform authentication enables Burp Scanner to automatically authenticate to destination web servers at the platform level.


 You can add platform authentication credentials when you add or edit a site or folder:


1.
 Under **Scan settings**, go to **Authentication > Platform authentication**.

1.
 Click **Add credentials**.

1.

 In the dialog, specify the platform authentication credentials:


  - **Destination host** - Enter the destination web server address that you want the rule to apply to, for example, `ginandjuice.shop`. Note that you cannot specify an HTTP protocol in this field.
  - **Type** - Choose from **Basic**, **NTLM v1**, or **NTLM v2**.
  - **Username** - Enter a username.
  - **Password** - Enter a password.
  - **Domain** - Only required for NTLM authentication. Enter your domain name.
  - **Domain hostname** - Only required for NTLM authentication. Enter the name of your domain server.
  - **SPNEGO encoding** - Only applies to NTLM authentication.
  - **Negotiate auth scheme** - Only applies to NTLM authentication.

1.
 Click **Save**.


 The credentials are added to the list in the **Platform authentication** tab. Burp Scanner now automatically authenticates all traffic to the destination host.


 To add additional credentials, click  **Add credentials**, then follow the steps above. Burp uses the first credentials in the list that match the destination web server. This enables you to configure authentication for individual hosts, or disable platform authentication for a specific host.


 To delete platform authentication credentials, click the trash icon .


#### Related pages

- For information about how settings from folders and sites combine, see [How scan configurations are combined](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/folder-configs#how-scan-configurations-are-combined).
- You can also configure platform authentication credentials in a custom scan configuration. For more information, see [Custom scan configuration settings](https://portswigger.net/burp/documentation/dast/user-guide/reference/custom-scan-config-settings#platform-authentication).
