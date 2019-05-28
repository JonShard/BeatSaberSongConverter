#!/usr/bin/env bash

folders=(./CustomSongs/*)
for folder in "${folders[@]}"; do                                     # Loop through songs in old CustomSong folder.
  if [ -f "$folder/info.json" ]; then                                 # if missing subfolder. Meaning it looks like: CustomSongs/XXXX-XX/info.json.
    cp -r "${folder}" "Beat Saber_Data/CustomLevels"
    filename=$(echo "$folder" | awk -F'/' '{print$3}')                # Get just the directory file name. Eg XXXX-XX.
    name=$(cat "$folder/info.json" | awk -F'"' '{print$4}')           # Get name of song by reading "songName" fields from info.json.
    echo
    echo Copying: "$folder"
    echo Missing subfolder, adding one for: $name
    mv "Beat Saber_Data/CustomLevels/$filename" "Beat Saber_Data/CustomLevels/$name"  # Rename XXXX-XX to actual song name.
    echo
  else
    subfolders=($folder/*);
    echo Copying: "${subfolders[0]}"
    cp -r "${subfolders[0]}" "Beat Saber_Data/CustomLevels"           # The song has a subfolder name. We can just copy it.
  fi
done

folders=(./Beat\ Saber_Data/CustomLevels/*);                         # Loop through the newly copied folders with songs in them.
for folder in "${folders[@]}"; do
  if [ -d "$folder" ]; then
    ./converter.exe "${folder}"                                # convert each song.
  fi
done
