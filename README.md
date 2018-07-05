# StumbleUpon Favorites Rescuer

As you may have already heard StumbleUpon (SU) is shutting down it's service on June 30th. It is sad news for anyone who is/was using SU actively. For some SU also doubled as a bookmark keeping tool. 

As a SU user you may wish to save your data. This small program takes your SU username as input and produces a JSON file that contains all the information regarding pages you like using SU. 

## Build

You could build this program as a command line tool.  Go compiler should be installed to build this program.

From program directory type: 
``` 
go build -o surescue main.go
```

## Usage

Then type from the same directory:
``` 
./surescue su_username
```
Change su_username accordingly

## Caveats

* Only publicly visible SU accounts likes data could be extracted using the tool. If your account is private, before using the tool change your account's visibility to public.
* The tool is not tested for accounts that have over 12K likes. 


### Author
Guvenc Usanmaz

### Licence
This project is licensed under the MIT License.