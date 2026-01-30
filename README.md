The aim of this project is two-fold:

1. Build from scratch the basic functionality of a database. Heavily inspired by [link](https://github.com/codecrafters-io/build-your-own-x)
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
4. Creating a front end UI 

Notes:
- Anything in same package is accessible, no need to import
- No need to modify go mod or go init for importing a package, just import, and ensure it is upper case
- What should be the entrypoint of this application, should it be the cli main.go, or should the cli live separately? What is the basic (and easiest) folder structure to start with?

Folder structure:

```
go-database/
├── cli/
│   ├── cmd/
│   │   ├── retrieve.go
│   │   ├── root.go
│   │   ├── version.go
│   │   └── write.go
│   ├── data/
│   │   └── text.txt
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── Makefile
├── ui/
│   └── README.md
├── .gitignore
└── README.md
```

Using the CLI:

Navigate to the cli folder:

```
cd .\cli
```

Specify the command you want to use:

```
go run main.go [verion][write][retrieve]
```

Importing the go package:

```
go install
```

```
database version
database write
...
```
Useful links:

Project layouts:
- https://go.dev/doc/modules/layout
- https://github.com/golang-standards/project-layout 
- https://refactoring.guru/design-patterns (can I try and incoprorate some of these design patterns into the project)


How do databases work?
- https://cstack.github.io/db_tutorial/ 
- https://build-your-own.org/database/
