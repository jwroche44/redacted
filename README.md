# Redacted
Built using Go Version 1.16
## Building
`go build .`

## Running
### Example for testing
```bash
$ ./redacted example, thingy "another phrase" "I created an ExaMple thingy with another Phrase"

I created an XXXX XXXX with XXXX
```

### Running with the contents of a file
```bash
$ ./redacted example, thingy "another phrase" "`cat example_file.txt`"
```

## Notes for Improvements
- Could implement Boyer-Moore for better efficiency
- Make unit tests for `SanitizeFile()` and `Sanitize()`
- Make it more configurable
    - JSON config?
- Statistics
    - How many instances of each key were replaced?
    - Store stats to a DB?
        - What files had what keys removed
- Storing the final file somewhere searchable
    - Elastic Search
    - Mongo
    - As a blob in a RDB