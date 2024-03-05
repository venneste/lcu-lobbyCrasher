# Lobby Crasher for League of Legends
Dodge your ranked games without any penalty

## But why you leaked it? ##
_It doesn't work since 14.4 patch, and then I decided to make this project public.
Code has some inconsistencies and I don't want to improve it since Riot fixed this exploit.
Use this project for study purposes only_

## How DID it work?
1. Quit your current gameflow
2. Get these values through LCU requests:
 - Game version
 - RSO Inventory JWT
 - RSO Id token
 - RSO Access token
 - League Session token
3. Generate custom lobby based on your gotten values
4. Send custom lobby to the client
5. Start champion selection stage in this custom lobby
6. Wait some time (12-15 seconds is enough I guess)
7. Again, quit your current gameflow
8. Send `setClientReceivedGameMessage` LCDS request to `gameService` destination
