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
 
 * A bare bones frontend for Telepathy built with React and Next JS is available [here](https://github.com/lajosdeme/telepathy-frontend). (This has to be run locally for now, but can easily be [deployed to IPFS later](https://github.com/Velenir/nextjs-ipfs-example).)
 * When an action is executed on the front-end the request is made through an Nginx reverse proxy because CORS policies are disabled by default for Cosmos chains because of security. More info on this can be found [here](https://docs.cosmos.network/v0.39/interfaces/rest.html).
 * Thoughts, comments etc. are all recorded as transactions on the Telepathy blockchain.
 * When a user sets an avatar or uploads a picture, we upload it to IPFS, and once we receive the CID, we pass that CID with a transaction to our blockchain. When the image has to be retreived we request it from IPFS with the CID stored on chain.

## Get started
Before running Telepathy you have to take a few preliminary steps:
1. Make sure **Go 1.16+** is installed on your workstation
2. Install [starport](https://docs.starport.network/intro/install.html) (Optional)
3. Install the [IPFS](https://docs.ipfs.io/install/command-line/) CLI
4. Install [Nginx](https://www.nginx.com/resources/wiki/start/topics/tutorials/install/)
5. Edit your ```nginx.conf``` file to include the contents of the ```/scripts/nginx/nginx.conf``` file 
6. Run the ```scripts/ipfs-config``` script to config IPFS
7. In your terminal run ```nginx``` and ```ipfs daemon``` to start the nginx server and a local IPFS node
8. Clone the repo:
```
git clone github.com/lajosdeme/Telepathy.git
```
Now you can ```cd``` into the repo and proceed with:
```
make
```
or
```
starport serve
```
(If you're not using ```starport``` you also have to run ```telepathyd init test --chain-id=telepathy```)

You will see that two wallets are created with their respective mnemonics. You can use these to log in to the front-end web app.

Run ```telepathyd help``` and ```telepathycli help``` to make sure everything is ok.

### `accounts`
Initialization parameters are stored in `config.yml`.

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |


### Docker Images And Pi Images

In order for Docker images and Raspberry Pi images to build successfully, please add your docker hub credentials as [secrets](https://github.com/lajosdeme/telepathy/settings/secrets/actions)

Add these:

DOCKERHUB_USERNAME
DOCKERHUB_TOKEN

You can get the token [here](https://hub.docker.com/settings/security)

## Try Telepathy with the CLI

Run ```telepathyd start``` to spin up a local node.

### Creating a user
The best place to start is to create a user for one of your addresses. Each address can have one user, with an avatar, username and bio.
```
telepathycli tx telepathy create-user "alice" "This is a short bio for alice" --from alice
```

### Creating a thought
Thought are just like tweets and you can create one like this:
```
telepathycli tx telepathy create-thought "Posting my first thought on Telepathy" --from alice
```

## Try Telepathy as a web app

Telepathy has a bare bones (yet) front-end which makes it easy to interact with the underlying blockchain.

Get it with:
```
git clone github.com/lajosdeme/telepathy-frontend.git
```

Then ```cd``` into the folder and start the rest server:
```
telepathycli rest-server --chain-id telepathy --trust-node --unsafe-cors
```
Run one of the start scripts like ```npm run dev```.
You can now go to ```localhost:3000``` and use the app.


## Roadmap
#### Detecting offensive behaviour
 * The current version contains no moderation mechanism for weeding out hate speech, racism, violence and other problematic user behaviour. On Twitter this is addressed by a centralized force deciding what is allowed and what is not. There are various options for solving this in a decentralized context. I propose one such solution here, but there might be other way better options to work this out.
 * We could have an open-source algorithm that analyses thoughts and determines whether they are offensive or not. This algorithm would be available to inspect and propose modifications to by everyone. Each such modification should be voted and and approved/rejected by the Telepathy community.
 * When a user wants to share a thought (that is, make a transaction) the system can use an oracle (like that provided by [Chainlink](https://chain.link/)) that using the above mentioned algorithm can securely verify whether the thought is offensive or not. 
 * If it is offensive, the transaction is rejected and the user incurs a penalty, thereby discouraging this type of behaviour.

#### Importing tweets
* We all already use Twitter, and it would be very hard to leave that platform behind and migrate to a web 3.0 solution like Telepathy. One roadblock is that people wouldn't want to leave the tweets they have crafted over the years behind. This can be solved by using an oracle (again, possibly from [Chainlink](https://chain.link/)) to query the Twitter API and import all your tweets as thoughts into Telepathy.
* This would also make sure that the wisdom and humour that we humans have generated on Twitter over the years is taken from that walled garden and preserved on the blockchain as our collective heritage.
##### Other
- [ ] Resharing thoughts posted by other users
- [ ] Adding photos to thoughts
- [ ] Redesigning the frontend to look cooler
 
 
