# Release Flow

Merging to `main` triggers the stable release workflow.

The workflow:

1. Finds the latest `v*-stable` tag.
2. Increments the patch version.
3. Builds Linux `amd64` and `arm64` binaries.
4. Injects the release tag into `terror version`.
5. Generates release notes from commit messages.
6. Publishes the GitHub release as latest.

## Installer Assets

Installer assets are deployed separately. After a successful stable release, `.github/workflows/deploy-installers.yml` copies changed installer assets to:

```txt
/var/www/terrorserver
```

The hosted installer is then used by both fresh installs and `terror update`.
