> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/managing-scanning-machines

DAST

# Managing self-hosted scanning machines with a Cloud instance

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

 You can manage the scanning pools for your self-hosted scanning machines. For more information, see [Managing scanning pools](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/scanning-pools).


## Updating self-hosted scanning machines

 Your self-hosted scanning machines update automatically.


## Deleting a self-hosted scanning machine

 To delete a self-hosted scanning machine:


1.
 Run the uninstaller for your scanning machine.

1.
 Go to **Scanning resources** and select **Manage scanning machines**. Under **Self-hosted
 scanning machines**, the scanning machine will show as disconnected.

1.
 Click the trash icon  next to the scanning machine.


## Managing authentication tokens

 You can only revoke a token if no scans are active on scanning machines that use that token. To revoke an authentication token:


1.
 Go to **Scanning resources** and select **Manage scanning machines**.

1.
 Under **Authentication tokens**, click the trash icon  next to the token.


 Any scanning machines that were using the authentication token will show as **Disconnected**.


#### Related pages

-  [Managing scanning pools](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/scanning-pools)
-  [Assigning scan limits](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/assigning-scan-limits)
