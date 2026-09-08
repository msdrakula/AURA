> Source: https://portswigger.net/burp/documentation/dast/user-guide/post-installation-config/config-db-backup

DAST

# Configuring database backups

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you use the bundled H2 database, you can control your database backup settings from within Burp Suite DAST as follows:


1.
 Log in to Burp Suite DAST.

1.
 From the settings menu , select **Database backup**.

1.
 To change the location for your saved backup files, edit the **Location of backups** field.

1.
 To change the number of backup files to retain, edit the **Number of backups to store** field.

1.
 If necessary, set how often you want the backups to repeat.

1.
 To backup your database manually, click **Backup now**.


#### Note

 If you use an external database, your database administrator manages your backup settings outside of Burp Suite DAST.
