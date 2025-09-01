# Tic-Tac-Tournament

Multiplayer Code vs. Code Tic-Tac-Toe server.

## Gameplay

Each player gets an `ID` and some form of authentication (token for example) from the event organizers.

`Active` player clients call the Ready API to signal their presence. The server automatically creates lobbies and organizes matches.

`Scheduled` players should only call `Get grid`, until the game starts. When the game lobby is ready and `running`, players become `in-game`, and should call `Get grid` and `Place` to play the game.

When a lobby's status changes to `finished`, all players must call `Ready` to indicate they are ready again.

## API

### Get player information

`GET /playerinfo/{ID}`

Returns the most important upcoming game info and player statistics.

- when the next game is scheduled (?)
- wins and losses count

### Ready

`GET /ready/{ID}`

Signal the server, that the player is ready for a match, like a heartbeat. Authentication required.

Recommended refresh period: 1-60s. Players idling for more than 60s will be marked as `inactive`.

If the player is scheduled for a match, the response includes a timestamp and lobby ID.

### Get grid

`GET /getgrid/{lobby ID}`

Get grid status. Response includes
- current game map
- status (pending, running, finished)
- if the player is expected to place its mark

A client's reaction time is measured from the first `Get grid` call, when it's expected to make a move.

### Place

`POST /place/{lobby ID}`

Place the player's mark.

Timeout, or placing the mark on an illegal coordinate immediately ends the round and the opponent wins.
