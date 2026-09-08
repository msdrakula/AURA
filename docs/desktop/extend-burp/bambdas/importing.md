> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/importing

ProfessionalCommunity Edition

# Importing scripts into your Bambda library

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can import scripts into your Bambda library that have been shared with you or downloaded from our [Bambdas
 GitHub repository](https://github.com/PortSwigger/bambdas). Once imported, scripts can be loaded and applied in compatible tools across Burp and in different projects.

#### Warning

 Bambda scripts can run arbitrary code. For security reasons, please be cautious when using scripts from unverified sources.


To import scripts into your Bambda library:

1.

Go to **Extensions > Bambda library**.
1.

Click **Import**. The **Import scripts** dialog opens.
1.

Select `.bambda` files or a folder containing `.bambda` files.
1.

Click **Open**.

Burp adds the selected files to your Bambda library. If you select a folder, Burp identifies any `.bambda` files within the folder and its subfolders and adds them to your library.

#### Note

 If you import a script that defines a scan check function, Burp also adds it to your custom scan checks library, under
 **Extensions > Custom scan checks**. The new scan check is immediately available from the scan launcher.


## Importing the full GitHub repository

 To quickly import all scripts from our [GitHub repository](https://github.com/PortSwigger/bambdas):

1.

Download the GitHub repository as a ZIP file.
1.

Extract the ZIP file contents.
1.

In **Extensions > Bambda library**, click **Import**. The **Import scripts** dialog opens.
1.

Select the extracted folder containing the GitHub repository files.
1.

Click **Open**.
1.

Burp identifies all `.bambda` files in the folder and its subfolders and adds them to your library.

## Updating your scripts

 If your scripts have been modified outside Burp, you can re-import them. Burp gives you the option to replace existing scripts with the new versions.

#### Note

 As scripts in our GitHub repository may be updated frequently, we recommend re-importing these regularly to keep your library current.


### How Burp overwrites scripts

 When a script is created in or imported into Burp (version 2025.2 onwards), it's assigned its own unique ID. The unique ID isn't visible in Burp, but it acts as a behind-the-scenes mechanism to manage conflicts and synchronize scripts that are re-imported. If you export scripts from Burp, the unique ID is included in the metadata.

 When you import a script, Burp checks its unique ID to determine if a matching script exists in your library. If the unique IDs match, you're given the option to overwrite the existing script.

#### Note

 All scripts in our [Bambdas GitHub repository](https://github.com/PortSwigger/bambdas) have a unique identifier, regardless of when they were created.
