# POC Blockchain GO (Alpha)

### Stack of project

- Gvm
- Golang 1.9
- Httpie (rest request)
- Linux

### Fire up project

For run this PoC, just make sure the `$GOPATH` is this directory.
Just run `. ./env.sh`.

Start the server. In the root folder, execute.
```sh
go run src/cmd/main.go --port=8000
```

Make some transaction on blockchain. First you need install the httpie, just execute the apt, `sudo apt install httpie`.

Send value on blockchain.
```sh
http post localhost:8000/transactions/new sender=thiago recipient=luis amount:=1 
```

Execute the PoW, like minning the block.
```sh
http get localhost:8000/mine
```

List the chain.
```sh
http get localhost:8000/chain 
```

The focus is make a study and improve the code base for a final objective.

---
The MIT License (MIT)

Copyright (c) 2017 RHIZOMPLATFORM

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

