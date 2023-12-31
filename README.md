# NostrO
[![License: ODbL](https://img.shields.io/badge/License-PDDL-brightgreen.svg)](https://opendatacommons.org/licenses/pddl/)  [![Go Report Card](https://goreportcard.com/badge/github.com/r3drun3/nostro)](https://goreportcard.com/report/github.com/r3drun3/nostro)  [![goreleaser](https://github.com/R3DRUN3/nostro/actions/workflows/release.yaml/badge.svg)](https://github.com/R3DRUN3/nostro/actions/workflows/release.yaml)  [![Latest Release](https://img.shields.io/github/v/release/r3drun3/nostro?logo=github)](https://github.com/r3drun3/nostro/releases/latest)

**Nostr Osint Tool** 🔎 𓅦  
  

<img src="images/logo.png" alt="Nostr Logo" width="250" height="190">  

`NostrO` is a tool designed for conducting *Open Source Intelligence* (OSINT) operations on [Nostr](https://nostr.com/).  
NostrO facilitates operations such as retrieving relay or user infos, search on notes, and more.  



## Development
I welcome pull requests and contributions!  
This tool is a CLI implemented with [cobra](https://github.com/spf13/cobra) and [go-nostr](https://github.com/nbd-wtf/go-nostr) library.  
I suggest using `VS Code` for local development.  
For debugging you can create a `launch.json` file similar to the following (change the command you want to test):  
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go",
            "args": ["relay", "--info", "bitcoinmaximalists.online"],
            "env": {},
            "showLog": true
        }
    ]
}
```  

  



## Examples

<details>
  <summary>Command Help</summary>

```console
nostro --help
Welcome to NostrO 🔎 𓅦

Usage:
  nostro [flags]
  nostro [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dm          Operations on direct messages
  event       Operations on events
  help        Help about any command
  notes       Operations on notes
  relay       Operations on relays
  user        Operations on users

Flags:
  -h, --help   help for nostro

Use "nostro [command] --help" for more information about a command.

```  

```console
nostro relay --help
Retrieve data on nostr relays

Usage:
  nostro relay [flags]

Flags:
  -h, --help   help for relay
      --info   Retrieve relay information document (nip-11)
```
</details>  



<details>
  <summary>Retrieve relay info</summary>
  
  ```console
nostro relay --info relay.nostrview.com
####################### RELAY INFO #######################
NAME:  relay.nostrview.com
DESCRIPTION:  Nostrview relay
PUB KEY:  2e9397a8c9268585668b76479f88e359d0ee261f8e8ea07b3b3450546d1601c8
CONTACT:  2e9397a8c9268585668b76479f88e359d0ee261f8e8ea07b3b3450546d1601c8
SUPPORTED NIPS:  [1 2 4 9 11 12 15 16 20 22 26 28 33 40 111]
SOFTWARE:  git+https://github.com/Cameri/nostream.git
VERSION:  1.22.2
LIMITATION:  &{524288 10 10 5000 256 4 2500 102400 0 false true}
PAYMENTSURL:  https://relay.nostrview.com/invoices
##########################################################
```  
</details>  



<details>
  <summary>Retrieve user info from the specified relay</summary>

```console
nostro user --info npub1rusgp3upyrtpsy2pcqznl6e8hejg9ne8u2eg05gzc4n2cctsugksvcx2np nos.lol
####################### USER INFO #######################
[[i github:R3DRUN3 0f954e6fada304dacdb8e7389eefaf2b]]
Name: r3drun3
Picture: https://i.postimg.cc/rwTgJm0G/symbol-1.gif
Username: r3drun3
Display Name: r3drun3
Banner: https://i.postimg.cc/90FYS0D7/1327483.png
Website: 
About: ᚱᛊᚧᚱVᚺᛊ
Jus a collection of quantum bits,
constantly phasing between cyberspace and meatspace.
Running #Bitcoin
Nip05: r3drun3@vlt.ge
Lud16: me@ln.stackzoo.io
Lud06: 
Created At: 1689593935
Nip05 Valid: false
##########################################################
```
</details>  


<details>
  <summary>Retrieve the specified event from the specified relay</summary>

```console
nostro event --info note1se5g5crjxaaet4vzy3xtpurv4as3dsfd5dteglk4z3f2xafstl5qyry4m3 nos.lol

####################### EVENT INFO #######################
ID: 86688a6072377b95d582244cb0f06caf6116c12da357947ed51452a375305fe8
PubKey: 1f2080c78120d6181141c0053feb27be6482cf27e2b287d102c566ac6170e22d
Kind: 1
Created At: 1696529147
Tags: [[t osint] [t osint] [t Nostr] [t nostr] [t github] [t github] [r https://github.com/r3drun3/nostro] [r https://image.nostr.build/7a83f1b7b006bdecd047731c6b0fcec54d1a5186ae222f3e98e15953850712f4.jpg]]
Content: I believe that one of the best ways to learn a technology is to experiment and build upon it. That's why I've started developing a tool for performing #osint operations on #Nostr on my #github. 
Feel free to collaborate if you want ☺
https://github.com/r3drun3/nostro


https://image.nostr.build/7a83f1b7b006bdecd047731c6b0fcec54d1a5186ae222f3e98e15953850712f4.jpg
Signature: 9b3b4af0bac8df5f62dd54b8f5be34bdee7545e0a6453fe6e3462861d29390282e95a4e85f8d2bf801d1f0da3ccc955b3ecff0ffbc6786ffa7d1c7017650b34a
##########################################################
```
</details>  




<details>
  <summary>Retrieve a user contact list from the specified relay</summary>

```console
nostro user --contactlist npub1rusgp3upyrtpsy2pcqznl6e8hejg9ne8u2eg05gzc4n2cctsugksvcx2np nos.lol
returned events saved to user_contact_list.json
```
</details>  



<details>
  <summary>Retrieve from the specified relay the last 300 direct messages that the specified user received</summary>

```console
nostro dm --userreceived npub1rusgp3upyrtpsy2pcqznl6e8hejg9ne8u2eg05gzc4n2cctsugksvcx2np nos.lol
returned events saved to user_received_direct_messages.json
```
</details>  

<details>
  <summary>Retrieve from the specified relay the last 300 notes that the specified user wrote</summary>

```console
nostro notes --userwritten npub1rusgp3upyrtpsy2pcqznl6e8hejg9ne8u2eg05gzc4n2cctsugksvcx2np nos.lol
returned events saved to user_written_notes.json
```
</details>  


<details>
  <summary>Retrieve from the specified relay the last 300 notes in which the specified user has been tagged</summary>
  
```console
nostro notes --usertagged npub1rusgp3upyrtpsy2pcqznl6e8hejg9ne8u2eg05gzc4n2cctsugksvcx2np nos.lol
returned events saved to user_tagged_notes.json
```  
</details>  

<details>
  <summary>Retrieve from the specified relay the last 300 notes from the specified user that have been reposted</summary>
  
```console
nostro notes --userreposted npub1rusgp3upyrtpsy2pcqznl6e8hejg9ne8u2eg05gzc4n2cctsugksvcx2np nos.lol
returned events saved to user_reposted_notes.json
```  
</details>


<details>
  <summary>Retrieve from the specified relay the last 300 reaction received by notes from the specified user</summary>
  
```console
nostro notes --userreacted npub1rusgp3upyrtpsy2pcqznl6e8hejg9ne8u2eg05gzc4n2cctsugksvcx2np nos.lol
returned events saved to user_reacted_notes.json
```  
</details>  

<br/>  

For all available command use the cli `help` function.

