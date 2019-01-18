# Cloudflare Dynamic DNS 

This app will fetch your current public IP address and update records in cloudflare with that address. You can easily stick
this in a cronjob to automate the updates.

There are other tools out there similar but I didn't see one in golang and wanted to write some golang.


basic usage is:
```
./bin/cfdyndns --api-key=bacon --email=person@example.com --zone=example.com
```

More flags can be found in `./bin/cfdyndns --help`

## Build

To build the app you have to have go 1.11 install and make. You can do a simple `make` after checkout which 
will create the binary in `./bin`
