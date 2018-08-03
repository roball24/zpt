# zpt
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

## Development
built using cobra https://github.com/spf13/cobra
