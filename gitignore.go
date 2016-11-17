package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)


func copyGitignoreUrl(url string) {
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

	// write to file
	f, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
	    panic(err)
	}

	defer f.Close()

	if _, err = f.Write(body); err != nil {
	    panic(err)
	}

	fmt.Println("Updated .gitignore with", url)
}

func main() {
	if (len(os.Args)) >= 2 {
		for i := 1; i < len(os.Args); i++ {
			// for each arg add its env to the gitignore
			// this way we can chain multiple gitignore together
			// like: gitignore macos python node
			switch env := strings.ToLower(os.Args[i]); env {
				case "actionscript":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Actionscript.gitignore")
				case "ada":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Ada.gitignore")
				case "agda":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Agda.gitignore")
				case "android":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Android.gitignore")
				case "appengine":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/AppEngine.gitignore")
				case "appceleratortitanium":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/AppceleratorTitanium.gitignore")
				case "archlinuxpackages":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ArchLinuxPackages.gitignore")
				case "autotools":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Autotools.gitignore")
				
				case "c++":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/C%2B%2B.gitignore")
				case "c":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/C.gitignore")
				case "cfwheels":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CFWheels.gitignore")
				case "cmake":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CMake.gitignore")
				case "cuda":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CUDA.gitignore")
				case "cakephp":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CakePHP.gitignore")
				case "chefcookbook":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ChefCookbook.gitignore")
				case "clojure":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Clojure.gitignore")
				case "codeigniter":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CodeIgniter.gitignore")
				case "commonlisp":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CommonLisp.gitignore")
				case "composer":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Composer.gitignore")
				case "concrete5":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Concrete5.gitignore")
				case "coq":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Coq.gitignore")
				case "craftcms":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/CraftCMS.gitignore")
				
				case "d":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/D.gitignore")
				case "dm":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/DM.gitignore")
				case "dart":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Dart.gitignore")
				case "delphi":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Delphi.gitignore")
				case "drupal":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Drupal.gitignore")
				case "episerver":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/EPiServer.gitignore")
				case "eagle":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Eagle.gitignore")
				case "elisp":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Elisp.gitignore")
				case "elixir":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Elixir.gitignore")
				case "elm":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Elm.gitignore")
				case "erlang":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Erlang.gitignore")
				case "expressionengine":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ExpressionEngine.gitignore")
				case "extjs":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ExtJs.gitignore")
				case "fancy":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Fancy.gitignore")
				case "finale":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Finale.gitignore")
				case "forcedotcom":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ForceDotCom.gitignore")
				case "fortran":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Fortran.gitignore")
				case "fuelphp":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/FuelPHP.gitignore")

				case "gwt":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/GWT.gitignore")
				case "gcov":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Gcov.gitignore")
				case "gitbook":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/GitBook.gitignore")
				case "go":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore")
				case "gradle":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Gradle.gitignore")
				case "grails":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Grails.gitignore")

				case "haskell":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Haskell.gitignore")

				case "igorpro":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/IGORPro.gitignore")
				case "idris":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Idris.gitignore")

				case "java":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Java.gitignore")
				case "jboss":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Jboss.gitignore")
				case "jekyll":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Jekyll.gitignore")
				case "joomla":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Joomla.gitignore")
				case "julia":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Julia.gitignore")

				case "kicad":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/KiCad.gitignore")
				case "kohana":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/kohana.gitignore")

				case "labview":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/LabVIEW.gitignore")
				case "laravel":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Laravel.gitignore")
				case "leiningen":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Leiningen.gitignore")
				case "lemonstand":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/LemonStand.gitignore")
				case "lilypond":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Lilypond.gitignore")
				case "lithium":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Lithium.gitignore")
				case "lua":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Lua.gitignore")
				
				case "magento":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Magento.gitignore")
				case "maven":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Maven.gitignore")
				case "mercury":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Mercury.gitignore")
				case "metaprogrammingsystem":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/MetaProgrammingSystem.gitignore")

				case "nanoc":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Nanoc.gitignore")
				case "nim":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Nim.gitignore")
				case "node":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Node.gitignore")

				case "ocaml":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/OCaml.gitignore")
				case "objective-c":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Objective-C.gitignore")
				case "opa":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Opa.gitignore")
				case "OpenCart":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/OpenCart.gitignore")
				case "oracleforms":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/OracleForms.gitignore")
				
				case "packer":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Packer.gitignore")
				case "perl":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Perl.gitignore")
				case "phalcon":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Phalcon.gitignore")
				case "playframework":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/PlayFramework.gitignore")
				case "plone":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Plone.gitignore")
				case "prestashop":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Prestashop.gitignore")
				case "processing":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Processing.gitignore")
				case "python":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Python.gitignore")

				case "qooxdoo":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Qooxdoo.gitignore")
				case "qt":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Qt.gitignore")

				case "r":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/R.gitignore")
				case "ros":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ROS.gitignore")
				case "rails":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Rails.gitignore")
				case "rhodesrhomobile":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/RhodesRhomobile.gitignore")
				case "ruby":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Ruby.gitignore")
				case "rust":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Rust.gitignore")
				
				case "scons":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SCons.gitignore")
				case "sass":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Sass.gitignore")
				case "scala":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Scala.gitignore")
				case "scheme":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Scheme.gitignore")
				case "scrivener":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Scrivener.gitignore")
				case "sdcc":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Sdcc.gitignore")
				case "seamgen":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SeamGen.gitignore")
				case "smalltalk":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Smalltalk.gitignore")
				case "stella":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Stella.gitignore")
				case "sugarcrm":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SugarCRM.gitignore")
				case "swift":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Swift.gitignore")
				case "symfony":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Symfony.gitignore")
				case "symphonycms":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/SymphonyCMS.gitignore")

				case "tex":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/TeX.gitignore")
				case "terraform":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Terraform.gitignore")
				case "textpattern":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Textpattern.gitignore")
				case "turbogears2":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/TurboGears2.gitignore")
				case "typo3":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Typo3.gitignore")

				case "umbraco":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Umbraco.gitignore")
				case "unity":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Unity.gitignore")
				case "unrealengine":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/UnrealEngine.gitignore")
				
				case "vvvv":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/VVVV.gitignore")
				case "visualstudio":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/VisualStudio.gitignore")

				case "waf":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Waf.gitignore")
				case "wordpress":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/WordPress.gitignore")

				case "xojo":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Xojo.gitignore")

				case "yeoman":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Yeoman.gitignore")
				case "yii":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Yii.gitignore")

				case "zendframework":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/ZendFramework.gitignore")
				case "zephir":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Zephir.gitignore")

				// global
				case "anjuta":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Anjuta.gitignore")
				case "ansible":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Ansible.gitignore")
				case "archives":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Archives.gitignore")

				case "bazaar":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Bazaar.gitignore")
				case "bricxcc":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/BricxCC.gitignore")

				case "cvs":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/CVS.gitignore")
				case "calabash":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Calabash.gitignore")
				case "cloud9":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Cloud9.gitignore")
				case "codekit":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/CodeKit.gitignore")

				case "darteditor":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/DartEditor.gitignore")
				case "dreamweaver":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Dreamweaver.gitignore")
				case "dropbox":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Dropbox.gitignore")

				case "eclipse":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Eclipse.gitignore")
				case "eiffelstudio":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/EiffelStudio.gitignore")
				case "emacs":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Emacs.gitignore")
				case "ensime":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Ensime.gitignore")
				case "espresso":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Espresso.gitignore")

				case "flexbuilder":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/FlexBuilder.gitignore")

				case "gpg":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/GPG.gitignore")

				case "jdeveloper":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/JDeveloper.gitignore")
				case "jetbrains":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/JetBrains.gitignore")

				case "kdevelop4":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/KDevelop4.gitignore")
				case "kate":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Kate.gitignore")

				case "lazarus":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Lazarus.gitignore")
				case "libreoffice":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/LibreOffice.gitignore")
				case "linux":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Linux.gitignore")
				case "lyx":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/LyX.gitignore")

				case "matlab":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Matlab.gitignore")
				case "mercurial":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Mercurial.gitignore")
				case "microsoftoffice":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/MicrosoftOffice.gitignore")
				case "modelsim":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/ModelSim.gitignore")
				case "momentics":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Momentics.gitignore")
				case "monodevelop":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/MonoDevelop.gitignore")

				case "netbeans":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/NetBeans.gitignore")
				case "ninja":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Ninja.gitignore")
				case "notepadpp":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/NotepadPP.gitignore")

				case "otto":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Otto.gitignore")

				case "redcar":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Redcar.gitignore")
				case "redis":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Redis.gitignore")

				case "sbt":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SBT.gitignore")
				case "svn":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SVN.gitignore")
				case "slickedit":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SlickEdit.gitignore")
				case "sublimetext":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SublimeText.gitignore")
				case "synopsysvcs":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/SynopsysVCS.gitignore")

				case "tags":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Tags.gitignore")
				case "textmate":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/TextMate.gitignore")
				case "tortoisegit":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/TortoiseGit.gitignore")

				case "vagrant":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Vagrant.gitignore")
				case "vim":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Vim.gitignore")
				case "virtualenv":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/VirtualEnv.gitignore")
				case "visualstudiocode":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/VisualStudioCode.gitignore")

				case "webmethods":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/WebMethods.gitignore")

				case "windows":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Windows.gitignore")

				case "xcode":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/Xcode.gitignore")
				case "xilinxise":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/XilinxISE.gitignore")

				case "macos":
					copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Global/macOS.gitignore")

				default:
					fmt.Println("Please pass type of gitignore to create, like: gitignore python")
		
			}
		}
	} else {
		fmt.Println("Please pass type of gitignore to create, like: gitignore python")
	}
	
}
