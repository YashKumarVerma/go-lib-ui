# go-lib-ui
This repository is supposed to be consumed as a sub-module in other cli projects. The main motive behind forming sub-modules for my commonly used libraries / packages like ui, config etc is to align with DRY principles and have a central library which provides apis to all the projects.

## Trimming unwanted content
As we are implementing submodules, there might be unwanted files in the final codebase (`README.md`, `.github`, `.git`). This should not be an issue, as the binary build won't be affected by those files, but just in case you are bothered about a few extra files, there is a `Makefile` available to clear the unwanted files.