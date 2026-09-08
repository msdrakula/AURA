> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts

ProfessionalCommunity Edition

# Submitting scripts to our GitHub repository

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Sharing your scripts to our [Bambdas repository on GitHub](https://github.com/PortSwigger/bambdas) contributes to a shared library of tools for the entire community to use.

 The submission process is as follows:

1.

Add documentation to your script.
1.

Format and refine your script.
1.

Export your script from Burp.
1.

Validate and submit your script.

#### Note

 The Bambdas repository is for Java-based Bambda scripts. To contribute custom scan checks in the BChecks language, use the BChecks repository. For more information, see
 [Submitting BChecks to GitHub](https://portswigger.net/burp/documentation/scanner/bchecks/contribute-bchecks).


## Step 1: Add documentation to your script

 At the top of your script, add a Javadoc block including the following, in this order:

1.

A short description (1 - 2 sentences) of what the script does.
1.

An `@author` tag in this format: `@author <your_name> (https://github.com/<your_profile>)`. Use a direct, unobscured GitHub profile link.

 This information is used to automatically generate your script's entry in the GitHub directory README.

 You may also include additional notes below the `@author` tag if necessary. These won't appear in the directory README.

### Example Javadoc block
`/**
 * A short description of what the script does.
 *
 * @author <your_name> (https://github.com/<your_profile>)
 *
 * Add further notes here.
 *
 **/
`

## Step 2: Format and refine your script

 Make sure your script meets the following quality standards:

-

It compiles successfully in Burp.
-

It is formatted in a way that is easy to read:

  -

Avoid long lines.
  -

Use consistent code styling.
  -

Avoid using tabs for indentation. Four spaces is preferred.
  -

Use clear, descriptive variable names to reduce the need for comments.

-

It considers performance by avoiding unnecessary complexity or resource usage, which can slow down Burp.
-

It doesn't replicate functionality that already exists in Burp Suite Professional.

#### Note

 For advanced functionality such as complex logic or external API access, consider converting your script to a Burp extension. Extensions offer more flexibility, better performance handling, and full access to Burp's Montoya API.


 For more information, see [Creating Burp extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating).


## Step 3: Export your script from Burp

 Export your script from Burp. The script is exported in YAML format with the required metadata automatically added.

#### Note

 If you wrote your script outside of Burp, we recommend adding it to the Bambda library and exporting it from there. This is the easiest way to make sure metadata is correctly added.


 For instructions on how to add your script to the Bambda library, see [Creating scripts in the Bambda library](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating#creating-scripts-in-the-bambda-library).


 To export your script from Burp:

1.

Save your script to the Bambda library.
1.

Go to **Extensions > Bambda library**.
1.

Select your script.
1.

Select .
1.

Select a directory.
1.

Enter a filename using camel case - capitalize each word and don't use spaces, hyphens, or underscores. For example, `MyCustomScript.bambda`.
1.

Click **Save**.

## Step 4: Validate and submit your script

 Once your script is ready, you'll need to validate it and submit it to the Bambdas GitHub repository for review. The validator checks that your script is in YAML format, includes all required metadata, and starts with the required Javadoc block.

#### Note

 Submits scripts individually. Bundling scripts into a single pull request can delay publication if some require changes while others are ready to merge.


 To validate and submit your script:

1.

In GitHub, fork the [Bambdas repository](https://github.com/PortSwigger/bambdas).
1.

Add your script to the forked repository:

  1.

Go to the appropriate directory for your script type.
  1.

Click **Add file > Upload files**, then select your script. Only upload `.bambda` files. Don't include or modify `README.md` files.
  1.

Add a descriptive commit message and select **Commit changes**.

1.

Run the validation workflow:

  1.

Go to the **Actions** tab.
  1.

Select the **Validate Bambdas** workflow.
  1.

Click **Run workflow** and wait for the results.

1.

Review the workflow results:

  1.

If the workflow passes, open a pull request to the main Bambdas repository. Include a concise description of your script in the request.
  1.

If the workflow fails, check the error message, edit the script as required, then re-run the **Validate Bambdas** workflow.

 Thanks for contributing to the Bambdas repository! We'll review your submission and get back to you with any feedback.
