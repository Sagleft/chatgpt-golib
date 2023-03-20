## example 

```go
    package main

    import gptlib "github.com/Sagleft/chatgpt-golib"

    func main() {
        client := gptlib.NewChatGPT("YOUR-OPEN-AI-TOKEN")

        response, err := client.SendRequest(gptlib.RequestData{
            Prompt: "How many stars are there in our galaxy?",
        })
        if err != nil {
            log.Fatalln(err)
        }

        fmt.Println("response: ", response)
    }
```

result:

```
response: According to recent estimates, our Milky Way galaxy contains about 100-400 billion stars. However, this is only a rough estimate and the exact number may be much higher or lower than this range.
```
