# StumbleUpon Favorites Rescuer

As you may have already heard StumbleUpon (SU) is shutting down it's service on June 30th. It is sad news for anyone who is/was using SU actively. For some SU also doubled as a bookmark keeping tool. 

As a SU user you may wish to save your data. This small program takes your SU username as input and produces a JSON file that contains all the information (most probably you won't need!) regarding pages you like using SU. 

## Build

You could build this program as a command line tool. 

From program directory type: 
``` 
go build -o surescue main.go
```

## Usage

Then type from same directory:
Change su_username accordingly
``` 
./surescue su_username
```

### Author
Guvenc Usanmaz

### Licence
This project is licensed under the MIT License.