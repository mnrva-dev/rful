# Really F*cking Useful Logger

The sister package to RFSL. While RFSL is meant to be a plug and play solution to my specific problem,
RFUL is meant to be a simple yet heavily customizable logging package that will hopefully be a good
solution for many programs.

## Features

- Useful included time formats
- Useful included location context functions
- Easy to use
- Totally customizable
- Includes RotateWriter from RFSL

## Basic Usage

```go
func main() {

    // create a new Logger that outputs to Stdout
    log := rful.New(os.Stdout)

    log.Info("Hello from RFUL!")
    // 2022-11-16 19:33:48 INFO     -> Hello from RFUL!

    // change the date format
    log.SetTimeFormat(rful.TimeRFC1123)

    log.Info("New time format set!")
    // Wed, 16 Nov 2022 19:33:48 INFO     -> New time format set!

    // change the separator
    log.SetSeparator("| ")

    log.Info("New separator!")
    // Wed, 16 Nov 2022 19:44:06 INFO     | New separator!

    // lets add some context to our log
    // we can add the function information to every log
    log.SetLocation(rful.LocShortFunc)
    // and add a prefix to every entry
    log.SetPrefix("Main Logs: ")
    // we can also chain methods
    log.SetLocation(rful.LocShortFunc).SetPrefix("Main Logs: ")

    log.Info("More context!")
    // Main Logs: Wed, 16 Nov 2022 19:44:06 main:71 INFO     | More context!

    // to change all of these at once, we can use the
    // set preface method
    log.SetPreface("PREFIX: ", rful.TimeRuby, rful.LocShortFile, "--> ")

    log.Info("Preface has been changed!")
    // PREFIX: Wed Nov 16 19:44:06 2022 main.go:77 INFO     --> Preface has been changed!
}
```