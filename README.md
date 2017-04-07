send-msg
=======

Service will: 
- watch specific dir (default is "source") for new .txt files
- read text from each file
- send a http post request (application/json) for each file to specific url (default is path of "receiver.go" which is "http://localhost:6061/api/msg")
- remove each file from watch dir after text is sent

Usage (skipping compiling and running native binary)
----------
    1. open terminal
    2. cd into projects root dir (send-msg)
    3. issue command: go run main/sender.go
    4. open new terminal
    5. cd into projects root dir (send-msg)
    6. issue command: go run main/receiver.go
    7. copy LoremIpsum.txt into "source" dir
    8. watch in second terminal for incoming msg
    
How to change watching file and/or changing url
-----------------------------------------------

- For step 3 above, issue command: go run main/sender.go -target=your/path/to/watch/dir -sendTo=your/url
    
