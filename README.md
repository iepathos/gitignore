Automatic Gitignore From Github
---------

When I start a new project, first thing after running git init is I realize I need to add a gitignore file before adding the initial files to the project.  So, I Google for basic gitignore like python gitignore and copy paste the sample provided from github/gitignore.  Everytime I do this I think wow I need to add a simple app to just do this for me.  After the thousandth time I created this application.

This is a Go application that creates a default gitignore based on the type of project.  Templates for the default .gitignore come from https://github.com/github/gitignore which is open source and actively contributed to, so you will always get an up to date gitignore for your projects using this tool!  In fact, the gitignore in this very project was the first .gitignore created using this tool.

Creates a .gitignore in the current working directory with the defaults for the specified type of project.

## Available Gitignores
````
gitignore actionscript
gitignore ada
gitignore agda
gitignore android
gitignore appengine
gitignore appceleratortitanium
gitignore archlinuxpackages
gitignore autotools

gitignore c++
gitignore c
gitignore cfwheels
gitignore cmake
gitignore cuda
gitignore cakephp
gitignore clojure


gitignore python
gitignore go
gitignore node
gitignore sass
gitignore unity
gitignore unrealengine
````





This is an open source project I've created to save us all from the hassle of copy-pasting default gitignore files for new projects.


