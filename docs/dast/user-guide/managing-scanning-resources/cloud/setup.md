> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/setup

DAST

# Setting up a self-hosted scanning machine for a Cloud instance

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

 We provide an installer for Windows and Linux operating systems. You can download these from Burp Suite DAST.


## Prerequisites

-
 The infrastructure meets the [system requirements for self-hosted scanning machines](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/system-requirements).

-
 You have configured your [network and firewall settings](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/network-settings).


## Downloading the installer

1.
 From the settings menu , select **Scanning resources**.

1.
 Click **Manage scanning machines**.

1.
 On the **Self-hosted scan settings** page, click **Add scanning machine**.

1.
 Click **Generate token**, and save the authentication token. You cannot retrieve the authentication token later, so keep it somewhere safe.

1.
 Choose your operating system and copy the URL.

1.
 Use the URL to download the installer.


## Running the installer

1.
 Unzip and run the installer. For Linux, run the installer from the terminal.

1.
 The wizard opens. Follow the wizard, and enter the authentication token when prompted.

1.

 Enter the hostname of your instance when prompted. The format should look something like this:
  `xxxxxx.portswigger.cloud`
1.
 Click **Next**. The scanning machine will be installed.


 In Burp Suite DAST, the new scanning machine is displayed under **Self-hosted scanning machines**. The
 **Health status** shows as **Starting**, and then **Connected**.


#### Note

 For Linux, you need to do some additional steps to enable browser-powered scanning. This gives you access to the full capabilities of Burp Scanner. For more information, see [Browser-powered scanning for Burp Suite DAST](https://portswigger.net/burp/documentation/dast/user-guide/reference/browser-powered#enabling-browser-powered-scanning-on-linux-machines).


## Scanning your sites with self-hosted scanning machines

 The new scanning machine is automatically added to a default self-hosted scanning pool. A scanning pool determines which sites are scanned by which machines.


 In order to use your self-hosted scanning machine, you need to assign your sites to use the same scanning pool.


 If you don't assign your site to a scanning pool, the PortSwigger-hosted scanning machines are used by default.


 To learn how to reassign a site to your scanning pool, see [Reassigning a site to a different pool](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/scanning-pools#reassigning-a-site-to-a-different-pool).


#### Related pages

-  [Managing self-hosted scanning machines with a Cloud instance](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/cloud/managing-scanning-machines)
-  [Managing scanning pools](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/scanning-pools)
-  [Assigning scan limits](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/assigning-scan-limits)
