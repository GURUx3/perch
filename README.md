# Perch

> Push a branch, get a live URL for it. Delete the branch and it shuts down on its own.

Perch is a small, self-hosted preview environment manager for Git repositories. It gives every branch its own temporary live URL without using Docker or external hosting services. It is a from-scratch rebuild of the preview environment feature found in Vercel and Netlify, built by one person in a day as a learning project.

## The idea

When working on different branches, you often need to test each version of your application separately.

Usually, this means setting up environments manually or deploying each branch somewhere.

Perch handles this on one server.

- Push a branch, and Perch starts its application.
- Make a new commit, and Perch restarts the application.
- Delete a branch, and Perch shuts down its preview.
- Leave a preview running for too long, and Perch cleans it up automatically.

Each branch gets its own port and live URL.

## Example

Here is what a normal workflow looks like:

1. **Create a branch:** Create `feature/login` in your Git repository.
2. **Push it:** Push the branch to Git. Perch detects it during its next check.
3. **Get a preview:** Perch picks an available port, starts the application, and makes it available at `http://server:4001`.
4. **Make changes:** Push another commit. Perch detects the new commit and restarts the application on the same port.
5. **Delete the branch:** Delete `feature/login` from Git. Perch detects the deletion, stops the application, and frees the port.

No manual cleanup is needed.

## How it works

Perch has one main Python program called the **controller**.

The controller runs continuously and checks Git every 30 seconds.

```text
                  Git Repository
                        |
                        |
                  +-----------+
                  | Controller|
                  |  Python   |
                  +-----------+
                        |
              Checks every 30 seconds
                        |
              +---------+---------+
              |         |         |
           Branch A  Branch B  Branch C
              |         |         |
           Port 4001 Port 4002 Port 4003
              |         |         |
              +---------+---------+
                        |
                  Running Apps
                        |
                        v
                      User
                        |
                 http://server:8080
                   Index Page
                        |
                Links to live previews
```

### The reconcile loop

A reconcile loop means repeatedly checking the actual state against the expected state and fixing any differences.

Every 30 seconds, the controller does roughly this:

```text
LOOP forever:

    Get all Git branches and their latest commits.

    FOR each branch:

        IF branch is new:
            Pick a free port.
            Start its app as an OS subprocess.
            Save its branch, port, PID and commit in SQLite.

        ELSE IF its latest commit has changed:
            Stop the old process.
            Start the updated app on the same port.
            Update the saved commit.

        ELSE IF its app has crashed:
            Restart the app on the same port.

    FOR each tracked preview:

        IF its Git branch no longer exists:
            Kill its process.
            Free its port.

        ELSE IF its preview is older than 1 hour:
            Kill its process.
            Free its port.

    Update the preview index page.

    Wait 30 seconds.
```

SQLite stores the branch-to-port mapping, process ID (PID), and latest commit. This lets the controller recover its tracking state after a restart.

The index page runs on port `8080` and lists all live previews with their URLs.

## What Perch does not do

Perch is intentionally small. It has some clear limitations:

- **No Docker:** Every branch runs as a normal operating system process using Python's `subprocess`. This is fine for a personal demo, but processes are not isolated.
- **No subdomains:** Previews use ports instead. For example, `http://server:4001`.
- **No webhooks:** Perch polls Git every 30 seconds. Changes are not detected instantly.
- **No login or authentication:** Anyone who knows a preview's port can open it. This is for personal demos only, not public or production use.

## Run it

Install the dependencies:

```bash
pip install flask psutil
```

Start the controller:

```bash
python controller.py
```

Open the preview index page:

```text
http://localhost:8080
```

Make sure Git is available and the repository is configured before starting the controller.

## What you learn by building it

- **Git automation:** Reading branches, tracking commits, and detecting deletions.
- **Process management:** Starting, stopping, monitoring, and restarting OS processes.
- **State management:** Using SQLite to persist preview information.
- **Reconciliation:** Building a background loop that continuously checks and corrects system state.
