# go-testcontainers-sample

[Testcontainers](https://testcontainers.com/) is an open source framework for providing throwaway, lightweight instances
of databases, message brokers,
web browsers, or just about anything that can run in a Docker container.

### How it works

Test dependencies as code
No more need for mocks or complicated environment configurations. Define your test dependencies as code, then simply run
your tests and containers will be created and then deleted.

With support for many languages and testing frameworks, all you need is Docker.

### Things todo list

1. Clone this repository: `git clone https://github.com/hendisantika/go-testcontainers-sample.git`
2. Navigate to the folder: `cd go-testcontainers-sample`
3. Run test: `go test -v ./...`
4. Check the console

```shell
hendisantika@Hendis-MacBook-Pro go-testcontainers-sample % go test -v ./...
?       go-testcontainers-sample/domain [no test files]
?       go-testcontainers-sample/infra/mem      [no test files]
?       go-testcontainers-sample/infra/mongo    [no test files]
2024/08/23 08:10:35 github.com/testcontainers/testcontainers-go - Connected to docker: 
  Server Version: 27.1.1 (via Testcontainers Desktop 1.15.1)
  API Version: 1.46
  Operating System: Docker Desktop
  Total Memory: 7837 MB
  Labels:
    com.docker.desktop.address=unix:///Users/hendisantika/Library/Containers/com.docker.docker/Data/docker-cli.sock
  Testcontainers for Go Version: v0.33.0
  Resolved Docker Host: tcp://127.0.0.1:49690
  Resolved Docker Socket Path: /var/run/docker.sock
  Test SessionID: 0c5fe099166c34cfc93f7dbd5abd1f6d0da59f0479b45353dac9afb7310a5f7d
  Test ProcessID: f13eab8f-1682-4104-b4a2-164bb7bbfabc
2024/08/23 08:10:35 🐳 Creating container for image testcontainers/ryuk:0.8.1
2024/08/23 08:10:35 ✅ Container created: ed3bf587b6cd
2024/08/23 08:10:35 🐳 Starting container: ed3bf587b6cd
2024/08/23 08:10:35 ✅ Container started: ed3bf587b6cd
2024/08/23 08:10:35 ⏳ Waiting for container id ed3bf587b6cd image: testcontainers/ryuk:0.8.1. Waiting for: &{Port:8080/tcp timeout:<nil> PollInterval:100ms skipInternalCheck:false}
2024/08/23 08:10:35 🔔 Container is ready: ed3bf587b6cd
2024/08/23 08:11:21 🐳 Creating container for image mongo:7.0.14-rc0-jammy
2024/08/23 08:11:21 ✅ Container created: 98dd34fc00ab
2024/08/23 08:11:21 🐳 Starting container: 98dd34fc00ab
2024/08/23 08:11:22 ✅ Container started: 98dd34fc00ab
2024/08/23 08:11:22 ⏳ Waiting for container id 98dd34fc00ab image: mongo:7.0.14-rc0-jammy. Waiting for: &{Port:27017 timeout:<nil> PollInterval:100ms skipInternalCheck:false}
2024/08/23 08:11:22 🔔 Container is ready: 98dd34fc00ab
2024/08/23 08:11:22 TestContainers: container mongo:7.0.14-rc0-jammy is now running at mongodb://127.0.0.1:53515
=== RUN   Test_Add
--- PASS: Test_Add (0.01s)
=== RUN   Test_Replace
--- PASS: Test_Replace (0.00s)
=== RUN   Test_By
--- PASS: Test_By (0.00s)
=== RUN   Test_All
--- PASS: Test_All (0.00s)
PASS
ok      go-testcontainers-sample/infra  (cached)

```