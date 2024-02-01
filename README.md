# Chat

a chat room app

## Funtionality

A user can create a instance or join one by typing in the id.

## How to try it out

1. clone the repo
2. install go if you dont have it installed
3. navigate to `chat/server`
4. make sure all the dependencies are there with `go mod tidy`
5. start the server with `go run .`
6. open `client/index.html` in your browser (clientId = 0)
7. open a second tab with `client/index.html`(clientId = 1)
8. open the terminal in booth tabs
9. sent messages by writing in the text box using the format `clientId; message`and then click send. For example "0;Welcome to the chat!"
10. you can now see the message in the other terminal
