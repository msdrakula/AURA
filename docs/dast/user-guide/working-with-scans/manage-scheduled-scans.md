> Source: https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans/manage-scheduled-scans

DAST

# Managing scheduled scans

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you schedule a recurring scan, Burp Suite DAST displays statistics and charts that enable you to [track your progress over time](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/track-issues). This enables you to identify areas for improvement and to see changes in your security posture.


-
 You can use the **Sites** page to select multiple folders and sites, and create a scheduled scan for them.

-
 You can use the **Scans > Scheduled scans** page to view and edit the schedule for recurring scans.


## Viewing scheduled scans

 To view schedule information for a recurring scan:


1.
 From the top menu, select **Scans > Scheduled scans**.

1.
 In the list of scans, click the scan you want to view.

1.

 Notice that the panel at the top of the screen displays the following information:


  -  **Schedule start**: The date and time that the scan was first scheduled.

  -  **Recurrence**: How often the scheduled scan repeats.

  -  **Next run**: The time and date that the scheduled scan will next start.


#### Related pages

-  [Working with scan results](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results).


## Editing scheduled scans

 To edit a scan's schedule:


1.
 From the top menu, select **Scans > Scheduled scans**.

1.
 Click the scheduled scan that you want to edit.

1.
 Click **Edit schedule** at the top of the window.

1.
 Under **Start scan**, select the time for the new schedule to start.

1.
 Select how often the scan repeats.

1.

 Select how many times the scan repeats:


  -
 Forever.

  -
 A specific number of times.

  -
 Until a particular date.


1.
 Click **Save**.


## Deleting scheduled scans

 To delete a scheduled scan:


1.
 From the top menu, select **Scans > Scheduled scans**.

1.
 Click the trash icon  next to the schedule you want to delete.

1.
 Click **Confirm**.


## Editing scheduled scan configurations

 You can edit the configurations for a scheduled scan from the **Sites** menu. For more information, see [Configuring site settings](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings).


#### Related pages

-  [Scans page](https://portswigger.net/burp/documentation/dast/user-guide/reference/scans-page#schedule-scan-configurations).
