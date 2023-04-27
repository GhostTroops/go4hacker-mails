# go4hacker-mails

Bulk email collection and retrieval

## Usage
searchText eg:
```
password
passwd
admin
vpn
```
50 50 indicates that the output will be searched to 50 bytes after the text is found

go4hacker-mails_macos_amd64 lists.txt searchText 50 MailServer
eg:
popmails lists.txt password 50 pop3.seclover.com

#### out results save to file:
SMResults.txt
```
lists.txt
format: user[tab]pswd
```
user1   pswd1
user2   pswd2
user3@xx.com    pswd1
username1 pswd1
username2 pswd2
username3 pswd3
```


