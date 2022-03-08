# secret-inject

Uses goreleaser for build, relase & homebrew management

Local build

```bash
goreleaser build --skip-validate --single-target --rm-dist --snapshot
find dist/plugins/ -type f | xargs -n1 -I {} cp {} dist/plugins
```

Pre publish build

```bash
goreleaser release --rm-dist --skip-validate --skip-publish
```

Release to GitHub/Homebrew
```bash
git commit
git tag X.X.X
git push
git push --tags
# git push --atomic origin <branch name> <tag>
goreleaser release --rm-dist
```
