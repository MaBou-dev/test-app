# test-app
Project to test an API backend in Golang and a frontend in Angular

## Git Workflow

```
master
|
| develop
| |
| | frontend
| | |
| |/
| |
| | backend
| | |
| |/
| |
|/
```

master
: stable version, increment version when merge on it, tag version on this branch, release

develop
: integration, tests, validation before release

frontend
: features and bugfix on the ui part, angular

backend
: features and bugfix on the backend / api part, golang

#TODO:
* vim modeline

## Sources

* https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
* https://tutorialedge.net/golang/creating-restful-api-with-golang/

* https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/
