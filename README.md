The aim of this project is two-fold:

1. Build from scratch the basic functionality of a database. Heavily inspired by [link]()
2. Become more familiar with go, go testing, go standard library etc.

(Initial) components:

1. A way to store data on disk
-  Data stored in one file
-  A few columns, of different types (int, str, date)

2. A way to query that data on disk using a cli

(Eventual) components:

1. Optimising storage by adding indexing
2. Create an API layer
3. Measure the performance of the database


Notes:
- Anything in same package is accessible, no need to import
- No need to modify go mod or go init for importing a package, just import, and ensure it is upper case
- What should be the entrypoint of this application, should it be the cli main.go, or should the cli live separately? What is the basic (and easiest) folder structure to start with?


Useful links:

Project layouts:
- https://go.dev/doc/modules/layout
- https://github.com/golang-standards/project-layout 
- https://refactoring.guru/design-patterns (can I try and incoprorate some of these design patterns into the project)