## dotenv

Exports the .env to enviroment variables.

```
APP_USER=DANILO
APP_KEY=thesecretkey
```

will load the `APP_USER` key to the enviroment variables that in go can be obtained using the os
package `os.Getenv("APP_USER")` and should return "DANILO".
