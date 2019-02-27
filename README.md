# Go GameServer
Is a package to build easly a simpe game server
but gives also the power to give a complex realtime server for any tipe of function
that has a lot of user who constantly communicate with each other

### Temp Doc

###### About exported functions 
**Std** Function that containt *Std* In the name are to create standard Interfaces with zero configuration
**New** Functions that contain *New* in the name are to create an interface wit full configuration
*Std* uses the *New* functions
###### About exported interfaces
**Server** Is the interface of the server, it allows to create rooms, add players,send events...
*Server* is a reference to a struct that contains also all the connected players
**Player**Is the player, unlike objects *coming soon*, there can be only one player per connection
the struct holds the connection and info of the player
**Message**Is the interface for the message, it allow to create and send different type of messages
there two types of messages, one in wich both client and server aggree on the message structure by the first byte of the message, 
and one other in wich the client doesent know the structure untile it reads it