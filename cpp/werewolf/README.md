# WereWolf

To compile the server code:
```bash
g++ server.cpp -o server -std=c++11
```

To compile the client code:
```bash
g++ client.cpp -o client -std=c++11
```

This will generate an executable file named `client`. The `-std=c++11` flag is necessary as it involves using threads.

## Execution

Run the server using:
```bash
./server
```

Run the client using:
```bash
./client
```

## Note

You need to have six clients connected to the server since the game setup is for six players. You can also modify the game's number of players within the server's code.
