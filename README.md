# qiwiDL

Very very simple program to download [qiwi.gg](https.//qiwi.gg) folders.

## Usage

```bash
git clone https://github.com/Masterjoona/qiwiDL
cd qiwiDL && go build
./qiwiDL
```

## Flags
- `-folder` - Folder url to download
- `-concurrent=3` - how many files to download at the same time.
- `-directory=.` - Directory to download to. Default is current directory.
- `-host=https://qiwi.lol/` - ~~qiwi seems to change their file host domain every now and then.~~ seems not to be the case anymore
> [!NOTE]
> ~~As I write this README and checked the previous domain `spyderrock.com` (which [these](https://github.com/jufantozzi/Qiwi.gg-downloader/blob/2d086e28eaf7f3a6972340a0f1c78b03c2f751d2/downloader.js#L54) [projects](https://github.com/gookie-dev/qiwi.ddl/blob/7a62b121acfc3f60818dfd53beba14e0681edeac/main.py#L8) use) it also apparently works still. So I guess you can use that too.~~
> looks like they finally made up their minds and settled on a domain.

## Screenshots
![image](https://bin.masterjoona.dev/u/f2JYRG.png)
![image](https://bin.masterjoona.dev/u/UYah4b.png)

## Modules
- [survey](https://github.com/AlecAivazis/survey) (Deprecated but still works)
- [mpb](https://github.com/vbauerster/mpb)

## Notes
- beginner project
- I tried to (copied) to make it download a file in parts, but I didn't see much difference in speed.
- I also tried to make a total byte progress bar, but I couldn't figure it out and didn't want to spend more time on it. There are remnants of it in the code commented out.
