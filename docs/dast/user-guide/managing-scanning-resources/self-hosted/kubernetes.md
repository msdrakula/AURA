> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/self-hosted/kubernetes

DAST

# Managing Kubernetes scanning resources

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This page explains how to configure scanning resources on Kubernetes instances of Burp Suite DAST.


## Setting concurrent scan limits

 In theory, Kubernetes enables Burp Suite DAST to scale to any size, based on your subscription.


 To control the amount of cloud resource that can potentially be dedicated to scanning at any one time, set a lower concurrent scan limit:


1.
 Sign in to Burp Suite DAST.

1.
 From the settings menu  select **Scanning resources**.

1.

 Select **Enable scanning**.


![scan resources page](https://portswigger.net/burp/documentation/dast/images/scan-enable-k8.png)

1.
 Click **View scans in progress**.

1.
 Select **Set a concurrent scan limit**.

1.
 Enter the maximum number of concurrent scans that you want Burp Suite DAST to run.

1.
 Click **Save**.


 The changes won't affect scans that are currently running.


#### Note

It can take a few minutes for new nodes to scale up after you initiate a scan.

## Disabling scanning

 If you disable scanning then Burp Suite DAST will not start any new scans. Scans that are in progress will continue.


 To disable scanning:


1.
 From the settings menu  select **Scanning resources**.

1.
 Under **Kubernetes scan containers**, click **View scans in progress**.

1.
 Deselect **Enable scanning**.


## Amending your license

 If you have a Classic license for a specific number of concurrent scans, you can change this number at any time:


1.
 In your browser, log in to `portswigger.net`.

1.
 On the **My account** page, scroll down to **Your Licenses and Products**.

1.
 Increase the number of concurrent scans that your license covers.

1.
 Click **Download license**.

1.
 Sign in to Burp Suite DAST.

1.
 From the settings menu , select **Licensing**.

1.
 Click **Upload license key**.

1.
 Navigate to the license you downloaded and click **Open**.


 Notice that Burp Suite DAST increases your scan limit automatically.


#### Note

 If you have a concurrent scan limit set on the **Kubernetes scan containers** page, Burp Suite DAST won't increase the scan limit automatically. Increase the limit manually in order to run the new maximum number of concurrent scans.


 If you decrease the number of scans covered by your license then the system finishes any scans that are already in progress, even if this exceeds your new limit.


## Viewing active scans

 To view your active scans:


1.
 From the settings menu  select **Scanning resources**.

1.
 Under **Kubernetes scan containers**, click **View scans in progress**.

1.
 Find a scan in the list of **Scans in progress**.

1.
 To view details about the scan, click **View**.


#### Related pages:

-  [Deploying Burp Suite DAST to Kubernetes](https://portswigger.net/burp/documentation/dast/setup/self-hosted/kubernetes).

-  [Deploying additional scanning machines](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/additional-scanning-machines) - explains how to manage scanning machines (as opposed to auto-scaling Kubernetes-based resources).
