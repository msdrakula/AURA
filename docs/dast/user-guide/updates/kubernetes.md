> Source: https://portswigger.net/burp/documentation/dast/user-guide/updates/kubernetes

DAST

# Updating Burp Suite DAST on Kubernetes

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Kubernetes instances of Burp Suite DAST do not update automatically. However, because all data is held within the database rather than the app itself, you can easily update the installation without affecting any of your configuration settings.


## Preparing to update

 Before you update Burp Suite DAST, you should disable scanning. Any queued scans (that is, scans that have been created by the system but not yet started) will fail if they attempt to run during the update process.


 To disable scanning, open the **Scan resources** settings page and set the **Enable scanning** toggle to off. This prevents Burp Suite DAST from creating any new scans.


#### Note

 Disabling scanning does not affect any scans that are currently in progress. You can safely let these scans run while updating, as any scans that have already started are not affected by the update process.


## Running the update command

 To update a Kubernetes installation of Burp Suite DAST:


1. Download the latest version of the Helm chart from the [Releases](https://portswigger.net/burp/releases/enterprise/latest?platform=kubernetes) page or the [Helm chart GitHub repository](https://github.com/PortSwigger/enterprise-helm-charts).
1. Unpack the new chart into a directory of your choice. Note that you will need the name of this directory when running commands against the chart.
1. Run the update command: `helm upgrade -n <namespace> <deployment name> <name of directory containing the chart>`

 For example, to use a chart located in `enterprise-helm-folder` to update a deployment of Burp Suite DAST called `bsee-deployment` with a namespace of `bsee-namespace`, you would run:
 `helm upgrade -n bsee-namespace bsee-deployment enterprise-helm-folder`

 After you have run the upgrade command, the application updates immediately. Note that there may be a few seconds of downtime during the process.


 Once the system has resumed, you can re-enable scanning by opening the **Scan resources** settings page and setting the **Enable scanning** toggle to on.


#### Note

 Although Burp Suite DAST does not update automatically, the Scanner component can update automatically if configured to do so. For more information on configuring Scanner updates, see the [Managing Updates](https://portswigger.net/burp/documentation/dast/user-guide/updates.html) page.
