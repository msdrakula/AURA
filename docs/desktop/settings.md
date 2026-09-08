> Source: https://portswigger.net/burp/documentation/desktop/settings

ProfessionalCommunity Edition

# Settings

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 You can access most of Burp Suite's settings via the **Settings** dialog. To access this dialog, click **Settings** on the top menu.


## User and project settings

 There are two types of setting in Burp:


- **Project settings** only apply to the current project. They are stored within the [project file](https://portswigger.net/burp/documentation/desktop/projects) itself.
- **User settings** apply to all installations of Burp on your machine. They affect all disk-based projects and any temporary projects.

 Some settings can be defined as both project and user settings. These settings have an **Override options for this project only** toggle. If this toggle is selected, then Burp applies the specified settings at the project level. Otherwise, the specified settings are treated as user settings and applied globally.


 The toggle enables you to configure defaults at the user level, and then override these settings on a project by project basis.


 For example, you might configure a corporate LAN proxy to connect to the Internet as a user setting. However, some projects might require you to use a different upstream proxy. You can configure the alternative details in your project settings.


#### Warning

 You can import project and configuration files from other users. However, for security reasons, we recommend only importing project and configuration files from trusted sources.


## Finding settings

 The navigation tree to the side of the **Settings** dialog enables you to find the settings you are looking for.


 To filter the content:


- Use the **Search** bar at the top of the panel.
- Alternatively, use the **All**, **User** or **Project** filter buttons.

#### Note

For a guide to some of the most important settings in Burp, see [Key settings](https://portswigger.net/burp/documentation/desktop/settings/key-settings).

## Managing settings

 From the **Settings** dialog, you can:


- Restore default settings.
- Save settings. The settings are saved as a configuration file in JSON format.
- Load settings. The configuration file needs to be in JSON format. The easiest way to generate the file is to create the desired configuration in Burp, then save a file from it.

 To manage all user or all project settings:


1. Click **Manage global settings**.
1. Choose between **User settings** or **Project settings**.
1. Select **Restore default settings**, **Save settings**, or **Load settings**.

 If you select **Restore default settings** for **Project settings**, you can choose which settings to restore defaults for:


- **All** - All project settings.
- **Target** - The [target scope](https://portswigger.net/burp/documentation/desktop/settings/project/scope#target-scope) and the
 [site map filter](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-filter).
- **Proxy** - All [Proxy settings](https://portswigger.net/burp/documentation/desktop/settings/tools/proxy).
- **Repeater** - All [Repeater settings](https://portswigger.net/burp/documentation/desktop/settings/tools/repeater).
- **Sequencer** - All [Sequencer settings](https://portswigger.net/burp/documentation/desktop/settings/tools/sequencer).
- **Logger** - The [capture filter](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/capture) and
 [view filter](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/view) for Burp Logger.
- **Other project settings** - All project settings that are not Target, Proxy, Repeater, Sequencer, or Logger settings.

 If you manage user settings, your action also applies to extensions and UI related settings, such as the [
 site map layout](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-layout).


 If you manage all project settings, your action also applies to non-UI-related settings in Burp tools.


 To manage an individual setting, click the icon  next to the setting and select an action.


#### Related pages

[Configuration files](https://portswigger.net/burp/documentation/desktop/settings/library#configuration-files).


## Settings pages

 The **Settings** dialog contains the following pages:


- [Tools](https://portswigger.net/burp/documentation/desktop/settings/tools).
- [Project](https://portswigger.net/burp/documentation/desktop/settings/project).
- [Sessions](https://portswigger.net/burp/documentation/desktop/settings/sessions).
- [Network](https://portswigger.net/burp/documentation/desktop/settings/network).
- [User interface](https://portswigger.net/burp/documentation/desktop/settings/ui).
- [Suite](https://portswigger.net/burp/documentation/desktop/settings/suite).
- [AI](https://portswigger.net/burp/documentation/desktop/settings/ai).
- [Extensions](https://portswigger.net/burp/documentation/desktop/settings/extensions).
- [Configuration library](https://portswigger.net/burp/documentation/desktop/settings/library).

 The [response extraction rules](https://portswigger.net/burp/documentation/desktop/settings/response-extraction) are used in various settings in Burp.
