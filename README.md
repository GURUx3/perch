# Perch

> Push a branch, get a live URL for it. Delete the branch and its
> container is torn down.

Perch is a self-hosted preview environment manager for Git repositories.
Every branch in your repo gets its own live, temporary URL, backed by
its own Docker container. It runs on a single cloud VM, uses the Docker
API to manage container lifecycles, and survives reboots as a systemd
service. It is a from-scratch rebuild of the preview environment feature
in Vercel and Netlify --- controller, container orchestration, cloud
deploy, and all --- built by one person in a day as a learning
experiment.

## The idea

Working across branches means you constantly need to _see_ each version
running. Usually that means manual deploys or one shared staging box.
Perch does it automatically, on one server, with real isolation:

- Push a branch → Perch builds a container from it and starts it on a
  free port.
- Push a new commit → Perch restarts the container with the new code.
- Delete the branch → Perch stops and removes the container.
- Leave a preview running past its TTL → Perch cleans it up on its
  own.

Each branch gets its own port and its own live URL. No shared state, no
manual cleanup, no Docker commands by hand.

## Example

1.  **Create a branch:** `feature/login` in your Git repository.
2.  **Push it.** Within 30 seconds, Perch detects it and starts a
    container: `perch-feature-login` running on port `4001`.
3.  **Open** `http://<vm-ip>:4001` --- you're looking at the live
    branch, isolated in its own container.
4.  **Push another commit.** Perch pulls the new code inside the
    container's mount and restarts it. Same port, new code.
5.  **Delete the branch.** Perch stops the container, removes it, and
    frees the port.
6.  **The index page** at `http://<vm-ip>:8080` lists every live preview
    with links, uptime, and a destroy button.

## Architecture
