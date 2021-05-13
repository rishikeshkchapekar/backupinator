# Backupinator

## Description

A tool for auto backing up storage devices to another storage device

## Compatibility

Binary built for Raspberry Pi OS (built for Buster, so it should work for all previous versions on Raspberry Pi 3B+ and below) and Ubuntu 20.04 (should work for all 64 bit systems running atleast Ubuntu 16.04)

Not built for Windows (don't intend to either, but contributions are most welcome)

## How to use

1. This repo contains a ".identifier" file. The contents of this file are irrelevant, however the file needs to be placed in the intended backup device.
2. Plug in your backup device (external HDD,external SSD, SD Cards whatever) and the devices you intend to backup (SD Cards, Pen drives whatever) and run the binary.
3. Multiple SD cards are supported at a time, but only 1 backup device. Make sure no other device other than the intended backup device contains the ".identifier" file.
4. The script runs forever, so if you add another SD card / Pen drive after you start the script, it'll automatically detect it and perform the backup.
5. **CURRENTLY DOES NOT SUPPORT COPYING FOLDERS**
