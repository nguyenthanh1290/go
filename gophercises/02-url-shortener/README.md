Exercise details can be found here: https://gophercises.com/exercises/urlshort

However, I recommend to check this Go doc first, https://golang.org/doc/articles/wiki/. By following the doc to build a Web App, you will have enough knowledge about http package to complete this exercise.

## Skills developed

* net/http package
* yaml

### Issue
"panic: yaml: line 2: found character that cannot start any token"

Fix:
- make sure there is no tab character in your yaml text
- that includes, do not indent the lines when declaring the variable storing the yaml text
