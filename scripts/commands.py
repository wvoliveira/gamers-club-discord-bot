import requests

"""Discord-bot: delete guild commands.

The idea is put all theses function in one project.
"""

API_URL = "https://discord.com/api/v8/applications"
TOKEN = "OTY0NzIzNzUyNTMxMjg0MDY4.YlozAg.popVPCZmhJDdudtW8Czv7I3_jzs"
APP_ID = "964723752531284068"  # Gamers Club app
GUIL_ID = "910359107582693436"  # elga.io server
# guild_id = "700713000209743924" # marijuanos

DEFAULT_COMMANDS = [{
    "name": "match",
    "type": 1,
    "description": "Get basic information from match.",
    "options": [
            {
                "name": "match_id",
                "description": "Match number. Ex.: 15693476",
                "type": 4,
                "required": True
            },
        {
                "name": "details",
                "description": "Get details from match.",
                "type": 5,
                "required": False
            }
    ]
}]


def list(token, app_id, guild_id, globall=False):
    url = f"{API_URL}/{app_id}/guilds/{guild_id}/commands"
    if globall:
        url = f"{API_URL}/{app_id}/commands"

    headers = {"Authorization": f"Bot {token}"}

    r = requests.get(url, headers=headers)
    if not (r.status_code >= 200 or r.status_code < 300):
        print(f"status_code != 200: {r.status_code}")
        exit(1)

    commands = []
    for command in r.json():
        commands.append(command.get("id"))

    return commands


def delete(token, app_id, guild_id, command_id, globall=False):
    url = f"{API_URL}/{app_id}/guilds/{guild_id}/commands/{command_id}"
    if globall:
        url = f"{API_URL}/{app_id}/commands/{command_id}"

    headers = {"Authorization": f"Bot {token}"}

    r = requests.delete(url, headers=headers)
    if not (r.status_code >= 200 or r.status_code < 300):
        print(f"status_code != 200: {r.status_code}")
        exit(1)

    return r.status_code


def create(token, app_id, guild_id, commands, globall=False):
    url = f"{API_URL}/{app_id}/guilds/{guild_id}/commands"
    if globall:
        url = f"{API_URL}/{app_id}/commands"

    headers = {"Authorization": f"Bot {token}"}

    for payload in commands:
        r = requests.post(url, headers=headers, payload=payload)
        if not (r.status_code >= 200 or r.status_code < 300):
            print(f"status_code != 200: {r.status_code}")
            exit(1)

    return r.status_code


def main():
    print("Starting script..")
    commands = list(TOKEN, APP_ID, GUIL_ID)
    for command in commands:
        status_code = delete(TOKEN, APP_ID, GUIL_ID, command)
        print(f"command: {command}")
        print(f"status_code: {status_code}\n")


main()
