## Setup Guide

This is a guide with all steps to install the Epic Games Launcher and Unreal Engine to get ready to work.

First up you will need to register on [Epic Games](https://www.epicgames.com/id/register?redirect_uri=https%3A%2F%2Fwww.epicgames.com%2Fstore%2Fen-US%2F&client_id=875a3b57d3a640a6b7f9b4e883463ab4).

Then depending on your operating system, you will need the following steps:

### Windows

- On the [Epic Games main page](https://www.epicgames.com/store/en-US/) you can find on the top right corner a button (`Get Epic Games`), click on it.
- An download pop up box should appear, so proceed to download the installer.
- After going through all the steps to install the Epic Games Launcher, open it and sign in using your account.

### Linux

- First of all you will need to install `wine` to run Windows programs on Linux systems.

  - If you have a 64-bit Linux system use: `sudo apt install wine64`
  - If you have a 32-bit Linux system use: `sudo apt install wine32`
  - Check if it's installed correctly with: `wine --version`

- Then, you will need to instal `Lutris` in order to download Epic Games with the following commands:

```sh
sudo add-apt-repository ppa:lutris-team/lutris
sudo apt update
sudo apt install lutris
```

- After `Lutris` is installed, open it and on the top bar, search for `Epic Games` and you will see the result `Epic Games Store`. Go ahead and click on install.

- After installing it, open it and sign up to your account.

### Unreal Engine

- Once the Epic Games Launcher is opened, go to the `Unreal Engine` tab and select the Library tab.
- Click on the plus icon to add `Unreal Engine 4` and choose the version you want to install, by clicking on the dropdown list. Then click on install.
- After installing the Unreal Engine, you can click on `Launch` and prepare for the next adventure.
