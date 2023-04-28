[![Tweet](https://img.shields.io/twitter/url/http/Hktalent3135773.svg?style=social)](https://twitter.com/intent/follow?screen_name=Hktalent3135773) [![Follow on Twitter](https://img.shields.io/twitter/follow/Hktalent3135773.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=Hktalent3135773) [![GitHub Followers](https://img.shields.io/github/followers/hktalent.svg?style=social&label=Follow)](https://github.com/hktalent/)

# go4hacker-mails

Bulk email collection and retrieval，Assist penetration and redness teams in searching for sensitive information in emails for multiple accounts, such as admin, VPN, and password information. All searches are completed in memory without landing, thus avoiding their own risks
帮助渗透、红对团队在若干账号的邮件中搜索敏感信息，例如admin、vpn账号及密码信息，所有的搜索在内存中完成，不落地，从而规避自己的风险

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

#### lists.txt
format: user[tab]pswd
```
user1   pswd1
user2   pswd2
user3@xx.com    pswd1
username1 pswd1
username2 pswd2
username3 pswd3
```


## 💖Star
[![Stargazers over time](https://starchart.cc/hktalent/go4hacker-mails.svg)](https://starchart.cc/hktalent/go4hacker-mails)

# Donation
| Wechat Pay | AliPay | Paypal | BTC Pay |BCH Pay |
| --- | --- | --- | --- | --- |
|<img src=https://raw.githubusercontent.com/hktalent/myhktools/main/md/wc.png>|<img width=166 src=https://raw.githubusercontent.com/hktalent/myhktools/main/md/zfb.png>|[paypal](https://www.paypal.me/pwned2019) **miracletalent@gmail.com**|<img width=166 src=https://raw.githubusercontent.com/hktalent/myhktools/main/md/BTC.png>|<img width=166 src=https://raw.githubusercontent.com/hktalent/myhktools/main/md/BCH.jpg>|

