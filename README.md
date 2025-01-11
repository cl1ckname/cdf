# CDF - 📁 OS path bookmarks manager

**CDF** is a fast and efficient terminal-based CLI tool, written in Go, designed to perform a single task with precision and simplicity. Following the Unix philosophy of "do one thing, but do it well," cdf has zero dependencies and provides a minimalistic solution to a specific problem. Its lightweight nature ensures that it integrates seamlessly into any Unix-like environment, delivering a reliable and focused user experience.

**Warning! CDF is under heavy development! Use for your own risk!**

## Usage
The CDF should be treated as a list of bookmarks in the browser, not as a search bar, unlike similar projects the CDF does not collect any data on its own, so the user himself must add points for quick navigation. 
### Add mark
The first step in using the program is to add a mark. 

```console
cdf add home /home/username
```

Both absolute and relative paths can be used for this.

```console
cdf add projects . # in /home/username/projects
```

### List marks

The list of brands can be easily retrieved by the `list` command

```console
cdf list
```

### Jump to mark
To move to the defined mark, use a shortened version of the command with the path alias

```
f home # we're in /home/username now
```

## Why not just use Z or Jump?

The short answer - it is different. 

The long answer is - well, that's really something else entirely. CDF uses only precisely marked points. The user does not need to heuristically evaluate how this or that incomplete path will resolve for him. Fuzzy Search is wonderful and is definitely comfortable to use. But there is such a way to guarantee yourself a quick result, even if only for a couple of symbols. Use them both to achieve unprecedented productivity!

## Inspired by
This project is inspired by these. Try them too!

- [zoxide](https://github.com/ajeetdsouza/zoxide)
- [yazi](https://github.com/sxyazi/yazi)
- [jump](https://github.com/gsamokovarov/jump)
- [z](https://github.com/rupa/z)