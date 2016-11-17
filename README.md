Automatic Gitignore
---------

When I start a new project, first thing after running git init is I realize I need to add a gitignore file before adding the initial files to the project.  So, I Google for basic gitignore like python gitignore and copy paste the sample provided from github/gitignore.  Everytime I do this I think wow I need to add a simple app to just do this for me.  After the thousandth time I created this application.

This is a Go application that creates a default gitignore based on the type of project.  Templates for the default .gitignore are copied from https://github.com/github/gitignore which is already open source and actively contributed to, so you will always get an up to date gitignore for your project using this tool!  In fact, the gitignore in this very project was the first .gitignore created using this tool.

Example:
````
gitignore python
````


````
gitignore go
````


````
gitignore node
````
