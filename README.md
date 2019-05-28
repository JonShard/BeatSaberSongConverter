## Beat Saber Song Converter

### Songe Converter by lolPants
**!! A more complete tool called *songe-converter* is being worked on by lolPants!!**  
https://github.com/lolPants/songe-converter/

When I searched this tool before I started making my own, I did not find anything. So I made my own very simple version. Which is what this repository is.

### Beat Saber Song Converter
A simple tool that converts a Beat Saber song from the old format to the new format.
Check out Beat Saber here:  
http://beatsaber.com/

Beat Saber recently hopped out of Early Access (May 2019), with the update came a native song loader.  
This means that the game can now load songs without the need for mods. This is great news.  
The issue is that **the native songs have a different format** for their `Info` files.

This tool is a very simple and limited tool to convert a song from the old, modding community's standard to Beat Saber's new official standard.

### Install instrucitons:
Install Golang on your system.  
https://golang.org/
```
git clone git@github.com:JonShard/BeatSaberSongConverter.git
cd BeatSaberSongConverter
go get github.com/bmatcuk/doublestar
go get github.com/TomOnTime/utfutil/catutf
go build
```

### Usage
1. Put `converter.exe` and an a songs `info.json` in the same folder.
1. Run `converter.exe`. It then generates an `Info.dat`. Which is in the new format. 
