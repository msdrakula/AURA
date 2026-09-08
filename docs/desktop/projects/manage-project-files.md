> Source: https://portswigger.net/burp/documentation/desktop/projects/manage-project-files

Professional

# Managing project files

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Once you create a project file, you can share it with other people. This section explains how to:

-  [Save a copy of a project file](https://portswigger.net/burp/documentation/desktop/projects/manage-project-files#saving-a-copy-of-a-project).

-  [Import a project file](https://portswigger.net/burp/documentation/desktop/projects/manage-project-files#importing-project-files).

-  [Recover data from a corrupted project file](https://portswigger.net/burp/documentation/desktop/projects/manage-project-files#recovering-data-from-corrupted-project-files).


#### Note

Professional
 If you only need to share a small number of HTTP messages (for example, to demonstrate a vulnerability or share test results),
 consider using a collection instead of a full project file.
 A collection contains selected messages from Organizer, and the shareable provides access to those messages without sharing the entire project.
 Collections are easier to share and carry fewer security and configuration risks.

 For more information, see [Collections](https://portswigger.net/burp/documentation/desktop/tools/organizer/collections).


## Saving a copy of a project

 This feature can help you to create a smaller project file after you refine your project scope, or delete unwanted data.

 You can continue to use Burp while the project saves. You may experience delays if you try to perform an operation on data while it is being saved. This is intended to prevent data corruption.

 To save a copy of the project file:

1.
 From the top menu bar, select **Project > Save copy**.

1.
 Select the tools you want to include data for, and whether you want to **Save in-scope items only**. Click **Next**.

1.
 Select whether to include Dashboard data and Burp Collaborator identifiers. Click **Next**.

1.
 Enter a filename, and click **Save**.

1.
 Close the wizard.


### Saving the Burp Collaborator identifier

 When you save a copy of a project file, you can include the [Burp Collaborator](https://portswigger.net/burp/documentation/collaborator) identifier. This enables Burp to retrieve ongoing Burp Collaborator interactions for the project.

 We recommend that you don't include the Burp Collaborator identifier in the following circumstances:

-
 You plan to share the project file with someone but you don't want them to receive details of ongoing Burp Collaborator interactions.

-
 You plan to run the project on two instances of Burp at the same time. This can cause some issues to be missed or reported incorrectly.


## Importing project files

 If you use a disk-based project, you can import another project file into your current project.

#### Warning

 You can import project and configuration files from other users. However, for security reasons, we recommend only importing project and configuration files from trusted sources.


 You can continue to use Burp while the project file is being imported. You may experience delays if you try to perform an operation on data while it is being imported. This is intended to prevent data corruption.

 To import a project:

1.
 From the top menu bar, select **Project > Import project file**.

1.
 Select the project file you want to import, and click **Open**.

1.
 Select the tools you want to import data for, and whether you want to **Pause Automated Tasks** and **Trust this project file**. Click **Next**.

1.
 Wait for the file to import, then click **Close**.


## Recovering data from corrupted project files

 Project files can become corrupted if Burp crashes, or if you open project files created in newer versions of Burp.

 If you try to open a corrupted project file, Burp automatically opens the project file recovery tool:

1.
 Select a location and filename for the data recovery file.

1.
 Click **Recover data** and wait for Burp to recover the data.

1.
 When prompted, click **Open recovered data file**.


 Burp recovers as much data as possible from the damaged project file and stores it in a recovery file. The recovery file may be larger than the original project file.

#### Note

 If you need to recover data from a corrupted project file, we recommend that you treat the recovery file as read only. If you continue to use the recovery file as a project file, you risk an increased chance of crashes and data loss.


 It is not always possible for the project file recovery tool to recover your data from a corrupted file.
