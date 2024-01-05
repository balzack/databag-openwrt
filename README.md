This is just a copy of Databag formatted for building within OpenWrt. The browser app has been built and resulting static web files copied to the web folder. The server is left as source to be built within the OpenWrt build system.

The device admin should update the configuration files specifying the service port and persistent storage path with the files '/etc/databag/service_port.txt' and '/etc/databag/store_path.txt'.

I hope to figure out how to have the Makefile do this work directly on Databag and eventually delete this repo.
