build:
	go build -ldflags "-X main.AppVersion=`autotag current` -X main.AppRevision=`autotag hash` -X main.AppBuildDate=`autotag date`"
install:
	go install -ldflags "-X main.AppVersion=`autotag current` -X main.AppRevision=`autotag hash` -X main.AppBuildDate=`autotag date`"
