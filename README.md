# Changelogger

Log your local file changes

## Usage

```bash
$ go get github.com/k0kubun/changelogger
$ cd PATH_TO_LOG
$ changelogger
```

## Example

```diff
 $ cd ~/go/src/changelogger
 $ changelogger
 2014/04/29 09:28:23 Changed: /Users/k0kubun/go/src/changelogger/diff.go
- lines := strings.Split(text, "\n")
-       for _, line := range lines {

 2014/04/29 09:28:35 Changed: /Users/k0kubun/go/src/changelogger/diff.go
+       dmp := diffmatchpatch.New()
+       diffs := dmp.DiffMain(oldText, newText, true)
```
