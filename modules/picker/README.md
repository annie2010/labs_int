<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

[![Build Status](https://travis-ci.org/derailed/picker.svg?branch=master)](https://travis-ci.org/derailed/picker)
[![Go Report Card](https://goreportcard.com/badge/github.com/derailed/picker)](https://goreportcard.com/report/github.com/derailed/picker)
[![GoDoc](https://godoc.org/github.com/derailed/picker?status.svg)](http://godoc.org/github.com/derailed/picker)

# Package Lab

---
## <img src="../../assets/lab.png" width="auto" height="32"/> Mission

> Package deal! Modulerize and publish a picker package to pick a word from a given word dictionary

* In this lab, you are going to leverage an existing implementation and deploy it as a go module.
* The picker package loads words from the given assets directory and randomly pick a new word.
* The application contains both lib code and a binary.
* You will need to use your own github account to publish your package

### Commands

* In your lab repo navigate to gopherland/labs/modules/picker
* Create a new module file using the following command:
  ```shell
  go mod init github.com/YOUR_GIT_USER_HANDLE/labs/picker
  ```
* Run the test command to make sure the tests are still passing and coverage is good!
* Run the main application and ensure it's working correctly
* Ensure all exported items have documentation and code coverage is adequate
* Using your own github account, create a new public repo named **picker**
* Setup your git repo using:
  ```shell
  git init
  git add .
  git commit -m 'Init drop'
  git remote add origin git@github.com:YOUR_GIT_USER_HANDLE/picker.git
  git push --set-upstream origin master
  ```
* In cmd/main.go update the import path to use your new repo
* In the root of your repo run your cli app
  ```shell
  go run cmd/main.go
  go run cmd/main.go -dic musicians -dir assets
  # Clean up your dependencies
  go mod tidy
  ```
* Using semantic versioning tag your release
  ```shell
  git tags v1.0.0
  git push --tags
  ```
* Checkout your docs and score card!
  ```shell
  open https://godoc.org/github.com/YOUR_GIT_USER_HANDLE/picker
  open https://goreportcard.com/report/github.com/YOUR_GIT_USER_HANDLE/picker
  ```
* Update the README.md badge section with your own github handle
* Congratulation! You've just built and published your very own go package!

* [BONUS]: Download and use, your classmate packages
* [BONUS] Use [Travis CI](https://travis-ci.org), add your repo and earn an extra badge!

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)