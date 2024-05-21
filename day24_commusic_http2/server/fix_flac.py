from mutagen.flac import FLAC
m = FLAC("./songslist/rightforyou.flac")
m["title"] = "Right For You"
m["artist"] = "Shawn Mendes"
m["album"] = "Wonder: Unreleased"
m["year"] = "2020"
m["TRACKNUMBER"] = "1"

m.save()