# DropBox File Sorter [![CircleCI](https://circleci.com/gh/jaredfholgate/DropBoxFileSorter.Go.svg?style=svg)](https://circleci.com/gh/jaredfholgate/DropBoxFileSorter.Go)

Moves files from the Camera Uploads folder to a defined hierarchy.



### Parameters

#### Arg 1: The source folder for Camera Uploads (e.g. c:\DropBox\Camera Uploads)

#### Arg 2: The target folder for the folder hierarchy (e.g. c:\DropBox)

The target folder hierarchy build folder in this format based on the file name and file extension;

* [TargetFolder]\Photos\YYYY MM MMMM (e.g. c:\DropBox\Photos\2018 06 June)

* [TargetFolder]\Video\YYYY MM MMMM 


### Usage Examples

CMD: DropBoxFileSorter.Go.exe "c:\DropBox\Camera Uploads" "c:\DropBox"

