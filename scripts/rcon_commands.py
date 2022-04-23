from rcon.source import Client

ENDPOINT = "borajogarcs.ddns.net"
PORT = 27015
PASSWORD = "CCWJu64ZV3JHDT8hZc"

with Client(ENDPOINT, PORT, passwd=PASSWORD) as client:
    response = client.run('status')
    print(response)
    response = client.run('users')
    print(response)
    response = client.run('stats')
    print(response)

