# telepathy

**telepathy** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

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
