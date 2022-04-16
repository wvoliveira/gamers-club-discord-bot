import requests


url = "https://discord.com/api/v8/applications/964723752531284068/guilds/910359107582693436/commands"

json = {
    "name": "demo",
    "type": 1,
    "description": "Parse and analyse demos (aka replays).",
    "options": [
        {
            "name": "match",
            "description": "Match number. Ex.: 15693476",
            "type": 4,
            "required": True
        }
    ],
    "name": "match",
        "type": 1,
        "description": "Get basic information from match.",
        "options": [
            {
                "name": "match",
                "description": "Match number. Ex.: 15693476",
                "type": 4,
                "required": True
            }
        ]
}

# For authorization, you can use either your bot token
headers = {
    "Authorization": "Bot OTY0NzIzNzUyNTMxMjg0MDY4.YlozAg.popVPCZmhJDdudtW8Czv7I3_jzs"
}

r = requests.post(url, headers=headers, json=json)

print(r.status_code)
print(r.text)
