## go build with a target file omits vcs info

```
vcspoc[master]$ make
rm -f vcspoc
go1.18 build vcspoc.go
./vcspoc
2022/02/23 00:13:35  git_commit=unknown; commit_date=unknown; go_version=go1.18rc1
go1.18 build
./vcspoc
2022/02/23 00:13:35  git_commit=b83e649709e4bdf1388f1b7702a97c82171d98e0; commit_date=2022-02-23T00:12:02Z; go_version=go1.18rc1
```
