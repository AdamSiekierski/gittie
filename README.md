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
