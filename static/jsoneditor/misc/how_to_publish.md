# How to publish jsoneditor

This document describes the steps required to publish a new version of jsoneditor.


## Update version number

Update the version number in package.json.


## Update history

Update the date and version number in the file HISTORY.md. Verify whether all
changes in the new version are described.


## Test the library

Run the unit tests and validate whether all tests pass:

    npm test


## Build library

Build the build (jsoneditor.js, jsoneditor.css, ...) files by running:

    npm run build

After the build is complete, verify if the files are updated and contain the
correct date and version number in the header.


## Test

Test whether the npm library is ok by installing it locally:

    cd ../tmp-folder
    npm install ./path/to/jsoneditor

Check whether the examples in the library work ok, and whether the necessary
files are included.


## Commit

- Commit the final code.
- Merge the develop branch into the master branch.
- Push to github.

If everything is well, create a tag for the new version, like:

    git tag v1.2.4
    git push --tags


## Publish

Publish to npm:

    npm publish

Publish at cdnjs: test after 30 to 60 minutes whether the new version is
published at cdnjs (should auto update).


## Test published library

Install the libraries locally and test whether they work correctly:

    cd tmp-folder
    npm install jsoneditor
    bower install jsoneditor


## Done

Congrats, be proud.

