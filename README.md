# zpt
install with `go get -u github.com/roball24/zpt`

Simple zip file encryption (wrapper on 'zip')

## Usage
```
zpt [files/dirs to zip] -o [output filename]

Enter Password:
Verify Password: 
```

unzip will request for the same password

## Requirements
UNIX 'zip' command. zpt runs zip with encryption on.

golang for installation using 'go get' https://golang.org/dl/

## Development
built using cobra https://github.com/spf13/cobra
