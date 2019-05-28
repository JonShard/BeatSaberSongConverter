## Beat Saber Song Converter

### Songe Converter by lolPants
**!! A more complete tool called *songe-converter* is being worked on by lolPants!!**  
https://github.com/lolPants/songe-converter/

When I searched this tool before I started making my own, I did not find anything. So I made my own very simple version so I could play my favorite songs.  Which is what this repository is.

### Beat Saber Song Converter
A simple tool that converts a Beat Saber song from the old format to the new format.
Check out Beat Saber here:  
http://beatsaber.com/

Beat Saber recently hopped out of Early Access (May 2019), with the update came a native song loader.  
This means that the game can now load songs without the need for mods. This is great news.  
The issue is that **the native songs have a different format** for their `Info` files.

This tool is a very simple(pretty stupid) and limited tool to convert a song from the old, modding community's standard to Beat Saber's new official standard.  
**It does not account for:**
- `noteJumpSpeed`: How fast the notes fly at you.
- Different `BPM` between difficulties: Beats per minute.
- Special characters in the info file.
Among other things I am too lazy to fix since a better version exist.

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
This will give you a converter.exe.

### Usage
1. Put `converter.exe` and a song's `info.json` in the same folder.
1. Run `converter.exe`. It then generates an `Info.dat`. Which is in the new format.

Or:
1. Put `converter.exe` and `convertAll.sh` in Beat Beat Saber's root folder.
  If installed with Steam on C drive: C:\Program Files (x86)\Steam\steamapps\common\Beat Saber
1. Run `convertAll.sh` to copy all songs from 'Beat Saber/CustomSongs ' into Beat Saber's new 'Beat Saber\Beat Saber_Data\CustomLevels' folder and convert them.
