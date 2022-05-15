from rcon.source import Client

ENDPOINT = "<dns or ip of cs server>"
PORT = 27015
PASSWORD = "<rcon password>"

with Client(ENDPOINT, PORT, passwd=PASSWORD) as client:
    response = client.run('status')
    print(response)
    response = client.run('users')
    print(response)
    response = client.run('stats')
    print(response)
