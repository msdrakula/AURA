> Source: https://portswigger.net/burp/documentation/scanner/bchecks/contribute-bchecks

DASTProfessional

# Submitting BChecks to the community

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 When you write a BCheck for Burp Suite, you can share it with the community. This collaboration enables the community to access an ever-growing library of BChecks, through our [BChecks repository on GitHub](https://github.com/PortSwigger/BChecks).


#### Note

 The BChecks repository is for custom scan checks written in the BChecks language. To contribute Java-based scan checks scripts, use the Bambdas repository. For more information, see
 [Submitting scripts to GitHub](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts).


## Step 1 - Check the submission guidelines

 Before you submit your BCheck, check that it meets the [submission guidelines](https://github.com/PortSwigger/BChecks/blob/main/CONTRIBUTING.md).


## Step 2 - Make a pull request

 Once you're happy with your BCheck, you can create a [pull request in GitHub](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request):


1.
 Export your BCheck from Burp. For more information, see [Exporting custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/managing#exporting-custom-scan-checks).

1.
 Log in to GitHub.

1.
 Create a fork from the [PortSwigger BChecks repository](https://github.com/PortSwigger/BChecks). For more information on creating a fork, see the [GitHub instructions for forking a repo](https://docs.github.com/en/get-started/quickstart/fork-a-repo).

1.
 Clone the fork onto your local machine, and make a new branch for your change.

1.
 Add your BCheck to your new branch, and commit and push your changes.

1.  [Create a pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request) against the PortSwigger BChecks repository.


## Step 3 - We review your BCheck

 We'll review your BCheck using a combination of automated and manual checks, to make sure it meets our [submission guidelines](https://github.com/PortSwigger/BChecks/blob/main/CONTRIBUTING.md). We'll add any feedback to your open pull request.


 Once the review is complete, we'll merge your BCheck into the PortSwigger BChecks repository.
