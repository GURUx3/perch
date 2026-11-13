# **Perch — 15-hour learning and building roadmap**

Your plan is **9:00 AM to 12:00 AM tomorrow (November 14)**.

That's a 15-hour project window. I would divide it into:

- **First 3 hours:** Understand the complete technology stack and architecture before writing code.
- **Next 10 hours:** Build Perch, learning specific concepts as you need them.
- **Final 2 hours:** Deploy, test, fix issues, and prepare the demo.

One important distinction: You do not need to master Go, Docker, cloud infrastructure, or DevOps in three hours. You need enough understanding to know what each component does, how they communicate, and what you need to implement.

The goal is to learn the concepts before coding, then learn the implementation details while building.

**Your time allocation**

# **3 hours**

Dedicated learning

3h Learning

10h Building

2h Final

| Time          | Activity           |
| ------------- | ------------------ |
| 9 AM – 12 PM  | Learn              |
| 12 PM – 10 PM | Build              |
| 10 PM – 12 AM | Deploy, test, demo |

## **Part 1 — The 3-hour prerequisite learning window**

9:00 AM – 12:00 PM

**Learn the architecture before opening your project folder.**

I would divide these three hours into six learning sessions.

### **9:00 – 9:35 AM | Go fundamentals for Perch**

[10 Reasons why developers love Go | TechGig](https://images.openai.com/static-rsc-4/fTkMvSADcL9IF1X2X9se1tQlCZ7cqWkJAa3_LMWov3MhQJdbVZixjE93JEipPxKRLofY6cDnKRCvoqRHhBnm7XtCoXl3xGxuDwEtzloXmLTbZ2Fgya56ecJarGjQek_hh_eljqZVRhidNZ4rthkJhQj91TcPcjUeSDfsl5PqOPo?purpose=inline)

**Goal:** Understand enough Go to write a long-running infrastructure application.

35 minutes · No need to complete the entire Go course.

Learn these concepts:

| Concept           | Why Perch needs it                          |
| ----------------- | ------------------------------------------- |
| Structs           | Represent a Git branch and its preview      |
| Functions         | Separate controller operations              |
| Interfaces        | Understand the Docker client abstraction    |
| Slices and maps   | Store branches and preview information      |
| Error handling    | Handle failed Git and Docker operations     |
| Goroutines        | Run concurrent tasks when needed            |
| `time.Ticker`     | Execute the reconcile loop every 30 seconds |
| `context.Context` | Control application shutdown and timeouts   |
| `os/exec`         | Run Git commands from Go                    |
| `net/http`        | Serve the preview index and health endpoint |

**Don't study:** Advanced generics, reflection, memory management, advanced concurrency patterns.

Resources:

- [A Tour of Go](https://go.dev/tour/?utm_source=chatgpt.com)
- [Effective Go](https://go.dev/doc/effective_go?utm_source=chatgpt.com)
- [Go os/exec documentation](https://pkg.go.dev/os/exec?utm_source=chatgpt.com)

Checkpoint

You should be able to explain how a Go program runs continuously, executes a command, handles errors, and waits for 30 seconds before repeating an operation.

### **9:35 – 10:15 AM | Docker fundamentals**

[Docker Logo, Icon, and Brand Guidelines](https://images.openai.com/static-rsc-4/e3JNUoCC2I1xaC-d5nwDwFPxo03f4SC-A_dcg0VcYGPSlNBluApjn3F1w4G1SUAHOur3UhwI9-36_T4QL4iPUdCfxf3fErvrYWd7MORRwQKvSETwAdnzA2E4p_0itH6LQDP-73tY3Ko171ZdEUh6VyHwMDIp5eVSVVV5YiV6kfo?purpose=inline)

**Goal:** Understand how Perch creates and manages isolated application environments.

40 minutes · This is one of the most important topics.

Learn these concepts in order:

1. What is Docker and why do we need it?
2. Docker architecture: Client, Docker CLI, Docker daemon.
3. Images vs containers.
4. Container lifecycle: Create, Start, Stop, Restart, Remove.
5. Port mapping: Host port vs container port.
6. Bind mounts: Sharing a host directory with a container.
7. Container logs and exit status.
8. Docker labels and how they identify Perch-managed containers.
9. Docker socket and how applications communicate with Docker.

**Understand this example:**

```
Your Git branch
      |
      v
  Host directory
      |
      | Bind mount
      v
 Docker container
      |
      | Application listens on 5000
      v
 Host port 4001
      |
      v
 http://server:4001
```

Resources:

- [Docker Get Started](https://docs.docker.com/get-started/?utm_source=chatgpt.com)
- [Docker containers](https://docs.docker.com/engine/containers/?utm_source=chatgpt.com)
- [Docker bind mounts](https://docs.docker.com/engine/storage/bind-mounts/?utm_source=chatgpt.com)
- [Docker Engine SDKs](https://docs.docker.com/reference/api/engine/sdk/?utm_source=chatgpt.com)

**Important:** Perch uses Docker programmatically through the Go SDK. You aren't just writing shell commands to run Docker.

Checkpoint

You should understand how Go can ask Docker to create a container, mount a directory, map a port, and restart a container.

### **10:15 – 10:25 AM | Break**

### **10:25 – 10:55 AM | Git automation and reconciliation**

[Git for Windows](https://images.openai.com/static-rsc-4/_6TzGCgiB44vGMLhF3m3Dp0SmJ3Edeb8ixErVPkEi48UurMr1pHVUGfA6qXJAgbcu_majlWdT-NuK8Ohx11syWv7iYlVB-fB_xQLqjoepATYJDPcu7utQIlBOLfMuNuWZTXp0XWYsmb1NbGQsRjKoR4CticsogvOZ8YlopInXIw?purpose=inline)

**Goal:** Understand how Perch discovers branches and decides what to do.

30 minutes · This is the heart of Perch.

Learn:

- Git repository, remote and local branches.
- What `git fetch` does.
- How to list remote branches.
- What a commit SHA is.
- How to detect a new commit.
- How to detect a deleted remote branch.
- What a polling mechanism is.
- What a reconciliation loop is.
- Desired state vs actual state.
- What happens when a container crashes.

Understand this table:

| Git state       | Docker state                  | Perch action       |
| --------------- | ----------------------------- | ------------------ |
| New branch      | No container                  | Create container   |
| Existing branch | Running container, old commit | Update and restart |
| Existing branch | Stopped container             | Start container    |
| Deleted branch  | Running container             | Stop and remove    |
| Existing branch | Running, TTL expired          | Stop and remove    |

Resources:

- [Git fetch documentation](https://git-scm.com/docs/git-fetch?utm_source=chatgpt.com)
- [Git remote branch discovery](https://git-scm.com/docs/git-ls-remote?utm_source=chatgpt.com)
- [Kubernetes controllers and reconciliation](https://kubernetes.io/docs/concepts/architecture/controller/?utm_source=chatgpt.com)

**Checkpoint**

You should be able to explain why Perch compares Git's state with Docker's state every 30 seconds instead of blindly starting containers.

### **10:55 – 11:25 AM | SQLite and persistence**

[Download SQLite Logo Vector Free | Logowik](https://images.openai.com/static-rsc-4/FNgt5tS7aCAtljy9lscSrsexgjF4pglv8m76CjorlrsvF2QY74RB5n0OKcfVEc1hSm_hKwVgeZTv9aAwnU8uOypGl7nn_6IAXTZYp5kaJUNWnm0ltOheJsZDp4_iDbuIvMAHwcNOppZ2cukLawB8fOm_UOB4oTdeGcgB0h6CsHc?purpose=inline)

**Goal:** Understand how Perch remembers previews across restarts.

30 minutes · Basic database operations only.

Learn:

- What SQLite is.
- Tables, rows and columns.
- Primary keys.
- CRUD operations.
- SQL `INSERT`, `SELECT`, `UPDATE`, `DELETE`.
- Go `database/sql`.
- Why state must survive controller restarts.
- What happens when SQLite and Docker disagree.

Think of the database as a notebook:

| Branch            | Port | Container ID | Commit SHA |
| ----------------- | ---- | ------------ | ---------- |
| feature-login     | 4001 | abc123       | 9a81f2     |
| feature-dashboard | 4002 | def456       | 7b42c1     |

Resources:

- [Go database/sql](https://go.dev/doc/database/?utm_source=chatgpt.com)
- [SQLite Go driver documentation](https://modernc.org/sqlite/?utm_source=chatgpt.com)
- [SQLite quickstart](https://www.sqlite.org/quickstart.html?utm_source=chatgpt.com)

**Checkpoint**

You should understand why Perch needs SQLite instead of keeping all preview information in a Go map.

### **11:25 AM – 12:00 PM | Cloud, Linux and systemd**

[How to setup a Jupyter server with EC2 instance | by Nandithags | Medium](https://images.openai.com/static-rsc-4/dvc1qLjhb8ldRMeUbcVmgpwngYzHzIeFW5F_HHa102a3Hn4GSdR-ZDnqy-8Q_Mwdk3Zgt_PUllvB2nXVg11IqCMenxY4mCJuKczznujSzbvrU2Uwz3cLJQwcUETwrUgIDLNS2xCZ4ypA-h6wddZKaWDxh421OInAfSaGFzuMtUM?purpose=inline)

**Goal:** Understand how your application becomes a service running on a Linux VM.

35 minutes · Operational basics.

Learn these:

| Topic                 | What you need to understand            |
| --------------------- | -------------------------------------- |
| Linux VM              | A remote Linux machine running Perch   |
| SSH                   | Connecting to your VM                  |
| SCP                   | Copying the Go binary to the VM        |
| Firewall              | Allowing specific inbound ports        |
| systemd               | Managing Perch as a background service |
| journalctl            | Reading application logs               |
| Health endpoint       | Checking if Perch is functioning       |
| Environment variables | Passing configuration to Perch         |

Resources:

- [Ubuntu command line tutorial](https://ubuntu.com/tutorials/command-line-for-beginners?utm_source=chatgpt.com)
- [systemd service documentation](https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html?utm_source=chatgpt.com)
- [journalctl documentation](https://www.freedesktop.org/software/systemd/man/latest/journalctl.html?utm_source=chatgpt.com)

**Checkpoint**

You should understand how to copy a Go binary to a VM, start it using systemd, inspect logs and restart it.

## **Part 2 — What you should understand before coding**

At 12 PM, you should be able to explain this entire architecture without looking at the README.

[Dubernetes: Vibe Coding A Dumb Container Orchestrator](https://images.openai.com/static-rsc-4/zBhwu9kfMu856qn7o2JJFzgK2ERCl_H891oS-X36JDtg_4XMS7gvEabg3wlKasnH_GeLB-4TeFV-2pME7VQNerQOQQ26ajsve8sg0eqydq_HacvigQYVsxlYoEK3MJ_R-EFdvhWU5lD8uwt0ec_EHQgg3kWnTvsW4c4bkNuTAFM?purpose=inline)

[blog.esc.sh](https://blog.esc.sh/vibecoding-dubernetes/?utm_source=chatgpt.com)

### **Perch architecture in plain English**

There are six main pieces:

| Component          | Responsibility                                                |
| ------------------ | ------------------------------------------------------------- |
| Go controller      | Coordinates everything                                        |
| Git                | Tells Perch which branches exist and which commits are latest |
| Docker             | Runs each branch in its own container                         |
| SQLite             | Remembers which branch uses which container and port          |
| HTTP server        | Displays previews and exposes health status                   |
| Linux VM + systemd | Keeps Perch running on a remote server                        |

The lifecycle looks like this:

**1. Developer pushes a Git branch**

feature/login · commit SHA

**2. Go controller discovers the branch**

Every 30 seconds

**3. Controller communicates with Docker**

Create · Start · Restart · Stop · Remove

**4. Docker runs the application**

Container port 5000 → host port 4001

**5. SQLite stores preview state**

Branch · port · container ID · commit SHA

**6. User opens the preview**

[http://server:4001](http://server:4001?utm_source=chatgpt.com)

### **Five concepts you must understand deeply**

These are more important than memorizing Go syntax.

**1. Reconciliation**

Perch continuously compares what Git says should exist with what Docker actually has running.

**2. Container lifecycle**

You must understand that creating a container and starting a container are different operations. The same applies to stopping and removing one.

**3. Persistent state**

Your controller can crash. Docker containers may still be running. SQLite and Docker labels help the controller recover.

**4. Port allocation**

Every active preview needs a unique host port. You must avoid assigning an occupied port or reusing a port before its container is removed.

**5. Idempotency**

Running the same reconciliation operation twice should not accidentally create duplicate containers. This is an important property of infrastructure controllers.

## **Part 3 — Your actual building roadmap**

12:00 PM – 9:00 PM

**Main implementation window**

The aim is to build one working feature at a time. Do not ask AI to generate the entire project. Use it to explain errors, clarify documentation, review individual functions and help debug.

| Time       | Build milestone       | What you should finish                                                        |
| ---------- | --------------------- | ----------------------------------------------------------------------------- |
| 12:00–1:00 | Project bootstrap     | Go module, Docker running, configuration and manual container test            |
| 1:00–2:00  | Controller foundation | Go HTTP server, `/health`, 30-second ticker                                   |
| 2:00–3:00  | Git integration       | Fetch repository, discover branches and read commit SHAs                      |
| 3:00–4:00  | Docker integration    | Connect Go to Docker, create/start a container, map a port and mount code     |
| 4:00–4:30  | Break                 | Food and rest                                                                 |
| 4:30–6:00  | Reconcile engine      | Handle new branches, changed commits, deleted branches and crashed containers |
| 6:00–7:00  | SQLite                | Persist state, recover records and implement TTL cleanup                      |
| 7:00–7:30  | Break                 | Dinner                                                                        |
| 7:30–8:30  | Preview dashboard     | Display live previews, ports and links                                        |
| 8:30–9:00  | Local integration     | Test the complete branch-to-preview lifecycle                                 |

### **Milestones in detail**

**Milestone 1 — Controller foundation**

Learn while building:

- Go HTTP server.
- `time.Ticker`.
- Graceful shutdown with `context`.
- Basic configuration.

Success: Perch starts, serves `localhost:8080/health`, and logs a message every 30 seconds.

**Milestone 2 — Git integration**

Learn while building:

- Calling Git through `os/exec`.
- Reading command output.
- Handling errors and exit codes.
- Working with branch names and commit SHAs.

Success: Your controller can discover remote branches and their latest commits.

**Milestone 3 — Docker integration**

Learn while building:

- Docker Go SDK.
- Docker container configuration.
- Bind mounts.
- Port bindings.
- Container labels.
- Container inspection.

Success: Your Go application creates and starts a container for one test branch.

**Milestone 4 — Reconciliation engine**

This is your biggest milestone.

Implement these operations individually:

1. Detect a new branch.
2. Create its working directory.
3. Start its container.
4. Detect a changed commit.
5. Restart the preview with updated code.
6. Detect crashed containers.
7. Remove previews for deleted branches.
8. Remove expired previews.

Success: You can push, update and delete a branch without manually managing its container.

**Milestone 5 — SQLite**

Store enough information to recover your previews after the controller restarts.

Success: Restart the controller and verify that it can rediscover existing Perch containers and their ports.

**Milestone 6 — Dashboard**

Build a basic HTML page. Don't waste time designing a beautiful dashboard.

Show:

- Branch name.
- Commit SHA.
- Preview URL.
- Container status.
- Creation time or expiry time.

Success: Open `localhost:8080` and see your running previews.

## **Part 4 — Cloud deployment and final testing**

9:00 PM – 12:00 AM

This is where Perch becomes more than a local application.

| Time        | Task                                           | Expected result                                         |
| ----------- | ---------------------------------------------- | ------------------------------------------------------- |
| 9:00–9:30   | Provision Linux VM                             | SSH connection works                                    |
| 9:30–10:00  | Install Docker and configure firewall          | Docker works on remote VM                               |
| 10:00–10:30 | Cross-compile Go binary and transfer using SCP | Perch binary runs on Linux                              |
| 10:30–11:00 | Configure systemd                              | Perch starts automatically as a service                 |
| 11:00–11:30 | Integration and recovery tests                 | Branch push, update, delete and controller restart work |
| 11:30–12:00 | Fix remaining issues and record demo           | Working demonstration and README                        |

Learn these during implementation:

- [Docker on Ubuntu](https://docs.docker.com/engine/install/ubuntu/?utm_source=chatgpt.com)
- [systemd service units](https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html?utm_source=chatgpt.com)
- [Building Go applications](https://go.dev/doc/tutorial/compile-install?utm_source=chatgpt.com)

**Deployment success criteria:**

- Perch runs on your Linux VM.
- systemd starts Perch.
- The index page is accessible.
- A pushed branch gets a preview.
- A new commit updates the preview.
- A deleted branch gets cleaned up.
- Restarting Perch does not lose track of existing containers.

## **Part 5 — Important architectural decisions to understand**

Before you start implementing, there are a few things in your revised README that need careful attention.

These are not additional technologies to learn. They are implementation decisions.

### **1. How will Perch run different applications?**

A Docker container needs an image, dependencies and a startup command.

For example:

```
Branch: feature/login
       |
       v
Host directory:
/var/perch/repos/feature-login
       |
       v
Docker image:
node:22-slim
       |
       v
Bind mount:
/app
       |
       v
Install dependencies
       |
       v
npm run start
       |
       v
Container port 5000
```

A base runtime image does not automatically contain your application's dependencies.

For the first version, support one known application type. For example, a simple Node.js application. Add other runtimes later.

### **2. How will Git branches become separate directories?**

You need a working directory for each branch.

Understand how to:

- Clone the repository.
- Fetch remote changes.
- Check out a specific branch or commit.
- Keep branches isolated in separate directories.
- Remove the directory when the preview expires.

### **3. What happens when Docker and SQLite disagree?**

Imagine:

- SQLite says a container exists.
- Docker says the container does not exist.

Or the reverse.

Your controller must inspect Docker's actual state and reconcile it with the stored state. Don't blindly trust the database.

### **4. How will TTL work?**

For the MVP, define TTL as one hour from preview creation.

Store a `created_at` timestamp and calculate expiry during reconciliation.

You can introduce last-access-based expiration later.

### **5. Security**

Your README explicitly says there is no authentication.

The Docker socket also grants substantial control over the host. Never expose it publicly. Restrict preview ports using firewall rules, preferably to your own IP during the demo.

## **Part 6 — Your three-hour learning checklist**

Use this as your actual checklist tomorrow. Don't move to the next topic just because you finished watching a video.

**Learning readiness**

# **0/23**

**1. Go fundamentals**

I understand structs, functions, slices and maps.

I understand error handling in Go.

I understand goroutines and context at a basic level.

I understand time.Ticker and os/exec.

I understand how net/http serves an endpoint.

**2. Docker**

I understand images versus containers.

I understand the Docker daemon and SDK.

I understand container lifecycle operations.

I understand bind mounts and port mapping.

I understand labels and container inspection.

**3. Git and reconciliation**

I understand remote branches and commit SHAs.

I understand fetching and checking out branches.

I understand polling and reconciliation.

I understand desired state versus actual state.

I understand branch deletion and crash detection.

**4. SQLite**

I understand tables and CRUD operations.

I understand Go database/sql.

I understand why state must persist.

I understand how Docker and SQLite state can disagree.

**5. Cloud and DevOps**

I understand SSH and SCP.

I understand Linux VM and firewall basics.

I understand systemd and journalctl.

I understand how to build a Linux Go binary.

Reset checklist

## **Part 7 — How to learn while building**

This is the approach I strongly recommend.

Do not spend the next 15 hours asking AI to generate code and pasting it into your project.

Use this cycle:

**1. Understand the feature**

Read official documentation or an article.

**2. Understand the implementation**

Ask AI to explain the API and its important concepts.

**3. Implement it yourself**

Write a small piece of code and test it.

**4. Debug and verify**

Use AI to understand errors, not hide them.

**Repeat for the next feature**

## **My recommendation: What NOT to learn tomorrow**

Don't get distracted by these:

- Kubernetes.
- Terraform.
- Helm.
- Nginx and reverse proxy configuration.
- Advanced Docker networking.
- Advanced Go concurrency.
- CI/CD pipelines.
- Microservices.
- Prometheus and Grafana.

They can all be useful later. None is necessary for the first working version of Perch.

## **Final expectation**

Your target for November 14

By midnight, aim to have:

1. A Go controller that runs continuously.
2. Git branch discovery.
3. Docker container management through Go.
4. A working reconciliation loop.
5. SQLite persistence.
6. A simple preview dashboard.
7. A Linux VM deployment using systemd.

**One honest warning:** A complete, reliable implementation and cloud deployment in 15 hours is ambitious, particularly if you're learning the Docker SDK and Go infrastructure programming along the way. Prioritize the local branch-to-container lifecycle first. If deployment takes longer, don't sacrifice the working core just to claim cloud deployment.

The real achievement is understanding how a small infrastructure controller works and building one yourself. That's much more valuable than having an AI generate a large repository you cannot explain.
