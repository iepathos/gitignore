package main

// gitignore is a simple .gitignore setup tool
// .gitignores are the latest from https://github.com/github/gitignore
// Written by Glen Baker <iepathos@gmail.com>
// This software is free to use and expand upon under
// a GNU General Public License v3.0

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func updateGitignore(body []byte) {
	// write to file
	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(body); err != nil {
		panic(err)
	}
}

func copyGitignoreUrl(url string, ch chan<- []byte) {
	// retrieve gitignore from given url
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ch <- body
}

func main() {
	if (len(os.Args)) >= 2 {
		// []byte to store multiple .gitignore request bodies
		ch := make(chan []byte)
		unknown_envs := 0
		fmt.Println("Retrieving the requested .gitignore patterns from github.com/github/gitignore")
		for i := 1; i < len(os.Args); i++ {
			// for each arg add retrieve the gitignore
			// this way we can chain multiple gitignore together
			// like: gitignore macos python node
			switch env := strings.ToLower(os.Args[i]); env {
			case "actionscript":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Actionscript.gitignore", ch)
			case "ada":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Ada.gitignore", ch)
			case "agda":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Agda.gitignore", ch)
			case "android":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Android.gitignore", ch)
			case "appengine":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/AppEngine.gitignore", ch)
			case "appceleratortitanium":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/AppceleratorTitanium.gitignore", ch)
			case "archlinuxpackages":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ArchLinuxPackages.gitignore", ch)
			case "autotools":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Autotools.gitignore", ch)

			case "c++":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/C%2B%2B.gitignore", ch)
			case "c":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/C.gitignore", ch)
			case "cfwheels":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CFWheels.gitignore", ch)
			case "cmake":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CMake.gitignore", ch)
			case "cuda":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CUDA.gitignore", ch)
			case "cakephp":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CakePHP.gitignore", ch)
			case "chefcookbook":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ChefCookbook.gitignore", ch)
			case "clojure":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Clojure.gitignore", ch)
			case "codeigniter":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CodeIgniter.gitignore", ch)
			case "commonlisp":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CommonLisp.gitignore", ch)
			case "composer":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Composer.gitignore", ch)
			case "concrete5":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Concrete5.gitignore", ch)
			case "coq":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Coq.gitignore", ch)
			case "craftcms":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CraftCMS.gitignore", ch)

			case "d":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/D.gitignore", ch)
			case "dm":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/DM.gitignore", ch)
			case "dart":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Dart.gitignore", ch)
			case "delphi":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Delphi.gitignore", ch)
			case "drupal":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Drupal.gitignore", ch)
			case "episerver":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/EPiServer.gitignore", ch)
			case "eagle":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Eagle.gitignore", ch)
			case "elisp":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Elisp.gitignore", ch)
			case "elixir":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Elixir.gitignore", ch)
			case "elm":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Elm.gitignore", ch)
			case "erlang":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Erlang.gitignore", ch)
			case "expressionengine":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ExpressionEngine.gitignore", ch)
			case "extjs":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ExtJs.gitignore", ch)
			case "fancy":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Fancy.gitignore", ch)
			case "finale":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Finale.gitignore", ch)
			case "forcedotcom":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ForceDotCom.gitignore", ch)
			case "fortran":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Fortran.gitignore", ch)
			case "fuelphp":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/FuelPHP.gitignore", ch)

			case "gwt":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/GWT.gitignore", ch)
			case "gcov":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Gcov.gitignore", ch)
			case "gitbook":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/GitBook.gitignore", ch)
			case "go":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore", ch)
			case "gradle":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Gradle.gitignore", ch)
			case "grails":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Grails.gitignore", ch)

			case "haskell":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Haskell.gitignore", ch)

			case "igorpro":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/IGORPro.gitignore", ch)
			case "idris":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Idris.gitignore", ch)

			case "java":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Java.gitignore", ch)
			case "jboss":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Jboss.gitignore", ch)
			case "jekyll":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Jekyll.gitignore", ch)
			case "joomla":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Joomla.gitignore", ch)
			case "julia":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Julia.gitignore", ch)

			case "kicad":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/KiCad.gitignore", ch)
			case "kohana":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/kohana.gitignore", ch)

			case "labview":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/LabVIEW.gitignore", ch)
			case "laravel":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Laravel.gitignore", ch)
			case "leiningen":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Leiningen.gitignore", ch)
			case "lemonstand":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/LemonStand.gitignore", ch)
			case "lilypond":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Lilypond.gitignore", ch)
			case "lithium":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Lithium.gitignore", ch)
			case "lua":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Lua.gitignore", ch)

			case "magento":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Magento.gitignore", ch)
			case "maven":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Maven.gitignore", ch)
			case "mercury":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Mercury.gitignore", ch)
			case "metaprogrammingsystem":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/MetaProgrammingSystem.gitignore", ch)

			case "nanoc":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Nanoc.gitignore", ch)
			case "nim":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Nim.gitignore", ch)
			case "node":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Node.gitignore", ch)

			case "ocaml":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/OCaml.gitignore", ch)
			case "objective-c":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Objective-C.gitignore", ch)
			case "opa":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Opa.gitignore", ch)
			case "opencart":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/OpenCart.gitignore", ch)
			case "oracleforms":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/OracleForms.gitignore", ch)

			case "packer":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Packer.gitignore", ch)
			case "perl":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Perl.gitignore", ch)
			case "phalcon":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Phalcon.gitignore", ch)
			case "playframework":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/PlayFramework.gitignore", ch)
			case "plone":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Plone.gitignore", ch)
			case "prestashop":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Prestashop.gitignore", ch)
			case "processing":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Processing.gitignore", ch)
			case "python":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Python.gitignore", ch)

			case "qooxdoo":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Qooxdoo.gitignore", ch)
			case "qt":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Qt.gitignore", ch)

			case "r":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/R.gitignore", ch)
			case "ros":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ROS.gitignore", ch)
			case "rails":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Rails.gitignore", ch)
			case "rhodesrhomobile":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/RhodesRhomobile.gitignore", ch)
			case "ruby":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Ruby.gitignore", ch)
			case "rust":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Rust.gitignore", ch)

			case "scons":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SCons.gitignore", ch)
			case "sass":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Sass.gitignore", ch)
			case "scala":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Scala.gitignore", ch)
			case "scheme":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Scheme.gitignore", ch)
			case "scrivener":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Scrivener.gitignore", ch)
			case "sdcc":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Sdcc.gitignore", ch)
			case "seamgen":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SeamGen.gitignore", ch)
			case "sketchup":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SketchUp.gitignore", ch)
			case "smalltalk":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Smalltalk.gitignore", ch)
			case "stella":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Stella.gitignore", ch)
			case "sugarcrm":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SugarCRM.gitignore", ch)
			case "swift":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Swift.gitignore", ch)
			case "symfony":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Symfony.gitignore", ch)
			case "symphonycms":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SymphonyCMS.gitignore", ch)

			case "tex":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/TeX.gitignore", ch)
			case "terraform":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Terraform.gitignore", ch)
			case "textpattern":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Textpattern.gitignore", ch)
			case "turbogears2":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/TurboGears2.gitignore", ch)
			case "typo3":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Typo3.gitignore", ch)

			case "umbraco":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Umbraco.gitignore", ch)
			case "unity":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Unity.gitignore", ch)
			case "unrealengine":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/UnrealEngine.gitignore", ch)

			case "vvvv":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/VVVV.gitignore", ch)
			case "visualstudio":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/VisualStudio.gitignore", ch)

			case "waf":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Waf.gitignore", ch)
			case "wordpress":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/WordPress.gitignore", ch)

			case "xojo":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Xojo.gitignore", ch)

			case "yeoman":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Yeoman.gitignore", ch)
			case "yii":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Yii.gitignore", ch)

			case "zendframework":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ZendFramework.gitignore", ch)
			case "zephir":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Zephir.gitignore", ch)

			// global
			case "anjuta":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Anjuta.gitignore", ch)
			case "ansible":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Ansible.gitignore", ch)
			case "archives":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Archives.gitignore", ch)

			case "bazaar":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Bazaar.gitignore", ch)
			case "bricxcc":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/BricxCC.gitignore", ch)

			case "cvs":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/CVS.gitignore", ch)
			case "calabash":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Calabash.gitignore", ch)
			case "cloud9":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Cloud9.gitignore", ch)
			case "codekit":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/CodeKit.gitignore", ch)

			case "darteditor":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/DartEditor.gitignore", ch)
			case "dreamweaver":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Dreamweaver.gitignore", ch)
			case "dropbox":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Dropbox.gitignore", ch)

			case "eclipse":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Eclipse.gitignore", ch)
			case "eiffelstudio":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/EiffelStudio.gitignore", ch)
			case "emacs":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Emacs.gitignore", ch)
			case "ensime":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Ensime.gitignore", ch)
			case "espresso":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Espresso.gitignore", ch)

			case "flexbuilder":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/FlexBuilder.gitignore", ch)

			case "gpg":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/GPG.gitignore", ch)

			case "jdeveloper":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/JDeveloper.gitignore", ch)
			case "jetbrains":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/JetBrains.gitignore", ch)

			case "kdevelop4":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/KDevelop4.gitignore", ch)
			case "kate":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Kate.gitignore", ch)

			case "lazarus":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Lazarus.gitignore", ch)
			case "libreoffice":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/LibreOffice.gitignore", ch)
			case "linux":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Linux.gitignore", ch)
			case "lyx":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/LyX.gitignore", ch)

			case "matlab":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Matlab.gitignore", ch)
			case "mercurial":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Mercurial.gitignore", ch)
			case "microsoftoffice":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/MicrosoftOffice.gitignore", ch)
			case "modelsim":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/ModelSim.gitignore", ch)
			case "momentics":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Momentics.gitignore", ch)
			case "monodevelop":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/MonoDevelop.gitignore", ch)

			case "netbeans":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/NetBeans.gitignore", ch)
			case "ninja":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Ninja.gitignore", ch)
			case "notepadpp":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/NotepadPP.gitignore", ch)

			case "otto":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Otto.gitignore", ch)

			case "redcar":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Redcar.gitignore", ch)
			case "redis":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Redis.gitignore", ch)

			case "sbt":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SBT.gitignore", ch)
			case "svn":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SVN.gitignore", ch)
			case "slickedit":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SlickEdit.gitignore", ch)
			case "sublimetext":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SublimeText.gitignore", ch)
			case "synopsysvcs":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SynopsysVCS.gitignore", ch)

			case "tags":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Tags.gitignore", ch)
			case "textmate":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/TextMate.gitignore", ch)
			case "tortoisegit":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/TortoiseGit.gitignore", ch)

			case "vagrant":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Vagrant.gitignore", ch)
			case "vim":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Vim.gitignore", ch)
			case "virtualenv":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/VirtualEnv.gitignore", ch)
			case "visualstudiocode":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/VisualStudioCode.gitignore", ch)

			case "webmethods":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/WebMethods.gitignore", ch)

			case "windows":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Windows.gitignore", ch)

			case "xcode":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Xcode.gitignore", ch)
			case "xilinxise":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/XilinxISE.gitignore", ch)

			case "macos":
				go copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/macOS.gitignore", ch)

			default:
				fmt.Println("Given argument was not a recognized .gitignore default")
				unknown_envs += 1
			}
		}
		// write []byte in channel ch to file
		for i := 0; i < len(os.Args[1:])-unknown_envs; i++ {
			updateGitignore(<-ch)
		}
		fmt.Println("Updated .gitignore with the requested patterns")
	} else {
		fmt.Println("Please pass type of gitignore to create, like: gitignore python")
	}

}
