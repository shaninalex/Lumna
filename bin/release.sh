#!/usr/bin/env bash
# Create an annotated release tag vX.Y.Z on the current commit of main.
#
#   bin/release.sh [patch|minor|major|X.Y.Z]   (default: minor)
#
# The tag is created locally only. Push it yourself:
#   git push origin vX.Y.Z
#
# The version is never written into files: the Makefile reads it back with
# `git describe` and passes it to the binary through -ldflags.
set -euo pipefail

bump="${1:-minor}"
release_branch="main"

die() { echo "release: $*" >&2; exit 1; }

cd "$(git rev-parse --show-toplevel)"

# --- preconditions -----------------------------------------------------------
branch="$(git symbolic-ref --short HEAD 2>/dev/null || true)"
[[ "$branch" == "$release_branch" ]] || die "must be on '$release_branch' (current: '${branch:-detached HEAD}')"

[[ -z "$(git status --porcelain)" ]] || die "working tree is not clean"

git fetch --quiet --tags origin "$release_branch"
local_head="$(git rev-parse HEAD)"
remote_head="$(git rev-parse "origin/$release_branch")"
[[ "$local_head" == "$remote_head" ]] || die "'$release_branch' differs from origin/$release_branch; pull or push first"

existing="$(git tag --points-at HEAD --list 'v[0-9]*')"
[[ -z "$existing" ]] || die "HEAD is already tagged: $existing"

# --- current version ---------------------------------------------------------
# Highest semver tag in the repository, with or without the "v" prefix
# (tags before v0.42.0 were created without it).
current="$(git tag --list --sort=-v:refname | grep -E '^v?[0-9]+\.[0-9]+\.[0-9]+$' | head -n1 || true)"
current="${current#v}"
current="${current:-0.0.0}"
IFS=. read -r major minor patch <<<"$current"

# --- next version ------------------------------------------------------------
case "$bump" in
  major) next="$((major + 1)).0.0" ;;
  minor) next="${major}.$((minor + 1)).0" ;;
  patch) next="${major}.${minor}.$((patch + 1))" ;;
  [0-9]*.[0-9]*.[0-9]*)
    [[ "$bump" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || die "invalid version '$bump'"
    next="$bump"
    ;;
  *) die "usage: $0 [patch|minor|major|X.Y.Z]" ;;
esac

newest="$(printf '%s\n%s\n' "$current" "$next" | sort -V | tail -n1)"
[[ "$next" != "$current" && "$newest" == "$next" ]] || die "v$next is not newer than $current"

tag="v$next"
git rev-parse -q --verify "refs/tags/$tag" >/dev/null && die "tag $tag already exists"

# --- release notes -----------------------------------------------------------
# Commits since the previous release that is reachable from HEAD.
previous="$(git describe --tags --abbrev=0 --match 'v[0-9]*' 2>/dev/null \
  || git describe --tags --abbrev=0 2>/dev/null || true)"
range="${previous:+$previous..}HEAD"
notes="$(git log --no-merges --pretty='- %s' "$range")"

echo "Current version: $current"
echo "New tag:         $tag on $(git rev-parse --short HEAD)"
echo "Changes since ${previous:-the beginning}: $(printf '%s\n' "$notes" | grep -c . || true) commits"
read -r -p "Create tag $tag? [y/N] " answer
[[ "$answer" =~ ^[Yy]$ ]] || die "aborted"

git tag -a "$tag" -F - <<EOF
Release $tag

$notes
EOF

echo
echo "Created $tag. Check it with:  git show $tag"
echo "Publish it with:              git push origin $tag"
