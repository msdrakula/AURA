> Source: https://portswigger.net/burp/documentation/dast/user-guide/reference/user-activity-logs

DAST

# User activity log

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

You can enable user activity logging to record user actions throughout Burp Suite. You can then download this log as a CSV file that contains a list of timestamped actions and successful logins.

If you deselect the user activity log setting, log entries are retained for the period that was specified.

## Enabling user activity logging

To enable user activity logging:

1. From the top menu, select **Settings** .
1. Select **User activity log**.
1. Under **Enable user activity logging**, select **Log user activity and retain data**.
1. Specify a number of calendar months. The maximum number is `999`.
1. Click **Save**.

## Downloading the user activity log

To be able to download the user activity log CSV file, the **Log user activity and retain data** setting must be selected.

To download the user activity log:

1. From the top menu, select **Settings** .
1. Select **User activity log**.
1. Click **Download log**.
