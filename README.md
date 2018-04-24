# Five Go Gotchas I Learned the Hard Way - Demos

On April 25th, 2018, the second Go Open Source conference took place in Tel-Aviv.
At this conference I gave a talk about five Go gotchas. This repository was created
to store the code and files needed for demonstrating the different gotchas.

Content:

 - **demo.go** - A short Go sourcefile we run in a container with go to
demonstrate Go's terrible dependency management.
 - **Dockerfile.go_container** - A Dockerfile for creating the image we use for
demonstrating Go's terrible dependency management
 - **Dockerfile.jupyter** - A Dockerfile for creating a Jupyter image with
notebooks for demonstrating the code related gotchas.
 - **notebooks** - A directory with notebooks used to demonstrate code related
gotchas.