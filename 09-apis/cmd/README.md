# /cmd

Main applications for the project

The directory name for each application should match the name of the executable you want to have (e.g. `/cmd/myapp`).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you dont want other to reuse it, put that code in the `/internal` directory.

It's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.
