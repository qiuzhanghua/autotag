# AutoTag
## Install
```bash
go install github.com/qiuzhanghua/autotag@latest
```

## SemVer
You should use Semantic Versioning in your projects

see https://semver.org/

Suppose you are using git, please use semver to tag, for example:
```bash
git tag 0.0.1
git tag 0.1.0-alpha.2
git tag 0.1.0-beta.0
git tag 0.1.0-rc.1

# git tag v0.9.0 
# also supported
```

## Rule of calculating next version

### Next Prerelease
```text
a.b.c - pre -> a.b.c
a.b.c-alpha.3 - pre -> a.b.c-alpha.4
a.b.c-beta.2 - pre -> a.b.c-beta.3
a.b.c-rc.1 - pre -> a.b.c-rc.2
```

### Next Phase
```text
a.b.c - phase -> a.b.(c+1)-alpha.0
a.b.c-alpha.2 - phase -> a.b.c-beta.0
a.b.c-beta.3 - phase -> a.b.c-rc.0
a.b.c-rc.4 - phase -> a.b.c
```

### Next Patch
```text
a.b.c - patch -> a.b.(c+1)
a.b.c-alpha.2 - phase -> a.b.(c+1)
a.b.c-beta.3 - phase -> a.b.(c+1)
a.b.c-rc.5 - phase -> a.b.(c+1)
```

### Next Minor
```text
a.b.c - minor -> a.(b+1).0
a.b.c-alpha.2 - minor -> a.(b+1).0
a.b.c-beta.3 - minor -> a.(b+1).0
a.b.c-rc.5 - minor -> a.(b+1).0
```

### Next Major
```text
a.b.c - major -> (a+1).0.0
a.b.c-alpha.2 - major -> (a+1).0.0
a.b.c-beta.3 - major -> (a+1).0.0
a.b.c-rc.5 - major -> (a+1).0.0
```

### Undefined version
suppose it is 0.0.0

## Command

### help
show autotag's help.
```bash
autotag help
```

### version
show autotag's version.
```bash
autotag version
```

### show
show autotag's version and next pre/phase/patch/minor/major version.
```bash
autotag show
```
result as:
```text
Current tag:  0.0.1-alpha.0
next pre   :  0.0.1-alpha.1
next phase :  0.0.1-beta.0
next patch :  0.0.1
next minor :  0.1.0
next major :  1.0.0
```

### next
add git tag using next pre/phase/patch/minor/major version as name
```bash
autotag next pre
autotag next phase
autotag next patch
autotag next minor
autotag next major
```

### write

**AppBuildDate is the date of the last commit**, AppRevision is the short commit hash, AppVersion is the latest tag.

```bash
autotag write
```
write autotag.go for Golang project managed by git, as
```go
package main

const AppVersion = "v0.3.3"
const AppRevision = "169b357"
const AppBuildDate = "2021-12-24"

```

write autotag.js for Node.js project managed by git, as
```javascript
const AppVersion = "v0.3.3"
const AppRevision = "169b357"
const AppBuildDate = "2021-12-24"

```

Learn a lot from https://github.com/hkloudou/git-autotag
