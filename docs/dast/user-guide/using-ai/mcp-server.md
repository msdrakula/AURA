> Source: https://portswigger.net/burp/documentation/dast/user-guide/using-ai/mcp-server

DAST

# Connecting your AI client to the MCP server

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 The Model Context Protocol (MCP) server lets your AI assistant, such as Claude or Cursor, work with Burp Suite DAST. Once connected, you can work with your scans and issues in natural language, using the same permissions and data your Burp Suite DAST account already has.


 The MCP server runs inside your existing Burp Suite DAST web server and is enabled by default. There is nothing extra to install or switch on: you point your AI client at it.


## Prerequisites

 To connect an AI client to the MCP server, you need the following:


- A licensed Burp Suite DAST instance.
- An API user account and an API key. An administrator can create these for you (see [Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user)).
- An MCP-compatible AI client, such as Claude Code, Claude Desktop, or Cursor.

## Connecting Claude Code

 To connect a client other than Claude Code, see [Connecting other clients](https://portswigger.net/burp/documentation/dast/user-guide/using-ai/mcp-server#connecting-other-clients).


1.

 Store your API key in a `DAST_API_KEY` environment variable. Set the variable in the environment that starts your AI client, so that the client can read it.

1.

 Add the MCP server to your project directory in one of the following ways:


  -
 Create an `.mcp.json` file:
 `{
  "mcpServers": {
    "burp-dast": {
      "type": "http",
      "url": "https://<your-dast-hostname>/mcp/v1",
      "headers": {
        "Authorization": "Bearer ${DAST_API_KEY}"
      }
    }
  }
}`
  -
 Run the `claude` CLI, at project scope:
 `claude mcp add --transport http burp-dast https://<your-dast-hostname>/mcp/v1 \
  --scope project \
  --header 'Authorization: Bearer ${DAST_API_KEY}'`

 Enclose the header in single quotes. Double quotes tell your shell to replace `${DAST_API_KEY}` with your API key, which writes the key into `.mcp.json` in plain text.


 Both methods create the same `.mcp.json` file.

1.

 Start Claude Code in your project directory, and approve the new server when Claude Code prompts you. Claude Code does not connect to a server listed in `.mcp.json` until you approve it. To check the status of the server, run `claude mcp list`.

1.

 Confirm the connection: ask your assistant to list your sites. If it returns your real site names, you are connected.


#### Note

 Treat your API key like a password: anyone with it can act as you through the MCP server. Teams often commit `.mcp.json` to source control, so make sure that the file refers to your `DAST_API_KEY` environment variable and never contains the key itself.


## Connecting other clients

 Any MCP-compatible client can connect to the MCP server. Configuration formats differ between clients, so to connect a client other than Claude Code, refer to its documentation for how to add a remote MCP server, and supply the following details:


- Transport: streamable HTTP.
- URL: `https://<your-dast-hostname>/mcp/v1`
- Authentication: an `Authorization` header with the value `Bearer <your-api-key>`.

 If your client's configuration format is unfamiliar, you can ask an AI assistant to build the configuration from these details and your client's own documentation. Refer to an environment variable in your prompt, and never paste your API key into it.


#### Note

 Use the exact `/mcp/v1` path. Any other path returns the Burp Suite DAST web interface instead of an error, so an incorrect URL can be hard to spot.


## Limitations

 If your Burp Suite DAST instance restarts, your client's connection needs to be re-established. Most MCP clients reconnect automatically.


#### Related pages

- [Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user)
- [API overview](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation)
- [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control)
