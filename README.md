<h1 align="center">
  Discord Blacklist Bot
</h1>

<p align="center">
  <a href="#about">About</a>
  •
  <a href="#features">Features</a>
  •
  <a href="#commands">Commands</a>
  •
  <a href="#setup">Setup</a>
  •
  <a href="#todo">Todo</a>
</p>

## About

This Discord bot can automatically remove images that are posted.
Once the image has been checked against the database and identified as being banned so will it be removed.

The bot checks images with four techniques and can thus detect alterations, making attemps to fool the bot and bypass the blacklist difficult.

Each individual detection thresholds can be edited in the configuration file.

Detection methods used:
1. [SHA-1 Hash](https://en.wikipedia.org/wiki/SHA-1)
2. Average hashing
3. Difference hashing
4. Perception hashing

You can read more about these methods [here](https://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)

All data is stored in a local SQLite database

## Features

* Automatic removal of inappropriate images

* Smart detection that can detect subtle changes

## Commands

* Ban

Bans the provided image. Can be an URL or uplaoded as an attachment

Supported formats are: PNG & JPEG

* Unban 

(Not implemented yet)

* Reload

Updates the settings from the configuration file

## Setup

Clone the repo

```
git clone https://github.com/CarlFlo/blacklisterBot.git
```

Install all the requirements

```
go mod download
```

### Configuration

Running the bot will create a `config.json` file in the directory where it was run.

Insert:
1. The bot token
2. The Discord user ID for the user(s) that are allowed to run the commands
3. The bots ID

## Todo

- [X] Basic functionality
- [ ] List of channels that the bot will ignore alt. listen to
- [ ] Make the bot aware of what discord server it is in
- [ ] Automated action once banned image is posted (banned, kick, timeout etc)