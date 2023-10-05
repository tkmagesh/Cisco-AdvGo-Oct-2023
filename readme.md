## Advanced Golang ##
- Magesh Kuppan

## Schedule ##
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 3:00 PM (20 mins)
- Wind up       : 5:00 PM

## Methodology ##
- No powerpoint
- 100% Code driven class
- No dedicated Q&A time

## Repository ##
- https://github.com/tkmagesh/cisco-advgo-oct-2023

## Software Requirements ##
- Go Tools
- Visual Studio Code (Or any other editor)
- Go Extension (VS Code) - https://marketplace.visualstudio.com/items?itemName=golang.Go
- Docker Desktop



## Go Basics ##
- Language constructs
- Collection Types (Array, Slice, Map)
- Error handling 
- Deferred functions
- Higher Order Functions
- Panic & Recovery
- Interfaces
- Modules & Packages

- Structs
- Concurrency

## Struct ##
- Value Types
- Can be composed (No inheritence)

## Concurrency ##
- What is Concurrency?
    - Ability to have more than one execution path in an application
- Support for concurrency is offered as a language feature (not thorugh APIs)
    - go keyword, chan data type, chan operator (<-), range construct, select-case construct
- Extensive API support
    - sync package
    - sync/atomic package

- Concurrency Model
![image concurrency_model](./images/concurrency_model.png)

- Race Detection
    - go run --race <app>
    - go build --race <app>
    - Note : DO NOT create production builds with race detector

- Communication between goroutines (channels)
    - Channel is a "data type"
    - Declaration
    ```
        var <channel_variable> chan <data_type>
        ex: var ch chan int
    ```
    - Initialize
    ```
        <channel_variable> = make(chan <data_type>)
        ex: ch = make(chan int)
    ```
    - Channel Operations (using channle operator (<-) )
        - Send Operation
        ```
            <channel_variable> <- <value>
            ex: ch <- 100
        ```
        - Receive Operation
        ```
            <- <channel_variable>
            ex: <- ch
        ```
    - Channel Behavior
        - Receive Operation is ALWAYS a blocking operation (operation is blocked until the data becomes available in the channel)
        - Send Operation is blocked until a receive operation is initiated (conditional)
    ![image channel_behaviors](./images/channel_behaviors.png)
    - Buffered Channel
        - A "send" operation can be successful even if a "receive" operation is not initiated
        - Use case:
            "Receive" operation need to be optional