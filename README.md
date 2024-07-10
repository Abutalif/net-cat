# Description:

This project implements a simple TCP chat that allows up to 10 people to communicate using termnial window. Clone of UNIX util **net-cat** (**nc**)

# Features:

1. Can work as both client and server/host.
2. Compatible with the original **nc**.
3. Ensures uniqueness of usernames.
4. Stores history of the current chat session. Messages visible even after leaving and logging back in. History is shown for all users.
5. Awesome Linux penguin for all new users.


# Usage:
 1. You can run the code using

    ```console
    go run ./cmd/main.go
    ```

    In this case the program will launch in a server mode on localhost in the default port 8989
 

 2. If you want to specify the port you can right your 4 digit port number right after the command above. For example:

    ```console
    go run ./cmd/main.go 1111
    ```


 3. You can also launch the program in the client mode by running the code with a flag "-c" and specifying server's port

    ```console
    go run ./cmd/main.go -c 8989
    ```
 
 4. To join the running server one can simply use "nc" command in their terminal

    ```console
    nc localhost <port_number>
    ```


# Suggestions:
You are welcome to try and break the working of the app. Key points:
1. Try logging in with more than 10 users.
2. Try entering with the same log-in several times.
3. Try sending empty messages.

# Authors

- [@Abutalif](https://github.com/Abutalif)
