# walk issue reproduction

I put together this repo to demonstrate two issues that I've encountered with [tailscale/walk](https://github.com/tailscale/walk). The [parent fork](https://github.com/lxn/walk) does not demonstrate these issues.

* The `tab` key cannot be used to move between elements in a main window
* There is an "overlay" bar at the top of the MainWindow. This happens most of the examples (which use MainWindow) that I've tested.

## Demonstration

The video below was run (`go run .`) on Windows Server 2022 using Go 1.24.1 via RDP.

https://github.com/user-attachments/assets/ae83a334-c4f2-40ca-a4b8-7bcf86703578



The video below was run (`GOOS=windows go build -o test.exe . && wine64 test.exe`) locally on Linux.

https://github.com/user-attachments/assets/487a9ebf-0809-45fc-b71d-0e3884b6c29c

