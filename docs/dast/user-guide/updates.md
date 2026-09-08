> Source: https://portswigger.net/burp/documentation/dast/user-guide/updates

DAST

# Managing updates

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This page explains how to manage updates for standard instances of Burp Suite DAST. Applying updates might result in [some downtime](https://portswigger.net/burp/documentation/dast/user-guide/updates#downtime-during-updates).


#### Note

 To perform an update, the DAST server needs network access to `https://portswigger.net`.


 Updates work differently on Kubernetes. For information on updating Kubernetes instances, see [Updating Burp Suite DAST on Kubernetes](https://portswigger.net/burp/documentation/dast/user-guide/updates/kubernetes).


## Configuring automatic updates

 You can manage updates for Burp Suite DAST and Burp Scanner separately:


1.
 From the settings menu , select **Updates**.

1.
 Look for pending updates for either of the two components:


  -
 Burp Suite DAST: This comprises the DAST server, scanning machines, and web server (including the web UI, REST API, and GraphQL API).

  -  [Burp Scanner](https://portswigger.net/burp/vulnerability-scanner): This is used by scanning machines to perform scans.


1.
 To install a pending update immediately, click **Install update** for the relevant component.


## Manually checking for updates

 Burp Suite DAST automatically checks for updates periodically. You can also check for updates manually:


1.
 From the settings menu , select **Updates**.

1.
 To manually check for updates, click **Check for updates**.

1.
 To install an available update, click **Install now**.


## Manually installing updates

 If your DAST server does not have network access, you can manually update both Burp Scanner and Burp Suite DAST. This ensures that you benefit from the latest improvements and scan checks.


1.
 Open the [Burp Suite Releases](https://portswigger.net/burp/releases) web page and click the **DAST** filter to view the latest Burp Suite DAST releases. Scroll down to the release you wish to install - this does not have to be the latest version.

1.

 To download the installer as a ZIP file, select the relevant option
 from the drop-down menu and click **Download**.


  - Select **Server component updater** to download the Burp Suite DAST update.
  - Select **Scanner component updater** to download the Burp Scanner update.

 The other options in the drop-down menu (Linux installer, Windows installer, and Helm chart) are for fresh installations or Kubernetes deployments, not for updates.

1.
 Log in to Burp Suite DAST.

1.
 From the settings menu , select **Updates**.

1.
 Click the **Upload zip file (offline update)** button and select the installer zip file.

1.
 Once the file has been verified and uploaded, notice that the **Updates** page shows the version details.

1.
 To complete the installation, click the **Install now** button and follow the on-screen instructions.


#### Note

 Burp Suite DAST sends out automated email notifications whenever a new update is available. In order to receive these updates, you will need to connect Burp Suite DAST to your email server. For more information, see [Configuring your SMTP server](https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/configure-smtp-server).


## Downtime during updates

 The impact of applying an update in terms of application downtime is as follows:


-
 Updates to Burp Suite DAST will cause some downtime while the update is applied. The web UI, REST API, and GraphQL API will be unavailable during the update and any scans that are configured to start during the update will be delayed until it is completed. Scans that are already running at the time of the update are unaffected.

-
 To reduce the impact of downtime, you can restrict automatic updates of Burp Suite DAST to specific days and times.

-
 Updates to Burp Scanner do not cause any downtime. Scans that are already running at the time of the update will continue using the version of Burp Scanner that they started with. New scans that start after the update will use the updated version of Burp Scanner.

-
 We recommend that you enable automatic updates of Burp Scanner to make sure that the latest scan checks are available.


 Over time, there might be backwards compatibility limitations on either component in relation to the other. If one component has not been updated for some time, you may have to update it before any further updates can be made to the other component. The **Updates** page indicates if this is the case.


 Some updates might require manual intervention. For example, you may have to accept changes to the software license agreement before an update can be applied. The **Updates** page indicates if this is the case.


## Updating a Cloud instance

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

 Cloud instances update automatically, ensuring that you always have the latest version of both Burp Suite DAST and Burp Scanner.


 Note that you may experience a short period of downtime (typically around 30-60 seconds) when your instance updates. This downtime does not affect any scans that are currently in progress. Any self-hosted scanning machines also update automatically once the PortSwigger-hosted infrastructure finishes updating.


 To view the current version number for your instance, click **? > About**.
