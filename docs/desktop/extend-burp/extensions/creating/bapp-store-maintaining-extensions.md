> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating/bapp-store-maintaining-extensions

ProfessionalCommunity Edition

# Maintaining extensions on the BApp Store

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Once your extension has been accepted into the BApp store, it may require occasional updates to maintain its quality and improve its functionality. To support this, please monitor your GitHub repository for any issues or feedback submitted by the community.

#### Note

 We only accept pull requests from the extension's original repository. If you want to fix a bug in someone else's extension, submit a pull request to their fork first. They'll then need to raise a pull request to PortSwigger. For full instructions, see [Troubleshooting extensions - Suggest a fix](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/troubleshooting#option-2-suggest-a-fix).


To update your extension:

1.

Make changes to your extension's code in your original repository.
1.

Create a pull request against PortSwigger's fork of your repository.
1.

Raise an **Update for existing extension** issue on the [extension-portal repository](https://github.com/PortSwigger/extension-portal).

We'll review your changes. If they're approved we'll merge them into the PortSwigger fork. We'll then test the updated extension to make sure it works correctly before publishing it to the BApp Store.
 If we have feedback, we'll provide it directly on the raised issue.

## Using AI to support extension maintenance

 You can use an LLM to help you make changes to your extension more efficiently. To support this, we provide a `CLAUDE.md` file that includes essential context on how Burp extensions are structured for the model.

 Before you start, install or access your preferred LLM. These instructions use Claude Code, but you can adapt them for use with other models. For more information on Claude Code, see the [Anthropic documentation](https://www.anthropic.com/claude-code).

To use Claude Code to update your extension

1.

Download our [CLAUDE.md](https://github.com/PortSwigger/extension-template-project/blob/main/CLAUDE.md) file and
 [supporting documentation](https://github.com/PortSwigger/extension-template-project/tree/main/docs) from GitHub.
1.

Add the `CLAUDE.md` file to the root of your extension project folder.
1.

Create a `docs` folder in your extension project and add the supporting documentation inside it.
1.

Open a terminal and navigate to your extension's project folder.
1.

Run Claude Code using the following command: `claude`
1.

Prompt Claude to update your extension as required.
1.

Review and test the AI-suggested updates.

 Claude should automatically read the contents of the `ExtensionTemplateProject` folder, including the `CLAUDE.md` file, then create draft code for you to review. If you think it hasn't read the `CLAUDE.md` file, directly prompt it to read the file before continuing.

 To use the `CLAUDE.md` file with an LLM other than Claude Code, prompt the LLM to read the file and supporting documentation, or provide their contents as part of your context window.

#### Note

 Make sure to review and test any code suggested by the LLM before submitting it, as it may not always produce correct or secure output.
