# Description:

This project implements a TCP chat that allows up to 10 people to communicate using termnial window.

# Authors

- [@Abutalif](https://01.alem.school/git/Abutalif) and [@mixturegg](https://01.alem.school/git/mixturegg)


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
