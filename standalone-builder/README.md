# STANDALONE-BUILDER

This container-util is used to build musl cross-compilers for linux alpine and then use them to cross-compile standalone executables.

# Create your standalone executable

## Run a container

To start a container just run:

```bash
    $ docker-compose run standalone-builder
```

If you want to clean all containers and images consider to run clean script (**scripts directory in this repository root**).

## Compile a cross-compiler

First you must configure your **config.mak** file inside **cross-compilers/musl-cross-make**. Provide a *TARGET* and *OUTPUT* variables. You can just use **config.mak.dist** as template and then rename it to **config.mak**.

***Note:*** It is advised to point *OUTPUT* variable somewhere inside **cross-compilers** as it is a mounted volume.

Check **cross-compilers/musl-cross-make** README for more details.

Go to **cross-compilers/musl-cross-make** and run:

```bash
    $ make
    $ make install
```

## Build executable

To build executable you must configure its build system (often make or automake), choose a good compiler, etc...

# Utils

## Submodules



To remove a submodule you need to:

- Delete the relevant section from the .gitmodules file.
- Stage the .gitmodules changes git add .gitmodules
- Delete the relevant section from .git/config.
- Run git rm --cached path_to_submodule (no trailing slash).
- Run rm -rf .git/modules/path_to_submodule (no trailing slash).
- Commit git commit -m "Removed submodule "
- Delete the now untracked submodule files rm -rf path_to_submodule

