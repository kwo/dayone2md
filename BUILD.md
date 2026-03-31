# Build and Release Maintenance

This document covers the tagged release workflow and the manual Homebrew tap update process for `dayone2md`.

## Release workflow

Tagged releases are created from GitHub Actions.

Example:

```bash
git tag v1.2.3
git push origin v1.2.3
```

The release workflow publishes archives with this naming pattern:

- `dayone2md_v<version>_darwin_amd64.tar.gz`
- `dayone2md_v<version>_darwin_arm64.tar.gz`
- `SHA256SUMS`

Archive layout is binary-only:

- macOS archives extract a single `dayone2md` binary at the archive root

## Generate the Homebrew formula

From this repository, generate a ready-to-commit `dayone2md.rb` formula on stdout.

Latest stable release:

```bash
scripts/homebrew-formula
```

Specific tag:

```bash
scripts/homebrew-formula --tag v1.2.3
```

The script:

- uses `gh release list` to select the latest published non-draft, non-prerelease release by default
- uses `gh release view` to inspect release assets
- downloads and parses the release `SHA256SUMS`
- fails if `gh` is missing, unusable, or the release is missing required assets

## Update the tap

Tap repository:

- `kwo/homebrew-tools`

Tap command for users:

```bash
brew tap kwo/tools
```

Manual update flow:

1. Cut and publish the `dayone2md` GitHub release.
2. Generate the formula from this repo:
   ```bash
   scripts/homebrew-formula --tag v1.2.3 > ../homebrew-tools/dayone2md.rb
   ```
3. Review the generated formula in `../homebrew-tools/dayone2md.rb`.
4. Commit and push the tap update from `kwo/homebrew-tools`.

User install command after the tap contains the formula:

```bash
brew install dayone2md
```
