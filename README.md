Automatic Gitignore From Github
---------

When I start a new project, first thing after running git init is I realize I need to add a gitignore file before adding the initial files to the project.  So, I Google for basic gitignore like python gitignore and copy paste the sample provided from github/gitignore.  Everytime I do this I think wow I need to add a simple app to just do this for me.  After the thousandth time I created this application.

This is a Go application that creates a default gitignore based on the type of project.  Templates for the default .gitignore come from https://github.com/github/gitignore which is open source and actively contributed to, so you will always get an up to date gitignore for your projects using this tool!  In fact, the gitignore in this very project was the first .gitignore created using this tool.

Creates a .gitignore in the current working directory with the defaults for the specified type of project.

````shell
gitignore python
gitignore go
````

But wait, there's more!  You can chain multiple gitignores together.
````shell
gitignore macos linux windows python node vim
````

Now your .gitignore contains all of the default .gitignore patterns from all of the github gitignores for MacOS, Linux, Windows, Python, Node, and Vim, so you're ready to commit w/e you need to on a cross-platform polyglot application.

## Long List of Available .gitignores
Any .gitignore available on https://github.com/github/gitignore should be available here.  If you find one that isn't, shoot a pull request or message to Glen Baker <iepathos@gmail.com>.

````shell
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
gitignore chefcookbook
gitignore clojure
gitignore codeigniter
gitignore commonlisp
gitignore composer
gitignore concrete5
gitignore coq
gitignore craftcms


gitignore d
gitignore dm
gitignore dart
gitignore delphi
gitignore drupal

gitignore episerver

gitignore eagle
gitignore elisp
gitignore elixir
gitignore elm
gitignore erlang
gitignore expressionengine
gitignore extjs

gitignore fancy
gitignore finale

gitignore forcedotcom
gitignore fortran
gitignore fuelphp

gitignore gwt
gitignore gcov
gitignore gitbook
gitignore go
gitignore gradle
gitignore grails

gitignore haskell

gitignore igorpro
gitignore idris

gitignore java
gitignore jboss
gitignore jekyll
gitignore joomla
gitignore julia

gitignore kicad
gitignore kohana

gitignore labview
gitignore laravel
gitignore leiningen
gitignore lemonstand
gitignore lilypond
gitignore lithium
gitignore lua

gitignore magento
gitignore maven
gitignore mercury
gitignore metaprogrammingsystem

gitignore nanoc
gitignore nim
gitignore node

gitignore ocaml
gitignore objective-c
gitignore opa
gitignore opencart
gitignore oracleforms

gitignore packer
gitignore perl
gitignore phalcon
gitignore playframework
gitignore plone
gitignore prestashop
gitignore processing
gitignore python

gitignore qooxdoo
gitignore qt

gitignore r
gitignore ros
gitignore rails
gitignore rhodesrhomobile
gitignore ruby
gitignore rust

gitignore scons
gitignore sass
gitignore scala
gitignore scheme
gitignore scrivener
gitignore sdcc
gitignore seamgen
gitignore sketchup
gitignore smalltalk
gitignore stella
gitignore sugarcrm
gitignore swift
gitignore symfony
gitignore symphonycms

gitignore text
gitignore terraform
gitignore textpattern
gitignore turbogears2
gitignore typo3

gitignore umbraco
gitignore unity
gitignore unrealengine

gitignore vvvv
gitignore visualstudio

gitignore waf
gitignore wordpress

gitignore xojo

gitignore yeoman
gitignore yii

gitignore zendframework
gitignore zephir


gitignore anjuta
gitignore ansible
gitignore archives

gitignore bazaar
gitignore bricxcc

gitignore cvs
gitignore calabash
gitignore cloud9
gitignore codekit

gitignore darteditor
gitignore dreamweaver
gitignore dropbox

gitignore eclipse
gitignore eiffelstudio
gitignore emacs
gitignore ensime
gitignore espresso

gitignore flexbuilder
gitignore gpg

gitignore jdeveloper
gitignore jetbrains

gitignore kdevelop4
gitignore kate

gitignore lazarus
gitignore libreoffice
gitignore linux
gitignore lyx

gitignore matlab
gitignore mercurial
gitignore microsoftoffice
gitignore modelsim
gitignore momentics
gitignore monodevelop

gitignore netbeans
gitignore ninja
gitignore notepadpp

gitignore otto

gitignore redcar
gitignore redis

gitignore sbt
gitignore svn
gitignore slickedit
gitignore sublimetext
gitignore synopsysvcs

gitignore tags
gitignore textmate
gitignore tortoisegit

gitignore vagrant
gitignore vim
gitignore virtualenv
gitignore visualstudiocode

gitignore webmethods
gitignore windows

gitignore xcode
gitignore xilinxise

gitignore macos
````





This is an open source project I've created to save us all from the hassle of copy-pasting default gitignore files for new projects.


