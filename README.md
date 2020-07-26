<p align="center">
  <img src="https://git-scm.com/images/logos/downloads/Git-Icon-1788C.png" width="200" />
</p>
<h2 align="center">gittie</h2>
<p align="center"><i>Quickly create .gitignore in your project using one simple command</i></p>

### About

I've always had a problem with creating a `.gitignore` file for my projects - it took too long, to enter the gitignore repo on GitHub, and find the appropiate `.gitignore` for my project. So I created this - **Gittie**

Gittie uses GitHub's REST API to get ready-to-use `.gitignore` templates, which then you can use inside your repo.

### Installation

Coming soon

### Usage

#### To retrieve a list of available templates:

```
gittie --list
```

or

```
gittie -l
```

#### To generate a .gitignore from a template

```
gittie --template [template]
```

For example, to create a .gitignore with Go template:

```
gittie -t Go
```

### Credits

- [Me](https://siekierski.ml) - Doing 2% of the work
- [GitHub](https://github.com/github/gitignore) - Preparing the .gitignore templates

### Templates

All available templates are:

- Actionscript
- Ada
- Agda
- Android
- AppEngine
- AppceleratorTitanium
- ArchLinuxPackages
- Autotools
- C
- C++
- CFWheels
- CMake
- CUDA
- CakePHP
- ChefCookbook
- Clojure
- CodeIgniter
- CommonLisp
- Composer
- Concrete5
- Coq
- CraftCMS
- D
- DM
- Dart
- Delphi
- Drupal
- EPiServer
- Eagle
- Elisp
- Elixir
- Elm
- Erlang
- ExpressionEngine
- ExtJs
- Fancy
- Finale
- ForceDotCom
- Fortran
- FuelPHP
- GWT
- GitBook
- Go
- Godot
- Gradle
- Grails
- Haskell
- IGORPro
- Idris
- JENKINS_HOME
- Java
- Jboss
- Jekyll
- Joomla
- Julia
- KiCAD
- Kohana
- Kotlin
- LabVIEW
- Laravel
- Leiningen
- LemonStand
- Lilypond
- Lithium
- Lua
- Magento
- Maven
- Mercury
- MetaProgrammingSystem
- Nim
- Node
- OCaml
- Objective-C
- Opa
- OracleForms
- Packer
- Perl
- Perl6
- Phalcon
- PlayFramework
- Plone
- Prestashop
- Processing
- PureScript
- Python
- Qooxdoo
- Qt
- R
- ROS
- Rails
- RhodesRhomobile
- Ruby
- Rust
- SCons
- Sass
- Scala
- Scheme
- Scrivener
- Sdcc
- SeamGen
- SketchUp
- Smalltalk
- SugarCRM
- Swift
- Symfony
- SymphonyCMS
- TeX
- Terraform
- Textpattern
- TurboGears2
- Typo3
- Umbraco
- Unity
- UnrealEngine
- VVVV
- VisualStudio
- Waf
- WordPress
- Xojo
- Yeoman
- Yii
- ZendFramework
- Zephir
- gcov
- nanoc
- opencart
- stella

_Thanks for reading the `readme.md` till the end! For more stuff from me, check out my [website](https://siekierski.ml) and [podcast](https://require.podcast.gq)_
