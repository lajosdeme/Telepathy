# Telepathy
![telepathy_logo](https://user-images.githubusercontent.com/44027725/120887037-0fee5580-c5f1-11eb-803e-ec29fa00997c.png)


**Telepathy** is a blockchain built with the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk). It aims to be a decentralized microblogging platform free from censorship, ads and middlemen.

While the Telepathy blockchain is functional and can be used, tested, tweaked and experimented with, it is still in development mode and has a long way to go before being used in production. (Though I hope that one day it will be.)

## Motivation
[Twitter](https://twitter.com) is an amazing platform for people all around the world to share their thoughts and ideas with everyone else, however it is rooted in the web 2.0/society 4.0 system of the world, and thus it has some serious handicaps, which can be remedied with the use of blockchain technology. These handicaps are mainly cemented in the centralized authority that Twitter is, mediating the thoughts of more than 300 million people and at the end of the day deciding who gets to say what. Twitter was an incredible invention for the web 2.0 era, however, as we move more and more to web 3.0 we need a thought sharing platform that is in sync with the spirit of the times. 

<b>tele·pathy</b> /tɪˈlɛpəθi/
<i>[noun]</i>
Communication of thoughts or ideas from one mind to another without the normal use of the senses. <br> <i>(Source: Oxford Advanced Learner's Dictionary)</i>

In the case of telepathy, you convey a thought without using speech or writing, while with Telepathy, you do it without using any third party.

## Architecture
![architecture2](https://user-images.githubusercontent.com/44027725/120889990-57c8a900-c600-11eb-8392-3a964c791920.jpg)
 
 * A bare bones frontend for Telepathy built with React and Next JS is available [here](https://github.com/lajosdeme/telepathy-frontend). (This has to be run locally for now, but can easily be deployed to IPFS later.)
 * When an action is executed on the front-end the request is made through an Nginx reverse proxy because CORS policies are disabled by default for Cosmos chains because of security. More info on this can be found [here](https://docs.cosmos.network/v0.39/interfaces/rest.html).
 * Thoughts, comments etc. are all recorded as transactions on the Telepathy blockchain.
 * When a user sets an avatar or uploads a picture, we upload it to IPFS, and once we receive the CID, we pass that CID with a transaction to our blockchain. When the image has to be retreived we request it from IPFS with the CID stored on chain.

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

### UI on Github Pages

Click the link below, and scroll down until you see it get her pages. Then, select the branch gh-pages.

[Github Pages Setings](https://github.com/(lajosdeme/telepathy/settings/)

After you do that you can visit your chain's UI at:

https://lajosdeme.github.io/telepathy

This is especially useful when you would like to rapidly iterate on a live user interface. Remember, each community member can have their own github pages instance, allowing your community to mix-and-match front ends.

### CI

By default, this chain includes a github action that builds for amd64 and arm64 on Windows, Mac, and Linux.

### Docker Images And Pi Images

In order for Docker images and Raspberry Pi images to build successfully, please add your docker hub credentials as [secrets](https://github.com/lajosdeme/telepathy/settings/secrets/actions)

Add these:

DOCKERHUB_USERNAME
DOCKERHUB_TOKEN

You can get the token [here](https://hub.docker.com/settings/security)

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
