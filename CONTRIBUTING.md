# Contributing Guidelines

## General

* Contributions of all kinds (issues, ideas, proposals), not just code, are highly appreciated.
* Pull requests are welcome with the understanding that major changes will be carefully evaluated and discussed, and may not always be accepted. Starting with a discussion is always best!
* All contributions including documentation, filenames and discussions should be written in English language.

### Issues
Our issue tracker can be used to report issues and propose changes to the current or next version of the data.world Go SDK.

Please follow these guidelines before opening an issue:

- Make sure your issue is not a duplicate.
- Make sure your issue is relevant to the specification.

## Contribute Code

### Review Relevant Docs

* [data.world API](https://apidocs.data.world/api)

### Fork the Project

Fork the project [on Github](https://github.com/datadotworld/dwapi-go) and check out your copy.

```sh
$ git clone https://github.com/[YOUR_GITHUB_NAME]/dwapi-go.git
$ cd dwapi-go
$ git remote add upstream https://github.com/datadotworld/dwapi-go.git
```

### Test

Run tests:

```sh
$ make test
```

### Create a Feature Branch

```sh
$ git checkout master
$ git pull upstream master
$ git checkout -b my-feature-branch
```

### Write Tests

Try to write a test that reproduces the problem you're trying to fix or describes a feature that you want to build.

We definitely appreciate pull requests that highlight or reproduce a problem, even without a fix.

### Write Code

Implement your feature or bug fix. Make sure that all tests pass without errors.

Also, to make sure that your code follows our coding style guide and best practices, run the commands;
```sh
$ go get -u github.com/alecthomas/gometalinter (only the first time to install some dependencies)
$ make fmtcheck
$ make lint
```
Make sure to fix any errors that appear, if any. `make fmt` will fix files that fail `make fmtcheck`.

### Write Documentation

Document any external behavior in the [README](README.md).

### Commit Changes

Make sure git knows your name and email address:

```sh
git config --global user.name "Your Name"
git config --global user.email "contributor@example.com"
```

Writing good commit logs is important. A commit log should describe what changed and why.
```sh
git add ...
git commit
```

### Push

```sh
git push origin my-feature-branch
```

### Make a Pull Request

Go to https://github.com/[YOUR_GITHUB_NAME]/dwapi-go.git and select your feature branch. Click the 'Pull Request' button and fill out the form. Pull requests are usually reviewed within a few days.

### Thank you!

Thank you in advance, for contributing to this project!
