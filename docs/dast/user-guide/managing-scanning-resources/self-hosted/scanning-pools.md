> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/self-hosted/scanning-pools

DAST

# Scanning pools for self-hosted instances of Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can use scanning pools to manage your scanning machines. Scanning pools stop the problem of a scan failing because the relevant machine is busy elsewhere, or an assigned machine being unable to access a restricted site for a scan.


 Some example uses for scanning pools are:


-

Keeping the scanning machines and sites for one geographic area together.
-

Cordoning off the resources of one team.
-

Scanning sites with restricted access.
-

Reserving scanning machines for specific purposes, such as a CI/CD pipeline or ad-hoc scanning.

#### Note

 Scanning pools are only available for scanning machines used in standard (non-Kubernetes) instances of Burp Suite DAST. You cannot allocate CI-driven scans to scanning pools.


## Features of scanning pools

 Each scanning machine and site is assigned to a pool. The important features of scanning pools are:


-

Each scanning machine must belong to a scanning pool.
-

Each site to be scanned must belong to a scanning pool.
-

Only scanning machines in the same pool as a site can scan that site.

## Creating a new scanning pool

 To create a new scanning pool:


1.

From the settings menu , select **Scanning resources**.
1.

Under **Scanning machines**, click **Manage scanning machines**.
1.

Click on the **Scanning pools** tab.
1.

Click **Create pool**.
1.

Enter a name and description for the new scanning pool.
1.

Assign the relevant scanning machines and sites to the new pool.
1.

Click **Save**.

## Reassigning a scanning machine to a different pool

 To reassign an existing scanning machine to a different pool, do the following:


1.

From the settings menu , select **Scanning resources**.
1.

Under **Scanning machines**, click **Manage scanning machines**.
1.

Make sure you are on the **Scanning machines** tab.
1.

In the list of scanning machines, select the **Scanning pool** drop-down menu for the scanning machine you want to reassign, and click on the name of the scanning pool to place it in that pool.

## Reassigning a site to a different pool

 To reassign an existing site to a different scanning pool, do the following:


1.

Click on the **Sites** tab.
1.

Select the relevant site.
1.

Click on the **Details** tab.
1.

Click **Edit**.
1.

Under **Scan settings**, go to the **Scanning pool** tab.
1.

Click on the **Scanning pool** drop-down and select the scanning machine pool you want the site to move to.
1.

Click **Save**.

#### Related pages

[Assigning scan limits](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/assigning-scan-limits)
