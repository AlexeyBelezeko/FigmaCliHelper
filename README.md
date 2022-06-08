# FigmaCliHelper
Small cli tool to help users with some routine sigma actions.

## Actions

For now this tool support only one action:

- Delete all your comments from file

## How to use

Get your personal access token from https://www.figma.com/developers/api#intro

Build tool with go. Or use binary from bin folder according to you OS and Arch.

```shell
# This command will only print count of comments. You can use it to check token and file key. 
./figma-cli-arm64-darwin --token="figd_CW7zSdWGZPQeQVtpdN-9shu70RYrjriZZJRhgdcr" --file-key="2E83Evqg2r226hArzZql5m"
# If you specify userName it will delete all your comments. N.B.: you can't delete comments of other users.
./figma-cli-arm64-darwin --token="figd_CW7zSdWGZPQeQVtpdN-9shu70RYrjriZZJRhgdcr" --file-key="2E83Evqg2r226hArzZql5m" --user="Oleg"

```
