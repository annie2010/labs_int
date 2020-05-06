<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Package Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Mission

> Package deal!

> Implement two separate modules aka dictionary and picker. The dictionary module produces a list of words from a given dictionary. The picker modules picks a random word leveraging the dictionary module.

* Clone the [Labs Repo](https://github.com/gopherland/labs_int)
* cd modules
* Dictionary Module (modules/dictionary)
  * Leveraging the given tests, implement a dictionary loader.
    * The loader loads a collection of words from file (see testdata)
    * The loader should take a dictionary location and a list of words to exclude from the returned list
* Picker Module (modules/picker)
  * The picker package loads words from the directory module and randomly pick a new word.
  * Initially make sure you can reference the dictionary from your own repo using the `replace` directive in go.mod
* Using the provided test suites. Test your entire application
* Make sure the picker cli works as expected.
* Using your own github repo user, version and publish your own dictionary module.
* BONUS!
  * Change your picker to reference a class mate dictionary module.

### Commands

* Publish your `dictionary` module (modules/dictionary):
  * On github, create your own git repo using your own GIT_USER_HANDLE

  ```shell
  cd modules/dictionary
  git init
  git add .
  git commit -m 'Init drop'
  git remote add origin git@github.com:YOUR_GIT_USER_HANDLE/dictionary.git
  git push -u origin master
  # Tag your repo with version 0.1.0
  git tag -a v0.1.0 -m 'Init drop'
  git push origin v0.1.0
  ```

* Remove the replace directive in picker/go.mod as we are now going to use your published
  dictionary module
* In picker/main.go update the import path to use your published dictionary module handle
* Run the picker app. You should see your dependency getting pulled!
* Verify your go.mod file is now correctly referencing your new dictionary module!
* BONUS: Rinse/Repeat the last 3 steps using a classmate dictionary module.

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)