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


## Features

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

## Todo